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

	PacketIn  chan packet.BgpPacket
	PacketOut chan packet.BgpPacket
	Signal    chan Signal
}

// Craft an OPEN packet based upon the configuration.
func (s *Session) CraftOpen() packet.BgpPacket {
	o := packet.MakeOpen()
	o.AutonomousSystem.Write(s.Config.Router.HoldTime)
	o.HoldTime.Write(s.Config.Router.HoldTime)

	o.Identifier.Write(s.Config.Router.Address.To4())

	h := packet.MakeHeader()
	h.Type.Write(packet.MESSAGE_OPEN)
	h.Length.Write(h.GetLength() + o.GetLength())

	bgp := packet.BgpPacket{
		Header:  h,
		Message: o,
	}

	return bgp
}

// StartSessionLister begins listening for incoming BGP sessions.
func (s *Session) StartSessionListener() {
	l, err := net.Listen("tcp", s.Config.Router.Address.String()+":179")
	errors.CheckError(err)
	for {
		conn, err := l.Accept()
		errors.CheckError(err)

		bgp := packet.ParseFromConn(conn)
		fmt.Printf("Got a packet with type %v\n", bgp.Header.Type.Value())
		s.PacketIn <- bgp
	}
}

// StartSessionSender attempts to initiate a TCP connection to a remote address
func (s *Session) StartSessionSender() {
	for _, neighbor := range s.Config.Neighbors {
		conn, err := net.DialTimeout("tcp", neighbor.Remote.Address.String()+":179", s.Config.ConnTimeout)
		errors.CheckError(err)

		open := s.CraftOpen()
		fmt.Printf("DEBUG: %v\n", open.Serialize())
		conn.Write(open.Serialize())
	}

}

// SessionInit begins a BGP session.
// It starts a listener, on port 179, and attempts to peer to all configured neighbors.
func SessionInit(c config.Config) {
	s := Session{
		Config:   &c,
		PacketIn: make(chan packet.BgpPacket),
	}
	go s.StartSessionListener()
	for {
		select {
		case bgp := <-s.PacketIn:
			fmt.Printf("Got a packet with type %v\n", bgp.Header.Type.Value())
		case s := <-s.Signal:
			fmt.Printf("Recieved signal %v. Exiting.", s.SignalType)
		}
	}
}
