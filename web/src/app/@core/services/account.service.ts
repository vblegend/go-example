import { Injectable } from "@angular/core";
import { RestfulService } from "./restful.service";
import { LoginRequest, LoginResponse } from '@core/types/restful';

@Injectable()
export class AccountService extends RestfulService {

    public async login(loginInfo: LoginRequest): Promise<LoginResponse> {
        const res = await this.post<LoginResponse>(`/api/login/`, loginInfo);
        return res.data;
    }
}