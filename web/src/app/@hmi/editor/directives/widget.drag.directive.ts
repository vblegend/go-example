import { ComponentRef, Directive, HostListener, Injector, Input } from '@angular/core';
import { BaseDirective } from '@core/directives/base.directive';
import { ObjectUtil } from '@core/util/object.util';
import { WidgetAddCommand } from '@hmi/editor/commands/widget.add.command';
import { DragPreviewComponent } from '@hmi/editor/components/drag-preview/drag.preview.component';
import { WidgetConfigure } from '@hmi/configuration/widget.configure';
import { WidgetSchema } from '@hmi/configuration/widget.schema';
import { HmiEditorComponent } from '@hmi/editor/hmi.editor.component';
import { HmiMath } from '@hmi/utility/hmi.math';
import { BasicWidgetComponent } from '@hmi/components/basic-widget/basic.widget.component';
import { WidgetSchemaService } from '@hmi/services/widget.schema.service';
import { WidgetDefaultVlaues } from '@hmi/configuration/widget.meta.data';
import { DataTransferService } from '@hmi/editor/services/data.transfer.service';
import { SelectionFillCommand } from '../commands/selection.fill.command';


@Directive({
    selector: '[widget-drag]'
})
/**
 * 编辑器的缩放指令
 * 实现了鼠标滚轮的缩放功能。
 */
export class WidgetDragDirective extends BaseDirective {
    private dragPreview!: ComponentRef<DragPreviewComponent> | null;

    constructor(protected injector: Injector, private provider: WidgetSchemaService, private dataTransferService: DataTransferService, private editor: HmiEditorComponent) {
        super(injector);
    }



    public onInit(): void {

    }

    /**
     * 悬停
     * @param ev 
     */
    @HostListener('dragover', ['$event'])
    public onDragOver(ev: DragEvent): void {
        const typed = this.dataTransferService.getText('json/widget');
        const schema = this.provider.getType(typed)!;
        if (schema && this.dragPreview) {
            const defaultValue = schema.component.prototype.metaData.default as WidgetDefaultVlaues;
            const rect = this.element.getBoundingClientRect();
            const x = (ev.clientX - rect.x) / this.editor.canvas.zoomScale;
            const y = (ev.clientY - rect.y) / this.editor.canvas.zoomScale;
            // 这里应该获取拖拽的小部件默认尺寸，但是。。。。这个事件里的数据是获取不到的。。另想它法
            const width = defaultValue.size.width;
            const height = defaultValue.size.height;
            this.dragPreview.instance.updateRectangle({ left: Math.floor(x - width / 2), top: Math.floor(y - height / 2), width, height });
            ev.preventDefault();
            ev.stopPropagation();
        }
    }

    private state: number = 0;

    /**
     * 拖拽着进来了
     * @param ev 
     */
    @HostListener('dragenter', ['$event'])
    public onDragEnter(ev: DragEvent): void {
        // do some thing
        const target = ev.target as HTMLElement;
        this.state++;
        if (!target.classList.contains('continer')) {
            return;// !target.classList.contains('scrollViewer') &&
        }
        if (this.dragPreview == null) {
            const componentFactory = this.componentFactoryResolver.resolveComponentFactory<DragPreviewComponent>(DragPreviewComponent);
            this.dragPreview = this.editor.canvas.container.createComponent<DragPreviewComponent>(componentFactory, undefined, this.injector);
            this.dragPreview.instance.rect = { left: 0, top: 0, width: 100, height: 100 };
            this.dragPreview.instance.updateRectangle({ left: 0, top: 0, width: 100, height: 100 });
        }
        ev.preventDefault();
        ev.stopPropagation();
    }


    /**
     * 拖拽着离开了
     * @param ev 
     */
    @HostListener('dragleave', ['$event'])
    public onDragLeave(ev: DragEvent): void {
        const target = ev.target as HTMLElement;
        if (!target.classList.contains('continer')) {
            return; //!target.classList.contains('scrollViewer') && 
        }
        this.state--;
        if (this.dragPreview && this.state <= 0) {
            const index = this.editor.canvas.container.indexOf(this.dragPreview.hostView);
            this.editor.canvas.container.remove(index);
            this.dragPreview = null;
        }
        ev.preventDefault();
        ev.stopPropagation();
    }

    /**
     * 放置
     * @param ev 
     */
    @HostListener('drop', ['$event'])
    public onDrop(ev: DragEvent): void {
        this.state = 0;
        this.onDragLeave(ev);
        // do some thing
        const rect = this.element.getBoundingClientRect();
        const x = (ev.clientX - rect.x) / this.editor.canvas.zoomScale;
        const y = (ev.clientY - rect.y) / this.editor.canvas.zoomScale;
        const typed = this.dataTransferService.getText('json/widget');
        const schema = this.provider.getType(typed)!;
        const defaultValue = schema.component.prototype.metaData.default as WidgetDefaultVlaues;
        const configure: WidgetConfigure = {
            id: this.generateId(),
            name: this.generateName(schema.name),
            type: schema.type!,
            zIndex: this.editor.canvas.children.length,
            style: ObjectUtil.clone(defaultValue.style)!,
            data: ObjectUtil.clone(defaultValue.data)!,
            interval: defaultValue.interval,
            rect: {
                left: Math.floor(x - defaultValue.size.width / 2),
                top: Math.floor(y - defaultValue.size.height / 2),
                width: defaultValue.size.width,
                height: defaultValue.size.height
            },
            events: ObjectUtil.clone(defaultValue.events)!
        };
        // 删除取消选中的命令
        this.editor.history.undoAndRemoveLast(SelectionFillCommand);
        const compRef = this.editor.canvas.parseComponent(configure);
        if (compRef) {
            this.editor.execute(new WidgetAddCommand(this.editor, [compRef], true));
        }
        this.dataTransferService.result = true;
        ev.preventDefault();
        ev.stopPropagation();
    }






    private generateName(baseName: string): string {
        let i = 1;
        for (; ;) {
            const name = baseName + i.toString();
            if (this.editor.canvas.findWidgetByName(name) == null) return name;
            i++;
        }
    }

    private generateId(): string {
        for (; ;) {
            const id = HmiMath.randomString(6);
            if (this.editor.canvas.findWidgetById(id) == null) return id;
        }
    }







}