package service

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewProtoResponseSerializer() ISerializer {
	return new(protoResponseSerializer)
}

type protoResponseSerializer struct {
}

func (resp *protoResponseSerializer) Serialize(message interface{}) ([]byte, error) {
	if message == nil {
		return []byte(""), nil
	}
	respData := message.(protoreflect.ProtoMessage)
	return proto.Marshal(respData)
}
