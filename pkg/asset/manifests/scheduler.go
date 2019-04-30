package manifests

import (
	"path/filepath"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	configv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/installer/pkg/asset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	schedulerCfgFilename = filepath.Join(manifestDir, "cluster-scheduler-02-config.yml")
)

type Scheduler struct{ FileList []*asset.File }

var _ asset.WritableAsset = (*Scheduler)(nil)

func (*Scheduler) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Scheduler Config"
}
func (*Scheduler) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (s *Scheduler) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := &configv1.Scheduler{TypeMeta: metav1.TypeMeta{APIVersion: configv1.SchemeGroupVersion.String(), Kind: "Scheduler"}, ObjectMeta: metav1.ObjectMeta{Name: "cluster"}, Spec: configv1.SchedulerSpec{}}
	configData, err := yaml.Marshal(config)
	if err != nil {
		return errors.Wrapf(err, "failed to create %s manifests from InstallConfig", s.Name())
	}
	s.FileList = []*asset.File{{Filename: schedulerCfgFilename, Data: configData}}
	return nil
}
func (s *Scheduler) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.FileList
}
func (s *Scheduler) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false, nil
}
