


package avs

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPlacementPolicy(ctx *pulumi.Context, args *LookupPlacementPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPlacementPolicyResult, error) {
	var rv LookupPlacementPolicyResult
	err := ctx.Invoke("azure-native:avs:getPlacementPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPlacementPolicyArgs struct {
	ClusterName         string `pulumi:"clusterName"`
	PlacementPolicyName string `pulumi:"placementPolicyName"`
	PrivateCloudName    string `pulumi:"privateCloudName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupPlacementPolicyResult struct {
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}
