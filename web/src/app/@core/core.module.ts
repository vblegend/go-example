import { APP_BOOTSTRAP_LISTENER, APP_INITIALIZER, ComponentRef, ErrorHandler, Injector, ModuleWithProviders, NgModule, Optional, Provider, SkipSelf } from '@angular/core';
import { CommonModule } from '@angular/common';
// import { MatNativeDateModule, MAT_RIPPLE_GLOBAL_OPTIONS } from '@angular/material/core';
// import { NbAuthModule } from '@nebular/auth';
// import { NbSecurityModule, NbRoleProvider } from '@nebular/security';


import { RestfulService } from './services/restful.service';
import { AuthGuardService } from './services/auth.guard.service';
import { DialogService } from './services/dialog.service';
import { AccountService } from './services/account.service';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { DragDropModule } from '@angular/cdk/drag-drop';
import { RouterModule } from '@angular/router';
import { LoginPageComponent } from './components/login/loginpage.component';
import { ThemeSettingComponent } from './components/windows/theme.setting.component/theme.setting.component';
import { TerminalComponent } from './components/windows/terminal.component/terminal.component';
import { DefaultPipe } from './pipes/default.pipe';
import { TranslatorPipe } from './pipes/translator.pipe';
import { DocumentTitleService } from './services/document.title.service';
import { NetWorkService } from './services/network.service';
import { NotFoundComponent } from './components/notfound/not-found.component';
import { NzLayoutModule } from 'ng-zorro-antd/layout';
import { NzInputModule } from 'ng-zorro-antd/input';
import { ThemeService } from './services/theme.service';
import { BootstrapService } from './services/bootstrap.service';
import { ErrorComponent } from './components/error/error.component';
import { NzSpaceModule } from 'ng-zorro-antd/space';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzGridModule } from 'ng-zorro-antd/grid';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzMessageModule } from 'ng-zorro-antd/message';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzMenuModule } from 'ng-zorro-antd/menu';
import { MenuService } from './services/menu.service';
import { HoverDirective } from './directives/hover.directive';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { GlobalErrorHandler } from './private/GlobalErrorHandler';
import { NzNotificationModule, NzNotificationService } from 'ng-zorro-antd/notification';
import { NzAvatarModule } from 'ng-zorro-antd/avatar';
import { NzPopoverModule } from 'ng-zorro-antd/popover';
import { UnSelectedDirective } from './directives/unselected.directive';
import { NzDrawerService } from 'ng-zorro-antd/drawer';
import { EVENT_MANAGER_PLUGINS } from '@angular/platform-browser';
import { OutSideEventPluginService } from './services/outside.event.plugin.service';
import { SessionService } from './services/session.service';
import { CacheService } from './services/cache.service';
import { EventBusService } from './services/event.bus.service';
import { AngularSplitModule } from 'angular-split';
import { CollapseComponent } from './components/ngx-collapse/ngx.collapse.component';
import { TimerPoolService } from './services/timer.pool.service';
import { NotificationService } from './services/notification.service';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { HttpHandlerInterceptor, } from './interceptors/http.handler.interceptor';
import { ProsmDirective } from './directives/prism.directive';
import { IconFontService } from './services/iconfont.service';
import { NzCheckboxModule } from 'ng-zorro-antd/checkbox';
import { WindowComponent } from './components/basic/window.component';
import { PortalModule } from '@angular/cdk/portal';
import { NgxIFrameComponent } from './components/ngx-iframe/ngx.iframe.component';

const EXPORT_PIPES: Provider[] = [
  DefaultPipe,
  TranslatorPipe
];


const EXPORT_DIRECTIVES: Provider[] = [
  HoverDirective,
  UnSelectedDirective,
  ProsmDirective
];


/**
 * EXPORT CONPONENTS
 */
const EXPORT_COMPONENTS = [
  WindowComponent,
  LoginPageComponent,
  NotFoundComponent,
  ThemeSettingComponent,
  TerminalComponent,
  ErrorComponent,
  CollapseComponent,
  NgxIFrameComponent
];


/**
 * custom providers services
 */
const PROVIDERS: Provider[] = [
  SessionService,
  DocumentTitleService,
  RestfulService,
  AuthGuardService,
  DialogService,
  AccountService,
  NetWorkService,
  ThemeService,
  BootstrapService,
  MenuService,
  NzDrawerService,
  CacheService,
  EventBusService,
  TimerPoolService,
  NotificationService,
  IconFontService,
];





@NgModule({
  imports: [
    CommonModule,
    // BrowserModule,
    HttpClientModule,
    RouterModule,
    FormsModule,
    ReactiveFormsModule,
    DragDropModule,
    PortalModule,
    NzIconModule,
    NzMenuModule,
    NzLayoutModule,
    NzGridModule,
    NzInputModule,
    NzFormModule,
    NzModalModule,
    NzSpaceModule,
    NzButtonModule,
    NzFormModule,
    NzNotificationModule,
    NzMessageModule,
    NzAvatarModule,
    NzPopoverModule,
    AngularSplitModule,
    NzCheckboxModule,

  ],
  exports: [
    EXPORT_COMPONENTS,
    EXPORT_DIRECTIVES,
    EXPORT_PIPES
  ],
  declarations: [
    EXPORT_COMPONENTS,
    EXPORT_DIRECTIVES,
    EXPORT_PIPES,
  ]
})

export class CoreModule {
  constructor(@Optional() @SkipSelf() parentModule: CoreModule, private themeService: ThemeService, private documentTitleService: DocumentTitleService, private networkService: NetWorkService) { }


  public static forRoot(): ModuleWithProviders<CoreModule> {
    return {
      ngModule: CoreModule,
      providers: [
        ...PROVIDERS,
        {
          provide: APP_BOOTSTRAP_LISTENER,
          useFactory: BootstrapService.BootstrapFactory,
          deps: [BootstrapService],
          multi: true
        },
        {
          provide: APP_INITIALIZER,
          useFactory: BootstrapService.InitializerFactory,
          deps: [BootstrapService],
          multi: true
        },
        {
          provide: ErrorHandler,
          useClass: GlobalErrorHandler,
          deps: [NzNotificationService]
        },
        {
          provide: EVENT_MANAGER_PLUGINS,
          useClass: OutSideEventPluginService,
          multi: true
        },
        {
          provide: HTTP_INTERCEPTORS,
          useClass: HttpHandlerInterceptor,
          multi: true
        },


      ]
    };
  }
}
