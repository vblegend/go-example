import { ModuleWithProviders, NgModule, Provider } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CoreModule } from '@core/core.module';

import { ViewCanvasComponent } from './components/view-canvas/view.canvas.component';

import { HmiEditorComponent } from './editor/hmi.editor.component';
import { WidgetSchemaService } from './services/widget.schema.service';

import { DisignerCanvasComponent } from './editor/components/disigner-canvas/disigner.canvas.component';
import { DisignerHotkeysDirective } from './editor/directives/disigner.hotkeys.directive';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';
import { SelectionAreaComponent } from './editor/components/selection-area/selection.area.component';
import { RubberbandComponent } from './editor/components/rubber-band/rubber.band.component';
import { NzDropDownModule } from 'ng-zorro-antd/dropdown';
import { Component } from '@angular/core';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzTabsModule } from 'ng-zorro-antd/tabs';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzToolTipModule } from 'ng-zorro-antd/tooltip';
import { NzSelectModule } from 'ng-zorro-antd/select';
import { NzInputNumberModule } from 'ng-zorro-antd/input-number';
import { NzSliderModule } from 'ng-zorro-antd/slider';
import { NzSpaceModule } from 'ng-zorro-antd/space';
import { NzSwitchModule } from 'ng-zorro-antd/switch';
import { NzElementPatchModule } from 'ng-zorro-antd/core/element-patch';
import { NzDrawerModule } from 'ng-zorro-antd/drawer';


import { SnapLineComponent } from './editor/components/snap-line/snap.line.component';
import { PanControlComponent } from './editor/components/pan-control/pan.control.component';
import { HmiViewerComponent } from './hmi.viewer.component';
import { AngularSplitModule } from 'angular-split';
import { ObjectListComponent } from './editor/components/object-list/object.list.component';
import { WidgetListComponent } from './editor/components/widget-list/widget.list.component';

import { DragPreviewComponent } from './editor/components/drag-preview/drag.preview.component';
import { PropertyGridComponent } from './editor/components/property-grid/property.grid.component';
import { TextInputPropertyComponent } from './editor/property.components/text.input.property/text.input.property.component';
import { NumberInputPropertyComponent } from './editor/property.components/number.input.property/number.input.property.component';
import { PropertyElementComponent } from './editor/components/property-element/property.element.component';
import { SelectColorPropertyComponent } from './editor/property.components/select.color.property/select.color.property.component';
import { NumberSliderPropertyComponent } from './editor/property.components/number.slider.property/number.slider.property.component';
import { SelectBooleanPropertyComponent } from './editor/property.components/select.boolean.property/select.boolean.property.component';
import { SelectStringPropertyComponent } from './editor/property.components/select.string.property/select.string.property.component';
import { SelectNumberPropertyComponent } from './editor/property.components/select.number.property/select.number.property.component';
import { BasicPropertiesComponent } from './editor/basic.properties/basic.properties.component';
import { WidgetPropertiesService } from './editor/services/widget.properties.service';
import { PropertieDefineTemplateDirective } from './editor/directives/properties.template.directive';
import { ButtonActionPropertyComponent } from './editor/property.components/button.action.property/button.action.property.component';
import { DataTransferService } from './editor/services/data.transfer.service';
import { WidgetEventComponent } from './editor/property.components/common.event.property/widget.event.component';
import { HmiEditorService } from './services/hmi.editor.service';
import { SwitchBooleanPropertyComponent } from './editor/property.components/switch.boolean.property/switch.boolean.property.component';
import { EditorToolbarComponent } from './editor/components/toolbar/toolbar.component';

import { NzDividerModule } from 'ng-zorro-antd/divider';
import { NzFormModule } from 'ng-zorro-antd/form';
import { ReSizeAnchorDirective } from './editor/directives/resize.anchor.directive';
import { MoveAnchorDirective } from './editor/directives/move.anchor.directive';
import { ZoomControlDirective } from './editor/directives/zoom.control.directive';
import { RubberBandDirective } from './editor/directives/rubber.band.directive';
import { WidgetDragDirective } from './editor/directives/widget.drag.directive';
import { IconfontSelectPropertyComponent } from './editor/property.components/iconfont.select.property/iconfont.select.property.component';
import { ImageSelectPropertyComponent } from './editor/property.components/image.select.property/image.select.property.component';

export declare const HMI_COMPONENT_SCHEMA_DECLARES: WidgetSchemaService;

/**
 * services
 */
const PROVIDERS: Provider[] = [
  WidgetPropertiesService,
  DataTransferService,
  HmiEditorService
];


const EXPORT_PIPES: Provider[] = [

];


const EXPORT_DIRECTIVES: Provider[] = [
  ReSizeAnchorDirective,
  MoveAnchorDirective,
  ZoomControlDirective,
  RubberBandDirective,
  DisignerHotkeysDirective,
  WidgetDragDirective,
  PropertieDefineTemplateDirective
  // DataPropertyDirective
];


/**
 * EXPORT CONPONENTS
 */
const EXPORT_COMPONENTS = [
  ViewCanvasComponent,
  DisignerCanvasComponent,
  HmiEditorComponent,
  HmiViewerComponent,
  SelectionAreaComponent,
  RubberbandComponent,
  SnapLineComponent,
  PanControlComponent,
  ObjectListComponent,
  WidgetListComponent,
  DragPreviewComponent,
  PropertyGridComponent,
  PropertyElementComponent,
  BasicPropertiesComponent,
  EditorToolbarComponent,
  // propertys
  WidgetEventComponent,
  TextInputPropertyComponent,
  NumberInputPropertyComponent,
  SelectColorPropertyComponent,
  NumberSliderPropertyComponent,
  SelectBooleanPropertyComponent,
  SelectStringPropertyComponent,
  SelectNumberPropertyComponent,
  ButtonActionPropertyComponent,
  SwitchBooleanPropertyComponent,
  IconfontSelectPropertyComponent,
  ImageSelectPropertyComponent
];

/**
 *  Dynamic Module
 */
@NgModule({
  imports: [
    CommonModule,
    NzElementPatchModule,
    RouterModule,
    FormsModule,
    ReactiveFormsModule,
    CoreModule,
    NzDropDownModule,
    AngularSplitModule,
    NzTabsModule,
    NzInputModule,
    NzButtonModule,
    NzIconModule,
    NzToolTipModule,
    NzSelectModule,
    NzInputNumberModule,
    NzSliderModule,
    NzSpaceModule,
    NzSwitchModule,
    NzDividerModule,
    NzFormModule,
    NzDrawerModule,
    // 
  ],
  exports: [
    EXPORT_COMPONENTS,
    EXPORT_DIRECTIVES,
    EXPORT_PIPES
  ],
  declarations: [
    EXPORT_COMPONENTS,
    EXPORT_DIRECTIVES,
    EXPORT_PIPES
  ]
})
export class HmiModule {

  public static forRoot(): ModuleWithProviders<HmiModule> {
    return {
      ngModule: HmiModule,
      providers: [
        ...PROVIDERS
      ]
    };
  }
}
