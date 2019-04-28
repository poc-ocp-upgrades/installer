package plugins

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-local/local"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	localProvider := func() {
		plugin.Serve(&plugin.ServeOpts{ProviderFunc: local.Provider})
	}
	KnownPlugins["terraform-provider-local"] = localProvider
}
