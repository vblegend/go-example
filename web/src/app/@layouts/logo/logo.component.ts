import { Component, Injector, } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';
import { MenuService } from '@core/services/menu.service';

@Component({
  selector: 'app-logo',
  templateUrl: './logo.component.html',
  styleUrls: ['./logo.component.less']
})
export class LogoComponent extends GenericComponent {

  constructor(injector: Injector, public menuService: MenuService) {
    super(injector);
  }


  protected onInit(): void {
    // console.log(`app-logo onInit`);
  }
  protected onDestroy(): void {
    console.log(`app-welcome onDestroy`);
  }
}
