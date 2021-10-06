


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspaceSqlAadAdmin(ctx *pulumi.Context, args *LookupWorkspaceSqlAadAdminArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceSqlAadAdminResult, error) {
	var rv LookupWorkspaceSqlAadAdminResult
	err := ctx.Invoke("azure-native:synapse/v20201201:getWorkspaceSqlAadAdmin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceSqlAadAdminArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupWorkspaceSqlAadAdminResult struct {
	AdministratorType *string `pulumi:"administratorType"`
	Id                string  `pulumi:"id"`
	Login             *string `pulumi:"login"`
	Name              string  `pulumi:"name"`
	Sid               *string `pulumi:"sid"`
	TenantId          *string `pulumi:"tenantId"`
	Type              string  `pulumi:"type"`
}
