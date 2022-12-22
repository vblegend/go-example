import { Component, Injector } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';
import { RouteConfigure } from '@core/models/route.configure';
import { MenuService } from '@core/services/menu.service';

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.less']
})
export class SidebarComponent extends GenericComponent {
  constructor(injector: Injector, public menuService: MenuService) {
    super(injector);
  }

  public menus: RouteConfigure[] = [];


  protected onInit(): void {
    this.menus = this.menuService.menus;
  }
}
