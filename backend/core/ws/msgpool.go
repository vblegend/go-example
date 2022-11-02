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

// MallocRequestMessage 申请一个请求消息对象
func MallocRequestMessage() (*RequestMessage, error) {
	msg, err := requestPool.Malloc()
	if err == nil {
		return msg.(*RequestMessage), nil
	}
	return nil, err
}

// FreeRequestMessage 释放一个请求消息对象
func FreeRequestMessage(msg *RequestMessage) error {
	if !msg.managed {
		return nil
	}
	msg.Payload = ""
	msg.Channel = ""
	msg.TraceID = ""
	msg.Method = ""
	msg.Action = 0
	return requestPool.Free(msg)
}

// MallocResponseMessage 申请一个响应消息对象
func MallocResponseMessage() (*ResponseMessage, error) {
	msg, err := responsePool.Malloc()
	if err == nil {
		return msg.(*ResponseMessage), nil
	}
	return nil, err
}

// FreeResponseMessage 释放一个响应消息对象
func FreeResponseMessage(msg *ResponseMessage) error {
	if !msg.managed {
		return nil
	}
	msg.Payload = ""
	msg.Code = 0
	msg.TraceID = ""
	return responsePool.Free(msg)
}
