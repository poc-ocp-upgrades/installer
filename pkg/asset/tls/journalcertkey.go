package tls

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"github.com/openshift/installer/pkg/asset"
)

type JournalCertKey struct{ SignedCertKey }

var _ asset.WritableAsset = (*JournalCertKey)(nil)

func (a *JournalCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&RootCA{}}
}
func (a *JournalCertKey) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ca := &RootCA{}
	dependencies.Get(ca)
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "journal-gatewayd", Organization: []string{"OpenShift Bootstrap"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, ExtKeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}, Validity: ValidityTenYears}
	return a.SignedCertKey.Generate(cfg, ca, "journal-gatewayd", DoNotAppendParent)
}
func (a *JournalCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (journal-gatewayd)"
}
