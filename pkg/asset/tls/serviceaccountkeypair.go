package tls

import "github.com/openshift/installer/pkg/asset"

type ServiceAccountKeyPair struct{ KeyPair }

var _ asset.WritableAsset = (*ServiceAccountKeyPair)(nil)

func (a *ServiceAccountKeyPair) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (a *ServiceAccountKeyPair) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.KeyPair.Generate("service-account")
}
func (a *ServiceAccountKeyPair) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Key Pair (service-account.pub)"
}
func (a *ServiceAccountKeyPair) Load(asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false, nil
}
