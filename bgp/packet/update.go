package packet

import "net"

// Routes removed from this BGP relationship
type WithdrawnRoutes struct {
	Length uint16
	Routes []WithdrawnRoute
}
type WithdrawnRoute struct {
	Length uint8
	Prefix net.IP
}

// BGP Path attributes (AS-PATH, Origin, MED, etc.)
type PathAttributes struct {
	Length     uint16
	Attributes []PathAttribute
}

type PathAttribute struct {
}
