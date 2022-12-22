import { Component, Injector, Input } from '@angular/core';
import { BasicPropertyComponent } from '@hmi/editor/components/basic-property/basic.property.component';

/**
 * 属性为Number的对象属性
 * ```
 * <hmi-number-slider-property [max]="10000" [min]="0" [precision]="0" [step]="1" unit="秒" [nullValue]="0">
 * </hmi-number-slider-property>
 * ```
 * @max **可以设置的最大值**
 * @min **可以设置的最小值**
 * @precision **数值精度位数**
 * @step **每次调整的步进量**
 * @unit **属性单位，为空则不显示**
 * @nullValue **当属性值为null时，属性框内默认使用该值**
 * @diffValue **当支持多选时，对象属性不统一时则使用该值**
 */
@Component({
  selector: 'hmi-number-slider-property',
  templateUrl: './number.slider.property.component.html',
  styleUrls: ['./number.slider.property.component.less'],
})
export class NumberSliderPropertyComponent extends BasicPropertyComponent<number> {

  @Input()
  public max: number = Number.MAX_VALUE;

  @Input()
  public min: number = Number.MIN_VALUE;

  @Input()
  public step: number = 1;

  @Input()
  public unit: string = '';

  public formatter: (value: number) => string;
  /**
   *
   */
  constructor(protected injector: Injector) {
    super(injector);
    this.formatter = this.formatterFunc.bind(this);
    this.nullValue = 0;
  }

  private formatterFunc(value: number): string {
    return (this.unit && this.unit.length) ? `${value} ${this.unit}` : `${value}`;
  }
}
