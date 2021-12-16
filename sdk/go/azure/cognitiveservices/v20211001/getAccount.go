


package v20211001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:cognitiveservices/v20211001:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccountResult struct {
	Etag       string                    `pulumi:"etag"`
	Id         string                    `pulumi:"id"`
	Identity   *IdentityResponse         `pulumi:"identity"`
	Kind       *string                   `pulumi:"kind"`
	Location   *string                   `pulumi:"location"`
	Name       string                    `pulumi:"name"`
	Properties AccountPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse              `pulumi:"sku"`
	SystemData SystemDataResponse        `pulumi:"systemData"`
	Tags       map[string]string         `pulumi:"tags"`
	Type       string                    `pulumi:"type"`
}


func (val *LookupAccountResult) Defaults() *LookupAccountResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
