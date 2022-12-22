import { Injectable, Type } from "@angular/core";
import { Action } from "@core/common/delegate";
import { AnyObject } from "@core/common/types";
import { BasicWidgetComponent } from "@hmi/components/basic-widget/basic.widget.component";
import { ViewCanvasComponent } from "@hmi/components/view-canvas/view.canvas.component";
import { MethodMeta } from "../../configuration/widget.meta.data";

/**
 * 小部件的事件服务\
 * 用于小部件之间的事件联动通讯
 */
@Injectable({
  providedIn: 'root'
})
export class WidgetEventService {
  private canvas!: ViewCanvasComponent;

  /**
   * 初始化 事件服务所属canvas
   * @param canvas 
   */
  public initCanvas$(canvas: ViewCanvasComponent): void {
    this.canvas = canvas;
  }

  /**
   * 派遣一个事件至目标对象
   * @param sender 发送人
   * @param receiver 接收人，为null时广播给所有小部件对象
   * @param method 事件数据
   * @param params 事件参数
   */
  public dispatch(sender: BasicWidgetComponent, receiver: string | undefined, method: string, params: AnyObject): void {
    const targets = this.getEventTargets(receiver, sender);
    for (let i = 0; i < targets.length; i++) {
      const comp = targets[i].instance;
      try {
        this.eventHandle(comp, method, params);
      }
      catch (e) {
        console.error(e);
      }
    }
  }

  /**
   * 获取事件接收人
   * @param receiver 
   * @param sender 
   * @returns 
   */
  private getEventTargets(receiver: string | undefined, sender: BasicWidgetComponent) {
    if (receiver == null) return this.canvas.children.filter(e => e.instance != sender);
    const children = this.canvas.children;
    const comp = children.find(e => e.instance.configure.id === receiver);
    if (comp == null) return [];
    return [comp];
  }



  /**
   * 事件处理
   * @param receiver 接收人
   * @param method 事件方法名
   * @param params 事件参数
   */
  private eventHandle(receiver: BasicWidgetComponent, method: string, params: AnyObject): void {
    if (method == null) return;
    const methodName = method.split(':').pop()!;
    const methodMeta = receiver.metaData.interface[methodName];
    const methodFunc = receiver.getMemberFunction(methodName);
    if (methodName && methodFunc) {
      const args = this.getEventParams(methodMeta, params);
      if (args != null) methodFunc.apply(receiver, args);
    }
  }


  /**
   * 验证/获取事件参数
   * @param method 事件元数据
   * @param params 调用参数 
   * @returns 
   */
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  private getEventParams(method: MethodMeta, params: any): AnyObject[] | undefined {
    const args: AnyObject[] = [];
    for (let i = 0; i < method!.args!.length; i++) {
      const arg = method.args![i];
      let value = params[arg.argName];
      if (arg.typed && value != null) {
        if (arg.typed === 'boolean') {
          value = value === true || value === 'true' || value === 1 || value === '1';
        } else if (arg.typed === 'number') {
          value = Number(value);
        } else if (arg.typed === 'array') {
          value = JSON.parse(value);
        }
      }
      // 严格模式下，参数不匹配则跳出
      if (method.strict && value === undefined) return undefined;
      args[i] = value;
    }
    return args;
  }



}
