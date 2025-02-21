package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"gitee.com/andyxt/gona/boot"
	"gitee.com/andyxt/gona/boot/boots"
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/logger"
)

func TestClient(t *testing.T) {
	testServer()
}

func testServer() {
	params := make(map[string]interface{})
	params["key"] = "serverValue"
	params[boot.KeyPacketBytesCount] = 2
	params[boot.SkipPacketBytesCount] = true
	boots.Serve(
		boots.WithHttpAddr(":20000"),
		boots.WithChannelParams(params),
		boots.WithInitializer(NewTestChannelInitializer()),
		boots.WithLogger(logger.GetLogger()),
		boots.WithReadTimeOut(30),
		boots.WithWriteTimeOut(30),
		boots.WithReadLimit(512),
	)
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
