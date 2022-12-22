import { Component, ElementRef, HostListener, Injector, OnInit, QueryList, ViewChild, ViewChildren, ViewContainerRef } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';
import { ViewCanvasComponent } from '@hmi/components/view-canvas/view.canvas.component';
import { HmiEditorComponent } from '@hmi/editor/hmi.editor.component';
import { HmiViewerComponent } from '@hmi/hmi.viewer.component';
import { HmiEditorService } from '@hmi/services/hmi.editor.service';
import { CustomPropertiesComponent } from '../dashboard/custom.properties/custom.properties.component';

@Component({
  selector: 'ngx-examples',
  templateUrl: './examples.component.html',
  styleUrls: ['./examples.component.less'],
  providers: [HmiEditorService]
})
export class ExamplesComponent extends GenericComponent {

  // @ViewChild('viewer', { static: true }) viewer!: HmiViewerComponent;

  @ViewChildren('viewer') viewers!: QueryList<HmiViewerComponent>;

  public portal_Closed(): void {
    if (this.showPortal) this.showPortal = false;
  }


  public showPortal: Boolean = false;

  public show_window(): void {
    this.showPortal = !this.showPortal;
  }



  public edit_click(): void {
    const viewer = this.viewers.first;
    const json = viewer.toJson();
    const editor = this.hmiEditorService.open(HmiEditorComponent);
    editor.instance.templates.install(CustomPropertiesComponent);
    editor.instance.loadFromJson(json, false);
    editor.instance.onSave.subscribe(e => {
      viewer.loadFromJson(e);
    });
  }


  /**
   *
   */
  constructor(protected injector: Injector, private hmiEditorService: HmiEditorService) {
    super(injector);
  }

}
