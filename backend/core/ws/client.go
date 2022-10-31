package ws

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"

	"github.com/gorilla/websocket"
)

type WSClient struct {
	Socket  *websocket.Conn
	Context context.Context
	Cancel  context.CancelFunc

	// 当前连接唯一ID
	ConnectId string
	// 当前连接信道
	Channel string
	// 当前连接所有参数
	Params url.Values

	channels map[string]*WSChannel
}

func NewWSClient(conn *websocket.Conn, ctx context.Context, cancel context.CancelFunc, clientId string) *WSClient {
	wsc := WSClient{
		Socket:    conn,
		Context:   ctx,
		Cancel:    cancel,
		ConnectId: clientId,
		channels:  make(map[string]*WSChannel),
	}
	return &wsc
}

func (wsc *WSClient) JoinChannel(channel *WSChannel) {
	if wsc.channels[channel.Name] == nil {
		wsc.channels[channel.Name] = channel
	}
}

func (wsc *WSClient) HasChannel(channelName string) bool {
	return wsc.channels[channelName] != nil
}

func (wsc *WSClient) LeaveChannel(channel *WSChannel) {
	if wsc.channels[channel.Name] != nil {
		delete(wsc.channels, channel.Name)
	}
}

func (wsc *WSClient) Write(channel *WSChannel, code ResponseCode, traceId string, data []byte) error {
	if channel != nil && !wsc.HasChannel(channel.Name) {
		return errors.New("未在此频道内")
	}
	if msg, err := MallocResponseMessage(); err == nil {
		msg.Code = code
		msg.TraceId = traceId
		msg.Data = data
		json.Marshal(msg)
		err = wsc.Socket.WriteMessage(int(TextMessage), data)
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
