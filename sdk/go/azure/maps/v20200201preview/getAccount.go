


package v20200201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:maps/v20200201preview:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccountResult struct {
	Id         string                        `pulumi:"id"`
	Location   string                        `pulumi:"location"`
	Name       string                        `pulumi:"name"`
	Properties MapsAccountPropertiesResponse `pulumi:"properties"`
	Sku        SkuResponse                   `pulumi:"sku"`
	SystemData SystemDataResponse            `pulumi:"systemData"`
	Tags       map[string]string             `pulumi:"tags"`
	Type       string                        `pulumi:"type"`
}
