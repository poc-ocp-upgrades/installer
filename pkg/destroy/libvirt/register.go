package libvirt

import (
	"github.com/openshift/installer/pkg/destroy"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	destroy.Registry["libvirt"] = New
}
