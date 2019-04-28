package defaults

import (
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/libvirt"
)

func SetMachinePoolDefaults(p *types.MachinePool, platform string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defaultReplicaCount := int64(3)
	if platform == libvirt.Name {
		defaultReplicaCount = 1
	}
	if p.Replicas == nil {
		p.Replicas = &defaultReplicaCount
	}
	if p.Hyperthreading == "" {
		p.Hyperthreading = types.HyperthreadingEnabled
	}
}
