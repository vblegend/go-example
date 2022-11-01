package ws

type WSChannel struct {
	clients map[string]*WSClient
}

func (wc *WSChannel) Length() int {
	return len(wc.clients)
}
