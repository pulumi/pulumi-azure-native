


package providerhub

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSkusNestedResourceTypeThird(ctx *pulumi.Context, args *LookupSkusNestedResourceTypeThirdArgs, opts ...pulumi.InvokeOption) (*LookupSkusNestedResourceTypeThirdResult, error) {
	var rv LookupSkusNestedResourceTypeThirdResult
	err := ctx.Invoke("azure-native:providerhub:getSkusNestedResourceTypeThird", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSkusNestedResourceTypeThirdArgs struct {
	NestedResourceTypeFirst  string `pulumi:"nestedResourceTypeFirst"`
	NestedResourceTypeSecond string `pulumi:"nestedResourceTypeSecond"`
	NestedResourceTypeThird  string `pulumi:"nestedResourceTypeThird"`
	ProviderNamespace        string `pulumi:"providerNamespace"`
	ResourceType             string `pulumi:"resourceType"`
	Sku                      string `pulumi:"sku"`
}

type LookupSkusNestedResourceTypeThirdResult struct {
	Id         string                        `pulumi:"id"`
	Name       string                        `pulumi:"name"`
	Properties SkuResourceResponseProperties `pulumi:"properties"`
	Type       string                        `pulumi:"type"`
}
