package tls

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"github.com/openshift/installer/pkg/asset"
)

type EtcdCA struct{ SelfSignedCertKey }

var _ asset.Asset = (*EtcdCA)(nil)

func (a *EtcdCA) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (a *EtcdCA) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "etcd", OrganizationalUnit: []string{"etcd"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, Validity: ValidityTenYears, IsCA: true}
	return a.SelfSignedCertKey.Generate(cfg, "etcd-client-ca")
}
func (a *EtcdCA) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (etcd)"
}

type EtcdClientCertKey struct{ SignedCertKey }

var _ asset.Asset = (*EtcdClientCertKey)(nil)

func (a *EtcdClientCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&EtcdCA{}}
}
func (a *EtcdClientCertKey) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	etcdCA := &EtcdCA{}
	dependencies.Get(etcdCA)
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "etcd", OrganizationalUnit: []string{"etcd"}}, KeyUsages: x509.KeyUsageKeyEncipherment, ExtKeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}, Validity: ValidityTenYears}
	return a.SignedCertKey.Generate(cfg, etcdCA, "etcd-client", DoNotAppendParent)
}
func (a *EtcdClientCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (etcd)"
}

type EtcdSignerCertKey struct{ SelfSignedCertKey }

var _ asset.WritableAsset = (*EtcdSignerCertKey)(nil)

func (c *EtcdSignerCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (c *EtcdSignerCertKey) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "etcd-signer", OrganizationalUnit: []string{"openshift"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, Validity: ValidityTenYears, IsCA: true}
	return c.SelfSignedCertKey.Generate(cfg, "etcd-signer")
}
func (c *EtcdSignerCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (etcd-signer)"
}

type EtcdCABundle struct{ CertBundle }

var _ asset.Asset = (*EtcdCABundle)(nil)

func (a *EtcdCABundle) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&EtcdSignerCertKey{}}
}
func (a *EtcdCABundle) Generate(deps asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var certs []CertInterface
	for _, asset := range a.Dependencies() {
		deps.Get(asset)
		certs = append(certs, asset.(CertInterface))
	}
	return a.CertBundle.Generate("etcd-ca-bundle", certs...)
}
func (a *EtcdCABundle) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (etcd-ca-bundle)"
}

type EtcdSignerClientCertKey struct{ SignedCertKey }

var _ asset.Asset = (*EtcdSignerClientCertKey)(nil)

func (a *EtcdSignerClientCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&EtcdSignerCertKey{}}
}
func (a *EtcdSignerClientCertKey) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ca := &EtcdSignerCertKey{}
	dependencies.Get(ca)
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "etcd", OrganizationalUnit: []string{"etcd"}}, KeyUsages: x509.KeyUsageKeyEncipherment, ExtKeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}, Validity: ValidityTenYears}
	return a.SignedCertKey.Generate(cfg, ca, "etcd-signer-client", DoNotAppendParent)
}
func (a *EtcdSignerClientCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (etcd-signer-client)"
}
