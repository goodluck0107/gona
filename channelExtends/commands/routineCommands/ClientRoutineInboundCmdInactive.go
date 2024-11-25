package routineCommands

import (
	"gona/channelExtends/channelConsts"
	"gona/channelExtends/extends"
	"gona/session"

	bootstrap "gona/bootStrap"
	"gona/bootStrap/bootStrapClient/listener"
	"gona/logger"
)

type ClientRoutineInboundCmdInactive struct {
	routinePoolId int64
	routineId     int64
	connector     listener.IConnector
	ChlCtx        extends.OutterChannelHandlerContext
}

func NewClientRoutineInboundCmdInactive(routinePoolId int64, routineId int64,
	ChlCtx extends.OutterChannelHandlerContext, connector listener.IConnector) (this *ClientRoutineInboundCmdInactive) {
	this = new(ClientRoutineInboundCmdInactive)
	this.routinePoolId = routinePoolId
	this.routineId = routineId
	this.ChlCtx = ChlCtx
	this.connector = connector
	return
}

func (inactiveEvent *ClientRoutineInboundCmdInactive) PoolId() int64 {
	return inactiveEvent.routinePoolId
}

func (inactiveEvent *ClientRoutineInboundCmdInactive) QueueId() int64 {
	return inactiveEvent.routineId
}

func (inactiveEvent *ClientRoutineInboundCmdInactive) Wait() (result interface{}, ok bool) {
	return nil, true
}

func (inactiveEvent *ClientRoutineInboundCmdInactive) Exec() {
	logger.Debug("ClientRoutineInboundCmdInactive Exec")
	if extends.IsConflict(inactiveEvent.ChlCtx) {
		logger.Debug("ClientRoutineInboundCmdInactive 连接已经被其他连接挤下线，不需要重连", extends.UID(inactiveEvent.ChlCtx))
		return
	}
	if extends.IsClose(inactiveEvent.ChlCtx) {
		logger.Debug("ClientRoutineInboundCmdInactive 连接已经被主动关闭，不需要重连", extends.UID(inactiveEvent.ChlCtx))
		return
	}
	if extends.IsLogout(inactiveEvent.ChlCtx) {
		logger.Debug("ClientRoutineInboundCmdInactive 连接已经被主动登出，不需要重连", extends.UID(inactiveEvent.ChlCtx))
		return
	}
	if extends.IsSystemKick(inactiveEvent.ChlCtx) {
		logger.Debug("ClientRoutineInboundCmdInactive 连接已经被踢出，不需要重连", extends.UID(inactiveEvent.ChlCtx))
		return
	}
	uId := extends.UID(inactiveEvent.ChlCtx)
	iSession := session.GetSession(0, uId)
	if iSession == nil {
		logger.Debug("ClientRoutineInboundCmdInactive 连接已经被主动关闭，不需要重连", extends.UID(inactiveEvent.ChlCtx))
		return
	}
	//default: broken reconnect
	logger.Debug("ClientRoutineInboundCmdInactive 连接中断，需要重连", extends.UID(inactiveEvent.ChlCtx))
	ip := inactiveEvent.ChlCtx.ContextAttr().GetString(channelConsts.ChannelIp)
	port := inactiveEvent.ChlCtx.ContextAttr().GetInt(channelConsts.ChannelPort)
	channelTag := inactiveEvent.ChlCtx.ContextAttr().GetString(channelConsts.ChannelTag)
	jsonData := inactiveEvent.ChlCtx.ContextAttr().GetString(channelConsts.ChannelParams)
	channelReadLimit := inactiveEvent.ChlCtx.ContextAttr().GetString(bootstrap.KeyChannelReadLimit)
	params := make(map[string]interface{})
	params[channelConsts.ChannelIp] = ip
	params[channelConsts.ChannelPort] = port
	params[channelConsts.ChannelFireUser] = uId
	params[channelConsts.ChannelTag] = channelTag
	params[channelConsts.ChannelParams] = jsonData
	params[bootstrap.KeyChannelReadLimit] = channelReadLimit
	inactiveEvent.connector.Connect(ip, port, params)
}
