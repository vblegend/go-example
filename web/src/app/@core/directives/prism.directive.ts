import { AfterViewInit, Directive, Input } from '@angular/core';
import * as Prism from 'prismjs';
import { BaseDirective } from './base.directive';

/***
 * ProsmJS 高亮代码块指令
 * ``` html
  <pre>
       <code prism [lineNumbers]="true" language="typescript" >{{content}}</code>
  </pre>
 * ```
 */
@Directive({
    selector: '[prism]'
})
export class ProsmDirective extends BaseDirective implements AfterViewInit {

    @Input()
    public language: string = 'javascript';

    @Input()
    public content: string = '';

    @Input()
    public lineNumbers: boolean = false;


    public ngAfterViewInit(): void {
        if (this.lineNumbers && !Prism.plugins.lineNumbers) require('prismjs/plugins/line-numbers/prism-line-numbers.js');
        if (!this.lineNumbers && Prism.plugins.lineNumbers) delete Prism.plugins.lineNumbers;
        if (this.lineNumbers) {
            this.element.classList.add('line-numbers');
        }
        require(`prismjs/components/prism-${this.language}.js`);
        this.element.classList.add(`language-${this.language}`);
        Prism.highlightElement(this.element);
    }




}