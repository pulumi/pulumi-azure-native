


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIotHubResource(ctx *pulumi.Context, args *LookupIotHubResourceArgs, opts ...pulumi.InvokeOption) (*LookupIotHubResourceResult, error) {
	var rv LookupIotHubResourceResult
	err := ctx.Invoke("azure-native:devices/v20210201preview:getIotHubResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotHubResourceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupIotHubResourceResult struct {
	Etag       *string                  `pulumi:"etag"`
	Id         string                   `pulumi:"id"`
	Identity   *ArmIdentityResponse     `pulumi:"identity"`
	Location   string                   `pulumi:"location"`
	Name       string                   `pulumi:"name"`
	Properties IotHubPropertiesResponse `pulumi:"properties"`
	Sku        IotHubSkuInfoResponse    `pulumi:"sku"`
	Tags       map[string]string        `pulumi:"tags"`
	Type       string                   `pulumi:"type"`
}
