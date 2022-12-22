/**
 * @license
 * Copyright Akveo. All Rights Reserved.
 * Licensed under the MIT License. See License.txt in the project root for license information.
 */
import { Route, Data } from '@angular/router';
import { RouteTitle } from './app/@core/models/route.titlee'
/* SystemJS module definition */
// declare var module: NodeModule;
// interface NodeModule {
//   id: string;
// }



// export declare var echarts: any;

declare module '@angular/router' {
    /**
     * document title
     */
    export interface Route {
        title?: RouteTitle;
    }
}