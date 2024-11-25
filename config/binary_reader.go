package config

type BinaryReader struct {
	names func() []string
	bytes func(string) ([]byte, error)
}

func NewBinaryReader(names func() []string, bytes func(string) ([]byte, error)) *BinaryReader {
	return &BinaryReader{
		names: names,
		bytes: bytes,
	}
}

func (b *BinaryReader) Names() []string {
	return b.names()
}

func (b *BinaryReader) Bytes(name string) ([]byte, error) {
	return b.bytes(name)
}
