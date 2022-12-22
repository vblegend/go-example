import { Component, Injector, ViewChild } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';
import { HmiEditorComponent } from '@hmi/editor/hmi.editor.component';
import { HmiViewerComponent } from '@hmi/hmi.viewer.component';
import { HmiEditorService } from '@hmi/services/hmi.editor.service';
import { CustomPropertiesComponent } from '../custom.properties/custom.properties.component';


@Component({
  selector: 'app-hmi-viewer',
  styleUrls: ['./viewer.component.less'],
  templateUrl: './viewer.component.html',
})
export class ViewerComponent extends GenericComponent {

  @ViewChild('viewer', { static: true }) viewer!: HmiViewerComponent;

  constructor(injector: Injector, private hmiEditorService: HmiEditorService) {
    super(injector);

  }

  protected onInit(): void {
    if (this.queryParams.get('pageId') == null) {
      this.navigateByUrl('/pages/dashboard', { pageId: 10000, deviceId: 456 });
    }
  }





  public edit_click(): void {
    const json = this.viewer.toJson();
    const editor = this.hmiEditorService.open(HmiEditorComponent);
    editor.instance.templates.install(CustomPropertiesComponent);
    editor.instance.loadFromJson(json, false);
    editor.instance.onSave.subscribe(e => {
      this.viewer.loadFromJson(e);
    });
  }






}
