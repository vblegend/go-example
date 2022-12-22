
import { ChangeDetectionStrategy, Component, ComponentRef, ElementRef, Host, HostListener, Injector, Input, OnChanges, OnInit, SimpleChanges, ViewChild, ViewContainerRef } from '@angular/core';
import { Vector2 } from '@hmi/editor/core/common';
import { WidgetSchemaService } from '@hmi/services/widget.schema.service';
import { NzDropdownMenuComponent } from 'ng-zorro-antd/dropdown';
import { HmiEditorComponent } from '../../hmi.editor.component';
import { ViewCanvasComponent } from '../../../components/view-canvas/view.canvas.component';
import { SelectionAreaComponent } from '../selection-area/selection.area.component';
import { WidgetEventService } from '@hmi/editor/services/widget.event.service';
import { TimerPoolService } from '@core/services/timer.pool.service';


@Component({
  selector: 'hmi-disigner-canvas',
  templateUrl: './disigner.canvas.component.html',
  styleUrls: ['./disigner.canvas.component.less'],
  // changeDetection: ChangeDetectionStrategy.OnPush,
  providers: [
    /* 为每个canvas 生成单独的管理服务 */
    WidgetEventService,
    TimerPoolService
  ]
})
export class DisignerCanvasComponent extends ViewCanvasComponent {

  @ViewChild('ChildrenView', { static: true, read: ViewContainerRef })
  public container!: ViewContainerRef;

  @ViewChild('scrollViewer', { static: true })
  public scrollViewer!: ElementRef<HTMLDivElement>;

  @ViewChild('selectionArea', { static: true })
  public selectionArea!: SelectionAreaComponent;

  public ignoreContextMenu?: boolean;
  /**
   * 设置/获取 视图的缩放倍率
   */
  public zoomScale: number;
  private _ctrlPressed!: boolean;
  public spaceKeyDown!: boolean;

  /**
   * 水平对齐线的位置
   */
  public readonly hSnapLines: Vector2[] | undefined[] = [undefined, undefined, undefined];

  /**
   * 垂直对齐线位置
   */
  public readonly vSnapLines: Vector2[] | undefined[] = [undefined, undefined, undefined];


  /**
   * 隐藏/重置 所有对齐辅助线
   */
  public hideSnapLines(): void {
    this.hSnapLines[0] = undefined;
    this.vSnapLines[0] = undefined;
    this.hSnapLines[1] = undefined;
    this.vSnapLines[1] = undefined;
    this.hSnapLines[2] = undefined;
    this.vSnapLines[2] = undefined;
  }


  public get isEditor(): boolean {
    return true;
  }
  public get ctrlPressed(): boolean {
    return this._ctrlPressed;
  }


  /**
   *
   */
  constructor(protected injector: Injector, public provider: WidgetSchemaService, public editor: HmiEditorComponent) {
    super(injector, provider);
    this.zoomScale = 1;
    this.editor = editor;
  }




  protected onInit(): void {
    super.onInit();
    // this.selectionArea = this.container.createComponent(SelectionAreaComponent);
    this.scrollViewer.nativeElement.setAttribute('tabindex', '0');
  }

  public focus(): void {
    this.scrollViewer.nativeElement.focus();
  }

  @HostListener('document:keydown', ['$event'])
  public onKeybordDown(event: KeyboardEvent): void {
    if (!(event.target instanceof HTMLDivElement)) return;
    if (event.code == 'Control') {
      this._ctrlPressed = true;
    }
    if (event.code == 'Space') {
      if (!this.spaceKeyDown) {
        this.spaceKeyDown = true;
      }
      event.preventDefault();
      event.stopPropagation();
    }
  }


  @HostListener('document:keyup', ['$event'])
  public onKeybordUp(event: KeyboardEvent): void {
    if (!(event.target instanceof HTMLDivElement)) return;
    if (event.code == 'Control') {
      this._ctrlPressed = false;
    }
    if (this.spaceKeyDown && event.code == 'Space') {
      this.spaceKeyDown = false;
      event.preventDefault();
      event.stopPropagation();
    }
  }






  public contextMenu($event: MouseEvent, menu: NzDropdownMenuComponent): void {
    if (!this.ignoreContextMenu) {
      this.contextMenuService.create($event, menu);
      $event.preventDefault();
      $event.stopPropagation();
    }
  }

  public closeMenu(): void {
    this.contextMenuService.close();
  }
}
