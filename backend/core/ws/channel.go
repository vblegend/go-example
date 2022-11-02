package ws

import (
	"fmt"
	"reflect"
	"sync"
)

type WSChannel struct {
	Name    string
	Clients map[string]*WSClient
	Handler IWSChannel
	lock    sync.Mutex
	//
	Perm AuthType
}

func (wc *WSChannel) JoinClient(client *WSClient, params Params) error {
	wc.lock.Lock()
	defer wc.lock.Unlock()
	if wc.GetClient(client.ClientId) != nil {
		return ErrorCannotJoinChannelRepeated
	}
	err := wc.Handler.OnJoin(client, params)
	if err == nil {
		wc.Clients[client.ClientId] = client
	}
	return err
}

func (wc *WSChannel) LeaveClient(client *WSClient) error {
	wc.lock.Lock()
	defer wc.lock.Unlock()
	if wc.GetClient(client.ClientId) == nil {
		return NotInChannel
	}
	delete(wc.Clients, client.ClientId)
	wc.Handler.OnLeave(client)
	return nil
}

func (wc *WSChannel) Broadcast(msg *ResponseMessage) {
	wc.lock.Lock()
	defer wc.lock.Unlock()
	for _, client := range wc.Clients {
		client.Write(msg)
	}
}

func (wc *WSChannel) KickedOut(client *WSClient) {
	go func() {
		err := wc.LeaveClient(client)
		if err == nil {
			// 告诉客户端被踢出去了
			fmt.Printf("客户端%s被踢出去咯。", client.ClientId)
		}
	}()
}

func (wc *WSChannel) GetClient(clientId string) *WSClient {
	return wc.Clients[clientId]
}

func (wc *WSChannel) Length() int {
	return len(wc.Clients)
}

func (wc *WSChannel) MessagePost(client *WSClient, msg *RequestMessage) {
	wc.Handler.OnMessagePost(client, msg)
}

func (wc *WSChannel) MessageCall(client *WSClient, msg *RequestMessage) {
	res, err := wc.Handler.OnMessageCall(client, msg)
	if err != nil {
		client.Error(msg.TraceId, err)
	} else {
		client.Write(res)
	}
}

func (wc *WSChannel) PostMap(method string, lfunc interface{}) error {
	v := reflect.ValueOf(lfunc)
	t := v.Type()
	p1 := t.In(0)
	pe1 := p1.Elem()
	fmt.Println(pe1.Name())
	return nil
}

func (wc *WSChannel) SendMap(method string, lfunc interface{}) error {
	return nil
}
