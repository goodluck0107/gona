package service

import (
	"reflect"

	"gona/channelExtends/extends"
	"gona/session"
	"gona/utils"
)

type IHandleChecker interface {
	// IsHandlerMethod decide a method is suitable handler method
	IsHandlerMethod(method reflect.Method) bool
	// AdaptArgs create the params a handler method need
	AdaptArgs(types []reflect.Type, data []interface{}) []reflect.Value
}

type IRouteMapper interface {
	GetCodeForPath(reqPath string) int32
	GetPathForCode(code int32) string
}
type IServiceContext utils.IAttr
type ISession session.ISession
type IChannelContext extends.OutterChannelHandlerContext

type IServiceRequest interface {
	ChannelContext() IChannelContext
	ReqContext() IServiceContext
}
type IServiceResponse interface {
	Write(interface{}, ISerializer)
}

type ISerializer interface {
	Serialize(message interface{}) ([]byte, error)
}
