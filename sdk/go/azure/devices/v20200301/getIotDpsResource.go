


package v20200301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIotDpsResource(ctx *pulumi.Context, args *LookupIotDpsResourceArgs, opts ...pulumi.InvokeOption) (*LookupIotDpsResourceResult, error) {
	var rv LookupIotDpsResourceResult
	err := ctx.Invoke("azure-native:devices/v20200301:getIotDpsResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotDpsResourceArgs struct {
	ProvisioningServiceName string `pulumi:"provisioningServiceName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupIotDpsResourceResult struct {
	Etag       *string                             `pulumi:"etag"`
	Id         string                              `pulumi:"id"`
	Location   string                              `pulumi:"location"`
	Name       string                              `pulumi:"name"`
	Properties IotDpsPropertiesDescriptionResponse `pulumi:"properties"`
	Sku        IotDpsSkuInfoResponse               `pulumi:"sku"`
	Tags       map[string]string                   `pulumi:"tags"`
	Type       string                              `pulumi:"type"`
}
