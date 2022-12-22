import { ErrorHandler, EventEmitter } from "@angular/core";
import { Exception } from "@core/common/exception";
import { NzNotificationService } from "ng-zorro-antd/notification";

export class GlobalErrorHandler implements ErrorHandler {
    // public em : EventEmitter;
    constructor(private notification: NzNotificationService) {

    }

    public handleError(error: Object): void {
        if (error instanceof Exception) {
            this._notification(error);
        }
        console.error(error);
    }

    private _notification(exception: Exception): void {
        this.notification.create('error', exception.host, exception.message);
    }


}