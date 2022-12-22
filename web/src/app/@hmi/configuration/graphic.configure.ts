import { HmiZoomMode } from "@hmi/editor/core/common";
import { WidgetConfigure } from "./widget.configure";


export interface GraphicOptions {
    /**
     * 组态画面宽度
     */
    width: number;
    /**
     * 组态画面高度
     */
    height: number;

    /**
     * 组态画面缩放模式
     */
    zoomMode: HmiZoomMode;
}


export interface GraphicConfigure extends GraphicOptions {

    magic: string;

    version: number[];

    widgets: WidgetConfigure[];
}

export interface ClipboardData {

    magic: string;

    version: number[];

    widgets: WidgetConfigure[];
}
