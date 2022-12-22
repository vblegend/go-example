import { ComponentRef } from "@angular/core";
import { BasicWidgetComponent } from "@hmi/components/basic-widget/basic.widget.component";
import { HmiEditorComponent } from "../hmi.editor.component";
import { BasicCommand } from "./basic.command";

/**
 * 编辑器选区对象状态切换命令\
 * 在编辑器选取内 切换当前数组的选中状态
 */
export class SelectionToggleCommand extends BasicCommand {
    public objects!: ComponentRef<BasicWidgetComponent>[];

    constructor(editor: HmiEditorComponent, objects: ComponentRef<BasicWidgetComponent>[]) {
        super(editor);
        this.objects = objects;
    }

    public execute(): void {
        this.editor.selection.toggle(this.objects);
    }

    public undo(): void {
        this.editor.selection.toggle(this.objects);
    }

    public update(cmd: BasicCommand): void {

    }
    
}