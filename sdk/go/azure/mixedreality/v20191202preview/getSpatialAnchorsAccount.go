


package v20191202preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSpatialAnchorsAccount(ctx *pulumi.Context, args *LookupSpatialAnchorsAccountArgs, opts ...pulumi.InvokeOption) (*LookupSpatialAnchorsAccountResult, error) {
	var rv LookupSpatialAnchorsAccountResult
	err := ctx.Invoke("azure-native:mixedreality/v20191202preview:getSpatialAnchorsAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSpatialAnchorsAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSpatialAnchorsAccountResult struct {
	AccountDomain string            `pulumi:"accountDomain"`
	AccountId     string            `pulumi:"accountId"`
	Id            string            `pulumi:"id"`
	Identity      *IdentityResponse `pulumi:"identity"`
	Location      string            `pulumi:"location"`
	Name          string            `pulumi:"name"`
	Tags          map[string]string `pulumi:"tags"`
	Type          string            `pulumi:"type"`
}
