


package v20190701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRemediationDeploymentsAtManagementGroup(ctx *pulumi.Context, args *ListRemediationDeploymentsAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*ListRemediationDeploymentsAtManagementGroupResult, error) {
	var rv ListRemediationDeploymentsAtManagementGroupResult
	err := ctx.Invoke("azure-native:policyinsights/v20190701:listRemediationDeploymentsAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRemediationDeploymentsAtManagementGroupArgs struct {
	ManagementGroupId         string `pulumi:"managementGroupId"`
	ManagementGroupsNamespace string `pulumi:"managementGroupsNamespace"`
	RemediationName           string `pulumi:"remediationName"`
	Top                       *int   `pulumi:"top"`
}


type ListRemediationDeploymentsAtManagementGroupResult struct {
	NextLink string                          `pulumi:"nextLink"`
	Value    []RemediationDeploymentResponse `pulumi:"value"`
}
