package bgp

import (
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
	conn   net.TCPConn
	Parser packet.Parser
}
