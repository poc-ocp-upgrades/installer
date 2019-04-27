package openstack

import (
	"os"
	"strings"
	"time"
	"github.com/openshift/installer/pkg/destroy"
	"github.com/openshift/installer/pkg/types"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/floatingips"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
	sg "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/trunks"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
	"github.com/gophercloud/gophercloud/openstack/objectstorage/v1/containers"
	"github.com/gophercloud/gophercloud/openstack/objectstorage/v1/objects"
	"github.com/gophercloud/utils/openstack/clientconfig"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/wait"
)

type Filter map[string]string
type ObjectWithTags struct {
	ID	string
	Tags	map[string]string
}
type deleteFunc func(opts *clientconfig.ClientOpts, filter Filter, logger logrus.FieldLogger) (bool, error)
type ClusterUninstaller struct {
	Cloud	string
	Filter	Filter
	Logger	logrus.FieldLogger
}

func (o *ClusterUninstaller) Run() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	deleteFuncs := map[string]deleteFunc{}
	populateDeleteFuncs(deleteFuncs)
	returnChannel := make(chan string)
	opts := &clientconfig.ClientOpts{Cloud: o.Cloud}
	for name, function := range deleteFuncs {
		go deleteRunner(name, function, opts, o.Filter, o.Logger, returnChannel)
	}
	for i := 0; i < len(deleteFuncs); i++ {
		select {
		case res := <-returnChannel:
			o.Logger.Debugf("goroutine %v complete", res)
		}
	}
	return nil
}
func deleteRunner(deleteFuncName string, dFunction deleteFunc, opts *clientconfig.ClientOpts, filter Filter, logger logrus.FieldLogger, channel chan string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	backoffSettings := wait.Backoff{Duration: time.Second * 10, Factor: 1.3, Steps: 100}
	err := wait.ExponentialBackoff(backoffSettings, func() (bool, error) {
		return dFunction(opts, filter, logger)
	})
	if err != nil {
		logger.Fatalf("Unrecoverable error/timed out: %v", err)
		os.Exit(1)
	}
	channel <- deleteFuncName
	return
}
func populateDeleteFuncs(funcs map[string]deleteFunc) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	funcs["deleteServers"] = deleteServers
	funcs["deleteTrunks"] = deleteTrunks
	funcs["deletePorts"] = deletePorts
	funcs["deleteSecurityGroups"] = deleteSecurityGroups
	funcs["deleteRouters"] = deleteRouters
	funcs["deleteSubnets"] = deleteSubnets
	funcs["deleteNetworks"] = deleteNetworks
	funcs["deleteContainers"] = deleteContainers
}
func filterObjects(osObjects []ObjectWithTags, filters Filter) []ObjectWithTags {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	objectsWithTags := []ObjectWithTags{}
	filteredObjects := []ObjectWithTags{}
	for _, object := range osObjects {
		allTagsFound := true
		for key := range filters {
			if _, ok := object.Tags[key]; !ok {
				allTagsFound = false
				break
			}
		}
		if allTagsFound {
			objectsWithTags = append(objectsWithTags, object)
		}
	}
	for _, object := range objectsWithTags {
		valuesMatch := true
		for key, val := range filters {
			if object.Tags[key] != val {
				valuesMatch = false
				break
			}
		}
		if valuesMatch {
			filteredObjects = append(filteredObjects, object)
		}
	}
	return filteredObjects
}
func filterTags(filters Filter) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tags := []string{}
	for k, v := range filters {
		tags = append(tags, strings.Join([]string{k, v}, "="))
	}
	return tags
}
func deleteServers(opts *clientconfig.ClientOpts, filter Filter, logger logrus.FieldLogger) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logger.Debug("Deleting openstack servers")
	defer logger.Debugf("Exiting deleting openstack servers")
	conn, err := clientconfig.NewServiceClient("compute", opts)
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	listOpts := servers.ListOpts{}
	allPages, err := servers.List(conn, listOpts).AllPages()
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	allServers, err := servers.ExtractServers(allPages)
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	serverObjects := []ObjectWithTags{}
	for _, server := range allServers {
		serverObjects = append(serverObjects, ObjectWithTags{ID: server.ID, Tags: server.Metadata})
	}
	filteredServers := filterObjects(serverObjects, filter)
	for _, server := range filteredServers {
		logger.Debugf("Deleting Server: %+v", server.ID)
		err = servers.Delete(conn, server.ID).ExtractErr()
		if err != nil {
			logger.Fatalf("%v", err)
			os.Exit(1)
		}
	}
	return len(filteredServers) == 0, nil
}
func deletePorts(opts *clientconfig.ClientOpts, filter Filter, logger logrus.FieldLogger) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logger.Debug("Deleting openstack ports")
	defer logger.Debugf("Exiting deleting openstack ports")
	conn, err := clientconfig.NewServiceClient("network", opts)
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	tags := filterTags(filter)
	listOpts := ports.ListOpts{TagsAny: strings.Join(tags, ",")}
	allPages, err := ports.List(conn, listOpts).AllPages()
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	allPorts, err := ports.ExtractPorts(allPages)
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	for _, port := range allPorts {
		listOpts := floatingips.ListOpts{PortID: port.ID}
		allPages, err := floatingips.List(conn, listOpts).AllPages()
		if err != nil {
			logger.Fatalf("%v", err)
			os.Exit(1)
		}
		allFIPs, err := floatingips.ExtractFloatingIPs(allPages)
		if err != nil {
			logger.Fatalf("%v", err)
			os.Exit(1)
		}
		for _, fip := range allFIPs {
			logger.Debugf("Deleting Floating IP: %+v", fip.ID)
			err = floatingips.Delete(conn, fip.ID).ExtractErr()
			if err != nil {
				logger.Fatalf("%v", err)
				os.Exit(1)
			}
		}
		logger.Debugf("Deleting Port: %+v", port.ID)
		err = ports.Delete(conn, port.ID).ExtractErr()
		if err != nil {
			return false, nil
		}
	}
	return len(allPorts) == 0, nil
}
func deleteSecurityGroups(opts *clientconfig.ClientOpts, filter Filter, logger logrus.FieldLogger) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logger.Debug("Deleting openstack security-groups")
	defer logger.Debugf("Exiting deleting openstack security-groups")
	conn, err := clientconfig.NewServiceClient("network", opts)
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	tags := filterTags(filter)
	listOpts := sg.ListOpts{TagsAny: strings.Join(tags, ",")}
	allPages, err := sg.List(conn, listOpts).AllPages()
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	allGroups, err := sg.ExtractGroups(allPages)
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	for _, group := range allGroups {
		logger.Debugf("Deleting Security Group: %+v", group.ID)
		err = sg.Delete(conn, group.ID).ExtractErr()
		if err != nil {
			return false, nil
		}
	}
	return len(allGroups) == 0, nil
}
func deleteRouters(opts *clientconfig.ClientOpts, filter Filter, logger logrus.FieldLogger) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logger.Debug("Deleting openstack routers")
	defer logger.Debugf("Exiting deleting openstack routers")
	conn, err := clientconfig.NewServiceClient("network", opts)
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	tags := filterTags(filter)
	listOpts := routers.ListOpts{TagsAny: strings.Join(tags, ",")}
	allPages, err := routers.List(conn, listOpts).AllPages()
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	allRouters, err := routers.ExtractRouters(allPages)
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	for _, router := range allRouters {
		portListOpts := ports.ListOpts{DeviceID: router.ID, DeviceOwner: "network:router_interface"}
		allPagesPort, err := ports.List(conn, portListOpts).AllPages()
		if err != nil {
			logger.Fatalf("%v", err)
			os.Exit(1)
		}
		allPorts, err := ports.ExtractPorts(allPagesPort)
		if err != nil {
			logger.Fatalf("%v", err)
			os.Exit(1)
		}
		HAportListOpts := ports.ListOpts{DeviceID: router.ID, DeviceOwner: "network:ha_router_replicated_interface"}
		HAallPagesPort, err := ports.List(conn, HAportListOpts).AllPages()
		if err != nil {
			logger.Fatalf("%v", err)
			os.Exit(1)
		}
		HAPorts, err := ports.ExtractPorts(HAallPagesPort)
		if err != nil {
			logger.Fatalf("%v", err)
			os.Exit(1)
		}
		allPorts = append(allPorts, HAPorts...)
		for _, port := range allPorts {
			for _, IP := range port.FixedIPs {
				removeOpts := routers.RemoveInterfaceOpts{SubnetID: IP.SubnetID}
				logger.Debugf("Removing Subnet %v from Router %v\n", IP.SubnetID, router.ID)
				_, err = routers.RemoveInterface(conn, router.ID, removeOpts).Extract()
				if err != nil {
					return false, nil
				}
			}
		}
		logger.Debugf("Deleting Router: %+v\n", router.ID)
		err = routers.Delete(conn, router.ID).ExtractErr()
		if err != nil {
			logger.Fatalf("%v", err)
			os.Exit(1)
		}
	}
	return len(allRouters) == 0, nil
}
func deleteSubnets(opts *clientconfig.ClientOpts, filter Filter, logger logrus.FieldLogger) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logger.Debug("Deleting openstack subnets")
	defer logger.Debugf("Exiting deleting openstack subnets")
	conn, err := clientconfig.NewServiceClient("network", opts)
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	tags := filterTags(filter)
	listOpts := subnets.ListOpts{TagsAny: strings.Join(tags, ",")}
	allPages, err := subnets.List(conn, listOpts).AllPages()
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	allSubnets, err := subnets.ExtractSubnets(allPages)
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	for _, subnet := range allSubnets {
		logger.Debugf("Deleting Subnet: %+v", subnet.ID)
		err = subnets.Delete(conn, subnet.ID).ExtractErr()
		if err != nil {
			return false, nil
		}
	}
	return len(allSubnets) == 0, nil
}
func deleteNetworks(opts *clientconfig.ClientOpts, filter Filter, logger logrus.FieldLogger) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logger.Debug("Deleting openstack networks")
	defer logger.Debugf("Exiting deleting openstack networks")
	conn, err := clientconfig.NewServiceClient("network", opts)
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	tags := filterTags(filter)
	listOpts := networks.ListOpts{TagsAny: strings.Join(tags, ",")}
	allPages, err := networks.List(conn, listOpts).AllPages()
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	allNetworks, err := networks.ExtractNetworks(allPages)
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	for _, network := range allNetworks {
		logger.Debugf("Deleting network: %+v", network.ID)
		err = networks.Delete(conn, network.ID).ExtractErr()
		if err != nil {
			return false, nil
		}
	}
	return len(allNetworks) == 0, nil
}
func deleteContainers(opts *clientconfig.ClientOpts, filter Filter, logger logrus.FieldLogger) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logger.Debug("Deleting openstack containers")
	defer logger.Debugf("Exiting deleting openstack containers")
	conn, err := clientconfig.NewServiceClient("object-store", opts)
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	listOpts := containers.ListOpts{Full: false}
	allPages, err := containers.List(conn, listOpts).AllPages()
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	allContainers, err := containers.ExtractNames(allPages)
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	for _, container := range allContainers {
		metadata, err := containers.Get(conn, container, nil).ExtractMetadata()
		if err != nil {
			logger.Fatalf("%v", err)
			os.Exit(1)
		}
		for key, val := range filter {
			titlekey := strings.Title(strings.ToLower(key))
			if metadata[titlekey] == val {
				listOpts := objects.ListOpts{Full: false}
				allPages, err := objects.List(conn, container, listOpts).AllPages()
				if err != nil {
					logger.Fatalf("%v", err)
					os.Exit(1)
				}
				allObjects, err := objects.ExtractNames(allPages)
				if err != nil {
					logger.Fatalf("%v", err)
					os.Exit(1)
				}
				for _, object := range allObjects {
					logger.Debugf("Deleting object: %+v\n", object)
					_, err = objects.Delete(conn, container, object, nil).Extract()
					if err != nil {
						logger.Fatalf("%v", err)
						os.Exit(1)
					}
				}
				logger.Debugf("Deleting container: %+v\n", container)
				_, err = containers.Delete(conn, container).Extract()
				if err != nil {
					logger.Fatalf("%v", err)
					os.Exit(1)
				}
				break
			}
		}
	}
	return true, nil
}
func deleteTrunks(opts *clientconfig.ClientOpts, filter Filter, logger logrus.FieldLogger) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logger.Debug("Deleting openstack trunks")
	defer logger.Debugf("Exiting deleting openstack trunks")
	conn, err := clientconfig.NewServiceClient("network", opts)
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	tags := filterTags(filter)
	listOpts := trunks.ListOpts{TagsAny: strings.Join(tags, ",")}
	allPages, err := trunks.List(conn, listOpts).AllPages()
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	allTrunks, err := trunks.ExtractTrunks(allPages)
	if err != nil {
		logger.Fatalf("%v", err)
		os.Exit(1)
	}
	for _, trunk := range allTrunks {
		logger.Debugf("Deleting Trunk: %+v", trunk.ID)
		err = trunks.Delete(conn, trunk.ID).ExtractErr()
		if err != nil {
			return false, nil
		}
	}
	return len(allTrunks) == 0, nil
}
func New(logger logrus.FieldLogger, metadata *types.ClusterMetadata) (destroy.Destroyer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ClusterUninstaller{Cloud: metadata.ClusterPlatformMetadata.OpenStack.Cloud, Filter: metadata.ClusterPlatformMetadata.OpenStack.Identifier, Logger: logger}, nil
}
