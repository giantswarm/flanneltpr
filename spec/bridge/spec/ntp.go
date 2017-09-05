package spec

import "net"

type NTP struct {
	Servers []string `json:"servers" yaml:"servers"`
}
