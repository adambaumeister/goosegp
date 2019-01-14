package bgp

import "encoding/binary"

const MARKER_LENGTH = 16
const PLENGTH_LENGTH = 2

// Marker
// Used for interoperability.
type Marker struct {
	fieldBase
}

func MakeMarker() *Marker {
	m := Marker{}
	m.length = 16
	return &m
}

// This func does nothing. Who cares, it's a marker!
func (f *Marker) Read([]byte) {
}

// Length
// Total BGP message length, including headers.
type Length struct {
	fieldBase
	value uint16
}

func MakeLength() *Length {
	l := Length{}
	l.length = 2
	return &l
}
func (f *Length) Read(b []byte) {
	f.value = binary.BigEndian.Uint16(b)
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
