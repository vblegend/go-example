/* eslint-disable @typescript-eslint/no-explicit-any */



export function Sealed(): (target: any, methodName: string, descriptor: PropertyDescriptor) => void {
    return function (target: Function, methodName: string, descriptor: PropertyDescriptor) {
        descriptor.writable = false;
        descriptor.configurable = false;
        descriptor.enumerable = false;
    };
}