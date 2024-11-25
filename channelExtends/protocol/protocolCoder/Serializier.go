package protocolCoder

import (
	"gona/channelExtends/protocol"
)

type Serializier interface {
	Serialize(protocol.IProtocol) []byte
	Deserialize(b []byte) (bool, protocol.IProtocol)
}
