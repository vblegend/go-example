import { ComponentRef, Injectable, Type } from "@angular/core";
import { WidgetPropertiesService } from "@hmi/editor/services/widget.properties.service";
import { PropertiesTemplatesComponent } from "../components/properties-templates/properties.templates.component";
import { HmiEditorComponent } from "../hmi.editor.component";

interface TemplateTypeInstance {
    template: Type<PropertiesTemplatesComponent>;
    compRef: ComponentRef<PropertiesTemplatesComponent>;
}


/**
 * 属性模板管理器\
 * 用于安装和卸载属性模板
 */
@Injectable({
    providedIn: 'root'
})
export class PropertyTemplateManager {

    private readonly types: TemplateTypeInstance[] = [];
    private readonly preInstalled: Type<PropertiesTemplatesComponent>[] = [];
    /**
     *
     */
    constructor(private readonly editor: HmiEditorComponent) {

    }

    /**
     * 添加一个模板 并不安装
     * @param template 
     */
    public addTemplate(template: Type<PropertiesTemplatesComponent>): void {
        this.preInstalled.push(template);
    }



    /**
     * 添加并安装一个属性模板
     * @param template 
     * @returns 
     */
    public install(template?: Type<PropertiesTemplatesComponent>): void {
        if (template) this.preInstalled.push(template);
        this.installTemplates();
    }



    private installTemplates(): void {
        if (this.preInstalled.length > 0) {
            for (const template of this.preInstalled) {
                if (this.types.some(e => e.template === template)) return;
                const compRef: ComponentRef<PropertiesTemplatesComponent> = this.editor.viewContainerRef.createComponent(template, { injector: this.editor.injector });
                this.types.push({ template, compRef });
            }
            this.preInstalled.length = 0;
        }
    }



    /**
     * 返回当前已安装模板中是否包含模板@template
     * @param template 
     */
    public contains(template: Type<PropertiesTemplatesComponent>): boolean {
        return this.types.some(e => e.template === template);
    }



    /**
     * 卸载属性模板
     * @param template 
     */
    public unInstall(template: Type<PropertiesTemplatesComponent>): void {
        const index = this.types.findIndex(e => e.template === template);
        if (index > -1) {
            const ts = this.types.splice(index, 1);
            const indexRef = this.editor.viewContainerRef.indexOf(ts[0].compRef.hostView);
            this.editor.viewContainerRef.remove(indexRef);
            ts[0].compRef.hostView.detach();
            ts[0].compRef.hostView.destroy();
            ts[0].compRef.destroy();
        }
    }
}