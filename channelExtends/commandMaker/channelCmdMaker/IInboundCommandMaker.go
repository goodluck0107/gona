package channelCmdMaker

import "gona/channelExtends/extends"

type IInboundCommandMaker interface {
	//触发异常
	MakeExceptionCommand(ctx extends.OutterChannelHandlerContext, err error) ICommand

	//新连接
	MakeActiveCommand(routinePoolId int64, Ctx extends.OutterChannelHandlerContext) ICommand
	//连接中断
	MakeInActiveCommand(routinePoolId int64, Ctx extends.OutterChannelHandlerContext) ICommand
	//收到消息包
	MakeMessageReceivedCommand(routinePoolId int64, Ctx extends.OutterChannelHandlerContext, Data interface{}) ICommand
}
