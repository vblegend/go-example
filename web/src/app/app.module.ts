/**
 * @license
 * Copyright Akveo. All Rights Reserved.
 * Licensed under the MIT License. See License.txt in the project root for license information.
 */
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ApplicationRef, APP_BOOTSTRAP_LISTENER, APP_INITIALIZER, ChangeDetectorRef, ComponentRef, NgModule, NgZone } from '@angular/core';
import { HttpClient, HttpClientModule } from '@angular/common/http';

import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';
import { LOCALE_ID } from '@angular/core';
import { CommonModule, registerLocaleData } from '@angular/common';
import { NZ_I18N } from 'ng-zorro-antd/i18n';
import { zh_CN } from 'ng-zorro-antd/i18n';
import zh from '@angular/common/locales/zh';
import { IconsProviderModule } from './icons-provider.module';
import { BootstrapService } from '@core/services/bootstrap.service';
import { NetWorkService } from '@core/services/network.service';
import { CoreModule } from '@core/core.module';
import { DocumentTitleService } from '@core/services/document.title.service';
import { Exception } from '@core/common/exception';
import { ThemeService } from '@core/services/theme.service';
import { NzIconService } from 'ng-zorro-antd/icon';
import { SessionService } from '@core/services/session.service';
import { CacheService } from '@core/services/cache.service';
import { SchedulingTask, TaskMode } from './pages/tasks/task-model/tasks.model';
import { LocalCache } from '@core/cache/local.cache';
import { HmiModule } from 'app/@hmi/hmi.module';
import { NgxEchartsModule } from 'ngx-echarts';
import { MenuService } from '@core/services/menu.service';


registerLocaleData(zh);



@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    CommonModule,
    BrowserModule,
    AppRoutingModule,
    CoreModule.forRoot(),
    HmiModule.forRoot(),
    BrowserAnimationsModule,
    HttpClientModule,
    NgxEchartsModule.forRoot({
      echarts: import('echarts')
    }),
    // IconsProviderModule,
  ],
  providers: [
    { provide: LOCALE_ID, useValue: 'zh-CN' }, // replace "de-at" with your locale
    { provide: NZ_I18N, useValue: zh_CN },

  ],
  bootstrap: [AppComponent]
})
export class AppModule {


  constructor(private bootstrapService: BootstrapService,
    private netWorkService: NetWorkService,
    private sessionService: SessionService,
    private http: HttpClient,
    private zone: NgZone,
    private appRef: ApplicationRef,
    private themeService: ThemeService,
    private cacheService: CacheService,
    private documentTitleService: DocumentTitleService,
    private networkService: NetWorkService,
    private iconService: NzIconService,
    public menuService: MenuService) {


    // this.testClone();
    // this.zone.runOutsideAngular(this.testClone.bind(this));


    // initialization theme
    themeService.registerTheme({ dark: '黑暗主题', white: '亮色主题' });
    const theme = this.sessionService.get<string>('theme');
    if (theme) {
      themeService.loadTheme(theme);
    } else {
      themeService.changeTheme('dark');
    }
    // loading 
    bootstrapService.loadingElement = document.getElementById('global-spinner');
    bootstrapService.runAtBootstrap(this.init, this);
    // title service
    documentTitleService.defaultTitle = { value: 'Administrator System', needsTranslator: false };
    documentTitleService.register();

    // register iconfont
    this.iconService.fetchFromIconfont({
      scriptUrl: 'assets/fonts/iconfont.js'
    });

  }




  private testClone() {


    // new WebWorker().Test();

    this.cacheService.register("tasks", new LocalCache('', e => e.taskId));
    // .get('seg');

    const tasks = [];
    for (let i = 0; i < 100000; i++) {
      tasks.push({
        taskId: i, taskName: '张三 - ' + i,
        service: 'data.service',
        online: true,
        serviceId: 'a001',
        mode: TaskMode.Manual,
        ipAddress: '123@gmail.com',
        data: ['abc', 'ABC', 1, 2, 3]
      });
    }

    const t1 = this.cacheService.get("tasks").subscribe(e => {
      console.log(`1+-type:${e.type},data.length:${e.data.length} this:${this.constructor.name}`);
    }, f => f.taskId > 100 && f.taskId < 50000);
    const t2 = this.cacheService.get("tasks").subscribe(e => {
      console.log(`2+-type:${e.type},data.length:${e.data.length} this:${this.constructor.name}`);
    }, f => f.taskId == 5);

    // t1();
    // t2();

    // console.time('clone.test');
    // this.cacheService.get("tasks").load(tasks);

    // for (let v = 0; v < 20; v++) {
    //   tasks[v].taskName = '李四' + v;
    // }
    // this.cacheService.get("tasks").batchPut(tasks.slice(0, 20));
    // this.cacheService.get("tasks").batchPut(tasks);
    // // for (let i = 0; i < 1000; i++) {
    // //   this.cacheService.tasks.remove(100000 + i);
    // // }
    // console.timeEnd('clone.test');
    // const lst = this.cacheService.get("tasks").getAll();
    // console.log(lst.length);
  }








  private async init(): Promise<void> {
    this.netWorkService.url = 'ws://127.0.0.1:8000/ws/test';

    // await this.menuService.load();
    // try {
    // const state = await this.netWorkService.connection();
    // if (!state) throw Exception.build('app init', 'failed to connect to server!');
    // console.time('websocket');
    // const list: Promise<string>[] = [];

    // for (let i = 0; i < 1000; i++) {
    //   list.push(this.networkService.send<string, string>('dasds', `data-${Math.random() * 1000000}`, 10000));
    //   // .then(result => {
    //   //   console.log(`result：${result}`);
    //   // });
    // }
    // await Promise.all(list);
    // console.timeEnd('websocket');
    // } catch (e) {
    // }
  }






}
