package packet

type BgpPacket struct {
	Header BgpHeader

	Message interface{}
}
