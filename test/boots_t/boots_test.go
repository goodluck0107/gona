package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gona/boot/boots"
	"gitee.com/andyxt/gona/boot/channel"
)

func TestClient(t *testing.T) {
	testServer()
}

func testServer() {
	params := make(map[string]interface{})
	params["key"] = "serverValue"
	params[boot.KeyPacketBytesCount] = 2
	params[boot.KeyChannelReadLimit] = 512
	params[boot.KeyReadTimeOut] = 30
	params[boot.KeyWriteTimeOut] = 30
	boots.Serve(boots.WithTCPAddr(":20000"), boots.WithChannelParams(params),
		boots.WithInitializer(NewTestChannelInitializer()))
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
