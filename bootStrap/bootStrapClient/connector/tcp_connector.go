package connector

import "github.com/gox-studio/gona/executor"

type SocketType int

const (
	NormalSocket SocketType = 0
	WebSocket    SocketType = 1
)

type IConnector interface {
	Connect(ip string, port int, success IConnectSuccess, fail IConnectFail)
}

func NewTcpConnector(socketType SocketType, routinePoolID int64) IConnector {
	switch socketType {
	case WebSocket:
		instance := new(websocketConnector)
		instance.routinePoolID = routinePoolID
		return instance
	}
	instance := new(tcpConnector)
	instance.routinePoolID = routinePoolID
	return instance
}

type tcpConnector struct {
	routinePoolID int64
}

func (connector *tcpConnector) Connect(ip string, port int, success IConnectSuccess, fail IConnectFail) {
	executor.FireEvent(newConnectEvent(connector.routinePoolID, ip, port, -1, success, fail))
}

type websocketConnector struct {
	routinePoolID int64
}

func (connector *websocketConnector) Connect(ip string, port int, success IConnectSuccess, fail IConnectFail) {
	executor.FireEvent(newWebsocketConnectEvent(connector.routinePoolID, ip, port, -1, success, fail))
}
