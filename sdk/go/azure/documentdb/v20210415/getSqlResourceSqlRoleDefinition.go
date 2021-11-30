


package v20210415

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlResourceSqlRoleDefinition(ctx *pulumi.Context, args *LookupSqlResourceSqlRoleDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupSqlResourceSqlRoleDefinitionResult, error) {
	var rv LookupSqlResourceSqlRoleDefinitionResult
	err := ctx.Invoke("azure-native:documentdb/v20210415:getSqlResourceSqlRoleDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlResourceSqlRoleDefinitionArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RoleDefinitionId  string `pulumi:"roleDefinitionId"`
}


type LookupSqlResourceSqlRoleDefinitionResult struct {
	AssignableScopes []string             `pulumi:"assignableScopes"`
	Id               string               `pulumi:"id"`
	Name             string               `pulumi:"name"`
	Permissions      []PermissionResponse `pulumi:"permissions"`
	RoleName         *string              `pulumi:"roleName"`
	Type             string               `pulumi:"type"`
}
