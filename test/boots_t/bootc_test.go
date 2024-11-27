package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"gitee.com/andyxt/gona/boot/boots"
	"gitee.com/andyxt/gona/boot/channel"
)

func TestClient(t *testing.T) {
	testServer()
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
