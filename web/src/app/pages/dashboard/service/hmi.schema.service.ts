import { Injectable, Injector } from "@angular/core";
import { WidgetSchemaService } from "@hmi/services/widget.schema.service";
import { WidgetSchema } from "@hmi/configuration/widget.schema";
import { WidgetCategory } from "@hmi/configuration/widget.category";
import { ImageWidgetComponent } from "../widgets/image.widget/image.widget.component";
import { TextWidgetComponent } from "../widgets/text.widget/text.widget.component";
import { ButtonWidgetComponent } from "../widgets/button.widget/button.widget.component";
import { CurveChartWidgetComponent } from "../widgets/curvechart.widget/curvechart.widget.component";
import { SingalWidgetComponent } from "../widgets/signal.widget/signal.widget.component";

/**
 * 图片类小部件
 */
 export const CustomImageWidgets: WidgetSchema[] = [
    {
        icon: 'grace-tupian1',
        name: "图片部件",
        classify: WidgetCategory.Common,
        component: ImageWidgetComponent
    }, {
        icon: 'grace-fuwenben',
        name: "单行文本部件",
        classify: WidgetCategory.Text,
        component: TextWidgetComponent
    },
    {
        icon: 'grace-anniu',
        name: "按钮部件",
        classify: WidgetCategory.Common,
        component: ButtonWidgetComponent
    }
];


/**
 * 其他类小部件
 */
export const CustomOtherWidgets: WidgetSchema[] = [
    {
        icon: 'grace-iconfonticon-baobia1',
        name: "数据曲线图",
        classify: WidgetCategory.Shapes,
        component: CurveChartWidgetComponent
    },
    {
        icon: 'grace-base-station',
        name: "设备实时信号",
        classify: WidgetCategory.Text,
        component: SingalWidgetComponent
    }
];


/*
 * 小部件注册服务\
 * 用于注册程序内所有小部件
 */
@Injectable({
    providedIn: 'root',
})
export class HmiSchemaService extends WidgetSchemaService {
    constructor(protected injector: Injector) {
        super(injector);
        this.register(CustomImageWidgets);
        this.register(CustomOtherWidgets);
    }
}


