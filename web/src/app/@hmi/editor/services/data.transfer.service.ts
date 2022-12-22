import { Injectable } from "@angular/core";
/**
 * 数据拖拽服务\
 * 用于替换系统drag相关事件的dataTransfer
 */
@Injectable({
    providedIn: 'root'
})
export class DataTransferService {

    private readonly _values: Record<string, string | null> = {};

    /**
     * 设置一个字符串键值
     * @param key 
     * @param value 
     */
    public setText(key: string, value: string | null): void {
        this._values[key] = value;
        if (value == null) {
            delete this._values[key];
        }
    }
    /**
     * 获取字符串的键值
     * @param key 
     * @returns 
     */
    public getText(key: string): string | null {
        return this._values[key]
    }


    public result: boolean = false;
}
