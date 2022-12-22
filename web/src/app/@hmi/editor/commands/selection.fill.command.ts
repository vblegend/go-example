import { ComponentRef } from "@angular/core";
import { BasicWidgetComponent } from "@hmi/components/basic-widget/basic.widget.component";
import { HmiEditorComponent } from "../hmi.editor.component";
import { BasicCommand } from "./basic.command";

/**
 * 编辑器选区对象填充命令\
 * 将编辑器的选区内所有选中对象全部取消选中，选中当前对象数组
 */
export class SelectionFillCommand extends BasicCommand {

    protected oldObjects: ComponentRef<BasicWidgetComponent>[];
    public objects!: ComponentRef<BasicWidgetComponent>[];



    constructor(editor: HmiEditorComponent, objects: ComponentRef<BasicWidgetComponent>[]) {
        super(editor);
        this.oldObjects = editor.selection.objects;
        this.objects = objects;
    }

    public execute(): void {
        this.editor.selection.fill(this.objects);
    }

    public undo(): void {
        this.editor.selection.fill(this.oldObjects);
    }

    public update(cmd: BasicCommand): void {

    }

}