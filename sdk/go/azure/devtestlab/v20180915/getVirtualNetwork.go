


package v20180915

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetwork(ctx *pulumi.Context, args *LookupVirtualNetworkArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkResult, error) {
	var rv LookupVirtualNetworkResult
	err := ctx.Invoke("azure-native:devtestlab/v20180915:getVirtualNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupVirtualNetworkResult struct {
	AllowedSubnets             []SubnetResponse         `pulumi:"allowedSubnets"`
	CreatedDate                string                   `pulumi:"createdDate"`
	Description                *string                  `pulumi:"description"`
	ExternalProviderResourceId *string                  `pulumi:"externalProviderResourceId"`
	ExternalSubnets            []ExternalSubnetResponse `pulumi:"externalSubnets"`
	Id                         string                   `pulumi:"id"`
	Location                   *string                  `pulumi:"location"`
	Name                       string                   `pulumi:"name"`
	ProvisioningState          string                   `pulumi:"provisioningState"`
	SubnetOverrides            []SubnetOverrideResponse `pulumi:"subnetOverrides"`
	Tags                       map[string]string        `pulumi:"tags"`
	Type                       string                   `pulumi:"type"`
	UniqueIdentifier           string                   `pulumi:"uniqueIdentifier"`
}
