import { Component, HostListener, Injector } from '@angular/core';
import { TimerTask } from '@core/common/timer.task';
import { CaptureTimeAsync } from '@core/decorators/capturetime';
import { BasicWidgetComponent } from '@hmi/components/basic-widget/basic.widget.component';
import { DefaultData, Params, DefineEvents, WidgetInterface, DefineProperties, CommonWidgetPropertys, DefaultStyle, DefaultInterval, DefaultSize, DefaultEvents } from '@hmi/editor/core/common';

@Component({
  selector: 'app-image-widget',
  templateUrl: './image.widget.component.html',
  styleUrls: ['./image.widget.component.less']
})
// 定义事件
@DefineEvents([
  { event: 'click', eventName: '单击事件', eventParams: ['deviceId'] },
  { event: 'mouseEnter', eventName: '鼠标移入', eventParams: ['deviceId'] },
  { event: 'mouseLeave', eventName: '鼠标移出', eventParams: ['deviceId'] },
])
@DefineProperties([...CommonWidgetPropertys, 'data.imgSrc'], 'Interval')
// 默认数据
@DefaultSize(200, 100)
@DefaultStyle({ bkImage: './assets/images/test.png' })
@DefaultInterval(0)
@DefaultEvents({})
@DefaultData({})
export class ImageWidgetComponent extends BasicWidgetComponent {


  constructor(injector: Injector) {
    super(injector)
  }



  protected onQueryChanges(): void {
    console.log(this.queryParams);
  }

  protected onDataChanged(attributePath: string[], value: string): void {
    console.log(`reload [${attributePath}] => ${value}`);
  }


  @CaptureTimeAsync()
  protected async onDefaultTimer(task: TimerTask): Promise<void> {
    // 模拟数据加载延迟
    await this.sleep(5000);
  }


  @WidgetInterface('更换图片', '使用URL更换图片内容', true)
  public updateImage(@Params('url') url?: string): void {
    console.log(`触发事件 updateImage => url:${url}`);
  }



  @HostListener('click', ['$event'])
  public onMouseDown(ev: MouseEvent): void {
    this.dispatchEvent('click', { deviceId: 0 });
  }


  /**
   * 鼠标移入事件
   * @param ev 
   */
  @HostListener('mouseenter@outside', ['$event'])
  public onMouseEnter(ev: MouseEvent): void {
    this.dispatchEvent('mouseEnter', { deviceId: 0 });
  }

  /**
   * 鼠标移出事件
   * @param ev 
   */
  @HostListener('mouseleave@outside', ['$event'])
  public onMouseLeave(ev: MouseEvent): void {
    this.dispatchEvent('mouseLeave', { deviceId: 0 });
  }

}
