package azure

type authConfig struct {
	Cloud				string	`json:"cloud" yaml:"cloud"`
	TenantID			string	`json:"tenantId" yaml:"tenantId"`
	AADClientID			string	`json:"aadClientId" yaml:"aadClientId"`
	AADClientSecret			string	`json:"aadClientSecret" yaml:"aadClientSecret"`
	AADClientCertPath		string	`json:"aadClientCertPath" yaml:"aadClientCertPath"`
	AADClientCertPassword		string	`json:"aadClientCertPassword" yaml:"aadClientCertPassword"`
	UseManagedIdentityExtension	bool	`json:"useManagedIdentityExtension" yaml:"useManagedIdentityExtension"`
	UserAssignedIdentityID		string	`json:"userAssignedIdentityID" yaml:"userAssignedIdentityID"`
	SubscriptionID			string	`json:"subscriptionId" yaml:"subscriptionId"`
}
type config struct {
	authConfig
	ResourceGroup				string	`json:"resourceGroup" yaml:"resourceGroup"`
	Location				string	`json:"location" yaml:"location"`
	VnetName				string	`json:"vnetName" yaml:"vnetName"`
	VnetResourceGroup			string	`json:"vnetResourceGroup" yaml:"vnetResourceGroup"`
	SubnetName				string	`json:"subnetName" yaml:"subnetName"`
	SecurityGroupName			string	`json:"securityGroupName" yaml:"securityGroupName"`
	RouteTableName				string	`json:"routeTableName" yaml:"routeTableName"`
	PrimaryAvailabilitySetName		string	`json:"primaryAvailabilitySetName" yaml:"primaryAvailabilitySetName"`
	VMType					string	`json:"vmType" yaml:"vmType"`
	PrimaryScaleSetName			string	`json:"primaryScaleSetName" yaml:"primaryScaleSetName"`
	CloudProviderBackoff			bool	`json:"cloudProviderBackoff" yaml:"cloudProviderBackoff"`
	CloudProviderBackoffRetries		int	`json:"cloudProviderBackoffRetries" yaml:"cloudProviderBackoffRetries"`
	CloudProviderBackoffExponent		float64	`json:"cloudProviderBackoffExponent" yaml:"cloudProviderBackoffExponent"`
	CloudProviderBackoffDuration		int	`json:"cloudProviderBackoffDuration" yaml:"cloudProviderBackoffDuration"`
	CloudProviderBackoffJitter		float64	`json:"cloudProviderBackoffJitter" yaml:"cloudProviderBackoffJitter"`
	CloudProviderRateLimit			bool	`json:"cloudProviderRateLimit" yaml:"cloudProviderRateLimit"`
	CloudProviderRateLimitQPS		float32	`json:"cloudProviderRateLimitQPS" yaml:"cloudProviderRateLimitQPS"`
	CloudProviderRateLimitBucket		int	`json:"cloudProviderRateLimitBucket" yaml:"cloudProviderRateLimitBucket"`
	CloudProviderRateLimitQPSWrite		float32	`json:"cloudProviderRateLimitQPSWrite" yaml:"cloudProviderRateLimitQPSWrite"`
	CloudProviderRateLimitBucketWrite	int	`json:"cloudProviderRateLimitBucketWrite" yaml:"cloudProviderRateLimitBucketWrite"`
	UseInstanceMetadata			bool	`json:"useInstanceMetadata" yaml:"useInstanceMetadata"`
	LoadBalancerSku				string	`json:"loadBalancerSku" yaml:"loadBalancerSku"`
	ExcludeMasterFromStandardLB		*bool	`json:"excludeMasterFromStandardLB" yaml:"excludeMasterFromStandardLB"`
	DisableOutboundSNAT			*bool	`json:"disableOutboundSNAT" yaml:"disableOutboundSNAT"`
	MaximumLoadBalancerRuleCount		int	`json:"maximumLoadBalancerRuleCount" yaml:"maximumLoadBalancerRuleCount"`
}
