import { Injectable, Injector, TemplateRef } from "@angular/core";
import { PropertiesTemplatesComponent } from "@hmi/editor/components/properties-templates/properties.templates.component";
import { PropertyElementComponent } from "@hmi/editor/components/property-element/property.element.component";


export interface PropertiesGroup {
    /**
     * 分组名称
     */
    name: string;
    /**
     * 分组内的模板
     */
    templates: TemplateRef<PropertyElementComponent>[];

}
export interface PropertiesTable {
    /**
     * Table 名称
     */
    name: string;
    /**
     * 分组列表
     */
    groups: PropertiesGroup[];
    /**
     * 未分组内容列表
     */
    templates: TemplateRef<PropertyElementComponent>[];
}




/**
 * 小部件 部件属性注册服务
 */
@Injectable({
    providedIn: 'root'
})
export class WidgetPropertiesService {
    /**
     *
     */
    constructor(protected injector: Injector) {
    }


    /**
     * 可用部件属性
     */
    public readonly tables: PropertiesTable[] = [];


    private _loaded: Function[] = [];

    /**
     * 从模板注册部件属性
     * @param properties 
     */
    public register(properties: PropertiesTemplatesComponent) {
        const directives = properties.directives.toArray();
        for (const directive of directives) {
            let table = this.tables.find(e => e.name == directive.tableName);
            if (table == null) {
                table = { name: directive.tableName, groups: [], templates: [] };
                this.tables.push(table);
            }
            if (directive.groupName) {
                let group = table.groups.find(e => e.name == directive.groupName);
                if (group == null) {
                    group = { name: directive.groupName, templates: [] };
                    table.groups.push(group!);
                }
                group.templates.push(directive.templateRef);
            } else {
                table.templates.push(directive.templateRef);
            }
        }
    }

    /**
     * 从模板注销部件属性
     */
    public unRegister(properties: PropertiesTemplatesComponent): void {
        const directives = properties.directives.toArray();
        for (const directive of directives) {
            let table = this.tables.find(e => e.name == directive.tableName);
            if (table == null) continue;
            if (directive.groupName) {
                let group = table.groups.find(e => e.name == directive.groupName);
                if (group == null) continue;
                this.removeTemplate(group.templates, directive.templateRef);
            } else {
                this.removeTemplate(table.templates, directive.templateRef);
            }
        }
    }

    private removeTemplate(templates: TemplateRef<PropertyElementComponent>[], templateRef: TemplateRef<PropertyElementComponent>): void {
        const index = templates.indexOf(templateRef);
        if (index > -1) {
            templates.splice(index, 1);
        }
    }



}