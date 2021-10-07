


package policyinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRemediationDeploymentsAtSubscription(ctx *pulumi.Context, args *ListRemediationDeploymentsAtSubscriptionArgs, opts ...pulumi.InvokeOption) (*ListRemediationDeploymentsAtSubscriptionResult, error) {
	var rv ListRemediationDeploymentsAtSubscriptionResult
	err := ctx.Invoke("azure-native:policyinsights:listRemediationDeploymentsAtSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRemediationDeploymentsAtSubscriptionArgs struct {
	RemediationName string `pulumi:"remediationName"`
	Top             *int   `pulumi:"top"`
}


type ListRemediationDeploymentsAtSubscriptionResult struct {
	NextLink string                          `pulumi:"nextLink"`
	Value    []RemediationDeploymentResponse `pulumi:"value"`
}
