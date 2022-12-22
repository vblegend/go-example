import { Component, HostBinding, HostListener, Injector, ViewChild, ViewContainerRef } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';
import { HmiEditorComponent } from '@hmi/editor/hmi.editor.component';
import { DisignerCanvasComponent } from '../disigner-canvas/disigner.canvas.component';
import { NzDropdownMenuComponent } from 'ng-zorro-antd/dropdown';
import { SelectionFillCommand } from '@hmi/editor/commands/selection.fill.command';
import { AnchorPosition } from '@hmi/editor/directives/resize.anchor.directive';

@Component({
  selector: 'hmi-selection-area',
  templateUrl: './selection.area.component.html',
  styleUrls: ['./selection.area.component.less']
})
export class SelectionAreaComponent extends GenericComponent {
  @ViewChild('ChildrenView', { static: true, read: ViewContainerRef })
  public container!: ViewContainerRef;
  public AnchorPosition: any = AnchorPosition;


  /**
   *
   */
  constructor(protected injector: Injector, public canvas: DisignerCanvasComponent, public editor: HmiEditorComponent) {
    super(injector);
  }


  public contextMenu($event: MouseEvent, menu: NzDropdownMenuComponent): void {
    if (!this.canvas.ignoreContextMenu) {
      this.contextMenuService.create($event, menu);
      $event.preventDefault();
      $event.stopPropagation();
    }
  }

  public closeMenu(): void {
    this.contextMenuService.close();
  }



  @HostListener('document:keydown', ['$event'])
  public onKeyDown(event: KeyboardEvent): void {
    if (!(event.target instanceof HTMLDivElement)) return;
    switch (event.code) {
      case 'Escape':
        if (this.editor.selection.length === 0) return;
        this.contextMenuService.close();
        this.editor.execute(new SelectionFillCommand(this.editor, []));
        // 往上微调  留空
        event.preventDefault();
        event.stopPropagation();
        break;
      case 'ArrowUp':
        if (this.editor.selection.length === 0) return;
        // 往上微调  留空
        event.preventDefault();
        event.stopPropagation();
        break;
      case 'ArrowDown':
        if (this.editor.selection.length === 0) return;
        // 往下微调  留空
        event.preventDefault();
        event.stopPropagation();
        break;
      case 'ArrowLeft':
        if (this.editor.selection.length === 0) return;
        // 往左微调  留空
        event.preventDefault();
        event.stopPropagation();
        break;
      case 'ArrowRight':
        if (this.editor.selection.length === 0) return;
        // 往右微调  留空
        event.preventDefault();
        event.stopPropagation();
        break;
    }
  }

  protected onDestroy(): void {

  }

  /**
   * host的绑定数据，不可修改。
   */
  @HostBinding('style.position')
  public readonly CONST_DEFAULT_HOST_POSITION_VALUE: string = 'absolute';


  /**
   * get component left px
   * binding host position
   */
  @HostBinding('style.left')
  public get left(): string {
    return `${this.editor.selection.bounds.left}px`;
  }

  /**
   * get component top px
   * binding host position
   */
  @HostBinding('style.top')
  public get top(): string {
    return `${this.editor.selection.bounds.top}px`;
  }

  /**
   * get component width px
   * binding host position
   */
  @HostBinding('style.width')
  public get width(): string {
    return `${this.editor.selection.bounds.width}px`;
  }

  /**
   * get component height px
   * binding host position
   */
  @HostBinding('style.height')
  public get height(): string {
    return `${this.editor.selection.bounds.height}px`;
  }

  /**
   * get/set zIndex
   * 当按下ctrl时，当前组件置于最底端，。
   */
  @HostBinding('style.zIndex')
  public get zIndex(): number {
    return (this.canvas && this.canvas.ctrlPressed) ? SelectionAreaComponent.MIN_ZINDEX : SelectionAreaComponent.MAX_ZINDEX;
  }

  /**
   * 用于控制框选组件是否可见
   */
  @HostBinding('style.display')
  public get display(): string {
    return this.editor.selection.bounds.height > 0 && this.editor.selection.bounds.width > 0 ? '' : 'none';
  }


  public static readonly MAX_ZINDEX: number = 999999;
  public static readonly MIN_ZINDEX: number = -999999;

}