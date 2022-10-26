package runtime

import (
	"net/http"

	"github.com/go-redis/redis/v7"

	"backend/core/storage"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type Runtime interface {
	SetDb(key string, db *gorm.DB)
	GetDb(key string) *gorm.DB

	SetRedisClient(redis *redis.Client)
	GetRedisClient() *redis.Client

	SetEngine(engine http.Handler)
	GetEngine() http.Handler

	SetCrontab(key string, crontab *cron.Cron)
	GetCrontab(key string) *cron.Cron

	SetCacheAdapter(storage.AdapterCache)
	GetMemoryCache(string) storage.AdapterCache

	SetQueueAdapter(storage.AdapterQueue)
	GetMemoryQueue(string) storage.AdapterQueue
}
