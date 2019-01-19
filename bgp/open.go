package bgp

import (
	"encoding/binary"
	"fmt"
	"github.com/adamb/go_osegp/bgp/errors"
	"net"
)

type BgpMsgOpen struct {
	Version          *Version
	AutonomousSystem *AutonomousSystem
	HoldTime         *HoldTime
	Identifier       *Identifier
	Params           *OptionalParams
	fields           []Field

	Length uint16
}

// Serialize a BGP Header for the wire
func (bgp *BgpMsgOpen) Serialize() []byte {
	var b []byte
	for _, f := range bgp.fields {
		fb := f.Serialize()
		b = append(b, fb...)
	}
	return b
}

// Populate all the fields within an OPEN message
func (bgp *BgpMsgOpen) Init() {
	bgp.Version = MakeVersion()
	bgp.AutonomousSystem = MakeAutonomousSystem()
	bgp.HoldTime = MakeHoldTime()
	bgp.Identifier = MakeIdentifier()
	bgp.Params = MakeOptionalParams()
	bgp.fields = []Field{
		bgp.Version,
		bgp.AutonomousSystem,
		bgp.HoldTime,
		bgp.Identifier,
		bgp.Params,
	}
}

// Read (unmarshal) an OPEN message
func ReadMsgOpen(b []byte) BgpMsgOpen {
	bgp := BgpMsgOpen{}
	bgp.Init()
	// Start byte offset
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

// Instantiate an OPEN message
func MakeOpen() BgpMsgOpen {
	bgp := BgpMsgOpen{}
	bgp.Init()
	return bgp
}

// Return the total length of the OPEN Message, not including the packet header
func (bgp *BgpMsgOpen) GetLength() uint16 {
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
const VERSION_LENGTH = 1
const AUTONOMOUS_SYSTEM_LENGTH = 2

// Version //
// Used for interoperability.
type Version struct {
	fieldBase
	value uint8
}

func MakeVersion() *Version {
	f := Version{}
	f.length = VERSION_LENGTH
	// Default
	f.value = 4
	return &f
}
func (f *Version) Read(b []byte) {
	f.value = uint8(b[0])
}
func (f *Version) Value() interface{} {
	return f.value
}
func (f *Version) Serialize() []byte {
	return []byte{f.value}
}

// AS Number //
// BGP Autonomous sytem of remote router.
type AutonomousSystem struct {
	fieldBase
	value uint16
}

func MakeAutonomousSystem() *AutonomousSystem {
	f := AutonomousSystem{}
	f.length = AUTONOMOUS_SYSTEM_LENGTH
	return &f
}
func (f *AutonomousSystem) Read(b []byte) {
	l := f.GetLength()
	f.value = binary.BigEndian.Uint16(b[:l])
}
func (f *AutonomousSystem) Value() interface{} {
	return f.value
}
func (f *AutonomousSystem) Write(v uint16) {
	f.value = v
}
func (f *AutonomousSystem) Serialize() []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, f.value)
	return b
}

// AS Number //
// BGP Autonomous sytem of remote router.
type Identifier struct {
	fieldBase
	value net.IP
}

// Remote identifier of BGP session
// IP address format so it's converted to a net.IP object
func MakeIdentifier() *Identifier {
	f := Identifier{}
	f.length = 4
	return &f
}
func (f *Identifier) Read(b []byte) {
	l := f.GetLength()
	f.value = net.IP(b[:l])
}
func (f *Identifier) Value() interface{} {
	return f.value
}
func (f *Identifier) Write(v []byte) {
	f.value = v
}
func (f *Identifier) Serialize() []byte {
	return f.value
}

// Hold time
// BGP  Hold time,
type HoldTime struct {
	fieldBase
	value uint16
}

func MakeHoldTime() *HoldTime {
	f := HoldTime{}
	f.length = 2
	return &f
}
func (f *HoldTime) Read(b []byte) {
	l := f.GetLength()
	f.value = binary.BigEndian.Uint16(b[:l])
}
func (f *HoldTime) Value() interface{} {
	return f.value
}
func (f *HoldTime) Write(v uint16) {
	f.value = v
}
func (f *HoldTime) Serialize() []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, f.value)
	return b
}

// Optional Params is a struct composed of all the provided optional fields
type OptionalParams struct {
	fieldBase
	params map[uint8]OptionalParam
}

func MakeOptionalParams() *OptionalParams {
	f := OptionalParams{
		params: make(map[uint8]OptionalParam),
	}
	f.length = 0
	return &f
}
func (f *OptionalParams) Read(b []byte) {
	// Set the length to the correct value based on OptionalParamLength field
	f.length = uint16(b[0])
	offset := uint16(0)
	for offset < f.length {
		op := OptionalParam{}
		op.Read(b[offset:])
		offset = offset + uint16(op.Length)
		f.params[op.Type] = op
	}
}
func (f *OptionalParams) Write(op OptionalParam) {
	f.params[op.Type] = op
}
func (f *OptionalParams) Value() interface{} {
	return f.params
}
func (f *OptionalParams) Serialize() []byte {
	var b []byte
	for _, op := range f.params {
		b = append(b, op.Serialize()...)
	}
	return b
}

// OptionalParam represents a single paramater
// OP's are mapped to actual params such as capabilities elsewhere.
type OptionalParam struct {
	Type   uint8
	Length uint8
	Value  []byte
}

func (f *OptionalParam) Read(b []byte) {
	f.Type = b[0]
	f.Length = b[1]
	f.Value = b[1:f.Length]
}
func (f *OptionalParam) Serialize() []byte {
	b := []byte{f.Type, f.Length}
	b = append(b, f.Value...)
	return b
}
