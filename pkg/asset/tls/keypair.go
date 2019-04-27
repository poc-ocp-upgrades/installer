package tls

import (
	"github.com/openshift/installer/pkg/asset"
	"github.com/pkg/errors"
)

type KeyPairInterface interface {
	Private() []byte
	Public() []byte
}
type KeyPair struct {
	Pvt		[]byte
	Pub		[]byte
	FileList	[]*asset.File
}

func (k *KeyPair) Generate(filenameBase string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := PrivateKey()
	if err != nil {
		return errors.Wrap(err, "failed to generate private key")
	}
	pubkeyData, err := PublicKeyToPem(&key.PublicKey)
	if err != nil {
		return errors.Wrap(err, "failed to get public key data from private key")
	}
	k.Pvt = PrivateKeyToPem(key)
	k.Pub = pubkeyData
	k.FileList = []*asset.File{{Filename: assetFilePath(filenameBase + ".key"), Data: k.Pvt}, {Filename: assetFilePath(filenameBase + ".pub"), Data: k.Pub}}
	return nil
}
func (k *KeyPair) Public() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return k.Pub
}
func (k *KeyPair) Private() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return k.Pvt
}
func (k *KeyPair) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return k.FileList
}
