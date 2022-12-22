import { Component, Injector, Input } from '@angular/core';
import { nzSelectItem } from '@core/common/types';
import { BasicPropertyComponent } from '@hmi/editor/components/basic-property/basic.property.component';

/**
 * 属性为数字的 Select列表对象
 * ```
 * <hmi-select-number-property [options]="[{ label: '开启', value: 0 },{ label: '关闭', value: 1 }]"
 *    [nullValue]="1" [diffValue]="0">
 * </hmi-select-number-property>
 * ```
 * @options **Select 数据列表**
 * @nullValue **当属性值为null时，属性框内默认使用该值**
 * @diffValue **当支持多选时，对象属性不统一时则使用该值**
 * 
 */
@Component({
  selector: 'hmi-select-number-property',
  templateUrl: './select.number.property.component.html',
  styleUrls: ['./select.number.property.component.less'],
})
export class SelectNumberPropertyComponent extends BasicPropertyComponent<number> {

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
