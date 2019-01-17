package bgp

import (
	"encoding/binary"
	"fmt"
	"github.com/adamb/go_osegp/bgp/errors"
)

/*
This file contains definitions for parsing a BGP packet header
The packet parsing flow looks like this:
-> Init the struct for the part of the packet to parse, passing a byte slice starting at the offset of the part
	-> Init the fields that correspond to the struct
		-> Iterate through the fields, reading their values in
*/
// Field lengths
const MARKER_LENGTH = 16
const PLENGTH_LENGTH = 2
const TYPE_LENGTH = 1

// Field constants
const MESSAGE_OPEN = 1
const MESSAGE_UPDATE = 2
const MESSAGE_NOTIFICATION = 3
const MESSAGE_KEEPALIVE = 4

// BGP Packet header
type BgpHeader struct {
	Marker *Marker
	Length *Length
	Type   *Type
	fields []Field
}

// Init the header and fields.
// Requires: byte slice, starting at header beginning
func ReadHeader(b []byte) BgpHeader {
	bgp := BgpHeader{}
	bgp.Marker = MakeMarker()
	bgp.Length = MakeLength()
	bgp.Type = MakeType()

	bgp.fields = []Field{
		bgp.Marker,
		bgp.Length,
		bgp.Type,
	}
	// Start byte offset for the header is aaaallways zero
	offset := uint16(0)
	// Iterate through each field and populate the values
	for _, f := range bgp.fields {
		l := f.GetLength()
		if int(offset+l) > len(b) {
			errors.RaiseError(fmt.Sprintf("Invalid BGP packet header. Expected %v more bytes.", l))
		}
		f.Read(b[offset:])
		offset = offset + l
	}

	fmt.Printf("Header length: %v", bgp.Length.Value())
	return bgp
}

// DummyHeader returns a byte array of a header filled with dummy data
// Useful for testing the unmarshaling of headers.
func MakeDummyHeader() []byte {
	var b []byte
	fields := []Field{
		MakeMarker(),
		MakeLength(),
		MakeType(),
	}
	for _, f := range fields {
		b = append(b, f.Dummy()...)
	}
	return b
}

/*
########
Field definitions
########
*/

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
func (f *Marker) Value() interface{} {
	return 0
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
func (f *Length) Value() interface{} {
	return f.value
}

func (f *Length) Dummy() []byte {
	var b []byte
	b = []byte{0, 45}
	return b
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
	t.length = TYPE_LENGTH
	return &t
}
func (f *Type) Read(b []byte) {
	f.value = uint8(b[0])
}
func (f *Type) Value() interface{} {
	return f.value
}
func (f *Type) Dummy() []byte {
	var b []byte
	b = []byte{1}
	return b
}
