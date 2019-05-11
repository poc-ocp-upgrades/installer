package installconfig

import (
	"fmt"
	"regexp"
	"github.com/pborman/uuid"
	utilrand "k8s.io/apimachinery/pkg/util/rand"
	"github.com/openshift/installer/pkg/asset"
)

const (
	maxNameLen	= 32 - 5
	randomLen	= 5
	maxBaseLen	= maxNameLen - randomLen - 1
)

type ClusterID struct {
	UUID	string
	InfraID	string
}

var _ asset.Asset = (*ClusterID)(nil)

func (a *ClusterID) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&InstallConfig{}}
}
func (a *ClusterID) Generate(dep asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ica := &InstallConfig{}
	dep.Get(ica)
	a.InfraID = generateInfraID(ica.Config.ObjectMeta.Name)
	a.UUID = uuid.New()
	return nil
}
func (a *ClusterID) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Cluster ID"
}
func generateInfraID(base string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(base) > maxBaseLen {
		base = base[:maxBaseLen]
	}
	re := regexp.MustCompile("[^A-Za-z0-9-]")
	base = re.ReplaceAllString(base, "-")
	return fmt.Sprintf("%s-%s", base, utilrand.String(randomLen))
}
