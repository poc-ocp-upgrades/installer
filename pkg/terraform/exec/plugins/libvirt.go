package plugins

import (
	"github.com/dmacvicar/terraform-provider-libvirt/libvirt"
	"github.com/hashicorp/terraform/plugin"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	exec := func() {
		defer libvirt.CleanupLibvirtConnections()
		plugin.Serve(&plugin.ServeOpts{ProviderFunc: libvirt.Provider})
	}
	KnownPlugins["terraform-provider-libvirt"] = exec
}
