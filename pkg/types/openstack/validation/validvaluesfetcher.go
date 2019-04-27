package validation

type ValidValuesFetcher interface {
	GetCloudNames() ([]string, error)
	GetRegionNames(cloud string) ([]string, error)
	GetNetworkNames(cloud string) ([]string, error)
	GetFlavorNames(cloud string) ([]string, error)
	GetNetworkExtensionsAliases(cloud string) ([]string, error)
}
