package packet

type Message interface {
	Serialize() []byte
}

type BgpPacket struct {
	Header BgpHeader

	Message Message
}

// Serialize returns the byte value of an entire bgp message
func (bgp BgpPacket) Serialize() []byte {
	var b []byte
	b = append(b, bgp.Header.Serialize()...)
	b = append(b, bgp.Message.Serialize()...)
	return b
}
