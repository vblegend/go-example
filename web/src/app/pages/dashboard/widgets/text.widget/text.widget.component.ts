import { Component, HostBinding, HostListener, Injector } from '@angular/core';
import { BasicWidgetComponent } from '@hmi/components/basic-widget/basic.widget.component';
import { DefaultData, DefineEvents, DefineProperties, CommonWidgetPropertys, DefaultStyle, DefaultInterval, DefaultSize, DefaultEvents } from '@hmi/editor/core/common';
import { TextWidgetDataModel } from './data.model';

@Component({
  selector: 'app-text-widget',
  templateUrl: './text.widget.component.html',
  styleUrls: ['./text.widget.component.less'],

})
@DefineProperties([...CommonWidgetPropertys, 'data.text', 'data.align'])
@DefineEvents([
  { event: 'click', eventName: '鼠标单击', eventParams: ['objectId'] },
  { event: 'mouseEnter', eventName: '鼠标悬停', eventParams: ['objectId'] },
  { event: 'mouseLeave', eventName: '鼠标离开', eventParams: ['objectId'] },
])
// 默认数据
@DefaultSize(80, 24)
@DefaultStyle({})
@DefaultInterval(0)
@DefaultEvents({})
@DefaultData<TextWidgetDataModel>({ text: '单行文本', align: 'center' })
export class TextWidgetComponent extends BasicWidgetComponent {
  public data!: TextWidgetDataModel;


  constructor(injector: Injector) {
    super(injector)


  }

  /**
   * 鼠标移入事件
   * @param ev 
   */
  @HostListener('mouseenter@outside', ['$event'])
  public onMouseEnter(ev: MouseEvent): void {
    this.dispatchEvent('mouseEnter', { objectId: 0 });
  }

  /**
   * 鼠标移出事件
   * @param ev 
   */
  @HostListener('mouseleave@outside', ['$event'])
  public onMouseLeave(ev: MouseEvent): void {
    this.dispatchEvent('mouseLeave', { objectId: 0 });
  }
  
  /**
   * get component text-align
   * binding host position
   */
  @HostBinding('style.text-align')
  public get textAlign$(): string | undefined {
    return this.data.align;
  }

  /**
   * get component line-height px
   * binding host position
   */
  @HostBinding('style.line-height')
  public get lineHeight$(): string | undefined {
    return `${this.configure.rect!.height}px`;
  }

}
