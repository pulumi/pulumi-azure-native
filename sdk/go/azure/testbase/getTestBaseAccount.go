


package testbase

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTestBaseAccount(ctx *pulumi.Context, args *LookupTestBaseAccountArgs, opts ...pulumi.InvokeOption) (*LookupTestBaseAccountResult, error) {
	var rv LookupTestBaseAccountResult
	err := ctx.Invoke("azure-native:testbase:getTestBaseAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTestBaseAccountArgs struct {
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
}


type LookupTestBaseAccountResult struct {
	AccessLevel       string                     `pulumi:"accessLevel"`
	Etag              string                     `pulumi:"etag"`
	Id                string                     `pulumi:"id"`
	Location          string                     `pulumi:"location"`
	Name              string                     `pulumi:"name"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	Sku               TestBaseAccountSKUResponse `pulumi:"sku"`
	SystemData        SystemDataResponse         `pulumi:"systemData"`
	Tags              map[string]string          `pulumi:"tags"`
	Type              string                     `pulumi:"type"`
}
