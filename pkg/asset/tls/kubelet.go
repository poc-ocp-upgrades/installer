package tls

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"github.com/openshift/installer/pkg/asset"
)

type KubeletCSRSignerCertKey struct{ SelfSignedCertKey }

var _ asset.WritableAsset = (*KubeletCSRSignerCertKey)(nil)

func (c *KubeletCSRSignerCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (c *KubeletCSRSignerCertKey) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "kubelet-signer", OrganizationalUnit: []string{"openshift"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, Validity: ValidityOneDay, IsCA: true}
	return c.SelfSignedCertKey.Generate(cfg, "kubelet-signer")
}
func (c *KubeletCSRSignerCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kubelet-signer)"
}

type KubeletClientCABundle struct{ CertBundle }

var _ asset.Asset = (*KubeletClientCABundle)(nil)

func (a *KubeletClientCABundle) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&KubeletCSRSignerCertKey{}}
}
func (a *KubeletClientCABundle) Generate(deps asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var certs []CertInterface
	for _, asset := range a.Dependencies() {
		deps.Get(asset)
		certs = append(certs, asset.(CertInterface))
	}
	return a.CertBundle.Generate("kubelet-client-ca-bundle", certs...)
}
func (a *KubeletClientCABundle) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kubelet-client-ca-bundle)"
}

type KubeletServingCABundle struct{ CertBundle }

var _ asset.Asset = (*KubeletServingCABundle)(nil)

func (a *KubeletServingCABundle) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&KubeletCSRSignerCertKey{}}
}
func (a *KubeletServingCABundle) Generate(deps asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var certs []CertInterface
	for _, asset := range a.Dependencies() {
		deps.Get(asset)
		certs = append(certs, asset.(CertInterface))
	}
	return a.CertBundle.Generate("kubelet-serving-ca-bundle", certs...)
}
func (a *KubeletServingCABundle) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kubelet-serving-ca-bundle)"
}

type KubeletBootstrapCertSigner struct{ SelfSignedCertKey }

var _ asset.WritableAsset = (*KubeletBootstrapCertSigner)(nil)

func (c *KubeletBootstrapCertSigner) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (c *KubeletBootstrapCertSigner) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "kubelet-bootstrap-kubeconfig-signer", OrganizationalUnit: []string{"openshift"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, Validity: ValidityTenYears, IsCA: true}
	return c.SelfSignedCertKey.Generate(cfg, "kubelet-bootstrap-kubeconfig-signer")
}
func (c *KubeletBootstrapCertSigner) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kubelet-bootstrap-kubeconfig-signer)"
}

type KubeletBootstrapCABundle struct{ CertBundle }

var _ asset.Asset = (*KubeletBootstrapCABundle)(nil)

func (a *KubeletBootstrapCABundle) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&KubeletBootstrapCertSigner{}}
}
func (a *KubeletBootstrapCABundle) Generate(deps asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var certs []CertInterface
	for _, asset := range a.Dependencies() {
		deps.Get(asset)
		certs = append(certs, asset.(CertInterface))
	}
	return a.CertBundle.Generate("kubelet-bootstrap-kubeconfig-ca-bundle", certs...)
}
func (a *KubeletBootstrapCABundle) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kubelet-bootstrap-kubeconfig-ca-bundle)"
}

type KubeletClientCertKey struct{ SignedCertKey }

var _ asset.Asset = (*KubeletClientCertKey)(nil)

func (a *KubeletClientCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&KubeletBootstrapCertSigner{}}
}
func (a *KubeletClientCertKey) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ca := &KubeletBootstrapCertSigner{}
	dependencies.Get(ca)
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "system:serviceaccount:openshift-machine-config-operator:node-bootstrapper", Organization: []string{"system:serviceaccounts:openshift-machine-config-operator"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, ExtKeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}, Validity: ValidityOneDay}
	return a.SignedCertKey.Generate(cfg, ca, "kubelet-client", DoNotAppendParent)
}
func (a *KubeletClientCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kubelet-client)"
}
