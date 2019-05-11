package tls

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"github.com/openshift/installer/pkg/asset"
)

type EtcdMetricSignerCertKey struct{ SelfSignedCertKey }

var _ asset.WritableAsset = (*EtcdMetricSignerCertKey)(nil)

func (c *EtcdMetricSignerCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (c *EtcdMetricSignerCertKey) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "etcd-metric-signer", OrganizationalUnit: []string{"openshift"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, Validity: ValidityTenYears, IsCA: true}
	return c.SelfSignedCertKey.Generate(cfg, "etcd-metric-signer")
}
func (c *EtcdMetricSignerCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (etcd-metric-signer)"
}

type EtcdMetricCABundle struct{ CertBundle }

var _ asset.Asset = (*EtcdMetricCABundle)(nil)

func (a *EtcdMetricCABundle) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&EtcdMetricSignerCertKey{}}
}
func (a *EtcdMetricCABundle) Generate(deps asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var certs []CertInterface
	for _, asset := range a.Dependencies() {
		deps.Get(asset)
		certs = append(certs, asset.(CertInterface))
	}
	return a.CertBundle.Generate("etcd-metric-ca-bundle", certs...)
}
func (a *EtcdMetricCABundle) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (etcd-metric-ca-bundle)"
}

type EtcdMetricSignerClientCertKey struct{ SignedCertKey }

var _ asset.Asset = (*EtcdMetricSignerClientCertKey)(nil)

func (a *EtcdMetricSignerClientCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&EtcdMetricSignerCertKey{}}
}
func (a *EtcdMetricSignerClientCertKey) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ca := &EtcdMetricSignerCertKey{}
	dependencies.Get(ca)
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "etcd-metric", OrganizationalUnit: []string{"etcd-metric"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, ExtKeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}, Validity: ValidityTenYears}
	return a.SignedCertKey.Generate(cfg, ca, "etcd-metric-signer-client", DoNotAppendParent)
}
func (a *EtcdMetricSignerClientCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (etcd-metric-signer-client)"
}
