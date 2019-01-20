package bgp

import (
	"github.com/adamb/go_osegp/bgp/config"
	"testing"
)

func TestSession(t *testing.T) {
	c := config.GetConfig()
	SessionInit(c)
}
