package routineCommands

import (
	"github.com/gox-studio/gona/channelExtends/channelConsts"
	"github.com/gox-studio/gona/session"

	"github.com/gox-studio/gona/bootStrap/bootStrapClient/listener"
	"github.com/gox-studio/gona/logger"
	"github.com/gox-studio/gona/utils"
)

type ClientRoutineInboundCmdConnect struct {
	routinePoolId int64
	routineId     int64
	uID           int64
	ip            string
	port          int
	params        map[string]interface{}
	connector     listener.IConnector
}

func NewClientRoutineInboundCmdConnect(routinePoolId int64, routineId int64, uID int64, ip string, port int, params map[string]interface{}, connector listener.IConnector) (this *ClientRoutineInboundCmdConnect) {
	this = new(ClientRoutineInboundCmdConnect)
	this.routinePoolId = routinePoolId
	this.routineId = routineId
	this.uID = uID
	this.ip = ip
	this.port = port
	this.params = params
	this.connector = connector
	return
}

func (upConnectEvent *ClientRoutineInboundCmdConnect) PoolId() int64 {
	return upConnectEvent.routinePoolId

}
func (upConnectEvent *ClientRoutineInboundCmdConnect) QueueId() int64 {
	return upConnectEvent.routineId
}

func (upConnectEvent *ClientRoutineInboundCmdConnect) Wait() (result interface{}, ok bool) {
	return nil, true
}

func (upConnectEvent *ClientRoutineInboundCmdConnect) Exec() {
	logger.Debug("ClientRoutineInboundCmdConnect Exec")
	iSession := session.GetSession(0, upConnectEvent.uID) // find in online players
	if iSession == nil {
		logger.Debug("ClientRoutineInboundCmdConnect iSession == nil")
		iSession = session.NewSession(utils.GetRandomPassword(), upConnectEvent.uID)
		session.AddSession(0, iSession)
	}
	upConnectEvent.params[channelConsts.ChannelIp] = upConnectEvent.ip
	upConnectEvent.params[channelConsts.ChannelPort] = upConnectEvent.port
	upConnectEvent.params[channelConsts.ChannelFireUser] = upConnectEvent.uID
	upConnectEvent.params[channelConsts.ChannelTag] = ""
	upConnectEvent.params[channelConsts.ChannelParams] = ""
	upConnectEvent.connector.Connect(upConnectEvent.ip, upConnectEvent.port, upConnectEvent.params)
}
