import { Component, EventEmitter, Injector, OnInit, Output } from '@angular/core';
import { GenericComponent } from '@core/components/basic/generic.component';

// import txt from './123.txt';
const $import = require as (url: string) => string;




@Component({
  selector: 'app-welcome',
  templateUrl: './welcome.component.html',
  styleUrls: ['./welcome.component.less']
})
export class WelcomeComponent extends GenericComponent {
  @Output() public deleteRequest: EventEmitter<Object> = new EventEmitter<Object>(true);
  public id!: string | null;
  public name!: string;
  public content: string = '';

  // private subscription: Subscription;

  constructor(injector: Injector) {
    super(injector)

  }

  protected onQueryChanges(): void {
    this.id = this.queryParams.get('id');
    console.log(`app-welcome onRouter ${this.id}`);
    // this.deleteRequest.emit(this.id);

    // this.subscribe(this.deleteRequest);
  }


  protected onInit(): void {
    this.id = this.queryParams.get('id');
    console.log(`app-welcome onInit ${this.id}`);
    this.content = $import('./script.txt');
    // const ref = this.generateComponent(NotFoundComponent);
    // ref.destroy();
    // this.subscription = this.deleteRequest.subscribe(e => {
    //   console.log(`app-welcome subscribe ${this.id}`);
    // })

  }

  protected onDestroy(): void {
    // this.subscription.unsubscribe();
    console.log(`app-welcome onDestroy`);
  }



}
