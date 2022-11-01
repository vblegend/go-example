package ws

import (
	"backend/core/mpool"
	"time"
)

var requestPool = mpool.NewObjectPool(&mpool.Options{
	Capacity: 100,
	MaxIdle:  90,
	New: func() interface{} {
		return &RequestMessage{
			managed: true,
		}
	},
	MinIdleTime: time.Hour,
})

var responsePool = mpool.NewObjectPool(&mpool.Options{
	Capacity: 100,
	MaxIdle:  90,
	New: func() interface{} {
		return &ResponseMessage{
			managed: true,
		}
	},
	MinIdleTime: time.Hour,
})

func MallocRequestMessage() (*RequestMessage, error) {
	msg, err := requestPool.Malloc()
	if err == nil {
		return msg.(*RequestMessage), nil
	}
	return nil, err
}

func FreeRequestMessage(msg *RequestMessage) error {
	if !msg.managed {
		return nil
	}
	msg.Payload = nil
	msg.Channel = ""
	msg.TraceId = ""
	msg.Action = 0
	return requestPool.Free(msg)
}

func MallocResponseMessage() (*ResponseMessage, error) {
	msg, err := responsePool.Malloc()
	if err == nil {
		return msg.(*ResponseMessage), nil
	}
	return nil, err
}

func FreeResponseMessage(msg *ResponseMessage) error {
	if !msg.managed {
		return nil
	}
	msg.Data = nil
	msg.Code = 0
	msg.TraceId = ""
	return responsePool.Free(msg)
}