package ws

type WSChannel struct {
	// Name           string
	// channelLock    sync.Mutex
	clients map[string]*WSClient
	// eventListeners []WSEventListener
	// parameters     []string
}

// func (ws *WSChannel) AddParameters(params ...string) *WSChannel {
// 	for _, param := range params {
// 		ws.parameters = append(ws.parameters, param)
// 	}
// 	return ws
// }

// func (wc *WSChannel) checkParameters(c *gin.Context) error {
// 	for _, s := range wc.parameters {
// 		if _, ok := c.GetQuery(s); !ok {
// 			return fmt.Errorf("necessary request parameter %s not found", s)
// 		}
// 	}
// 	return nil
// }

func (wc *WSChannel) Length() int {
	return len(wc.clients)
}

// func (wc *WSChannel) BroadcastTextMessage(code ResponseCode, traceId string, data string) {
// 	for _, client := range wc.clients {
// 		client.Write(wc, code, traceId, []byte(data))
// 	}
// }

type hWSChannel struct {
}
