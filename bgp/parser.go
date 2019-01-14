package bgp

import "fmt"

/*
Go BGP Parser

Takes incoming BGP messages in Byte format and decodes them returning useable structs
*/
type Parser struct {
}

func (p *Parser) Parse([]byte) {
	// Use field Initilization methods to retrieve refs
	headerFields := []Field{
		MakeMarker(),
		MakeLength(),
	}

	for _, h := range headerFields {
		fmt.Printf("Len: %v\n", h.GetLength())
	}
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
// GetLength() is implemented by FieldBase
type Field interface {
	GetLength() uint16
}
