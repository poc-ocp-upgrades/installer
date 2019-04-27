package defaults

import (
	"github.com/openshift/installer/pkg/types/libvirt"
)

const (
	DefaultURI = "qemu+tcp://192.168.122.1/system"
)

func SetPlatformDefaults(p *libvirt.Platform) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if p.URI == "" {
		p.URI = DefaultURI
	}
	if p.Network == nil {
		p.Network = &libvirt.Network{}
	}
	SetNetworkDefaults(p.Network)
}
