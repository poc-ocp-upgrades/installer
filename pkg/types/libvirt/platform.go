package libvirt

type Platform struct {
	URI						string			`json:"URI,omitempty"`
	DefaultMachinePlatform	*MachinePool	`json:"defaultMachinePlatform,omitempty"`
	Network					*Network		`json:"network,omitempty"`
}
