import { Component, Injector, Input, Type } from '@angular/core';
import { DefaultButtonActionFactory, ButtonActionFactory } from '@hmi/editor/common';
import { BasicPropertyComponent } from '@hmi/editor/components/basic-property/basic.property.component';



/**
 * 属性按钮对象属性
 * ```
 * <hmi-button-action-property [factory]="DefaultButtonActionFactory">
 * </hmi-button-action-property>
 * ```
 * @factory **按钮动作工厂，继承自`ButtonActionFactory`**
 */
@Component({
  selector: 'hmi-button-action-property',
  templateUrl: './button.action.property.component.html',
  styleUrls: ['./button.action.property.component.less'],
})
export class ButtonActionPropertyComponent extends BasicPropertyComponent<any> {
  @Input()
  public factory!: Type<ButtonActionFactory>;
  public factoryObject!: ButtonActionFactory;


  /**
   *
   */
  constructor(protected injector: Injector) {
    super(injector);
    this.nullValue = '按钮';
    this.factoryObject = new DefaultButtonActionFactory<any>(injector);
  }

  protected onInit(): void {
    super.onInit();
    if (this.factory != null) {
      this.factoryObject = new this.factory(this.injector);
    }
  }



  public btn_click(): void {
    this.factoryObject.execute(this);
  }

}
