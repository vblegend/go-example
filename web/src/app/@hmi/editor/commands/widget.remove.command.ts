import { ComponentRef } from "@angular/core";
import { BasicWidgetComponent } from "@hmi/components/basic-widget/basic.widget.component";
import { HmiEditorComponent } from "../hmi.editor.component";
import { BasicCommand } from "./basic.command";


export class WidgetRemoveCommand extends BasicCommand {
    protected selecteds: boolean[];
    public objects: ComponentRef<BasicWidgetComponent>[];


    constructor(editor: HmiEditorComponent, objects: ComponentRef<BasicWidgetComponent>[]) {
        super(editor);
        this.objects = objects;
        this.selecteds = [];
        for (const obj of this.objects) {
            this.selecteds.push(this.editor.selection.contains(obj));
        }
    }

    public execute(): void {
        for (let i = 0; i < this.objects.length; i++) {
            this.editor.canvas.remove(this.objects[i]);
            if (this.selecteds[i]) {
                this.editor.selection.remove(this.objects[i]);
            }
        }
    }

    public undo(): void {
        for (let i = 0; i < this.objects.length; i++) {
            this.editor.canvas.add(this.objects[i]);
            if (this.selecteds[i]) {
                this.editor.selection.add(this.objects[i]);
            }
        }
    }

    public update(cmd: BasicCommand): void {

    }

}