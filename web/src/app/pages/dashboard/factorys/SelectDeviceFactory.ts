import { Injector } from "@angular/core";
import { ButtonActionFactory } from "@hmi/editor/common";
import { BasicPropertyComponent } from "@hmi/editor/components/basic-property/basic.property.component";
/***
 * 自定义小部件的按钮绑定设备属性工厂
 */
export class SelectDeviceFactory extends ButtonActionFactory {

    /**
     * init get service
     * @param injector 
     */
    constructor(protected injector: Injector) {
        super(injector);
        // 示例
        // const deviceService = this.injector.get(DeviceService);
    }

    public toBinding(value: number): string {
        return `选择设备`;
    }


    public canExecute(value: number): boolean {
        return true;
    }

    public execute(propertie: BasicPropertyComponent<number>): void {
        console.log('这里打开选择设备对话框：');
        console.log('调用 propertie.saveAndUpdate(1000000); 保存结果！');
        propertie.saveAndUpdate(1000000);
    }
}