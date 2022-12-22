import { Component, Injector, Input, QueryList, ViewChildren } from '@angular/core';
import { NgModel } from '@angular/forms';
import { GenericComponent } from '@core/components/basic/generic.component';
import { GenericAttributeCommand } from '@hmi/editor/commands/generic.attribute.command';
import { WidgetConfigure } from '@hmi/configuration/widget.configure';
import { HmiEditorComponent } from '@hmi/editor/hmi.editor.component';
import { PropertyElementComponent } from '../property-element/property.element.component';
@Component({
  selector: 'app-basic-comp',
  template: '<div></div>',
  styles: []
})

/**
 * 
 */
export abstract class BasicPropertyComponent<TPropertyVlaue> extends GenericComponent {

  protected batchNo?: number | undefined;
  private _attributePath: string = 'data';
  private _attrPaths: string[] = [];
  /**
   * 
   */
  @ViewChildren(NgModel)
  protected modelDirectives: QueryList<NgModel> = new QueryList();


  /**
   * 属性的路径\
   * 填写configure下的层级，层级之间用/分割\
   * 如：\
   * “name”\
   * “rect/left”\
   * “data/deviceId”
   */
  @Input()
  public set attributePath(value: string) {
    this._attributePath = value;
    this._attrPaths = value.split('/');
  }

  public get attributePath(): string {
    return this._attributePath;
  }


  /**
   * 当绑定的数据为null时 使用当前值
   */
  @Input()
  public nullValue: TPropertyVlaue | undefined;


  /**
   * 当选中多个对象时，如果当前对象支持多个对象且对象的属性不同时 则使用该值
   */
  @Input()
  public diffValue: TPropertyVlaue | undefined;

  public editor!: HmiEditorComponent;
  public parent: PropertyElementComponent;
  /**
   *, @Host() private parent: PropertyElementComponent
   */
  constructor(protected injector: Injector) {
    super(injector);
    this.batchNo = undefined;
    this.editor = this.injector.get(HmiEditorComponent);
    this.parent = this.injector.get(PropertyElementComponent);
    this.editor = this.parent.editor;
    this.attributePath = this.parent.attributePath;
  }

  /**
   * 禁止重写此方法\
   * 组件附加至容器后触发\
   * 本方法已 no catch ，方法内触发catch不影响主逻辑\
   */
  protected readonly onAfterViewInit = (): void => {
    for (let i = 0; i < this.modelDirectives.length; i++) {
      const ngModel = this.modelDirectives.get(i);
      const subscription = ngModel!.update.subscribe(e => { this.saveAndUpdate(e); });
      this.managedSubscription(subscription);
    }
  }







  /**
   * 保存变更并更新
   */
  public saveAndUpdate(value: TPropertyVlaue): void {
    const fixValue = this.dataModify_fix(value);
    const obejcts = this.editor.selection.objects.map(e => e.instance.configure);
    const command = new GenericAttributeCommand(this.editor, obejcts, this.attributePath, [fixValue], this.batchNo);
    if (this._attrPaths[0] === 'data') {
      this.editor.selection.objects.map(widget => {
        const onDataChanged = widget.instance.getMemberFunction('onDataChanged');
        onDataChanged.apply(widget.instance, [this._attrPaths, value]);
      });
    }
    this.editor.execute(command);
    this.detectChanges();
  }


  /**
   * 数据修改修复扩展\
   * 此事件发生在用户输入完成后，且在数据命令创建之前。
   * @param value 用户实际输入的数值
   * @returns 返回被修复的数值
   */
  protected dataModify_fix(value: TPropertyVlaue): TPropertyVlaue {
    return value;
  }

  /**
   * 
   * 数据绑定修复扩展\
   * 此事件发生在组件绑定模型数据之前\
   * @param value 绑定模型的实际数据
   * @returns 组件中显示的数据
   */
  protected dataBinding_fix(value: TPropertyVlaue | undefined): TPropertyVlaue | undefined {
    return value != null ? value : this.nullValue;
  }



  /**
   * 获取第一个选中对象的属性
   */
  public get defaultProperty(): TPropertyVlaue | undefined {
    let value = <TPropertyVlaue><unknown>this.configure;
    if (value == null) return undefined;
    for (let i = 0; i < this._attrPaths.length; i++) {
      if (value != null) {
        value = <TPropertyVlaue><unknown>value[this._attrPaths[i] as keyof TPropertyVlaue];
      }
    }
    return this.dataBinding_fix(value);
  }

  /**
   * 获取第一个选中对象的配置文件
   */
  public get configure(): WidgetConfigure | null {
    if (this.editor.selection.objects && this.editor.selection.objects.length > 0) {
      return this.editor.selection.objects[0].instance.configure;
    }
    return null;
  }


}

