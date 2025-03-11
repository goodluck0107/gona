package boot

const (
	KeyConnType string = "ConnType" // 链接类型 http ws tcp
	KeyURLPath  string = "URLPath"  // 请求路径

	ConnTypeTcp  string = "tcp"
	ConnTypeWs   string = "ws"
	ConnTypeHttp string = "http"
)

func IsConnTcp(connType string) bool {
	return connType == ConnTypeTcp
}
func IsConnWs(connType string) bool {
	return connType == ConnTypeWs
}
func IsConnHttp(connType string) bool {
	return connType == ConnTypeHttp
}
