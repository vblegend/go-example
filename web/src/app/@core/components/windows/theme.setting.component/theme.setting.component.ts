import { Component, ElementRef, Injector, ViewChild } from '@angular/core';
import { StringNameValue } from '@core/common/types';
import { GenericComponent } from '@core/components/basic/generic.component';



@Component({
  selector: 'ngx-theme-setting-component',
  styleUrls: ['./theme.setting.component.less'],
  templateUrl: './theme.setting.component.html'
})
export class ThemeSettingComponent extends GenericComponent {
  @ViewChild('TerminalParent', { static: true })
  private terminalDiv!: ElementRef;
  public themes: StringNameValue[] = [
    {
      value: 'default',
      name: 'Light'
    },
    {
      value: 'dark',
      name: 'Dark'
    },
    {
      value: 'cosmic',
      name: 'Cosmic'
    },
    {
      value: 'corporate',
      name: 'Corporate'
    },
    {
      value: 'material-light',
      name: 'Material Light'
    },
    {
      value: 'material-dark',
      name: 'Material Dark'
    }
  ];
  constructor(injector: Injector) {
    super(injector)
  }

  public changeTheme(themeName: string): void {
    // this.themeService.changeTheme(themeName);
    // this.windowRef.close();
  }


  protected onInit(): void {

  }

  protected onDestroy(): void {

  }


}
