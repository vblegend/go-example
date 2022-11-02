package socket

import (
	"backend/core/ws"
	"fmt"
)

type JobSocketChannel struct {
}

func (wd *JobSocketChannel) OnJoin(channel ws.IWSChannel, client ws.IWSClient, params ws.Params) error {
	pwd := "xxdsa"
	params.Parse("pwd", &pwd)
	fmt.Println(pwd)
	fmt.Printf("新连接加入：%s\n", client.ClientID())
	return nil
}

func (wd *JobSocketChannel) OnLeave(channel ws.IWSChannel, client ws.IWSClient) {
	fmt.Printf("连接断开：id:%s\n", client.ClientID())
}

func (wd *JobSocketChannel) OnMessagePost(channel ws.IWSChannel, client ws.IWSClient, msg *ws.RequestMessage) {
	if msg.Method == "add" {
		fmt.Printf("收到消息Add：id:%s, content:%s\n", client.ClientID(), msg.Payload)
	} else {
		fmt.Printf("收到无效消息：id:%s, content:%s\n", client.ClientID(), msg.Payload)
	}

	client.Write(msg.Success("LLLLSX"))

}

func (wd *JobSocketChannel) OnMessageCall(channel ws.IWSChannel, client ws.IWSClient, msg *ws.RequestMessage) (*ws.ResponseMessage, error) {
	fmt.Printf("收到消息：id:%s, content:%s\n", client.ClientID(), msg.Payload)
	response := msg.Success("OKOK")
	return response, nil
}
