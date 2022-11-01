package ws

import "errors"

var InvalidChannelName = errors.New("invalid channel name")
var NotInChannel = errors.New("not in channel")
var ErrorCannotJoinChannelRepeated = errors.New("cannot join channel repeated")

type WSEventListener interface {
	// websocket  连接建立
	OnJoin(client *WSClient)
	// websocket  连接断开
	OnLeave(client *WSClient)

	// websocket  连接断开
	OnMessage(client *WSClient, msg *RequestMessage)
}

type IWSChannel interface {
	Name() string
	OnJoin(client *WSClient)
	// websocket  连接断开
	OnLeave(client *WSClient)
	// websocket  连接断开
	OnMessagePost(client *WSClient, msg *RequestMessage)
	OnMessageCall(client *WSClient, msg *RequestMessage) (*ResponseMessage, error)
}

type IChannelCollection interface {
	JoinChannel(channel *WSChannel)
	LeaveChannel(channel *WSChannel)
	HasChannel(channelName string) bool
}
