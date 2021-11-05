


package operationalinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupQueryPack(ctx *pulumi.Context, args *LookupQueryPackArgs, opts ...pulumi.InvokeOption) (*LookupQueryPackResult, error) {
	var rv LookupQueryPackResult
	err := ctx.Invoke("azure-native:operationalinsights:getQueryPack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueryPackArgs struct {
	QueryPackName     string `pulumi:"queryPackName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupQueryPackResult struct {
	Id                string            `pulumi:"id"`
	Location          string            `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	QueryPackId       string            `pulumi:"queryPackId"`
	Tags              map[string]string `pulumi:"tags"`
	TimeCreated       string            `pulumi:"timeCreated"`
	TimeModified      string            `pulumi:"timeModified"`
	Type              string            `pulumi:"type"`
}
