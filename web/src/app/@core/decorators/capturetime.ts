/* eslint-disable @typescript-eslint/no-explicit-any */
import { environment } from "environments/environment";



/**
 * 监控对象异步方法的执行时间\
 * 在生产模式时runOfProduction=false的装饰器对性能无影响
 * @param runOfProduction 是否可以运行在生产环境，默认值：false
 * @logOfMillisecond 打印条件执行大于等于 logOfMillisecond 毫秒 才输出显示  默认值 0
 * @returns 
 */
export function CaptureTimeAsync(runOfProduction?: boolean, logOfMillisecond: number = 0): (target: any, methodName: string, descriptor: PropertyDescriptor) => void {
    return function (target: Function, methodName: string, descriptor: PropertyDescriptor) {
        if (!environment.production || runOfProduction) {
            const original = descriptor.value;
            descriptor.value = async function (...args: any[]) {
                const before = window.performance.now();
                const result = await original.apply(this, args);
                const value = window.performance.now() - before;
                if (value >= logOfMillisecond) {
                    console.log(`%c${target.constructor.name}::${methodName}%c => ${value.toFixed(2)}ms`, "color:orange;font-weight:bold", "color:red;font-weight:bold");
                }
                return result;
            };
        }
    };
}




/**
 * 监控对象方法的执行时间\
 * 在生产模式时runOfProduction=false的装饰器对性能无影响
 * @param runOfProduction 是否可以运行在生产环境，默认值：false
 * @logOfMillisecond 打印条件执行大于等于 logOfMillisecond 毫秒 才输出显示  默认值 0
 * @returns 
 */
export function CaptureTime(runOfProduction?: boolean, logOfMillisecond: number = 0): (target: any, methodName: string, descriptor: PropertyDescriptor) => void {
    return function (target: Function, methodName: string, descriptor: PropertyDescriptor) {
        if (!environment.production || runOfProduction) {
            const original = descriptor.value;
            descriptor.value = function (...args: any[]) {
                const before = window.performance.now();
                const result = original.apply(this, args);
                const value = window.performance.now() - before;
                if (value >= logOfMillisecond) {
                    console.log(`%c${target.constructor.name}::${methodName}%c => ${value.toFixed(2)}ms`, "color:orange;font-weight:bold", "color:red;font-weight:bold");
                }
                return result;
            };
        }
    };
}

