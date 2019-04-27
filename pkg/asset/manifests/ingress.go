package manifests

import (
	"fmt"
	"path/filepath"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	configv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	ingCfgFilename = filepath.Join(manifestDir, "cluster-ingress-02-config.yml")
)

type Ingress struct{ FileList []*asset.File }

var _ asset.WritableAsset = (*Ingress)(nil)

func (*Ingress) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Ingress Config"
}
func (*Ingress) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&installconfig.InstallConfig{}}
}
func (ing *Ingress) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	installConfig := &installconfig.InstallConfig{}
	dependencies.Get(installConfig)
	config := &configv1.Ingress{TypeMeta: metav1.TypeMeta{APIVersion: configv1.SchemeGroupVersion.String(), Kind: "Ingress"}, ObjectMeta: metav1.ObjectMeta{Name: "cluster"}, Spec: configv1.IngressSpec{Domain: fmt.Sprintf("apps.%s", installConfig.Config.ClusterDomain())}}
	configData, err := yaml.Marshal(config)
	if err != nil {
		return errors.Wrapf(err, "failed to create %s manifests from InstallConfig", ing.Name())
	}
	ing.FileList = []*asset.File{{Filename: ingCfgFilename, Data: configData}}
	return nil
}
func (ing *Ingress) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ing.FileList
}
func (ing *Ingress) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false, nil
}
