package ws

import (
	"context"
	"encoding/json"

	"github.com/gorilla/websocket"
)

type wsClient struct {
	Socket  *websocket.Conn
	Context context.Context
	Cancel  context.CancelFunc
	// 当前连接唯一ID
	clientID string
}

func newWSClient(ctx context.Context, conn *websocket.Conn, cancel context.CancelFunc, clientId string) *wsClient {
	wsc := wsClient{
		Socket:   conn,
		Context:  ctx,
		Cancel:   cancel,
		clientID: clientId,
	}
	return &wsc
}

func (wsc *wsClient) ClientID() string {
	return wsc.clientID
}

func (wsc *wsClient) Success(traceID string, message string, data []byte) error {
	if msg, err := MallocResponseMessage(); err == nil {
		msg.Code = Success
		msg.TraceID = traceID
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
func (wsc *wsClient) Write(msg *ResponseMessage) error {
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

func (wsc *wsClient) OK(traceID string, data []byte, message string) error {
	if msg, err := MallocResponseMessage(); err == nil {
		msg.Code = Success
		msg.TraceID = traceID
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

func (wsc *wsClient) Error(traceID string, data error) error {
	if msg, err := MallocResponseMessage(); err == nil {
		msg.Code = Failure
		msg.TraceID = traceID
		msg.Message = data.Error()
		if bytes, err := json.Marshal(msg); err == nil {
			err = wsc.Socket.WriteMessage(int(TextMessage), bytes)
		}
		FreeResponseMessage(msg)
		return err
	}
	return nil
}

func (wsc *wsClient) Close() error {
	wsc.Socket.WriteMessage(int(CloseMessage), nil)
	wsc.Cancel()
	return wsc.Socket.Close()
}
