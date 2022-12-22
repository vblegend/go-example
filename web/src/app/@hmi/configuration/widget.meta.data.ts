import { Type } from "@angular/core";
import { ObjectUtil } from "@core/util/object.util";
import { BasicWidgetComponent } from "@hmi/components/basic-widget/basic.widget.component";
import { Size, WidgetDataConfigure, WidgetEventConfigure, WidgetStyles } from "@hmi/configuration/widget.configure";
import { AnyObject } from "chart.js/types/basic";
import { Vector2 } from "../editor/core/common";

export interface MethodArgument {
    /**
     * 接口的参数名
     */
    argName: string;
    /**
     * 接口参数索引
     */
    paramIndex: number;

    /**
     * 参数的类型
     */
    typed?: 'number' | 'string' | 'boolean' | 'array';
}


export interface MethodMeta {
    /**
     * 接口方法名称
     */
    methodName: string | undefined;
    /**
     * 接口名
     */
    name: string | undefined;
    /**
     * 是否为严格模式（默认为false）
     * 严格模式下需所有参数都被匹配才会触发
     */
    strict: boolean | undefined;
    /**
     * 接口说明
     */
    description: string | undefined;
    /**
     * 接口方法对象
     */
    descriptor: PropertyDescriptor | undefined;
    /**
    * 接口参数
    */
    args: MethodArgument[] | undefined;

}

export interface EventMeta {
    /**
     * 事件名
     */
    event: string;
    /**
     * 事件的参数列表
     */
    eventParams: string[];
    /**
     * 事件说明
     */
    eventName: string;
}

/**
 * 小部件默认数据对象
 */
export class WidgetDefaultVlaues {
    /**
     * 定义Widget对象的默认触发事件模板\
     * 用于特殊对象的预定义事件
     */
    public events: Record<string, WidgetEventConfigure[]> = {};

    /**
     * 定义Widget对象的默认刷新周期 单位：秒
     */
    public interval: number = 0;

    /**
     * 定义Widget对象的默认样式表
     */
    public style: WidgetStyles = {};

    /**
     * 定义Widget对象的默认大小 单位：px
     */
    public size: Size = { width: 0, height: 0 };

    /**
     * 定义Widget对象的默认Data对象结构\
     * **注意：** 此默认数据必须包含所有字段\
     * **默认数据内的空值必须使用null表示**\
     * 因为data在upgrade环节会将结构中为undefined字段替换为 defaultData 中的字段
     */
    public data: WidgetDataConfigure = {};




    public copy(target: WidgetDefaultVlaues): void {
        this.events = ObjectUtil.clone(target.events)!;
        this.interval = target.interval;
        this.style = ObjectUtil.clone(target.style)!;
        this.size = ObjectUtil.clone(target.size)!;
        this.data = ObjectUtil.clone(target.data)!;
    }


}




/**
 * 部件的元数据对象
 */
export class WidgetMetaObject {

    /**
     * 
     * @param typed 元数据所属类型
     */
    constructor(public readonly typed: Type<BasicWidgetComponent>) {

    }


    /**
     * Widget 接口声明
     */
    public readonly interface: Record<string, MethodMeta> = {}

    /**
     * Widget 事件声明
     */
    public readonly events: EventMeta[] = [];


    /**
     * Widget 默认数据
     */
    public readonly default: WidgetDefaultVlaues = new WidgetDefaultVlaues();


    /**
   * 对象支持的属性key
   */
    public properties: string[] = [];



    public upgrade(typed: Type<BasicWidgetComponent>): WidgetMetaObject {
        const meta = new WidgetMetaObject(typed);
        //
        const $interface = ObjectUtil.clone(this.interface);
        for (const key in $interface) meta.interface[key] = $interface[key];
        //
        const events = ObjectUtil.clone(this.events)!;
        meta.events.push(...events);
        // 
        meta.default.copy(this.default);
        //
        meta.properties = this.properties.slice();
        return meta;
    }



}