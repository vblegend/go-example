import { Injectable } from "@angular/core";

@Injectable({
    providedIn: 'root'
})
export class IconFontService {

    /**
     * 获取所有图标的类名
     */
    public readonly icons: string[] = [];
    /**
     *
     */
    constructor() {

    }

    /**
     * 匹配所有图标类名
     * @param prefix 
     */
    public matchIconClassNames(prefix: string = 'grace-'): void {
        this.icons.length = 0;
        const elements = document.querySelectorAll("body svg symbol");
        elements.forEach(element => {
            if (prefix == null || (prefix && element.id.startsWith(prefix))) {
                this.icons.push(element.id);
            }
        });
    }

}