package vsphere

type Platform struct {
	VCenter				string	`json:"vCenter"`
	Username			string	`json:"username"`
	Password			string	`json:"password"`
	Datacenter			string	`json:"datacenter"`
	DefaultDatastore	string	`json:"defaultDatastore"`
}
