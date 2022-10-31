package ws

import (
	"backend/core/mpool"
	"time"
)

var msgpool = mpool.NewObjectPool(&mpool.Options{
	Capacity: 100,
	MaxIdle:  90,
	New: func() interface{} {
		return &RequestMessage{}
	},
	MinIdleTime: time.Hour,
})

func MallocMessage() (*RequestMessage, error) {
	msg, err := msgpool.Malloc()
	if err == nil {
		return msg.(*RequestMessage), nil
	}
	return nil, err
}

func FreeMessage(msg *RequestMessage) error {
	return msgpool.Free(msg)
}
