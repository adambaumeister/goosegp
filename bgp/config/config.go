package config

import (
	"github.com/adamb/go_osegp/bgp/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net"
	"os"
	"time"
)

const DEFAULT_HOLDTIME = 60

type Config struct {
	Router    Router
	Neighbors []Neighbor

	ConnTimeout time.Duration
}

type Router struct {
	AS       uint16
	HoldTime uint16
	Address  net.IP
}

type Neighbor struct {
	Remote Router
}

func GetFromEnv(key string) (string, bool) {
	if len(os.Getenv(key)) > 0 {
		return os.Getenv(key), true
	} else {
		return "", false
	}
}

func GetConfig() Config {
	s, result := GetFromEnv("GOOSE_CONFIG")
	if result {
		return Read(s)
	}
	errors.RaiseError("GOOSE_STARTUP: Invalid or missing configuration file path.")
	// Is there a better way of doing this?
	panic("")
}

// REad a yaml configuration
func Read(p string) Config {
	// Setup defaults
	c := Config{
		Router: Router{
			HoldTime: DEFAULT_HOLDTIME,
		},
	}
	data, err := ioutil.ReadFile(p)
	errors.CheckError(err)

	err = yaml.Unmarshal(data, &c)
	errors.CheckError(err)

	return c
}
