package bgp

import (
	"encoding/binary"
)

const MARKER_LENGTH = 16
const PLENGTH_LENGTH = 2

// BGP Packet header
type BgpHeader struct {
	Marker *Marker
	Length *Length
	Type   *Type
	fields []Field
}

// Init the header and fields.
// Requires: byte slice, starting at header beginning
func MakeHeader(b []byte) BgpHeader {
	bgp := BgpHeader{}
	bgp.Marker = MakeMarker()
	bgp.Length = MakeLength()
	bgp.Type = MakeType()
	bgp.fields = []Field{
		bgp.Marker,
		//bgp.Length,
		//bgp.Type,
	}
	// Start byte offset for the header is aaaallways zero
	offset := uint16(0)
	// Iterate through each field and populate the values
	for _, f := range bgp.fields {
		l := f.GetLength()
		f.Read(b[offset:])
		offset = l
	}

	return bgp
}

// DummyHeader returns a byte array of a header filled with dummy data
// Useful for testing the unmarshaling of headers.
func MakeDummyHeader() []byte {
	var b []byte
	m := Marker{}
	b = m.Dummy()
	return b
}

// Marker
// Used for interoperability.
type Marker struct {
	fieldBase
}

// Init new Marker field
func MakeMarker() *Marker {
	m := Marker{}
	m.length = MARKER_LENGTH
	return &m
}

// This func does nothing. Who cares, it's a marker!
func (f *Marker) Read([]byte) {
}

// Dummy returns a byte representation of this field filled with dummy data
// Useful for debuggin
func (f *Marker) Dummy() []byte {
	var b []byte
	b = []byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
	return b
}

// Length
// Total BGP message length, including headers.
type Length struct {
	fieldBase
	value uint16
}

// Init new Length field
func MakeLength() *Length {
	l := Length{}
	l.length = PLENGTH_LENGTH
	return &l
}
func (f *Length) Read(b []byte) {
	l := f.GetLength()
	f.value = binary.BigEndian.Uint16(b[:l])
}

// Type
// Type of BGP Message that follows.
//
// 1 - OPEN
// 2 - UPDATE
// 3 - NOTIFICATION
// 4 - KEEPALIVE
type Type struct {
	fieldBase
	value uint8
}

func MakeType() *Type {
	t := Type{}
	t.length = 1
	return &t
}
func (f *Type) Read(b []byte) {
	f.value = uint8(b[0])
}
