


package v20201120

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSkusNestedResourceTypeFirst(ctx *pulumi.Context, args *LookupSkusNestedResourceTypeFirstArgs, opts ...pulumi.InvokeOption) (*LookupSkusNestedResourceTypeFirstResult, error) {
	var rv LookupSkusNestedResourceTypeFirstResult
	err := ctx.Invoke("azure-native:providerhub/v20201120:getSkusNestedResourceTypeFirst", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSkusNestedResourceTypeFirstArgs struct {
	NestedResourceTypeFirst string `pulumi:"nestedResourceTypeFirst"`
	ProviderNamespace       string `pulumi:"providerNamespace"`
	ResourceType            string `pulumi:"resourceType"`
	Sku                     string `pulumi:"sku"`
}

type LookupSkusNestedResourceTypeFirstResult struct {
	Id         string                        `pulumi:"id"`
	Name       string                        `pulumi:"name"`
	Properties SkuResourceResponseProperties `pulumi:"properties"`
	Type       string                        `pulumi:"type"`
}
