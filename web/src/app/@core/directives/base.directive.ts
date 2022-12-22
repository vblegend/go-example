import { Directive, ElementRef, HostListener, Input, Output, EventEmitter, OnInit, Optional, ViewContainerRef, ViewRef, ComponentRef, Injector, ComponentFactoryResolver, NgZone } from '@angular/core';


@Directive({
    selector: '[selectRectangle]'
})

export class BaseDirective implements OnInit {

    protected element: HTMLElement;
    protected elementRef: ElementRef;
    protected viewContainerRef: ViewContainerRef;
    protected componentFactoryResolver: ComponentFactoryResolver;
    protected zone: NgZone;

    constructor(protected injector: Injector) {
        this.zone = injector.get(NgZone);
        this.elementRef = injector.get(ElementRef);
        this.viewContainerRef = injector.get(ViewContainerRef);
        this.componentFactoryResolver = injector.get(ComponentFactoryResolver);
        this.element = this.elementRef.nativeElement;
    }

    public ngOnInit(): void {
        this.onInit();
    }

    protected onInit(): void {

    }


    // @HostListener('mousedown', ['$event'])
    // public onMouseDown(ev: MouseEvent): void {
    // }

    // @HostListener('document:mousemove', ['$event'])
    // public onMouseMove(ev: MouseEvent): void {
    // }

    // @HostListener('document:mouseup', ['$event'])
    // public onMouseUp(ev: MouseEvent): void {
    // }

}