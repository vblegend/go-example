package initialize

import (
	"server/common/config"
	"server/sugar/sdk"
	"server/sugar/storage"
)

// Setup 配置storage组件
func InitCache() {
	// 设置缓存
	sdk.Runtime.SetCacheAdapter(storage.NewMemoryCache())
	// 设置队列
	if config.Queue != nil && config.Queue.PoolSize > 0 {
		queueAdapter := storage.NewMemoryQueue(config.Queue.PoolSize)
		sdk.Runtime.SetQueueAdapter(queueAdapter)
		defer func() {
			go queueAdapter.Run()
		}()
	}

}
