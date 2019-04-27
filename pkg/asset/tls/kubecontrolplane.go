package tls

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"github.com/openshift/installer/pkg/asset"
)

type KubeControlPlaneSignerCertKey struct{ SelfSignedCertKey }

var _ asset.WritableAsset = (*KubeControlPlaneSignerCertKey)(nil)

func (c *KubeControlPlaneSignerCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (c *KubeControlPlaneSignerCertKey) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "kube-control-plane-signer", OrganizationalUnit: []string{"openshift"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, Validity: ValidityOneYear, IsCA: true}
	return c.SelfSignedCertKey.Generate(cfg, "kube-control-plane-signer")
}
func (c *KubeControlPlaneSignerCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-control-plane-signer)"
}

type KubeControlPlaneCABundle struct{ CertBundle }

var _ asset.Asset = (*KubeControlPlaneCABundle)(nil)

func (a *KubeControlPlaneCABundle) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&KubeControlPlaneSignerCertKey{}, &KubeAPIServerLBSignerCertKey{}, &KubeAPIServerLocalhostSignerCertKey{}, &KubeAPIServerServiceNetworkSignerCertKey{}}
}
func (a *KubeControlPlaneCABundle) Generate(deps asset.Parents) error {
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
	return a.CertBundle.Generate("kube-control-plane-ca-bundle", certs...)
}
func (a *KubeControlPlaneCABundle) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-control-plane-ca-bundle)"
}

type KubeControlPlaneKubeControllerManagerClientCertKey struct{ SignedCertKey }

var _ asset.Asset = (*KubeControlPlaneKubeControllerManagerClientCertKey)(nil)

func (a *KubeControlPlaneKubeControllerManagerClientCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&KubeControlPlaneSignerCertKey{}}
}
func (a *KubeControlPlaneKubeControllerManagerClientCertKey) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ca := &KubeControlPlaneSignerCertKey{}
	dependencies.Get(ca)
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "system:admin", Organization: []string{"system:masters"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, ExtKeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}, Validity: ValidityOneYear}
	return a.SignedCertKey.Generate(cfg, ca, "kube-control-plane-kube-controller-manager-client", DoNotAppendParent)
}
func (a *KubeControlPlaneKubeControllerManagerClientCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-control-plane-kube-controller-manager-client)"
}

type KubeControlPlaneKubeSchedulerClientCertKey struct{ SignedCertKey }

var _ asset.Asset = (*KubeControlPlaneKubeSchedulerClientCertKey)(nil)

func (a *KubeControlPlaneKubeSchedulerClientCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&KubeControlPlaneSignerCertKey{}}
}
func (a *KubeControlPlaneKubeSchedulerClientCertKey) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ca := &KubeControlPlaneSignerCertKey{}
	dependencies.Get(ca)
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "system:admin", Organization: []string{"system:masters"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, ExtKeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}, Validity: ValidityOneYear}
	return a.SignedCertKey.Generate(cfg, ca, "kube-control-plane-kube-scheduler-client", DoNotAppendParent)
}
func (a *KubeControlPlaneKubeSchedulerClientCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-control-plane-kube-scheduler-client)"
}
