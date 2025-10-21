package listener

import (
	"net"

	"github.com/goodluck0107/gona/boot/bootc/connector"
)

const (
	defaultRetryTimes int = 3 // 连接失败时的默认重试次数
)

type IConnector interface {
	Connect(connType connector.SocketType, ip string, port int, params map[string]interface{}, failFunc func(error, map[string]interface{}))
}

// iConnReceiver 连接接收
type iConnReceiver interface {
	ReceiveConn(connWrapper *ConnWrapper)
}

type tcpConnector struct {
	mConnReceiver iConnReceiver
}

func newConnector(mConnReceiver iConnReceiver) IConnector {
	instance := new(tcpConnector)
	instance.mConnReceiver = mConnReceiver
	return instance
}

func (wapper *tcpConnector) Connect(socketType connector.SocketType, ip string, port int, params map[string]interface{}, failFunc func(error, map[string]interface{})) {
	connector.Connect(socketType, ip, port, defaultRetryTimes, newConnectSuccess(params, wapper.mConnReceiver), newConnectFail(params, failFunc))
}

type connectSuccess struct {
	params        map[string]interface{}
	mConnReceiver iConnReceiver
}

func newConnectSuccess(params map[string]interface{}, mConnReceiver iConnReceiver) connector.IConnectSuccess {
	instance := new(connectSuccess)
	instance.params = params
	instance.mConnReceiver = mConnReceiver
	return instance
}

func (success *connectSuccess) Handle(conn net.Conn) {
	success.mConnReceiver.ReceiveConn(newConnWrapper(conn, success.params))
}

type connectFail struct {
	params   map[string]interface{}
	failFunc func(error, map[string]interface{})
}

func newConnectFail(params map[string]interface{}, failFunc func(error, map[string]interface{})) connector.IConnectFail {
	instance := new(connectFail)
	instance.params = params
	instance.failFunc = failFunc
	return instance
}

func (fail *connectFail) Handle(err error) {
	fail.failFunc(err, fail.params)
}
