import { Component } from '@angular/core';
import { PropertiesTemplatesComponent } from '@hmi/editor/components/properties-templates/properties.templates.component';
import { SelectDeviceFactory } from '../factorys/SelectDeviceFactory';


/***
 * 小部件基础属性模板列表\
 * 自定义小部件属性的模板
 */
@Component({
  selector: 'widget-custom-properties',
  templateUrl: './custom.properties.component.html'
})
export class CustomPropertiesComponent extends PropertiesTemplatesComponent {
  public readonly SelectDeviceFactory = SelectDeviceFactory;
}