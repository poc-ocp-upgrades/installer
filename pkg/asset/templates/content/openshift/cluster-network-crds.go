package openshift

import (
	"os"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	netopCRDfilename = "cluster-networkconfig-crd.yaml"
)

var _ asset.WritableAsset = (*NetworkCRDs)(nil)

type NetworkCRDs struct{ FileList []*asset.File }

func (t *NetworkCRDs) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (t *NetworkCRDs) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Network CRDs"
}
func (t *NetworkCRDs) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	data, err := content.GetOpenshiftTemplate(netopCRDfilename)
	if err != nil {
		return err
	}
	t.FileList = append(t.FileList, &asset.File{Filename: filepath.Join(content.TemplateDir, netopCRDfilename), Data: []byte(data)})
	return nil
}
func (t *NetworkCRDs) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.FileList
}
func (t *NetworkCRDs) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, netopCRDfilename))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = append(t.FileList, file)
	return true, nil
}
