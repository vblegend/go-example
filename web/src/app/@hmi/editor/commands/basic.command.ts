import { HmiEditorComponent } from "../hmi.editor.component";


export abstract class BasicCommand {
    public id!: number;
    public type!: string;
    public name!: string;
    public batchNo!: number;
    public attributeName!: string;
    public attributePaths!: string[];
    protected editor!: HmiEditorComponent;
    public executeTime!: Date;
    public objects!: any[];
    /**
     *
     */
    constructor(editor: HmiEditorComponent) {
        this.id = -1;
        this.editor = editor;
        // this.objects = [];
        // this.oldValues = [];
        // this.newValues = [];
        this.batchNo = Math.round((Math.random() * Number.MAX_VALUE));
    }


    /**
     * 执行，重做
     */
    public execute(): void {

    }

    /**
     * 撤销
     */
    public undo(): void {

    }

    /**
     * 更新
     */
    public update(cmd: BasicCommand): void {

    }


    /**
     * 判断两条命令是否相同
     * @param command
     */
    // public abstract equal(command: BasicCommand): boolean;



    public equal(command: BasicCommand): boolean {
        if (command.type != this.type) return false;
        if (command.attributeName != this.attributeName) return false;
        if (command.objects.length != this.objects.length) return false;
        for (let i = 0; i < command.objects.length; i++) {
            if (command.objects[i] != this.objects[i]) {
                return false;
            }
        }
        return true;
    }
}