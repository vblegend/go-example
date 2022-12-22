import { Component, Injector, Input } from '@angular/core';
import { BasicPropertyComponent } from '@hmi/editor/components/basic-property/basic.property.component';

/**
 * 属性为字符串的 Select列表对象
 * ```
 * <hmi-text-input-property nullValue="" diffValue="***">
 * </hmi-text-input-property>
 * ```
 * @maxLength **可以输入文本的最大长度 默认64**
 * @nullValue **当属性值为null时，属性框内默认使用该值**
 * @diffValue **当支持多选时，对象属性不统一时则使用该值**
 */
@Component({
  selector: 'hmi-text-input-property',
  templateUrl: './text.input.property.component.html',
  styleUrls: ['./text.input.property.component.less'],
})
export class TextInputPropertyComponent extends BasicPropertyComponent<string> {

  @Input()
  public maxLength: number = 64;

  /**
   *
   */
  constructor(protected injector: Injector) {
    super(injector);
    this.nullValue = '';
  }


}
