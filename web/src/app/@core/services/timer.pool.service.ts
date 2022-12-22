import { Injectable, NgZone, OnDestroy } from "@angular/core";
import { FixedTimer } from "@core/common/fixedtimer";
import { FixedTimerTask, TimerTask, TimerTaskEventHandler, TimeState } from "@core/common/timer.task";


/**
 * 定时器池服务\
 * 用于统一管理定时器机制
 */
@Injectable()
export class TimerPoolService implements OnDestroy {
    private readonly defaultInterial: number = 1000;
    private _tasks: FixedTimerTask[];
    private fTimer: FixedTimer;
    /**
     * canvas: ViewCanvasComponent
     */
    constructor(private ngZone: NgZone) {
        this._tasks = [];
        // console.log('new TimerPoolService()');
        this.fTimer = new FixedTimer(this.timer_doWork, undefined, undefined, this);
        this.ngZone.runOutsideAngular(() => {
            this.fTimer.start(this.defaultInterial);
        });
    }


    private timer_doWork() {
        const tasks = this._tasks.slice();
        const tickcount = new Date().getTime();
        for (const task of tasks) {
            if (!task.isExecuteing && tickcount - task.runOfLastTime >= task.interval * 1000) {
                // 异步调用任务回调函数
                task.execute();
            }
        }
        tasks.length = 0;
        // console.log('=======================================================================================');
        this.ngZone.runOutsideAngular(() => {
            this.fTimer.restart(this.defaultInterial);
        })
    }



    /**
     * **分配一个定时器任务** \
     * 并返回一个可被管理的定时器任务对象
     * @param callback 回调函数
     * @param _this 回调上下文对象
     * @param interval 间隔 毫秒
     * @returns 任务对象
     */
    public allocTimer(callback: TimerTaskEventHandler, _this: Object, interval: number): TimerTask {
        const task = new FixedTimerTask(this, callback, _this, interval, this.ngZone);
        this._tasks.push(task);
        return task;
    }




    /**
     * 释放一个Timer任务
     * @param task 
     */
    public freeTimer(task: TimerTask): void {
        if (task.state != TimeState.Cancelled) {
            task.cancel();
        }
        const index = this._tasks.indexOf(<FixedTimerTask>task);
        if (index > -1) {
            this._tasks.splice(index, 1);
        }
    }





    public ngOnDestroy(): void {
        this.fTimer.stop();
        const tasks = this._tasks.slice();
        for (const task of tasks) {
            task.cancel();
        }
        this._tasks.length = 0;
        // console.log('ngOnDestroy TimerPoolService()');
    }
}