/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable @typescript-eslint/explicit-module-boundary-types */
/**
 * 
 * 
 * 
 * 
 */

export class ObjectUtil {



    /**
     * regardless of object references
     * determine whether the contents of two objects are equal 
     * @param object1 
     * @param object2 
     * @returns 
     */
    public static equals<T>(object1: T, object2: T): boolean {
        const entrie1 = Object.entries(object1).toString();
        const entrie2 = Object.entries(object2).toString();
        return entrie1 === entrie2;
    }

    /**
     *  deep freeze object
     * @param target 
     */
    public static freeze<T>(target: T): void {
        if (typeof target === 'object') {
            if (target instanceof Array) {
                for (let i = 0; i < target.length; i++) {
                    if (typeof target[i] === 'object') this.freeze(target[i]);
                }
            } else {
                const keys = Object.keys(target);
                for (let i = 0; i < keys.length; i++) {
                    const key = keys[i];
                    if (typeof target[key as keyof T] === 'object') this.freeze(target[key as keyof T]);
                }
            }
            Object.freeze(target);
        }
    }



    /**
     * Merge object properties to target object 
     * @param target 
     * @param object 
     * @param override Overwrite existing options 
     */
    public static merge<T extends Object>(target: T, object: any, override?: boolean): void {
        if (target == null || object == null) return;
        const deObject = this.clone(object);
        const keys = Object.keys(deObject);
        for (let i = 0; i < keys.length; i++) {
            const key = keys[i];
            const hasBarProperty = Object.prototype.hasOwnProperty.call(target, key);
            if (!hasBarProperty || override) {
                target[key as keyof T] = deObject[key];
            }
        }
    }



    /**
     * deep clone object 
     * @param target 
     * @param thisContext 
     * @returns 
     */
    public static clone<T>(target: T, thisContext?: Object): T | undefined | null {
        if (typeof target === 'undefined') return undefined;
        if (typeof target === 'string') return target.substring(0) as any;
        if (typeof target === 'number') return target;
        if (typeof target === 'boolean') return target;
        if (typeof target === 'function') return target.bind(thisContext);
        if (target === null) return null;
        let result: Record<string, any>;
        if (target instanceof Array) {
            result = [];
            for (let i = 0; i < target.length; i++) {
                result.push(this.clone(target[i]));
            }
        } else {
            result = {};
            const keys = Object.keys(target);
            for (let i = 0; i < keys.length; i++) {
                const key = keys[i];
                result[key] = this.clone(target[key as keyof T], result);
            }
        }
        return <T>result
    }


    /**
     * upgrade object\
     * 当_target字段为undefined时，赋值为_default中的字段\
     * @param _target 
     * @param _default 
     */
    public static upgrade<T>(_target: T, _default: T): void {
        if (typeof _target === 'undefined') return;
        if (typeof _target === 'string') return;
        if (typeof _target === 'number') return;
        if (typeof _target === 'boolean') return;
        if (typeof _target === 'function') return;
        if (_target === null) return;
        for (const key in _default) {
            if (_default[key] === undefined) continue;
            if (_target[key] === undefined) {
                _target[key as keyof T] = ObjectUtil.clone(_default[key])!;
            } else {
                this.upgrade(_target[key], _default[key]);
            }
        }
    }








}