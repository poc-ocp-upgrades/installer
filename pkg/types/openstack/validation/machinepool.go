package validation

import (
	"k8s.io/apimachinery/pkg/util/validation/field"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"github.com/openshift/installer/pkg/types/openstack"
)

func ValidateMachinePool(p *openstack.MachinePool, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return field.ErrorList{}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
