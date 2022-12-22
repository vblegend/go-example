import { Component, Injector, Input, ViewChild } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';
import { WidgetSchemaService } from './services/widget.schema.service';
import { ViewCanvasComponent } from './components/view-canvas/view.canvas.component';
import { MetaDataService } from './editor/services/meta.data.service';
import { GraphicConfigure } from './configuration/graphic.configure';
import { CurrentVersion, DocumentMagicCode, HmiZoomMode } from './editor/core/common';
import { verifyDocument } from './configuration/global.default.configure';

@Component({
  selector: 'hmi-viewer',
  templateUrl: './hmi.viewer.component.html',
  styleUrls: ['./hmi.viewer.component.less'],
  providers: [MetaDataService],
})
export class HmiViewerComponent extends GenericComponent {
  @ViewChild('canvas', { static: true })
  public canvas!: ViewCanvasComponent;

  /**
   * 组态宽度
   */
  public width: number = 1920;

  /**
   * 组态高度
   */
  public height: number = 1080;

  /**
   * 组态缩放类型
   */
  public zoomMode: HmiZoomMode = HmiZoomMode.Scale;

  /**
   *
   */
  constructor(protected injector: Injector, public provider: WidgetSchemaService) {
    super(injector);
  }

  /**
   * 从Json对象加载对象列表至画布
   * @param json 
   */
  public loadFromJson(json: GraphicConfigure): void {
    verifyDocument(json);
    if (json.zoomMode != null) this.zoomMode = json.zoomMode;
    if (json.width != null) this.width = json.width;
    if (json.height != null) this.height = json.height;
    this.canvas.clear();
    for (const config of json.widgets) {
      const compRef = this.canvas.parseComponent(config);
      if (compRef) this.canvas.add(compRef);
    }
    this.canvas.updatezIndexs();
  }



  /**
   * 将画布的内容转换为Json对象
   */
  public toJson(): GraphicConfigure {
    const widgets = this.canvas.children.map(e => e.instance.configure);
    return {
      magic: DocumentMagicCode,
      version: CurrentVersion,
      width: this.width,
      height: this.height,
      zoomMode: this.zoomMode,
      widgets: widgets
    };
  }


  public getStyle(): Record<string, string> {
    const result: Record<string, string> = {};
    const element = this.viewContainerRef.element.nativeElement as HTMLDivElement;
    const parentRect = element.getBoundingClientRect();
    const rx = parentRect.width / this.width;
    const ry = parentRect.height / this.height;
    if (this.zoomMode === HmiZoomMode.Scale) {
      const zoomScale = Math.min(rx, ry);
      result['transform'] = `scale(${zoomScale})`;
    } else if (this.zoomMode === HmiZoomMode.Stretch) {
      result['transform'] = `scale(${rx},${ry})`;
    } else if (this.zoomMode === HmiZoomMode.None) {
      result['transform'] = `scale(1)`;
    }
    return result;
  }

}
