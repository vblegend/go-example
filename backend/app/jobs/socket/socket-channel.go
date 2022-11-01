package socket

import (
	"backend/core/ws"
	"fmt"
)

type JobSocketChannel struct {
	ws.WSChannel
}

func (wd *JobSocketChannel) Name() string {
	return "jobs"
}

func (wd *JobSocketChannel) OnJoin(client *ws.WSClient) {
	fmt.Printf("新连接加入：%s\n", client.Params)
	// 发送所有定时任务的状态至客户端
	// client.SendJsonMessage(wd)
}

// websocket  连接断开
func (wd *JobSocketChannel) OnLeave(client *ws.WSClient) {
	fmt.Printf("连接断开：id:%s\n", client.Params)
}

// websocket  连接断开
func (wd *JobSocketChannel) OnMessagePost(client *ws.WSClient, msg *ws.RequestMessage) {
	fmt.Printf("收到消息：id:%s, content:%s\n", client.ClientId, string(msg.Payload))
	// wd.Channel.BroadcastTextMessage("Hello!")
}

func (wd *JobSocketChannel) OnMessageCall(client *ws.WSClient, msg *ws.RequestMessage) (*ws.ResponseMessage, error) {
	fmt.Printf("收到消息：id:%s, content:%s\n", client.ClientId, string(msg.Payload))
	// wd.Channel.BroadcastTextMessage("Hello!")
	response := msg.Response(ws.Success)
	response.Message = "OKOK"
	return response, nil
}
