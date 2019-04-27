package plugins

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-ignition/ignition"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	exec := func() {
		plugin.Serve(&plugin.ServeOpts{ProviderFunc: ignition.Provider})
	}
	KnownPlugins["terraform-provider-ignition"] = exec
}
