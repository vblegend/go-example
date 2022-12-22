
import { Component, EventEmitter, Input, Output, ViewChild } from '@angular/core';
import { CdkPortal, DomPortalOutlet } from '@angular/cdk/portal';
import { GenericComponent } from './generic.component';
import { TimerTask } from '@core/common/timer.task';

/**
 * This component template wrap the projected content
 * with a 'cdkPortal'.
 *     
 * <!-- Whatever you wrap with the '<window>' component will be rendered on a new window  -->
 * <ng-window *ngIf="showPortal">
 *   <h2>Hello world from amother window!!</h2>
 *   <button (click)="this.showPortal = false">Close me!</button>
 * </ng-window>
 * 
 */

@Component({
    selector: 'ng-window',
    template: `
    <ng-container *cdkPortal>
      <ng-content></ng-content>
    </ng-container>
  `
})
export class WindowComponent extends GenericComponent {
    // STEP 1: get a reference to the portal
    @ViewChild(CdkPortal, { static: true, read: CdkPortal }) public portal?: CdkPortal;

    @Input()
    public title: string = '';

    // STEP 2: save a reference to the window so we can close it
    private externalWindow: Window | null = null;
    @Output() public afterClose: EventEmitter<WindowComponent> = new EventEmitter<WindowComponent>();

    // STEP 3: Inject all the required dependencies for a PortalHost
    protected onInit(): void {
        // STEP 4: create an external window
        this.externalWindow = window.open('', '', 'width=600,height=400,left=200,top=200')!;
        this.externalWindow.onbeforeunload = (e) => {
            this.afterClose.emit(this);
        };
        if (this.title && this.title.length > 0) {
            this.externalWindow!.document.title = this.title;
        }
        // STEP 5: create a PortalHost with the body of the new window document    
        const host = new DomPortalOutlet(
            this.externalWindow!.document.body,
            this.componentFactoryResolver,
            this.applicationRef,
            this.injector
        );
        // STEP 6: Attach the portal
        host.attach(this.portal);
    }


    protected onDestroy(): void {
        // STEP 7: close the window when this component destroyed
        this.externalWindow!.close()
    }
}