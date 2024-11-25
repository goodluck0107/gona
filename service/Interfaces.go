package service

import (
	"reflect"

	"github.com/gox-studio/gona/channelExtends/extends"
	"github.com/gox-studio/gona/session"
	"github.com/gox-studio/gona/utils"
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
