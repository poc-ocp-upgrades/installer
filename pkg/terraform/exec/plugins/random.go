package plugins

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-random/random"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	randomProvider := func() {
		plugin.Serve(&plugin.ServeOpts{ProviderFunc: random.Provider})
	}
	KnownPlugins["terraform-provider-random"] = randomProvider
}
