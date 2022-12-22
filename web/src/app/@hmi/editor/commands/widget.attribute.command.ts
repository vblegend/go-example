import { ComponentRef } from '@angular/core';
import { AnyObject } from '@core/common/types';
import { BasicWidgetComponent } from '@hmi/components/basic-widget/basic.widget.component';
import { HmiEditorComponent } from '../hmi.editor.component';
import { BasicCommand } from './basic.command';




export class WidgetAttributeCommand<TValue> extends BasicCommand {
    public objects: ComponentRef<BasicWidgetComponent>[];
    private oldValues!: TValue[];
    private newValues!: TValue[];

    /**
     * 设置对象属性值
     * 如果objects 对象含有 setXXX() 方法时会调用该方法更新
     * @param editor 
     * @param objects 
     * @param path  属性路径 如‘configure/rect’
     * @param newValues
     */
    public constructor(editor: HmiEditorComponent, objects: ComponentRef<BasicWidgetComponent>[], path: string, newValues: TValue[], batchNo?: number) {
        super(editor);
        this.name = '修改属性';
        this.objects = objects;
        this.oldValues = [];
        const paths = path.split('/');
        this.attributeName = paths.pop()!;
        this.attributePaths = paths;
        if (batchNo != null) this.batchNo = batchNo;
        for (const object of objects) {
            const target = this.getTarget(object.instance);
            if (this.attributeName) {
                const oldValue = target[this.attributeName as keyof AnyObject];
                this.oldValues.push(<TValue><unknown>oldValue);
            }
        }
        if (newValues.length == 1 && this.oldValues.length > 1) {
            this.newValues = Array(this.oldValues.length).fill(newValues[0]);
        } else if (this.oldValues.length == newValues.length) {
            this.newValues = newValues;
        } else {
            throw 'error';
        }
    }


    public getTarget(object: BasicWidgetComponent): AnyObject {
        let root = <AnyObject><unknown>object;
        for (let i = 0; i < this.attributePaths.length; i++) {
            const key = this.attributePaths[i];
            root = root[key];
        }
        return root;
    }




    public execute(): void {
        for (let i = 0; i < this.objects.length; i++) {
            const rootObject = this.objects[i].instance;
            const object = this.getTarget(rootObject);
            object[this.attributeName] = <AnyObject><unknown>this.newValues[i];
            // this.objects[i].changeDetectorRef.markForCheck();
            // call setXXX();
            // const setter = `set${this.attributeName[0].toUpperCase()}${this.attributeName.substring(1)}`;
            // if (rootObject[setter as key of AnyObject] && typeof rootObject[setter] === 'function') {
            //     rootObject[setter](this.newValues[i]);
            // }
        }
    }

    public undo(): void {
        for (let i = 0; i < this.objects.length; i++) {
            const rootObject = this.objects[i].instance;
            const object = this.getTarget(rootObject);
            object[this.attributeName] = <AnyObject><unknown>this.oldValues[i];
            // this.objects[i].changeDetectorRef.markForCheck();
            // call setXXX();
            // const setter = `set${this.attributeName![0].toUpperCase()}${this.attributeName!.substring(1)}`;
            // if (rootObject[setter] && typeof rootObject[setter] === 'function') {
            //     rootObject[setter](this.oldValues[i]);
            // }
        }
    }

    public update(command: WidgetAttributeCommand<TValue>) {
        this.newValues = command.newValues;
    }

}
