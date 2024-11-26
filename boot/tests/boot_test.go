package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"gitee.com/andyxt/gona/boot/bootc"
	"gitee.com/andyxt/gona/boot/bootc/connector"
	"gitee.com/andyxt/gona/boot/boots"
	"gitee.com/andyxt/gona/boot/channel"
)

func TestClient(t *testing.T) {
	// testServer()
	testClient()
}

func testClient() {
	bc :=
		bootc.NewClientBootStrap(connector.NormalSocket, 1)
	connector := bc.GetConnector()
	bc.
		ChannelInitializer(
			NewTestChannelInitializer())
	bc.Listen()
	fmt.Println("Connect")
	params := make(map[string]interface{})
	params["key"] = "clientValue"
	connector.Connect("127.0.0.1", 20000, params)

	for {
		fmt.Println("当前协程数：", runtime.NumGoroutine())
		time.Sleep(time.Second * 1)
	}
}

func testServer() {
	params := make(map[string]interface{})
	params["key"] = "serverValue"
	bs :=
		boots.NewServerBootStrap().
			Params(params).
			Port(":20000").
			ChannelInitializer(
				NewTestChannelInitializer())
	go bs.Listen()
	for {
		fmt.Println("当前协程数：", runtime.NumGoroutine())
		time.Sleep(time.Second * 1)
	}
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
