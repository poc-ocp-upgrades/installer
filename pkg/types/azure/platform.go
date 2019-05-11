package azure

import "strings"

type Platform struct {
	Region						string				`json:"region"`
	BaseDomainResourceGroupName	string				`json:"baseDomainResourceGroupName,omitempty"`
	UserTags					map[string]string	`json:"userTags,omitempty"`
	DefaultMachinePlatform		*MachinePool		`json:"defaultMachinePlatform,omitempty"`
}

func (p *Platform) SetBaseDomain(baseDomainID string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	parts := strings.Split(baseDomainID, "/")
	p.BaseDomainResourceGroupName = parts[4]
	return nil
}
