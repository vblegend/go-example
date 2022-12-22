import { Component, Injector } from '@angular/core';
import { GenericComponent } from '../basic/generic.component';

@Component({
  selector: 'ngx-error',
  templateUrl: './error.component.html',
  styleUrls: ['./error.component.less']
})
export class ErrorComponent extends GenericComponent {

  constructor(injector: Injector) {
    super(injector);
  }

}
