/**
 * disable zone event plugin servic
 * 
 * <button class="btn btn-primary" (click@outside)="onClick()">
 *   Click me!
 * </button>
 * 
 */
import { Injectable } from "@angular/core";
import { EventManager } from "@angular/platform-browser";

@Injectable()
export class OutSideEventPluginService {
    public manager!: EventManager;

    public supports(eventName: string): boolean {
        return eventName.endsWith("@outside");
    }

    public addEventListener(element: HTMLElement, eventName: string, originalHandler: EventListener): Function {
        const [nativeEventName] = eventName.split("@");
        this.manager.getZone().runOutsideAngular(() => {
            element.addEventListener(nativeEventName, originalHandler);
        });
        return () => element.removeEventListener(nativeEventName, originalHandler);
    }
}