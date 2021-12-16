


package v20210201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBastionHost(ctx *pulumi.Context, args *LookupBastionHostArgs, opts ...pulumi.InvokeOption) (*LookupBastionHostResult, error) {
	var rv LookupBastionHostResult
	err := ctx.Invoke("azure-native:network/v20210201:getBastionHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupBastionHostArgs struct {
	BastionHostName   string `pulumi:"bastionHostName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBastionHostResult struct {
	DnsName           *string                              `pulumi:"dnsName"`
	Etag              string                               `pulumi:"etag"`
	Id                *string                              `pulumi:"id"`
	IpConfigurations  []BastionHostIPConfigurationResponse `pulumi:"ipConfigurations"`
	Location          *string                              `pulumi:"location"`
	Name              string                               `pulumi:"name"`
	ProvisioningState string                               `pulumi:"provisioningState"`
	Sku               *SkuResponse                         `pulumi:"sku"`
	Tags              map[string]string                    `pulumi:"tags"`
	Type              string                               `pulumi:"type"`
}


func (val *LookupBastionHostResult) Defaults() *LookupBastionHostResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
}
