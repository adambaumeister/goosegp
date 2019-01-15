package bgp

/*
Go BGP Parser

Takes incoming BGP messages in Byte format and decodes them returning useable structs
*/
type Parser struct {
}

func (p *Parser) Parse(b []byte) {
	MakeHeader(b)
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

func (fb *fieldBase) GetLength() uint16 {
	return fb.length
}

// Field is the interface that all packet fields must implement.
// Field is a single field/value within a packet and associated methods to help with manipulating said fields
type Field interface {
	GetLength() uint16
	Read([]byte)
	Dummy() []byte
	Value() interface{}
}
