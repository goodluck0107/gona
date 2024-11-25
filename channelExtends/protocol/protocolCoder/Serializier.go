package protocolCoder

import (
	"github.com/gox-studio/gona/channelExtends/protocol"
)

type Serializier interface {
	Serialize(protocol.IProtocol) []byte
	Deserialize(b []byte) (bool, protocol.IProtocol)
}
