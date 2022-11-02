package ws

import (
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

type wsManager struct {
	channels map[string]*WSChannel
	clients  map[string]*wsClient
	chanLock sync.Mutex
}

// Default 默认的 websocket 管理器
var Default IWSManager = NewWebSocketManager()

// NewWebSocketManager 创建一个websocket 管理器
func NewWebSocketManager() IWSManager {
	ws := &wsManager{}
	ws.channels = make(map[string]*WSChannel)
	return ws
}

// RegisterChannel 注册一个频道 ， 使用perm控制频道的消息处理授权
func (ws *wsManager) RegisterChannel(name string, handler IWSMessageHandler, perm AuthType) error {
	if name == "" {
		return errors.New("无效的频道")
	}
	if ws.channels[name] != nil {
		return errors.New("重复注册频道")
	}
	ws.channels[name] = newWSChannel(name, perm, handler)
	return nil
}

func (ws *wsManager) parseParams(c *gin.Context) url.Values {
	params := c.Request.URL.Query()
	if _, ok := c.GetQuery("clientId"); !ok {
		return nil
	}
	return params
}

// AcceptHandler websocket 应答处理器
func (ws *wsManager) AcceptHandler(c *gin.Context) {
	params := ws.parseParams(c)
	if params == nil {
		c.Writer.Header().Set("error", "missing necessary request parameters")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "missing necessary request parameters"})
		return
	}
	clientID := params.Get("clientId")
	if ws.clients[clientID] != nil {
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

	client := newWSClient(ctx, conn, cancel, clientID)
	// go channel.register(client)
	go ws.readLoop(client)
}

func (ws *wsManager) readLoop(client *wsClient) {
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

func (ws *wsManager) datarecv(client *wsClient, msg *RequestMessage) {
	defer func() {
		_ = recover()
	}()
	channel := ws.channels[msg.Channel]
	if channel == nil {
		client.Error(msg.TraceID, errorInvalidChannelName)
		return
	}
	switch msg.Action {
	case JoinChannel:
		{
			p := Params{}
			json.Unmarshal([]byte(msg.Payload), &p)
			err := channel.joinClient(client, p)
			if err != nil {
				client.Error(msg.TraceID, err)
				return
			}
			client.OK(msg.TraceID, nil, "welcome")
		}
	case LevelChannel:
		{
			err := channel.leaveClient(client)
			if err != nil {
				client.Error(msg.TraceID, err)
				return
			}
			client.OK(msg.TraceID, nil, "goodbye")
		}
	case TransferPost:
		{
			if channel.Perm != Auth_Anonymous && (channel.Perm&Auth_PostNeedJoin) == Auth_PostNeedJoin && channel.GetClient(client.ClientID()) == nil {
				client.Error(msg.TraceID, errors.New("当前动作不被允许"))
				return
			}
			go channel.messagePost(client, msg)
		}
	case TransferSend:
		{
			if channel.Perm != Auth_Anonymous && (channel.Perm&Auth_SendNeedJoin) == Auth_SendNeedJoin && channel.GetClient(client.ClientID()) == nil {
				client.Error(msg.TraceID, errors.New("当前动作不被允许"))
				return
			}
			go channel.messageCall(client, msg)
		}
	default:
		{
			client.Error(msg.TraceID, errors.New("无效的权限"))
		}
	}
	// join channel  增加 参数
	// 数据传输 send  post 增加参数
}

func (ws *wsManager) clientOffline(client *wsClient) {
	for _, channel := range ws.channels {
		channel.leaveClient(client)
	}
	delete(ws.clients, client.clientID)
}
