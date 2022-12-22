import { Directive, HostListener, Injector, Input } from "@angular/core";
import { BaseDirective } from "@core/directives/base.directive";
import { WidgetRemoveCommand } from "@hmi/editor/commands/widget.remove.command";
import { HmiEditorComponent } from "@hmi/editor/hmi.editor.component";


@Directive({
    selector: '[hmi-disigner-hotkey]'
})
/**
 * 快捷键指令
 * 用于在编辑器下快捷键的实现
 */
export class DisignerHotkeysDirective extends BaseDirective {

    /**
     *
     */
    constructor(protected injector: Injector, private editor: HmiEditorComponent) {
        super(injector);

    }


    protected onInit(): void {

    }

    @HostListener('document:keydown', ['$event'])
    public onKeyDown(event: KeyboardEvent): void {
        if ((event.target instanceof HTMLInputElement)) return;
        switch (event.code.toLowerCase()) {
            case 'delete':
                this.editor.executeDeleteCommand();
                event.preventDefault();
                event.stopPropagation();
                break;
            case 'keya':
                if (event.ctrlKey) {
                    this.editor.executeSelectAll();
                    event.preventDefault();
                    event.stopPropagation();
                }
                break;
            case 'keyx':
                if (event.ctrlKey) {
                    this.editor.executeCutCommand();
                    event.preventDefault();
                    event.stopPropagation();
                }
                break;
            case 'keyc':
                if (event.ctrlKey) {
                    this.editor.executeCopyCommand();
                    event.preventDefault();
                    event.stopPropagation();
                }
                break;
            case 'keyv':
                if (event.ctrlKey) {
                    this.editor.executePasteCommand();
                    event.preventDefault();
                    event.stopPropagation();
                }
                break;
            case 'keyz':
                if (event.ctrlKey) {
                    this.editor.executeUndo();
                    event.preventDefault();
                    event.stopPropagation();
                }
                break;
            case 'keyy':
                if (event.ctrlKey) {
                    this.editor.executeRedo();
                    event.preventDefault();
                    event.stopPropagation();
                }
                break;

        }

    }





}