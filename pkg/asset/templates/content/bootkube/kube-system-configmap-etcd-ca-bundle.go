package bootkube

import (
	"os"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	kubeSystemConfigmapEtcdCAFileName = "kube-system-configmap-etcd-ca-bundle.yaml.template"
)

var _ asset.WritableAsset = (*KubeSystemConfigmapEtcdCA)(nil)

type KubeSystemConfigmapEtcdCA struct{ FileList []*asset.File }

func (t *KubeSystemConfigmapEtcdCA) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (t *KubeSystemConfigmapEtcdCA) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "KubeSystemConfigmapEtcdCA"
}
func (t *KubeSystemConfigmapEtcdCA) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fileName := kubeSystemConfigmapEtcdCAFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{{Filename: filepath.Join(content.TemplateDir, fileName), Data: []byte(data)}}
	return nil
}
func (t *KubeSystemConfigmapEtcdCA) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.FileList
}
func (t *KubeSystemConfigmapEtcdCA) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, kubeSystemConfigmapEtcdCAFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
