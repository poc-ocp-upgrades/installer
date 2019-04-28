package kubeconfig

import (
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/tls"
)

var (
	kubeconfigKubeletPath = filepath.Join("auth", "kubeconfig-kubelet")
)

type Kubelet struct{ kubeconfig }

var _ asset.WritableAsset = (*Kubelet)(nil)

func (k *Kubelet) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&tls.KubeAPIServerCompleteCABundle{}, &tls.KubeletClientCertKey{}, &installconfig.InstallConfig{}}
}
func (k *Kubelet) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ca := &tls.KubeAPIServerCompleteCABundle{}
	clientcertkey := &tls.KubeletClientCertKey{}
	installConfig := &installconfig.InstallConfig{}
	parents.Get(ca, clientcertkey, installConfig)
	return k.kubeconfig.generate(ca, clientcertkey, getIntAPIServerURL(installConfig.Config), installConfig.Config.GetName(), "kubelet", kubeconfigKubeletPath)
}
func (k *Kubelet) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Kubeconfig Kubelet"
}
func (k *Kubelet) Load(asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false, nil
}
