


package v20190228preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSpatialAnchorsAccount(ctx *pulumi.Context, args *LookupSpatialAnchorsAccountArgs, opts ...pulumi.InvokeOption) (*LookupSpatialAnchorsAccountResult, error) {
	var rv LookupSpatialAnchorsAccountResult
	err := ctx.Invoke("azure-native:mixedreality/v20190228preview:getSpatialAnchorsAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSpatialAnchorsAccountArgs struct {
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	SpatialAnchorsAccountName string `pulumi:"spatialAnchorsAccountName"`
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
