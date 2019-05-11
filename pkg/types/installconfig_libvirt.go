package types

import (
	"sort"
	"github.com/openshift/installer/pkg/types/libvirt"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	PlatformNames = append(PlatformNames, libvirt.Name)
	sort.Strings(PlatformNames)
}
