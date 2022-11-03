package mpool

func NewObjectPool(options *Options) IPool {
	pool := objectPool{}
	pool.setOptions(options)
	return &pool
}
