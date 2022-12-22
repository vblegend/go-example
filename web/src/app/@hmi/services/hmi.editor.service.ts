import { Overlay, OverlayRef } from "@angular/cdk/overlay";
import { ComponentPortal, ComponentType } from "@angular/cdk/portal";
import { ComponentRef, Injectable, Injector, Type } from "@angular/core";
import { GenericComponent } from "@core/components/basic/generic.component";
import { HmiEditorComponent } from "@hmi/editor/hmi.editor.component";


/**
 * 打开HMI编辑器服务\
 * 在overlay层 全屏打开编辑器。
 */
@Injectable({
    providedIn: 'root'
})
export class HmiEditorService {
    private editorRef: ComponentRef<HmiEditorComponent> | null;
    private overlayRef: OverlayRef | null;
    /**
     * 当前编辑器是否处于打开状态
     */
    public get opened(): boolean {
        return this.editorRef != null;
    }

    /**
     *
     */
    constructor(private overlay: Overlay, private injector: Injector) {
        this.editorRef = null;
        this.overlayRef = null;
    }

    public open<TEditor>(editorComponent: Type<TEditor>): ComponentRef<TEditor> {
        if (this.editorRef == null) {
            this.overlayRef = this.overlay.create({
                hasBackdrop: false,
                scrollStrategy: this.overlay.scrollStrategies.noop(),
                positionStrategy: this.overlay.position().global()
            });
            const componentPortal = new ComponentPortal(editorComponent, null, this.injector);
            this.overlayRef.overlayElement.style.zIndex = '1000';
            return this.overlayRef.attach(componentPortal);
        }
        return null!;
    }



    public close(): void {
        if (this.overlayRef) {
            this.overlayRef.detach();
        }
    }



}