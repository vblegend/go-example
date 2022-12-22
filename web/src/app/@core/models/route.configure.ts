import { Params } from "@angular/router";




export interface RouteConfigure {
    id?: number;
    title: string;
    icon?: string;
    path?: string;
    opened?: boolean;
    disabled?: boolean;
    selected?: boolean;
    queryParams?: Params;
    children?: RouteConfigure[];
}