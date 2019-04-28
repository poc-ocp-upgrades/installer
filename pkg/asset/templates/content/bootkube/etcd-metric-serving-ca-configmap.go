package bootkube

import (
	"os"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	etcdMetricServingCAConfigMapFileName = "etcd-metric-serving-ca-configmap.yaml.template"
)

var _ asset.WritableAsset = (*EtcdMetricServingCAConfigMap)(nil)

type EtcdMetricServingCAConfigMap struct{ FileList []*asset.File }

func (t *EtcdMetricServingCAConfigMap) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (t *EtcdMetricServingCAConfigMap) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "EtcdMetricServingCAConfigMap"
}
func (t *EtcdMetricServingCAConfigMap) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fileName := etcdMetricServingCAConfigMapFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{{Filename: filepath.Join(content.TemplateDir, fileName), Data: []byte(data)}}
	return nil
}
func (t *EtcdMetricServingCAConfigMap) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.FileList
}
func (t *EtcdMetricServingCAConfigMap) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, etcdMetricServingCAConfigMapFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
