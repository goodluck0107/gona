package protocol

/*
一条socket链接建立以后，发送给服务端的第一条消息必须实现该接口，用于分配处理该链接的协程
*/
type ILoginMsg interface {
	GetLoginUid() int64
	IsValid() bool
	GetLngType() int8
}
