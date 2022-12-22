import { Component, Injector, Input } from "@angular/core";
import { NzDrawerRef } from "ng-zorro-antd/drawer";
import { GenericComponent } from "./generic.component";



@Component({
    selector: 'ngx-generic-drawer',
    template: `<ng-container #view></ng-container>`
})
export class GenericDrawerComponent<TInput, TOut> extends GenericComponent {
    @Input()
    public readonly input?: TInput;

    protected readonly drawerRef: NzDrawerRef<TOut>;

    constructor(injector: Injector) {
        super(injector);
        this.drawerRef = injector.get(NzDrawerRef);
    }

    public submit(): void {
        this.close();
    }

    public cancel(): void {
        this.close();
    }


    /**
     * close drawer and result value
     * @param value 
     */
    public close(value?: TOut): void {
        this.drawerRef.close(value);
    }
}