package bgp

import (
	"fmt"
	"github.com/adamb/go_osegp/bgp/config"
	"github.com/adamb/go_osegp/bgp/errors"
	"github.com/adamb/go_osegp/bgp/packet"
	"net"
)

/*
Session implements the connection handling required to set up and maintain a BGP sesion.

It uses the Packet library to parse incoming packets.

Session should not communicate with the RIB directly but rather relay information back to a "Router" object

Router -> session -> packet
*/

type Session struct {
	Config *config.Config

	PacketIn chan packet.BgpPacket
}

func (s *Session) StartSessionListener() {
	l, err := net.Listen("tcp", s.Config.Router.Address.String()+":179")
	errors.CheckError(err)
	for {
		conn, err := l.Accept()
		errors.CheckError(err)
		fmt.Printf("Received conn from %v\n", conn.RemoteAddr())
		b := packet.BgpPacket{}
		s.PacketIn <- b
	}
}

func SessionInit(c config.Config) {
	s := Session{
		Config: &c,
	}
	go s.StartSessionListener()
	_ = <-s.PacketIn
}
