import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { WidgetSchemaService } from '@hmi/services/widget.schema.service';
import { ErrorComponent } from './@core/components/error/error.component';
import { LoginPageComponent } from './@core/components/login/loginpage.component';
import { NotFoundComponent } from './@core/components/notfound/not-found.component';
import { AuthGuardService } from './@core/services/auth.guard.service';
import { HmiSchemaService } from './pages/dashboard/service/hmi.schema.service';
import { CodeingComponent } from './pages/codeing/codeing.component';


const routes: Routes = [
  {
    path: 'login',
    title: 'login',
    component: LoginPageComponent
  },
  {
    path: 'notfound',
    title: 'not found',
    component: NotFoundComponent,
    canActivate: [AuthGuardService]
  },
  {
    path: 'error',
    title: 'error',
    component: ErrorComponent,
  },
  {
    path: 'pages',
    loadChildren: () => import('./pages/pages.module').then(m => m.PagesModule),
    canActivate: [AuthGuardService]
  },
  {
    path: '',
    redirectTo: 'pages',
    pathMatch: 'full'
  },
  {
    path: '**',
    component: NotFoundComponent,
    canActivate: [AuthGuardService]
  },
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes, { useHash: true })
  ],
  exports: [RouterModule],
  providers:[
    {
      provide: WidgetSchemaService,
      useClass: HmiSchemaService,
    }
  ]
})
export class AppRoutingModule { }
