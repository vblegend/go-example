import { Component, HostBinding, HostListener, Injector } from '@angular/core';
import { TimerTask } from '@core/common/timer.task';
import { CaptureTime, CaptureTimeAsync } from '@core/decorators/capturetime';
import { BasicWidgetComponent } from '@hmi/components/basic-widget/basic.widget.component';
import { WidgetDataConfigure } from '@hmi/configuration/widget.configure';
import { DefaultData, Params, DefineEvents, WidgetInterface, DefaultValue, DefineProperties, CommonWidgetPropertys, DefaultStyle, DefaultInterval, DefaultSize, DefaultEvents } from '@hmi/editor/core/common';
import { SignalWidgetDataModel } from './data.model';

@Component({
  selector: 'app-signal-widget',
  templateUrl: './signal.widget.component.html',
  styleUrls: ['./signal.widget.component.less']
})
@DefineProperties([...CommonWidgetPropertys, 'data.eqid', 'data.sgid', 'data.textformat', 'data.displayunit', 'data.align'])

@DefineEvents([])
// 默认数据
@DefaultSize(86, 22)
@DefaultStyle({})
@DefaultInterval(2)
@DefaultEvents({})
@DefaultData<SignalWidgetDataModel>({ align: 'left' })
export class SingalWidgetComponent extends BasicWidgetComponent {
  public data!: SignalWidgetDataModel;
  public output: string = "0V";


  constructor(injector: Injector) {
    super(injector)
  }

  protected onWidgetInit(data: WidgetDataConfigure): void {

  }

  protected onDestroy(): void {

  }

  @WidgetInterface('更新信号', '刷新部件数据')
  public updateSvg(@Params('stationId') stationId?: number, @Params('roomId') roomId?: number, @Params('deviceId') deviceId?: number): void {
    // console.log(`触发事件 updateSvg => stationId:${stationId}，roomId:${roomId}，deviceId:${deviceId}`);
  }

  @CaptureTimeAsync()
  protected async onDefaultTimer(task: TimerTask): Promise<void> {
    task.run(() => {
      this.output = `${(Math.random() * 40 + 180).toFixed(2)}V`;
    });
  }

  @HostListener('mousedown@outside', ['$event'])
  public onMouseDown(ev: MouseEvent): void {
    this.dispatchEvent('click', {});
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
