import { Pipe, PipeTransform } from '@angular/core';
import { Router, RouterState, RouterStateSnapshot } from '@angular/router';
import { GenericComponent } from '@core/components/basic/generic.component';

@Pipe({ name: 'translator' })
export class TranslatorPipe implements PipeTransform {
    /**
     *
     */
    constructor(private router: Router) {
        // console.log(router);
    }


    public transform(value: string, component: GenericComponent): string {
        // component.translationKey 
        return value.toString() + '...';
    }



}