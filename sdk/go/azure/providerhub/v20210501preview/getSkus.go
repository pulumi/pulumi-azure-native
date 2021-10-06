


package v20210501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSkus(ctx *pulumi.Context, args *LookupSkusArgs, opts ...pulumi.InvokeOption) (*LookupSkusResult, error) {
	var rv LookupSkusResult
	err := ctx.Invoke("azure-native:providerhub/v20210501preview:getSkus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSkusArgs struct {
	ProviderNamespace string `pulumi:"providerNamespace"`
	ResourceType      string `pulumi:"resourceType"`
	Sku               string `pulumi:"sku"`
}

type LookupSkusResult struct {
	Id         string                        `pulumi:"id"`
	Name       string                        `pulumi:"name"`
	Properties SkuResourceResponseProperties `pulumi:"properties"`
	Type       string                        `pulumi:"type"`
}
