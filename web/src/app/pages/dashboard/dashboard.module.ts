import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { NzLayoutModule } from 'ng-zorro-antd/layout';
import { NzToolTipModule } from 'ng-zorro-antd/tooltip';
import { HmiModule } from '@hmi/hmi.module';
import { CoreModule } from '@core/core.module';
import { NzInputNumberModule } from 'ng-zorro-antd/input-number';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { DashboardRoutingModule } from './dashboard-routing.module';
import { CustomPropertiesComponent } from './custom.properties/custom.properties.component';
import { ViewerComponent } from './viewer/viewer.component';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzSpaceModule } from 'ng-zorro-antd/space';
import { ButtonWidgetComponent } from './widgets/button.widget/button.widget.component';
import { ImageWidgetComponent } from './widgets/image.widget/image.widget.component';
import { SingalWidgetComponent } from './widgets/signal.widget/signal.widget.component';
import { CurveChartWidgetComponent } from './widgets/curvechart.widget/curvechart.widget.component';
import { TextWidgetComponent } from './widgets/text.widget/text.widget.component';
import { NgxEchartsModule } from 'ngx-echarts';


/**
 * widget import ng-zerro-ant modules
 */
const NZMODULES = [
  NzInputNumberModule,
  NzButtonModule,
  NzIconModule,
  NzSpaceModule,
  NzToolTipModule
];

/**
 * widgets declare
 */
const WIDGETS = [
  CustomPropertiesComponent,
  ViewerComponent,
  ButtonWidgetComponent,
  ImageWidgetComponent,
  SingalWidgetComponent,
  CurveChartWidgetComponent,
  TextWidgetComponent,

];

@NgModule({
  declarations: [
    ...WIDGETS,
  ],
  exports: [
    ...WIDGETS,
  ],
  imports: [
    CommonModule,
    NzLayoutModule,
    NgxEchartsModule,
    DashboardRoutingModule,
    CoreModule.forRoot(),
    HmiModule.forRoot(),
    ...NZMODULES,
  ],
  providers: [
    //   {
    //     provide: WidgetSchemaService,
    //     useClass: HmiSchemaService,
    //   }
  ]
})
export class DashboardModule {

  // public static forRoot(): ModuleWithProviders<EditorModule> {
  //   return {
  //     ngModule: EditorModule,
  //     providers: [
  //       // ...PROVIDERS
  //     ]
  //   };
  // }
}
