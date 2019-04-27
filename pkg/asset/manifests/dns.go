package manifests

import (
	"fmt"
	"path/filepath"
	"strings"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	configv1 "github.com/openshift/api/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
	icaws "github.com/openshift/installer/pkg/asset/installconfig/aws"
	icazure "github.com/openshift/installer/pkg/asset/installconfig/azure"
	awstypes "github.com/openshift/installer/pkg/types/aws"
	azuretypes "github.com/openshift/installer/pkg/types/azure"
	libvirttypes "github.com/openshift/installer/pkg/types/libvirt"
	nonetypes "github.com/openshift/installer/pkg/types/none"
	openstacktypes "github.com/openshift/installer/pkg/types/openstack"
	vspheretypes "github.com/openshift/installer/pkg/types/vsphere"
)

var (
	dnsCfgFilename = filepath.Join(manifestDir, "cluster-dns-02-config.yml")
)

type DNS struct{ FileList []*asset.File }

var _ asset.WritableAsset = (*DNS)(nil)

func (*DNS) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "DNS Config"
}
func (*DNS) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&installconfig.InstallConfig{}, &installconfig.ClusterID{}, &installconfig.PlatformCredsCheck{}}
}
func (d *DNS) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	installConfig := &installconfig.InstallConfig{}
	clusterID := &installconfig.ClusterID{}
	dependencies.Get(installConfig, clusterID)
	config := &configv1.DNS{TypeMeta: metav1.TypeMeta{APIVersion: configv1.SchemeGroupVersion.String(), Kind: "DNS"}, ObjectMeta: metav1.ObjectMeta{Name: "cluster"}, Spec: configv1.DNSSpec{BaseDomain: installConfig.Config.ClusterDomain()}}
	switch installConfig.Config.Platform.Name() {
	case awstypes.Name:
		zone, err := icaws.GetPublicZone(installConfig.Config.BaseDomain)
		if err != nil {
			return errors.Wrapf(err, "getting public zone for %q", installConfig.Config.BaseDomain)
		}
		config.Spec.PublicZone = &configv1.DNSZone{ID: strings.TrimPrefix(*zone.Id, "/hostedzone/")}
		config.Spec.PrivateZone = &configv1.DNSZone{Tags: map[string]string{fmt.Sprintf("kubernetes.io/cluster/%s", clusterID.InfraID): "owned", "Name": fmt.Sprintf("%s-int", clusterID.InfraID)}}
	case azuretypes.Name:
		dnsConfig, err := icazure.NewDNSConfig()
		if err != nil {
			return err
		}
		config.Spec.PublicZone = &configv1.DNSZone{ID: dnsConfig.GetDNSZoneID(installConfig.Config.Azure.BaseDomainResourceGroupName, installConfig.Config.BaseDomain)}
		config.Spec.PrivateZone = &configv1.DNSZone{ID: dnsConfig.GetDNSZoneID(clusterID.InfraID+"-rg", installConfig.Config.ClusterDomain())}
	case libvirttypes.Name, openstacktypes.Name, nonetypes.Name, vspheretypes.Name:
	default:
		return errors.New("invalid Platform")
	}
	configData, err := yaml.Marshal(config)
	if err != nil {
		return errors.Wrapf(err, "failed to create %s manifests from InstallConfig", d.Name())
	}
	d.FileList = []*asset.File{{Filename: dnsCfgFilename, Data: configData}}
	return nil
}
func (d *DNS) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return d.FileList
}
func (d *DNS) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false, nil
}
