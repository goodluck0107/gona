package channelHandlers

import (
	"gitee.com/andyxt/gona/channelExtends/channelConsts"
	"gitee.com/andyxt/gona/channelExtends/protocol/protocolCoder"
	"gitee.com/andyxt/gona/channelExtends/protocol/protocolCoderImpl"
	"gitee.com/andyxt/gona/channelExtends/protocol/protocolDefine"

	"gitee.com/andyxt/gona/channel"
	"gitee.com/andyxt/gona/logger"
)

// *ProtocolBuffer-->*ProtocolBuffer
type SecurityEncoderHandler struct {
	securitierMap map[int8]protocolCoder.Securitier
}

func NewSecurityEncoderHandler() (this *SecurityEncoderHandler) {
	this = new(SecurityEncoderHandler)
	this.securitierMap = this.createSerializierMap()
	return
}

func (encoder *SecurityEncoderHandler) createSerializierMap() map[int8]protocolCoder.Securitier {
	mSecurityMap := make(map[int8]protocolCoder.Securitier)
	mSecurityMap[protocolDefine.CommonSecurityType] = protocolCoderImpl.NewDefualtSecuritier()
	return mSecurityMap
}

func (encoder *SecurityEncoderHandler) Write(ctx channel.ChannelHandlerContext, e interface{}) (ret interface{}) {
	//logger.Debug("SecurityEncoder Write-0")
	buf := e.([]byte)
	encryptType := ctx.ContextAttr().GetInt8(channelConsts.ChannelSecurityType)
	if encryptType <= 0 { //不加密
		ret = buf
		//logger.Debug("SecurityEncoder Write1")
	} else {
		//logger.Debug("SecurityEncoder Write2")
		securitier := encoder.securitierMap[encryptType]
		if securitier == nil {
			logger.Error("协议安全类型无效:encryptType=", encryptType)
			ret = buf
			//logger.Debug("SecurityEncoder Write3")
		} else {
			securityBuf := securitier.Encrypt(buf)
			securityTypeBuf := []byte{byte(encryptType)}
			buf = append(securityTypeBuf, securityBuf...)
			ret = buf
			//logger.Debug("SecurityEncoder Write4")
		}

	}
	return
}

func (encoder *SecurityEncoderHandler) Close(ctx channel.ChannelHandlerContext) {
	//	logger.Debug("SecurityEncoder Close")
}
func (encoder *SecurityEncoderHandler) ExceptionCaught(ctx channel.ChannelHandlerContext, err error) {
	//	logger.Debug("SecurityEncoder ExceptionCaught")
}
