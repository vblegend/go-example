import { Component, Injector, Input, Type } from '@angular/core';
import { IconFontService } from '@core/services/iconfont.service';
import { FileUtil } from '@core/util/file.util';
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
  selector: 'hmi-image-select-property',
  templateUrl: './image.select.property.component.html',
  styleUrls: ['./image.select.property.component.less'],
})
export class ImageSelectPropertyComponent extends BasicPropertyComponent<any> {
  public visible: boolean = false;

  public images: string[] = [];
  /**
   *
   */
  public constructor(protected injector: Injector) {
    super(injector);
    this.nullValue = '';
    this.images.push('/assets/images/login.bg.jpeg');

    this.images.push('/assets/images/team.png');

    this.images.push('/assets/images/test.png');

    this.images.push('/assets/images/texture.png');

  }

  public btn_click(): void {
    this.visible = true;
  }


  public async upload(): Promise<void> {
    const files = await FileUtil.selectLocalFile('.jpg,.png,.gif,.bmp', true);

    console.log(files);

  }


  public select(icon: string): void {
    this.saveAndUpdate(icon);
    this.visible = false;
  }

  public clear(): void {
    this.saveAndUpdate(null);
    this.visible = false;
  }


  public close(): void {
    this.visible = false;
  }



}
