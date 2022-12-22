import { Injectable, Injector } from '@angular/core';




/**
 * General service collection 
 */
@Injectable({
    providedIn: 'root'
})
export class CommonService {

    constructor(protected injector: Injector) {

    }


}