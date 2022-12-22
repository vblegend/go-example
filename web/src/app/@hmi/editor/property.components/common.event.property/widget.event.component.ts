import { ChangeDetectionStrategy, Component, ComponentRef, DoCheck, Injector, OnChanges, SimpleChanges } from '@angular/core';
import { nzSelectItem } from '@core/common/types';
import { ObjectUtil } from '@core/util/object.util';
import { BasicWidgetComponent } from '@hmi/components/basic-widget/basic.widget.component';
import { WidgetEventConfigure } from '@hmi/configuration/widget.configure';
import { EventMeta } from '@hmi/configuration/widget.meta.data';
import { GenericAttributeCommand } from '@hmi/editor/commands/generic.attribute.command';
import { WidgetAttributeCommand } from '@hmi/editor/commands/widget.attribute.command';
import { BasicPropertyComponent } from '@hmi/editor/components/basic-property/basic.property.component';
import { PropertyElementComponent } from '@hmi/editor/components/property-element/property.element.component';
import { HmiEditorComponent } from '@hmi/editor/hmi.editor.component';
import { MetaDataService } from '@hmi/editor/services/meta.data.service';
import { WidgetSchemaService } from '@hmi/services/widget.schema.service';
import { NzNotificationService } from 'ng-zorro-antd/notification';






@Component({
  selector: 'hmi-event-property',
  templateUrl: './widget.event.component.html',
  styleUrls: ['./widget.event.component.less'],
  // changeDetection: ChangeDetectionStrategy.OnPush
})
/**
 * 橡皮筋套选工具
 */
export class WidgetEventComponent extends PropertyElementComponent {
  public selectedEvent: string;
  public readonly metaData: MetaDataService;

  /**
   *
   */
  constructor(protected injector: Injector, protected provider: WidgetSchemaService, editor: HmiEditorComponent) {
    super(injector,editor);
    this.metaData = injector.get(MetaDataService);
    this.selectedEvent = 'click';
  }


  /**
   * 事件的目标对象变更后 清空下目标接口方法属性
   * @param eventName 
   * @param index 
   * @param value 
   */
  public targetChanged(eventName: string, index: number, value: string): void {
    const data = this.object.instance.configure.events[eventName][index]!;
    this.editor.executes(
      new GenericAttributeCommand(this.editor, [data], 'target', [value]),
      new GenericAttributeCommand(this.editor, [data], 'method', [null]),
      new GenericAttributeCommand(this.editor, [data], 'params', [{}])
    );
  }

  /**
   * 目标方法属性变更后 清空下目标接口的参数
   * @param eventName 
   * @param index 
   * @param value 
   */
  public methodChanged(eventName: string, index: number, value: string): void {
    const data = this.object.instance.configure.events[eventName][index]!;
    this.editor.executes(
      new GenericAttributeCommand(this.editor, [data], 'method', [value]),
      new GenericAttributeCommand(this.editor, [data], 'params', [{}])
    );
  }

  /**
   * 事件的参数变更
   * @param eventName 
   * @param index 
   * @param paramName 
   * @param value 
   */
  public paramsChanged(eventName: string, index: number, paramName: string, value: string): void {
    const data = this.object.instance.configure.events[eventName][index].params!;
    this.editor.execute(new GenericAttributeCommand(this.editor, [data], paramName, [value]));
  }




  /**
   * 删除一条事件
   * @param eventName 
   * @param index 
   */
  public deleteEvent(eventName: string, index: number): void {
    const events = ObjectUtil.clone(this.object.instance.configure.events)!;
    const event = events[eventName];
    const value: WidgetEventConfigure[] = [];
    if (event.length == 1) {
      delete events[eventName];
    } else {
      event.splice(index, 1)
    }
    this.editor.execute(new GenericAttributeCommand(this.editor, [this.object.instance.configure], 'events', [events]));
  }


  /**
   * 添加一条事件
   * @param event 
   */
  public addEvent(event: string): void {
    const result = this.object.instance.metaData.events.filter(e => e.event === event);
    if (event == null || result.length == 0) return;
    const events = ObjectUtil.clone(this.object.instance.configure.events)!;
    if (events[event] == null) events[event] = [];
    if (events[event].length >= 3) {
      this.notification.error('失败', '每个事件最大支持3个动作');
      return;
    }
    events[event].push({ method: null, target: null, params: {} });
    this.editor.execute(new WidgetAttributeCommand(this.editor, [this.object], 'configure/events', [events]));
  }

  public randomString(): string {
    return `${Math.random() * Number.MAX_VALUE}`;
  }



  public trackByEvent(index: number, value: EventMeta): string {
    return value.event;
  }

  public trackByValue(index: number, value: any): any {
    return value.value;
  }

  public trackByAny(index: number, value: any): any {
    return value;
  }

  public trackByKey(index: number, value: any): any {
    return value.key;
  }
  public trackByName(index: number, value: any): any {
    return value.name;
  }

  public get object(): ComponentRef<BasicWidgetComponent> {
    return this.editor.selection.objects[0];
  }

  public get events(): Record<string, WidgetEventConfigure[]> {
    return this.editor.selection.objects[0].instance.configure.events;
  }

  public get metaEvents(): EventMeta[] {
    return this.editor.selection.objects[0].instance.metaData.events;
  }

}


