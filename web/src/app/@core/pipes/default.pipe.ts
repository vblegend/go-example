import { Pipe, PipeTransform } from '@angular/core';

@Pipe({ name: 'default' })
export class DefaultPipe implements PipeTransform {
    public transform(value: string | number, defaultValue: string | number): string | number {
        return value != null ? value : defaultValue;
    }
}