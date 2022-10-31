package ws

import (
	"fmt"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func Test_WebSocket(t *testing.T) {
	var r = gin.New()
	ws := NewWebSocketManager()
	ws.RegisterRouter(r)
	demo := wsDemo{}
	channel := ws.GetChannel("games")
	channel.AddParameters("gameId")
	channel.AddEventListen(&demo)
	go func() {
		for i := 0; i < 10000; i++ {
			for cancel := 0; cancel < 100; cancel++ {
				channel.BroadcastTextMessage(fmt.Sprintf("这是第%d条消息", cancel))
			}
			time.Sleep(time.Second)
		}
	}()
	fmt.Println("Start")
	r.Run(":10086")
	fmt.Println("hello")
}

type wsDemo struct {
}

func (wd *wsDemo) OnJoin(client *WSClient) {
	fmt.Printf("新连接加入：%s\n", client.Params)
}

// websocket  连接断开
func (wd *wsDemo) OnLeave(client *WSClient) {
	fmt.Printf("连接断开：id:%s\n", client.Params)
}

// websocket  连接断开
func (wd *wsDemo) OnMessage(client *WSClient, msgType MessageType, message []byte) {

	fmt.Printf("收到消息：id:%s, content:%s\n", client.ConnectId, string(message))
	client.SendJsonMessage(wd)
}