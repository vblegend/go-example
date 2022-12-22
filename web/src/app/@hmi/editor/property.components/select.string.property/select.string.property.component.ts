import { Component, Injector, Input } from '@angular/core';
import { nzSelectItem } from '@core/common/types';
import { BasicPropertyComponent } from '@hmi/editor/components/basic-property/basic.property.component';


/**
 * 属性为字符串的 Select列表对象
 * ```
 * <hmi-select-string-property [options]="[{ label: '小红', value: 'XiaoHong' },{ label: '小明', value: 'XiaoMing' }]"
 *    nullValue="XiaoHong" diffValue="XiaoHong">
 * </hmi-select-string-property>
 * ```
 * @options **Select 数据列表**
 * @nullValue **当属性值为null时，属性框内默认使用该值**
 * @diffValue **当支持多选时，对象属性不统一时则使用该值**
 */
@Component({
  selector: 'hmi-select-string-property',
  templateUrl: './select.string.property.component.html',
  styleUrls: ['./select.string.property.component.less'],
})
export class SelectStringPropertyComponent extends BasicPropertyComponent<string> {

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
