import { Injectable } from '@angular/core';
import { Exception } from '../common/exception';
import { ThemeConfigure } from '../common/themeconfigure';
import { CrossOriginService, IActionResponse } from './cross.origin.service';
import { SessionService } from './session.service';

@Injectable({
    providedIn: 'root',
})
export class ThemeService {
    private readonly _themes!: Record<string, string>;
    private _currentTheme!: string;


    public get currentTheme(): string {
        return this._currentTheme;
    }

    constructor(private sessionService: SessionService) {
        this._themes = {
            default: '默认主题'
        }
        CrossOriginService.registerRequestHandler('theme.change', { callback: this.theme_change, context: this, shared: true });
    }

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    private theme_change(action: string, data: string): IActionResponse {
        this.loadTheme(data);
        return { success: true, data: 'OK' };
    }

    /**
     * 更改当前主题，支持跨域IFrame窗口的更改
     * @param themeId 
     */
    public changeTheme(themeId: string): void {
        CrossOriginService.request('theme.change', themeId);
    }

    /**
     * 仅本地切换主题 \
     * 不涉及跨域窗口交互，处理跨域IFrame窗口主题调用 @changeTheme
     * @param themeId 
     * @returns 
     */
    public async loadTheme(themeId: string): Promise<void> {
        if (this._currentTheme == themeId) return;
        const themeClass = `theme-${themeId}`;
        await this.loadCssFile(`${themeClass}.css`, themeClass);
        document.documentElement.classList.add(themeId);
        if (this._currentTheme) {
            this.removeTheme(this._currentTheme);
        }
        this._currentTheme = themeId;
        this.sessionService.localSet('theme', themeId);
    }



    public get themes(): Record<string, string> {
        return this._themes;
    }


    /**
     * register themes
     * @param _themes 
     */
    public registerTheme(_themes: Record<string, string>): void {
        for (const key in _themes) {
            this._themes[key] = _themes[key];
        }
    }



    /**
     * 移除一个主题
     * @param themeId 
     */
    private removeTheme(themeId: string) {
        const themeClass = `theme-${themeId}`;
        document.documentElement.classList.remove(themeId);
        const removedThemeStyle = document.getElementById(themeClass);
        if (removedThemeStyle) {
            document.head.removeChild(removedThemeStyle);
        }
    }


    /**
     * 改变主题
     * @param themeId 
     * #throw Exception
     */
    // public async changeTheme2(themeId: string = 'default'): Promise<void> {
    //     if (!CrossOriginService.isTopLevel) {
    //         CrossOriginService.request({ action: 'theme.change', data: themeId });
    //         return;
    //     }
    //     if (this._themes[themeId] == null) throw Exception.build('change Theme', "don't try to load an undeclared theme");
    //     try {
    //         if (this._currentTheme === themeId) return;
    //         await this.loadTheme(themeId);
    //         if (this._currentTheme) {
    //             this.removeTheme(this._currentTheme);
    //         }
    //         this._currentTheme = themeId;
    //          this.sessionService.set('theme', themeId);
    //         CrossOriginService.broadcast({ action: 'theme.change', data: themeId });
    //     } catch (ex) {
    //         throw Exception.fromCatch('change Theme', ex, 'A fatal error was encountered while loading the theme file.')
    //     }
    // }





    private loadCssFile(href: string, id: string): Promise<Event> {
        return new Promise((resolve, reject) => {
            const style = document.createElement('link');
            style.rel = 'stylesheet';
            style.href = href;
            style.id = id;
            style.onload = resolve;
            style.onerror = reject;
            document.head.append(style);
        });
    }
}
