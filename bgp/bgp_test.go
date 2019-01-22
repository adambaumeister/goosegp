package bgp

import (
	"github.com/adamb/go_osegp/bgp/config"
	"github.com/adamb/go_osegp/bgp/packet"
	"testing"
)

// Test BGP session objects
func TestSession(t *testing.T) {
	c := config.GetConfig()
	// Start a session listener, and thread it.
	s := Session{
		Config:   &c,
		PacketIn: make(chan packet.BgpPacket),
	}
	go s.StartSessionListener()

	// Init a sender, to validate it works in that direction

	// Try sending.
	s.StartSessionSender()

}
