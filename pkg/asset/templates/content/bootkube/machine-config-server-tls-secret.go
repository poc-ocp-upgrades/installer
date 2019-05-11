package bootkube

import (
	"os"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	machineConfigServerTLSSecretFileName = "machine-config-server-tls-secret.yaml.template"
)

var _ asset.WritableAsset = (*MachineConfigServerTLSSecret)(nil)

type MachineConfigServerTLSSecret struct{ FileList []*asset.File }

func (t *MachineConfigServerTLSSecret) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (t *MachineConfigServerTLSSecret) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "MachineConfigServerTLSSecret"
}
func (t *MachineConfigServerTLSSecret) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fileName := machineConfigServerTLSSecretFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{{Filename: filepath.Join(content.TemplateDir, fileName), Data: []byte(data)}}
	return nil
}
func (t *MachineConfigServerTLSSecret) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.FileList
}
func (t *MachineConfigServerTLSSecret) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, machineConfigServerTLSSecretFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
