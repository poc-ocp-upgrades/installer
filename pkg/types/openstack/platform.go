package openstack

type Platform struct {
	Region					string			`json:"region"`
	DefaultMachinePlatform	*MachinePool	`json:"defaultMachinePlatform,omitempty"`
	Cloud					string			`json:"cloud"`
	ExternalNetwork			string			`json:"externalNetwork"`
	FlavorName				string			`json:"computeFlavor"`
	LbFloatingIP			string			`json:"lbFloatingIP"`
	TrunkSupport			string			`json:"trunkSupport"`
}
