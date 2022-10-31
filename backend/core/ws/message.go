package ws

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
	// 传输数据
	TransferData = RequestAction(0)
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

type ResponseMessage struct {
	// 响应代码
	Code ResponseCode `json:"code"`
	// 请求ID
	TraceId string `json:"traceId"`
	// 响应数据
	Data []byte `json:"data"`
}
