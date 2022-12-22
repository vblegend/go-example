import { Injectable } from '@angular/core';
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal';

import { TerminalComponent } from '../components/windows/terminal.component/terminal.component';
import { ThemeSettingComponent } from '../components/windows/theme.setting.component/theme.setting.component';

@Injectable({
  providedIn: 'root'
})
export class DialogService {
  constructor(private modalService: NzModalService) {


  }


  public createTerminalWindow(): NzModalRef<TerminalComponent, boolean> {
    return this.modalService.create<TerminalComponent, boolean>({
      nzTitle: 'terminal',
      nzWidth:'646px',
      // nzClassName : 'terminal',
      nzContent: TerminalComponent,
    });
  }


  public createThemeWindow(): void {
    // return this.windowService.open(ThemeSettingComponent, {
    //   title: 'Theme Setting',
    //   initialState: NbWindowState.FULL_SCREEN,
    //   windowClass: 'terminal',
    //   hasBackdrop: true,
    //   closeOnBackdropClick: false,
    //   closeOnEsc: false,
    //   buttons: {
    //     maximize: false,
    //     minimize: false
    //   }
    // });
  }


}
