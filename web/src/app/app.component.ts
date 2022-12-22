import { Component, Injector } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';

@Component({
  selector: 'app-root',
  template: '<router-outlet></router-outlet>',
  styleUrls: ['./app.component.less']
})
export class AppComponent extends GenericComponent {

  constructor(injector: Injector) {
    super(injector);
  }





}
