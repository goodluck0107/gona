package channelHandlers

import (
	"gona/channelExtends/protocol/protocolCoder"
	"gona/channelExtends/protocol/protocolCoderImpl"
	"gona/channelExtends/protocol/protocolDefine"

	"gona/channel"
	"gona/logger"
)

// *buffer.ProtocolBuffer ---> UpBase
type MessageDecoderHandler struct {
	serializierMap map[int8]protocolCoder.Serializier
}

func NewMessageDecoderHandler(messageFactory protocolCoder.IMessageFactory) (this *MessageDecoderHandler) {
	this = new(MessageDecoderHandler)
	this.serializierMap = this.createSerializierMap(messageFactory)
	return
}

func (decoder *MessageDecoderHandler) createSerializierMap(messageFactory protocolCoder.IMessageFactory) map[int8]protocolCoder.Serializier {
	serializerMap := make(map[int8]protocolCoder.Serializier)
	serializerMap[protocolDefine.CommonSerilizeType] = protocolCoderImpl.NewDefualtSerializier(messageFactory)
	return serializerMap
}

func (decoder *MessageDecoderHandler) MessageReceived(ctx channel.ChannelHandlerContext, e interface{}) (ret interface{}, goonNext bool) {
	//	logger.Debug("MessageDecoder MessageReceived")
	byteSlice := e.([]byte)
	serializeType := int8(byteSlice[0]) //序列化类型
	serializier := decoder.serializierMap[serializeType]
	if serializier == nil {
		logger.Info("关闭连接：", " 关闭原因：协议结构类型无效:IP=", ctx.RemoteAddr(), "结构类型：", serializeType)
		ctx.Close()
		return
	}
	byteSlice = byteSlice[1:]
	valid, msg := serializier.Deserialize(byteSlice)
	if !valid {
		logger.Info("关闭连接：", " 关闭原因：协议解析失败:IP=", ctx.RemoteAddr()) //, " , 协议号：", VersionId, UserId, AppId, MessageId)
		ctx.Close()
		return
	}
	ret = msg
	goonNext = true
	return
}

func (decoder *MessageDecoderHandler) ChannelActive(ctx channel.ChannelHandlerContext) (goonNext bool) {
	//	logger.Debug("MessageDecoder ChannelActive")
	return true
}
func (decoder *MessageDecoderHandler) ChannelInactive(ctx channel.ChannelHandlerContext) (goonNext bool) {
	//	logger.Debug("MessageDecoder ChannelInactive")
	return true
}

func (decoder *MessageDecoderHandler) ExceptionCaught(ctx channel.ChannelHandlerContext, err error) {
	//	logger.Debug("MessageDecoder ExceptionCaught")
}
