


export declare type FixedTimerHandler = (timer: FixedTimer) => void;

export class FixedTimer {

    private _interval: number;
    private _callback: Function;
    private _thisContext?: Object;
    private _timeId?: number;
    private _thisfunction: Function;

    private _onStart?: Function;
    private _onStop?: Function;



    /**
     *
     */
    constructor(callback: FixedTimerHandler, onStart?: FixedTimerHandler, onStop?: FixedTimerHandler, thisContext?: Object) {
        this._callback = callback;
        this._onStart = onStart;
        this._onStop = onStop;
        this._interval = 1000;
        this._thisContext = thisContext;
        this._timeId = undefined;
        this._thisfunction = this.time_callback.bind(this);
    }

    public get interval(): number {
        return this._interval;
    }

    public set interval(value: number) {
        this._interval = value;
    }

    private time_callback() {
        this.emitStopEvent();
        this._timeId = undefined;
        this._callback.apply(this._thisContext, [this]);
    }


    public get isRuning(): boolean {
        return this._timeId != null;
    }

    public restart(_interval?: number): void {
        if (_interval) {
            this._interval = _interval;
        }
        this.stop();
        this.start();
    }

    public start(_interval?: number): void {
        if (_interval) {
            this._interval = _interval;
        }
        if (this._timeId == null) {
            this.emitStartEvent();
            this._timeId = window.setTimeout(this._thisfunction, this._interval | 0);
        }
    }

    public startAs(): void {
        if (this._timeId == null) {
            this.emitStartEvent();
            this._timeId = window.setTimeout(this._thisfunction, this._interval);
        }
    }



    public stop(): void {
        if (this._timeId) {
            window.clearTimeout(this._timeId);
            this.emitStopEvent();
            this._timeId = undefined;
        }
    }

    private emitStartEvent() {
        if (this._onStart) this._onStart.apply(this._thisContext, [this]);
    }

    private emitStopEvent() {
        if (this._onStop) this._onStop.apply(this._thisContext, [this]);
    }


}