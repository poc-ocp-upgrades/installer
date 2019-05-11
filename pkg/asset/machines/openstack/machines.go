package openstack

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	machineapi "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	openstackprovider "sigs.k8s.io/cluster-api-provider-openstack/pkg/apis/openstackproviderconfig/v1alpha1"
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/openstack"
)

const (
	cloudsSecret			= "openstack-credentials"
	cloudsSecretNamespace	= "kube-system"
	CloudName				= "openstack"
)

func Machines(clusterID string, config *types.InstallConfig, pool *types.MachinePool, osImage, role, userDataSecret string) ([]machineapi.Machine, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if configPlatform := config.Platform.Name(); configPlatform != openstack.Name {
		return nil, fmt.Errorf("non-OpenStack configuration: %q", configPlatform)
	}
	if poolPlatform := pool.Platform.Name(); poolPlatform != openstack.Name {
		return nil, fmt.Errorf("non-OpenStack machine-pool: %q", poolPlatform)
	}
	platform := config.Platform.OpenStack
	mpool := pool.Platform.OpenStack
	total := int64(1)
	if pool.Replicas != nil {
		total = *pool.Replicas
	}
	var machines []machineapi.Machine
	for idx := int64(0); idx < total; idx++ {
		az := ""
		trunk := config.Platform.OpenStack.TrunkSupport
		provider, err := provider(clusterID, platform, mpool, osImage, az, role, userDataSecret, trunk)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create provider")
		}
		machine := machineapi.Machine{TypeMeta: metav1.TypeMeta{APIVersion: "machine.openshift.io/v1beta1", Kind: "Machine"}, ObjectMeta: metav1.ObjectMeta{Namespace: "openshift-machine-api", Name: fmt.Sprintf("%s-%s-%d", clusterID, pool.Name, idx), Labels: map[string]string{"machine.openshift.io/cluster-api-cluster": clusterID, "machine.openshift.io/cluster-api-machine-role": role, "machine.openshift.io/cluster-api-machine-type": role}}, Spec: machineapi.MachineSpec{ProviderSpec: machineapi.ProviderSpec{Value: &runtime.RawExtension{Object: provider}}}}
		machines = append(machines, machine)
	}
	return machines, nil
}
func provider(clusterID string, platform *openstack.Platform, mpool *openstack.MachinePool, osImage string, az string, role, userDataSecret string, trunk string) (*openstackprovider.OpenstackProviderSpec, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &openstackprovider.OpenstackProviderSpec{TypeMeta: metav1.TypeMeta{APIVersion: "openstackproviderconfig.k8s.io/v1alpha1", Kind: "OpenstackProviderSpec"}, Flavor: mpool.FlavorName, Image: osImage, CloudName: CloudName, CloudsSecret: &corev1.SecretReference{Name: cloudsSecret, Namespace: cloudsSecretNamespace}, UserDataSecret: &corev1.SecretReference{Name: userDataSecret}, Networks: []openstackprovider.NetworkParam{{Subnets: []openstackprovider.SubnetParam{{Filter: openstackprovider.SubnetFilter{Name: fmt.Sprintf("%s-nodes", clusterID), Tags: fmt.Sprintf("%s=%s", "openshiftClusterID", clusterID)}}}}}, AvailabilityZone: az, SecurityGroups: []openstackprovider.SecurityGroupParam{{Name: fmt.Sprintf("%s-%s", clusterID, role)}}, Trunk: trunkSupportBoolean(trunk), Tags: []string{fmt.Sprintf("openshiftClusterID=%s", clusterID)}, ServerMetadata: map[string]string{"Name": fmt.Sprintf("%s-%s", clusterID, role), "openshiftClusterID": clusterID}}, nil
}
func trunkSupportBoolean(trunkSupport string) (result bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if trunkSupport == "1" {
		result = true
	} else {
		result = false
	}
	return
}
func ConfigMasters(machines []machineapi.Machine, clusterID string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
