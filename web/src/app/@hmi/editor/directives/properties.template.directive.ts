import { Directive, Input, TemplateRef } from "@angular/core";
import { PropertyElementComponent } from "../components/property-element/property.element.component";

@Directive({
    selector: "ng-template.[hmi-properties]"
})
export class PropertieDefineTemplateDirective {

    /**
     * 属性所在选项卡名称
     */
    @Input()
    public tableName!: string;

    /**
     * 属性所在选项卡下分组名称
     */
    @Input()
    public groupName!: string;

    /**
     * 属性所在分组显示的顺序ID
     */
    @Input()
    public index!: number;


    constructor(public templateRef: TemplateRef<PropertyElementComponent>) { }
}