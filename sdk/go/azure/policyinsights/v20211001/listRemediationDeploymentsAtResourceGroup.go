


package v20211001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRemediationDeploymentsAtResourceGroup(ctx *pulumi.Context, args *ListRemediationDeploymentsAtResourceGroupArgs, opts ...pulumi.InvokeOption) (*ListRemediationDeploymentsAtResourceGroupResult, error) {
	var rv ListRemediationDeploymentsAtResourceGroupResult
	err := ctx.Invoke("azure-native:policyinsights/v20211001:listRemediationDeploymentsAtResourceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRemediationDeploymentsAtResourceGroupArgs struct {
	RemediationName   string `pulumi:"remediationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Top               *int   `pulumi:"top"`
}


type ListRemediationDeploymentsAtResourceGroupResult struct {
	NextLink string                          `pulumi:"nextLink"`
	Value    []RemediationDeploymentResponse `pulumi:"value"`
}
