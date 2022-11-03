package ws

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var errorJSONUnmarshalFail = errors.New("invalid json string")
var errorInvalidChannelName = errors.New("invalid channel name")
var errorNotInChannel = errors.New("not in channel")
var errorCannotJoinChannelRepeated = errors.New("cannot join channel repeated")

// IWSManager websocket 管理器
type IWSManager interface {
	// AcceptHandler
	AcceptHandler(c *gin.Context)
	// Broadcast
	Broadcast(msg *ResponseMessage)
	// RegisterChannel
	RegisterChannel(name string, handler IWSMessageHandler, perm AuthType) error
}

// IWSClient websocket 客户端
type IWSClient interface {

	// ClientID
	ClientID() string
	// Write
	Write(msg *ResponseMessage) error
	// OK
	OK(traceID string, data []byte, message string) error
	// Error
	Error(traceID string, data error) error
	// Success
	Success(traceID string, message string, data []byte) error
	// Close
	Close() error
}

// IWSChannel websocket 频道
type IWSChannel interface {
	Name() string
	KickedOut(client IWSClient)
	Broadcast(msg *ResponseMessage)
	GetClient(clientID string) IWSClient
	Length() int
}

// IWSMessageHandler websocket 频道
type IWSMessageHandler interface {
	OnJoin(channel IWSChannel, client IWSClient, params Params) error
	// websocket  连接断开
	OnLeave(channel IWSChannel, client IWSClient)
	// websocket  连接断开
	OnMessagePost(channel IWSChannel, client IWSClient, msg *RequestMessage)
	OnMessageCall(channel IWSChannel, client IWSClient, msg *RequestMessage) (*ResponseMessage, error)
}

// Default 默认的 websocket 管理器
var Default IWSManager = NewWebSocketManager()

// NewWebSocketManager 创建一个websocket 管理器
func NewWebSocketManager() IWSManager {
	ws := &wsManager{}
	ws.channels = make(map[string]*WSChannel)
	return ws
}
