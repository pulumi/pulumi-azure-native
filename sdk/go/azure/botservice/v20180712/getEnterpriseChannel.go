


package v20180712

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEnterpriseChannel(ctx *pulumi.Context, args *LookupEnterpriseChannelArgs, opts ...pulumi.InvokeOption) (*LookupEnterpriseChannelResult, error) {
	var rv LookupEnterpriseChannelResult
	err := ctx.Invoke("azure-native:botservice/v20180712:getEnterpriseChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnterpriseChannelArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupEnterpriseChannelResult struct {
	Etag       *string                             `pulumi:"etag"`
	Id         string                              `pulumi:"id"`
	Kind       *string                             `pulumi:"kind"`
	Location   *string                             `pulumi:"location"`
	Name       string                              `pulumi:"name"`
	Properties EnterpriseChannelPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse                        `pulumi:"sku"`
	Tags       map[string]string                   `pulumi:"tags"`
	Type       string                              `pulumi:"type"`
}
