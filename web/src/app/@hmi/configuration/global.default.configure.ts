import { ClipDocumentMagicCode, DocumentMagicCode } from "@hmi/editor/core/common";
import { ClipboardData, GraphicConfigure } from "./graphic.configure";
import { WidgetStyles } from "./widget.configure";


/**
 * 小部件的全局默认样式\
 * 所有部件将会应用\
 * 如果部件定义了默认样式重写了属性则不会被应用\
 * 值为undefined的属性将不会被处理
 */
export const DefaultGlobalWidgetStyle: WidgetStyles = {
    bkColor: undefined,
    color: undefined,
    opacity: undefined,
    border: undefined,
    radius: undefined,
    fontFamily: undefined,
    fontSize: undefined,
    textAlign: undefined,
    rotate: undefined
}

  /**
   * 校验配置文件格式
   * @param json 
   */
   export function verifyDocument (json: GraphicConfigure): void {
    if (json.magic != DocumentMagicCode) throw new Error('加载配置失败，数据文档不是有效格式。');
    if (json.version == null || json.version.length != 3) throw new Error('加载配置失败，数据文档不是有效格式。');
    // if (json.version != CurrentVersion){
    //   throw Exception.build('编辑器', '加载配置失败，数据文档不是有效格式。');
    // }
  }
  /**
   * 校验配置文件格式
   * @param json 
   */
   export function verifyClipboardDocument (json: ClipboardData): void {
    if (json.magic != ClipDocumentMagicCode) throw new Error('加载配置失败，数据文档不是有效格式。');
    if (json.version == null || json.version.length != 3) throw new Error('加载配置失败，数据文档不是有效格式。');
    // if (json.version != CurrentVersion){
    //   throw Exception.build('编辑器', '加载配置失败，数据文档不是有效格式。');
    // }
  }