package installconfig

import (
	"fmt"
	"github.com/gophercloud/utils/openstack/clientconfig"
	"github.com/openshift/installer/pkg/asset"
	awsconfig "github.com/openshift/installer/pkg/asset/installconfig/aws"
	azureconfig "github.com/openshift/installer/pkg/asset/installconfig/azure"
	"github.com/openshift/installer/pkg/types/aws"
	"github.com/openshift/installer/pkg/types/azure"
	"github.com/openshift/installer/pkg/types/libvirt"
	"github.com/openshift/installer/pkg/types/none"
	"github.com/openshift/installer/pkg/types/openstack"
	"github.com/openshift/installer/pkg/types/vsphere"
	"github.com/pkg/errors"
)

type PlatformCredsCheck struct{}

var _ asset.Asset = (*PlatformCredsCheck)(nil)

func (a *PlatformCredsCheck) Dependencies() []asset.Asset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []asset.Asset{&InstallConfig{}}
}
func (a *PlatformCredsCheck) Generate(dependencies asset.Parents) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ic := &InstallConfig{}
	dependencies.Get(ic)
	var err error
	platform := ic.Config.Platform.Name()
	switch platform {
	case aws.Name:
		ssn, err := awsconfig.GetSession()
		if err != nil {
			return errors.Wrap(err, "creating AWS session")
		}
		err = awsconfig.ValidateCreds(ssn)
		if err != nil {
			return errors.Wrap(err, "validate AWS credentials")
		}
	case openstack.Name:
		opts := new(clientconfig.ClientOpts)
		opts.Cloud = ic.Config.Platform.OpenStack.Cloud
		_, err = clientconfig.GetCloudFromYAML(opts)
	case libvirt.Name, none.Name, vsphere.Name:
	case azure.Name:
		_, err = azureconfig.GetSession()
		if err != nil {
			return errors.Wrap(err, "creating Azure session")
		}
	default:
		err = fmt.Errorf("unknown platform type %q", platform)
	}
	return err
}
func (a *PlatformCredsCheck) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "Platform Credentials Check"
}
