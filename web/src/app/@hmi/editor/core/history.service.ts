import { Type } from "@angular/core";
import { BasicCommand } from "../commands/basic.command";
import { HmiEditorComponent } from "../hmi.editor.component";



export class HistoryService {

    private editor: HmiEditorComponent;
    private undos: BasicCommand[];
    private redos: BasicCommand[];
    public disabled: boolean;
    private idCounter: number;


    public constructor(editor: HmiEditorComponent) {
        this.editor = editor;
        this.undos = [];
        this.redos = [];
        this.idCounter = 0;
        this.disabled = false;
    }


    /**
     * 当前是否可以执行撤销操作
     */
    public get canUndo(): boolean {
        return this.undos.length > 0;
    }

    /**
     * 当前是否可以执行重做操作
     */
    public get canRedo(): boolean {
        return this.redos.length > 0;
    }

    /**
     * 上一条命令
     */
    public get lastCommand(): BasicCommand {
        if (this.undos.length == 0) return null!;
        return this.undos[this.undos.length - 1];
    }

    /**
     * 执行某条可以被插销、重做的指令
     * 与上一条指令相同  且batchNo相同 合并 
     * 与上一条指令相同 且 stickTime内  合并
     * @param cmd 
     * @param optionalName 
     */
    public execute(cmd: BasicCommand): void {
        let newCommand = true;
        if (this.undos.length > 0) {
            const lastCmd = this.undos[this.undos.length - 1];
            if (cmd.equal(lastCmd)) {
                const timeDifference = new Date().getTime() - lastCmd.executeTime.getTime();
                // 两次操作间隔小于300毫秒  或属于同一批次的 合并
                if (lastCmd.batchNo === cmd.batchNo || timeDifference < 300) {
                    lastCmd.update(cmd);
                    cmd = lastCmd;
                    newCommand = false;
                }
            }
        }
        if (newCommand) {
            this.undos.push(cmd);
            cmd.id = ++this.idCounter;
        }
        cmd.execute();
        cmd.executeTime = new Date();
        this.redos = [];
    }




    /**
     * 撤销并删除最后一条命令\
     * 命令必须为 type 类型
     * @param type 最后一条命令类型
     */
    public undoAndRemoveLast(type: Type<BasicCommand>): void {
        if (this.editor.history.lastCommand instanceof type) {
            this.editor.history.undo();
            this.editor.history.clearRedos();
        }
    }





    /**
     * 执行撤销操作  并返回撤销的指令
     */
    public undo(): BasicCommand | null {
        if (this.disabled) {
            alert("Undo/Redo disabled while scene is playing.");
            return null;
        }
        let cmd = null;
        if (this.undos.length > 0) {
            cmd = this.undos.pop()!;
        }
        if (cmd != null) {
            cmd.undo();
            this.redos.push(cmd);
            // this._onChanged.dispatch(cmd);
        }
        return cmd;
    }

    /*
    * 执行重做操作 并返回重做的指令
    */
    public redo(): BasicCommand | null {
        if (this.disabled) {
            alert("Undo/Redo disabled while scene is playing.");
            return null;
        }
        let cmd: BasicCommand | null = null;
        if (this.redos.length > 0) {
            cmd = this.redos.pop()!;
        }
        if (cmd != null) {
            cmd!.execute();
            this.undos.push(cmd);
            // this._onChanged.dispatch(cmd);
        }
        return cmd;
    }


    public clearRedos(): void {
        this.redos = [];
    }



    /**
     * 清除所有（撤销、重做）历史指令
     */
    public clear(): void {
        this.undos = [];
        this.redos = [];
        this.idCounter = 0;
        // this._onChanged.dispatch();
    }


}