import { Component } from '@angular/core';
import { PropertiesTemplatesComponent } from '@hmi/editor/components/properties-templates/properties.templates.component';
import { DefaultButtonActionFactory } from '../common';

/***
 * 小部件基础属性模板列表\
 * 用于定义小部件属性的模板
 * html 
 * ```
 * <ng-template let-table="close" let-group="dismiss">
 *     <hmi-property-element [multiple]="false" attributePath="name" title="部件名称" tooltip="小部件的名称，用于识别">
 *         <hmi-text-input-property nullValue="">
 *         </hmi-text-input-property>
 *     </hmi-property-element>
 * </ng-template>
 * ```
 */
@Component({
  selector: 'hmi-editor-basic-properties',
  templateUrl: './basic.properties.component.html'
})
export class BasicPropertiesComponent extends PropertiesTemplatesComponent {
  public readonly DefaultButtonActionFactory = DefaultButtonActionFactory;



  
}