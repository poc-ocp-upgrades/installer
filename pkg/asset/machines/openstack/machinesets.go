package openstack

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clusterapi "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/openstack"
	"github.com/pkg/errors"
)

func MachineSets(clusterID string, config *types.InstallConfig, pool *types.MachinePool, osImage, role, userDataSecret string) ([]*clusterapi.MachineSet, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	total := int32(0)
	if pool.Replicas != nil {
		total = int32(*pool.Replicas)
	}
	var machinesets []*clusterapi.MachineSet
	az := ""
	trunk := config.Platform.OpenStack.TrunkSupport
	provider, err := provider(clusterID, platform, mpool, osImage, az, role, userDataSecret, trunk)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create provider")
	}
	name := fmt.Sprintf("%s-%s", clusterID, pool.Name)
	mset := &clusterapi.MachineSet{TypeMeta: metav1.TypeMeta{APIVersion: "machine.openshift.io/v1beta1", Kind: "MachineSet"}, ObjectMeta: metav1.ObjectMeta{Namespace: "openshift-machine-api", Name: name, Labels: map[string]string{"machine.openshift.io/cluster-api-cluster": clusterID, "machine.openshift.io/cluster-api-machine-role": role, "machine.openshift.io/cluster-api-machine-type": role}}, Spec: clusterapi.MachineSetSpec{Replicas: &total, Selector: metav1.LabelSelector{MatchLabels: map[string]string{"machine.openshift.io/cluster-api-machineset": name, "machine.openshift.io/cluster-api-cluster": clusterID}}, Template: clusterapi.MachineTemplateSpec{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"machine.openshift.io/cluster-api-machineset": name, "machine.openshift.io/cluster-api-cluster": clusterID, "machine.openshift.io/cluster-api-machine-role": role, "machine.openshift.io/cluster-api-machine-type": role}}, Spec: clusterapi.MachineSpec{ProviderSpec: clusterapi.ProviderSpec{Value: &runtime.RawExtension{Object: provider}}}}}}
	machinesets = append(machinesets, mset)
	return machinesets, nil
}
