


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSkusNestedResourceTypeThird(ctx *pulumi.Context, args *LookupSkusNestedResourceTypeThirdArgs, opts ...pulumi.InvokeOption) (*LookupSkusNestedResourceTypeThirdResult, error) {
	var rv LookupSkusNestedResourceTypeThirdResult
	err := ctx.Invoke("azure-native:providerhub/v20210901preview:getSkusNestedResourceTypeThird", args, &rv, opts...)
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
	SystemData SystemDataResponse            `pulumi:"systemData"`
	Type       string                        `pulumi:"type"`
}
