import { ComponentRef, Injectable } from "@angular/core";
import { BasicWidgetComponent } from "@hmi/components/basic-widget/basic.widget.component";


export interface MetaDataItem {
    /**
     * 部件ID
     */
    widgetId: string;
    /**
     * 接口或事件名
     */
    name: string;

    /**
     * 接口会事件的id
     */
    value: string;
}






/**
 * 部件元数据服务
 */
@Injectable()
export class MetaDataService {

    /**
     * canvas 上所有对象的接口元数据
     */
    public readonly interfaces: MetaDataItem[] = [];
    /**
     * canvas 上所有对象的事件元数据
     */
    public readonly events: MetaDataItem[] = [];


    /**
     * canvas 上所有对象的事件元数据
     */
    public readonly params: Record<string, string[]> = {};





    /**
     * 
     * @param ref 
     */
    public add(ref: ComponentRef<BasicWidgetComponent>): void {

        for (const key in ref.instance.metaData.interface) {
            const intf = ref.instance.metaData.interface[key];
            this.interfaces.push({
                widgetId: ref.instance.configure.id,
                name: intf.name!,
                value: `${ref.instance.configure.type}:${intf.methodName}`
            });
            if (intf.args) {
                if (this.params[`${ref.instance.configure.type}:${intf.methodName}`] == null) {
                    this.params[`${ref.instance.configure.type}:${intf.methodName}`] = [];
                }
                const args = this.params[`${ref.instance.configure.type}:${intf.methodName}`];
                for (const arg of intf.args) {
                    if (args.indexOf(arg.argName) == -1) args.push(arg.argName);
                }
            }
        }

        for (const event of ref.instance.metaData.events) {
            this.events.push({
                widgetId: ref.instance.configure.id,
                name: event.eventName!,
                value: event.event!
            });
        }

    }


    /**
     * 
     * @param ref 
     * @returns 
     */
    public remove(ref: ComponentRef<BasicWidgetComponent>): void {
        for (let i = this.interfaces.length - 1; i >= 0; i--) {
            if (this.interfaces[i].widgetId === ref.instance.configure.id) {
                this.interfaces.splice(i, 1);
            }
        }
        for (let i = this.events.length - 1; i >= 0; i--) {
            if (this.events[i].widgetId === ref.instance.configure.id) {
                this.events.splice(i, 1);
            }
        }

    }
}
