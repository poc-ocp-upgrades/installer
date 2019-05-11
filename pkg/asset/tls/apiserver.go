package tls

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"net"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/pkg/errors"
)

type KubeAPIServerToKubeletSignerCertKey struct{ SelfSignedCertKey }

var _ asset.WritableAsset = (*KubeAPIServerToKubeletSignerCertKey)(nil)

func (c *KubeAPIServerToKubeletSignerCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (c *KubeAPIServerToKubeletSignerCertKey) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "kube-apiserver-to-kubelet-signer", OrganizationalUnit: []string{"openshift"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, Validity: ValidityOneYear, IsCA: true}
	return c.SelfSignedCertKey.Generate(cfg, "kube-apiserver-to-kubelet-signer")
}
func (c *KubeAPIServerToKubeletSignerCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-apiserver-to-kubelet-signer)"
}

type KubeAPIServerToKubeletCABundle struct{ CertBundle }

var _ asset.Asset = (*KubeAPIServerToKubeletCABundle)(nil)

func (a *KubeAPIServerToKubeletCABundle) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&KubeAPIServerToKubeletSignerCertKey{}}
}
func (a *KubeAPIServerToKubeletCABundle) Generate(deps asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var certs []CertInterface
	for _, asset := range a.Dependencies() {
		deps.Get(asset)
		certs = append(certs, asset.(CertInterface))
	}
	return a.CertBundle.Generate("kube-apiserver-to-kubelet-ca-bundle", certs...)
}
func (a *KubeAPIServerToKubeletCABundle) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-apiserver-to-kubelet-ca-bundle)"
}

type KubeAPIServerToKubeletClientCertKey struct{ SignedCertKey }

var _ asset.Asset = (*KubeAPIServerToKubeletClientCertKey)(nil)

func (a *KubeAPIServerToKubeletClientCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&KubeAPIServerToKubeletSignerCertKey{}}
}
func (a *KubeAPIServerToKubeletClientCertKey) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ca := &KubeAPIServerToKubeletSignerCertKey{}
	dependencies.Get(ca)
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "system:kube-apiserver", Organization: []string{"kube-master"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, ExtKeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}, Validity: ValidityOneYear}
	return a.SignedCertKey.Generate(cfg, ca, "kube-apiserver-to-kubelet-client", DoNotAppendParent)
}
func (a *KubeAPIServerToKubeletClientCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-apiserver-to-kubelet-client)"
}

type KubeAPIServerLocalhostSignerCertKey struct{ SelfSignedCertKey }

var _ asset.WritableAsset = (*KubeAPIServerLocalhostSignerCertKey)(nil)

func (c *KubeAPIServerLocalhostSignerCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (c *KubeAPIServerLocalhostSignerCertKey) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "kube-apiserver-localhost-signer", OrganizationalUnit: []string{"openshift"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, Validity: ValidityTenYears, IsCA: true}
	return c.SelfSignedCertKey.Generate(cfg, "kube-apiserver-localhost-signer")
}
func (c *KubeAPIServerLocalhostSignerCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-apiserver-localhost-signer)"
}

type KubeAPIServerLocalhostCABundle struct{ CertBundle }

var _ asset.Asset = (*KubeAPIServerLocalhostCABundle)(nil)

func (a *KubeAPIServerLocalhostCABundle) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&KubeAPIServerLocalhostSignerCertKey{}}
}
func (a *KubeAPIServerLocalhostCABundle) Generate(deps asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var certs []CertInterface
	for _, asset := range a.Dependencies() {
		deps.Get(asset)
		certs = append(certs, asset.(CertInterface))
	}
	return a.CertBundle.Generate("kube-apiserver-localhost-ca-bundle", certs...)
}
func (a *KubeAPIServerLocalhostCABundle) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-apiserver-localhost-ca-bundle)"
}

type KubeAPIServerLocalhostServerCertKey struct{ SignedCertKey }

var _ asset.Asset = (*KubeAPIServerLocalhostServerCertKey)(nil)

func (a *KubeAPIServerLocalhostServerCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&KubeAPIServerLocalhostSignerCertKey{}}
}
func (a *KubeAPIServerLocalhostServerCertKey) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ca := &KubeAPIServerLocalhostSignerCertKey{}
	dependencies.Get(ca)
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "system:kube-apiserver", Organization: []string{"kube-master"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, ExtKeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, Validity: ValidityOneDay, DNSNames: []string{"localhost"}, IPAddresses: []net.IP{net.ParseIP("127.0.0.1"), net.ParseIP("::1")}}
	return a.SignedCertKey.Generate(cfg, ca, "kube-apiserver-localhost-server", AppendParent)
}
func (a *KubeAPIServerLocalhostServerCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-apiserver-localhost-server)"
}

type KubeAPIServerServiceNetworkSignerCertKey struct{ SelfSignedCertKey }

var _ asset.WritableAsset = (*KubeAPIServerServiceNetworkSignerCertKey)(nil)

func (c *KubeAPIServerServiceNetworkSignerCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (c *KubeAPIServerServiceNetworkSignerCertKey) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "kube-apiserver-service-network-signer", OrganizationalUnit: []string{"openshift"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, Validity: ValidityTenYears, IsCA: true}
	return c.SelfSignedCertKey.Generate(cfg, "kube-apiserver-service-network-signer")
}
func (c *KubeAPIServerServiceNetworkSignerCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-apiserver-service-network-signer)"
}

type KubeAPIServerServiceNetworkCABundle struct{ CertBundle }

var _ asset.Asset = (*KubeAPIServerServiceNetworkCABundle)(nil)

func (a *KubeAPIServerServiceNetworkCABundle) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&KubeAPIServerServiceNetworkSignerCertKey{}}
}
func (a *KubeAPIServerServiceNetworkCABundle) Generate(deps asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var certs []CertInterface
	for _, asset := range a.Dependencies() {
		deps.Get(asset)
		certs = append(certs, asset.(CertInterface))
	}
	return a.CertBundle.Generate("kube-apiserver-service-network-ca-bundle", certs...)
}
func (a *KubeAPIServerServiceNetworkCABundle) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-apiserver-service-network-ca-bundle)"
}

type KubeAPIServerServiceNetworkServerCertKey struct{ SignedCertKey }

var _ asset.Asset = (*KubeAPIServerServiceNetworkServerCertKey)(nil)

func (a *KubeAPIServerServiceNetworkServerCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&KubeAPIServerServiceNetworkSignerCertKey{}, &installconfig.InstallConfig{}}
}
func (a *KubeAPIServerServiceNetworkServerCertKey) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ca := &KubeAPIServerServiceNetworkSignerCertKey{}
	installConfig := &installconfig.InstallConfig{}
	dependencies.Get(ca, installConfig)
	serviceAddress, err := cidrhost(installConfig.Config.Networking.ServiceNetwork[0].IPNet, 1)
	if err != nil {
		return errors.Wrap(err, "failed to get service address for kube-apiserver from InstallConfig")
	}
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "system:kube-apiserver", Organization: []string{"kube-master"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, ExtKeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, Validity: ValidityOneDay, DNSNames: []string{"kubernetes", "kubernetes.default", "kubernetes.default.svc", "kubernetes.default.svc.cluster.local"}, IPAddresses: []net.IP{net.ParseIP(serviceAddress)}}
	return a.SignedCertKey.Generate(cfg, ca, "kube-apiserver-service-network-server", AppendParent)
}
func (a *KubeAPIServerServiceNetworkServerCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-apiserver-service-network-server)"
}

type KubeAPIServerLBSignerCertKey struct{ SelfSignedCertKey }

var _ asset.WritableAsset = (*KubeAPIServerLBSignerCertKey)(nil)

func (c *KubeAPIServerLBSignerCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (c *KubeAPIServerLBSignerCertKey) Generate(parents asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "kube-apiserver-lb-signer", OrganizationalUnit: []string{"openshift"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, Validity: ValidityTenYears, IsCA: true}
	return c.SelfSignedCertKey.Generate(cfg, "kube-apiserver-lb-signer")
}
func (c *KubeAPIServerLBSignerCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-apiserver-lb-signer)"
}

type KubeAPIServerLBCABundle struct{ CertBundle }

var _ asset.Asset = (*KubeAPIServerLBCABundle)(nil)

func (a *KubeAPIServerLBCABundle) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&KubeAPIServerLBSignerCertKey{}}
}
func (a *KubeAPIServerLBCABundle) Generate(deps asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var certs []CertInterface
	for _, asset := range a.Dependencies() {
		deps.Get(asset)
		certs = append(certs, asset.(CertInterface))
	}
	return a.CertBundle.Generate("kube-apiserver-lb-ca-bundle", certs...)
}
func (a *KubeAPIServerLBCABundle) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-apiserver-lb-ca-bundle)"
}

type KubeAPIServerExternalLBServerCertKey struct{ SignedCertKey }

var _ asset.Asset = (*KubeAPIServerExternalLBServerCertKey)(nil)

func (a *KubeAPIServerExternalLBServerCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&KubeAPIServerLBSignerCertKey{}, &installconfig.InstallConfig{}}
}
func (a *KubeAPIServerExternalLBServerCertKey) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ca := &KubeAPIServerLBSignerCertKey{}
	installConfig := &installconfig.InstallConfig{}
	dependencies.Get(ca, installConfig)
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "system:kube-apiserver", Organization: []string{"kube-master"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, ExtKeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, Validity: ValidityOneDay, DNSNames: []string{apiAddress(installConfig.Config)}}
	return a.SignedCertKey.Generate(cfg, ca, "kube-apiserver-lb-server", AppendParent)
}
func (a *KubeAPIServerExternalLBServerCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-apiserver-external-lb-server)"
}

type KubeAPIServerInternalLBServerCertKey struct{ SignedCertKey }

var _ asset.Asset = (*KubeAPIServerInternalLBServerCertKey)(nil)

func (a *KubeAPIServerInternalLBServerCertKey) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&KubeAPIServerLBSignerCertKey{}, &installconfig.InstallConfig{}}
}
func (a *KubeAPIServerInternalLBServerCertKey) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ca := &KubeAPIServerLBSignerCertKey{}
	installConfig := &installconfig.InstallConfig{}
	dependencies.Get(ca, installConfig)
	cfg := &CertCfg{Subject: pkix.Name{CommonName: "system:kube-apiserver", Organization: []string{"kube-master"}}, KeyUsages: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, ExtKeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, Validity: ValidityOneDay, DNSNames: []string{internalAPIAddress(installConfig.Config)}}
	return a.SignedCertKey.Generate(cfg, ca, "kube-apiserver-internal-lb-server", AppendParent)
}
func (a *KubeAPIServerInternalLBServerCertKey) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-apiserver-internal-lb-server)"
}

type KubeAPIServerCompleteCABundle struct{ CertBundle }

var _ asset.Asset = (*KubeAPIServerCompleteCABundle)(nil)

func (a *KubeAPIServerCompleteCABundle) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&KubeAPIServerLocalhostCABundle{}, &KubeAPIServerServiceNetworkCABundle{}, &KubeAPIServerLBCABundle{}}
}
func (a *KubeAPIServerCompleteCABundle) Generate(deps asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var certs []CertInterface
	for _, asset := range a.Dependencies() {
		deps.Get(asset)
		certs = append(certs, asset.(CertInterface))
	}
	return a.CertBundle.Generate("kube-apiserver-complete-server-ca-bundle", certs...)
}
func (a *KubeAPIServerCompleteCABundle) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-apiserver-complete-server-ca-bundle)"
}

type KubeAPIServerCompleteClientCABundle struct{ CertBundle }

var _ asset.Asset = (*KubeAPIServerCompleteClientCABundle)(nil)

func (a *KubeAPIServerCompleteClientCABundle) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&AdminKubeConfigCABundle{}, &KubeletClientCABundle{}, &KubeControlPlaneCABundle{}, &KubeAPIServerToKubeletCABundle{}, &KubeletBootstrapCABundle{}}
}
func (a *KubeAPIServerCompleteClientCABundle) Generate(deps asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var certs []CertInterface
	for _, asset := range a.Dependencies() {
		deps.Get(asset)
		certs = append(certs, asset.(CertInterface))
	}
	return a.CertBundle.Generate("kube-apiserver-complete-client-ca-bundle", certs...)
}
func (a *KubeAPIServerCompleteClientCABundle) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Certificate (kube-apiserver-complete-client-ca-bundle)"
}
