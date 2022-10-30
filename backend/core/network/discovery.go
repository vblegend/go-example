package network

import (
	"errors"
	"fmt"
	"net"

	"golang.org/x/net/ipv4"
)

type RecvHandler func(net.Addr, int, []byte)
type IDiscoveryService interface {
	Listen(localAddr string) error
	SetLoopback(state bool) error
	GetLoopback() (bool, error)
	SetRecvHandler(fn RecvHandler)
	Send(buffer []byte) (int, error)
	Shutdown() error
}

//
//UDP发现服务
//
type DiscoveryService struct {
	udpConn          *net.UDPConn
	packetConn       *ipv4.PacketConn
	multicastAddress net.UDPAddr
	listenAddress    net.UDPAddr
	fn               RecvHandler
	shutdown         chan bool
}

//
// 用一个组播地址创建一个udp发现服务 如 239.255.255.250
//
func NewDiscoveryServer(multicastIPAddr string, recvfn RecvHandler) (IDiscoveryService, error) {
	server := DiscoveryService{
		fn:       recvfn,
		shutdown: make(chan bool, 10),
	}
	gip := net.ParseIP(multicastIPAddr)
	if gip == nil {
		return nil, fmt.Errorf("'%s' Invalid multicast IP address", multicastIPAddr)
	}
	server.multicastAddress = net.UDPAddr{IP: gip, Port: 0}
	return &server, nil
}

//
// 监听本地端口 0.0.0.0:12345 端口号为组播通讯端口号
//
//
func (ds *DiscoveryService) Listen(localAddr string) error {
	addr, err := net.ResolveUDPAddr("udp", localAddr)
	if err != nil {
		return err
	}
	ds.listenAddress = *addr
	ds.multicastAddress.Port = addr.Port
	conn, err := net.ListenUDP("udp4", &ds.listenAddress)
	if err != nil {
		return err
	}
	ds.udpConn = conn
	err = ds.joinGroup()
	if err != nil {
		return nil
	}
	go ds.recvdata()
	return nil
}

func (ds *DiscoveryService) recvdata() {
	buf := make([]byte, 65536)
	for {
		select {
		case <-ds.shutdown:
			return
		default:
			n, from, err := ds.udpConn.ReadFrom(buf)
			if err != nil {
				continue
			}
			if ds.fn != nil {
				ds.fn(from, n, buf[:n])
			}
		}
	}
}

// 拟定一个回调 用以接收数据 此函数为协程函数,调用shutdown时退出
// go server.RecvDataOn(fn)
func (ds *DiscoveryService) SetRecvHandler(fn RecvHandler) {
	ds.fn = fn
}

// 加入广播组
func (ds *DiscoveryService) joinGroup() error {
	pc := ipv4.NewPacketConn(ds.udpConn)
	ifaces, err := net.Interfaces()
	if err != nil {
		return err
	}
	errCount := 0
	for _, iface := range ifaces {
		if err := pc.JoinGroup(&iface, &net.UDPAddr{IP: ds.multicastAddress.IP}); err != nil {
			errCount++
		}
	}
	if len(ifaces) == errCount {
		return errors.New("failed to join multicast group on all interfaces")
	}
	ds.packetConn = pc
	return nil
}

// 设置广播消息是否环回
func (ds *DiscoveryService) SetLoopback(state bool) error {
	if ds.packetConn == nil {
		return fmt.Errorf("connection not initialized")
	}
	return ds.packetConn.SetMulticastLoopback(state)
}

// 获取广播消息环回状态
func (ds *DiscoveryService) GetLoopback() (bool, error) {
	if ds.packetConn == nil {
		return false, fmt.Errorf("connection not initialized")
	}
	return ds.packetConn.MulticastLoopback()
}

// 发送广播消息，此消息是发送给广播组内所有成员的
func (ds *DiscoveryService) Send(buffer []byte) (int, error) {
	return ds.udpConn.WriteTo(buffer, &ds.multicastAddress)
}

// 关闭并退出UDP
func (ds *DiscoveryService) Shutdown() error {
	ds.shutdown <- true
	if ds.packetConn != nil {
		ds.packetConn.Close()
		ds.packetConn = nil
	}
	if ds.udpConn != nil {
		ds.udpConn.Close()
		ds.udpConn = nil
	}
	return nil
}
