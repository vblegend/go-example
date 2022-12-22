import { WidgetDataConfigure } from "@hmi/configuration/widget.configure";


export enum HorizontalAlign {
    Left = 'left',
    Center = 'center',
    Right = 'right'
}

export enum VerticalAlign {
    Top = 0,
    Middle = 1,
    Bottom = 2
}






/**
 * 图片小部件的数据属性
 */
export interface TextWidgetDataModel extends WidgetDataConfigure {
    /**
     * 图片地址
     */
    text: string | null | undefined;

    /**
     * 字符串左右对齐
     */
    align: 'left' | 'center' | 'right';

}