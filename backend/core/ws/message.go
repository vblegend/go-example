package ws

import (
	"encoding/json"
	"errors"
)

var ErrorJsonUnmarshalFail = errors.New("invalid json string")

type MessageType int

const (
	TextMessage = MessageType(1)

	// BinaryMessage denotes a binary data message.
	BinaryMessage = MessageType(2)

	// CloseMessage denotes a close control message. The optional message
	// payload contains a numeric code and text. Use the FormatCloseMessage
	// function to format a close message payload.
	CloseMessage = MessageType(8)

	// PingMessage denotes a ping control message. The optional message payload
	// is UTF-8 encoded text.
	PingMessage = MessageType(9)

	// PongMessage denotes a pong control message. The optional message payload
	// is UTF-8 encoded text.
	PongMessage = MessageType(10)
)

type RequestAction int8
type ResponseCode int8

const (
	// 操作成功
	Success = ResponseCode(0)
	// 操作失败
	Failure = ResponseCode(-1)
)

const (
	// 加入频道
	JoinChannel = RequestAction(-127)
	// 离开频道
	LevelChannel = RequestAction(127)
	// 传输数据 仅发送 不需要考虑响应
	TransferPost = RequestAction(64)
	// 传输数据 发送后需要回应
	TransferSend = RequestAction(-64)
)

type RequestMessage struct {
	// 动作
	Action RequestAction `json:"action"`
	// 消息属于哪个频道
	Channel string `json:"channel"`
	// 请求ID
	TraceId string `json:"traceId"`
	// 数据负载
	Payload []byte `json:"payload"`
}

func (r *RequestMessage) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return err
	}
	// Name不能为空
	if r.Action == 0 {
		return ErrorJsonUnmarshalFail
	}
	if r.Channel == "" {
		return ErrorJsonUnmarshalFail
	}
	if r.TraceId == "" {
		return ErrorJsonUnmarshalFail
	}
	return nil
}

type ResponseMessage struct {
	// 响应代码
	Code ResponseCode `json:"code"`
	// 请求ID
	TraceId string `json:"traceId"`
	// 响应数据
	Data []byte `json:"data"`
}
