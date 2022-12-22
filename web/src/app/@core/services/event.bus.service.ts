import { Injectable } from "@angular/core";
import { EventMessage } from "@core/common/event.message";
import { Observable, Subject, Subscription } from "rxjs";
import { map, filter } from 'rxjs/operators';

@Injectable({
    providedIn: 'root'
})
/**
 * 全局事件总线服务
 */
export class EventBusService {
    private _subscriber: Subject<EventMessage> = new Subject<EventMessage>();

    /**
     *
     */
    constructor() {
        this._subscriber = new Subject();
    }


    /**
     * 订阅事件总线
     * @param next 
     * @returns 
     */
    public subscribe(next?: (value: EventMessage) => void): Subscription {
        return this._subscriber.subscribe(next);
    }

    public filter(predicate: (value: EventMessage, index: number) => boolean): Observable<EventMessage> {
        return this._subscriber.pipe(filter(predicate, this));
    }


    /**
     * 把事件派遣给指定ID的对象 type: MessageTypes, sender: BasicWidgetComponent, receiver: string, data: EventMessageData
     * @param target 
     * @param type 
     * @param data 
     */
    public dispatch(message: EventMessage): void {
        this._subscriber.next(message);
    }

}