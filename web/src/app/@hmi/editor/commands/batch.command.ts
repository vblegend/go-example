import { HmiEditorComponent } from "../hmi.editor.component";
import { BasicCommand } from "./basic.command";


/**
 * 批处理命令 用于一次执行多个命令
 * 由传入的先后顺序执行
 */
export class BatchCommand extends BasicCommand {
    private commands: BasicCommand[] = [];

    /**
     * 批处理命令，依次执行多个命令
     * @param editor 
     * @param _commands 
     */
    constructor(editor: HmiEditorComponent, ..._commands: BasicCommand[]) {
        super(editor);
        this.commands = _commands;
        this.objects = [];
    }


    /**
     * 执行，重做
     */
    public execute(): void {
        for (let i = 0; i < this.commands.length; i++) {
            this.commands[i].execute();
        }
    }

    /**
     * 撤销
     */
    public undo(): void {
        for (let i = this.commands.length - 1; i >= 0; i--) {
            this.commands[i].undo();
        }
    }

    /**
     * 更新
     */
    public update(cmd: BasicCommand): void {

    }


    /**
     * 批处理命令 永不合并
     * @param command 
     * @returns 
     */
    public equal(command: BasicCommand): boolean {
        return false;
    }
}