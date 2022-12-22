/* eslint-disable @angular-eslint/use-lifecycle-interface */
import { ChangeDetectorRef, Component, ComponentRef, ElementRef, HostBinding, Injector, Input, OnChanges, OnInit, SimpleChanges, ViewChild, ViewContainerRef } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';
import { BasicCommand } from '@hmi/editor/commands/basic.command';
import { SelectionFillCommand } from '@hmi/editor/commands/selection.fill.command';
import { SelectionToggleCommand } from '@hmi/editor/commands/selection.toggle.command';
import { BasicWidgetComponent } from '@hmi/components/basic-widget/basic.widget.component';
import { DisignerCanvasComponent } from '@hmi/editor/components/disigner-canvas/disigner.canvas.component';

import { HmiEditorComponent } from '@hmi/editor/hmi.editor.component';
import { WidgetSchemaService } from '@hmi/services/widget.schema.service';


@Component({
  selector: 'hmi-object-list',
  templateUrl: './object.list.component.html',
  styleUrls: ['./object.list.component.less']
})
/**
 * 橡皮筋套选工具
 */
export class ObjectListComponent extends GenericComponent {
  @Input()
  public canvas!: DisignerCanvasComponent;

  public searchText: string = '';
  /**
   *
   */
  constructor(protected injector: Injector, protected provider: WidgetSchemaService, public editor: HmiEditorComponent) {
    super(injector);
  }


  public filterWidgets(searchText: string): ComponentRef<BasicWidgetComponent>[] {
    if (searchText === '') return this.editor.canvas.children;
    const regExp = new RegExp(searchText, 'i');
    return this.editor.canvas.children.filter(e => {
      return regExp.test(e.instance.configure.name);
    });
  }


  public widget_click(event: MouseEvent, widget: ComponentRef<BasicWidgetComponent>):void {
    let command: BasicCommand | null = null;
    const selecteds: ComponentRef<BasicWidgetComponent>[] = [widget];
    // 分组过滤选中
    const groupId = widget.instance.configure.group;
    if (groupId) {
      const result = this.editor.canvas.children.filter(e => e.instance.groupId === groupId);
      if (result.length > 0) selecteds.push(...result);
    }
    if (event.ctrlKey) {
      command = new SelectionToggleCommand(this.editor, selecteds);
    } else if (event.shiftKey) {
      //
    } else {
      command = new SelectionFillCommand(this.editor, selecteds);
    }
    this.editor.execute(command);
    this.canvas.selectionArea.detectChanges();
  }

  public trackByName(index: number, value: ComponentRef<BasicWidgetComponent>): any {
    return value.instance.configure.name;
  }

}