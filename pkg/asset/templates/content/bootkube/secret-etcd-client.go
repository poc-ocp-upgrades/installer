package bootkube

import (
	"os"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	openshiftConfigSecretEtcdClientFileName = "openshift-config-secret-etcd-client.yaml.template"
)

var etcdClientCertFiles = []string{openshiftConfigSecretEtcdClientFileName}
var _ asset.WritableAsset = (*KubeSystemSecretEtcdClient)(nil)

type KubeSystemSecretEtcdClient struct{ FileList []*asset.File }

func (t *KubeSystemSecretEtcdClient) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (t *KubeSystemSecretEtcdClient) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "SecretEtcdClient"
}
func (t *KubeSystemSecretEtcdClient) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func (t *KubeSystemSecretEtcdClient) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.FileList
}
func (t *KubeSystemSecretEtcdClient) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
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
