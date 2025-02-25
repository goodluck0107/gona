package connector

func Connect(socketType SocketType, ip string, port int, retryTimes int, success IConnectSuccess, fail IConnectFail) {
	switch socketType {
	case WebSocket:
		eventPool.FireEvent(newWebsocketConnectEvent(ip, port, retryTimes, success, fail))
		return
	}
	eventPool.FireEvent(newConnectEvent(ip, port, retryTimes, success, fail))
}

type SocketType int

const (
	NormalSocket SocketType = 0
	WebSocket    SocketType = 1
)
