

import { HttpClient } from '@angular/common/http';
import { Injectable, Injector } from '@angular/core';
import { RestfulResponse } from '@core/types/restful';





@Injectable({
    providedIn: 'root'
})
export class RestfulService {


    /**
     *
     */
    constructor(protected client: HttpClient) {
        this.onInit()
    }


    protected onInit(): void {

    }


    public async get<T>(url: string): Promise<RestfulResponse<T>> {
        return new Promise<RestfulResponse<T>>((resolve, reject) => {
            this.client.get<RestfulResponse<T>>(url).subscribe({ next: resolve, error: reject });
        });
    }



    public async post<T>(url: string, body: Object): Promise<RestfulResponse<T>> {
        return new Promise<RestfulResponse<T>>((resolve, reject) => {
            this.client.post<RestfulResponse<T>>(url, body).subscribe({ next: resolve, error: reject });
        });
    }





}