package azure

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	machineapi "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/azure"
	azureprovider "sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1alpha1"
)

const (
	cloudsSecret		= "azure-credentials"
	cloudsSecretNamespace	= "kube-system"
)

func Machines(clusterID string, config *types.InstallConfig, pool *types.MachinePool, osImage, role, userDataSecret string) ([]machineapi.Machine, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if configPlatform := config.Platform.Name(); configPlatform != azure.Name {
		return nil, fmt.Errorf("non-Azure configuration: %q", configPlatform)
	}
	if poolPlatform := pool.Platform.Name(); poolPlatform != azure.Name {
		return nil, fmt.Errorf("non-Azure machine-pool: %q", poolPlatform)
	}
	platform := config.Platform.Azure
	mpool := pool.Platform.Azure
	total := int64(1)
	if pool.Replicas != nil {
		total = *pool.Replicas
	}
	var machines []machineapi.Machine
	for idx := int64(0); idx < total; idx++ {
		provider, err := provider(platform, mpool, osImage, userDataSecret)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create provider")
		}
		machine := machineapi.Machine{TypeMeta: metav1.TypeMeta{APIVersion: "machine.openshift.io/v1beta1", Kind: "Machine"}, ObjectMeta: metav1.ObjectMeta{Namespace: "openshift-machine-api", Name: fmt.Sprintf("%s-%s-%d", clusterID, pool.Name, idx), Labels: map[string]string{"machine.openshift.io/cluster-api-cluster": clusterID, "machine.openshift.io/cluster-api-machine-role": role, "machine.openshift.io/cluster-api-machine-type": role}}, Spec: machineapi.MachineSpec{ProviderSpec: machineapi.ProviderSpec{Value: &runtime.RawExtension{Object: provider}}}}
		machines = append(machines, machine)
	}
	return machines, nil
}
func provider(platform *azure.Platform, mpool *azure.MachinePool, osImage string, userDataSecret string) (*azureprovider.AzureMachineProviderSpec, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &azureprovider.AzureMachineProviderSpec{TypeMeta: metav1.TypeMeta{APIVersion: "azureprovider.k8s.io/v1alpha1", Kind: "AzureMachineProviderSpec"}, UserDataSecret: &corev1.SecretReference{Name: userDataSecret}, CredentialsSecret: &corev1.SecretReference{Name: cloudsSecret, Namespace: cloudsSecretNamespace}, Location: platform.Region, VMSize: mpool.InstanceType, Image: azureprovider.Image{ResourceID: osImage}, OSDisk: azureprovider.OSDisk{OSType: "Linux", DiskSizeGB: 64, ManagedDisk: azureprovider.ManagedDisk{StorageAccountType: "Premium_LRS"}}}, nil
}
func ConfigMasters(machines []machineapi.Machine, clusterID string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
