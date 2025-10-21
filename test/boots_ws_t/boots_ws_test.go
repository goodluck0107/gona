package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/goodluck0107/gona/boot/boots"
	"github.com/goodluck0107/gona/boot/channel"
)

func TestClient(t *testing.T) {
	testServer()
}

func testServer() {
	boots.Serve(
		boots.WithHttpAddr(":20000"),
		boots.WithInitializer(NewTestChannelInitializer()),
		boots.WithReadTimeOut(30),
		boots.WithWriteTimeOut(30),
		boots.WithReadLimit(512),
		boots.WithPacketBytesCount(2),
		boots.WithSkipPacketBytesCount(),
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
