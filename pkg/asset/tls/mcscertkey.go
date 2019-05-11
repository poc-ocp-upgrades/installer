package tls

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
)

type MCSCertKey struct{ SignedCertKey }

var _ asset.Asset = (*MCSCertKey)(nil)

func (a *MCSCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&RootCA{}, &installconfig.InstallConfig{}}
}
func (a *MCSCertKey) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ca := &RootCA{}
	installConfig := &installconfig.InstallConfig{}
	dependencies.Get(ca, installConfig)
	hostname := apiAddress(installConfig.Config)
	cfg := &CertCfg{Subject: pkix.Name{CommonName: hostname}, ExtKeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, Validity: ValidityTenYears, DNSNames: []string{hostname}}
	return a.SignedCertKey.Generate(cfg, ca, "machine-config-server", DoNotAppendParent)
}
func (a *MCSCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (mcs)"
}
