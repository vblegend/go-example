import { Component, ElementRef, Injector, OnInit, ViewChild, ViewContainerRef } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';
import { Rectangle } from '@hmi/editor/core/common';


@Component({
  selector: 'hmi-drag-preview',
  templateUrl: './drag.preview.component.html',
  styleUrls: ['./drag.preview.component.less']
})
/**
 * 橡皮筋套选工具
 */
export class DragPreviewComponent extends GenericComponent {
  // @ViewChild('selectionDiv', { static: true }) selectionDiv: ElementRef<HTMLDivElement>;
  public rect: Rectangle;

  /**
   *
   */
  constructor(protected injector: Injector) {
    super(injector);
    this.rect = { left: 0, top: 0, width: 0, height: 0 };
  }

  /**
   * 不使用绑定的方式而是直接使用nativeElement 是为了提高debug下的体验
   * @param rect 
   */
  public updateRectangle(rect: Rectangle) {
    this.rect = rect;
    this.viewContainerRef.element.nativeElement.style.left = `${this.rect.left}px`;
    this.viewContainerRef.element.nativeElement.style.top = `${this.rect.top}px`;
    this.viewContainerRef.element.nativeElement.style.width = `${this.rect.width}px`;
    this.viewContainerRef.element.nativeElement.style.height = `${this.rect.height}px`;
    this.viewContainerRef.element.nativeElement.style.display = (this.rect.width === 0 && this.rect.height === 0) ? `none` : '';
  }



}