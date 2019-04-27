package tls

import (
	"bytes"
	"encoding/pem"
	"github.com/openshift/installer/pkg/asset"
	"github.com/pkg/errors"
)

type CertBundle struct {
	BundleRaw	[]byte
	FileList	[]*asset.File
}

func (b *CertBundle) Cert() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return b.BundleRaw
}
func (b *CertBundle) Generate(filename string, certs ...CertInterface) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(certs) < 1 {
		return errors.New("atleast one certificate required for a bundle")
	}
	buf := bytes.Buffer{}
	for _, c := range certs {
		cert, err := PemToCertificate(c.Cert())
		if err != nil {
			return errors.Wrap(err, "decoding certificate from PEM")
		}
		if err := pem.Encode(&buf, &pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw}); err != nil {
			return errors.Wrap(err, "encoding certificate to PEM")
		}
	}
	b.BundleRaw = buf.Bytes()
	b.FileList = []*asset.File{{Filename: assetFilePath(filename + ".crt"), Data: b.BundleRaw}}
	return nil
}
func (b *CertBundle) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return b.FileList
}
func (b *CertBundle) Load(asset.FileFetcher) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false, nil
}
