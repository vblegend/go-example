/* eslint-disable @typescript-eslint/no-explicit-any */

// eslint-disable-next-line @typescript-eslint/no-empty-interface
export interface AnyObject {
    [key: string]: AnyObject;
}
export interface StringNameValue {
    name: string;
    value: string;
}

export interface nzSelectItem {
    label: string;
    value: any;
}

export enum NotificationType {
    Success = 'success',
    Info = 'info',
    Warning = 'warning',
    Error = 'error',
    Blank = 'blank'
}

/**
 * Getter函数的委托声明
 */
export declare type GetterType<T> = () => T;

/**
 * Setter函数的委托声明
 */
export declare type SetterType<T> = (value: T) => void;

/**
 * 数组 
 */
export declare type Array<T> = T[];