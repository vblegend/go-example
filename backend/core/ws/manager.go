package ws

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WSManager struct {
	channels map[string]*WSChannel
}

// 默认的 websocket 管理器
var Default = NewWebSocketManager()

func NewWebSocketManager() *WSManager {
	ws := &WSManager{}
	ws.channels = make(map[string]*WSChannel)
	return ws
}

func (ws *WSManager) RegisterRouter(r *gin.Engine) {
	r.GET("/ws", ws.AcceptHandler)
}

/*
 * 获取一个信道，如果信道不存在则创建
 */
func (ws *WSManager) GetChannel(channel string) *WSChannel {
	if ws.channels[channel] == nil {
		ws.channels[channel] = NewWebSocketChannel(channel)
	}
	return ws.channels[channel]
}

func (ws *WSManager) parseParams(c *gin.Context) url.Values {
	params := c.Request.URL.Query()
	if _, ok := c.GetQuery("clientId"); !ok {
		return nil
	}
	if _, ok := c.GetQuery("channel"); !ok {
		return nil
	}
	return params
}

// websocket 应答处理器
func (ws *WSManager) AcceptHandler(c *gin.Context) {
	params := ws.parseParams(c)
	if params == nil {
		c.Writer.Header().Set("error", "missing necessary request parameters")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "missing necessary request parameters"})
		return
	}
	clientId := params.Get("clientId")
	channelName := params.Get("channel")
	channel := ws.GetChannel(channelName)
	err := channel.checkParameters(c)
	if err != nil {
		c.Writer.Header().Set("error", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	upGrader := websocket.Upgrader{
		CheckOrigin:  func(r *http.Request) bool { return true },
		Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
	}
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Writer.Header().Set("error", err.Error())
		fmt.Printf("websocket connect error: %s", c.Param("channel"))
		cancel()
		return
	}
	client := &WSClient{conn, ctx, cancel, clientId, channelName, params}
	go channel.register(client)
	go ws.readLoop(channel, client)
}

func (ws *WSManager) readLoop(channel *WSChannel, client *WSClient) {
	defer func() {
		channel.unRegister(client)
		if channel.CanDestroy() {
			delete(ws.channels, channel.Name)
		}
	}()
	for {
		select {
		case <-client.Context.Done():
			return
		default:
			if client.Context.Err() != nil {
				return
			}
			messageType, message, err := client.Socket.ReadMessage()
			if err != nil || messageType == websocket.CloseMessage {
				return
			}
			channel.lock.Lock()
			for i := 0; i < len(channel.eventListeners); i++ {
				channel.eventListeners[i].OnMessage(client, MessageType(messageType), message)
			}
			channel.lock.Unlock()
		}
	}
}
