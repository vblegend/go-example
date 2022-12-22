import { ObjectUtil } from "@core/util/object.util";
import { Observable, Observer } from "rxjs";

export type DataCacheChangedHandle = () => void;


export enum DataChangeTypes {
    Add,
    Remove,
    Update
}

export interface CacheChangeEvents<T> {
    /**
     * data change event typed 
     */
    type: DataChangeTypes;

    /**
     * datas form updata
     */
    data: T[];
}


class SubscriptionContext<TObject> {
    /**
     *
     */
    constructor(private thisContext: Object,
        private callback: (arg: CacheChangeEvents<TObject>) => void,
        private filter: (value: TObject, index: number, array: TObject[]) => boolean) {
    }


    public exec(type: DataChangeTypes, data: TObject[]): void {
        if (data == null || type == null) return;
        if (this.filter) data = data.filter(this.filter);
        if (data.length > 0) {
            const userdata: CacheChangeEvents<TObject> = { type, data };
            this.callback?.apply(this.thisContext, [userdata]);
        }
    }


}

export declare interface LocalCacheType<TObject, TKey> extends Function {
    new(tableName: string, primaryKey: (ev: TObject) => TKey, writable: boolean): LocalCache<TObject, TKey>;
}

/**
 * Local data cache manager 
 */
export class LocalCache<TObject, TKey> {
    public readonly tableName: string;
    private readonly _indexs: Map<TKey, number> = new Map();
    private readonly _buffer: TObject[] = [];
    private readonly _primaryKeyFunc: (ev: TObject) => TKey;
    private readonly _subscribers: SubscriptionContext<TObject>[] = [];


    /**
     * 
     * @param tableName 
     * @param primaryKey 
     * @param writable 
     */
    public constructor(tableName: string, primaryKey: (ev: TObject) => TKey, private writable: boolean = false) {
        this.tableName = tableName;
        this._primaryKeyFunc = primaryKey;
    }

    /**
     * subscribe data change
     */
    public subscribe(callback: (arg: CacheChangeEvents<TObject>) => void, filter?: (value: TObject, index: number, array: TObject[]) => boolean): () => void {
        const sub = new SubscriptionContext<TObject>(this, callback, filter!);
        this._subscribers.push(sub);
        return () => {
            const index = this._subscribers.indexOf(sub);
            if (index > -1) this._subscribers.splice(index, 1);
        };
    }

    /**
     * emit change event
     * @param type 
     * @param data 
     */
    private emit(type: DataChangeTypes, data: TObject[]): void {
        if (this._subscribers.length == 0) return;
        Object.freeze(data);
        const ems = this._subscribers.slice();
        for (let i = 0; i < ems.length; i++) {
            ems[i].exec(type, data);
        }
        // Object.freeze(data);
        ems.length = 0;
    }


    /**
     * load data from object array
     * @param entries 
     */
    public load(entries: TObject[]): void {
        this.batchPut(entries);
    }


    /**
     * batch update object
     * @param entries 
     * @returns 
     */
    public batchPut(entries: TObject[]): boolean {
        if (entries == null) return false;
        const updates: TObject[] = [];
        const inserts: TObject[] = [];
        for (let i = 0; i < entries.length; i++) {
            const key = this._primaryKeyFunc(entries[i]);
            const duplicate = ObjectUtil.clone(entries[i]);
            if (!this.writable) ObjectUtil.freeze(duplicate);
            if (key == null) return false;
            const index = this._indexs.get(key);
            if (index == null) {
                // append record
                this._buffer.push(duplicate!);
                this._indexs.set(key, this._buffer.length - 1);
                inserts.push(duplicate!);
            } else {
                // override record
                this._buffer.splice(index, 1, duplicate!);
                // this._buffer[index] = duplicate;
                updates.push(duplicate!);
            }
        }
        if (inserts.length > 0) this.emit(DataChangeTypes.Add, inserts);
        if (updates.length > 0) this.emit(DataChangeTypes.Update, updates);
        return true;
    }


    public put(entrie: TObject): boolean {
        return this.batchPut([entrie]);
    }

    public remove(key: TKey): TObject | null {
        const index = this._indexs.get(key);
        if (index == null) return null;
        const result = this._buffer[index];
        this._buffer.splice(index, 1);
        this._indexs.delete(key);
        this.emit(DataChangeTypes.Remove, [result]);
        return result;
    }

    public clear(): void {
        this.emit(DataChangeTypes.Remove, this._buffer.slice());
        this._indexs.clear();
        this._buffer.length = 0;
    }

    public filter(predicate: (value: TObject, index: number, array: TObject[]) => value is TObject, thisArg?: Object): TObject[] {
        return this._buffer.filter(predicate, thisArg);
    }


    public find(key: TKey): TObject | null {
        const index = this._indexs.get(key);
        if (index == null) return null;
        return this._buffer[index];
    }


    public getAll(): TObject[] {
        return this._buffer.slice();
    }

    public dispose(): void {
        if (this._subscribers) {
            this._subscribers.length = 0;
        }
        if (this._indexs) {
            this._indexs.clear();
        }
        if (this._buffer) {
            this._buffer.length = 0;
        }
    }



}