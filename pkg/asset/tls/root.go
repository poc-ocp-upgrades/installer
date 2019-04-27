package tls

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"github.com/openshift/installer/pkg/asset"
)

type RootCA struct{ SelfSignedCertKey }

var _ asset.WritableAsset = (*RootCA)(nil)

func (c *RootCA) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (c *RootCA) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "root-ca", OrganizationalUnit: []string{"openshift"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, Validity: ValidityTenYears, IsCA: true}
	return c.SelfSignedCertKey.Generate(cfg, "root-ca")
}
func (c *RootCA) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Root CA"
}
