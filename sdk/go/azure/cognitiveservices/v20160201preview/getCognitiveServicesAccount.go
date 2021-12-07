


package v20160201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCognitiveServicesAccount(ctx *pulumi.Context, args *LookupCognitiveServicesAccountArgs, opts ...pulumi.InvokeOption) (*LookupCognitiveServicesAccountResult, error) {
	var rv LookupCognitiveServicesAccountResult
	err := ctx.Invoke("azure-native:cognitiveservices/v20160201preview:getCognitiveServicesAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCognitiveServicesAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCognitiveServicesAccountResult struct {
	Endpoint          *string           `pulumi:"endpoint"`
	Etag              *string           `pulumi:"etag"`
	Id                *string           `pulumi:"id"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	Sku               *SkuResponse      `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
}
