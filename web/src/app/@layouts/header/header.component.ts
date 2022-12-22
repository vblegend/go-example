import { KeyValue } from '@angular/common';
import { Component, Injector } from '@angular/core';
import { Exception } from '@core/common/exception';
import { FixedTimer } from '@core/common/fixedtimer';
import { TimerTask } from '@core/common/timer.task';
import { GenericComponent } from '@core/components/basic/generic.component';
import { DialogService } from '@core/services/dialog.service';
import { MenuService } from '@core/services/menu.service';
import { ThemeService } from '@core/services/theme.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.less']
})
export class HeaderComponent extends GenericComponent {

  public currentTheme: string = '';
  public today: Date;
  public themes: KeyValue<string, string>[] = [];
  constructor(injector: Injector, public menuService: MenuService, public themeService: ThemeService, private dialogService: DialogService) {
    super(injector);
    this.today = new Date();
  }


  protected onInit(): void {
    this.themes = [];
    this.currentTheme = this.themeService.currentTheme;
    for (const key in this.themeService.themes) {
      this.themes.push({ key, value: this.themeService.themes[key] })
    }
    const timer = this.createTimer(this.timerUpdate, 1000);
  }

  private timerUpdate(task: TimerTask): void {
    this.today = new Date();
  }



  public changeTheme(event: string): void {
    this.themeService.changeTheme(event);
  }

  public serachClick(): void {

  }





  public tclick(): void {
    this.dialogService.createTerminalWindow();
  }

  public throwError(): void {
    throw Exception.build('test Global Error Handle', 'ZHE......');
  }


  public logout(): void {
    this.sessionService.remove('user');
    this.navigate('../login');
  }



  public systemSetting(): void {
    this.showPortal = !this.showPortal;
  }

  public portal_Closed(): void {
    if (this.showPortal) this.showPortal = false;
  }


  public showPortal: Boolean = false;

}
