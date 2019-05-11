package validation

import (
	"github.com/gophercloud/gophercloud/openstack/common/extensions"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/regions"
	netext "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/gophercloud/utils/openstack/clientconfig"
)

type realValidValuesFetcher struct{}

func NewValidValuesFetcher() ValidValuesFetcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return realValidValuesFetcher{}
}
func (f realValidValuesFetcher) GetCloudNames() ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clouds, err := clientconfig.LoadCloudsYAML()
	if err != nil {
		return nil, err
	}
	i := 0
	cloudNames := make([]string, len(clouds))
	for k := range clouds {
		cloudNames[i] = k
		i++
	}
	return cloudNames, nil
}
func (f realValidValuesFetcher) GetRegionNames(cloud string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts := &clientconfig.ClientOpts{Cloud: cloud}
	conn, err := clientconfig.NewServiceClient("identity", opts)
	if err != nil {
		return nil, err
	}
	listOpts := regions.ListOpts{}
	allPages, err := regions.List(conn, listOpts).AllPages()
	if err != nil {
		return nil, err
	}
	allRegions, err := regions.ExtractRegions(allPages)
	if err != nil {
		return nil, err
	}
	regionNames := make([]string, len(allRegions))
	for x, region := range allRegions {
		regionNames[x] = region.ID
	}
	return regionNames, nil
}
func (f realValidValuesFetcher) GetNetworkNames(cloud string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts := &clientconfig.ClientOpts{Cloud: cloud}
	conn, err := clientconfig.NewServiceClient("network", opts)
	if err != nil {
		return nil, err
	}
	listOpts := networks.ListOpts{}
	allPages, err := networks.List(conn, listOpts).AllPages()
	if err != nil {
		return nil, err
	}
	allNetworks, err := networks.ExtractNetworks(allPages)
	if err != nil {
		return nil, err
	}
	networkNames := make([]string, len(allNetworks))
	for x, network := range allNetworks {
		networkNames[x] = network.Name
	}
	return networkNames, nil
}
func (f realValidValuesFetcher) GetFlavorNames(cloud string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts := &clientconfig.ClientOpts{Cloud: cloud}
	conn, err := clientconfig.NewServiceClient("compute", opts)
	if err != nil {
		return nil, err
	}
	listOpts := flavors.ListOpts{}
	allPages, err := flavors.ListDetail(conn, listOpts).AllPages()
	if err != nil {
		return nil, err
	}
	allFlavors, err := flavors.ExtractFlavors(allPages)
	if err != nil {
		return nil, err
	}
	flavorNames := make([]string, len(allFlavors))
	for i, flavor := range allFlavors {
		flavorNames[i] = flavor.Name
	}
	return flavorNames, nil
}
func (f realValidValuesFetcher) GetNetworkExtensionsAliases(cloud string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts := &clientconfig.ClientOpts{Cloud: cloud}
	conn, err := clientconfig.NewServiceClient("network", opts)
	if err != nil {
		return nil, err
	}
	allPages, err := netext.List(conn).AllPages()
	if err != nil {
		return nil, err
	}
	allExts, err := extensions.ExtractExtensions(allPages)
	if err != nil {
		return nil, err
	}
	extAliases := make([]string, len(allExts))
	for i, ext := range allExts {
		extAliases[i] = ext.Alias
	}
	return extAliases, err
}
