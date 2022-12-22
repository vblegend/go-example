

export interface LoginRequest {
    username: string;
    password: string;
}

export interface LoginResponse {
    expire: string;
    nice: string;
    token: string;
    username: string;
    userId: number;
}


export interface RestfulResponse<T> {
    traceId: string;
    code: number;
    msg: string;
    data: T;
}
