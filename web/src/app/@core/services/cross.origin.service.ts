import { CancelFunc } from "@core/common/delegate";
import { Exception } from "@core/common/exception";
import { Guid } from "@core/util/guid";



interface CrossOriginHeader {
    transId: string;
    type: 'request' | 'response';
}


interface CrossOriginRequest<T> extends CrossOriginHeader {
    body: CrossOriginMessage<T>;
}

interface CrossOriginResponse<T> extends CrossOriginHeader {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    success: boolean;
    data: T;
}

interface CrossOriginMessage<T> {
    action: string;
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    data: T | null;
}

interface TaskPromise<T> {
    timeId: number | null;
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    reject: (reason?: any) => void;
    resolve: (value: CrossOriginMessage<T>) => void;
}

export interface CallbackOptions<T> {
    /**
     * 该处理函数的执行上下文环境
     */
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    context: any;

    /**
     * 共享的 \
     * 此状态设为true后，在主窗口的事件处理完成之后会广播给所有子窗口 \
     * 一般用于set操作和状态变更操作来共享状态给所有iframe子窗口 \
     * 对于get操作，应默认为false \
     */
    shared: boolean;

    /**
     * 处理函数回调方法
     */
    callback: (action: string, data: T | null) => IActionResponse;

}

/**
 * 请求动作的响应结果
 */
export interface IActionResponse {
    /**
     * 响应状态
     */
    success: boolean;
    /**
     * 返回的数据
     */
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    data: any
}


/**
 * 跨域通讯服务
 */
export class CrossOriginService {

    private static messageHandle: (event: MessageEvent) => void;
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    private static tasks: Record<string, TaskPromise<any>> = {};
    private static windows?: Window[];

    private static isShutdown: boolean = false;
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    private static requestHandlers: Record<string, CallbackOptions<any>> = {};
    /**
     * 当前环境是否为顶级窗口
     */
    public static get isTopLevel(): boolean {
        return window.top == window.self;
    }


    /**
     * 初始化跨域信息交互环境
     * 宿主与IFrame的页面在启动app之前要确保执行此函数成功
     */
    public static async setup(): Promise<boolean> {
        if (this.windows != null) throw Exception.build('', '环境已初始化', -1);
        this.windows = [];
        this.messageHandle = this.onMessage.bind(this);
        window.addEventListener("message", this.messageHandle);
        if (!this.isTopLevel) {
            const result = await this.request("domain.register");
            return result == 'OK';
        } else {
            this.checkWindowActivityStatus();
            return true;
        }
    }


    /**
     * 销毁环境
     */
    public static shutdown(): void {
        window.removeEventListener("message", this.messageHandle);
        this.windows = [];
        this.isShutdown = true;
    }


    /**
     * 注册跨域请求处理器
     * @param action 请求动作
     * @param options 动作处理器相关选项
     * @returns 
     */
    // eslint-disable-next-line @typescript-eslint/explicit-module-boundary-types, @typescript-eslint/no-explicit-any
    public static registerRequestHandler(action: string, options: CallbackOptions<any>): CancelFunc {
        const handler = this.requestHandlers[action];
        if (handler == null) {
            this.requestHandlers[action] = options;
        }
        return () => {
            delete this.requestHandlers[action];
        }
    }



    /**
     * window 注册事件（仅在顶级窗口触发）
     * @param origin 
     */
    private static domainRegister(origin: Window) {
        // 注册 window
        if (this.windows!.indexOf(origin) == -1) {
            this.windows!.push(origin);
        }
    }
    /**
     * window 注销事件（仅在顶级窗口触发）
     * @param origin 
     */
    private static domainShutdown(origin: Window) {
        // 注册 window
        const index = this.windows!.indexOf(origin);
        if (index > -1) {
            this.windows!.splice(index, 1);
        }
    }



    /**
     * 清理已关闭的窗口
     */
    private static checkWindowActivityStatus() {
        for (let i = this.windows!.length - 1; i >= 0; i--) {
            if (this.windows![i].closed) {
                this.domainShutdown(this.windows![i]);
            }
        }
        if (!this.isShutdown) {
            window.setTimeout(this.checkWindowActivityStatus.bind(this), 500);
        }
    }




    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    private static parseCrossMessage(data: any): CrossOriginHeader | null {
        const coh = data as CrossOriginHeader;
        if (coh.type == 'request' || coh.type == 'response') {
            return data;
        }
        return null;
    }



    private static onMessage(event: MessageEvent): void {
        const origin = event.source as Window;
        const header = this.parseCrossMessage(event.data);
        if (header == null || origin == null) return;
        if (header.type == 'request') {
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            this.handleRequest(header as CrossOriginRequest<any>, origin);
        } else if (header.type == 'response') {
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            const response = header as CrossOriginResponse<any>;
            const task = this.tasks[response.transId];
            if (response.transId && task) {
                delete this.tasks[response.transId];
                window.clearTimeout(task.timeId!);
                if (response.success) {
                    task.resolve(response.data);
                } else {
                    task.reject(Exception.build('cross origin service result fail.', response.data, 800));
                }
            }
        }
    }



    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    private static handleRequest(request: CrossOriginRequest<any>, origin: Window): void {
        if (request.body.action == 'domain.register' && CrossOriginService.isTopLevel) {
            this.domainRegister(origin);
            CrossOriginService.response(request, origin, true, 'OK');
            return;
        }
        const handler = this.requestHandlers[request.body.action];
        if (handler) {
            let result: IActionResponse;
            try {
                result = handler.callback.apply(handler.context, [request.body.action, request.body.data]) as IActionResponse;
                // 处理函数指定了共享此状态且当前为顶级窗口时
                if (CrossOriginService.isTopLevel && handler.shared && result && result.success) {
                    CrossOriginService.broadcast(request.body);
                }
            } catch (ex) {
                result = { success: false, data: ex };
            }
            if (result != null) {
                CrossOriginService.response(request, origin, result.success, result.data);
            }
        }
    }

    /**
     * 发送一个结果消息给target窗口不等待目标反馈立即返回。
     * @param request 跨域请求消息
     * @param target 要返回的目标窗口对象
     * @param success 返回状态
     * @param data 返回数据
     */
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    public static response<T>(request: CrossOriginRequest<any>, target: Window, success: boolean, data: T): void {
        const reponse: CrossOriginResponse<T> = {
            transId: request ? request.transId : Guid.custom(8, 16),
            type: 'response',
            success: success,
            data: data
        }
        target.postMessage(reponse, '*');
    }

    /**
     * 发送一条动作请求消息给target窗口对象，并等待target窗口反馈返回结果。\
     * 返回失败时触发catch
     * @param action 跨域发送的请求动作
     * @param data 跨域发送的请求动作所携带的数据
     * @param target 发送目标窗口对象，默认为window.top 顶级窗口对象
     * @param timeout 超时时间(毫秒)，默认为60000毫秒，为null或为0时发送立即完成不等待数据返回
     * @returns 动作返回结果
     */
    // eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/explicit-module-boundary-types
    public static async request<T>(action: string, data: any | null = null, target: Window | null = null, timeout: number | null = 60000): Promise<T> {
        return new Promise((resolve, reject) => {
            if (target == null) target = window.top;
            if (window.top?.self == window && target == window) {
                console.warn(`[${'master -> master'}] cross-origin: ${action}, ${data}`)
            } else {
                console.warn(`[${this.isTopLevel ? 'master -> child' : 'child -> master'}] cross-origin: ${action}, ${data}`)
            }
            const request: CrossOriginRequest<T> = {
                transId: Guid.custom(8, 16),
                type: 'request',
                body: { action, data }
            }
            let timeId: number | null = null;
            if (timeout != null && timeout > 0) {
                timeId = window.setTimeout(() => {
                    const task = this.tasks[request.transId];
                    if (task) {
                        delete this.tasks[request.transId];
                        window.clearTimeout(task.timeId!);
                        if (target!.closed) {
                            task.reject(Exception.build("cross-origin", 'window is closed' + action, 500));
                        } else {
                            task.reject(Exception.build("cross-origin", 'wait data time out' + action, 500));
                        }
                    }
                }, timeout!);
                this.tasks[request.transId] = { resolve: resolve as never, reject, timeId };
                target!.postMessage(request, '*');
            } else {
                target!.postMessage(request, '*');
                // eslint-disable-next-line @typescript-eslint/no-explicit-any
                resolve(null as any);
            }
        });
    }

    /**
     * 作为顶级窗口，广播给所有子窗口
     * @param data 
     */
    private static broadcast<T>(data: CrossOriginMessage<T>): void {
        if (this.isTopLevel) {
            const windows = this.windows!.slice();
            for (const fwind of windows) {
                if (!fwind.closed) {
                    this.request(data.action, data.data, fwind);
                }
            }
        }
    }


}