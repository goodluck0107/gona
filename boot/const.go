package boot

const (
	KeyChannelReadLimit       = "KeyChannelReadLimit" // 连接读取消息长度限制
	ChannelReadLimit    int32 = 256                   // 256个字节
	KeyPacketBytesCount       = "KeyPacketBytesCount" // 消息长度占用字节数
	PacketBytesCount    int32 = 4                     // 4个字节
	KeyReadTimeOut            = "KeyReadTimeOut"      // 连接读取消息超时时间
	ReadTimeOut         int32 = 30                    // 30秒
	KeyWriteTimeOut           = "KeyWriteTimeOut"     // 连接写入消息超时时间
	WriteTimeOut        int32 = 30                    // 30秒
	KeyConnType               = "KeyConnType"         // 连接类型
	KeyIP                     = "KeyIP"               // 连接IP
	KeyPort                   = "KeyPort"             // 连接端口
	KeyURLPath                = "URLPath"             // 请求路径
)
