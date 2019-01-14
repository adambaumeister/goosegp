package main

import "github.com/adamb/go_osegp/bgp"

func main() {
	b := bgp.Parser{}
	b.Parse([]byte{255})
}
