/* eslint-disable @typescript-eslint/no-explicit-any */
import { HttpErrorResponse, HttpEvent, HttpHandler, HttpInterceptor, HttpRequest, HttpResponse, HttpResponseBase } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { LoginResponse } from '@core/types/restful';
import { SessionService } from '@core/services/session.service';
import { Observable, of, throwError } from 'rxjs';
import { catchError, filter, map, mergeMap, tap } from 'rxjs/operators';
import { Router } from '@angular/router';
import { NzMessageService } from 'ng-zorro-antd/message';


export interface RestfulResponse<T> {
    traceId: string;
    code: number;
    msg: string;
    data: T;
}



@Injectable()
export class HttpHandlerInterceptor implements HttpInterceptor {

    /**
     *
     */
    constructor(protected session: SessionService, protected router: Router, protected messageService: NzMessageService) {
        console.log("New HttpHandlerInterceptor");
        console.log(session);

    }


    public intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
        req = this.handleRequest(req);
        return next.handle(req).pipe(
            mergeMap(evt => this.handleResponse(evt)),
            catchError((error: HttpErrorResponse) => {
                this.handleError(error);
                return throwError(() => error);
            })
        );
    }

    /**
     * 请求参数拦截处理
     */
    public handleRequest(req: HttpRequest<any>): HttpRequest<any> {
        // console.log(`拦截器A在请求发起前的拦截处理`);
        const usetInfo = this.session.get<LoginResponse>("user");
        if (usetInfo && usetInfo.token) {
            return req.clone({ setHeaders: { Authorization: `Bearer ${usetInfo.token}` } });
        }
        return req
    }


    /**
     * 返回结果拦截处理
     */
    public handleResponse(evt: HttpEvent<any>): Observable<HttpEvent<any>> {
        return new Observable<HttpEvent<any>>(observer => {
            if (evt instanceof HttpResponse) {
                // console.log("拦截器A在数据返回后的拦截处理");
                const response = evt.body as RestfulResponse<any>
                if (response.code != 0) {
                    observer.error(response.msg);
                    return;
                }
            } else {
                // console.log(`拦截器A接收到请求发出状态：${evt}`);
            }
            observer.next(evt);
        });
    }
    /**
     * 返回结果拦截处理
     */
    public handleError(error: HttpErrorResponse): void {
        switch (error.status) {
            case 200:
                break;
            case 401: // Unauthorized
                // todo
                this.session.remove("user");
                this.router.navigate(["/"], undefined);
                break;
            // case 403: // Forbidden
            //     // todo
            //     break;
            // case 404: // Forbidden
            //     // todo
            //     break;
            default:
                if (typeof error == 'string'){
                    this.messageService.error(`Bad Request\n${error}`)
                }else{
                    this.messageService.error(`Bad Request\nCode:${error.status}\nMessage:${error.message}`)
                }
            // todo
        }
    }
}