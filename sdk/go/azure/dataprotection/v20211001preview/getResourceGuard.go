


package v20211001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupResourceGuard(ctx *pulumi.Context, args *LookupResourceGuardArgs, opts ...pulumi.InvokeOption) (*LookupResourceGuardResult, error) {
	var rv LookupResourceGuardResult
	err := ctx.Invoke("azure-native:dataprotection/v20211001preview:getResourceGuard", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceGuardArgs struct {
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	ResourceGuardsName string `pulumi:"resourceGuardsName"`
}

type LookupResourceGuardResult struct {
	ETag       *string                     `pulumi:"eTag"`
	Id         string                      `pulumi:"id"`
	Identity   *DppIdentityDetailsResponse `pulumi:"identity"`
	Location   *string                     `pulumi:"location"`
	Name       string                      `pulumi:"name"`
	Properties ResourceGuardResponse       `pulumi:"properties"`
	SystemData SystemDataResponse          `pulumi:"systemData"`
	Tags       map[string]string           `pulumi:"tags"`
	Type       string                      `pulumi:"type"`
}
