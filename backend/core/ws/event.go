package ws

type WSEventListener interface {
	// websocket  连接建立
	OnJoin(client *WSClient)
	// websocket  连接断开
	OnLeave(client *WSClient)

	// websocket  连接断开
	OnMessage(client *WSClient, msgType MessageType, message []byte)
}
