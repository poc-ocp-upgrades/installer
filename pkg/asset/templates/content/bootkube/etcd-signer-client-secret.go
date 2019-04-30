package bootkube

import (
	"os"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	etcdSignerClientSecretFileName = "etcd-signer-client-secret.yaml.template"
)

var _ asset.WritableAsset = (*EtcdSignerClientSecret)(nil)

type EtcdSignerClientSecret struct{ FileList []*asset.File }

func (t *EtcdSignerClientSecret) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (t *EtcdSignerClientSecret) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "EtcdSignerClientSecret"
}
func (t *EtcdSignerClientSecret) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fileName := etcdSignerClientSecretFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{{Filename: filepath.Join(content.TemplateDir, fileName), Data: []byte(data)}}
	return nil
}
func (t *EtcdSignerClientSecret) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.FileList
}
func (t *EtcdSignerClientSecret) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, etcdSignerClientSecretFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
