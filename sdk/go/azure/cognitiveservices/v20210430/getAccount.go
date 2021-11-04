


package v20210430

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:cognitiveservices/v20210430:getAccount", args, &rv, opts...)
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
