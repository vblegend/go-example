// import { RouteTitle } from "@core/models/route.title";
import { Route, Data } from '@angular/router';
import { SessionService } from '@core/services/session.service'

declare module 'raw-loader!*' {
    const contents: string;
    export = contents;
}

declare module '!raw-loader!*' {
    const contents: string;
    export = contents;
}


declare module '!!raw-loader!*' {
    const contents: string;
    export = contents;
}


declare module "*.txt" {
    const content: string;
    export default content;
}


export interface RouteTitle {
    value: string;
    needsTranslator?: boolean;
}

declare module '@angular/router' {
    /**
     * document title
     */
    export interface Route {
        /**
         * menu title configure
         */
        title?: RouteTitle;
    }
}

declare global {
    export interface Window {
        windowId: string;
        sessionService: SessionService;
        call(): void



    }
}
