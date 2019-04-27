package bootkube

import (
	"os"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	hostEtcdServiceEndpointsKubeSystemFileName = "host-etcd-service-endpoints.yaml.template"
)

var _ asset.WritableAsset = (*HostEtcdServiceEndpointsKubeSystem)(nil)

type HostEtcdServiceEndpointsKubeSystem struct{ FileList []*asset.File }

func (t *HostEtcdServiceEndpointsKubeSystem) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (t *HostEtcdServiceEndpointsKubeSystem) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "HostEtcdServiceEndpointsKubeSystem"
}
func (t *HostEtcdServiceEndpointsKubeSystem) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fileName := hostEtcdServiceEndpointsKubeSystemFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{{Filename: filepath.Join(content.TemplateDir, fileName), Data: []byte(data)}}
	return nil
}
func (t *HostEtcdServiceEndpointsKubeSystem) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.FileList
}
func (t *HostEtcdServiceEndpointsKubeSystem) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, hostEtcdServiceEndpointsKubeSystemFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
