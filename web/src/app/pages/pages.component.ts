import { Component, Injector } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';
import { MenuService } from '@core/services/menu.service';

// import { MENU_ITEMS } from './pages-menu';

@Component({
  selector: 'ngx-pages',
  styleUrls: ['pages.component.less'],
  templateUrl: 'pages.component.html',
})
export class PagesComponent extends GenericComponent {

  /**
   *
   */
  constructor(injector: Injector, public menuService: MenuService) {
    super(injector);
  }

  protected onQueryChanges(): void {
    console.log(`ngx-pages onRouter ${this.queryParams.get('id')}`);

  }




}


