
/* eslint-disable @typescript-eslint/no-explicit-any */

export class Exception {
    [name: string]: any;

    public readonly host: string;

    public readonly code: number;

    public readonly message: string;

    /**
     *
     */
    protected constructor(host: string, message: string, code?: number) {
        this.host = host;
        this.code = code != null ? code : -1;
        this.message = message
    }



    public static build(host: string, message: string, code?: number): Exception {
        return new Exception(host, message, code);
    }

    // eslint-disable-next-line @typescript-eslint/explicit-module-boundary-types
    public static fromCatch(host: string, ex: any, message: string): Exception {
        const exc = new Exception(host, message);
        exc.copy(ex);
        return exc;
    }


    private copy(target: Object) {
        if (target == null) return;
        const keys = Object.keys(target)
        for (let i = 0; i < keys.length; i++) {
            this[keys[i]] = target[keys[i] as keyof Object];
        }
    }



    public toString(): string {
        return `[${this.host}:${this.code}]=>${this.message}`;
    }


}