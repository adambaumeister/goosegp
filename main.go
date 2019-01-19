package main

import (
	"github.com/adamb/go_osegp/bgp"
	"net"
)

func main() {
	b := bgp.Parser{}

	o := bgp.MakeOpen()
	o.AutonomousSystem.Write(6262)
	o.HoldTime.Write(60)

	o.Identifier.Write(net.ParseIP("1.1.1.1").To4())
	op := bgp.OptionalParam{}
	op.Type = 2
	op.Length = 16
	op.Value = []byte{1, 4, 0, 1, 0, 1}
	o.Params.Write(op)

	h := bgp.MakeHeader()
	h.Type.Write(bgp.MESSAGE_OPEN)
	h.Length.Write(h.GetLength() + o.GetLength())

	sp := append(h.Serialize(), o.Serialize()...)
	b.Parse(sp)
}
