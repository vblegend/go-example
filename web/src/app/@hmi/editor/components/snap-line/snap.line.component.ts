import { Component, ElementRef, HostBinding, Injector, Input, OnInit, ViewChild, ViewContainerRef } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';
import { Rectangle, Vector2 } from '@hmi/editor/core/common';
import { DisignerCanvasComponent } from '../disigner-canvas/disigner.canvas.component';


@Component({
  selector: 'hmi-snap-line',
  templateUrl: './snap.line.component.html',
  styleUrls: ['./snap.line.component.less']
})
/**
 * 橡皮筋套选工具
 */
export class SnapLineComponent extends GenericComponent {

  @Input() position!: Vector2;
  /**
   *
   */
  constructor(protected injector: Injector,canvas: DisignerCanvasComponent) {
    super(injector);
  }




  
  /**
   * get component left px
   * binding host position
   */
  @HostBinding('style.left')
  public get $left(): string {
    return `${this.position.x}px`;
  }

  /**
   * get component top px
   * binding host position
   */
  @HostBinding('style.top')
  public get $top(): string {
    return `${this.position.y}px`;
  }






}