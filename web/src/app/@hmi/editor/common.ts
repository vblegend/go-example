/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable @typescript-eslint/explicit-module-boundary-types */
import { Injector } from "@angular/core";
import { BasicPropertyComponent } from "./components/basic-property/basic.property.component";

/**
 * 按钮的动作工厂
 */
export abstract class ButtonActionFactory {

    /**
     * 固定的构造函数。\
     * 不要修改
     */
    constructor(protected injector: Injector) {

    }


    /**
     * 属性到绑定元素的转换函数
     * @param value 显示到按钮中
     */
    public abstract toBinding(value: any): string;

    /**
     * 按钮是否可以按下
     * @param value 
     */
    public abstract canExecute(value: any): boolean;

    /**
     * 按钮执行函数\
     * 可以在这里调用 **propertie.saveAndUpdate(value);** 函数对属性赋值
     * @param propertie 属性控件
     */
    public abstract execute(propertie: BasicPropertyComponent<any>): void;
}

/**
 * 一个默认按钮属性工厂
 */
export class DefaultButtonActionFactory<TValue> extends ButtonActionFactory {

    public toBinding(value: TValue): string {
        return `默认按钮工厂`;
    }

    public canExecute(value: TValue): boolean {
        return true;
    }

    public execute(propertie: BasicPropertyComponent<TValue>): void {
        console.warn('该事件出自默认的按钮动作工厂，请绑定属性的指定工厂[factory]=""');
        // propertie.saveAndUpdate(<TValue><unknown>'test');
    }
}