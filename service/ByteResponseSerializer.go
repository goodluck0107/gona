package service

func NewByteResponseSerializer() ISerializer {
	return new(byteResponseSerializer)
}

type byteResponseSerializer struct {
}

func (resp *byteResponseSerializer) Serialize(message interface{}) ([]byte, error) {
	if message == nil {
		return []byte(""), nil
	}
	return message.([]byte), nil
}
