import { Injectable } from "@angular/core";
import { Guid } from "@core/util/guid";

import * as CryptoJS from 'crypto-js'
import { CrossOriginService, IActionResponse } from "./cross.origin.service";


@Injectable({
    providedIn: 'root'
})
export class SessionService {

    public readonly confuseCode = {
        version: 1.0,
        date: '2020-12-12 00:00:00',
        password: '~!@#$%^&*(*))/**/'
    }

    /**
     *{ callback: this.theme_change, context: this, broadcast: true }
     */
    constructor() {
        CrossOriginService.registerRequestHandler('session.getAll', { callback: this.session_getall, context: this, shared: false });
        CrossOriginService.registerRequestHandler('session.removeItem', { callback: this.session_remove, context: this, shared: true });
        CrossOriginService.registerRequestHandler('session.setItem', { callback: this.session_set, context: this, shared: true });
    }
    // 主机端处理
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    private session_getall(action: string, data: string): IActionResponse {
        const result: Record<string, unknown> = {};
        for (const key of Object.keys(sessionStorage)) {
            const value = sessionStorage.getItem(key);
            result[key] = value;
        }
        return { success: true, data: result };
    }


    private session_remove(action: string, data: string): IActionResponse {
        this._remove(data);
        return { success: true, data: 'OK' };
    }

    private session_set(action: string, data: { key: string, value: string }): IActionResponse {
        this.localSet(data.key, data.value);
        return { success: true, data: 'OK' };
    }

    // 或者不管 是来自顶层的事件或子窗口的事件统一postmessage到top处理
    public changeTheme(themeId: string): void {
        CrossOriginService.request('theme.change', themeId);
    }




    /**
     * set the key content in the session
     * @param key 
     * @param value 
     */
    public async set<T>(key: string, value: T): Promise<void> {
        await CrossOriginService.request('session.setItem', { key: key, value: value });
    }

    /**
     * set the key content to the session 
     * @param key 
     * @returns 
     */
    public get<T>(key: string): T | null {
        const _key = this.generateKey(key);
        const value = sessionStorage.getItem(_key);
        if (value == null) return null;
        const data = CryptoJS.AES.decrypt(value, this.generateIV(key));
        const result = data.toString(CryptoJS.enc.Utf8);
        return JSON.parse(result);
    }

    /**
     * remove the key from session
     * @param key 
     */
    public remove(key: string): void {
        const _key = this.generateKey(key);
        CrossOriginService.request('session.removeItem', _key);
    }


    public localSet(key: string, value: string): void {
        const _key = this.generateKey(key);
        if (value == null) {
            this.remove(_key);
            return;
        }
        const data = CryptoJS.AES.encrypt(JSON.stringify(value), this.generateIV(key)).toString();
        sessionStorage.setItem(_key, data);
    }

    public _remove(key: string): void {
        sessionStorage.removeItem(key);
    }


    public generateKey(key: string): string {
        return CryptoJS.MD5(key + JSON.stringify(this.confuseCode)).toString();
    }


    public generateIV(key: string): string {
        return CryptoJS.MD5(JSON.stringify(this.confuseCode) + key + '$').toString();
    }
}