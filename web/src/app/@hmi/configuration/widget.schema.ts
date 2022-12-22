import { Type } from "@angular/core";
import { BasicWidgetComponent } from "@hmi/components/basic-widget/basic.widget.component";
import { BasicPropertyComponent } from "@hmi/editor/components/basic-property/basic.property.component";
import { WidgetCategory } from "./widget.category";
import { WidgetDefaultConfigure } from "./widget.configure";


export interface WidgetSchema {
    /**
     * 组件类型
     * schema文件中不需要定义，当service load时会自动填充
     */
    type?: string | null;

    /**
     * 组件的图标
     * 用于在对象列表中显示
     */
    icon: string;

    /**
     * 组件的分类
     * 用于在对象列表中分类显示
     */
    classify: WidgetCategory | string;

    /**
     * 预览图片
     */
    previewImage?: string;
    /**
     * 组件的名称
     * 默认名称，对象列表中显示
     */
    name: string;

    /**
     * 组建的实例
     */
    component: Type<BasicWidgetComponent>;

    /**
     * 定义data结构中所用到的属性选择器\
     * 属性对话框中所显示的属性设置由此定义\
     * 未定义的属性将无法设置属性
     */
    // properties?: Record<string, Type<BasicPropertyComponent>>;
}

