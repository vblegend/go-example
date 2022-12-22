import { AnyObject } from "@core/common/types";
import { Rectangle } from "@hmi/editor/core/common";


/**
 * 2D组态的绑定数据结构
 */
export interface WidgetDataConfigure {




}

/**
 * 2D组态元素样式
 */
export interface WidgetStyles {

    /**
     * 组件背景色
     */
    bkColor?: string;

    /**
     * 背景图片
     */
    bkImage?: string

    /**
     * 背景大小  仅支持 拉伸 和 平铺
     */
    bkSize?: 'stretch' | 'tile';

    /**
     * 前景颜色
     */
    color?: string;

    /**
     * 组件不透明度
     */
    opacity?: number;

    /**
     * 边框 solid 1px #ffffff
     */
    border?: string;

    /**
     * 边框圆角
     */
    radius?: string;

    /**
     * 字体
     */
    fontFamily?: string;

    /**
     * 字体大小
     */
    fontSize?: number;

    /**
     * 是否为粗体
     */
    fontBold?: boolean;



    /**
     * 文本对齐方式
     */
    textAlign?: string;


    /**
     * 对象的旋转方位
     */
    rotate?: number;




}

export interface Position {
    /**
     * 左侧间距
     */
    left: number;
    /**
     * 顶侧间距
     */
    top: number;
}

export interface Size {
    /**
     * 组件宽度
     */
    width: number;
    /**
     * 组件高度
     */
    height: number;
}






export interface WidgetDefaultConfigure {
    /**
     * 位置与大小
     */
    rect: Rectangle | null;
    /**
     * 
     * 样式属性，主动更新。
     * 所有为空的属性请置为null，不要使用undefined\
     * 因为设置为undefined会触发组件的数据兼容从而恢复为默认值
     */
    style: WidgetStyles;

    /**
     * 组态绑定数据，被动更新 需事件通知\
     * 所有为空的属性请置为null，不要使用undefined\
     * 因为设置为undefined会触发组件的数据兼容从而恢复为默认值
     */
    data: WidgetDataConfigure;

    /**
     * 刷新间隔(秒)  为0时不刷新
     */
    interval: number;

    /**
     * 部件的事件触发
     * 一个事件可以触发多个接口方法
     */
    events: Record<string, WidgetEventConfigure[]>;
}


export interface WidgetEventConfigure {
    /**
     * 目标对象，为空则广播给所有部件
     */
    target: string | null;
    /**
     * 对象的接口方法名
     */
    method: string | null;
    /**
     * 指定参数，重写参数
     */
    params?: Record<string, any>;
}


/**
 * 2D组态元素配置
 */
export interface WidgetConfigure extends WidgetDefaultConfigure {

    /**
     * 对象唯一ID
     */
    id: string;
    /**
     * 对象名字
     */
    name: string;
    /**
     * 对象类型
     */
    type: string | null;
    /**
     * 被锁定的
     */
    locked?: boolean;
    /**
     * 是否忽略鼠标事件 默认为 false
     * 为true时鼠标事件穿透
     */
    ignoreEvent?: boolean;
    /**
     * 组件所属分组
     */
    group?: number;
    /**
     * 组件层级索引
     */
    zIndex?: number;
    /**
     * 是否可见
     */
    visible?: boolean;



}