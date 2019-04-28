package aws

type Metadata struct {
	Region		string			`json:"region"`
	Identifier	[]map[string]string	`json:"identifier"`
}
