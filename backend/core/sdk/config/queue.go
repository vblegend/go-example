package config

import (
	"backend/core/storage"
	"backend/core/storage/queue"
)

type Queue struct {
	Memory *QueueMemory `yaml:"memory"`
}

type QueueMemory struct {
	PoolSize uint `yaml:"poolSize"`
}

var QueueConfig = new(Queue)

// Empty 空设置
func (e Queue) Empty() bool {
	return e.Memory == nil || e.Memory.PoolSize <= 0
}

// Setup 启用顺序 redis > 其他 > memory
func (e Queue) Setup() storage.AdapterQueue {
	return queue.NewMemory(e.Memory.PoolSize)
}
