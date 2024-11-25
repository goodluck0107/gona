package main

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"time"

	"gitee.com/andyxt/gona/bootStrap/bootStrapClient"
	"gitee.com/andyxt/gona/bootStrap/bootStrapClient/connector"
	"gitee.com/andyxt/gona/bootStrap/bootStrapServer"
	"gitee.com/andyxt/gona/channel"
	"gitee.com/andyxt/gona/executor"
)

func main() {
	// testServer()
	testClient()
}

func testClient() {
	executor.Init(NewRoutinePoolBuilder())
	bootStrap :=
		bootStrapClient.NewClientBootStrap(connector.NormalSocket, 1)
	connector := bootStrap.GetConnector()
	bootStrap.
		ChannelInitializer(
			NewTestChannelInitializer(1))
	bootStrap.Listen()
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
	executor.Init(NewRoutinePoolBuilder())
	params := make(map[string]interface{})
	params["key"] = "serverValue"
	bootStrap :=
		bootStrapServer.NewServerBootStrap().
			Params(params).
			Port(":20000").
			ChannelInitializer(
				NewTestChannelInitializer(1))
	go bootStrap.Listen()
	for {
		fmt.Println("当前协程数：", runtime.NumGoroutine())
		time.Sleep(time.Second * 1)
	}
}

type TestChannelInitializer struct {
	mEventRoutinePoolID int64
}

func NewTestChannelInitializer(eventRoutinePoolID int64) (this *TestChannelInitializer) {
	this = new(TestChannelInitializer)
	this.mEventRoutinePoolID = eventRoutinePoolID
	return
}

func (initializer *TestChannelInitializer) InitChannel(pipeline channel.ChannelPipeline) {
	if pipeline == nil {
		return
	}
	fmt.Println("param key:", pipeline.ContextAttr().GetString("key"))
}

type RoutinePoolBuilder struct {
	mRoutinePools map[int64]*executor.RoutinePool
}

func NewRoutinePoolBuilder() (ret *RoutinePoolBuilder) {
	ret = new(RoutinePoolBuilder)
	ret.mRoutinePools = make(map[int64]*executor.RoutinePool)
	ret.init()
	return
}

func (builder *RoutinePoolBuilder) init() {
	builder.mRoutinePools[0] = executor.NewRoutinePool(64, 10000)
	builder.mRoutinePools[1] = executor.NewRoutinePool(64, 10000)
}

func (builder *RoutinePoolBuilder) GetRoutinePool(key int64) *executor.RoutinePool {
	if routinePool, ok := builder.mRoutinePools[key]; ok {
		return routinePool
	}
	panic(errors.New("Invalid RoutinePoolKey:" + strconv.Itoa(int(key))))
}
