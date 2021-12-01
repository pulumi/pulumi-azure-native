


package v20170119

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIotHubResource(ctx *pulumi.Context, args *LookupIotHubResourceArgs, opts ...pulumi.InvokeOption) (*LookupIotHubResourceResult, error) {
	var rv LookupIotHubResourceResult
	err := ctx.Invoke("azure-native:devices/v20170119:getIotHubResource", args, &rv, opts...)
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
	Etag           *string                  `pulumi:"etag"`
	Id             string                   `pulumi:"id"`
	Location       string                   `pulumi:"location"`
	Name           string                   `pulumi:"name"`
	Properties     IotHubPropertiesResponse `pulumi:"properties"`
	Resourcegroup  string                   `pulumi:"resourcegroup"`
	Sku            IotHubSkuInfoResponse    `pulumi:"sku"`
	Subscriptionid string                   `pulumi:"subscriptionid"`
	Tags           map[string]string        `pulumi:"tags"`
	Type           string                   `pulumi:"type"`
}
