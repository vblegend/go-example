package ws

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

type WSChannel struct {
	Name           string
	lock           sync.Mutex
	clients        map[string]*WSClient
	eventListeners []WSEventListener
	parameters     []string
}

func NewWebSocketChannel(chanName string) *WSChannel {
	wsc := WSChannel{}
	wsc.Name = chanName
	wsc.clients = make(map[string]*WSClient)
	wsc.parameters = make([]string, 0)
	wsc.eventListeners = make([]WSEventListener, 0)
	return &wsc
}

func (ws *WSChannel) AddParameters(params ...string) *WSChannel {
	for _, param := range params {
		ws.parameters = append(ws.parameters, param)
	}
	return ws
}

func (wc *WSChannel) checkParameters(c *gin.Context) error {
	for _, s := range wc.parameters {
		if _, ok := c.GetQuery(s); !ok {
			return fmt.Errorf("necessary request parameter %s not found", s)
		}
	}
	return nil
}

func (wc *WSChannel) register(client *WSClient) {
	has := wc.clients[client.ConnectId]
	if has != nil {
		client.Close()
		return
	}
	wc.clients[client.ConnectId] = client
	client.JoinChannel(wc)
	wc.lock.Lock()
	for i := 0; i < len(wc.eventListeners); i++ {
		wc.eventListeners[i].OnJoin(client)
	}
	wc.lock.Unlock()
}

func (wc *WSChannel) unRegister(client *WSClient) {
	has := wc.clients[client.ConnectId]
	// 区分相同ID的不同客户端连接
	if has != nil && has == client {
		client.LeaveChannel(wc)
		wc.lock.Lock()
		for i := 0; i < len(wc.eventListeners); i++ {
			wc.eventListeners[i].OnLeave(client)
		}
		wc.lock.Unlock()
		delete(wc.clients, client.ConnectId)
	}
	client.Close()
}

func (wc *WSChannel) Length() int {
	return len(wc.clients)
}

func (wc *WSChannel) CanDestroy() bool {
	return len(wc.clients) == 0 && len(wc.eventListeners) == 0
}

func (wc *WSChannel) BroadcastTextMessage(code ResponseCode, traceId string, data string) {
	for _, client := range wc.clients {
		client.Write(wc, code, traceId, []byte(data))
	}
}

func (ws *WSChannel) AddEventListen(listener WSEventListener) *WSChannel {
	ws.lock.Lock()
	defer ws.lock.Unlock()
	for i := 0; i < len(ws.eventListeners); i++ {
		lst := ws.eventListeners[i]
		if lst == listener {
			return ws
		}
	}
	ws.eventListeners = append(ws.eventListeners, listener)
	return ws
}

func (ws *WSChannel) RemoveEventListen(listener WSEventListener) {
	ws.lock.Lock()
	defer ws.lock.Unlock()
	for i := 0; i < len(ws.eventListeners); i++ {
		lst := ws.eventListeners[i]
		if lst == listener {
			ws.eventListeners = append(ws.eventListeners[:i], ws.eventListeners[i+1:]...)
			return
		}
	}
}

type hWSChannel struct {
}
