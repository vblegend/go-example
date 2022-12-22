import { Directive, ElementRef, HostListener, Input, Output, EventEmitter, OnInit, ViewContainerRef, Injector } from '@angular/core';
import { BaseDirective } from '@core/directives/base.directive';
import { WidgetAttributeCommand } from '@hmi/editor/commands/widget.attribute.command';
import { BasicWidgetComponent } from '@hmi/components/basic-widget/basic.widget.component';
import { SelectionAreaComponent } from '@hmi/editor/components/selection-area/selection.area.component';
import { Position } from '@hmi/configuration/widget.configure';
import { Rectangle, Vector2 } from '@hmi/editor/core/common';
import { HmiEditorComponent } from '@hmi/editor/hmi.editor.component';



export enum AnchorPosition {
    Top = 't',
    RightTop = 'rt',
    Right = 'r',
    RightDown = 'rd',
    Down = 'd',
    LeftDown = 'ld',
    Left = 'l',
    LeftTop = 'lt'
}

@Directive({
    selector: '[resizeAnchor]'
})
/**
 * 组件大小调整指令\
 * 用于在编辑态调整组件大小
 */
export class ReSizeAnchorDirective extends BaseDirective {
    @Input()
    public position!: AnchorPosition;
    private buttonDown!: boolean;
    /* 鼠标按下时的坐标 */
    private pressedPoint!: Vector2;
    /** 被选中对象的实时的矩形位置 */
    private rectRealTime!: Rectangle;
    /** 当鼠标按下任意锚点时采集的当前选中区域矩形位置 */
    private rectOrigin!: Rectangle;
    /* resize 命令的批次号  每次鼠标按下后为一个新的批次 */
    private batchNo!: number;
    /* 选中对象的位置大小在选框内的位置快照 */
    private snapshots: Rectangle[] = [];


    constructor(protected injector: Injector, private editor: HmiEditorComponent) {
        super(injector);
    }

    /**
     * 抓取选中对象在选定区域的位置(0-1)
     */
    private captureComponents(): void {
        this.snapshots = [];
        for (const comp of this.editor.selection.objects) {
            const left = (comp.instance.left! - this.editor.selection.bounds.left!) / this.editor.selection.bounds.width;
            const top = (comp.instance.top! - this.editor.selection.bounds.top!) / this.editor.selection.bounds.height;
            const width = comp.instance.width! / this.editor.selection.bounds.width;
            const height = comp.instance.height! / this.editor.selection.bounds.height;
            this.snapshots.push({ left, top, width, height });
        }
    }

    @HostListener('mousedown', ['$event'])
    public onMouseDown(ev: MouseEvent): void {
        if (ev.button === 0) {
            document.body.style.cursor = this.getCursor();
            this.rectRealTime = { left: 0, top: 0, width: 0, height: 0 };
            this.batchNo = Math.floor(Math.random() * Number.MAX_VALUE);
            this.rectOrigin = {
                left: this.editor.selection.bounds.left,
                top: this.editor.selection.bounds.top,
                width: this.editor.selection.bounds.width,
                height: this.editor.selection.bounds.height
            };
            this.captureComponents();
            this.editor.adsorbService.captureAnchors();
            this.buttonDown = true;
            const scale = this.editor.canvas.zoomScale;
            this.pressedPoint = {
                x: ev.clientX / scale + this.editor.canvas.scrollViewer.nativeElement.scrollLeft,
                y: ev.clientY / scale + this.editor.canvas.scrollViewer.nativeElement.scrollTop
            }
            ev.preventDefault();
            ev.stopPropagation();
        }
    }

    @HostListener('document:mousemove', ['$event'])
    public onMouseMove(ev: MouseEvent): void {
        if (this.buttonDown) {
            this.editor.canvas.hideSnapLines();
            this.rectRealTime.left = this.rectOrigin.left;
            this.rectRealTime.top = this.rectOrigin.top;
            this.rectRealTime.width = this.rectOrigin.width;
            this.rectRealTime.height = this.rectOrigin.height;
            const scale = this.editor.canvas.zoomScale;
            const currentPoint: Vector2 = {
                x: ev.clientX / scale + this.editor.canvas.scrollViewer.nativeElement.scrollLeft,
                y: ev.clientY / scale + this.editor.canvas.scrollViewer.nativeElement.scrollTop
            }
            let xLen = Math.floor(currentPoint.x - this.pressedPoint.x);
            let yLen = Math.floor(currentPoint.y - this.pressedPoint.y);
            if (Number.isNaN(xLen)) return;
            if (Number.isNaN(yLen)) return;

            if (this.position.indexOf(AnchorPosition.Left) > -1) {
                xLen = -Math.min(this.rectRealTime.left!, -xLen);
                this.rectRealTime.left = this.rectOrigin.left! + xLen;
                this.rectRealTime.width = this.rectOrigin.width - xLen;
                const result = this.editor.adsorbService.matchXAxis(this.rectRealTime.left, this.editor.DEFAULT_ADSORB_THRESHOLD);
                if (result != null) {
                    this.rectRealTime.left = result;
                    this.rectRealTime.width = this.rectOrigin.left! + this.rectOrigin.width - result;
                    this.editor.canvas.vSnapLines[2] = { x: this.rectRealTime.left * scale, y: this.editor.canvas.scrollViewer.nativeElement.scrollTop };
                }
            }
            if (this.position.indexOf(AnchorPosition.Top) > -1) {
                yLen = -Math.min(this.rectRealTime.top!, -yLen);
                this.rectRealTime.top = this.rectOrigin.top! + yLen;
                this.rectRealTime.height = this.rectOrigin.height - yLen;
                const result = this.editor.adsorbService.matchYAxis(this.rectRealTime.top, this.editor.DEFAULT_ADSORB_THRESHOLD);
                if (result != null) {
                    this.rectRealTime.top = result;
                    this.rectRealTime.height = this.rectOrigin.top! + this.rectOrigin.height - result;
                    this.editor.canvas.hSnapLines[2] = { x: this.editor.canvas.scrollViewer.nativeElement.scrollLeft, y: this.rectRealTime.top * scale };
                }
            }

            if (this.position.indexOf(AnchorPosition.Right) > -1) {
                this.rectRealTime.width = this.rectOrigin.width + xLen;
                const result = this.editor.adsorbService.matchXAxis(this.rectRealTime.left! + this.rectRealTime.width, this.editor.DEFAULT_ADSORB_THRESHOLD);
                if (result != null) {
                    this.rectRealTime.width = result - this.rectOrigin.left!;
                    this.editor.canvas.vSnapLines[2] = { x: (this.rectRealTime.left! + this.rectRealTime.width) * scale, y: this.editor.canvas.scrollViewer.nativeElement.scrollTop };
                }
            }

            if (this.position.indexOf(AnchorPosition.Down) > -1) {
                this.rectRealTime.height = this.rectOrigin.height + yLen;
                const result = this.editor.adsorbService.matchYAxis(this.rectRealTime.top! + this.rectRealTime.height, this.editor.DEFAULT_ADSORB_THRESHOLD);
                if (result != null) {
                    this.rectRealTime.height = result - this.rectOrigin.top!;
                    this.editor.canvas.hSnapLines[2] = { x: this.editor.canvas.scrollViewer.nativeElement.scrollLeft, y: (this.rectRealTime.top! + this.rectRealTime.height) * scale };
                }
            }
            this.executeResizeCommand();
            ev.preventDefault();
            ev.stopPropagation();
        }

    }

    @HostListener('document:mouseup', ['$event'])
    public onMouseUp(ev: MouseEvent): void {
        if (this.buttonDown) {
            this.buttonDown = false;
            document.body.style.cursor = '';
            this.editor.canvas.hideSnapLines();
            ev.preventDefault();
        }
    }


    /** 获取各个方位的鼠标样式 */
    private getCursor(): string {
        switch (this.position) {
            case AnchorPosition.Left:
                return 'w-resize';
            case AnchorPosition.LeftTop:
                return 'se-resize';
            case AnchorPosition.Top:
                return 's-resize';
            case AnchorPosition.RightTop:
                return 'ne-resize';
            case AnchorPosition.Right:
                return 'w-resize';
            case AnchorPosition.RightDown:
                return 'se-resize';
            case AnchorPosition.Down:
                return 's-resize';
            case AnchorPosition.LeftDown:
                return 'ne-resize';
        }
    }



    /**
     * 执行调整大小命令
     * @param ev 
     * @returns 
     */
    private executeResizeCommand() {
        const attrs: Rectangle[] = [];
        for (let i = 0; i < this.editor.selection.objects.length; i++) {
            const rect: Rectangle = {
                left: this.rectRealTime.left! + this.rectRealTime.width * this.snapshots[i].left!,
                top: this.rectRealTime.top! + this.rectRealTime.height * this.snapshots[i].top!,
                width: this.rectRealTime.width * this.snapshots[i].width,
                height: this.rectRealTime.height * this.snapshots[i].height
            };
            attrs.push(rect);
        }
        this.editor.execute(new WidgetAttributeCommand(this.editor,
            this.editor.selection.objects,
            'configure/rect',
            attrs,
            this.batchNo
        ));
    }










}