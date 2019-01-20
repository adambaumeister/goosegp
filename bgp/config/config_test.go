package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	c := GetConfig()
	fmt.Printf("Configured: %v\n", c.Router.Address.String())
	return
}
