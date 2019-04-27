package plugins

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-openstack/openstack"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	exec := func() {
		plugin.Serve(&plugin.ServeOpts{ProviderFunc: openstack.Provider})
	}
	KnownPlugins["terraform-provider-openstack"] = exec
}
