/* eslint-disable @typescript-eslint/no-explicit-any */


/**
 * 注解对象的方法使对象方法可以通过 for(const key in object) 枚举到
 * @param value 是否可以被 for in 枚举  默认为false
 * @returns 
 */
 export function Enumable(value: boolean = false): (target: any, methodName: string, descriptor: PropertyDescriptor) => void {
    return function (target: any, methodName: string, descriptor: PropertyDescriptor) {
        descriptor.enumerable = value;
    };
}