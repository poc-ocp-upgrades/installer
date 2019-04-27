package cluster

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/cluster/aws"
	"github.com/openshift/installer/pkg/asset/cluster/azure"
	"github.com/openshift/installer/pkg/asset/cluster/libvirt"
	"github.com/openshift/installer/pkg/asset/cluster/openstack"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/types"
	awstypes "github.com/openshift/installer/pkg/types/aws"
	azuretypes "github.com/openshift/installer/pkg/types/azure"
	libvirttypes "github.com/openshift/installer/pkg/types/libvirt"
	nonetypes "github.com/openshift/installer/pkg/types/none"
	openstacktypes "github.com/openshift/installer/pkg/types/openstack"
	vspheretypes "github.com/openshift/installer/pkg/types/vsphere"
	"github.com/pkg/errors"
)

const (
	metadataFileName = "metadata.json"
)

type Metadata struct{ File *asset.File }

var _ asset.WritableAsset = (*Metadata)(nil)

func (m *Metadata) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Metadata"
}
func (m *Metadata) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&installconfig.ClusterID{}, &installconfig.InstallConfig{}}
}
func (m *Metadata) Generate(parents asset.Parents) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	clusterID := &installconfig.ClusterID{}
	installConfig := &installconfig.InstallConfig{}
	parents.Get(clusterID, installConfig)
	metadata := &types.ClusterMetadata{ClusterName: installConfig.Config.ObjectMeta.Name, ClusterID: clusterID.UUID, InfraID: clusterID.InfraID}
	switch installConfig.Config.Platform.Name() {
	case awstypes.Name:
		metadata.ClusterPlatformMetadata.AWS = aws.Metadata(clusterID.UUID, clusterID.InfraID, installConfig.Config)
	case libvirttypes.Name:
		metadata.ClusterPlatformMetadata.Libvirt = libvirt.Metadata(installConfig.Config)
	case openstacktypes.Name:
		metadata.ClusterPlatformMetadata.OpenStack = openstack.Metadata(clusterID.InfraID, installConfig.Config)
	case azuretypes.Name:
		metadata.ClusterPlatformMetadata.Azure = azure.Metadata(installConfig.Config)
	case nonetypes.Name, vspheretypes.Name:
	default:
		return errors.Errorf("no known platform")
	}
	data, err := json.Marshal(metadata)
	if err != nil {
		return errors.Wrap(err, "failed to Marshal ClusterMetadata")
	}
	m.File = &asset.File{Filename: metadataFileName, Data: data}
	return nil
}
func (m *Metadata) Files() []*asset.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.File != nil {
		return []*asset.File{m.File}
	}
	return []*asset.File{}
}
func (m *Metadata) Load(f asset.FileFetcher) (found bool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false, nil
}
func LoadMetadata(dir string) (*types.ClusterMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	path := filepath.Join(dir, metadataFileName)
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var metadata *types.ClusterMetadata
	if err = json.Unmarshal(raw, &metadata); err != nil {
		return nil, errors.Wrapf(err, "failed to Unmarshal data from %q to types.ClusterMetadata", path)
	}
	return metadata, err
}
