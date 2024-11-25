package channel

type MessageSpliter interface {
	GetBytesCountForMessageLength() int8
}
