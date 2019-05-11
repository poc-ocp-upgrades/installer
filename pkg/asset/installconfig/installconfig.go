package installconfig

import (
	"os"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/conversion"
	"github.com/openshift/installer/pkg/types/defaults"
	openstackvalidation "github.com/openshift/installer/pkg/types/openstack/validation"
	"github.com/openshift/installer/pkg/types/validation"
)

const (
	installConfigFilename = "install-config.yaml"
)

type InstallConfig struct {
	Config	*types.InstallConfig	`json:"config"`
	File	*asset.File				`json:"file"`
}

var _ asset.WritableAsset = (*InstallConfig)(nil)

func (a *InstallConfig) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&sshPublicKey{}, &baseDomain{}, &clusterName{}, &pullSecret{}, &platform{}}
}
func (a *InstallConfig) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sshPublicKey := &sshPublicKey{}
	baseDomain := &baseDomain{}
	clusterName := &clusterName{}
	pullSecret := &pullSecret{}
	platform := &platform{}
	parents.Get(sshPublicKey, baseDomain, clusterName, pullSecret, platform)
	a.Config = &types.InstallConfig{TypeMeta: metav1.TypeMeta{APIVersion: types.InstallConfigVersion}, ObjectMeta: metav1.ObjectMeta{Name: clusterName.ClusterName}, SSHKey: sshPublicKey.Key, BaseDomain: baseDomain.BaseDomain, PullSecret: pullSecret.PullSecret}
	a.Config.AWS = platform.AWS
	a.Config.Libvirt = platform.Libvirt
	a.Config.None = platform.None
	a.Config.OpenStack = platform.OpenStack
	a.Config.VSphere = platform.VSphere
	a.Config.Azure = platform.Azure
	if err := a.setDefaults(); err != nil {
		return errors.Wrap(err, "failed to set defaults for install config")
	}
	if err := validation.ValidateInstallConfig(a.Config, openstackvalidation.NewValidValuesFetcher()).ToAggregate(); err != nil {
		return errors.Wrap(err, "invalid install config")
	}
	data, err := yaml.Marshal(a.Config)
	if err != nil {
		return errors.Wrap(err, "failed to Marshal InstallConfig")
	}
	a.File = &asset.File{Filename: installConfigFilename, Data: data}
	return nil
}
func (a *InstallConfig) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Install Config"
}
func (a *InstallConfig) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a.File != nil {
		return []*asset.File{a.File}
	}
	return []*asset.File{}
}
func (a *InstallConfig) Load(f asset.FileFetcher) (found bool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	file, err := f.FetchByName(installConfigFilename)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	config := &types.InstallConfig{}
	if err := yaml.Unmarshal(file.Data, config); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal")
	}
	a.Config = config
	if err := a.convert(); err != nil {
		return false, errors.Wrap(err, "failed to upconvert install config")
	}
	if err := a.setDefaults(); err != nil {
		return false, errors.Wrap(err, "failed to set defaults for install config")
	}
	if err := validation.ValidateInstallConfig(a.Config, openstackvalidation.NewValidValuesFetcher()).ToAggregate(); err != nil {
		return false, errors.Wrapf(err, "invalid %q file", installConfigFilename)
	}
	data, err := yaml.Marshal(a.Config)
	if err != nil {
		return false, errors.Wrap(err, "failed to Marshal InstallConfig")
	}
	a.File = &asset.File{Filename: installConfigFilename, Data: data}
	return true, nil
}
func (a *InstallConfig) setDefaults() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defaults.SetInstallConfigDefaults(a.Config)
	return nil
}
func (a *InstallConfig) convert() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return conversion.ConvertInstallConfig(a.Config)
}
