


package mixedreality

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupObjectAnchorsAccount(ctx *pulumi.Context, args *LookupObjectAnchorsAccountArgs, opts ...pulumi.InvokeOption) (*LookupObjectAnchorsAccountResult, error) {
	var rv LookupObjectAnchorsAccountResult
	err := ctx.Invoke("azure-native:mixedreality:getObjectAnchorsAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupObjectAnchorsAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupObjectAnchorsAccountResult struct {
	AccountDomain      string                                `pulumi:"accountDomain"`
	AccountId          string                                `pulumi:"accountId"`
	Id                 string                                `pulumi:"id"`
	Identity           *ObjectAnchorsAccountResponseIdentity `pulumi:"identity"`
	Kind               *SkuResponse                          `pulumi:"kind"`
	Location           string                                `pulumi:"location"`
	Name               string                                `pulumi:"name"`
	Plan               *IdentityResponse                     `pulumi:"plan"`
	Sku                *SkuResponse                          `pulumi:"sku"`
	StorageAccountName *string                               `pulumi:"storageAccountName"`
	SystemData         SystemDataResponse                    `pulumi:"systemData"`
	Tags               map[string]string                     `pulumi:"tags"`
	Type               string                                `pulumi:"type"`
}
