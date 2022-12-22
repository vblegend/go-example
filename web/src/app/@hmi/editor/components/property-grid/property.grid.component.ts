import { Component, Injector } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';
import { HmiEditorComponent } from '@hmi/editor/hmi.editor.component';
import { WidgetPropertiesService } from '@hmi/editor/services/widget.properties.service';

@Component({
  selector: 'hmi-property-grid',
  templateUrl: './property.grid.component.html',
  styleUrls: ['./property.grid.component.less']
})
/**
 * 橡皮筋套选工具
 */
export class PropertyGridComponent extends GenericComponent {

  constructor(protected injector: Injector, public editor: HmiEditorComponent, public properties: WidgetPropertiesService) {
    super(injector);
  }
}