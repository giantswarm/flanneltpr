package spec

import "net"

type NTP struct {
	Servers []net.IP `json:"servers" yaml:"servers"`
}
