


package v20150701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Permission struct {
	Actions    []string `pulumi:"actions"`
	NotActions []string `pulumi:"notActions"`
}





type PermissionInput interface {
	pulumi.Input

	ToPermissionOutput() PermissionOutput
	ToPermissionOutputWithContext(context.Context) PermissionOutput
}

type PermissionArgs struct {
	Actions    pulumi.StringArrayInput `pulumi:"actions"`
	NotActions pulumi.StringArrayInput `pulumi:"notActions"`
}

func (PermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Permission)(nil)).Elem()
}

func (i PermissionArgs) ToPermissionOutput() PermissionOutput {
	return i.ToPermissionOutputWithContext(context.Background())
}

func (i PermissionArgs) ToPermissionOutputWithContext(ctx context.Context) PermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionOutput)
}





type PermissionArrayInput interface {
	pulumi.Input

	ToPermissionArrayOutput() PermissionArrayOutput
	ToPermissionArrayOutputWithContext(context.Context) PermissionArrayOutput
}

type PermissionArray []PermissionInput

func (PermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Permission)(nil)).Elem()
}

func (i PermissionArray) ToPermissionArrayOutput() PermissionArrayOutput {
	return i.ToPermissionArrayOutputWithContext(context.Background())
}

func (i PermissionArray) ToPermissionArrayOutputWithContext(ctx context.Context) PermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionArrayOutput)
}

type PermissionOutput struct{ *pulumi.OutputState }

func (PermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Permission)(nil)).Elem()
}

func (o PermissionOutput) ToPermissionOutput() PermissionOutput {
	return o
}

func (o PermissionOutput) ToPermissionOutputWithContext(ctx context.Context) PermissionOutput {
	return o
}

func (o PermissionOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Permission) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

func (o PermissionOutput) NotActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Permission) []string { return v.NotActions }).(pulumi.StringArrayOutput)
}

type PermissionArrayOutput struct{ *pulumi.OutputState }

func (PermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Permission)(nil)).Elem()
}

func (o PermissionArrayOutput) ToPermissionArrayOutput() PermissionArrayOutput {
	return o
}

func (o PermissionArrayOutput) ToPermissionArrayOutputWithContext(ctx context.Context) PermissionArrayOutput {
	return o
}

func (o PermissionArrayOutput) Index(i pulumi.IntInput) PermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Permission {
		return vs[0].([]Permission)[vs[1].(int)]
	}).(PermissionOutput)
}

type PermissionResponse struct {
	Actions    []string `pulumi:"actions"`
	NotActions []string `pulumi:"notActions"`
}

type PermissionResponseOutput struct{ *pulumi.OutputState }

func (PermissionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionResponse)(nil)).Elem()
}

func (o PermissionResponseOutput) ToPermissionResponseOutput() PermissionResponseOutput {
	return o
}

func (o PermissionResponseOutput) ToPermissionResponseOutputWithContext(ctx context.Context) PermissionResponseOutput {
	return o
}

func (o PermissionResponseOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionResponse) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

func (o PermissionResponseOutput) NotActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionResponse) []string { return v.NotActions }).(pulumi.StringArrayOutput)
}

type PermissionResponseArrayOutput struct{ *pulumi.OutputState }

func (PermissionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PermissionResponse)(nil)).Elem()
}

func (o PermissionResponseArrayOutput) ToPermissionResponseArrayOutput() PermissionResponseArrayOutput {
	return o
}

func (o PermissionResponseArrayOutput) ToPermissionResponseArrayOutputWithContext(ctx context.Context) PermissionResponseArrayOutput {
	return o
}

func (o PermissionResponseArrayOutput) Index(i pulumi.IntInput) PermissionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PermissionResponse {
		return vs[0].([]PermissionResponse)[vs[1].(int)]
	}).(PermissionResponseOutput)
}

type RoleAssignmentProperties struct {
	PrincipalId      string `pulumi:"principalId"`
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
}





type RoleAssignmentPropertiesInput interface {
	pulumi.Input

	ToRoleAssignmentPropertiesOutput() RoleAssignmentPropertiesOutput
	ToRoleAssignmentPropertiesOutputWithContext(context.Context) RoleAssignmentPropertiesOutput
}

type RoleAssignmentPropertiesArgs struct {
	PrincipalId      pulumi.StringInput `pulumi:"principalId"`
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
}

func (RoleAssignmentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleAssignmentProperties)(nil)).Elem()
}

func (i RoleAssignmentPropertiesArgs) ToRoleAssignmentPropertiesOutput() RoleAssignmentPropertiesOutput {
	return i.ToRoleAssignmentPropertiesOutputWithContext(context.Background())
}

func (i RoleAssignmentPropertiesArgs) ToRoleAssignmentPropertiesOutputWithContext(ctx context.Context) RoleAssignmentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentPropertiesOutput)
}

type RoleAssignmentPropertiesOutput struct{ *pulumi.OutputState }

func (RoleAssignmentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleAssignmentProperties)(nil)).Elem()
}

func (o RoleAssignmentPropertiesOutput) ToRoleAssignmentPropertiesOutput() RoleAssignmentPropertiesOutput {
	return o
}

func (o RoleAssignmentPropertiesOutput) ToRoleAssignmentPropertiesOutputWithContext(ctx context.Context) RoleAssignmentPropertiesOutput {
	return o
}

func (o RoleAssignmentPropertiesOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v RoleAssignmentProperties) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o RoleAssignmentPropertiesOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v RoleAssignmentProperties) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type RoleAssignmentPropertiesWithScopeResponse struct {
	PrincipalId      *string `pulumi:"principalId"`
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	Scope            *string `pulumi:"scope"`
}

type RoleAssignmentPropertiesWithScopeResponseOutput struct{ *pulumi.OutputState }

func (RoleAssignmentPropertiesWithScopeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleAssignmentPropertiesWithScopeResponse)(nil)).Elem()
}

func (o RoleAssignmentPropertiesWithScopeResponseOutput) ToRoleAssignmentPropertiesWithScopeResponseOutput() RoleAssignmentPropertiesWithScopeResponseOutput {
	return o
}

func (o RoleAssignmentPropertiesWithScopeResponseOutput) ToRoleAssignmentPropertiesWithScopeResponseOutputWithContext(ctx context.Context) RoleAssignmentPropertiesWithScopeResponseOutput {
	return o
}

func (o RoleAssignmentPropertiesWithScopeResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoleAssignmentPropertiesWithScopeResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o RoleAssignmentPropertiesWithScopeResponseOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoleAssignmentPropertiesWithScopeResponse) *string { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

func (o RoleAssignmentPropertiesWithScopeResponseOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoleAssignmentPropertiesWithScopeResponse) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PermissionOutput{})
	pulumi.RegisterOutputType(PermissionArrayOutput{})
	pulumi.RegisterOutputType(PermissionResponseOutput{})
	pulumi.RegisterOutputType(PermissionResponseArrayOutput{})
	pulumi.RegisterOutputType(RoleAssignmentPropertiesOutput{})
	pulumi.RegisterOutputType(RoleAssignmentPropertiesWithScopeResponseOutput{})
}
