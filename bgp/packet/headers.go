package packet

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

// Serialize a BGP Header for the wire
func (bgp *BgpHeader) Serialize() []byte {
	var b []byte
	for _, f := range bgp.fields {
		fb := f.Serialize()
		b = append(b, fb...)
	}
	return b
}

func (bgp *BgpHeader) Init() {
	bgp.Marker = MakeMarker()
	bgp.Length = MakeLength()
	bgp.Type = MakeType()

	bgp.fields = []Field{
		bgp.Marker,
		bgp.Length,
		bgp.Type,
	}
}

// Parse (unmarshal) a BGP packet Header and fields.
// Requires: byte slice, starting at header beginning
func ReadHeader(b []byte) BgpHeader {
	bgp := BgpHeader{}
	bgp.Init()

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
	return bgp
}

// MakeHeader crafts an empty packet header
func MakeHeader() BgpHeader {
	bgp := BgpHeader{}
	bgp.Init()

	return bgp
}

func (bgp *BgpHeader) GetLength() uint16 {
	var l uint16
	for _, f := range bgp.fields {
		l = l + f.GetLength()
	}
	return l
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
func (f *Marker) Serialize() []byte {
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
func (f *Length) Write(v uint16) {
	f.value = v
}
func (f *Length) Serialize() []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, f.value)
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
func (f *Type) Write(v uint8) {
	f.value = v
}
func (f *Type) Serialize() []byte {
	return []byte{f.value}
}
