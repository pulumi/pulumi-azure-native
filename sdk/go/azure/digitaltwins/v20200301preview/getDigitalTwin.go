


package v20200301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDigitalTwin(ctx *pulumi.Context, args *LookupDigitalTwinArgs, opts ...pulumi.InvokeOption) (*LookupDigitalTwinResult, error) {
	var rv LookupDigitalTwinResult
	err := ctx.Invoke("azure-native:digitaltwins/v20200301preview:getDigitalTwin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDigitalTwinArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupDigitalTwinResult struct {
	CreatedTime       string                       `pulumi:"createdTime"`
	HostName          string                       `pulumi:"hostName"`
	Id                string                       `pulumi:"id"`
	LastUpdatedTime   string                       `pulumi:"lastUpdatedTime"`
	Location          string                       `pulumi:"location"`
	Name              string                       `pulumi:"name"`
	ProvisioningState string                       `pulumi:"provisioningState"`
	Sku               *DigitalTwinsSkuInfoResponse `pulumi:"sku"`
	Tags              map[string]string            `pulumi:"tags"`
	Type              string                       `pulumi:"type"`
}
