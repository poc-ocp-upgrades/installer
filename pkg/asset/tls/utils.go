package tls

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/pkg/errors"
)

func PrivateKeyToPem(key *rsa.PrivateKey) []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	keyInBytes := x509.MarshalPKCS1PrivateKey(key)
	keyinPem := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: keyInBytes})
	return keyinPem
}
func CertToPem(cert *x509.Certificate) []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	certInPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw})
	return certInPem
}
func CSRToPem(cert *x509.CertificateRequest) []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	certInPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: cert.Raw})
	return certInPem
}
func PublicKeyToPem(key *rsa.PublicKey) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	keyInBytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to MarshalPKIXPublicKey")
	}
	keyinPem := pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: keyInBytes})
	return keyinPem, nil
}
func PemToPrivateKey(data []byte) (*rsa.PrivateKey, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.Errorf("could not find a PEM block in the private key")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}
func PemToCertificate(data []byte) (*x509.Certificate, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.Errorf("could not find a PEM block in the certificate")
	}
	return x509.ParseCertificate(block.Bytes)
}
