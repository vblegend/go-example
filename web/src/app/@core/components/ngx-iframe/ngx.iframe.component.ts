import { Component, ElementRef, HostListener, Injector, Input, OnInit, ViewChild, ViewContainerRef } from '@angular/core';
import { SafeResourceUrl } from '@angular/platform-browser';
import { GenericComponent } from '@core/components/basic/generic.component';
import { Guid } from '@core/util/guid';
import { ViewCanvasComponent } from '@hmi/components/view-canvas/view.canvas.component';

@Component({
  selector: 'ngx-iframe',
  templateUrl: './ngx.iframe.component.html',
  styleUrls: ['./ngx.iframe.component.less']
})
export class NgxIFrameComponent extends GenericComponent {
  @ViewChild('iFrame', { static: true }) iFrame!: ElementRef<HTMLIFrameElement>;

  @Input()
  public src?: SafeResourceUrl;

  public windowId: string;

  /**
   *
   */
  constructor(protected injector: Injector) {
    super(injector);
    this.windowId = Guid.custom(8, 62);
  }



  protected onAfterViewInit(): void {
    // (<any>window).do =  (text: string)=> {
    //   console.log(`do ${text}`)
    // }
    // this.iFrame.nativeElement.contentWindow!.windowId = this.windowId;
  }


  public PostMessage(): void {
    // this.iFrame.nativeElement.contentWindow?.postMessage("", "")
  }



  public iFrameLoadComplete(): void {
    // this.iFrame.nativeElement.contentWindow?.postMessage("Hello IFrame", '*');
  }
}
