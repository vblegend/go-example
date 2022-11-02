package socket

import (
	"backend/core/ws"
	"fmt"
)

type JobSocketChannel struct {
	*ws.WSChannel
}

func (wd *JobSocketChannel) OnMethodMap() {
	// wd.PostMap("post", wd.OnMessageCall)
	// wd.SendMap("send", wd.OnMessageCall)
}

func (wd *JobSocketChannel) Name() string {
	return "jobs"
}

func (wd *JobSocketChannel) OnJoin(client *ws.WSClient, params ws.Params) error {
	pwd := "xxdsa"
	params.Parse("pwd", &pwd)
	fmt.Println(pwd)
	fmt.Printf("新连接加入：%s\n", client.ClientId)
	return nil
}

func (wd *JobSocketChannel) OnLeave(client *ws.WSClient) {
	fmt.Printf("连接断开：id:%s\n", client.ClientId)
}

func (wd *JobSocketChannel) OnMessagePost(client *ws.WSClient, msg *ws.RequestMessage) {
	if msg.Method == "add" {
		fmt.Printf("收到消息Add：id:%s, content:%s\n", client.ClientId, msg.Payload)
	} else {
		fmt.Printf("收到无效消息：id:%s, content:%s\n", client.ClientId, msg.Payload)
	}
}

func (wd *JobSocketChannel) OnMessageCall(client *ws.WSClient, msg *ws.RequestMessage) (*ws.ResponseMessage, error) {
	fmt.Printf("收到消息：id:%s, content:%s\n", client.ClientId, msg.Payload)
	response := msg.Success("OKOK")
	return response, nil
}
