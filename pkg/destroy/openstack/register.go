package openstack

import (
	"github.com/openshift/installer/pkg/destroy"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	destroy.Registry["openstack"] = New
}
