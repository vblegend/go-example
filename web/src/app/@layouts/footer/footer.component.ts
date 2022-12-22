import { Component } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';

@Component({
  selector: 'app-footer',
  templateUrl: './footer.component.html',
  styleUrls: ['./footer.component.less']
})
export class FooterComponent extends GenericComponent {



  protected onInit(): void {
  }

}
