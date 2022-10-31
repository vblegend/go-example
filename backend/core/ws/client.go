package ws

import (
	"context"
	"encoding/json"
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
}

func NewWSClient() *WSClient {
	wsc := WSClient{}
	return &wsc
}

func (wsc *WSClient) SendMessage(messageType MessageType, data []byte) error {
	return wsc.Socket.WriteMessage(int(messageType), data)
}

func (wsc *WSClient) SendJsonMessage(object interface{}) error {
	data, err := json.Marshal(object)
	if err != nil {
		return err
	}
	return wsc.Socket.WriteMessage(int(TextMessage), data)
}

func (wsc *WSClient) Close() error {
	wsc.Socket.WriteMessage(int(CloseMessage), nil)
	wsc.Cancel()
	return wsc.Socket.Close()
}
