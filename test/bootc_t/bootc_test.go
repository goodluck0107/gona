package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gona/boot/bootc"
	"gitee.com/andyxt/gona/boot/bootc/connector"
	"gitee.com/andyxt/gona/boot/bootc/listener"
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/logger"
)

func TestClient(t *testing.T) {
	testClient()
}

var connConnector listener.IConnector

func testClient() {
	connConnector :=
		bootc.Serv(bootc.WithInitializer(NewTestChannelInitializer()), bootc.WithLogger(logger.GetLogger()))
	fmt.Println("Connect")
	IP := "127.0.0.1" // 连接IP
	Port := 20000     // 连接端口
	params := make(map[string]interface{})
	params["key"] = "clientValue"
	params[boot.KeyPacketBytesCount] = 2
	params[boot.KeyConnType] = connector.NormalSocket
	params[boot.KeyIP] = IP
	params[boot.KeyPort] = Port
	connConnector.Connect(connector.NormalSocket, IP, Port, params, failFunc)
	for {
		fmt.Println("当前协程数：", runtime.NumGoroutine())
		time.Sleep(time.Second * 1)
	}
}

func failFunc(err error, params map[string]interface{}) {
	fmt.Println(params, "connectFail", err)
	connType := params[boot.KeyConnType].(connector.SocketType)
	ip := params[boot.KeyIP].(string)
	port := params[boot.KeyPort].(int)
	connConnector.Connect(connType, ip, port, params, failFunc)
}

type TestChannelInitializer struct {
}

func NewTestChannelInitializer() (this *TestChannelInitializer) {
	this = new(TestChannelInitializer)
	return
}

func (initializer *TestChannelInitializer) InitChannel(pipeline channel.ChannelPipeline) {
	if pipeline == nil {
		return
	}
	fmt.Println("param key:", pipeline.ContextAttr().GetString("key"))
}
