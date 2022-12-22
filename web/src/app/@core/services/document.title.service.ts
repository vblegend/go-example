import { Injectable } from '@angular/core';
import { Router, NavigationEnd, RouterState, RouterStateSnapshot, RouterEvent, ActivatedRouteSnapshot, NavigationCancel, NavigationStart } from '@angular/router';
import { Title } from '@angular/platform-browser';
import { RouteTitle } from '../models/route.title';
import { Subscription } from 'rxjs';

@Injectable({
    providedIn: 'root'
})
export class DocumentTitleService {

    private subscription!: Subscription | null;
    private _defaultTitle!: RouteTitle;
    private _globalSuffix!: RouteTitle;
    private _globalPrefix!: RouteTitle;

    constructor(private router: Router,
        public titleService: Title) {

    }

    public get defaultTitle(): RouteTitle {
        return this._defaultTitle;
    }

    public set defaultTitle(value: RouteTitle) {
        this._defaultTitle = value;
    }

    /**
     * get/set global title suffix
     */
    public get globalSuffix(): RouteTitle {
        return this._globalSuffix;
    }

    public set globalSuffix(value: RouteTitle) {
        this._globalSuffix = value;
    }

    /**
     * get/set global title prefix
     */
    public get globalPrefix(): RouteTitle {
        return this._globalPrefix;
    }

    public set globalPrefix(value: RouteTitle) {
        this._globalPrefix = value;
    }

    /**
     * load page default title from router
     */
    public register(): void {
        if (this.subscription == null) {
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            this.subscription = this.router.events.subscribe(<any>this.router_event.bind(this));
        }
    }


    public getUrl(route: ActivatedRouteSnapshot): string {
        let next = this.getTruthRoute(route);
        const segments: string[] = [];
        while (next) {
            segments.unshift(...next.url.map(e => e.path));
            next = next.parent!;
        }
        return segments.join('/');
    }

    private getTruthRoute(route: ActivatedRouteSnapshot): ActivatedRouteSnapshot {
        let next = route;
        while (next.firstChild) {
            next = next.firstChild;
        }
        return next;
    }

    public unRegister(): void {
        if (this.subscription) {
            this.subscription.unsubscribe();
            this.subscription = null;
        }
    }

    private router_event(event: RouterEvent) {
        // console.log(event.constructor.name);

        if(event instanceof NavigationStart){

            // console.log(event);

        }



        if (event instanceof NavigationEnd || event instanceof NavigationCancel) {
            let title = this.getCurrentTitle(this.router);
            // console.log(title);s
            if (title == null) title = this._defaultTitle;
            const suffixText = this.getTitleText(this._globalSuffix);
            const titleText = this.getTitleText(title);
            const prefixText = this.getTitleText(this._globalPrefix);
            if (event instanceof NavigationCancel) this.titleService.setTitle('');
            this.titleService.setTitle(suffixText + titleText + prefixText);
        }
    }

    private getTitleText(title: RouteTitle): string {
        let value: string = '';
        if (title != null) {
            value = title.value;
            if (title.needsTranslator) {
                value = value + '...';
            }
        }
        return value;
    }

    private getCurrentTitle(router: Router): RouteTitle | undefined {
        let title: string = "";
        const state: RouterState = router.routerState;
        const snapshot: RouterStateSnapshot = state.snapshot;
        let node = snapshot.root;
        while (node != null) {
            if (node.routeConfig && node.routeConfig.title) {
                title = node.routeConfig.title as string;
            }
            node = node.firstChild!;
        }
        return {
            value :title,
            needsTranslator : false
        };
    }

}