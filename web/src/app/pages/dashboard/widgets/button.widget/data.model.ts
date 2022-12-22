import { WidgetDataConfigure } from "@hmi/configuration/widget.configure";

/**
 * 图片小部件的数据属性
 */
export interface ButtonWidgetDataModel extends WidgetDataConfigure {
    /**
     * 按钮文本
     */
    text: string;

    icon?: string;
    /**
     * 字符串左右对齐
     */
    align: 'left' | 'center' | 'right';

}