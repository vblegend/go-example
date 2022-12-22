import { Component, ElementRef, HostListener, Injector, OnInit, ViewChild, ViewContainerRef } from '@angular/core';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';
import { GenericComponent } from '@core/components/basic/generic.component';
import { ViewCanvasComponent } from '@hmi/components/view-canvas/view.canvas.component';

@Component({
  selector: 'ngx-iFrame',
  templateUrl: './iframe.component.html',
  styleUrls: ['./iframe.component.less']
})
export class IFrameComponent extends GenericComponent {
  // @ViewChild('iFrame', { static: true }) iFrame!: ElementRef<HTMLIFrameElement>;

  public url: SafeResourceUrl = "http://localhost:8000/#/pages/welcome/1";

  /**
   *
   */
  constructor(protected injector: Injector, private sanitizer: DomSanitizer,
  ) {
    super(injector);
    this.url = this.sanitizer.bypassSecurityTrustResourceUrl("http://10.169.42.139:8000/#/pages/welcome/1");
  }


  // protected onAfterViewInit(): void {
  //   const fWindow = this.iFrame.nativeElement.contentWindow as any;
  //   fWindow["__Angular-Admin__"] = {
  //     callback: () => {
  //       super.goBack();
  //     }
  //   } 
  //   this.iFrame.nativeElement.src = "http://localhost:8000/#/pages/welcome/1";
  // }



}
