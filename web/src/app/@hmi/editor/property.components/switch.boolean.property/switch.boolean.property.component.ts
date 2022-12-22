import { Component, Injector, Input } from '@angular/core';
import { nzSelectItem } from '@core/common/types';
import { BasicPropertyComponent } from '@hmi/editor/components/basic-property/basic.property.component';


/**
 * 属性为布尔类型的 switch对象
 * ```
 * <hmi-switch-boolean-property checked="是" unChecked="否" nullValue="false" diffValue="true">
 * </hmi-switch-boolean-property>
 * ```
 * @options **Select 数据列表**
 * @nullValue **当属性值为null时，属性框内默认使用该值**
 * @diffValue **当支持多选时，对象属性不统一时则使用该值**
 */
@Component({
  selector: 'hmi-switch-boolean-property',
  templateUrl: './switch.boolean.property.component.html',
  styleUrls: ['./switch.boolean.property.component.less'],
})
export class SwitchBooleanPropertyComponent extends BasicPropertyComponent<boolean> {


  @Input()
  public checked: string = '是';


  @Input()
  public unChecked: string = '否';




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
