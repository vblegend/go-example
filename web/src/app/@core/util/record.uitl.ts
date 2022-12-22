

export class RecordUtil {



    /**
     *  从数组转换一个HashMap 对象
     * @param array 数组对象
     * @param keyFunc  获取记录主键
     * @returns 
     */
    public static fromArray<TKey extends keyof TType, TType>(array: TType[], keyFunc: (v: TType) => TKey): Record<TKey, TType> {
        const result: Record<TKey, TType> = <Record<TKey, TType>>{};
        for (const item of array) {
            if (item) {
                const key = keyFunc(item);
                result[key] = item;
            }
        }
        return result;
    }


    public static intoFromArray<TKey extends keyof TType, TType>(result: Record<TKey, TType>, array: TType[], keyFunc: (v: TType) => TKey): void {
        for (const item of array) {
            if (item) {
                const key = keyFunc(item);
                result[key] = item;
            }
        }
    }



    public static from<TKey extends keyof TType, TType, TValue>(array: TType[], keyFunc: (v: TType) => TKey, valueFunc: (v: TType) => TValue): Record<TKey, TValue> {
        const result: Record<TKey, TValue> = <Record<TKey, TValue>>{};
        for (const item of array) {
            if (item) {
                const key = keyFunc(item);
                result[key] = valueFunc(item);
            }
        }
        return result;
    }

    public static toArray<TType>(record: Record<string, TType>): TType[] {
        const result = [];
        for (const key in record) {
            result.push(record[key]);
        }
        return result;
    }

    /**
     * 拷贝对象内属性到另一个对象中
     * @param target 
     * @param source 
     */
    public static copy<TObject>(source: TObject, target: TObject): void {
        if (target == null) throw new Error('Parameter ‘target’ cannot be null');
        if (source == null) throw new Error('Parameter ‘source’ be null');
        for (const key in source) {
            target[key] = source[key];
        }
    }

    /**
     * 返回一个克隆的对象，浅克隆
     * @param source 
     * @returns 
     */
    public static clone<TObject>(source: TObject): TObject {
        const result: TObject = <TObject>{};
        if (source == null) throw new Error('Parameter ‘source’ be null');
        for (const key in source) {
            result[key] = source[key];
        }
        return result;
    }


    /**
     * 清除对象内所有属性
     * @param target 
     */
    public static clear<TObject>(target: TObject): void {
        if (target == null) throw new Error('Parameter ‘target’ cannot be null');
        for (const key in target) {
            delete target[key];
        }
    }
}