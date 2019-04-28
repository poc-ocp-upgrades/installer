package openstack

type Metadata struct {
	Region		string			`json:"region"`
	Cloud		string			`json:"cloud"`
	Identifier	map[string]string	`json:"identifier"`
}
