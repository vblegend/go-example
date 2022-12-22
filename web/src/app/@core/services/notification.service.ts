import { Injectable } from '@angular/core';
import { NotificationType } from '@core/common/types';

import { NzNotificationRef, NzNotificationService } from 'ng-zorro-antd/notification';
@Injectable({
  providedIn: 'root'
})
export class NotificationService {

  constructor(private notification: NzNotificationService) {
  }



  public success(titleKey: string, contentKey: string): NzNotificationRef {
    // this.translate.instant
    return this.notification.success(titleKey, contentKey);
  }


  public error(titleKey: string, contentKey: string): NzNotificationRef {
    return this.notification.error(titleKey, contentKey);
  }



}
