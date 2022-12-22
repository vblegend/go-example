import { Component, HostListener, Injector } from '@angular/core';
import { BasicWidgetComponent } from '@hmi/components/basic-widget/basic.widget.component';
import { WidgetDataConfigure } from '@hmi/configuration/widget.configure';
import { DefaultData, Params, DefineEvents, WidgetInterface, DefaultValue, DefineProperties, CommonWidgetPropertys, DefaultStyle, DefaultInterval, DefaultSize, DefaultEvents } from '@hmi/editor/core/common';
import * as echarts from 'echarts';


@Component({
  selector: 'app-curvechart-widget',
  templateUrl: './curvechart.widget.component.html',
  styleUrls: ['./curvechart.widget.component.less']
})
@DefineProperties([...CommonWidgetPropertys, 'data.deviceId', 'data.signalId', 'data.imgSrc'])
@DefineEvents([])
// 默认数据
@DefaultSize(400, 300)
@DefaultStyle({})
@DefaultInterval(4)
@DefaultEvents({})
@DefaultData({})
export class CurveChartWidgetComponent extends BasicWidgetComponent {
  public options: echarts.EChartsOption = {};
  public updateOptions: echarts.EChartsOption = {};


  constructor(injector: Injector) {
    super(injector)
  }

  protected onWidgetInit(data: WidgetDataConfigure): void {
    this.options = {
      tooltip: {
        trigger: 'axis'
      },
      title: {
        text: 'eChart'
      },
      axisPointer: {
        animation: false
      },
      xAxis: {
        type: 'category',
        data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
        splitLine: {
          show: false
        }
      },
      yAxis: {
        type: 'value',
        boundaryGap: [0, '100%'],
        splitLine: {
          show: false
        }
      },
      series: [
        {
          type: 'line',
          showSymbol: false,
          data: this.randomArray(),
          markPoint: {
            data: [
              { type: 'max', name: 'Max' },
              { type: 'min', name: 'Min' }
            ]
          },
        }
      ]
    };
  }

  @WidgetInterface('设置小部件的显示状态', '通过接口事件刷新图表数据')
  public setWidgetVisible(@Params('visible', 'boolean') visible?: string): void {
    this.viewContainerRef.element.nativeElement.style.opacity = visible ? 1 : 0;
  }

  @WidgetInterface('更新曲线图表数据', '通过接口事件刷新图表数据')
  public updateSvg(@Params('stationId') stationId?: number): void {
    this.updateOptions = { series: [{ data: this.randomArray() }] };
  }



  public randomArray(): number[] {
    const datas: number[] = [];
    for (let i = 0; i < 7; i++) {
      datas.push(Math.floor(Math.random() * 110 + 150));
    }
    return datas;
  }



  protected onDestroy(): void {

  }


  @HostListener('mousedown', ['$event'])
  public onMouseDown(ev: MouseEvent): void {
    this.dispatchEvent('click', { taskId: 123 });
  }


}
