package vsphere

type MachinePool struct{}

func (p *MachinePool) Set(required *MachinePool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if required == nil || p == nil {
		return
	}
}
