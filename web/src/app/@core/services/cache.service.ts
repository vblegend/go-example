import { Injectable, Injector } from '@angular/core';
import { LocalCache } from '@core/cache/local.cache';
import { Exception } from '@core/common/exception';
import { SchedulingTask } from 'app/pages/tasks/task-model/tasks.model';



export interface CacheEKeyMaps {
    tasks: LocalCache<SchedulingTask, number>;
    users: LocalCache<SchedulingTask, number>;
}


/**
 * 数据缓存服务 
 */
@Injectable({
    providedIn: 'root'
})
export class CacheService {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    private _caches: Map<string, any> = new Map();
    constructor(protected injector: Injector) {
    }

    //, 
    public register<K extends keyof CacheEKeyMaps>(type: K, cache: CacheEKeyMaps[K]): void {
        if (this._caches.has(type)) {
            throw Exception.build('', '');
        }
        this._caches.set(type, cache);
    }

    public get<K extends keyof CacheEKeyMaps>(type: K): CacheEKeyMaps[K] {
        return this._caches.get(type) as CacheEKeyMaps[K];
    }
}