package tls

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"github.com/pkg/errors"
	"github.com/openshift/installer/pkg/asset"
)

type CertInterface interface{ Cert() []byte }
type CertKeyInterface interface {
	CertInterface
	Key() []byte
}
type CertKey struct {
	CertRaw		[]byte
	KeyRaw		[]byte
	FileList	[]*asset.File
}

func (c *CertKey) Cert() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.CertRaw
}
func (c *CertKey) Key() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.KeyRaw
}
func (c *CertKey) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.FileList
}
func (c *CertKey) CertFile() *asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.FileList[1]
}
func (c *CertKey) generateFiles(filenameBase string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.FileList = []*asset.File{{Filename: assetFilePath(filenameBase + ".key"), Data: c.KeyRaw}, {Filename: assetFilePath(filenameBase + ".crt"), Data: c.CertRaw}}
}
func (c *CertKey) Load(asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false, nil
}

type AppendParentChoice bool

const (
	AppendParent		AppendParentChoice	= true
	DoNotAppendParent	AppendParentChoice	= false
)

type SignedCertKey struct{ CertKey }

func (c *SignedCertKey) Generate(cfg *CertCfg, parentCA CertKeyInterface, filenameBase string, appendParent AppendParentChoice) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var key *rsa.PrivateKey
	var crt *x509.Certificate
	var err error
	caKey, err := PemToPrivateKey(parentCA.Key())
	if err != nil {
		return errors.Wrap(err, "failed to parse rsa private key")
	}
	caCert, err := PemToCertificate(parentCA.Cert())
	if err != nil {
		return errors.Wrap(err, "failed to parse x509 certificate")
	}
	key, crt, err = GenerateSignedCertificate(caKey, caCert, cfg)
	if err != nil {
		return errors.Wrap(err, "failed to generate signed cert/key pair")
	}
	c.KeyRaw = PrivateKeyToPem(key)
	c.CertRaw = CertToPem(crt)
	if appendParent {
		c.CertRaw = bytes.Join([][]byte{c.CertRaw, CertToPem(caCert)}, []byte("\n"))
	}
	c.generateFiles(filenameBase)
	return nil
}

type SelfSignedCertKey struct{ CertKey }

func (c *SelfSignedCertKey) Generate(cfg *CertCfg, filenameBase string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, crt, err := GenerateSelfSignedCertificate(cfg)
	if err != nil {
		return errors.Wrap(err, "failed to generate self-signed cert/key pair")
	}
	c.KeyRaw = PrivateKeyToPem(key)
	c.CertRaw = CertToPem(crt)
	c.generateFiles(filenameBase)
	return nil
}
