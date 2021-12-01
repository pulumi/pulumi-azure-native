


package v20200301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBastionHost(ctx *pulumi.Context, args *LookupBastionHostArgs, opts ...pulumi.InvokeOption) (*LookupBastionHostResult, error) {
	var rv LookupBastionHostResult
	err := ctx.Invoke("azure-native:network/v20200301:getBastionHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
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
	Tags              map[string]string                    `pulumi:"tags"`
	Type              string                               `pulumi:"type"`
}
