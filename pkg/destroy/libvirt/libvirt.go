package libvirt

import (
	"strings"
	libvirt "github.com/libvirt/libvirt-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/openshift/installer/pkg/destroy"
	"github.com/openshift/installer/pkg/types"
)

type filterFunc func(name string) bool

var ClusterIDPrefixFilter = func(clusterid string) filterFunc {
	if clusterid == "" {
		panic("clusterid cannot be empty")
	}
	return func(name string) bool {
		return strings.HasPrefix(name, clusterid)
	}
}
var AlwaysTrueFilter = func() filterFunc {
	return func(name string) bool {
		return name != "default"
	}
}

type deleteFunc func(conn *libvirt.Connect, filter filterFunc, logger logrus.FieldLogger) error
type ClusterUninstaller struct {
	LibvirtURI	string
	Filter		filterFunc
	Logger		logrus.FieldLogger
}

func (o *ClusterUninstaller) Run() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn, err := libvirt.NewConnect(o.LibvirtURI)
	if err != nil {
		return errors.Wrap(err, "failed to connect to Libvirt daemon")
	}
	for _, del := range []deleteFunc{deleteDomains, deleteNetwork, deleteVolumes} {
		err = del(conn, o.Filter, o.Logger)
		if err != nil {
			return err
		}
	}
	return nil
}
func deleteDomains(conn *libvirt.Connect, filter filterFunc, logger logrus.FieldLogger) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logger.Debug("Deleting libvirt domains")
	var err error
	nothingToDelete := false
	for !nothingToDelete {
		nothingToDelete, err = deleteDomainsSinglePass(conn, filter, logger)
		if err != nil {
			return err
		}
	}
	return nil
}
func deleteDomainsSinglePass(conn *libvirt.Connect, filter filterFunc, logger logrus.FieldLogger) (nothingToDelete bool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	domains, err := conn.ListAllDomains(0)
	if err != nil {
		return false, errors.Wrap(err, "list domains")
	}
	nothingToDelete = true
	for _, domain := range domains {
		defer domain.Free()
		dName, err := domain.GetName()
		if err != nil {
			return false, errors.Wrap(err, "get domain name")
		}
		if !filter(dName) {
			continue
		}
		nothingToDelete = false
		dState, _, err := domain.GetState()
		if err != nil {
			return false, errors.Wrapf(err, "get domain state %d", dName)
		}
		if dState != libvirt.DOMAIN_SHUTOFF && dState != libvirt.DOMAIN_SHUTDOWN {
			if err := domain.Destroy(); err != nil {
				return false, errors.Wrapf(err, "destroy domain %q", dName)
			}
		}
		if err := domain.Undefine(); err != nil {
			return false, errors.Wrapf(err, "undefine domain %q", dName)
		}
		logger.WithField("domain", dName).Info("Deleted domain")
	}
	return nothingToDelete, nil
}
func deleteVolumes(conn *libvirt.Connect, filter filterFunc, logger logrus.FieldLogger) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logger.Debug("Deleting libvirt volumes")
	pools, err := conn.ListStoragePools()
	if err != nil {
		return errors.Wrap(err, "list storage pools")
	}
	tpool := "default"
	for _, pname := range pools {
		if filter(pname) {
			tpool = pname
		}
	}
	pool, err := conn.LookupStoragePoolByName(tpool)
	if err != nil {
		return errors.Wrapf(err, "get storage pool %q", tpool)
	}
	defer pool.Free()
	switch tpool {
	case "default":
		vols, err := pool.ListAllStorageVolumes(0)
		if err != nil {
			return errors.Wrapf(err, "list volumes in %q", tpool)
		}
		for _, vol := range vols {
			defer vol.Free()
			vName, err := vol.GetName()
			if err != nil {
				return errors.Wrapf(err, "get volume names in %q", tpool)
			}
			if !filter(vName) {
				continue
			}
			if err := vol.Delete(0); err != nil {
				return errors.Wrapf(err, "delete volume %q from %q", vName, tpool)
			}
			logger.WithField("volume", vName).Info("Deleted volume")
		}
	default:
		if err := pool.Destroy(); err != nil {
			return errors.Wrapf(err, "destroy pool %q", tpool)
		}
		if err := pool.Undefine(); err != nil {
			return errors.Wrapf(err, "undefine pool %q", tpool)
		}
		logger.WithField("pool", tpool).Info("Deleted pool")
	}
	return nil
}
func deleteNetwork(conn *libvirt.Connect, filter filterFunc, logger logrus.FieldLogger) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logger.Debug("Deleting libvirt network")
	networks, err := conn.ListNetworks()
	if err != nil {
		return errors.Wrap(err, "list networks")
	}
	for _, nName := range networks {
		if !filter(nName) {
			continue
		}
		network, err := conn.LookupNetworkByName(nName)
		if err != nil {
			return errors.Wrapf(err, "get network %q", nName)
		}
		defer network.Free()
		if err := network.Destroy(); err != nil {
			return errors.Wrapf(err, "destroy network %q", nName)
		}
		if err := network.Undefine(); err != nil {
			return errors.Wrapf(err, "undefine network %q", nName)
		}
		logger.WithField("network", nName).Info("Deleted network")
	}
	return nil
}
func New(logger logrus.FieldLogger, metadata *types.ClusterMetadata) (destroy.Destroyer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ClusterUninstaller{LibvirtURI: metadata.ClusterPlatformMetadata.Libvirt.URI, Filter: ClusterIDPrefixFilter(metadata.InfraID), Logger: logger}, nil
}
