


package v20180101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoleDefinition(ctx *pulumi.Context, args *LookupRoleDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupRoleDefinitionResult, error) {
	var rv LookupRoleDefinitionResult
	err := ctx.Invoke("azure-native:authorization/v20180101preview:getRoleDefinition", args, &rv, opts...)
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

func LookupRoleDefinitionOutput(ctx *pulumi.Context, args LookupRoleDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupRoleDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoleDefinitionResult, error) {
			args := v.(LookupRoleDefinitionArgs)
			r, err := LookupRoleDefinition(ctx, &args, opts...)
			var s LookupRoleDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoleDefinitionResultOutput)
}

type LookupRoleDefinitionOutputArgs struct {
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
	Scope            pulumi.StringInput `pulumi:"scope"`
}

func (LookupRoleDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleDefinitionArgs)(nil)).Elem()
}


type LookupRoleDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupRoleDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleDefinitionResult)(nil)).Elem()
}

func (o LookupRoleDefinitionResultOutput) ToLookupRoleDefinitionResultOutput() LookupRoleDefinitionResultOutput {
	return o
}

func (o LookupRoleDefinitionResultOutput) ToLookupRoleDefinitionResultOutputWithContext(ctx context.Context) LookupRoleDefinitionResultOutput {
	return o
}

func (o LookupRoleDefinitionResultOutput) AssignableScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) []string { return v.AssignableScopes }).(pulumi.StringArrayOutput)
}

func (o LookupRoleDefinitionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRoleDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRoleDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRoleDefinitionResultOutput) Permissions() PermissionResponseArrayOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) []PermissionResponse { return v.Permissions }).(PermissionResponseArrayOutput)
}

func (o LookupRoleDefinitionResultOutput) RoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) *string { return v.RoleName }).(pulumi.StringPtrOutput)
}

func (o LookupRoleDefinitionResultOutput) RoleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) *string { return v.RoleType }).(pulumi.StringPtrOutput)
}

func (o LookupRoleDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoleDefinitionResultOutput{})
}
