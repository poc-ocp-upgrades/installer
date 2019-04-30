package aws

type MachinePool struct {
	Zones		[]string	`json:"zones,omitempty"`
	InstanceType	string		`json:"type"`
	EC2RootVolume	`json:"rootVolume"`
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
	if required.EC2RootVolume.IOPS != 0 {
		a.EC2RootVolume.IOPS = required.EC2RootVolume.IOPS
	}
	if required.EC2RootVolume.Size != 0 {
		a.EC2RootVolume.Size = required.EC2RootVolume.Size
	}
	if required.EC2RootVolume.Type != "" {
		a.EC2RootVolume.Type = required.EC2RootVolume.Type
	}
}

type EC2RootVolume struct {
	IOPS	int	`json:"iops"`
	Size	int	`json:"size"`
	Type	string	`json:"type"`
}
