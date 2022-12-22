import { ChangeDetectionStrategy, Component, ComponentRef, ElementRef, Injector, Input, OnInit, ViewChild, ViewContainerRef } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';
import { WidgetConfigure } from '../../configuration/widget.configure';
import { WidgetSchemaService } from '@hmi/services/widget.schema.service';
import { BasicWidgetComponent } from '../basic-widget/basic.widget.component';
import { HmiMath } from '@hmi/utility/hmi.math';
import { CurrentVersion, DocumentMagicCode, Rectangle } from '@hmi/editor/core/common';
import { WidgetEventService } from '@hmi/editor/services/widget.event.service';
import { TimerPoolService } from '@core/services/timer.pool.service';
import { WidgetDefaultVlaues } from '@hmi/configuration/widget.meta.data';
import { MetaDataService } from '@hmi/editor/services/meta.data.service';
import { GraphicConfigure } from '@hmi/configuration/graphic.configure';


@Component({
  selector: 'hmi-view-canvas',
  templateUrl: './view.canvas.component.html',
  styleUrls: ['./view.canvas.component.less'],
  // changeDetection: ChangeDetectionStrategy.OnPush,
  providers: [
    /* 为每个canvas 生成单独的管理服务 */
    WidgetEventService,
    TimerPoolService
  ]
})
export class ViewCanvasComponent extends GenericComponent {
  @ViewChild('ChildrenView', { static: true, read: ViewContainerRef })
  public container!: ViewContainerRef;
  private _children: ComponentRef<BasicWidgetComponent>[];
  private readonly _eventHub: WidgetEventService;
  public readonly metaData: MetaDataService;


  /**
   *
   */
  constructor(protected injector: Injector, public provider: WidgetSchemaService) {
    super(injector);
    this._children = [];
    this._eventHub = injector.get(WidgetEventService);
    this._eventHub.initCanvas$(this);
    this.metaData = injector.get(MetaDataService);
  }


  /**
   * 获取容器内所有组件列表
   */
  public get children(): ComponentRef<BasicWidgetComponent>[] {
    return this._children.slice();
  }


  /**
   * 更新容器的zIndex属性
   * 数值越大越往前
   */
  public updatezIndexs(): void {
    this._children.sort((a, b) => b.instance.zIndex! - a.instance.zIndex!);
  }


  /**
   * 添加一个组件至容器中
   * @param ref 
   * @param index 
   * @returns 
   */
  public add(ref: ComponentRef<BasicWidgetComponent>): ComponentRef<BasicWidgetComponent> {
    const ofIndex = this._children.indexOf(ref);
    if (ofIndex === -1) {
      this.container.insert(ref.hostView);
      this._children.push(ref);
      ref.hostView.reattach();
      this.metaData.add(ref);
      if (ref.instance.zIndex == null) {
        ref.instance.configure.zIndex = this._children.length;
      }
      ref.changeDetectorRef.detectChanges();
    }
    return ref;
  }

  /**
   * 从容器中移除一个组件
   * @param ref 
   * @returns 
   */
  public remove(ref: ComponentRef<BasicWidgetComponent>): ComponentRef<BasicWidgetComponent> {
    const ofIndex = this._children.indexOf(ref);
    if (ofIndex > -1) {
      this._children.splice(ofIndex, 1);
      const v = this.container.indexOf(ref.hostView);
      this.container.detach(v);
      ref.hostView.detach();
      this.metaData.remove(ref);
    }
    return ref;
  }

  /**
   * 获取当前是否为编辑态
   */
  public get isEditor(): boolean {
    return false;
  }



  /**
   * 清理并销毁掉所有组件
   */
  public clear(): void {
    while (this._children.length > 0) {
      const compRef = this._children[0];
      this.remove(compRef);
      compRef.destroy();
    }
  }

  /**
   * 解析一个组件，返回组件对象。
   * 当解析到象失败时返回null
   * @param configure 
   * @returns 
   */
  public parseComponent(configure: WidgetConfigure): ComponentRef<BasicWidgetComponent> | null {
    let componentRef: ComponentRef<BasicWidgetComponent> | null = null;
    const comRef = this.provider.getType(configure.type);
    if (comRef) {
      componentRef = this.container.createComponent<BasicWidgetComponent>(comRef.component, { injector: this.injector });
      if (componentRef && componentRef.instance instanceof BasicWidgetComponent) {
        const v = this.container.indexOf(componentRef.hostView);
        this.container.detach(v);
        componentRef.hostView.detach();
        const widgetSchema = this.provider.getType(configure.type);
        const defaultValue = widgetSchema!.component.prototype.metaData.default as WidgetDefaultVlaues;
        componentRef.instance.$initialization(this, configure, defaultValue);
        // eslint-disable-next-line @typescript-eslint/no-non-null-asserted-optional-chain
        (<any>componentRef)['icon'] = widgetSchema?.icon!;
      }
    }
    if (componentRef == null) this.onError(new Error(`未知的组态类型：${configure.type}.`));
    return componentRef;
  }


  public findWidgetByName(name: string): ComponentRef<BasicWidgetComponent> | null {
    if (name == null) return null;
    return this.children.find(e => e.instance.configure && e.instance.configure.name === name)!;
  }

  public findWidgetById(id: string): ComponentRef<BasicWidgetComponent> | null {
    if (id == null) return null;
    return this.children.find(e => e.instance.configure && e.instance.configure.id === id)!;
  }


  protected onInit(): void {

  }


  protected onDestroy(): void {
    this.clear();
    super.onDestroy();
  }


  /**
   * 获取所有容器的总大小。
   * @returns 
   */
  public getComponentsBound(): Rectangle | null {
    let result: Rectangle | null = null;
    for (const comp of this.children) {
      if (result == null) {
        result = {
          left: comp.instance.configure.rect!.left,
          top: comp.instance.configure.rect!.top,
          width: comp.instance.configure.rect!.width,
          height: comp.instance.configure.rect!.height
        };
      } else {
        result = HmiMath.extendsRectangle(result, comp.instance.configure.rect);
      }
    }
    if (result == null) result = { left: 0, top: 0, width: 0, height: 0 };
    return result;
  }





  /**
   * 获取所有小部件的配置项
   * @returns 
   */
  public getConfigure(): WidgetConfigure[] {
    return this.children.map(e => e.instance.configure);
  }




}
