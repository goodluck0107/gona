package bootc

type DefaultMessageSpliter struct {
}

func NewDefaultMessageSpliter() (this *DefaultMessageSpliter) {
	this = new(DefaultMessageSpliter)
	return
}

func (builder *DefaultMessageSpliter) GetBytesCountForMessageLength() int8 {
	return 4
}
