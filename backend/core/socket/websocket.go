package socket

import (
	"backend/core/log"
	"context"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketEventListener interface {
	// websocket  连接建立
	WsOpen(client *Client)
	// websocket  连接断开
	WsClose(client *Client)
}

// Manager 所有 websocket 信息
type Manager struct {
	Group                   map[string]map[string]*Client
	groupCount, clientCount uint
	Lock                    sync.Mutex
	Register, UnRegister    chan *Client
	Message                 chan *MessageData
	GroupMessage            chan *GroupMessageData
	BroadCastMessage        chan *BroadCastMessageData
	EventListeners          []WebSocketEventListener
}

// Client 单个 websocket 信息
type Client struct {
	Id, Group  string
	Context    context.Context
	CancelFunc context.CancelFunc
	Socket     *websocket.Conn
	Message    chan []byte
}

// messageData 单个发送数据信息
type MessageData struct {
	Id, Group string
	Context   context.Context
	Message   []byte
}

// groupMessageData 组广播数据信息
type GroupMessageData struct {
	Group   string
	Message []byte
}

// 广播发送数据信息
type BroadCastMessageData struct {
	Message []byte
}

// 读信息，从 websocket 连接直接读取数据
func (c *Client) Read(ctx context.Context) {
	defer func(ctx context.Context) {
		WebsocketManager.UnRegister <- c
		log.Infof("client [%s] disconnect", c.Id)
		if err := c.Socket.Close(); err != nil {
			log.Errorf("client [%s] disconnect err: %s", c.Id, err)
		}
	}(ctx)

	for {
		select {
		case <-ctx.Done():
			// Context结束
			return
		default:
			if ctx.Err() != nil {
				return
			}
			messageType, message, err := c.Socket.ReadMessage()
			if err != nil || messageType == websocket.CloseMessage {
				return
			}
			log.Infof("client [%s] receive message: %s", c.Id, string(message))
			c.Message <- message
		}

	}
}

// 写信息，从 channel 变量 Send 中读取数据写入 websocket 连接
func (c *Client) Write(ctx context.Context) {
	defer func(ctx context.Context) {
		_ = recover()
		log.Infof("client [%s] disconnect", c.Id)
		if err := c.Socket.Close(); err != nil {
			log.Infof("client [%s] disconnect err: %s", c.Id, err)
		}
	}(ctx)

	for {
		select {
		case <-ctx.Done():
			// Context 结束
			return
		case message, ok := <-c.Message:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			// log.Printf("client [%s] write message: %s", c.Id, string(message))
			err := c.Socket.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Errorf("client [%s] writemessage err: %s", c.Id, err)
			}
		}
	}
}

// 启动 websocket 管理器
func (manager *Manager) Start() {
	log.Info("websocket manage start")
	for {
		select {
		// 注册
		case client := <-manager.Register:
			log.Infof("client [%s] connect", client.Id)
			log.Infof("register client [%s] to group [%s]", client.Id, client.Group)

			manager.Lock.Lock()
			for i := 0; i < len(manager.EventListeners); i++ {
				manager.EventListeners[i].WsOpen(client)
			}
			if manager.Group[client.Group] == nil {
				manager.Group[client.Group] = make(map[string]*Client)
				manager.groupCount += 1
			}
			manager.Group[client.Group][client.Id] = client
			manager.clientCount += 1
			manager.Lock.Unlock()

		// 注销
		case client := <-manager.UnRegister:
			log.Infof("unregister client [%s] from group [%s]", client.Id, client.Group)
			manager.Lock.Lock()
			if mGroup, ok := manager.Group[client.Group]; ok {
				if mClient, ok := mGroup[client.Id]; ok {
					close(mClient.Message)
					delete(mGroup, client.Id)
					manager.clientCount -= 1
					if len(mGroup) == 0 {
						//log.Printf("delete empty group [%s]", client.Group)
						delete(manager.Group, client.Group)
						manager.groupCount -= 1
					}
					mClient.CancelFunc()
				}
			}
			for i := 0; i < len(manager.EventListeners); i++ {
				manager.EventListeners[i].WsClose(client)
			}
			manager.Lock.Unlock()

			// 发送广播数据到某个组的 channel 变量 Send 中
			//case data := <-manager.boardCast:
			//	if groupMap, ok := manager.wsGroup[data.GroupId]; ok {
			//		for _, conn := range groupMap {
			//			conn.Send <- data.Data
			//		}
			//	}
		}
	}
}

// 处理单个 client 发送数据
func (manager *Manager) SendService() {
	for {
		select {
		case data := <-manager.Message:
			if groupMap, ok := manager.Group[data.Group]; ok {
				if conn, ok := groupMap[data.Id]; ok {
					conn.Message <- data.Message
				}
			}
		}
	}
}

// 处理 group 广播数据
func (manager *Manager) SendGroupService() {
	for {
		select {
		// 发送广播数据到某个组的 channel 变量 Send 中
		case data := <-manager.GroupMessage:
			if groupMap, ok := manager.Group[data.Group]; ok {
				for _, conn := range groupMap {
					conn.Message <- data.Message
				}
			}
		}
	}
}

// 处理广播数据
func (manager *Manager) SendAllService() {
	for {
		select {
		case data := <-manager.BroadCastMessage:
			for _, v := range manager.Group {
				for _, conn := range v {
					conn.Message <- data.Message
				}
			}
		}
	}
}

// 向指定的 client 发送数据
func (manager *Manager) Send(cxt context.Context, id string, group string, message []byte) {
	data := &MessageData{
		Id:      id,
		Context: cxt,
		Group:   group,
		Message: message,
	}
	manager.Message <- data
}

func (manager *Manager) SendGroupTextMessage(group string, message string) {
	manager.SendGroup(group, []byte(message))
}

// 向指定的 Group 广播
func (manager *Manager) SendGroup(group string, message []byte) {
	data := &GroupMessageData{
		Group:   group,
		Message: message,
	}
	manager.GroupMessage <- data
}

// 广播
func (manager *Manager) SendAll(message []byte) {
	data := &BroadCastMessageData{
		Message: message,
	}
	manager.BroadCastMessage <- data
}

// 注册
func (manager *Manager) RegisterClient(client *Client) {
	manager.Register <- client
}

// 注销
func (manager *Manager) UnRegisterClient(client *Client) {
	manager.UnRegister <- client
}

// 当前组个数
func (manager *Manager) LenGroup() uint {
	return manager.groupCount
}

// 当前连接个数
func (manager *Manager) LenClient() uint {
	return manager.clientCount
}

// 获取 wsManager 管理器信息
func (manager *Manager) Info() map[string]interface{} {
	managerInfo := make(map[string]interface{})
	managerInfo["groupLen"] = manager.LenGroup()
	managerInfo["clientLen"] = manager.LenClient()
	managerInfo["chanRegisterLen"] = len(manager.Register)
	managerInfo["chanUnregisterLen"] = len(manager.UnRegister)
	managerInfo["chanMessageLen"] = len(manager.Message)
	managerInfo["chanGroupMessageLen"] = len(manager.GroupMessage)
	managerInfo["chanBroadCastMessageLen"] = len(manager.BroadCastMessage)
	return managerInfo
}

// 初始化 wsManager 管理器
var WebsocketManager = Manager{
	Group:            make(map[string]map[string]*Client),
	Register:         make(chan *Client, 128),
	UnRegister:       make(chan *Client, 128),
	GroupMessage:     make(chan *GroupMessageData, 128),
	Message:          make(chan *MessageData, 128),
	BroadCastMessage: make(chan *BroadCastMessageData, 128),
	groupCount:       0,
	clientCount:      0,
	EventListeners:   make([]WebSocketEventListener, 0),
}

// gin 处理 websocket handler
func (manager *Manager) WsClient(c *gin.Context) {

	ctx, cancel := context.WithCancel(context.Background())

	upGrader := websocket.Upgrader{
		// cross origin domain
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// 处理 Sec-WebSocket-Protocol Header
		Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
	}

	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Errorf("websocket connect error: %s", c.Param("channel"))
		cancel()
		return
	}

	log.Infof("token:%s", c.Query("token"))

	client := &Client{
		Id:         c.Param("id"),
		Group:      c.Param("channel"),
		Context:    ctx,
		CancelFunc: cancel,
		Socket:     conn,
		Message:    make(chan []byte, 1024),
	}

	manager.RegisterClient(client)
	go client.Read(ctx)
	go client.Write(ctx)
	// time.Sleep(time.Second * 15)
	// 这个方法有bug 回调方法内容不准确
	// pkg.FileMonitoringById(ctx, "temp/logs/job/db-20200820.log", c.Param("id"), c.Param("channel"), SendOne)
}

func SendOne(ctx context.Context, id string, group string, msg string) {
	WebsocketManager.Send(ctx, id, group, []byte("{\"code\":200,\"data\":"+msg+"}"))
	log.Info(WebsocketManager.Info())
}

func (manager *Manager) UnWsClient(c *gin.Context) {
	id := c.Param("id")
	group := c.Param("channel")
	WsLogout(id, group)
	c.Set("result", "ws close success")
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": "ws close success",
		"msg":  "success",
	})
}

func (manager *Manager) AddEventListen(listener WebSocketEventListener) {
	for i := 0; i < len(manager.EventListeners); i++ {
		lst := manager.EventListeners[i]
		if lst == listener {
			return
		}
	}
	manager.EventListeners = append(manager.EventListeners, listener)
}

func (manager *Manager) RemoveEventListen(listener WebSocketEventListener) {
	for i := 0; i < len(manager.EventListeners); i++ {
		lst := manager.EventListeners[i]
		if lst == listener {
			manager.EventListeners = append(manager.EventListeners[:i], manager.EventListeners[i+1:]...)
			return
		}
	}
}

func SendGroup(groupName string, msg []byte) {
	WebsocketManager.SendGroup(groupName, msg)
	log.Info(WebsocketManager.Info())
}

func SendAll(msg []byte) {
	WebsocketManager.SendAll(msg)
	log.Info(WebsocketManager.Info())
}

func WsLogout(id string, group string) {
	WebsocketManager.UnRegisterClient(&Client{Id: id, Group: group})
	log.Info(WebsocketManager.Info())
}