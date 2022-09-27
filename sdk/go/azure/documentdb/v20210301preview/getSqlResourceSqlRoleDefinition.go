


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlResourceSqlRoleDefinition(ctx *pulumi.Context, args *LookupSqlResourceSqlRoleDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupSqlResourceSqlRoleDefinitionResult, error) {
	var rv LookupSqlResourceSqlRoleDefinitionResult
	err := ctx.Invoke("azure-native:documentdb/v20210301preview:getSqlResourceSqlRoleDefinition", args, &rv, opts...)
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

func LookupSqlResourceSqlRoleDefinitionOutput(ctx *pulumi.Context, args LookupSqlResourceSqlRoleDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupSqlResourceSqlRoleDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlResourceSqlRoleDefinitionResult, error) {
			args := v.(LookupSqlResourceSqlRoleDefinitionArgs)
			r, err := LookupSqlResourceSqlRoleDefinition(ctx, &args, opts...)
			var s LookupSqlResourceSqlRoleDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlResourceSqlRoleDefinitionResultOutput)
}

type LookupSqlResourceSqlRoleDefinitionOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RoleDefinitionId  pulumi.StringInput `pulumi:"roleDefinitionId"`
}

func (LookupSqlResourceSqlRoleDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlRoleDefinitionArgs)(nil)).Elem()
}


type LookupSqlResourceSqlRoleDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupSqlResourceSqlRoleDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlResourceSqlRoleDefinitionResult)(nil)).Elem()
}

func (o LookupSqlResourceSqlRoleDefinitionResultOutput) ToLookupSqlResourceSqlRoleDefinitionResultOutput() LookupSqlResourceSqlRoleDefinitionResultOutput {
	return o
}

func (o LookupSqlResourceSqlRoleDefinitionResultOutput) ToLookupSqlResourceSqlRoleDefinitionResultOutputWithContext(ctx context.Context) LookupSqlResourceSqlRoleDefinitionResultOutput {
	return o
}

func (o LookupSqlResourceSqlRoleDefinitionResultOutput) AssignableScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlRoleDefinitionResult) []string { return v.AssignableScopes }).(pulumi.StringArrayOutput)
}

func (o LookupSqlResourceSqlRoleDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlRoleDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSqlResourceSqlRoleDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlRoleDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSqlResourceSqlRoleDefinitionResultOutput) Permissions() PermissionResponseArrayOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlRoleDefinitionResult) []PermissionResponse { return v.Permissions }).(PermissionResponseArrayOutput)
}

func (o LookupSqlResourceSqlRoleDefinitionResultOutput) RoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlRoleDefinitionResult) *string { return v.RoleName }).(pulumi.StringPtrOutput)
}

func (o LookupSqlResourceSqlRoleDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlResourceSqlRoleDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlResourceSqlRoleDefinitionResultOutput{})
}
