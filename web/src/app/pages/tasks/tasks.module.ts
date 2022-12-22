import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { TaskListComponent } from './task-list/task-list.component';
import { TasksRoutingModule } from './tasks.routing.module';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzPageHeaderModule } from 'ng-zorro-antd/page-header';
import { NzTabsModule } from 'ng-zorro-antd/tabs';
import { NzStatisticModule } from 'ng-zorro-antd/statistic';
import { NzDescriptionsModule } from 'ng-zorro-antd/descriptions';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzPopconfirmModule } from 'ng-zorro-antd/popconfirm';
import { NzDropDownModule } from 'ng-zorro-antd/dropdown';
import { NzInputModule } from 'ng-zorro-antd/input';
import { FormsModule } from '@angular/forms';
import { NzDividerModule } from 'ng-zorro-antd/divider';
import { CoreModule } from '@core/core.module';
import { TaskAddComponent } from './task-add/task-add.component';
import { NzDrawerModule } from 'ng-zorro-antd/drawer';
import { NzGridModule } from 'ng-zorro-antd/grid';
import { ReactiveFormsModule } from '@angular/forms';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzDatePickerModule } from 'ng-zorro-antd/date-picker';
import { NzTagModule } from 'ng-zorro-antd/tag';
import { NzCheckboxModule } from 'ng-zorro-antd/checkbox';




@NgModule({
  declarations: [
    TaskListComponent,
    TaskAddComponent
  ],
  imports: [
    CommonModule,
    TasksRoutingModule,
    NzTableModule,
    NzFormModule,
    NzPageHeaderModule,
    NzTabsModule,
    CoreModule,
    NzStatisticModule,
    NzDescriptionsModule,
    NzButtonModule,
    NzPopconfirmModule,
    NzDropDownModule,
    NzInputModule,
    NzDividerModule,
    FormsModule,
    NzDrawerModule,
    NzGridModule,
    ReactiveFormsModule,
    NzIconModule,
    NzDatePickerModule,
    NzTagModule,
    NzCheckboxModule
  ]
})
export class TasksModule { }
