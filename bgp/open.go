package bgp

import "encoding/binary"

type BgpMsgOpen struct {
	Version          *Version
	AutonomousSystem *AutonomousSystem
	fields           []Field

	Length uint16
}

func (bgp *BgpMsgOpen) Init() {
	bgp.Version = MakeVersion()
	bgp.AutonomousSystem = MakeAutonomousSystem()
	bgp.fields = []Field{
		bgp.Version,
		bgp.AutonomousSystem,
	}
}

func ReadMsgOpen(b []byte) BgpMsgOpen {
	bgp := BgpMsgOpen{}
	bgp.Init()
	return bgp
}

func MakeOpen() BgpMsgOpen {
	bgp := BgpMsgOpen{}
	bgp.Init()
	return bgp
}

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
