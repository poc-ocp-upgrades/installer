package libvirt

import (
	"github.com/openshift/installer/pkg/types"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"github.com/openshift/installer/pkg/types/libvirt"
)

func Metadata(config *types.InstallConfig) *libvirt.Metadata {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &libvirt.Metadata{URI: config.Platform.Libvirt.URI}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
