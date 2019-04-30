package bootkube

import (
	"os"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	etcdCAConfigMapFileName = "etcd-ca-bundle-configmap.yaml.template"
)

var _ asset.WritableAsset = (*EtcdCAConfigMap)(nil)

type EtcdCAConfigMap struct{ FileList []*asset.File }

func (t *EtcdCAConfigMap) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (t *EtcdCAConfigMap) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "EtcdCAConfigMap"
}
func (t *EtcdCAConfigMap) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fileName := etcdCAConfigMapFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{{Filename: filepath.Join(content.TemplateDir, fileName), Data: []byte(data)}}
	return nil
}
func (t *EtcdCAConfigMap) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.FileList
}
func (t *EtcdCAConfigMap) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, etcdCAConfigMapFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
