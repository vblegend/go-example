/**
 * @license
 * Copyright Akveo. All Rights Reserved.
 * Licensed under the MIT License. See License.txt in the project root for license information.
 */
import { COMPILER_OPTIONS, enableProdMode, LOCALE_ID } from '@angular/core';
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';
import { registerLocaleData } from '@angular/common';
import zh from '@angular/common/locales/zh';
import { AppModule } from './app/app.module';
import { environment } from './environments/environment';
import 'zone.js';  // Included with Angular CLI.
import { CrossOriginService } from '@core/services/cross.origin.service';

if (environment.production) {
  enableProdMode();
}

document.oncontextmenu = (e: MouseEvent) => {
  const element = e.target as HTMLElement;
  if (element instanceof HTMLInputElement) return;
  const selection = document.getSelection()!.toString();
  if (selection === '') e.preventDefault();
}
registerLocaleData(zh);



async function init() {
  // 初始化跨域服务
  await CrossOriginService.setup();
  // 同步session
  if (!CrossOriginService.isTopLevel) {
    const session = await CrossOriginService.request<Record<string, string>>("session.getAll");
    sessionStorage.clear();
    for (const key in session) {
      sessionStorage.setItem(key, session[key]);
    }
  }
  // 启动应用
  platformBrowserDynamic().bootstrapModule(AppModule, {}).catch(err => console.error(err));
}


init();







