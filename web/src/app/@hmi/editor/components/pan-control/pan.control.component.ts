import { Component, ElementRef, HostBinding, HostListener, Injector, Input, OnInit, ViewChild, ViewContainerRef } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';
import { DisignerCanvasComponent } from '../disigner-canvas/disigner.canvas.component';


@Component({
  selector: 'hmi-pan-control',
  templateUrl: './pan.control.component.html',
  styleUrls: ['./pan.control.component.less']
})
/**
 * 移动工具
 */
export class PanControlComponent extends GenericComponent {
  private scrollViewer!: HTMLDivElement;
  private buttonDown!: boolean;
  private disX!: number;
  private disY!: number;


  /**
   *
   */
  constructor(protected injector: Injector, private canvas: DisignerCanvasComponent) {
    super(injector);
    this.buttonDown = false;
  }

  protected onInit(): void {
    this.scrollViewer = this.viewContainerRef.element.nativeElement.parentElement;
  }



  /**
   * get component left px
   * binding host position
   */
  @HostBinding('style.width')
  public get width(): string {
    return `${this.scrollViewer.scrollWidth}px`;
  }

  /**
   * get component top px
   * binding host position
   */
  @HostBinding('style.height')
  public get height(): string {
    return `${this.scrollViewer.scrollHeight}px`;
  }









  @HostListener('mousedown@outside', ['$event'])
  public onMouseDown(ev: MouseEvent): void {
    if (ev.buttons == 1) {
      this.buttonDown = true;
      this.disX = ev.clientX;
      this.disY = ev.clientY;
    }
    ev.stopPropagation();
    ev.preventDefault();
  }

  @HostListener('document:mousemove@outside', ['$event'])
  public onMouseMove(ev: MouseEvent): void {
    if (this.buttonDown) {
      const scrollWidth = this.scrollViewer.scrollWidth - this.scrollViewer.clientWidth;
      const scrollHeight = this.scrollViewer.scrollHeight - this.scrollViewer.clientHeight;
      let left = -(ev.clientX - this.disX);
      let top = -(ev.clientY - this.disY);
      this.scrollViewer.scrollLeft += left;
      this.scrollViewer.scrollTop += top;
      if (left < 0) left = 0;
      if (top < 0) top = 0;
      if (left > scrollWidth) left = scrollWidth;
      if (top > scrollHeight) top = scrollHeight;
      this.disX = ev.clientX;
      this.disY = ev.clientY;
    }
    ev.preventDefault();
    ev.stopPropagation();
  }

  @HostListener('document:mouseup@outside', ['$event'])
  public onMouseUp(ev: MouseEvent): void {
    if (this.buttonDown) {
      this.buttonDown = false;
    }
    ev.preventDefault();
    ev.stopPropagation();
  }

  /**
   * 屏蔽右键菜单
   * @param ev 
   */
  @HostListener('contextmenu@outside', ['$event'])
  public onContextMenu(ev: MouseEvent): void {
    ev.preventDefault();
    ev.stopPropagation();
  }


}