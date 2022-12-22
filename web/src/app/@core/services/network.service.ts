/* eslint-disable @typescript-eslint/no-explicit-any */
import { Injectable } from '@angular/core';
import { Exception } from '../common/exception';


export interface WebSocketTask {
    tickcount: number;
    timeout: number;
    complate: boolean;
    timer: number;
    resolve: (data: any) => void;
    rejects: (data: any) => void;
}
export interface WebSocketMessage<T> {
    sn: number;
    method: string;
    data: T;
}


@Injectable({
    providedIn: 'root'
})
export class NetWorkService {
    private _connectPromise: Promise<boolean> | null = null;
    private webSocket!: WebSocket | null;
    private serialNumber: number;
    private tasklist: Map<number, WebSocketTask>;
    public timeout = 120000;
    private _url!: string;

    constructor() {
        this.serialNumber = 0;
        this.tasklist = new Map();
        // console.log('NetWorkService');
    }

    private getSerialNumber(): number {
        this.serialNumber++;
        return this.serialNumber;
    }

    public get url(): string {
        return this._url;
    }

    public set url(value: string) {
        if (this._connectPromise) throw Exception.build('network service', 'websocket url must be modified when the connection request is not established.');
        this._url = value;
    }

    /**
     * connect to server
     * @param url 
     * @returns 
     */
    public async connection(): Promise<boolean> {
        if (this._connectPromise) return this._connectPromise;
        this._connectPromise = this.connectionAsync(this._url);
        return this._connectPromise;
    }



    private async connectionAsync(url: string): Promise<boolean> {
        if (this.webSocket) {
            if (this.webSocket.readyState === WebSocket.OPEN) return Promise.resolve(true);
            if (this.webSocket.readyState === WebSocket.CONNECTING) {
                this.close();
            }
            this.webSocket = null;
        }
        return new Promise((resolve, reject) => {
            this.webSocket = new WebSocket(url);
            this.webSocket.onopen = (_ev: Event) => {
                resolve(true);
                this.socket_opend(_ev);
            };
            this.webSocket.onerror = (_ev: Event) => {
                reject(_ev);
                this.socket_error(_ev);
                this._connectPromise = null;
            };
            this.webSocket.onclose = this.socket_closed.bind(this);
            this.webSocket.onmessage = this.socket_message.bind(this);
        });
    }







    /**
     * 
     * @param method 
     * @param data 
     * @param timeout  default = 120000
     * @returns 
     */
    public async send<TData, TResult>(method: string, data: TData, timeout?: number): Promise<TResult> {
        let _timeout = this.timeout;
        if (timeout != null) _timeout = this.timeout;
        const promise = new Promise<TResult>((resolve, rejects) => {
            const sn = this.getSerialNumber();
            if (this.webSocket == null) return rejects(Exception.build('network service', 'websocket is not initialized!'));
            if (this.webSocket.readyState != WebSocket.OPEN) return rejects(Exception.build('network service', 'websocket is not connected!'));
            const message: WebSocketMessage<TData> = { sn, method, data };
            this.webSocket.send(JSON.stringify(message));
            // console.log(sn);
            // check timeout
            const timer = window.setTimeout(() => {
                if (!task.complate) {
                    task.rejects(Exception.build('network service', 'websocket call timeout!'));
                    task.complate = true;
                    this.tasklist.delete(sn);
                }
            }, timeout);
            const task: WebSocketTask = { tickcount: this.tickCount, resolve, rejects, timeout: _timeout, complate: false, timer: timer };
            this.tasklist.set(sn, task);
        });
        return promise;
    }



    private socket_message(ev: MessageEvent): void {
        try {
            if (typeof ev.data != 'string') return;
            const message: WebSocketMessage<any> = JSON.parse(ev.data);
            if (message.sn == null) return;
            const task = this.tasklist.get(message.sn);
            if (task) {
                window.clearTimeout(task.timer);
                task.complate = true;
                this.tasklist.delete(message.sn);
                task.resolve(message.data);

            }
        } catch (e) {
            console.error(e);
        }
    }

    private get tickCount(): number {
        return (typeof performance === 'undefined' ? Date : performance).now();
    }





    public get isConnect(): boolean {
        return this.webSocket! && this.webSocket!.readyState === WebSocket.OPEN;
    }



    private close(): void {
        if (this.webSocket) {
            if (this.webSocket.readyState === WebSocket.OPEN || this.webSocket.readyState === WebSocket.CONNECTING) {
                this.webSocket.close();
            }
            this.webSocket = null;
        }
    }




    private socket_opend(_ev: Event): void {

    }

    private socket_closed(ev: CloseEvent): void {
        const code = ev.code;
        const reason = ev.reason;
        const wasClean = ev.wasClean;

    }

    private socket_error(_ev: Event): void {
        throw Exception.fromCatch('NetService websocket', _ev, 'Unable to connect to Websocket server !');
    }




}
