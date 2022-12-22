import { Component, Injector, Input } from '@angular/core';
import { BasicPropertyComponent } from '@hmi/editor/components/basic-property/basic.property.component';



/**
 * 属性为Number的对象属性
 * ```
 * <hmi-number-input-property [max]="10000" [min]="0" [precision]="0" [step]="1" unit="秒" [nullValue]="0">
 * </hmi-number-input-property>
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
  selector: 'hmi-number-input-property',
  templateUrl: './number.input.property.component.html',
  styleUrls: ['./number.input.property.component.less'],
})
export class NumberInputPropertyComponent extends BasicPropertyComponent<number> {

  @Input()
  public max: number = Number.MAX_VALUE;

  @Input()
  public min: number = Number.MIN_VALUE;

  @Input()
  public step: number = 1;

  @Input()
  public unit: string = '';

  @Input()
  public precision: number = 2;

  public readonly formatter: (value: number) => string = (value: number) => (this.unit && this.unit.length) ? `${value} ${this.unit}` : `${value}`;
  public readonly parser: (value: string) => string = (value: string) => value.replace(this.unit, '');


  /**
   *
   */
  constructor(protected injector: Injector) {
    super(injector);
    this.nullValue = 0;
  }



}
