import { Directive, ElementRef, HostListener, Injector, Input } from '@angular/core';
import { BaseDirective } from './base.directive';

@Directive({
    selector: '[selectdisable]'
})
export class UnSelectedDirective extends BaseDirective {
    constructor(protected injector: Injector) {
        super(injector);
        this.element.style.userSelect = 'none';
    }
}