import { Component, Injector, Input } from '@angular/core';
import { AnyObject, nzSelectItem } from '@core/common/types';
import { BasicPropertyComponent } from '@hmi/editor/components/basic-property/basic.property.component';



/**
 * 属性为String的颜色类型(#FFFFFF)的 Select列表对象
 * ```
 * <hmi-select-color-property nullValue="transparent">
 * </hmi-select-color-property>
 * ```
 * @colors **自定义颜色列表 Select 数据列表**
 * @nullValue **当属性值为null时，属性框内默认使用该值**
 * @diffValue **当支持多选时，对象属性不统一时则使用该值**
 */
@Component({
  selector: 'hmi-select-color-property',
  templateUrl: './select.color.property.component.html',
  styleUrls: ['./select.color.property.component.less'],
})
export class SelectColorPropertyComponent extends BasicPropertyComponent<string> {

  /**
    * 当绑定的数据为null时 使用当前值
    */
  @Input()
  public nullValue: string | undefined = 'transparent';

  /**
   * 默认颜色表
   */
  @Input()
  public colors: nzSelectItem[] = [];


  constructor(protected injector: Injector) {
    super(injector);
    this.nullValue = 'transparent';
    this.colors = [
      { label: 'transparent', value: 'transparent' },
      { label: 'Red', value: '#FF0000' },
      { label: 'Green', value: '#00FF00' },
      { label: 'Blue', value: '#0000FF' },
      { label: 'Yellow', value: '#FFFF00' },
      { label: '#339999', value: '#339999' },
      { label: 'Red', value: '#FF0000' },
      { label: 'Green', value: '#00FF00' },
      { label: 'Blue', value: '#0000FF' },
      { label: 'Yellow', value: '#FFFF00' },
      { label: '#339999', value: '#339999' },
      { label: 'Red', value: '#FF0000' },
      { label: 'Green', value: '#00FF00' },
      { label: 'Blue', value: '#0000FF' },
      { label: 'Yellow', value: '#FFFF00' },
      { label: '#339999', value: '#339999' },
      { label: 'Red', value: '#FF0000' },
      { label: 'Green', value: '#00FF00' },
      { label: 'Blue', value: '#0000FF' },
      { label: 'Yellow', value: '#FFFF00' },
      { label: '#339999', value: '#339999' },
      { label: 'Red', value: '#FF0000' },
      { label: 'Green', value: '#00FF00' },
      { label: 'Blue', value: '#0000FF' },
      { label: 'Yellow', value: '#FFFF00' },
      { label: '#339999', value: '#339999' },
      { label: 'Red', value: '#FF0000' },
      { label: 'Green', value: '#00FF00' },
      { label: 'Blue', value: '#0000FF' },
      { label: 'Yellow', value: '#FFFF00' },
      { label: '#339999', value: '#339999' },



    ];
  }


  /**
   * 这里透明色要显示为黑色 
   * transparent 会导致 input[type=color] 出警告
   */
  public get colorBoxValue(): string | undefined {
    const value = this.defaultProperty;
    if (value == 'transparent') return '#000000';
    return value;
  }


  /**
   * 预制默认颜色点击应用
   * @param value 
   */
  public colorClick(value: nzSelectItem): void {
    this.saveAndUpdate(value.value);
  }


}
