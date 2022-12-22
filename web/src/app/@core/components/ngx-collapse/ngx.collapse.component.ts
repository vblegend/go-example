import { Component, DoCheck, ElementRef, Injector, Input, OnInit, TemplateRef, ViewChild, ViewChildren } from '@angular/core';
import { GenericComponent } from '../basic/generic.component';

@Component({
  selector: 'ngx-collapse',
  templateUrl: './ngx.collapse.component.html',
  styleUrls: ['./ngx.collapse.component.less']
})
export class CollapseComponent extends GenericComponent implements DoCheck {

  /**
   * 展开的
   */
  @Input()
  public expanded: boolean = false;

  /**
   * 当内容高度为0时隐藏
   */
  @Input()
  public autoHide: boolean = false;


  @Input()
  public header: string = 'ngx-collapse';

  public hasContent: boolean = false;


  @ViewChild('content')
  public content!: ElementRef<any>;


  constructor(protected injector: Injector) {
    super(injector);
  }

  public header_click(): void {
    this.expanded = !this.expanded
  }


  public ngDoCheck(): void {
    this.hasContent = !this.autoHide || (this.autoHide && this.content && this.content.nativeElement.clientHeight > 0);
  }




}