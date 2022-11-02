package ws

import (
	"backend/core/random"
	"context"
	"encoding/json"
	"errors"
	"sync"

	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WSManager struct {
	channels map[string]*WSChannel
	clients  map[string]*WSClient
	chanLock sync.Mutex
}

// 默认的 websocket 管理器
var Default = NewWebSocketManager()

func NewWebSocketManager() *WSManager {
	ws := &WSManager{}
	ws.channels = make(map[string]*WSChannel)
	return ws
}

// 注册一个频道 ， 使用perm控制频道的消息处理授权
func (ws *WSManager) RegisterChannel(channel *WSChannel, perm AuthType) error {
	if channel.Name == "" {
		return errors.New("无效的频道")
	}
	if ws.channels[channel.Name] != nil {
		return errors.New("重复注册频道")
	}
	ws.channels[channel.Name] = channel
	channel.Perm = perm
	return nil
}

func (ws *WSManager) parseParams(c *gin.Context) url.Values {
	params := c.Request.URL.Query()
	if _, ok := c.GetQuery("clientId"); !ok {
		return nil
	}
	return params
}

func (ws *WSManager) newClientId() string {
	var clientId string
	for clientId == "" || ws.clients[clientId] != nil {
		clientId = random.RandomString(8)
	}
	return clientId
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
	if ws.clients[clientId] != nil {
		c.Writer.Header().Set("error", "duplicate client id")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "duplicate client id"})
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

	client := NewWSClient(conn, ctx, cancel, clientId)
	// go channel.register(client)
	go ws.readLoop(client)
}

func (ws *WSManager) readLoop(client *WSClient) {
	defer func() {
		// 下线 退出所有频道
		ws.clientOffline(client)
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
			if msg, err := MallocRequestMessage(); err == nil {
				if err = msg.Unmarshal(message); err == nil {
					ws.datarecv(client, msg)
				}
				FreeRequestMessage(msg)
			}
		}
	}
}

func (ws *WSManager) datarecv(client *WSClient, msg *RequestMessage) {
	defer func() {
		_ = recover()
	}()
	channel := ws.channels[msg.Channel]
	if channel == nil {
		client.Error(msg.TraceId, InvalidChannelName)
		return
	}
	switch msg.Action {
	case JoinChannel:
		{
			p := Params{}
			json.Unmarshal([]byte(msg.Payload), &p)
			err := channel.JoinClient(client, p)
			if err != nil {
				client.Error(msg.TraceId, err)
				return
			}
			client.OK(msg.TraceId, nil, "welcome")
		}
	case LevelChannel:
		{
			err := channel.LeaveClient(client)
			if err != nil {
				client.Error(msg.TraceId, err)
				return
			}
			client.OK(msg.TraceId, nil, "goodbye")
		}
	case TransferPost:
		{
			if channel.Perm != Auth_Anonymous && (channel.Perm&Auth_PostNeedJoin) == Auth_PostNeedJoin && !client.HasChannel(msg.Channel) {
				client.Error(msg.TraceId, errors.New("当前动作不被允许"))
				return
			}
			go channel.MessagePost(client, msg)
		}
	case TransferSend:
		{
			if channel.Perm != Auth_Anonymous && (channel.Perm&Auth_SendNeedJoin) == Auth_SendNeedJoin && !client.HasChannel(msg.Channel) {
				client.Error(msg.TraceId, errors.New("当前动作不被允许"))
				return
			}
			go channel.MessageCall(client, msg)
		}
	default:
		{
			client.Error(msg.TraceId, errors.New("无效的权限"))
		}
	}
	// join channel  增加 参数
	// 数据传输 send  post 增加参数
}

func (ws *WSManager) clientOffline(client *WSClient) {
	for _, channel := range ws.channels {
		channel.LeaveClient(client)
	}
	delete(ws.clients, client.ClientId)
}
