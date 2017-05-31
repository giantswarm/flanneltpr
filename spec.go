package flanneltpr

import (
	"github.com/giantswarm/flanneltpr/host"
	"github.com/giantswarm/flanneltpr/network"
)

type Spec struct {
	Host    host.Host       `json:"host" yaml:"host"`
	Network network.Network `json:"network" yaml:"network"`
}
