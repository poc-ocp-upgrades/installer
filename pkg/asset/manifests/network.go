package manifests

import (
	"fmt"
	"path/filepath"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	configv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/templates/content/openshift"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	noCrdFilename	= filepath.Join(manifestDir, "cluster-network-01-crd.yml")
	noCfgFilename	= filepath.Join(manifestDir, "cluster-network-02-config.yml")
)

type Networking struct {
	Config		*configv1.Network
	FileList	[]*asset.File
}

var _ asset.WritableAsset = (*Networking)(nil)

func (no *Networking) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Network Config"
}
func (no *Networking) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&installconfig.InstallConfig{}, &openshift.NetworkCRDs{}}
}
func (no *Networking) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	installConfig := &installconfig.InstallConfig{}
	crds := &openshift.NetworkCRDs{}
	dependencies.Get(installConfig, crds)
	netConfig := installConfig.Config.Networking
	clusterNet := []configv1.ClusterNetworkEntry{}
	if len(netConfig.ClusterNetwork) > 0 {
		for _, net := range netConfig.ClusterNetwork {
			clusterNet = append(clusterNet, configv1.ClusterNetworkEntry{CIDR: net.CIDR.String(), HostPrefix: uint32(net.HostPrefix)})
		}
	} else {
		return errors.Errorf("ClusterNetworks must be specified")
	}
	serviceNet := []string{}
	for _, sn := range netConfig.ServiceNetwork {
		serviceNet = append(serviceNet, sn.String())
	}
	no.Config = &configv1.Network{TypeMeta: metav1.TypeMeta{APIVersion: configv1.SchemeGroupVersion.String(), Kind: "Network"}, ObjectMeta: metav1.ObjectMeta{Name: "cluster"}, Spec: configv1.NetworkSpec{ClusterNetwork: clusterNet, ServiceNetwork: serviceNet, NetworkType: netConfig.NetworkType}}
	configData, err := yaml.Marshal(no.Config)
	if err != nil {
		return errors.Wrapf(err, "failed to create %s manifests from InstallConfig", no.Name())
	}
	crdContents := ""
	for _, crdFile := range crds.Files() {
		crdContents = fmt.Sprintf("%s\n---\n%s", crdContents, crdFile.Data)
	}
	no.FileList = []*asset.File{{Filename: noCrdFilename, Data: []byte(crdContents)}, {Filename: noCfgFilename, Data: configData}}
	return nil
}
func (no *Networking) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return no.FileList
}
func (no *Networking) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false, nil
}
