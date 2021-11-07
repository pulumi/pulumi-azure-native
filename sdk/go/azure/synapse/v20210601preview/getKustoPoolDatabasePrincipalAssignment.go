


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKustoPoolDatabasePrincipalAssignment(ctx *pulumi.Context, args *LookupKustoPoolDatabasePrincipalAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupKustoPoolDatabasePrincipalAssignmentResult, error) {
	var rv LookupKustoPoolDatabasePrincipalAssignmentResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getKustoPoolDatabasePrincipalAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoPoolDatabasePrincipalAssignmentArgs struct {
	DatabaseName            string `pulumi:"databaseName"`
	KustoPoolName           string `pulumi:"kustoPoolName"`
	PrincipalAssignmentName string `pulumi:"principalAssignmentName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	WorkspaceName           string `pulumi:"workspaceName"`
}


type LookupKustoPoolDatabasePrincipalAssignmentResult struct {
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
