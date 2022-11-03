package mpool

import (
	"context"
	"errors"
	"time"
)

var (
	ErrorOfNotSpace  = errors.New("there is no spare capacity position")
	ErrorOfDestroyed = errors.New("the object is destroyed")
)

type IPool interface {
	MallocWithContext(ctx context.Context) (interface{}, error)
	FreeWithContext(ctx context.Context, v interface{}) error
	Malloc() (interface{}, error)
	Free(v interface{}) error
	GetCapacity() int
	GetNumActive() int
	GetNumIdle() int
	Destroy()
	Clear() error
	ClearWithContext(ctx context.Context) error
}

type Options struct {
	// 池子所属上下文
	Ctx context.Context
	// 对象生成器
	New func() interface{}
	// 池子最大容量
	Capacity int
	// 归还时空闲数大于该数值 对象自动释放
	MaxIdle int
	// 最小的空闲时间，空闲时间大于该数值时自动释放
	MinIdleTime time.Duration
}
