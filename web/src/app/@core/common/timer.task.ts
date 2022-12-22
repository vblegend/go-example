/* eslint-disable @typescript-eslint/explicit-module-boundary-types */
import { NgZone } from "@angular/core";
import { TimerPoolService } from "@core/services/timer.pool.service";
import { Subject, Subscription } from "rxjs";
import { Delegate } from "./delegate";

/**
 * 定时器回调函数类型
 */
export declare type TimerTaskEventHandler = (task: TimerTask) => void;

/**
 * 表示一个定时器的状态
 */
export enum TimeState {
    /**
     * 正在运行中
     */
    Runing = 0,
    /**
     * 被挂起了
     */
    Suspend = 1,
    /**
     * 已取消/停止 无效的
     */
    Cancelled = 2
}

/**
 * 定时器的任务实例\
 * 此实例运行在ngZone之外
 */
export interface TimerTask {

    /**
     * 获取/设置 定时器间隔(秒)
     */
    get interval(): number;

    /**
     * 获取/设置 定时器间隔(秒)
     */
    set interval(value: number);

    /**
     * 定时器当前是否在执行回调函数
     */
    get isExecuteing(): boolean;

    /**
     * 定时器当前是否正在运行
     * 如果定时器被取消则返回false
     */
    get state(): TimeState;

    /**
     * 定时器任务运行计数器
     */
    get counter(): number;

    /**
     * 定时器最后一次任务的运行时间戳\
     * 当执行 @continue(true) 方法重置后此值为当前时间
     */
    get runOfLastTime(): number;

    /**
     * 挂起定时器，停止任务运行（不会停止当前正在执行的任务）
     * 可使用 @continue 方法继续执行定时器任务
     */
    suspend(): void;

    /**
    * 继续被挂起的定时器
    * @param reset 是否重置时间,重置后需继续等待 interval 秒,否则以上次执行时间继续  默认值不重置
    */
    continue(reset?: boolean): void;

    /**
     * 取消当前定时器\
     * 定时器取消后将销毁，无法继续使用
     */
    cancel(): void;

    /**
     * 订阅计时器取消的事件
     * @param next 
     */
    subscribeCancelEvent(next?: (value: TimerTask) => void): Subscription;

    /**
     * 在ngZone中运行callback回调函数 \
     * 使其代码被angular跟踪
     */
    run(callback: Delegate): void;
}




/**
 * 固定事件定时器任务
 */
export class FixedTimerTask implements TimerTask {
    private _counter: number;
    private _runOfLastTime: number;
    private _isExecing: boolean;
    private _state: TimeState;
    private _subscriber: Subject<TimerTask>;

    /**
     *
     */
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    public constructor(private _service: TimerPoolService, private _callback: TimerTaskEventHandler, private _thisContext: any, private _interval: number, private ngZone: NgZone) {
        this._counter = 0;
        this._runOfLastTime = new Date().getTime();
        this._state = TimeState.Runing;
        this._isExecing = false;
        this._subscriber = new Subject<TimerTask>();
    }

    public run(callback: Delegate): void {
        this.ngZone.run(callback);
    }

    public subscribeCancelEvent(next?: (value: TimerTask) => void): Subscription {
        return this._subscriber.subscribe(next);
    }

    public get state(): TimeState {
        return this._state;
    }

    public get isExecuteing(): boolean {
        return this._isExecing;
    }

    public get runOfLastTime(): number {
        return this._runOfLastTime;
    }

    public get counter(): number {
        return this._counter;
    }

    public get interval(): number {
        return this._interval;
    }
    public set interval(value: number) {
        this._interval = value;
    }
    /**
     * 临时中断任务
     */
    public suspend(): void {
        if (this._state == TimeState.Cancelled) throw new Error('该定时器已被取消，无法进行挂起操作');
        if (this._state == TimeState.Suspend) {
            console.warn('重复的定时器挂起操作');
        }
        this._state = TimeState.Suspend;
    }

    /**
     * 继续执行任务
     * @param reset 是否重置时间\
     * 重置后需继续等待 @interval 秒
     */
    public continue(reset?: boolean): void {
        if (this._state == TimeState.Cancelled) throw new Error('该定时器已被取消，无法进行挂起操作');
        if (this._state == TimeState.Runing) {
            console.warn('重复的定时器恢复操作');
        }
        if (reset) this._runOfLastTime = new Date().getTime();
        this._state = TimeState.Runing;
    }

    /**
     * 取消任务（取消后不能继续）
     */
    public cancel(): void {
        if (this._service) {
            this._state = TimeState.Cancelled;
            this._service.freeTimer(this);
            this._subscriber.next(this);
        }
    }

    public async execute(): Promise<void> {
        if (this._state != TimeState.Runing) return;
        if (this._isExecing) return;
        try {
            this._counter++;
            this._isExecing = true;
            await this._callback.apply(this._thisContext, [this]);
        } catch (e) {
            console.warn(e);
        } finally {
            this._isExecing = false;
            this._runOfLastTime = new Date().getTime();
        }
    }
}

