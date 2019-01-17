package bgp

import "encoding/binary"

type BgpMsgOpen struct {
	Version          *Version
	AutonomousSystem *AutonomousSystem
	fields           []Field
}

func ReadMsgOpen(b []byte) BgpMsgOpen {
	bgp := BgpMsgOpen{}
	bgp.Version = MakeVersion()
	bgp.AutonomousSystem = MakeAutonomousSystem()
	bgp.fields = []Field{
		bgp.Version,
		bgp.AutonomousSystem,
	}
	return bgp
}

func MakeDummyOpen() []byte {
	var b []byte
	fields := []Field{
		MakeVersion(),
		MakeAutonomousSystem(),
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
	return &f
}
func (f *Version) Read(b []byte) {
	f.value = uint8(b[0])
}
func (f *Version) Value() interface{} {
	return f.value
}
func (f *Version) Dummy() []byte {
	var b []byte
	b = []byte{4}
	return b
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
func (f *AutonomousSystem) Dummy() []byte {
	var b []byte
	b = []byte{2, 128}
	return b
}
