import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot } from '@angular/router';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { NetWorkService } from './network.service';
import { SessionService } from './session.service';
import { LoginResponse } from '@core/types/restful';

@Injectable()
export class AuthGuardService implements CanActivate {
  constructor(private router: Router,
    private sessionService: SessionService,
    private netWorkService: NetWorkService) {

  }

  public canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<boolean> | Promise<boolean> | boolean {
    const userCookie = this.sessionService.get<LoginResponse>("user");
    if (!this.verifyLoginState(userCookie!) || !this.verifyNetWorkState()) {
      this.router.navigateByUrl('/login');
      return false;
    } else {
      return true;
    }
  }


  private verifyNetWorkState(): boolean {
    return true;// this.netWorkService.isConnect;
  }

  private verifyLoginState(state: LoginResponse): boolean {
    if (state == null || state.expire == null) return false;
    const date = new Date(state.expire).getTime()
    const now = Date.now()
    return now < date;
  }

}
