package types

import (
	"fmt"
	"github.com/openshift/installer/pkg/ipnet"
	"github.com/openshift/installer/pkg/types/aws"
	"github.com/openshift/installer/pkg/types/azure"
	"github.com/openshift/installer/pkg/types/libvirt"
	"github.com/openshift/installer/pkg/types/none"
	"github.com/openshift/installer/pkg/types/openstack"
	"github.com/openshift/installer/pkg/types/vsphere"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	InstallConfigVersion = "v1"
)

var (
	PlatformNames		= []string{aws.Name}
	HiddenPlatformNames	= []string{none.Name, azure.Name, openstack.Name, vsphere.Name}
)

type InstallConfig struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata"`
	SSHKey			string	`json:"sshKey,omitempty"`
	BaseDomain		string	`json:"baseDomain"`
	*Networking		`json:"networking,omitempty"`
	ControlPlane		*MachinePool	`json:"controlPlane,omitempty"`
	Compute			[]MachinePool	`json:"compute,omitempty"`
	Platform		`json:"platform"`
	PullSecret		string	`json:"pullSecret"`
}

func (c *InstallConfig) ClusterDomain() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s.%s", c.ObjectMeta.Name, c.BaseDomain)
}

type Platform struct {
	AWS		*aws.Platform		`json:"aws,omitempty"`
	Libvirt		*libvirt.Platform	`json:"libvirt,omitempty"`
	None		*none.Platform		`json:"none,omitempty"`
	OpenStack	*openstack.Platform	`json:"openstack,omitempty"`
	VSphere		*vsphere.Platform	`json:"vsphere,omitempty"`
	Azure		*azure.Platform		`json:"azure,omitempty"`
}

func (p *Platform) Name() string {
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
	case p.None != nil:
		return none.Name
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

type Networking struct {
	MachineCIDR			*ipnet.IPNet		`json:"machineCIDR,omitempty"`
	NetworkType			string			`json:"networkType,omitempty"`
	ClusterNetwork			[]ClusterNetworkEntry	`json:"clusterNetwork,omitempty"`
	ServiceNetwork			[]ipnet.IPNet		`json:"serviceNetwork,omitempty"`
	DeprecatedType			string			`json:"type,omitempty"`
	DeprecatedServiceCIDR		*ipnet.IPNet		`json:"serviceCIDR,omitempty"`
	DeprecatedClusterNetworks	[]ClusterNetworkEntry	`json:"clusterNetworks,omitempty"`
}
type ClusterNetworkEntry struct {
	CIDR				ipnet.IPNet	`json:"cidr"`
	HostPrefix			int32		`json:"hostPrefix"`
	DeprecatedHostSubnetLength	int32		`json:"hostSubnetLength,omitempty"`
}
