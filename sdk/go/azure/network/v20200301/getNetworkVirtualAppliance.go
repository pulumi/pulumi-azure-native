


package v20200301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkVirtualAppliance(ctx *pulumi.Context, args *LookupNetworkVirtualApplianceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkVirtualApplianceResult, error) {
	var rv LookupNetworkVirtualApplianceResult
	err := ctx.Invoke("azure-native:network/v20200301:getNetworkVirtualAppliance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkVirtualApplianceArgs struct {
	Expand                      *string `pulumi:"expand"`
	NetworkVirtualApplianceName string  `pulumi:"networkVirtualApplianceName"`
	ResourceGroupName           string  `pulumi:"resourceGroupName"`
}


type LookupNetworkVirtualApplianceResult struct {
	BootStrapConfigurationBlob []string                                `pulumi:"bootStrapConfigurationBlob"`
	CloudInitConfigurationBlob []string                                `pulumi:"cloudInitConfigurationBlob"`
	Etag                       string                                  `pulumi:"etag"`
	Id                         *string                                 `pulumi:"id"`
	Identity                   *ManagedServiceIdentityResponse         `pulumi:"identity"`
	Location                   *string                                 `pulumi:"location"`
	Name                       string                                  `pulumi:"name"`
	ProvisioningState          string                                  `pulumi:"provisioningState"`
	Sku                        *VirtualApplianceSkuPropertiesResponse  `pulumi:"sku"`
	Tags                       map[string]string                       `pulumi:"tags"`
	Type                       string                                  `pulumi:"type"`
	VirtualApplianceAsn        *float64                                `pulumi:"virtualApplianceAsn"`
	VirtualApplianceNics       []VirtualApplianceNicPropertiesResponse `pulumi:"virtualApplianceNics"`
	VirtualHub                 *SubResourceResponse                    `pulumi:"virtualHub"`
}
