package ws

import (
	"context"
	"encoding/json"

	"github.com/gorilla/websocket"
)

type WSClient struct {
	Socket  *websocket.Conn
	Context context.Context
	Cancel  context.CancelFunc
	// 当前连接唯一ID
	ClientId string
	channels map[string]IWSChannel
}

func NewWSClient(conn *websocket.Conn, ctx context.Context, cancel context.CancelFunc, clientId string) *WSClient {
	wsc := WSClient{
		Socket:   conn,
		Context:  ctx,
		Cancel:   cancel,
		ClientId: clientId,
		channels: make(map[string]IWSChannel),
	}
	return &wsc
}

func (wsc *WSClient) JoinChannel(channel IWSChannel) {
	chanl := wsc.channels[channel.Name()]
	if chanl == nil {
		wsc.channels[channel.Name()] = channel
	}
}

func (wsc *WSClient) HasChannel(channelName string) bool {
	return wsc.channels[channelName] != nil
}

func (wsc *WSClient) LeaveChannel(channel IWSChannel) {
	chanl := wsc.channels[channel.Name()]
	if chanl != nil {
		delete(wsc.channels, channel.Name())
	}
}

func (wsc *WSClient) Success(traceId string, message string, data []byte) error {
	if msg, err := MallocResponseMessage(); err == nil {
		msg.Code = Success
		msg.TraceId = traceId
		msg.Payload = PayloadDomain(data)
		msg.Message = message
		if bytes, err := json.Marshal(msg); err == nil {
			err = wsc.Socket.WriteMessage(int(TextMessage), bytes)
		}
		FreeResponseMessage(msg)
		return err
	}
	return nil
}
func (wsc *WSClient) Write(msg *ResponseMessage) error {
	defer func() {
		FreeResponseMessage(msg)
	}()
	var err error
	var bytes []byte
	if bytes, err = json.Marshal(msg); err == nil {
		err = wsc.Socket.WriteMessage(int(TextMessage), bytes)
	}
	return err
}

func (wsc *WSClient) OK(traceId string, data []byte, message string) error {
	if msg, err := MallocResponseMessage(); err == nil {
		msg.Code = Success
		msg.TraceId = traceId
		msg.Payload = PayloadDomain(data)
		msg.Message = message
		if bytes, err := json.Marshal(msg); err == nil {
			err = wsc.Socket.WriteMessage(int(TextMessage), bytes)
		}
		FreeResponseMessage(msg)
		return err
	}
	return nil
}

func (wsc *WSClient) Error(traceId string, data error) error {
	if msg, err := MallocResponseMessage(); err == nil {
		msg.Code = Failure
		msg.TraceId = traceId
		msg.Message = data.Error()
		if bytes, err := json.Marshal(msg); err == nil {
			err = wsc.Socket.WriteMessage(int(TextMessage), bytes)
		}
		FreeResponseMessage(msg)
		return err
	}
	return nil
}

func (wsc *WSClient) Close() error {
	wsc.Socket.WriteMessage(int(CloseMessage), nil)
	wsc.Cancel()
	return wsc.Socket.Close()
}
