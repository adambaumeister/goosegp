package bgp

import (
	"net"
	"testing"
)

// Test function
// Instantiates a dummy packet, testing the write, and then parses it, testing the read.
func TestBackend(t *testing.T) {
	b := Parser{}

	// Try an OPEN packet first
	o := MakeOpen()
	o.AutonomousSystem.Write(6262)
	o.HoldTime.Write(60)

	o.Identifier.Write(net.ParseIP("1.1.1.1").To4())
	op := OptionalParam{}
	op.Type = 2
	op.Length = 16
	op.Value = []byte{1, 4, 0, 1, 0, 1}
	o.Params.Write(op)

	h := MakeHeader()
	h.Type.Write(MESSAGE_OPEN)
	h.Length.Write(h.GetLength() + o.GetLength())

	sp := append(h.Serialize(), o.Serialize()...)
	b.Parse(sp)

	// Now an UPDATE packet
}
