package types

import (
	"github.com/openshift/installer/pkg/types/aws"
	"github.com/openshift/installer/pkg/types/azure"
	"github.com/openshift/installer/pkg/types/libvirt"
	"github.com/openshift/installer/pkg/types/openstack"
	"github.com/openshift/installer/pkg/types/vsphere"
)

type HyperthreadingMode string

const (
	HyperthreadingEnabled	HyperthreadingMode	= "Enabled"
	HyperthreadingDisabled	HyperthreadingMode	= "Disabled"
)

type MachinePool struct {
	Name		string			`json:"name"`
	Replicas	*int64			`json:"replicas,omitempty"`
	Platform	MachinePoolPlatform	`json:"platform"`
	Hyperthreading	HyperthreadingMode	`json:"hyperthreading,omitempty"`
}
type MachinePoolPlatform struct {
	AWS		*aws.MachinePool	`json:"aws,omitempty"`
	Libvirt		*libvirt.MachinePool	`json:"libvirt,omitempty"`
	OpenStack	*openstack.MachinePool	`json:"openstack,omitempty"`
	VSphere		*vsphere.MachinePool	`json:"vsphere,omitempty"`
	Azure		*azure.MachinePool	`json:"azure,omitempty"`
}

func (p *MachinePoolPlatform) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case p == nil:
		return ""
	case p.AWS != nil:
		return aws.Name
	case p.Libvirt != nil:
		return libvirt.Name
	case p.OpenStack != nil:
		return openstack.Name
	case p.VSphere != nil:
		return vsphere.Name
	case p.Azure != nil:
		return azure.Name
	default:
		return ""
	}
}
