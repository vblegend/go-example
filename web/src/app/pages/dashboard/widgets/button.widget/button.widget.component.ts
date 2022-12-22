import { Component, HostBinding, HostListener, Injector } from '@angular/core';
import { BasicWidgetComponent } from '@hmi/components/basic-widget/basic.widget.component';
import { CommonWidgetPropertys, DefaultData, DefaultEvents, DefaultInterval, DefaultSize, DefaultStyle, DefineEvents, DefineProperties } from '@hmi/editor/core/common';
import { ButtonWidgetDataModel } from './data.model';

@Component({
  selector: 'app-button-widget',
  templateUrl: './button.widget.component.html',
  styleUrls: ['./button.widget.component.less']
})
@DefineEvents([
  { event: 'click', eventName: '单击事件', eventParams: [] },
  { event: 'mouseEnter', eventName: '鼠标悬停', eventParams: [] },
  { event: 'mouseLeave', eventName: '鼠标离开', eventParams: [] }
])
@DefineProperties([...CommonWidgetPropertys, 'data.text', 'data.icon', 'data.align'])
// 默认数据
@DefaultSize(86, 32)
@DefaultStyle({ border: '0px' })
@DefaultInterval(2)
@DefaultEvents({})
@DefaultData<ButtonWidgetDataModel>({ text: 'Button', icon: 'grace-duankailianjie', align: 'center' })
export class ButtonWidgetComponent extends BasicWidgetComponent {

  public data!: ButtonWidgetDataModel;

  public constructor(injector: Injector) {
    super(injector)
  }

  @HostListener('click@outside', ['$event'])
  public onClick(ev: MouseEvent): void {
    this.dispatchEvent('click', {});
  }
  /**
   * 鼠标移入事件
   * @param ev 
   */
  @HostListener('mouseenter@outside', ['$event'])
  public onMouseEnter(ev: MouseEvent): void {
    this.dispatchEvent('mouseEnter', {});
  }

  /**
   * 鼠标移出事件
   * @param ev 
   */
  @HostListener('mouseleave@outside', ['$event'])
  public onMouseLeave(ev: MouseEvent): void {
    this.dispatchEvent('mouseLeave', {});
  }

  /**
   * get component line-height px
   * binding host position
   */
  @HostBinding('style.line-height')
  public get lineHeight$(): string | undefined {
    return `${this.configure.rect!.height}px`;
  }

  /**
   * get component text-align
   * binding host position
   */
  @HostBinding('style.text-align')
  public get textAlign$(): string | undefined {
    return this.data.align;
  }
}
