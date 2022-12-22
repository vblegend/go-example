import { Component, Host, Injector, Input, QueryList, ViewChildren } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';
import { PropertieDefineTemplateDirective } from '@hmi/editor/directives/properties.template.directive';
import { HmiEditorComponent } from '@hmi/editor/hmi.editor.component';
import { WidgetPropertiesService } from '@hmi/editor/services/widget.properties.service';

/***
 * 小部件属性模板列表\
 * 用于定义小部件属性的模板
 * html 
 * ```
 * <ng-template defineTemplate table="属性" group="基础属性" index="1">
 *     <hmi-property-element [multiple]="false" attributePath="name" title="部件名称" tooltip="小部件的名称，用于识别">
 *         <hmi-text-input-property nullValue="">
 *         </hmi-text-input-property>
 *     </hmi-property-element>
 * </ng-template>
 * ```
 */
@Component({
  selector: 'hmi-properties-templates',
  template: ''
})
export abstract class PropertiesTemplatesComponent extends GenericComponent {

  /**
   * 模板中所有声明的 指令
   */
  @ViewChildren(PropertieDefineTemplateDirective)
  public directives!: QueryList<PropertieDefineTemplateDirective>;




  /**
   *
   */
  constructor(protected injector: Injector, protected provider: WidgetPropertiesService, public editor: HmiEditorComponent) {
    super(injector);
  }

  protected onAfterViewInit(): void {
    if (this.provider && this.directives && this.directives.length > 0) {
      this.provider.register(this);
    }
  }


  protected onDestroy(): void {
    if (this.provider && this.directives && this.directives.length > 0) {
      this.provider.unRegister(this);
    }
  }




}