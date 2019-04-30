package bootkube

import (
	"os"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	etcdHostServiceEndpointsFileName = "etcd-host-service-endpoints.yaml.template"
)

var _ asset.WritableAsset = (*EtcdHostServiceEndpoints)(nil)

type EtcdHostServiceEndpoints struct{ FileList []*asset.File }

func (t *EtcdHostServiceEndpoints) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (t *EtcdHostServiceEndpoints) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "EtcdHostServiceEndpoints"
}
func (t *EtcdHostServiceEndpoints) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fileName := etcdHostServiceEndpointsFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{{Filename: filepath.Join(content.TemplateDir, fileName), Data: []byte(data)}}
	return nil
}
func (t *EtcdHostServiceEndpoints) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.FileList
}
func (t *EtcdHostServiceEndpoints) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, etcdHostServiceEndpointsFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
