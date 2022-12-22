import { Type } from "@angular/core";
import { ObjectUtil } from "@core/util/object.util";
import { WidgetDataConfigure, WidgetDefaultConfigure, WidgetEventConfigure, WidgetStyles } from "@hmi/configuration/widget.configure";
import { EventMeta, WidgetDefaultVlaues, WidgetMetaObject } from "../../configuration/widget.meta.data";


/**
 * 当前版本
 */
export const CurrentVersion: number[] = [0, 0, 1];
/**
 * 配置文件魔法代码
 */
export const DocumentMagicCode: string = 'graphic.config.file';

export const ClipDocumentMagicCode: string = 'graphic.clip.file';


export declare type DecoratorFunction = (target: Function) => void;


/**
 * Hmi 组态缩放模式
 */
export enum HmiZoomMode {
    /**
     * 原始大小，不缩放
     */
    None = 0,
    /**
     * 保持比例缩放
     */
    Scale = 1,
    /**
     * 拉伸缩放
     */
    Stretch = 2,
}


/**
 * 表示一个2维向量坐标
 */
export interface Vector2 {
    x: number;
    y: number;
}
/**
 * 表示一块矩形区域。
 */
export interface Rectangle {
    left?: number;
    top?: number;
    width: number;
    height: number;
}

/**
 * 定义一个Widget对象的所有可触发的事件列表\
 * 使用小部件的 this.dispatchEvent()方法触发事件\
 * 事件所使用的参数必须与事件声明里参数签名一致
 * 
 * @param options 
 */
export function DefineEvents(events: EventMeta[]): DecoratorFunction {
    return function (target: any) {
        const metaData = target.prototype.upgradeWidgetMetaData(target) as WidgetMetaObject;
        if (events == null || events.length == 0) return;
        for (const event of events) {
            const result = metaData.events.find(e => e.event == event.event);
            if (result) throw new Error(`事件重复定义，[${target}][${event.event}]已被定义。`);
            metaData.events.push(event);
        }
    };
}

/**
 * 定义+排除Widget对象在编辑器内所支持的属性\
 * 列表中的属性必须在 hmi-property-element 定义\
 * 如下所示 ```key="signalId"```
 * ```
 *  <ng-template hmi-properties tableName="属性" groupName="数据绑定">
 *      <hmi-property-element key="signalId" [multiple]="false" attributePath="data/signalId"
 *          header="绑定信号" tooltip="绑定信号">
 *          <hmi-button-action-property>
 *          </hmi-button-action-property>
 *      </hmi-property-element>
 *  </ng-template>
 * ```
 * @param properties 提供一个属性列表
 * @param excludeds 属性排除列表  从 properties 排除掉
 * @returns 
 */
export function DefineProperties(properties: string[], ...excludeds: string[]): DecoratorFunction {
    return function (target: any) {
        const metaData = target.prototype.upgradeWidgetMetaData(target) as WidgetMetaObject;
        const props = properties.slice();
        if (excludeds && excludeds.length > 0) {
            for (let i = 0; i < excludeds.length; i++) {
                const index = props.indexOf(excludeds[i]);
                if (index > -1) props.splice(index, 1);
            }
        }
        metaData.properties = props;
    };
}

/**
 * 默认通用的部件属性
 * 
 */
export const CommonWidgetPropertys: string[] = [
    'Name', 'Index', 'Interval', 'Locked',
    'style.bkColor', 'style.bkImage', 'style.bkSize', 'style.opacity', 'style.color', 'style.fontSize', 'style.fontBold',

    'style.border', 'style.radius',


    'Events'
];


/**
 * 定义Widget对象的默认数据\
 *  属性定义查询 @WidgetDefaultVlaues 也可以使用下列装饰器
 * @DefaultInterval
 * @DefaultSize
 * @DefaultStyle
 * @DefaultEvents
 * 
 * @param key 默认属性名
 * @param value 默认值
 * @returns 
 */
export function DefaultValue<T extends keyof WidgetDefaultVlaues>(key: T, value: WidgetDefaultVlaues[T]): DecoratorFunction {
    return function (target: any) {
        const metaData = target.prototype.upgradeWidgetMetaData(target) as WidgetMetaObject;
        metaData.default[key] = value;
    };
}




/**
 * 定义Widget对象的默认Data对象结构\
 * **注意：** 此默认数据必须包含所有字段\
 * **默认数据内的空值必须使用null表示**\
 * 因为data在upgrade环节会将结构中为undefined字段替换为 defaultData 中的字段
 * 
 * @param defaultData 
 * @returns 
 */
export function DefaultData<TData extends WidgetDataConfigure>(defaultData: TData): DecoratorFunction {
    return function (target: any) {
        const metaData = target.prototype.upgradeWidgetMetaData(target) as WidgetMetaObject;
        metaData.default.data = defaultData;
    };
}
/**
 * 定义Widget对象的默认时钟刷新周期
 * @param interval 间隔 单位秒
 * @returns 
 */
export function DefaultInterval(interval: number): DecoratorFunction {
    return function (target: any) {
        const metaData = target.prototype.upgradeWidgetMetaData(target) as WidgetMetaObject;
        metaData.default.interval = interval;
    };
}

/**
 * 定义Widget对象添加时的默认大小\
 * 当手动拖放或双击添加部件时的默认大小
 * @param width 宽度 单位px
 * @param height 高度 单位px
 * @returns 
 */
export function DefaultSize(width: number, height: number): DecoratorFunction {
    return function (target: any) {
        const metaData = target.prototype.upgradeWidgetMetaData(target) as WidgetMetaObject;
        metaData.default.size = { width, height };
    };
}
/**
 * 定义Widget对象的默认样式表
 * @param style 样式表
 * @returns 
 */
export function DefaultStyle(style: WidgetStyles): DecoratorFunction {
    return function (target: any) {
        const metaData = target.prototype.upgradeWidgetMetaData(target) as WidgetMetaObject;
        metaData.default.style = style;
    };
}

/**
 * 定义Widget对象的默认触发事件模板\
 * 通常用于预制事件
 * @param events 事件表
 * @returns 
 */
export function DefaultEvents(events: Record<string, WidgetEventConfigure[]>): DecoratorFunction {
    return function (target: any) {
        const metaData = target.prototype.upgradeWidgetMetaData(target) as WidgetMetaObject;
        metaData.default.events = events;
    };
}


/**
 * 用于声明一个Widget对外接口\
 * 接口的参数需使用{@Params}进行标注\
 * 未标注的参数将不会收到任何数据
 * @param name 定义接口名。
 * @param description 定义接口说明。
 * @param strict 参数严格性（默认为false） 当为true时只有所有参数都匹配才会触发接口事件
 * @returns 
 */
export function WidgetInterface(name: string, description: string, strict?: boolean): (target: any, methodName: string, descriptor: PropertyDescriptor) => void {
    return function (prototype: any, methodName: string, descriptor: PropertyDescriptor) {
        const metaData = prototype.upgradeWidgetMetaData(prototype.constructor) as WidgetMetaObject;
        if (metaData.interface[methodName] == null) {
            metaData.interface[methodName] = { methodName, name, description, descriptor, args: [], strict: strict ? true : false };
        }
        const method = metaData.interface[methodName];
        method.name = name;
        method.strict = strict ? true : false;
        method.description = description;
        method.descriptor = descriptor;
        method.methodName = methodName;
        // ObjectUtil.freeze(metaData.interface);
    };
}

/**
 * 用于声明接口参数
 * @param argName 参数对外的名称
 * @param typed 对象类型 Number Boolean Array String
 * @returns 
 */
export function Params(argName: string, typed?: 'number' | 'string' | 'boolean' | 'array'): (target: Function, methodName: string, paramIndex: number) => void {
    return function (target: any, methodName: string, paramIndex: number) {
        const metaData = target.constructor.prototype.upgradeWidgetMetaData(target.constructor) as WidgetMetaObject;
        if (metaData.interface[methodName] == null) {
            metaData.interface[methodName] = { methodName: undefined, name: undefined, description: undefined, descriptor: undefined, args: [], strict: false };
        }
        const method = metaData.interface[methodName];
        method.args![paramIndex] = { argName, paramIndex, typed };
    }
}
