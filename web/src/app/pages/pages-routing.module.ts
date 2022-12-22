import { RouterModule, Routes } from '@angular/router';
import { NgModule } from '@angular/core';
import { PagesComponent } from './pages.component';
import { WelcomeComponent } from './welcome/welcome.component';
import { ExamplesComponent } from './examples/examples.component';
import { CodeingComponent } from './codeing/codeing.component';
import { NotFoundComponent } from '@core/components/notfound/not-found.component';
import { AuthGuardService } from '@core/services/auth.guard.service';
import { IFrameComponent } from './iframe/iframe.component';

const routes: Routes = [{
  path: '',
  component: PagesComponent,
  children: [
    {
      path: 'welcome/:id',
      title: 'Wel Come',
      component: WelcomeComponent
    },
    {
      path: 'dashboard',
      title: 'dashboard',
      loadChildren: () => import('./dashboard/dashboard.module').then(m => m.DashboardModule),
    },
    {
      path: 'codeing',
      title: 'codeing',
      component: CodeingComponent
    },
    {
      path: 'examples',
      title: 'dashboard',
      component: ExamplesComponent
    },
    {
      path: 'ifame',
      title: 'ifame',
      component: IFrameComponent
    },
    {
      path: 'tasks',
      title: '计划任务',
      loadChildren: () => import('./tasks/tasks.module').then(m => m.TasksModule),
    },
    {
      path: 'notfound',
      title: 'not found',
      component: NotFoundComponent,
      canActivate: [AuthGuardService]
    },
    {
      path: '',
      redirectTo: 'dashboard',
      pathMatch: 'full'
    },
    {
      path: '**',
      component: NotFoundComponent,
      canActivate: [AuthGuardService]
    },
  ]
}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PagesRoutingModule {
}
