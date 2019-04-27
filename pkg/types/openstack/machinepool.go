package openstack

type MachinePool struct {
	FlavorName string `json:"type"`
}

func (o *MachinePool) Set(required *MachinePool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if required == nil || o == nil {
		return
	}
	if required.FlavorName != "" {
		o.FlavorName = required.FlavorName
	}
}
