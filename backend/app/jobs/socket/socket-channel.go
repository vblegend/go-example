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
}

func (wd *JobSocketChannel) OnLeave(client *ws.WSClient) {
	fmt.Printf("连接断开：id:%s\n", client.Params)
}

func (wd *JobSocketChannel) OnMessagePost(client *ws.WSClient, msg *ws.RequestMessage) {
	fmt.Printf("收到消息：id:%s, content:%s\n", client.ClientId, string(msg.Payload))
}

func (wd *JobSocketChannel) OnMessageCall(client *ws.WSClient, msg *ws.RequestMessage) (*ws.ResponseMessage, error) {
	fmt.Printf("收到消息：id:%s, content:%s\n", client.ClientId, string(msg.Payload))
	response := msg.Success("OKOK")
	return response, nil
}
