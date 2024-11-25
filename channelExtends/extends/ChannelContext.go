package extends

import (
	"fmt"

	"github.com/gox-studio/gona/channelExtends/channelConsts"
)

func ChannelContextToString(chlCtx OutterChannelHandlerContext) string {
	chlCtxID := chlCtx.ID()
	chlCtxUID := UID(chlCtx)
	return fmt.Sprintf("chlCtxID=%s chlCtxUID=%d ", chlCtxID, chlCtxUID)
}

func ChannelContextEquals(chlCtx OutterChannelHandlerContext, other OutterChannelHandlerContext) (ret bool) {
	return chlCtx.ID() == other.ID()
}

// 放入用户信息代表后续Command事件中可以执行派发Event事件的行为了
func PutInUserInfo(chlCtx OutterChannelHandlerContext, uID int64, lngType int8) {
	chlCtx.ContextAttr().Set("poolKey", uID)     // 用户在连接列表中的Key，目前使用用户UID标识
	chlCtx.ContextAttr().Set("lngType", lngType) // 用户语言类型
	chlCtx.ContextAttr().Set("isInPool", true)   // 连接是否添加到连接池
}

func HasUserInfo(chlCtx OutterChannelHandlerContext) bool {
	return chlCtx.ContextAttr().GetBool("isInPool")
}

func UID(chlCtx OutterChannelHandlerContext) int64 {
	return chlCtx.ContextAttr().GetInt64("poolKey")
}

func GetLngType(chlCtx OutterChannelHandlerContext) int8 {
	return chlCtx.ContextAttr().GetInt8("lngType")
}

func SetFireUser(chlCtx OutterChannelHandlerContext, FireUserID int64) {
	chlCtx.ContextAttr().Set(channelConsts.ChannelFireUser, FireUserID) // 连接ID,通常是玩家uid
}

func GetFireUser(chlCtx OutterChannelHandlerContext) int64 {
	return chlCtx.ContextAttr().GetInt64(channelConsts.ChannelFireUser)
}

/**
 * 设置以前的用户连接为废弃，废弃连接的后续消息都将不处理
 * */
func Conflict(chlCtx OutterChannelHandlerContext) {
	chlCtx.ContextAttr().Set("userConflict", true) // 连接是否废弃,当有用户重连与异地登陆时候，以前的连接会被置为废弃
}

func IsConflict(chlCtx OutterChannelHandlerContext) (ret bool) {
	return chlCtx.ContextAttr().GetBool("userConflict")
}

func IsClose(chlCtx OutterChannelHandlerContext) (ret bool) {
	return chlCtx.ContextAttr().GetBool("isClose")
}

func Close(chlCtx OutterChannelHandlerContext) {
	chlCtx.Close()
	chlCtx.ContextAttr().Set("isClose", true) //是否已经关闭
}

func IsLogout(chlCtx OutterChannelHandlerContext) (ret bool) {
	return chlCtx.ContextAttr().GetBool("isLogout")
}

func Logout(chlCtx OutterChannelHandlerContext) {
	chlCtx.ContextAttr().Set("isLogout", true) //是否已经登出
}

func IsSystemKick(chlCtx OutterChannelHandlerContext) (ret bool) {
	return chlCtx.ContextAttr().GetBool("isSystemKick")
}

func SystemKick(chlCtx OutterChannelHandlerContext) {
	chlCtx.ContextAttr().Set("isSystemKick", true) //是否已经被系统踢出
}

func IsOfflie(chlCtx OutterChannelHandlerContext) (ret bool) {
	return chlCtx.ContextAttr().GetBool("isOfflie")
}

func Offlie(chlCtx OutterChannelHandlerContext) {
	chlCtx.ContextAttr().Set("isOfflie", true) //是否已经断开
}
