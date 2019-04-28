package openshift

import (
	"os"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	kubeadminPasswordSecretFileName = "kubeadmin-password-secret.yaml.template"
)

var _ asset.WritableAsset = (*KubeadminPasswordSecret)(nil)

type KubeadminPasswordSecret struct{ FileList []*asset.File }

func (t *KubeadminPasswordSecret) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (t *KubeadminPasswordSecret) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "KubeadminPasswordSecret"
}
func (t *KubeadminPasswordSecret) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fileName := kubeadminPasswordSecretFileName
	data, err := content.GetOpenshiftTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{{Filename: filepath.Join(content.TemplateDir, fileName), Data: []byte(data)}}
	return nil
}
func (t *KubeadminPasswordSecret) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.FileList
}
func (t *KubeadminPasswordSecret) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, kubeadminPasswordSecretFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
