


package v20150701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoleDefinition(ctx *pulumi.Context, args *LookupRoleDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupRoleDefinitionResult, error) {
	var rv LookupRoleDefinitionResult
	err := ctx.Invoke("azure-native:authorization/v20150701:getRoleDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoleDefinitionArgs struct {
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
	Scope            string `pulumi:"scope"`
}


type LookupRoleDefinitionResult struct {
	AssignableScopes []string             `pulumi:"assignableScopes"`
	Description      *string              `pulumi:"description"`
	Id               string               `pulumi:"id"`
	Name             string               `pulumi:"name"`
	Permissions      []PermissionResponse `pulumi:"permissions"`
	RoleName         *string              `pulumi:"roleName"`
	RoleType         *string              `pulumi:"roleType"`
	Type             string               `pulumi:"type"`
}
