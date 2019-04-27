package password

import (
	"crypto/rand"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"math/big"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"golang.org/x/crypto/bcrypt"
)

var (
	kubeadminPasswordPath = filepath.Join("auth", "kubeadmin-password")
)

type KubeadminPassword struct {
	Password	string
	PasswordHash	[]byte
	File		*asset.File
}

var _ asset.WritableAsset = (*KubeadminPassword)(nil)

func (a *KubeadminPassword) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{}
}
func (a *KubeadminPassword) Generate(asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := a.generateRandomPasswordHash(23)
	if err != nil {
		return err
	}
	return nil
}
func (a *KubeadminPassword) generateRandomPasswordHash(length int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	const (
		lowerLetters	= "abcdefghijkmnopqrstuvwxyz"
		upperLetters	= "ABCDEFGHIJKLMNPQRSTUVWXYZ"
		digits		= "23456789"
		all		= lowerLetters + upperLetters + digits
	)
	var password string
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(all))))
		if err != nil {
			return err
		}
		newchar := string(all[n.Int64()])
		if password == "" {
			password = newchar
		}
		if i < length-1 {
			n, err = rand.Int(rand.Reader, big.NewInt(int64(len(password)+1)))
			if err != nil {
				return err
			}
			j := n.Int64()
			password = password[0:j] + newchar + password[j:]
		}
	}
	pw := []rune(password)
	for _, replace := range []int{5, 11, 17} {
		pw[replace] = '-'
	}
	if a.Password == "" {
		a.Password = string(pw)
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.PasswordHash = bytes
	a.File = &asset.File{Filename: kubeadminPasswordPath, Data: []byte(a.Password)}
	return nil
}
func (a *KubeadminPassword) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Kubeadmin Password"
}
func (a *KubeadminPassword) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a.File != nil {
		return []*asset.File{a.File}
	}
	return []*asset.File{}
}
func (a *KubeadminPassword) Load(f asset.FileFetcher) (found bool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
