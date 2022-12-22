import { NgModule } from '@angular/core';


import { PagesComponent } from './pages.component';
// import { DashboardModule } from './dashboard/dashboard.module';
import { PagesRoutingModule } from './pages-routing.module';
import { NzLayoutModule } from 'ng-zorro-antd/layout';
import { NzMenuModule } from 'ng-zorro-antd/menu';
// import { ButtonsComponent } from './buttons/buttons.component';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { WelcomeComponent } from './welcome/welcome.component';
import { LayoutModule } from '@layouts/layout.module';
import { CoreModule } from '@core/core.module';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzFormModule } from 'ng-zorro-antd/form';

// import { BrowserModule } from '@angular/platform-browser';
import { CommonModule } from '@angular/common';
import { NzPageHeaderModule } from 'ng-zorro-antd/page-header';
import { NzTabsModule } from 'ng-zorro-antd/tabs';
import { NzStatisticModule } from 'ng-zorro-antd/statistic';
import { NzDescriptionsModule } from 'ng-zorro-antd/descriptions';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzPopconfirmModule } from 'ng-zorro-antd/popconfirm';
import { NzDropDownModule } from 'ng-zorro-antd/dropdown';
import { NzInputModule } from 'ng-zorro-antd/input';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NzDividerModule } from 'ng-zorro-antd/divider';
// import { DragDropModule } from '@angular/cdk/drag-drop';
import { HMI_COMPONENT_SCHEMA_DECLARES, HmiModule } from 'app/@hmi/hmi.module';
import { WidgetSchemaService } from '@hmi/services/widget.schema.service';
import { ExamplesComponent } from './examples/examples.component';
import { CodeingComponent } from './codeing/codeing.component';
import { IFrameComponent } from './iframe/iframe.component';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    // TranslationModule.forRoot({ root: './i18n/' }),
    HmiModule.forRoot(),
    LayoutModule.forRoot(),
    NzIconModule,
    ReactiveFormsModule,
    CoreModule,
    NzMenuModule,
    NzLayoutModule,
    PagesRoutingModule,
    NzTableModule,
    NzFormModule,
    NzPageHeaderModule,
    NzTabsModule,
    NzStatisticModule,
    NzDescriptionsModule,
    NzButtonModule,
    NzPopconfirmModule,
    NzDropDownModule,
    NzInputModule,
    NzDividerModule,
    // DragDropModule,
  ],
  declarations: [
    PagesComponent,
    WelcomeComponent,
    ExamplesComponent,
    CodeingComponent,
    IFrameComponent
  ],
  // providers: [
  //   {
  //     provide: WidgetSchemaService,
  //     useClass: HmiSchemaService,
  //   }
  // ]
})
export class PagesModule {



}
