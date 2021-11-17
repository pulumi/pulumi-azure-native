


package v20180701preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRemediationDeploymentsAtResource(ctx *pulumi.Context, args *ListRemediationDeploymentsAtResourceArgs, opts ...pulumi.InvokeOption) (*ListRemediationDeploymentsAtResourceResult, error) {
	var rv ListRemediationDeploymentsAtResourceResult
	err := ctx.Invoke("azure-native:policyinsights/v20180701preview:listRemediationDeploymentsAtResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRemediationDeploymentsAtResourceArgs struct {
	RemediationName string `pulumi:"remediationName"`
	ResourceId      string `pulumi:"resourceId"`
	Top             *int   `pulumi:"top"`
}


type ListRemediationDeploymentsAtResourceResult struct {
	NextLink string                          `pulumi:"nextLink"`
	Value    []RemediationDeploymentResponse `pulumi:"value"`
}
