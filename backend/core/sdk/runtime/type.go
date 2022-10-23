package runtime

import (
	"net/http"

	"github.com/go-redis/redis/v7"

	"backend/core/logger"
	"backend/core/storage"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type Runtime interface {
	// SetDb 多db设置，⚠️SetDbs不允许并发,可以根据自己的业务，例如app分库、host分库
	SetDb(key string, db *gorm.DB)
	GetDb() map[string]*gorm.DB
	GetDbByKey(key string) *gorm.DB

	SetRedisClient(redis *redis.Client)
	GetRedisClient() *redis.Client
	// SetEngine 使用的路由
	SetEngine(engine http.Handler)
	GetEngine() http.Handler

	GetRouter() []Router

	// SetLogger 使用backend定义的logger，参考来源go-micro
	SetLogger(logger logger.Logger)
	GetLogger() logger.Logger

	// SetCrontab crontab
	SetCrontab(key string, crontab *cron.Cron)
	GetCrontab() map[string]*cron.Cron
	GetCrontabKey(key string) *cron.Cron

	// SetMiddleware middleware
	// SetMiddleware(string, interface{})
	// GetMiddleware() map[string]interface{}
	// GetMiddlewareKey(key string) interface{}

	// SetCacheAdapter cache
	SetCacheAdapter(storage.AdapterCache)
	SetQueueAdapter(storage.AdapterQueue)

	GetMemoryCache(string) storage.AdapterCache
	GetMemoryQueue(string) storage.AdapterQueue
}
