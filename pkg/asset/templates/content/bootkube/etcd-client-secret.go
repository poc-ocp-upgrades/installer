package bootkube

import (
	"os"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	etcdClientSecretFileName = "etcd-client-secret.yaml.template"
)

var etcdClientCertFiles = []string{etcdClientSecretFileName}
var _ asset.WritableAsset = (*EtcdClientSecret)(nil)

type EtcdClientSecret struct{ FileList []*asset.File }

func (t *EtcdClientSecret) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (t *EtcdClientSecret) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "EtcdClientSecret"
}
func (t *EtcdClientSecret) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.FileList = []*asset.File{}
	for _, fileName := range etcdClientCertFiles {
		data, err := content.GetBootkubeTemplate(fileName)
		if err != nil {
			return err
		}
		t.FileList = append(t.FileList, &asset.File{Filename: filepath.Join(content.TemplateDir, fileName), Data: []byte(data)})
	}
	return nil
}
func (t *EtcdClientSecret) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.FileList
}
func (t *EtcdClientSecret) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.FileList = []*asset.File{}
	for _, fileName := range etcdClientCertFiles {
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
