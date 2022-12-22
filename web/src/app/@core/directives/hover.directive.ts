import { Directive, ElementRef, HostListener, Input } from '@angular/core';
import { BaseDirective } from './base.directive';

@Directive({
    selector: '[hitHover]'
})
export class HoverDirective extends BaseDirective {
    @Input('hitHover') color: string = '#007ACC';
    @Input() cursor?: string;

    private isHover?: boolean;
    private _oldColor?: string;
    private _oldCursor?: string;

    protected onInit(): void {
        this.isHover = false;
    }

    @HostListener('mouseenter')
    public onMouseEnter(): void {
        if (!this.isHover) {
            this.isHover = true;
            this._oldCursor = this.element.style.cursor;
            this._oldColor = this.element.style.color;
            this.element.style.color = this.color;
            if (this.cursor) this.element.style.cursor = this.cursor;
        }
    }


    @HostListener('mouseleave')
    public onMouseLeave(): void {
        if (this.isHover) {
            this.isHover = false;
            this.element.style.color = this._oldColor!;
            if (this.cursor) this.element.style.cursor = this._oldCursor!;
            this._oldColor = undefined;
        }
    }

}