package ws

import (
	"fmt"
	"sync"
)

// WSChannel 用户对象，定义了用户的基础信息
type WSChannel struct {
	name    string
	Clients map[string]*wsClient
	Handler IWSMessageHandler
	lock    sync.Mutex
	//
	Perm AuthType
}

// NewWSChannel 创建一个websocket频道
// name 频道名
// perm 消息权限
// handler 消息处理器
func newWSChannel(name string, perm AuthType, handler IWSMessageHandler) *WSChannel {
	channel := new(WSChannel)
	channel.name = name
	channel.Perm = perm
	channel.Handler = handler
	return channel
}

// Name 获取频道名
func (wc *WSChannel) Name() string {
	return wc.name
}

// joinClient 加入客户端
// client 客户端
// params 参数
func (wc *WSChannel) joinClient(client *wsClient, params Params) error {
	wc.lock.Lock()
	defer wc.lock.Unlock()
	if wc.GetClient(client.clientID) != nil {
		return errorCannotJoinChannelRepeated
	}
	err := wc.Handler.OnJoin(wc, client, params)
	if err == nil {
		wc.Clients[client.clientID] = client
	}
	return err
}

// g
func (wc *WSChannel) leaveClient(client IWSClient) error {
	wc.lock.Lock()
	defer wc.lock.Unlock()
	if wc.GetClient(client.ClientID()) == nil {
		return errorNotInChannel
	}
	delete(wc.Clients, client.ClientID())
	wc.Handler.OnLeave(wc, client)
	return nil
}

// Broadcast 在频道内广播一条消息
// msg 消息
func (wc *WSChannel) Broadcast(msg *ResponseMessage) {
	wc.lock.Lock()
	defer wc.lock.Unlock()
	for _, client := range wc.Clients {
		client.Write(msg)
	}
}

// KickedOut 把客户端踢出频道
// client
func (wc *WSChannel) KickedOut(client IWSClient) {
	go func() {
		err := wc.leaveClient(client)
		if err == nil {
			// 告诉客户端被踢出去了
			fmt.Printf("客户端%s被踢出去咯。", client.ClientID())
		}
	}()
}

// GetClient 获取指定ID 的客户端对象
// clientID 客户端ID
func (wc *WSChannel) GetClient(clientID string) IWSClient {
	return wc.Clients[clientID]
}

// Length 返回频道内客户端连接数
func (wc *WSChannel) Length() int {
	return len(wc.Clients)
}

func (wc *WSChannel) messagePost(client *wsClient, msg *RequestMessage) {
	wc.Handler.OnMessagePost(wc, client, msg)
}

func (wc *WSChannel) messageCall(client *wsClient, msg *RequestMessage) {
	res, err := wc.Handler.OnMessageCall(wc, client, msg)
	if err != nil {
		client.Error(msg.TraceID, err)
	} else {
		client.Write(res)
	}
}
