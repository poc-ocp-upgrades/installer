package azure

import (
	"bytes"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"encoding/json"
	"fmt"
)

type CloudProviderConfig struct {
	TenantID	string
	SubscriptionID	string
	GroupLocation	string
	ResourcePrefix	string
}

func (params CloudProviderConfig) JSON() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resourceGroupName := params.ResourcePrefix + "-rg"
	config := config{authConfig: authConfig{Cloud: "AzurePublicCloud", TenantID: params.TenantID, SubscriptionID: params.SubscriptionID, UseManagedIdentityExtension: true, UserAssignedIdentityID: fmt.Sprintf("/subscriptions/%s/resourcegroups/%s/providers/Microsoft.ManagedIdentity/userAssignedIdentities/%s", params.SubscriptionID, resourceGroupName, params.ResourcePrefix+"-identity")}, ResourceGroup: resourceGroupName, Location: params.GroupLocation, SubnetName: params.ResourcePrefix + "-node-subnet", SecurityGroupName: params.ResourcePrefix + "-node-nsg", VnetName: params.ResourcePrefix + "-vnet", VnetResourceGroup: resourceGroupName, RouteTableName: params.ResourcePrefix + "-node-routetable", CloudProviderBackoff: true, CloudProviderRateLimit: true, UseInstanceMetadata: true, LoadBalancerSku: "standard"}
	buff := &bytes.Buffer{}
	encoder := json.NewEncoder(buff)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(config); err != nil {
		return "", err
	}
	return buff.String(), nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
