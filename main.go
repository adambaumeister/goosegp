package main

import (
	"github.com/adamb/go_osegp/bgp"
)

func main() {
	b := bgp.Parser{}

	o := bgp.MakeOpen()
	o.AutonomousSystem.Write(6262)

	h := bgp.MakeHeader()
	h.Type.Write(bgp.MESSAGE_OPEN)
	h.Length.Write(h.GetLength() + o.GetLength())

	sp := append(h.Serialize(), o.Serialize()...)
	b.Parse(sp)
}
