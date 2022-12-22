import { Component, Injector, Input, Type } from '@angular/core';
import { IconFontService } from '@core/services/iconfont.service';
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
  selector: 'hmi-iconfont-select-property',
  templateUrl: './iconfont.select.property.component.html',
  styleUrls: ['./iconfont.select.property.component.less'],
})
export class IconfontSelectPropertyComponent extends BasicPropertyComponent<any> {
  public visible: boolean = false;

  /**
   *
   */
  public constructor(protected injector: Injector, public iconfont: IconFontService) {
    super(injector);
    this.nullValue = '';
  }

  public btn_click(): void {
    this.iconfont.matchIconClassNames();
    this.visible = true;
  }

  public select(icon: string): void {
    this.saveAndUpdate(icon);
    this.visible = false;
  }

  public clear():void {
    this.saveAndUpdate(null);
    this.visible = false;
  }


  public close(): void {
    this.visible = false;
  }



}
