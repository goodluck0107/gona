package routineCmdMaker

import (
	"gitee.com/andyxt/gona/channelExtends/extends"
	"gitee.com/andyxt/gona/channelExtends/protocol"

	"gitee.com/andyxt/gona/executor"
)

// Outbound
type IOutboundEventMaker interface {
	//发起连接
	MakeConnectEvent(routinePoolId int64, routineId int64, ip string, port int, uID int64, params map[string]interface{}) executor.Event
	//关闭连接
	MakeCloseEvent(routinePoolId int64, routineId int64, uID int64, Desc string) executor.Event
	//下发消息包:OnClose是否在消息发送完毕后关闭连接
	MakeSendMessageEvent(routinePoolId int64, routineId int64, Data protocol.IProtocol, OnClose bool, PoolKey int64, ChlCtx extends.OutterChannelHandlerContext, Desc string) executor.Event
}
