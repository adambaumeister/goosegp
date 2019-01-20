package packet

/*
Go BGP Parser

Takes incoming BGP messages in Byte format and decodes them returning useable structs
*/
type Parser struct {
}

// Take an incoming byte slice and convert to structs
func (p *Parser) Parse(b []byte) BgpPacket {
	packet := BgpPacket{}
	// Firt the headers are made
	h := ReadHeader(b)
	headerLength := h.GetLength()
	packet.Header = h

	switch h.Type.value {
	case MESSAGE_OPEN:
		m := ReadMsgOpen(b[headerLength:])
		packet.Message = m
	case MESSAGE_KEEPALIVE:
		return packet
	}
	return packet
}

// Field base is the basic construct for each field
// Contains:
//	- b: Byte array of original field data
//  - length: Length of field
// This struct is emebedded into the field types.
type fieldBase struct {
	b      []byte
	length uint16
}

// Return the length, in bytes, of this field.
func (fb *fieldBase) GetLength() uint16 {
	return fb.length
}

// Field is the interface that all packet fields must implement.
// Field is a single field/value within a packet and associated methods to help with manipulating said fields
type Field interface {
	GetLength() uint16
	Read([]byte)
	Value() interface{}
	Serialize() []byte
}
