package bootkube

import (
	"os"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	openshiftConfigSecretEtcdMetricClientFileName = "openshift-config-secret-etcd-metric-client.yaml.template"
)

var _ asset.WritableAsset = (*OpenshiftConfigSecretEtcdMetricClient)(nil)

type OpenshiftConfigSecretEtcdMetricClient struct{ FileList []*asset.File }

func (t *OpenshiftConfigSecretEtcdMetricClient) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (t *OpenshiftConfigSecretEtcdMetricClient) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "OpenshiftConfigSecretEtcdMetricClient"
}
func (t *OpenshiftConfigSecretEtcdMetricClient) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fileName := openshiftConfigSecretEtcdMetricClientFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{{Filename: filepath.Join(content.TemplateDir, fileName), Data: []byte(data)}}
	return nil
}
func (t *OpenshiftConfigSecretEtcdMetricClient) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.FileList
}
func (t *OpenshiftConfigSecretEtcdMetricClient) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, openshiftConfigSecretEtcdMetricClientFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
