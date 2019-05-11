package tls

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"github.com/openshift/installer/pkg/asset"
)

type AggregatorCA struct{ SelfSignedCertKey }

var _ asset.Asset = (*AggregatorCA)(nil)

func (a *AggregatorCA) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (a *AggregatorCA) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "aggregator", OrganizationalUnit: []string{"bootkube"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, Validity: ValidityOneDay, IsCA: true}
	return a.SelfSignedCertKey.Generate(cfg, "aggregator-ca")
}
func (a *AggregatorCA) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (aggregator)"
}

type APIServerProxyCertKey struct{ SignedCertKey }

var _ asset.Asset = (*APIServerProxyCertKey)(nil)

func (a *APIServerProxyCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&AggregatorCA{}}
}
func (a *APIServerProxyCertKey) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	aggregatorCA := &AggregatorCA{}
	dependencies.Get(aggregatorCA)
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "system:kube-apiserver-proxy", Organization: []string{"kube-master"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, ExtKeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}, Validity: ValidityOneDay}
	return a.SignedCertKey.Generate(cfg, aggregatorCA, "apiserver-proxy", DoNotAppendParent)
}
func (a *APIServerProxyCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (system:kube-apiserver-proxy)"
}

type AggregatorSignerCertKey struct{ SelfSignedCertKey }

var _ asset.WritableAsset = (*AggregatorSignerCertKey)(nil)

func (c *AggregatorSignerCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (c *AggregatorSignerCertKey) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "aggregator-signer", OrganizationalUnit: []string{"openshift"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, Validity: ValidityOneDay, IsCA: true}
	return c.SelfSignedCertKey.Generate(cfg, "aggregator-signer")
}
func (c *AggregatorSignerCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (aggregator-signer)"
}

type AggregatorCABundle struct{ CertBundle }

var _ asset.Asset = (*AggregatorCABundle)(nil)

func (a *AggregatorCABundle) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&AggregatorSignerCertKey{}}
}
func (a *AggregatorCABundle) Generate(deps asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var certs []CertInterface
	for _, asset := range a.Dependencies() {
		deps.Get(asset)
		certs = append(certs, asset.(CertInterface))
	}
	return a.CertBundle.Generate("aggregator-ca-bundle", certs...)
}
func (a *AggregatorCABundle) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (aggregator-ca-bundle)"
}

type AggregatorClientCertKey struct{ SignedCertKey }

var _ asset.Asset = (*AggregatorClientCertKey)(nil)

func (a *AggregatorClientCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&AggregatorSignerCertKey{}}
}
func (a *AggregatorClientCertKey) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ca := &AggregatorSignerCertKey{}
	dependencies.Get(ca)
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "system:kube-apiserver-proxy", Organization: []string{"kube-master"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, ExtKeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}, Validity: ValidityOneDay}
	return a.SignedCertKey.Generate(cfg, ca, "aggregator-client", DoNotAppendParent)
}
func (a *AggregatorClientCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (system:kube-apiserver-proxy)"
}
