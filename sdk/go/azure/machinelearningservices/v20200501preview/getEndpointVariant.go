


package v20200501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEndpointVariant(ctx *pulumi.Context, args *LookupEndpointVariantArgs, opts ...pulumi.InvokeOption) (*LookupEndpointVariantResult, error) {
	var rv LookupEndpointVariantResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200501preview:getEndpointVariant", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointVariantArgs struct {
	Expand            *bool  `pulumi:"expand"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupEndpointVariantResult struct {
	Id         string            `pulumi:"id"`
	Identity   *IdentityResponse `pulumi:"identity"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Sku        *SkuResponse      `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}
