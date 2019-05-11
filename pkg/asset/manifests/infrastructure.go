package manifests

import (
	"path/filepath"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	configv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/openshift/installer/pkg/types/aws"
	"github.com/openshift/installer/pkg/types/azure"
	"github.com/openshift/installer/pkg/types/libvirt"
	"github.com/openshift/installer/pkg/types/none"
	"github.com/openshift/installer/pkg/types/openstack"
	"github.com/openshift/installer/pkg/types/vsphere"
)

var (
	infraCrdFilename	= filepath.Join(manifestDir, "cluster-infrastructure-01-crd.yaml")
	infraCfgFilename	= filepath.Join(manifestDir, "cluster-infrastructure-02-config.yml")
)

type Infrastructure struct{ FileList []*asset.File }

var _ asset.WritableAsset = (*Infrastructure)(nil)

func (*Infrastructure) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Infrastructure Config"
}
func (*Infrastructure) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&installconfig.ClusterID{}, &installconfig.InstallConfig{}, &CloudProviderConfig{}}
}
func (i *Infrastructure) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clusterID := &installconfig.ClusterID{}
	installConfig := &installconfig.InstallConfig{}
	cloudproviderconfig := &CloudProviderConfig{}
	dependencies.Get(clusterID, installConfig, cloudproviderconfig)
	var platform configv1.PlatformType
	switch installConfig.Config.Platform.Name() {
	case aws.Name:
		platform = configv1.AWSPlatformType
	case none.Name:
		platform = configv1.NonePlatformType
	case libvirt.Name:
		platform = configv1.LibvirtPlatformType
	case openstack.Name:
		platform = configv1.OpenStackPlatformType
	case vsphere.Name:
		platform = configv1.VSpherePlatformType
	case azure.Name:
		platform = configv1.AzurePlatformType
	default:
		platform = configv1.NonePlatformType
	}
	config := &configv1.Infrastructure{TypeMeta: metav1.TypeMeta{APIVersion: configv1.SchemeGroupVersion.String(), Kind: "Infrastructure"}, ObjectMeta: metav1.ObjectMeta{Name: "cluster"}, Status: configv1.InfrastructureStatus{InfrastructureName: clusterID.InfraID, Platform: platform, APIServerURL: getAPIServerURL(installConfig.Config), EtcdDiscoveryDomain: getEtcdDiscoveryDomain(installConfig.Config)}}
	if cloudproviderconfig.ConfigMap != nil {
		config.Spec.CloudConfig = configv1.ConfigMapFileReference{Name: cloudproviderconfig.ConfigMap.Name, Key: cloudProviderConfigDataKey}
		i.FileList = append(i.FileList, cloudproviderconfig.File)
	}
	configData, err := yaml.Marshal(config)
	if err != nil {
		return errors.Wrapf(err, "failed to marshal config: %#v", config)
	}
	i.FileList = append(i.FileList, &asset.File{Filename: infraCfgFilename, Data: configData})
	return nil
}
func (i *Infrastructure) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return i.FileList
}
func (i *Infrastructure) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false, nil
}
