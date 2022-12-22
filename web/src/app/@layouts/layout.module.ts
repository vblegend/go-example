import { CommonModule } from "@angular/common";
import { ModuleWithProviders, NgModule } from "@angular/core";
import { RouterModule } from "@angular/router";
import { NzGridModule } from "ng-zorro-antd/grid";
import { NzIconModule } from "ng-zorro-antd/icon";
import { NzLayoutModule } from "ng-zorro-antd/layout";
import { NzMenuModule } from "ng-zorro-antd/menu";
import { FooterComponent } from "./footer/footer.component";
import { HeaderComponent } from "./header/header.component";
import { SidebarComponent } from "./sidebar/sidebar.component";
import { LogoComponent } from './logo/logo.component';
import { CoreModule } from "@core/core.module";
import { NzToolTipModule } from 'ng-zorro-antd/tooltip';
import { NzSelectModule } from 'ng-zorro-antd/select';
import { FormsModule } from "@angular/forms";
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzSpaceModule } from 'ng-zorro-antd/space';
import { NzDividerModule } from 'ng-zorro-antd/divider';
import { NzBadgeModule } from 'ng-zorro-antd/badge';
import { NzPopconfirmModule } from 'ng-zorro-antd/popconfirm';









@NgModule({
    imports: [
        CommonModule,
        CoreModule,
        FormsModule,
        RouterModule,
        NzIconModule,
        NzMenuModule,
        NzLayoutModule,
        NzGridModule,
        NzToolTipModule,
        NzButtonModule,
        NzSelectModule,
        NzSpaceModule,
        NzDividerModule,
        NzBadgeModule,
        NzPopconfirmModule
    ],
    exports: [
        FooterComponent,
        HeaderComponent,
        SidebarComponent,
        LogoComponent
    ],
    declarations: [
        FooterComponent,
        HeaderComponent,
        SidebarComponent,
        LogoComponent
    ],
})

export class LayoutModule {
    constructor() {

    }

    public static forRoot(): ModuleWithProviders<LayoutModule> {
        return {
            ngModule: LayoutModule,
            providers: [
            ]
        };
    }
}
