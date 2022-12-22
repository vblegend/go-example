import { Component, Host, HostBinding, Injector, Type } from '@angular/core';
import { AnyObject } from '@core/common/types';
import { GenericComponent } from '@core/components/basic/generic.component';
import { WidgetConfigure, WidgetDataConfigure, WidgetDefaultConfigure, WidgetStyles } from '../../configuration/widget.configure';
import { WidgetDefaultVlaues, WidgetMetaObject } from '@hmi/configuration/widget.meta.data';
import { WidgetEventService } from '@hmi/editor/services/widget.event.service';
import { ObjectUtil } from '@core/util/object.util';
import { ViewCanvasComponent } from '../view-canvas/view.canvas.component';
import { TimerTask } from '@core/common/timer.task';
import { DisignerCanvasComponent } from '../../editor/components/disigner-canvas/disigner.canvas.component';
import { DefaultGlobalWidgetStyle } from '@hmi/configuration/global.default.configure';
import { Params, Rectangle, Vector2, WidgetInterface } from '@hmi/editor/core/common';
import { HmiMath } from '@hmi/utility/hmi.math';


@Component({
  selector: 'app-basic-comp',
  template: '<div></div>',
  styles: []
})
/**
 * 小部件基类，实现了小部件的一些基础服务
 */
export abstract class BasicWidgetComponent extends GenericComponent {
  private _config!: WidgetConfigure;
  @Host()
  private viewParent!: ViewCanvasComponent;
  /**
   * 小部件内部事件处理服务\
   * 由canvas进行隔离\
   */
  private readonly eventService: WidgetEventService;

  /**
   * 部件的默认定时器
   */
  private defaultTimer !: TimerTask;
  /**
   * 小部件的元数据\
   * 
   */
  public get metaData(): WidgetMetaObject {
    if (this.constructor.prototype.METADATA == null) {
      this.constructor.prototype.METADATA = new WidgetMetaObject(this.constructor.prototype);
    }
    return this.constructor.prototype.METADATA;
  }


  /**
   * 小部件的元数据\
   */
  public upgradeWidgetMetaData(typed: Type<BasicWidgetComponent>): WidgetMetaObject {
    let meta = this.constructor.prototype.METADATA as WidgetMetaObject;
    if (meta == null) {
      this.constructor.prototype.METADATA = new WidgetMetaObject(typed);
      return this.constructor.prototype.METADATA;
    }
    if (typed != meta.typed) {
      this.constructor.prototype.METADATA = meta = meta.upgrade(typed);
      // console.log(' need upgrade!');
    }
    return meta;
  }




  /**
   * 获取当前小部件的样式属性
   * @returns 
   */
  public get style(): WidgetStyles {
    return this.configure.style;
  }
  /**
   * 当前是否为编辑模式
   */
  public get isEditMode(): boolean {
    return this.viewParent instanceof DisignerCanvasComponent;
  }

  /**
   * 获取成员函数
   * @param funcName 
   * @returns 
   */
  public getMemberFunction(funcName: string): Function {
    this.ifDisposeThrowException();
    return <Function><unknown>this[funcName as keyof this];
  }


  public get left(): number | undefined {
    return this.configure.rect!.left;
  }

  public get top(): number | undefined {
    return this.configure.rect!.top;
  }

  public get width(): number | undefined {
    return this.configure.rect!.width;
  }

  public get height(): number | undefined {
    return this.configure.rect!.height;
  }

  public get right(): number | undefined {
    return this.configure.rect!.left! + this.configure.rect!.width;
  }

  public get bottom(): number {
    return this.configure.rect!.top! + this.configure.rect!.height;
  }

  public get centerX(): number {
    return this.configure.rect!.left! + this.configure.rect!.width / 2;
  }

  public get centerY(): number | undefined {
    return this.configure!.rect!.top! + this.configure!.rect!.height / 2;
  }


  public get groupId(): number | undefined {
    return this.configure.group;
  }


  /**
   * host的绑定数据，不可修改。
   */
  @HostBinding('style.position')
  public readonly CONST_DEFAULT_HOST_POSITION_VALUE: string = 'absolute';

  /**
   * host的绑定数据，不可修改。
   */
  @HostBinding('style.overflow')
  public readonly CONST_DEFAULT_HOST_OVERFLOW_VALUE: string = 'hidden';

  /**
   * get component left px
   * binding host position
   */
  @HostBinding('style.left')
  public get left$(): string | undefined {
    return `${this.configure.rect!.left}px`;
  }

  /**
   * get component top px
   * binding host position
   */
  @HostBinding('style.top')
  public get top$(): string | undefined {
    return `${this.configure.rect!.top}px`;
  }

  /**
   * get component width px
   * binding host position
   */
  @HostBinding('style.width')
  public get width$(): string | undefined {
    return `${this.configure.rect!.width}px`;
  }

  /**
   * get component height px
   * binding host position
   */
  @HostBinding('style.height')
  public get height$(): string | undefined {
    return `${this.configure.rect!.height}px`;
  }

  /**
   * get component background
   * binding host position
   */
  @HostBinding('style.font-size')
  public get fontSize(): string | undefined {
    return `${this.configure.style.fontSize}px`;
  }

  /**
   * get component background
   * binding host position
   */
  @HostBinding('style.font-weight')
  public get fontWeight(): string | undefined {
    return this.configure.style.fontBold ? 'bold' : '';
  }

  /**
   * get component background
   * binding host position
   */
  @HostBinding('style.color')
  public get color(): string | undefined {
    return this.configure.style.color;
  }

  /**
   * get component opacity
   * binding host position
   */
  @HostBinding('style.opacity')
  public get opacity(): number | undefined {
    return this.configure.style.opacity;
  }

  /**
   * get component zIndex
   * binding host position
   */
  @HostBinding('style.zIndex')
  public get zIndex(): number | undefined {
    return this.configure.zIndex;
  }

  /**
   * get component zIndex
   * binding host position
   */
  @HostBinding('style.border')
  public get border(): string | undefined {
    return this.configure.style.border;
  }

  /**
   * get component zIndex
   * binding host position
   */
  @HostBinding('style.border-radius')
  public get radius(): string | undefined {
    return this.configure.style.radius;
  }

  /**
   * get component transform rotates
   * binding host position
   */
  @HostBinding('style.transform')
  public get transform(): string | undefined {
    if (this.configure.style.rotate) {
      return `rotateZ(${this.configure.style.rotate}deg)`;
    } else {
      return ``;
    }
  }



  /**
   * get component background
   * binding host position
   */
  @HostBinding('style.background-color')
  public get backgroundColor(): string | undefined {
    return this.configure.style.bkColor ?? '';
  }

  @HostBinding('style.background-image')
  public get backgroundImage(): string | undefined {
    if (this.configure.style.bkImage && this.configure.style.bkImage.length > 0) {
      return `url(${this.configure.style.bkImage})`;
    }
    return '';
  }

  /**
   * get component transform rotates
   * binding host position
   */
  @HostBinding('style.background-size')
  public get backgroundSize(): string {
    return this.style.bkSize === 'tile' ? '' : '100% 100%';
  }




  public get configure(): WidgetConfigure {
    return this._config;
  }





  constructor(injector: Injector) {
    super(injector)
    this.eventService = injector.get(WidgetEventService);
  }



  /**
   * 派遣一个事件
   * 使用此功能需要在class对象使用@WidgetEvent 注解
   * @param event 要触发的事件，必须是使用在{@WidgetEvent}列表内的
   * @param params 事件的参数 必须满足声明的事件参数
   */
  protected dispatchEvent<T extends Record<string, any>>(event: string, params: T): void {
    // 编辑模式不触发事件
    if (this.isEditMode) return;
    this.ifDisposeThrowException();
    const _eventMeta = this.metaData.events.find(e => e.event == event);
    if (_eventMeta == null) {
      throw `错误：事件派遣失败，部件“${this.constructor.name}”下未找到事件“${event}”的声明。`;
    }
    for (const key of _eventMeta.eventParams) {
      // eslint-disable-next-line no-prototype-builtins
      if (!params.hasOwnProperty(key)) {
        throw `错误：事件派遣失败，部件“${this.constructor.name}”下,事件“${event}”缺少参数“${key}”。`;
      }
    }

    this.run(() => {
      if (this.configure.events == null) return;
      const eventCfg = this.configure.events[event];
      if (eventCfg == null || eventCfg.length === 0) return;
      for (const cfg of eventCfg) {
        this.eventService.dispatch(this, cfg.target!, cfg.method!, Object.assign({}, params, cfg.params));
      }
    });

  }

  /**
   * 内部函数，禁止重载该方法
   * 保证变量data不可修改
   * @param _data 
   */
  public $initialization(parent: ViewCanvasComponent, _config: WidgetConfigure, _default: WidgetDefaultVlaues): void {
    if (this._config) throw 'This method is only available on first run ';
    this.viewParent = parent;
    this._config = ObjectUtil.clone(_config)!;
    // 升级数据属性
    ObjectUtil.upgrade(this._config.data, _default.data);
    // 升级部件的样式默认值
    ObjectUtil.upgrade(this._config.style, _default.style);
    // 升级全局的样式默认值
    ObjectUtil.upgrade(this._config.style, DefaultGlobalWidgetStyle);


    // 订阅定时器
    if (this.timerPool) {
      this.defaultTimer = this.createTimer(this.onDefaultTimer, this.configure.interval);
      // 在组态编辑器下 挂起默认计时器
      if (this.isEditMode) this.defaultTimer.suspend();
    }
  }


  /**
   * get component data profile
   */
  public data?: WidgetDataConfigure;




  /**
   * 禁止重写该方法\
   * 请重写 @onWidgetInit 方法以实现
   */
  protected readonly onInit = (): void => {
    this.data = this._config.data;
    this.onWidgetInit(this.configure.data);
  };




  /**
   * 部件的初始化事件
   * @param data  部件的绑定数据，使用时请在部件中重写此参数类型
   */
  protected onWidgetInit(data: WidgetDataConfigure): void {

  }



  @WidgetInterface('修改样式', '')
  public setStyleProperties(@Params('color') color?: string, @Params('show', 'boolean') show?: boolean): void {
    if (color) this.viewContainerRef.element.nativeElement.style.color = color;
    if (show != null) this.viewContainerRef.element.nativeElement.style.opacity = show ? 1 : 0;

  }


  /**
   * 默认定时器事件 `本参数运行于outside angular zone`\
   * 默认更新数据后不触发变更检测，需手动调用以下代码\
   * 重写时可选择async异步方法，或普通方法
   * ```
   * this.run(()=>{ 
   *   // some thing
   * });
   * ```
   * 由小部件的interval参数决定触发周期\
   * @param task 计时器任务对象 
   */
  protected async onDefaultTimer(task: TimerTask): Promise<void> {
    // console.log(`${DateUtil.currentDateToString()} [${this.configure.id}]=> ${this.constructor.name}`);
  }

  /**
   * 在编辑状态下的组件非绑定数据更新\
   * 适用于编辑状态下非绑定数据的更新事件通知
   * @param attributePath 更改的属性路径 
   * @param value 属性值 
   */
  protected onDataChanged(attributePath: string[], value: string | boolean | number): void {

  }


  /**
   * 获取小部件在canvas中的矩形区域
   * @returns 
   */
  public getRelativeRect(): Rectangle {
    const rect = this.configure.rect!;
    if (this.configure.style.rotate != null && this.configure.style.rotate != 0) {
      const center: Vector2 = { x: rect.left! + rect.width / 2, y: rect.top! + rect.height / 2 };
      const lt = HmiMath.rotateVector2({ x: rect.left!, y: rect.top! }, center, this.configure.style.rotate);
      const ld = HmiMath.rotateVector2({ x: rect.left!, y: rect.top! + rect.height }, center, this.configure.style.rotate);
      const rt = HmiMath.rotateVector2({ x: rect.left! + rect.width, y: rect.top! }, center, this.configure.style.rotate);
      const rd = HmiMath.rotateVector2({ x: rect.left! + rect.width, y: rect.top! + rect.height }, center, this.configure.style.rotate);
      const left = Math.floor(Math.min(lt.x, ld.x, rt.x, rd.x));
      const top = Math.floor(Math.min(lt.y, ld.y, rt.y, rd.y));
      const right = Math.floor(Math.max(lt.x, ld.x, rt.x, rd.x));
      const bottom = Math.floor(Math.max(lt.y, ld.y, rt.y, rd.y));
      return { left, top, width: right - left, height: bottom - top };
    } else {
      return { left: rect.left, top: rect.top, width: rect.width, height: rect.height };
    }
  }

}
