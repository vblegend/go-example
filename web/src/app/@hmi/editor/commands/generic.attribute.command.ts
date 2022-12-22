import { ObjectUtil } from '@core/util/object.util';
import { HmiEditorComponent } from '../hmi.editor.component';
import { BasicCommand } from './basic.command';

export class GenericAttributeCommand<TObject, TValue> extends BasicCommand {
    public objects: TObject[];
    private oldValues!: TValue[];
    private newValues!: TValue[];

    /**
     * 设置对象属性值
     * @param editor 
     * @param _objects 
     * @param path  属性路径 如‘configure/rect’
     * @param newValues
     */
    public constructor(editor: HmiEditorComponent, _objects: TObject[], path: string, newValues: TValue[], batchNo?: number | undefined) {
        super(editor);
        this.name = '修改属性';
        this.objects = _objects;
        this.oldValues = [];
        const paths = path.split('/');
        this.attributeName = paths.pop()!;
        this.attributePaths = paths;
        if (batchNo != null) this.batchNo = batchNo;
        for (const object of _objects) {
            const target = this.getTarget(object);
            const oldValue = ObjectUtil.clone(target[this.attributeName as keyof TObject]);
            this.oldValues.push(<TValue><unknown>oldValue);
        }
        if (newValues.length == 1 && this.oldValues.length > 1) {
            this.newValues = Array(this.oldValues.length).fill(newValues[0]);
        } else if (this.oldValues.length == newValues.length) {
            this.newValues = newValues;
        } else {
            throw '错误：命令参数不一致。';
        }
    }

    public getTarget(object: TObject): TObject {
        let root = object;
        for (let i = 0; i < this.attributePaths.length; i++) {
            const key = this.attributePaths[i];
            root = <TObject><unknown>root[key as keyof TObject];
        }
        return root;
    }




    public execute(): void {
        for (let i = 0; i < this.objects.length; i++) {
            const rootObject = this.objects[i];
            const object = this.getTarget(rootObject);

            const properties: Record<string, TValue> = {};
            properties[this.attributeName] = this.newValues[i];
            Object.assign(object, properties);


            // object[this.attributeName] 
            // object[this.attributeName] = this.newValues[i];
        }
    }

    public undo(): void {
        for (let i = 0; i < this.objects.length; i++) {
            const rootObject = this.objects[i];
            const object = this.getTarget(rootObject);
            // object[this.attributeName as keyof AnyObject] = this.oldValues[i];

            const properties: Record<string, TValue> = {};
            properties[this.attributeName] = this.oldValues[i];
            Object.assign(object, properties);

        }
    }

    public update(command: GenericAttributeCommand<TObject, TValue>) {
        this.newValues = command.newValues;
    }



}
