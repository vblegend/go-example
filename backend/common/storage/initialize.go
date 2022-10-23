/*
 * @Author: lwnmengjing
 * @Date: 2021/6/10 3:39 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/10 3:39 下午
 */

package storage

import (
	"backend/core/sdk"
	"backend/core/sdk/config"
	"backend/core/storage/cache"
)

// Setup 配置storage组件
func Setup() {
	// 设置缓存
	sdk.Runtime.SetCacheAdapter(cache.NewMemory())
	// 设置队列
	if !config.QueueConfig.Empty() {
		queueAdapter := config.QueueConfig.Setup()
		sdk.Runtime.SetQueueAdapter(queueAdapter)
		defer func() {
			go queueAdapter.Run()
		}()
	}

}
