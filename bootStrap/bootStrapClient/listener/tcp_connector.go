package listener

import (
	"net"

	"github.com/gox-studio/gona/bootStrap/bootStrapClient/connector"
	"github.com/gox-studio/gona/logger"
)

type IConnReceiver interface {
	ReceiveTCP(connWrapper *TCPConnWrapper)
}

type tcpConnector struct {
	mConnector    connector.IConnector
	mConnReceiver IConnReceiver
}

func newConnector(mConnector connector.IConnector, mConnReceiver IConnReceiver) IConnector {
	instance := new(tcpConnector)
	instance.mConnector = mConnector
	instance.mConnReceiver = mConnReceiver
	return instance
}

func (wapper *tcpConnector) Connect(ip string, port int, params map[string]interface{}) {
	wapper.mConnector.Connect(ip, port, newConnectSuccess(params, wapper.mConnReceiver), newConnectFail(params))
}

type connectSuccess struct {
	params        map[string]interface{}
	mConnReceiver IConnReceiver
}

func newConnectSuccess(params map[string]interface{}, mConnReceiver IConnReceiver) connector.IConnectSuccess {
	instance := new(connectSuccess)
	instance.params = params
	instance.mConnReceiver = mConnReceiver
	return instance
}

func (success *connectSuccess) Handle(conn net.Conn) {
	success.mConnReceiver.ReceiveTCP(newTCPConnWrapper(conn, success.params))
}

type connectFail struct {
	params map[string]interface{}
}

func newConnectFail(params map[string]interface{}) connector.IConnectFail {
	instance := new(connectFail)
	instance.params = params
	return instance
}

func (fail *connectFail) Handle(err error) {
	logger.Error("ConnectFail err ", err)
}
