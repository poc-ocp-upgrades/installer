package plugins

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	azurermProvider := func() {
		plugin.Serve(&plugin.ServeOpts{ProviderFunc: azurerm.Provider})
	}
	KnownPlugins["terraform-provider-azurerm"] = azurermProvider
}
