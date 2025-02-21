package boot

const (
	KeyChannelReadLimit  string = "KeyChannelReadLimit"  // 连接读取消息长度限制
	ChannelReadLimit     int32  = 256                    // 256个字节
	KeyPacketBytesCount  string = "KeyPacketBytesCount"  // 消息长度占用字节数
	PacketBytesCount     int32  = 4                      // 4个字节
	KeyReadTimeOut       string = "KeyReadTimeOut"       // 连接读取消息超时时间
	ReadTimeOut          int32  = 30                     // 30秒
	KeyWriteTimeOut      string = "KeyWriteTimeOut"      // 连接写入消息超时时间
	WriteTimeOut         int32  = 30                     // 30秒
	KeyConnType          string = "KeyConnType"          // 连接类型
	KeyIP                string = "KeyIP"                // 连接IP
	KeyPort              string = "KeyPort"              // 连接端口
	KeyURLPath           string = "URLPath"              // 请求路径
	KeyIsLD              string = "KeyIsLD"              // 是否小端
	KeyLengthInclude     string = "KeyLengthInclude"     // 包长度是否包含自己的长度
	SkipPacketBytesCount string = "SkipPacketBytesCount" // 是否跳过包长度
)
