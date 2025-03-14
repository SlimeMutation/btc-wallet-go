package utils

type Bytes []byte

func (b Bytes) Bytes() []byte {
	return b[:]
}
func (b *Bytes) SetBytes(bytes []byte) {
	*b = bytes
}
