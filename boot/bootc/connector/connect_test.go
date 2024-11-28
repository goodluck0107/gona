package connector

import (
	"fmt"
	"net"
	"testing"
	"time"

	"gitee.com/andyxt/gona/boot"
)

func TestConnect(t *testing.T) {
	params := make(map[string]interface{})
	params["key"] = "clientValue"
	params[boot.KeyPacketBytesCount] = 2
	params[boot.KeyChannelReadLimit] = 10240
	params[boot.KeyReadTimeOut] = -1
	params[boot.KeyWriteTimeOut] = -1
	Connect(NormalSocket, "127.0.0.1", 10086, 3, newConnectSuccess(params), newConnectFail(params, func(err error, params map[string]interface{}) {
		fmt.Println(params, "connectFail", err)
	}))
	for {
		time.Sleep(1 * time.Second)
	}
}

type connectSuccess struct {
	params map[string]interface{}
}

func newConnectSuccess(params map[string]interface{}) IConnectSuccess {
	instance := new(connectSuccess)
	instance.params = params
	return instance
}

func (success *connectSuccess) Handle(conn net.Conn) {
	fmt.Println(success.params, "connectSuccess")
}

type connectFail struct {
	params   map[string]interface{}
	failFunc func(error, map[string]interface{})
}

func newConnectFail(params map[string]interface{}, failFunc func(error, map[string]interface{})) IConnectFail {
	instance := new(connectFail)
	instance.params = params
	instance.failFunc = failFunc
	return instance
}

func (fail *connectFail) Handle(err error) {
	fail.failFunc(err, fail.params)
}
