package channelHandlers

import (
	"github.com/gox-studio/gona/channelExtends/channelConsts"
	"github.com/gox-studio/gona/channelExtends/protocol"
	"github.com/gox-studio/gona/channelExtends/protocol/protocolCoder"
	"github.com/gox-studio/gona/channelExtends/protocol/protocolCoderImpl"
	"github.com/gox-studio/gona/channelExtends/protocol/protocolDefine"

	"github.com/gox-studio/gona/channel"
	"github.com/gox-studio/gona/logger"
)

// DownBase ---> *buffer.ProtocolBuffer
type MessageEncoderHandler struct {
	serializierMap map[int8]protocolCoder.Serializier
}

func NewMessageEncoderHandler(messageFactory protocolCoder.IMessageFactory) (this *MessageEncoderHandler) {
	this = new(MessageEncoderHandler)
	this.serializierMap = this.createSerializierMap(messageFactory)
	return
}

func (encoder *MessageEncoderHandler) createSerializierMap(messageFactory protocolCoder.IMessageFactory) map[int8]protocolCoder.Serializier {
	serializerMap := make(map[int8]protocolCoder.Serializier)
	serializerMap[protocolDefine.CommonSerilizeType] = protocolCoderImpl.NewDefualtSerializier(messageFactory)
	return serializerMap
}

func (encoder *MessageEncoderHandler) ExceptionCaught(ctx channel.ChannelHandlerContext, err error) {
	//	logger.Debug("MessageEncoder ExceptionCaught")
}

func (encoder *MessageEncoderHandler) Write(ctx channel.ChannelHandlerContext, e interface{}) (ret interface{}) {
	//	logger.Debug("MessageEncoder Write")
	buf := e.(protocol.IProtocol)
	encryptType := buf.GetSecurityType()
	ctx.ContextAttr().Set(channelConsts.ChannelSecurityType, encryptType)
	serializeType := buf.GetSerializeType()
	serializier := encoder.serializierMap[encryptType]
	if serializier == nil {
		logger.Error("协议序列化类型无效:encryptType=", encryptType)
	} else {
		pBuf := serializier.Serialize(buf)
		serializeTypeBuf := []byte{byte(serializeType)}
		byteSlice := append(serializeTypeBuf, pBuf...)
		ret = byteSlice
	}
	return
}

func (encoder *MessageEncoderHandler) Close(ctx channel.ChannelHandlerContext) {
	//	logger.Debug("MessageEncoder Close")
}
