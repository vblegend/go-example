import { WidgetDataConfigure } from "@hmi/configuration/widget.configure";

/**
 * 图片小部件的数据属性
 */
export interface SignalWidgetDataModel extends WidgetDataConfigure {
    /**
     * 设备ID
     */
    eqid?: number;

    /**
     * 信号Id
     */
    sgid?: number;

    /**
     * 显示内容
     */
    text?: string;

    /**
     * 是否显示单位
     */
    unit?: boolean;
    /**
     * 字符串左右对齐
     */
    align: 'left' | 'center' | 'right';
}