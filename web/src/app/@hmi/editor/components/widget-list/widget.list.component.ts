import { Component, Injector, Input } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';
import { ObjectUtil } from '@core/util/object.util';
import { WidgetAddCommand } from '@hmi/editor/commands/widget.add.command';
import { DisignerCanvasComponent } from '@hmi/editor/components/disigner-canvas/disigner.canvas.component';
import { WidgetConfigure } from '@hmi/configuration/widget.configure';
import { WidgetSchema } from '@hmi/configuration/widget.schema';

import { HmiEditorComponent } from '@hmi/editor/hmi.editor.component';
import { WidgetSchemaService } from '@hmi/services/widget.schema.service';
import { HmiMath } from '@hmi/utility/hmi.math';
import { WidgetDefaultVlaues } from '@hmi/configuration/widget.meta.data';
import { DataTransferService } from '@hmi/editor/services/data.transfer.service';
import { SelectionFillCommand } from '@hmi/editor/commands/selection.fill.command';


@Component({
  selector: 'hmi-widget-list',
  templateUrl: './widget.list.component.html',
  styleUrls: ['./widget.list.component.less']
})
/**
 * 橡皮筋套选工具
 */
export class WidgetListComponent extends GenericComponent {
  @Input() canvas!: DisignerCanvasComponent;

  /**
   *
   */
  constructor(protected injector: Injector, public provider: WidgetSchemaService, private dataTransferService: DataTransferService, public editor: HmiEditorComponent) {
    super(injector);

  }

  protected onInit(): void {

  }






  public widget_dragstart(schema: WidgetSchema, event: DragEvent): void {
    event.dataTransfer!.setDragImage(new Image(), 0, 0);
    // event.dataTransfer!.setData('json/widget', schema.type!);
    this.dataTransferService.result = false;
    this.dataTransferService.setText('json/widget', schema.type!);
    // 设置所有widget不接受鼠标事件
    this.canvas.children.map(e => {
      const element = e.location.nativeElement as HTMLElement;
      element.style.pointerEvents = 'none';
    });
    this.editor.execute(new SelectionFillCommand(this.editor, []));
  }


  public widget_dragend(schema: WidgetSchema, event: DragEvent): void {
    // 设置所有widget接受鼠标事件
    this.dataTransferService.setText('json/widget', null);
    this.canvas.children.map(e => {
      const element = e.location.nativeElement as HTMLElement;
      element.style.pointerEvents = '';
    });
    if (!this.dataTransferService.result) {
      this.editor.history.undoAndRemoveLast(SelectionFillCommand);
    }

  }

  private generateName(baseName: string): string {
    let i = 1;
    // eslint-disable-next-line no-constant-condition
    while (true) {
      const name = baseName + i.toString();
      if (this.canvas.findWidgetByName(name) == null) return name;
      i++;
    }
  }

  private generateId(): string {
    // eslint-disable-next-line no-constant-condition
    while (true) {
      const id = HmiMath.randomString(6);
      if (this.canvas.findWidgetById(id) == null) return id;
    }
  }



  public widget_dblClick(schema: WidgetSchema, event: MouseEvent): void {
    const widgetSchema = this.provider.getType(schema.type!);
    const defaultValue = widgetSchema!.component.prototype.metaData.default as WidgetDefaultVlaues;
    const configure: WidgetConfigure = {
      id: this.generateId(),
      name: this.generateName(schema.name),
      type: schema.type!,
      interval: defaultValue.interval,
      zIndex: this.editor.canvas.children.length,
      style: ObjectUtil.clone(defaultValue.style)!,
      data: ObjectUtil.clone(defaultValue.data)!,
      rect: {
        left: Math.floor(this.canvas.scrollViewer.nativeElement.scrollLeft / this.canvas.zoomScale),
        top: Math.floor(this.canvas.scrollViewer.nativeElement.scrollTop / this.canvas.zoomScale),
        width: defaultValue.size.width,
        height: defaultValue.size.height
      },
      events: ObjectUtil.clone(defaultValue.events)!,
    };
    const compRef = this.canvas.parseComponent(configure);
    if (compRef) {
      this.editor.execute(new WidgetAddCommand(this.editor, [compRef], true));
    }
  }






}