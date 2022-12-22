import { ComponentRef } from "@angular/core";
import { CaptureTime } from "@core/decorators/capturetime";
import { BasicWidgetComponent } from "@hmi/components/basic-widget/basic.widget.component";
import { WidgetConfigure } from "@hmi/configuration/widget.configure";
import { HmiMath } from "@hmi/utility/hmi.math";
import { Rectangle } from "./common";


export class SelectionService {

    private components: ComponentRef<BasicWidgetComponent>[];
    private _selectionBounds: Rectangle;
    private _configures: WidgetConfigure[] = [];

    /**
     * 获取选区的边界线
     */
    public get bounds(): Rectangle {
        return this._selectionBounds;
    }

    /**
     * 获取所有已选中对象
     */
    public get objects(): ComponentRef<BasicWidgetComponent>[] {
        return this.components.slice();
    }


    public get configures(): WidgetConfigure[] {
        return this._configures.slice();
    }





    /**
     *
     */
    constructor() {
        this.components = [];
        this._selectionBounds = { left: 0, top: 0, width: 0, height: 0 };
    }

    /**
     * 填充一组对象至选区
     * @param comps 
     */
    public fill(comps: ComponentRef<BasicWidgetComponent>[]): void {
        let x = this.components.length;
        this.clear();
        for (let i = 0; i < comps.length; i++) {
            if (comps[i] && this.addItem(comps[i])) {
                x++;
            }
        }
        if (x > 0) this.update();
    }

    /**
     * 将多个对象的状态由选中/未选中之间切换
     * @param comps 
     */
    public toggle(comps: ComponentRef<BasicWidgetComponent>[]): void {
        for (let i = 0; i < comps.length; i++) {
            if (comps[i]) {
                const comp = this.addItem(comps[i]);
                if (comp == null) this.removeItem(comps[i]);
            }
        }
        this.update();
    }

    /**
     * 添加一个对象至选区。
     * @param component 
     */
    public add(component: ComponentRef<BasicWidgetComponent>): ComponentRef<BasicWidgetComponent> | null {
        const comp = this.addItem(component);
        if (comp) this.update();
        return comp;
    }

    /**
     * 是否包含指定对象，或是否选中指定对象
     * @param component 
     * @returns 
     */
    public contains(component: ComponentRef<BasicWidgetComponent>): boolean {
        return this.components.indexOf(component) > -1;
    }

    /**
     * 从选区删除一个对象，取消对象的选中状态
     * @param component 
     */
    public remove(component: ComponentRef<BasicWidgetComponent>): ComponentRef<BasicWidgetComponent> | null {
        const comp = this.removeItem(component);
        if (comp) this.update();
        return comp;
    }


    /**
     * 仅添加一个实例容器，但不做更新操作。\
     * 添加成功返回{容器对象实例}，失败（已在列表）则返回{null}
     * @param component 
     * @returns 
     */
    public addItem(component: ComponentRef<BasicWidgetComponent>): ComponentRef<BasicWidgetComponent> | null {
        if (component) {
            const index = this.components.indexOf(component);
            if (index === -1) {
                this.components.push(component);
                const element = component.location.nativeElement as HTMLElement;
                element.classList.add('hmi-selected-component');
                return component;
            }
        }
        return null;
    }
    /**
     * 仅移除一个实例容器，但不做更新操作。\
     * 移除成功返回{容器对象实例}，失败（未在列表）则返回{null}
     * @param component 
     * @returns 
     */
    private removeItem(component: ComponentRef<BasicWidgetComponent>): ComponentRef<BasicWidgetComponent> | null {
        const index = this.components.indexOf(component);
        if (index > -1) {
            this.components.splice(index, 1);
            const element = component.location.nativeElement as HTMLElement;
            element.classList.remove('hmi-selected-component');
            return component;
        }
        return null;
    }






    /**
     * 清除移除所有选中对象，并恢复其状态。
     */
    public clear(): void {
        while (this.components.length > 0) {
            this.removeItem(this.components[0]);
        }
        this.components = [];
        this.update();
    }

    /***
     * 获取选中对象的个数
     */
    public get length(): number {
        return this.components.length;
    }


    /**
     * 当前选中对象是否包含组合的
     */
    public get hasGrouping(): boolean {
        for (let i = 0; i < this.components.length; i++) {
            if (this.components[i].instance.configure.group != null) {
                return true;
            }
        }
        return false;
    }

    /**
     * 当前选中对象是否包含锁定的
     */
    public get hasLocking(): boolean {
        for (let i = 0; i < this.components.length; i++) {
            if (this.components[i].instance.configure.locked == true) {
                return true;
            }
        }
        return false;
    }





    /**
     * 获取所有选中的范围
     * @returns 
     */
    // @CaptureTime(false, 1)
    public update(): void {
        let bounds: Rectangle | null = null;
        for (let i = 0; i < this.components.length; i++) {
            const rect = this.components[i].instance.getRelativeRect();  // this.components[i].instance.configure.rect
            if (bounds == null) {
                bounds = rect;
            } else {
                bounds = HmiMath.extendsRectangle(bounds, rect);
            }
        }
        this._selectionBounds = bounds || { left: 0, top: 0, width: 0, height: 0 };
        this._configures = this.objects.map(e => e.instance.configure);
    }






}