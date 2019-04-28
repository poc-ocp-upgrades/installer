package azure

type MachinePool struct {
	Zones		[]string	`json:"zones,omitempty"`
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
