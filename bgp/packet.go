package bgp

type BgpPacket struct {
	Header BgpHeader

	Message interface{}
}
