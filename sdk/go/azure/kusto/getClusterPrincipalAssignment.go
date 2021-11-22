


package kusto

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupClusterPrincipalAssignment(ctx *pulumi.Context, args *LookupClusterPrincipalAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupClusterPrincipalAssignmentResult, error) {
	var rv LookupClusterPrincipalAssignmentResult
	err := ctx.Invoke("azure-native:kusto:getClusterPrincipalAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterPrincipalAssignmentArgs struct {
	ClusterName             string `pulumi:"clusterName"`
	PrincipalAssignmentName string `pulumi:"principalAssignmentName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupClusterPrincipalAssignmentResult struct {
	Id                string  `pulumi:"id"`
	Name              string  `pulumi:"name"`
	PrincipalId       string  `pulumi:"principalId"`
	PrincipalName     string  `pulumi:"principalName"`
	PrincipalType     string  `pulumi:"principalType"`
	ProvisioningState string  `pulumi:"provisioningState"`
	Role              string  `pulumi:"role"`
	TenantId          *string `pulumi:"tenantId"`
	TenantName        string  `pulumi:"tenantName"`
	Type              string  `pulumi:"type"`
}
