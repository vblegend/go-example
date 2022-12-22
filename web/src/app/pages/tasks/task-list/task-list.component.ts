import { Component, OnInit } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';

import { NzSafeAny } from 'ng-zorro-antd/core/types';
import { NzDrawerOptions, NzDrawerRef } from 'ng-zorro-antd/drawer';

import { NzTableFilterFn, NzTableFilterList, NzTableFilterValue, NzTableSortFn, NzTableSortOrder } from 'ng-zorro-antd/table';
import { TaskAddComponent } from '../task-add/task-add.component';
import { SchedulingTask, TaskMode } from '../task-model/tasks.model';



@Component({
  selector: 'ngx-tasks',
  templateUrl: './task-list.component.html',
  styleUrls: ['./task-list.component.less']
})
export class TaskListComponent extends GenericComponent {
  public searchValue: string = '';
  public serviceFilters: NzTableFilterList = [];
  public listOfRandomUser: SchedulingTask[] = [];
  public listOfDisplay: SchedulingTask[] = [];
  public loading: boolean = false;
  public visible: boolean = false;





  protected onInit(): void {
    this.searchValue = '';
    this.visible = false;
    this.loading = false;
    this.loadDataFromServer();
  }


  private async loadDataFromServer(): Promise<void> {
    this.loading = true;
    this.serviceFilters = [{ text: '数据服务', value: 'data.service' }, { text: '网关服务', value: 'gateway.service' }, { text: '登陆服务', value: 'login.service' }];
    this.loading = false;
    this.listOfRandomUser = this.cacheService.get('tasks').getAll();
    this.search();
  }


  public editRow(task: SchedulingTask): void {

  }

  public deleteRow(task: SchedulingTask): void {
    this.listOfRandomUser = this.listOfRandomUser.filter(d => d.taskId !== task.taskId);
    this.listOfDisplay = this.listOfDisplay.filter(d => d.taskId !== task.taskId);
  }




  public async newTask(): Promise<void> {
    // ngx-task-add
    // this.modalService.create({
    //   nzTitle: '创建任务',
    //   nzContent: TaskAddComponent
    // });
    const drawerRef = this.openDrawer<TaskAddComponent, string, number>({
      nzTitle: '创建任务',
      nzContent: TaskAddComponent,
      nzMaskClosable: false,
      nzWidth: 'auto',
      nzContentParams: {
        input: 'sfsdfsg'
      }
    });
    const result = await this.waitDrawer(drawerRef);
    console.log(result);

  }

  public trackBydata(index: number, item: SchedulingTask): number {
    return item.taskId;
  }



  /**
   * 
   */
  public search(): void {
    this.visible = false;
    this.listOfDisplay = this.listOfRandomUser.filter((item: SchedulingTask) => item.taskName.indexOf(this.searchValue!) !== -1);
  }

  public reset(): void {
    this.searchValue = '';
    this.search();
  }

  public serviceFilterFn(list: string[], item: SchedulingTask): boolean {
    return list.some(serviceName => item.service.indexOf(serviceName) !== -1)
  }

  public taskNameSortFn(a: SchedulingTask, b: SchedulingTask): number {
    return a.taskName.localeCompare(b.taskName);
  }

}
