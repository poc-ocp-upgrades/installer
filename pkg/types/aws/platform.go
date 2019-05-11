package aws

type Platform struct {
	Region					string				`json:"region"`
	UserTags				map[string]string	`json:"userTags,omitempty"`
	DefaultMachinePlatform	*MachinePool		`json:"defaultMachinePlatform,omitempty"`
}
