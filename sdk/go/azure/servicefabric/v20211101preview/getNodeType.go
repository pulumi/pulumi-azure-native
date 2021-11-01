


package v20211101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNodeType(ctx *pulumi.Context, args *LookupNodeTypeArgs, opts ...pulumi.InvokeOption) (*LookupNodeTypeResult, error) {
	var rv LookupNodeTypeResult
	err := ctx.Invoke("azure-native:servicefabric/v20211101preview:getNodeType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNodeTypeArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	NodeTypeName      string `pulumi:"nodeTypeName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupNodeTypeResult struct {
	AdditionalDataDisks          []VmssDataDiskResponse            `pulumi:"additionalDataDisks"`
	ApplicationPorts             *EndpointRangeDescriptionResponse `pulumi:"applicationPorts"`
	Capacities                   map[string]string                 `pulumi:"capacities"`
	DataDiskLetter               *string                           `pulumi:"dataDiskLetter"`
	DataDiskSizeGB               int                               `pulumi:"dataDiskSizeGB"`
	DataDiskType                 *string                           `pulumi:"dataDiskType"`
	EnableAcceleratedNetworking  *bool                             `pulumi:"enableAcceleratedNetworking"`
	EnableEncryptionAtHost       *bool                             `pulumi:"enableEncryptionAtHost"`
	EphemeralPorts               *EndpointRangeDescriptionResponse `pulumi:"ephemeralPorts"`
	FrontendConfigurations       []FrontendConfigurationResponse   `pulumi:"frontendConfigurations"`
	Id                           string                            `pulumi:"id"`
	IsPrimary                    bool                              `pulumi:"isPrimary"`
	IsStateless                  *bool                             `pulumi:"isStateless"`
	MultiplePlacementGroups      *bool                             `pulumi:"multiplePlacementGroups"`
	Name                         string                            `pulumi:"name"`
	NetworkSecurityRules         []NetworkSecurityRuleResponse     `pulumi:"networkSecurityRules"`
	PlacementProperties          map[string]string                 `pulumi:"placementProperties"`
	ProvisioningState            string                            `pulumi:"provisioningState"`
	Sku                          *NodeTypeSkuResponse              `pulumi:"sku"`
	SystemData                   SystemDataResponse                `pulumi:"systemData"`
	Tags                         map[string]string                 `pulumi:"tags"`
	Type                         string                            `pulumi:"type"`
	UseDefaultPublicLoadBalancer *bool                             `pulumi:"useDefaultPublicLoadBalancer"`
	VmExtensions                 []VMSSExtensionResponse           `pulumi:"vmExtensions"`
	VmImageOffer                 *string                           `pulumi:"vmImageOffer"`
	VmImagePublisher             *string                           `pulumi:"vmImagePublisher"`
	VmImageSku                   *string                           `pulumi:"vmImageSku"`
	VmImageVersion               *string                           `pulumi:"vmImageVersion"`
	VmInstanceCount              int                               `pulumi:"vmInstanceCount"`
	VmManagedIdentity            *VmManagedIdentityResponse        `pulumi:"vmManagedIdentity"`
	VmSecrets                    []VaultSecretGroupResponse        `pulumi:"vmSecrets"`
	VmSize                       *string                           `pulumi:"vmSize"`
}
