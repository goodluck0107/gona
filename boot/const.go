package boot

const (
	KeyChannelReadLimit  string = "KeyChannelReadLimit"  // 连接读取消息长度限制
	KeyPacketBytesCount  string = "KeyPacketBytesCount"  // 消息长度占用字节数
	KeyIsLD              string = "KeyIsLD"              // 是否小端
	KeyLengthInclude     string = "KeyLengthInclude"     // 包长度是否包含自己的长度
	SkipPacketBytesCount string = "SkipPacketBytesCount" // 是否跳过包长度

	KeyConnType string = "KeyConnType" // 连接类型
	KeyIP       string = "KeyIP"       // 连接IP
	KeyPort     string = "KeyPort"     // 连接端口
	KeyURLPath  string = "URLPath"     // 请求路径
)
