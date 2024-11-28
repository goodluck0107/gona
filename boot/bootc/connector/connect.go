package connector

import "net"

func Connect(socketType SocketType, ip string, port int, retryTimes int, success IConnectSuccess, fail IConnectFail) {
	switch socketType {
	case WebSocket:
		routinePool.FireEvent(newWebsocketConnectEvent(ip, port, retryTimes, success, fail))
		return
	}
	routinePool.FireEvent(newConnectEvent(ip, port, retryTimes, success, fail))
}

type IConnectSuccess interface {
	Handle(conn net.Conn)
}

type IConnectFail interface {
	Handle(err error)
}

type SocketType int

const (
	NormalSocket SocketType = 0
	WebSocket    SocketType = 1
)
