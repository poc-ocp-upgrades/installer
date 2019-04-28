package bootkube

import (
	"os"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	etcdServingCAConfigMapFileName = "etcd-serving-ca-configmap.yaml.template"
)

var etcdServingCAFiles = []string{etcdServingCAConfigMapFileName}
var _ asset.WritableAsset = (*EtcdServingCAConfigMap)(nil)

type EtcdServingCAConfigMap struct{ FileList []*asset.File }

func (t *EtcdServingCAConfigMap) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (t *EtcdServingCAConfigMap) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "EtcdServingCAConfigMap"
}
func (t *EtcdServingCAConfigMap) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.FileList = []*asset.File{}
	for _, fileName := range etcdServingCAFiles {
		data, err := content.GetBootkubeTemplate(fileName)
		if err != nil {
			return err
		}
		t.FileList = append(t.FileList, &asset.File{Filename: filepath.Join(content.TemplateDir, fileName), Data: []byte(data)})
	}
	return nil
}
func (t *EtcdServingCAConfigMap) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.FileList
}
func (t *EtcdServingCAConfigMap) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.FileList = []*asset.File{}
	for _, fileName := range etcdServingCAFiles {
		file, err := f.FetchByName(filepath.Join(content.TemplateDir, fileName))
		if err != nil {
			if os.IsNotExist(err) {
				return false, nil
			}
			return false, err
		}
		t.FileList = append(t.FileList, file)
	}
	return true, nil
}
