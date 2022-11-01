package ws

import (
	"backend/core/random"
	"context"
	"errors"
	"sync"

	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WSManager struct {
	channels    map[string]IWSChannel
	clients     map[string]*WSClient
	chanLock    sync.Mutex
	permissions map[string]AuthType
}

// 默认的 websocket 管理器
var Default = NewWebSocketManager()

func NewWebSocketManager() *WSManager {
	ws := &WSManager{}
	ws.channels = make(map[string]IWSChannel)
	ws.permissions = make(map[string]AuthType)
	return ws
}

func (ws *WSManager) RegisterChannel(channel IWSChannel, perm AuthType) error {
	if channel.Name() == "" {
		return errors.New("无效的频道")
	}
	if ws.channels[channel.Name()] != nil {
		return errors.New("重复注册频道")
	}
	ws.channels[channel.Name()] = channel
	ws.permissions[channel.Name()] = perm
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
	prem := ws.permissions[msg.Channel]
	switch msg.Action {
	case JoinChannel:
		{
			if client.HasChannel(msg.Channel) {
				client.Error(msg.TraceId, ErrorCannotJoinChannelRepeated)
				return
			}
			err := ws.joinChannel(channel, client)
			if err != nil {
				client.Error(msg.TraceId, err)
				return
			}
			client.OK(msg.TraceId, nil, "welcome")
		}
	case LevelChannel:
		{
			if !client.HasChannel(msg.Channel) {
				client.Error(msg.TraceId, NotInChannel)
			}
			err := ws.leaveChannel(channel, client)
			if err != nil {
				client.Error(msg.TraceId, err)
				return
			}
			client.OK(msg.TraceId, nil, "goodbye")
		}
	case TransferPost:
		{
			if prem != Auth_Anonymous && (prem&Auth_PostNeedJoin) == Auth_PostNeedJoin && !client.HasChannel(msg.Channel) {
				client.Error(msg.TraceId, errors.New("当前动作不被允许"))
				return
			}
			channel.OnMessagePost(client, msg)
		}
	case TransferSend:
		{
			if prem != Auth_Anonymous && (prem&Auth_SendNeedJoin) == Auth_SendNeedJoin && !client.HasChannel(msg.Channel) {
				client.Error(msg.TraceId, errors.New("当前动作不被允许"))
				return
			}
			res, err := channel.OnMessageCall(client, msg)
			if err != nil {
				client.Error(msg.TraceId, err)
			}
			client.Write2(res)
		}
	default:
		{
			client.Write(Failure, msg.TraceId, []byte("无效的权限"))
		}
	}
	// join channel  增加 参数
	// 数据传输 send  post 增加参数
}

func (ws *WSManager) joinChannel(channel IWSChannel, client *WSClient) error {
	ws.chanLock.Lock()
	defer ws.chanLock.Unlock()
	if client.HasChannel(channel.Name()) {
		return ErrorCannotJoinChannelRepeated
	}
	client.JoinChannel(channel)
	channel.OnJoin(client)
	return nil
}

func (ws *WSManager) leaveChannel(channel IWSChannel, client *WSClient) error {
	ws.chanLock.Lock()
	defer ws.chanLock.Unlock()
	if !client.HasChannel(channel.Name()) {
		return NotInChannel
	}
	client.LeaveChannel(channel)
	channel.OnLeave(client)
	return nil
}

func (ws *WSManager) clientOffline(client *WSClient) {
	for _, v := range client.channels {
		ws.leaveChannel(v, client)
	}
	delete(ws.clients, client.ClientId)
}
