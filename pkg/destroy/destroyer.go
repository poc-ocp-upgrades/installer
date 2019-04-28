package destroy

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/openshift/installer/pkg/asset/cluster"
	"github.com/openshift/installer/pkg/types"
)

type Destroyer interface{ Run() error }
type NewFunc func(logger logrus.FieldLogger, metadata *types.ClusterMetadata) (Destroyer, error)

var Registry = make(map[string]NewFunc)

func New(logger logrus.FieldLogger, rootDir string) (Destroyer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	metadata, err := cluster.LoadMetadata(rootDir)
	if err != nil {
		return nil, err
	}
	platform := metadata.Platform()
	if platform == "" {
		return nil, errors.New("no platform configured in metadata")
	}
	creator, ok := Registry[platform]
	if !ok {
		return nil, errors.Errorf("no destroyers registered for %q", platform)
	}
	return creator(logger, metadata)
}
