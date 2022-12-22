import { ComponentRef } from "@angular/core";
import { BasicWidgetComponent } from "@hmi/components/basic-widget/basic.widget.component";
import { HmiEditorComponent } from "../hmi.editor.component";
import { BasicCommand } from "./basic.command";


export class WidgetAddCommand extends BasicCommand {
    protected unselObjects!: ComponentRef<BasicWidgetComponent>[];
    public objects: ComponentRef<BasicWidgetComponent>[];

    constructor(editor: HmiEditorComponent, objects: ComponentRef<BasicWidgetComponent>[], private selected: boolean) {
        super(editor);
        this.objects = objects;
        if (this.selected) {
            this.unselObjects = this.editor.selection.objects;
        }
    }

    public execute(): void {
        for (let i = 0; i < this.objects.length; i++) {
            this.editor.canvas.add(this.objects[i]);
        }
        if (this.selected) {
            this.editor.selection.fill(this.objects);
        }
    }

    public undo(): void {
        for (let i = 0; i < this.objects.length; i++) {
            this.editor.canvas.remove(this.objects[i]);
        }
        if (this.selected) {
            this.editor.selection.fill(this.unselObjects);
        }
    }

    public update(cmd: BasicCommand): void {

    }
}