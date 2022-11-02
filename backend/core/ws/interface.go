package ws

import "errors"

var InvalidChannelName = errors.New("invalid channel name")
var NotInChannel = errors.New("not in channel")
var ErrorCannotJoinChannelRepeated = errors.New("cannot join channel repeated")

type IWSChannel interface {
	Name() string

	OnMethodMap()
	OnJoin(client *WSClient, params Params) error
	// websocket  连接断开
	OnLeave(client *WSClient)
	// websocket  连接断开
	OnMessagePost(client *WSClient, msg *RequestMessage)
	OnMessageCall(client *WSClient, msg *RequestMessage) (*ResponseMessage, error)
	//踢出客户端
	KickedOut(client *WSClient)
}

type IChannelCollection interface {
	JoinChannel(channel *WSChannel)
	LeaveChannel(channel *WSChannel)
	HasChannel(channelName string) bool
}
