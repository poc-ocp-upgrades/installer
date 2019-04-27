package bootkube

import (
	"os"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	openshiftConfigSecretPullSecretFileName = "openshift-config-secret-pull-secret.yaml.template"
)

var _ asset.WritableAsset = (*OpenshiftConfigSecretPullSecret)(nil)

type OpenshiftConfigSecretPullSecret struct{ FileList []*asset.File }

func (t *OpenshiftConfigSecretPullSecret) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (t *OpenshiftConfigSecretPullSecret) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "OpenshiftConfigSecretPullSecret"
}
func (t *OpenshiftConfigSecretPullSecret) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fileName := openshiftConfigSecretPullSecretFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{{Filename: filepath.Join(content.TemplateDir, fileName), Data: []byte(data)}}
	return nil
}
func (t *OpenshiftConfigSecretPullSecret) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.FileList
}
func (t *OpenshiftConfigSecretPullSecret) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, openshiftConfigSecretPullSecretFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
