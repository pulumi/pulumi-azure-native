


package v20180701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMediaServiceEdgePolicies(ctx *pulumi.Context, args *ListMediaServiceEdgePoliciesArgs, opts ...pulumi.InvokeOption) (*ListMediaServiceEdgePoliciesResult, error) {
	var rv ListMediaServiceEdgePoliciesResult
	err := ctx.Invoke("azure-native:media/v20180701:listMediaServiceEdgePolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMediaServiceEdgePoliciesArgs struct {
	AccountName       string  `pulumi:"accountName"`
	DeviceId          *string `pulumi:"deviceId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}

type ListMediaServiceEdgePoliciesResult struct {
	UsageDataCollectionPolicy *EdgeUsageDataCollectionPolicyResponse `pulumi:"usageDataCollectionPolicy"`
}
