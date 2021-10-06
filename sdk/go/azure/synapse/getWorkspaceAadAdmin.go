


package synapse

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspaceAadAdmin(ctx *pulumi.Context, args *LookupWorkspaceAadAdminArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceAadAdminResult, error) {
	var rv LookupWorkspaceAadAdminResult
	err := ctx.Invoke("azure-native:synapse:getWorkspaceAadAdmin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceAadAdminArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupWorkspaceAadAdminResult struct {
	AdministratorType *string `pulumi:"administratorType"`
	Id                string  `pulumi:"id"`
	Login             *string `pulumi:"login"`
	Name              string  `pulumi:"name"`
	Sid               *string `pulumi:"sid"`
	TenantId          *string `pulumi:"tenantId"`
	Type              string  `pulumi:"type"`
}
