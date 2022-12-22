import { Component, ElementRef, Injector, ViewChild } from '@angular/core';
import { UntypedFormBuilder, UntypedFormGroup, Validators } from '@angular/forms';
import { AccountService } from '@core/services/account.service';
import { MessageType } from '../../common/messagetype';
// import { NbAuthService } from '@nebular/auth';

import { GenericComponent } from '../basic/generic.component';


@Component({
  selector: 'ngx-login',
  styleUrls: ['./loginpage.component.less'],
  templateUrl: './loginpage.component.html',
})
export class LoginPageComponent extends GenericComponent {
  public background: string;
  public loadingWait: boolean;
  public validateForm!: UntypedFormGroup;


  constructor(injector: Injector, private fb: UntypedFormBuilder, private accessService: AccountService) {
    super(injector);
    this.loadingWait = false;
    this.background = 'url(/assets/images/team.png)';
  }

  protected onInit(): void {
    this.validateForm = this.fb.group({
      username: [null, [Validators.required]],
      password: [null, [Validators.required]],
      remember: [true]
    });
  }



  public submitForm(): void {
    if (this.validateForm.valid) {
      this.validateForm.value["code"] = "0";
      this.accessService.login(this.validateForm.value).then(e => {
        this.showMessage("login successful.", MessageType.Success);
        this.sessionService.set('user', e).then(e => {
          this.navigate('/');
        });
      }).catch(e => {
        this.sessionService.remove('user');
        // this.showMessage(e, MessageType.Error);
      })
    } else {
      Object.values(this.validateForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }
}
