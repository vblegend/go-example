/* eslint-disable @typescript-eslint/no-explicit-any */
export declare type Delegate = () => void;
export declare type CancelFunc = () => void;
export declare type Action = (...params: any[]) => void;
export interface DelegateContext {
    delegate: Delegate;
    context: any;
}

