package main

import "github.com/adamb/go_osegp/bgp"

func main() {
	b := bgp.Parser{}
	dp := append(
		bgp.MakeDummyHeader(),
		bgp.MakeDummyOpen()...,
	)
	b.Parse(dp)
}
