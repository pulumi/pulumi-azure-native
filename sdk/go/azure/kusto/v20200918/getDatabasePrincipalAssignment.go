


package v20200918

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabasePrincipalAssignment(ctx *pulumi.Context, args *LookupDatabasePrincipalAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupDatabasePrincipalAssignmentResult, error) {
	var rv LookupDatabasePrincipalAssignmentResult
	err := ctx.Invoke("azure-native:kusto/v20200918:getDatabasePrincipalAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabasePrincipalAssignmentArgs struct {
	ClusterName             string `pulumi:"clusterName"`
	DatabaseName            string `pulumi:"databaseName"`
	PrincipalAssignmentName string `pulumi:"principalAssignmentName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupDatabasePrincipalAssignmentResult struct {
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
