import { ComponentFactoryResolver, Injectable, Injector } from "@angular/core";
import { WidgetSchema } from "@hmi/configuration/widget.schema";

class WidgetSchemaCategory {
    public name: string;
    public children: WidgetSchema[];

    /**
     *
     */
    constructor(category: string) {
        this.children = [];
        this.name = category;
    }




}



/**
 * 小部件 部件架构注册服务
 */
@Injectable({
    providedIn: 'root'
})
export class WidgetSchemaService {

    private _widgetsMap: Record<string, WidgetSchema>;
    private componentFactoryResolver: ComponentFactoryResolver;
    private _categorys: WidgetSchemaCategory[];


    public get categorys(): WidgetSchemaCategory[] {
        return this._categorys.slice();
    }


    /**
     *
     */
    constructor(protected injector: Injector) {
        this._widgetsMap = {};
        this._categorys = [];
        this.componentFactoryResolver = injector.get(ComponentFactoryResolver);
    }

    public register(data: WidgetSchema[]): void {
        for (const widget of data) {
            const factory = this.componentFactoryResolver.resolveComponentFactory(widget.component);
            if (factory) {
                const type = factory.selector;
                if (this._widgetsMap[type] == null) {
                    widget.type = type;
                    this._widgetsMap[type] = widget;
                    Object.freeze(widget);
                    let category = this.findWidgetCategory(widget.classify);
                    if (category == null) {
                        category = new WidgetSchemaCategory(widget.classify);
                        this._categorys.push(category);
                    }
                    category.children.push(widget);
                } else {
                    console.warn(`关键字${type}已被注册，跳过当前组件：${widget.component}`);
                }
            }
            else {
                console.warn(`未找到部件对象${widget.component}`);
            }
        }
    }

    public findWidgetCategory(category: string): WidgetSchemaCategory | undefined {
        return this._categorys.find(e => e.name === category);
    }

    public getType(type: string | null): WidgetSchema | null {
        if (type == null) {
            return null;
        } else {
            return this._widgetsMap[type];
        }
    }

    public random(): WidgetSchema | undefined {
        if (this._categorys.length == 0) return undefined;
        // eslint-disable-next-line no-constant-condition
        while (true) {
            const categoryIndex = Math.floor(Math.random() * this._categorys.length);
            const category = this._categorys[categoryIndex];
            if (category && category.children.length > 0) {
                const index = Math.floor(Math.random() * category.children.length);
                return category.children[index];
            }
        }
    }
}