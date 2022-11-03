package mpool

import (
	"context"
	"time"

	pool "github.com/jolestar/go-commons-pool"
)

type NewObjectFunc func() interface{}

type ObjectPoolFactory struct {
	new NewObjectFunc
}

func (f *ObjectPoolFactory) MakeObject(ctx context.Context) (*pool.PooledObject, error) {
	return pool.NewPooledObject(f.new()), nil
}

func (f *ObjectPoolFactory) DestroyObject(ctx context.Context, object *pool.PooledObject) error {
	// do destroy
	return nil
}

func (f *ObjectPoolFactory) ValidateObject(ctx context.Context, object *pool.PooledObject) bool {
	// do validate
	return true
}

func (f *ObjectPoolFactory) ActivateObject(ctx context.Context, object *pool.PooledObject) error {
	// do activate
	return nil
}

func (f *ObjectPoolFactory) PassivateObject(ctx context.Context, object *pool.PooledObject) error {
	// do passivate
	return nil
}

type objectPool struct {
	pool *pool.ObjectPool
	ctx  context.Context
}

func (p *objectPool) setOptions(options *Options) {
	if options.Ctx == nil {
		options.Ctx = context.Background()
	}

	p.ctx = options.Ctx
	p.pool = pool.NewObjectPoolWithDefaultConfig(options.Ctx, &ObjectPoolFactory{new: options.New})
	// 连接池允许分配对象的上限。为负值时表示无限制
	p.pool.Config.MaxTotal = options.Capacity
	// 连接池中空闲对象的上限。如果 MaxIdle在重负载系统上设置得太低，可能会看到对象被销毁后，立即有新对象被创建。这是活跃的 goroutine 返回请求结果的速度比发起请求的速度更快，导致空闲的数量超过 maxIdle 的对象。 maxIdle 的最佳值在不同的系统会有所不同。
	p.pool.Config.MaxIdle = options.MaxIdle
	// MinIdle：连接池要维护的最小空闲对象数此设置仅TimeBetweenEvictionRuns 大于零的情况下有效。用于尝试确保连接池在空闲对象驱逐运行期间保留所需的最小实例数如果 MinIdle 的配置值大于 MaxIdle 的配置值则将使用 MaxIdle 的值。
	p.pool.Config.MinIdle = 0
	// 当活跃的商品池已经达到最大参数的限制时，执行 ObjectPool.BorrowObject() 是否阻塞。
	p.pool.Config.BlockWhenExhausted = false
	p.pool.Config.TestOnBorrow = false
	p.pool.Config.TestOnCreate = false
	p.pool.Config.TestOnReturn = false
	p.pool.Config.TestWhileIdle = false
	// 对象可以在池中闲置的最短时间，如果该参数为正值，对象驱逐器执行过程中会将存活时间超过这个参数值的对象销毁。若为非正值则不会因为空闲时间而将任何对象从池中逐出。
	p.pool.Config.MinEvictableIdleTime = time.Hour
	// 当前空闲连接数量大于 MinIdle 时，对象可以在池中闲置的最短时间如果MinEvictableIdleTime 为正数，SoftMinEvictableIdleTime 被忽略。
	p.pool.Config.SoftMinEvictableIdleTime = options.MinIdleTime
	// 每次驱逐器运行期间要检查的最大对象数（如果驱逐策略开启）如果该参数为正值，检测对象数量为该参数值和当前空闲连接数中较小的值。如果该参数为负值，检测次数执行将是 math.Ceil(ObjectPool.GetNumIdle()/math.Abs(PoolConfig.NumTestsPerEvictionRun)) 这意味着当值为 -n 时大约有 n 分之一的空闲对象将被检测。
	p.pool.Config.NumTestsPerEvictionRun = 1024

	// p.pool.Config.TimeBetweenEvictionRuns = time.Second * 30
}

func (p *objectPool) MallocWithContext(ctx context.Context) (obj interface{}, err error) {
	if p.pool == nil {
		return nil, ErrorOfDestroyed
	}
	defer func() {
		ERR := recover()
		if ERR != nil {
			err = ErrorOfNotSpace
		}
	}()
	obj, err = p.pool.BorrowObject(ctx)
	return
}

func (p *objectPool) FreeWithContext(ctx context.Context, v interface{}) error {
	if p.pool == nil {
		return ErrorOfDestroyed
	}
	p.pool.ReturnObject(ctx, v)
	return nil
}

func (p *objectPool) Malloc() (obj interface{}, err error) {
	if p.pool == nil {
		return nil, ErrorOfDestroyed
	}
	defer func() {
		ERR := recover()
		if ERR != nil {
			err = ErrorOfNotSpace
		}
	}()
	obj, err = p.pool.BorrowObject(p.ctx)
	return
}

func (p *objectPool) Free(v interface{}) error {
	if p.pool == nil {
		return ErrorOfDestroyed
	}
	p.pool.ReturnObject(p.ctx, v)
	return nil
}

func (p *objectPool) GetCapacity() int {
	return p.pool.Config.MaxTotal
}

func (p *objectPool) GetNumActive() int {
	return p.pool.GetNumActive()
}

func (p *objectPool) GetNumIdle() int {
	return p.pool.GetNumActive()
}

func (p *objectPool) Destroy() {
	p.pool.Close(p.ctx)
	p.ctx = nil
	p.pool = nil

}

func (p *objectPool) Clear() error {
	if p.pool == nil {
		return ErrorOfDestroyed
	}
	p.pool.Clear(p.ctx)
	return nil
}

func (p *objectPool) ClearWithContext(ctx context.Context) error {
	if p.pool == nil {
		return ErrorOfDestroyed
	}
	p.pool.Clear(ctx)
	return nil
}
