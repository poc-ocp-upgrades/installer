package bootkube

import (
	"os"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	openshiftConfigConfigmapEtcdMetricServingCAFileName = "openshift-config-configmap-etcd-metric-serving-ca.yaml.template"
)

var _ asset.WritableAsset = (*OpenshiftConfigConfigmapEtcdMetricServingCA)(nil)

type OpenshiftConfigConfigmapEtcdMetricServingCA struct{ FileList []*asset.File }

func (t *OpenshiftConfigConfigmapEtcdMetricServingCA) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (t *OpenshiftConfigConfigmapEtcdMetricServingCA) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "OpenshiftConfigConfigmapEtcdMetricServingCA"
}
func (t *OpenshiftConfigConfigmapEtcdMetricServingCA) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fileName := openshiftConfigConfigmapEtcdMetricServingCAFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{{Filename: filepath.Join(content.TemplateDir, fileName), Data: []byte(data)}}
	return nil
}
func (t *OpenshiftConfigConfigmapEtcdMetricServingCA) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.FileList
}
func (t *OpenshiftConfigConfigmapEtcdMetricServingCA) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, openshiftConfigConfigmapEtcdMetricServingCAFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
