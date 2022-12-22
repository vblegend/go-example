import { Component, Injector, Input } from '@angular/core';
import { nzSelectItem } from '@core/common/types';
import { BasicPropertyComponent } from '@hmi/editor/components/basic-property/basic.property.component';


/**
 * 属性为布尔类型的 Select列表对象
 * ```
 * <hmi-select-boolean-property [options]="[{ label: '是', value: true },{ label: '否', value: false }]"
 *    nullValue="false" diffValue="true">
 * </hmi-select-boolean-property>
 * ```
 * @options **Select 数据列表**
 * @nullValue **当属性值为null时，属性框内默认使用该值**
 * @diffValue **当支持多选时，对象属性不统一时则使用该值**
 */
@Component({
  selector: 'hmi-select-boolean-property',
  templateUrl: './select.boolean.property.component.html',
  styleUrls: ['./select.boolean.property.component.less'],
})
export class SelectBooleanPropertyComponent extends BasicPropertyComponent<boolean> {

  @Input()
  public options: nzSelectItem[] = [];

  /**
   *
   */
  constructor(protected injector: Injector) {
    super(injector);
    this.nullValue = undefined;
  }

  protected onInit(): void {
    super.onInit();
  }

}
