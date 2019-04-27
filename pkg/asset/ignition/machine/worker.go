package machine

import (
	"encoding/json"
	"os"
	igntypes "github.com/coreos/ignition/config/v2_2/types"
	"github.com/pkg/errors"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/tls"
)

const (
	workerIgnFilename = "worker.ign"
)

type Worker struct {
	Config	*igntypes.Config
	File	*asset.File
}

var _ asset.WritableAsset = (*Worker)(nil)

func (a *Worker) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&installconfig.InstallConfig{}, &tls.RootCA{}}
}
func (a *Worker) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	installConfig := &installconfig.InstallConfig{}
	rootCA := &tls.RootCA{}
	dependencies.Get(installConfig, rootCA)
	a.Config = pointerIgnitionConfig(installConfig.Config, rootCA.Cert(), "worker")
	data, err := json.Marshal(a.Config)
	if err != nil {
		return errors.Wrap(err, "failed to marshal Ignition config")
	}
	a.File = &asset.File{Filename: workerIgnFilename, Data: data}
	return nil
}
func (a *Worker) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Worker Ignition Config"
}
func (a *Worker) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a.File != nil {
		return []*asset.File{a.File}
	}
	return []*asset.File{}
}
func (a *Worker) Load(f asset.FileFetcher) (found bool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	file, err := f.FetchByName(workerIgnFilename)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	config := &igntypes.Config{}
	if err := json.Unmarshal(file.Data, config); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal")
	}
	a.File, a.Config = file, config
	return true, nil
}
