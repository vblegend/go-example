package runtime

import (
	"net/http"
	"sync"

	"github.com/go-redis/redis/v7"
	"gorm.io/gorm"

	"backend/core/storage"

	"github.com/robfig/cron/v3"
)

type Application struct {
	dbs         map[string]*gorm.DB
	engine      http.Handler
	crontab     map[string]*cron.Cron
	mux         sync.RWMutex
	memoryCache storage.AdapterCache
	memoryQueue storage.AdapterQueue
	redis       *redis.Client
}

type Router struct {
	HttpMethod, RelativePath, Handler string
}

type Routers struct {
	List []Router
}

// SetDb 设置对应key的db
func (e *Application) SetDb(key string, db *gorm.DB) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.dbs[key] = db
}

// GetDbByKey 根据key获取db
func (e *Application) GetDb(key string) *gorm.DB {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.dbs[key]
}

// SetEngine 设置路由引擎
func (e *Application) SetEngine(engine http.Handler) {
	e.engine = engine
}

// GetEngine 获取路由引擎
func (e *Application) GetEngine() http.Handler {
	return e.engine
}

// NewConfig 默认值

func NewConfig() *Application {
	return &Application{
		dbs:     make(map[string]*gorm.DB),
		crontab: make(map[string]*cron.Cron),
	}
}

func (e *Application) SetRedisClient(redis *redis.Client) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.redis = redis
}
func (e *Application) GetRedisClient() *redis.Client {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.redis
}

// SetCrontab 设置对应key的crontab
func (e *Application) SetCrontab(key string, crontab *cron.Cron) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.crontab[key] = crontab
}

// GetCrontabKey 根据key获取crontab
func (e *Application) GetCrontab(key string) *cron.Cron {
	e.mux.Lock()
	defer e.mux.Unlock()
	if e, ok := e.crontab["*"]; ok {
		return e
	}
	return e.crontab[key]
}

func (e *Application) SetCacheAdapter(cache storage.AdapterCache) {
	e.memoryCache = cache
}

func (e *Application) GetMemoryCache(key string) storage.AdapterCache {
	return NewCache(key, e.memoryCache)
}

func (e *Application) SetQueueAdapter(queue storage.AdapterQueue) {
	e.memoryQueue = queue
}

func (e *Application) GetMemoryQueue(prefix string) storage.AdapterQueue {
	return NewQueue(prefix, e.memoryQueue)
}
