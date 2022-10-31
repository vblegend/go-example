package socket

import (
	"backend/core/ws"
	"fmt"
)

type JobSocketService struct {
	Channel *ws.WSChannel
}

func (wd *JobSocketService) OnJoin(client *ws.WSClient) {
	fmt.Printf("新连接加入：%s\n", client.Params)
	// 发送所有定时任务的状态至客户端
	client.SendJsonMessage(wd)
}

// websocket  连接断开
func (wd *JobSocketService) OnLeave(client *ws.WSClient) {
	fmt.Printf("连接断开：id:%s\n", client.Params)
}

// websocket  连接断开
func (wd *JobSocketService) OnMessage(client *ws.WSClient, msgType ws.MessageType, message []byte) {
	fmt.Printf("收到消息：id:%s, content:%s\n", client.ConnectId, string(message))
	wd.Channel.BroadcastTextMessage("Hello!")
}
