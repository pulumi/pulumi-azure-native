


package v20210401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabasePrincipalAssignment(ctx *pulumi.Context, args *LookupDatabasePrincipalAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupDatabasePrincipalAssignmentResult, error) {
	var rv LookupDatabasePrincipalAssignmentResult
	err := ctx.Invoke("azure-native:synapse/v20210401preview:getDatabasePrincipalAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabasePrincipalAssignmentArgs struct {
	DatabaseName            string `pulumi:"databaseName"`
	KustoPoolName           string `pulumi:"kustoPoolName"`
	PrincipalAssignmentName string `pulumi:"principalAssignmentName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	WorkspaceName           string `pulumi:"workspaceName"`
}


type LookupDatabasePrincipalAssignmentResult struct {
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	PrincipalId       string             `pulumi:"principalId"`
	PrincipalName     string             `pulumi:"principalName"`
	PrincipalType     string             `pulumi:"principalType"`
	ProvisioningState string             `pulumi:"provisioningState"`
	Role              string             `pulumi:"role"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	TenantId          *string            `pulumi:"tenantId"`
	TenantName        string             `pulumi:"tenantName"`
	Type              string             `pulumi:"type"`
}
