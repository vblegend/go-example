import { Component, ContentChild, ContentChildren, ElementRef, HostBinding, Injector, Input, QueryList } from '@angular/core';
import { AnyObject } from '@core/common/types';
import { GenericComponent } from '@core/components/basic/generic.component';
import { BasicPropertyComponent } from '@hmi/editor/components/basic-property/basic.property.component';
import { HmiEditorComponent } from '@hmi/editor/hmi.editor.component';
@Component({
  selector: 'hmi-property-element',
  templateUrl: './property.element.component.html',
  styleUrls: ['./property.element.component.less']
})
/**
 * 字符串属性绑定
 */
export class PropertyElementComponent extends GenericComponent {
  private _attributePath: string = 'data';
  private _attrPaths: string[] = [];
  /**
   * 属性key
   * 用于确定属性的匹配
   */
  @Input()
  public key!: string;
  /**
   * 属性对象的属性名描述
   */
  @Input()
  public header: string = '标题';

  /**
   * 属性名称的toolip
   */
  @Input()
  public tooltip: string = '提示';

  /**
   * 是否支持多选设置
   */
  @Input()
  public multiple!: boolean;

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
   *
   */
  constructor(protected injector: Injector,public editor: HmiEditorComponent) {
    super(injector);
    this.multiple = false;
  }

  protected onInit(): void {

  }

  public ngDisplayed(): boolean {
    if (this.editor.selection.length == 0) return false;
    if (!this.multiple) {
      return this.editor.selection.length == 1;
    }
    return true;
  }

  /***
   * 选中对象是否包含此属性
   */
  public get hasProperty(): boolean {
    if (this.ngDisplayed() === false) return false;
    return !this.editor.selection.objects.some(e => e.instance.metaData.properties.indexOf(this.key) == -1);
  }


  private getPathValue(data: any): any {
    let result = data;
    for (let i = 0; i < this._attrPaths.length; i++) {
      if (result === undefined) return undefined;
      if (result === null) return null;
      result = result[this._attrPaths[i]];
    }
    return result;
  }





}
