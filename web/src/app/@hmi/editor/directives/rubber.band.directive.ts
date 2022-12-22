import { Directive, ElementRef, HostListener, Input, Output, EventEmitter, OnInit, Optional, ViewContainerRef, ViewRef, ComponentRef, Injector, ComponentFactoryResolver } from '@angular/core';
import { BaseDirective } from '@core/directives/base.directive';
import { BasicCommand } from '@hmi/editor/commands/basic.command';
import { SelectionFillCommand } from '@hmi/editor/commands/selection.fill.command';
import { SelectionToggleCommand } from '@hmi/editor/commands/selection.toggle.command';
import { BasicWidgetComponent } from '@hmi/components/basic-widget/basic.widget.component';
import { DisignerCanvasComponent } from '@hmi/editor/components/disigner-canvas/disigner.canvas.component';
import { RubberbandComponent } from '@hmi/editor/components/rubber-band/rubber.band.component';
import { Rectangle } from '@hmi/editor/core/common';
import { HmiEditorComponent } from '@hmi/editor/hmi.editor.component';
import { HmiMath } from '@hmi/utility/hmi.math';

@Directive({
    selector: '[rubber-band]'
})

export class RubberBandDirective extends BaseDirective {

    private rectComponent: ComponentRef<RubberbandComponent>;
    /**
     * 橡皮筋所选区域窗口坐标
     */
    private rubberBandArea: Rectangle;

    /**
     * 橡皮筋所选中的视图区域
     */
    private selectionArea: Rectangle;


    private widgets: ComponentRef<BasicWidgetComponent>[] = [];

    private buttonDown = false;
    private startX: number = 0;
    private startY: number = 0;
    private endX: number = 0;
    private endY: number = 0;
    constructor(protected injector: Injector,private editor: HmiEditorComponent, private canvas: DisignerCanvasComponent) {
        super(injector);
        const componentFactory = this.componentFactoryResolver.resolveComponentFactory<RubberbandComponent>(RubberbandComponent);
        this.rectComponent = this.viewContainerRef.createComponent<RubberbandComponent>(componentFactory, undefined, this.injector);
        this.rectComponent.hostView.detach();
        this.rubberBandArea = { left: 0, top: 0, width: 0, height: 0 };
        this.selectionArea = { left: 0, top: 0, width: 0, height: 0 };
    }

    @HostListener('mousedown', ['$event'])
    public onMouseDown(ev: MouseEvent): void {
        if (ev.buttons === 1 || ev.buttons == 2) {
            // 过滤滚动条上的点击事件
            const rect = this.element.getBoundingClientRect();
            if (ev.clientX - rect.left > this.element.clientWidth ||
                ev.clientY - rect.top > this.element.clientHeight) return;
            this.buttonDown = true;
            // 仅限左键更改指针,右键为菜单项 不修改指针样式
            if (ev.buttons == 1) this.element.style.cursor = 'crosshair';
            this.endX = this.startX = (ev.clientX - rect.left);
            this.endY = this.startY = (ev.clientY - rect.top);
            // 更新画布上所有的小部件
            this.widgets = this.canvas.children;
            this.widgets.sort((a, b) => a.instance.zIndex! - b.instance.zIndex!);
            this.updatePosition();
            this.viewContainerRef.insert(this.rectComponent.hostView);
            ev.preventDefault();
            ev.stopPropagation();
        }
    }

    /**
     * 一个逃离Anugular 数据检测框架之外的方法用于处理mousemove事件
     * 具体查看 OutSideEventPluginService  的 @outside
     * @param ev 
     */
    @HostListener('document:mousemove@outside', ['$event'])
    public onMouseMove(ev: MouseEvent): void {
        if (this.buttonDown && ev.buttons === 1) {
            const rect = this.element.getBoundingClientRect();
            this.endX = ev.clientX - rect.left;
            this.endY = ev.clientY - rect.top;
            this.updatePosition();
            ev.preventDefault();
            ev.stopPropagation();
        }
    }


    @HostListener('document:mouseup', ['$event'])
    public onMouseUp(ev: MouseEvent): void {
        if (this.buttonDown) {
            this.element.style.cursor = '';
            this.buttonDown = false;
            // const rect = this.element.getBoundingClientRect();
            this.element.style.cursor = '';
            this.updateSelectionArea(this.startX === this.endX && this.startY === this.endY, ev.ctrlKey);
            const index = this.viewContainerRef.indexOf(this.rectComponent.hostView);
            if (index >= 0) {
                this.viewContainerRef.detach(index);
            }
            this.rectComponent.instance.updateRectangle({ left: 0, top: 0, width: 0, height: 0 });
            ev.preventDefault();
            ev.stopPropagation();
            this.canvas.focus();
        }

    }


    /**
     * 更新橡皮筋等坐标
     */
    private updatePosition() {
        // 限制橡皮筋的坐标区域防止溢出
        this.startX = Math.max(Math.min(this.startX, this.element.clientWidth), 0);
        this.startY = Math.max(Math.min(this.startY, this.element.clientHeight), 0);
        this.endX = Math.max(Math.min(this.endX, this.element.clientWidth), 0);
        this.endY = Math.max(Math.min(this.endY, this.element.clientHeight), 0);
        // 更新橡皮筋
        const scale = this.editor.canvas.zoomScale;


        const left = Math.min(this.endX, this.startX);
        const top = Math.min(this.endY, this.startY);
        const width = Math.abs(this.endX - this.startX);
        const height = Math.abs(this.endY - this.startY);


        this.rubberBandArea = {
            left: Math.min(this.endX, this.startX),
            top: Math.min(this.endY, this.startY),
            width: Math.abs(this.endX - this.startX),
            height: Math.abs(this.endY - this.startY)
        };
        this.rectComponent.instance.updateRectangle(this.rubberBandArea);
        // 更新选中区域
        this.selectionArea = {
            left: left / scale + this.element.scrollLeft / scale,
            top: top / scale + this.element.scrollTop / scale,
            width: width / scale,
            height: height / scale
        };
    }



    /**
     * 更新Selection选择器
     */
    private updateSelectionArea(isClick: boolean, ctrlKey: boolean) {
        const selecteds: ComponentRef<BasicWidgetComponent>[] = [];
        let command: BasicCommand | null = null;
        if (isClick) {
            for (let i = this.widgets.length - 1; i >= 0; i--) {
                if (HmiMath.checkRectangleCross(this.widgets[i].instance.configure.rect, this.selectionArea)) {
                    selecteds.push(this.widgets[i]);
                    break;
                }
            }
        } else {
            for (let i = this.widgets.length - 1; i >= 0; i--) {
                if (HmiMath.checkRectangleCross(this.widgets[i].instance.configure.rect, this.selectionArea)) {
                    selecteds.push(this.widgets[i]);
                }
            }
        }
        // 分组过滤选中
        const groupIds = Array.from(new Set(selecteds.filter(e => e.instance.configure.group != null).map(e => e.instance.groupId)));
        if (groupIds.length > 0) {
            const result = this.widgets.filter(e => e.instance.groupId != null && selecteds.indexOf(e) == -1 && groupIds.indexOf(e.instance.groupId) > -1);
            if (result.length > 0) selecteds.push(...result);
        }
        if (ctrlKey) {
            command = new SelectionToggleCommand(this.editor, selecteds);
        } else {
            command = new SelectionFillCommand(this.editor, selecteds);
        }
        this.editor.execute(command);
    }







}