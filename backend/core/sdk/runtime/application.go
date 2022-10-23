package runtime

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"

	"backend/core/logger"

	"backend/core/storage"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type Application struct {
	dbs         map[string]*gorm.DB
	engine      http.Handler
	crontab     map[string]*cron.Cron
	mux         sync.RWMutex
	memoryCache storage.AdapterCache
	memoryQueue storage.AdapterQueue
	routers     []Router
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

// GetDb 获取所有map里的db数据
func (e *Application) GetDb() map[string]*gorm.DB {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.dbs
}

// GetDbByKey 根据key获取db
func (e *Application) GetDbByKey(key string) *gorm.DB {
	e.mux.Lock()
	defer e.mux.Unlock()
	// if db, ok := e.dbs["*"]; ok {
	// 	return db
	// }
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

// GetRouter 获取路由表
func (e *Application) GetRouter() []Router {
	return e.setRouter()
}

// setRouter 设置路由表
func (e *Application) setRouter() []Router {
	switch e.engine.(type) {
	case *gin.Engine:
		routers := e.engine.(*gin.Engine).Routes()
		for _, router := range routers {
			e.routers = append(e.routers, Router{RelativePath: router.Path, Handler: router.Handler, HttpMethod: router.Method})
		}
	}
	return e.routers
}

// SetLogger 设置日志组件
func (e *Application) SetLogger(l logger.Logger) {
	logger.DefaultLogger = l
}

// GetLogger 获取日志组件
func (e *Application) GetLogger() logger.Logger {
	return logger.DefaultLogger
}

// NewConfig 默认值

func NewConfig() *Application {
	return &Application{
		dbs:     make(map[string]*gorm.DB),
		crontab: make(map[string]*cron.Cron),
		routers: make([]Router, 0),
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

// GetCrontab 获取所有map里的crontab数据
func (e *Application) GetCrontab() map[string]*cron.Cron {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.crontab
}

// GetCrontabKey 根据key获取crontab
func (e *Application) GetCrontabKey(key string) *cron.Cron {
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
