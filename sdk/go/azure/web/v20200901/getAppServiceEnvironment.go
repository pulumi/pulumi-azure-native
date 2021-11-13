


package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAppServiceEnvironment(ctx *pulumi.Context, args *LookupAppServiceEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupAppServiceEnvironmentResult, error) {
	var rv LookupAppServiceEnvironmentResult
	err := ctx.Invoke("azure-native:web/v20200901:getAppServiceEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppServiceEnvironmentArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAppServiceEnvironmentResult struct {
	AllowedMultiSizes          string                              `pulumi:"allowedMultiSizes"`
	AllowedWorkerSizes         string                              `pulumi:"allowedWorkerSizes"`
	ApiManagementAccountId     *string                             `pulumi:"apiManagementAccountId"`
	ClusterSettings            []NameValuePairResponse             `pulumi:"clusterSettings"`
	DatabaseEdition            string                              `pulumi:"databaseEdition"`
	DatabaseServiceObjective   string                              `pulumi:"databaseServiceObjective"`
	DefaultFrontEndScaleFactor int                                 `pulumi:"defaultFrontEndScaleFactor"`
	DnsSuffix                  *string                             `pulumi:"dnsSuffix"`
	DynamicCacheEnabled        *bool                               `pulumi:"dynamicCacheEnabled"`
	EnvironmentCapacities      []StampCapacityResponse             `pulumi:"environmentCapacities"`
	EnvironmentIsHealthy       bool                                `pulumi:"environmentIsHealthy"`
	EnvironmentStatus          string                              `pulumi:"environmentStatus"`
	FrontEndScaleFactor        *int                                `pulumi:"frontEndScaleFactor"`
	HasLinuxWorkers            *bool                               `pulumi:"hasLinuxWorkers"`
	Id                         string                              `pulumi:"id"`
	InternalLoadBalancingMode  *string                             `pulumi:"internalLoadBalancingMode"`
	IpsslAddressCount          *int                                `pulumi:"ipsslAddressCount"`
	Kind                       *string                             `pulumi:"kind"`
	LastAction                 string                              `pulumi:"lastAction"`
	LastActionResult           string                              `pulumi:"lastActionResult"`
	Location                   string                              `pulumi:"location"`
	MaximumNumberOfMachines    int                                 `pulumi:"maximumNumberOfMachines"`
	MultiRoleCount             *int                                `pulumi:"multiRoleCount"`
	MultiSize                  *string                             `pulumi:"multiSize"`
	Name                       string                              `pulumi:"name"`
	NetworkAccessControlList   []NetworkAccessControlEntryResponse `pulumi:"networkAccessControlList"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	ResourceGroup              string                              `pulumi:"resourceGroup"`
	SslCertKeyVaultId          *string                             `pulumi:"sslCertKeyVaultId"`
	SslCertKeyVaultSecretName  *string                             `pulumi:"sslCertKeyVaultSecretName"`
	Status                     string                              `pulumi:"status"`
	SubscriptionId             string                              `pulumi:"subscriptionId"`
	Suspended                  *bool                               `pulumi:"suspended"`
	SystemData                 SystemDataResponse                  `pulumi:"systemData"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
	UpgradeDomains             int                                 `pulumi:"upgradeDomains"`
	UserWhitelistedIpRanges    []string                            `pulumi:"userWhitelistedIpRanges"`
	VipMappings                []VirtualIPMappingResponse          `pulumi:"vipMappings"`
	VirtualNetwork             VirtualNetworkProfileResponse       `pulumi:"virtualNetwork"`
	VnetName                   *string                             `pulumi:"vnetName"`
	VnetResourceGroupName      *string                             `pulumi:"vnetResourceGroupName"`
	VnetSubnetName             *string                             `pulumi:"vnetSubnetName"`
	WorkerPools                []WorkerPoolResponse                `pulumi:"workerPools"`
}
