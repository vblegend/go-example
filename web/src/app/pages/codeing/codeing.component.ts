import { Component, ElementRef, Injector, ViewChild } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';


@Component({
  selector: 'ngx-codeing',
  styleUrls: ['./codeing.component.less'],
  templateUrl: './codeing.component.html',
})
export class CodeingComponent extends GenericComponent {

  constructor(injector: Injector) {
    super(injector);

  }

  editorOptions = { theme: 'vs-dark', language: 'typescript', };
  code: string = 'function x() {\nconsole.log("Hello world!");\n}';

}
