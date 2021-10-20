


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PolicyAssignmentPropertiesResponse struct {
	Policy         *PolicyAssignmentPropertiesResponsePolicy         `pulumi:"policy"`
	RoleDefinition *PolicyAssignmentPropertiesResponseRoleDefinition `pulumi:"roleDefinition"`
	Scope          *PolicyAssignmentPropertiesResponseScope          `pulumi:"scope"`
}





type PolicyAssignmentPropertiesResponseInput interface {
	pulumi.Input

	ToPolicyAssignmentPropertiesResponseOutput() PolicyAssignmentPropertiesResponseOutput
	ToPolicyAssignmentPropertiesResponseOutputWithContext(context.Context) PolicyAssignmentPropertiesResponseOutput
}

type PolicyAssignmentPropertiesResponseArgs struct {
	Policy         PolicyAssignmentPropertiesResponsePolicyPtrInput         `pulumi:"policy"`
	RoleDefinition PolicyAssignmentPropertiesResponseRoleDefinitionPtrInput `pulumi:"roleDefinition"`
	Scope          PolicyAssignmentPropertiesResponseScopePtrInput          `pulumi:"scope"`
}

func (PolicyAssignmentPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAssignmentPropertiesResponse)(nil)).Elem()
}

func (i PolicyAssignmentPropertiesResponseArgs) ToPolicyAssignmentPropertiesResponseOutput() PolicyAssignmentPropertiesResponseOutput {
	return i.ToPolicyAssignmentPropertiesResponseOutputWithContext(context.Background())
}

func (i PolicyAssignmentPropertiesResponseArgs) ToPolicyAssignmentPropertiesResponseOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAssignmentPropertiesResponseOutput)
}

func (i PolicyAssignmentPropertiesResponseArgs) ToPolicyAssignmentPropertiesResponsePtrOutput() PolicyAssignmentPropertiesResponsePtrOutput {
	return i.ToPolicyAssignmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PolicyAssignmentPropertiesResponseArgs) ToPolicyAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAssignmentPropertiesResponseOutput).ToPolicyAssignmentPropertiesResponsePtrOutputWithContext(ctx)
}









type PolicyAssignmentPropertiesResponsePtrInput interface {
	pulumi.Input

	ToPolicyAssignmentPropertiesResponsePtrOutput() PolicyAssignmentPropertiesResponsePtrOutput
	ToPolicyAssignmentPropertiesResponsePtrOutputWithContext(context.Context) PolicyAssignmentPropertiesResponsePtrOutput
}

type policyAssignmentPropertiesResponsePtrType PolicyAssignmentPropertiesResponseArgs

func PolicyAssignmentPropertiesResponsePtr(v *PolicyAssignmentPropertiesResponseArgs) PolicyAssignmentPropertiesResponsePtrInput {
	return (*policyAssignmentPropertiesResponsePtrType)(v)
}

func (*policyAssignmentPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignmentPropertiesResponse)(nil)).Elem()
}

func (i *policyAssignmentPropertiesResponsePtrType) ToPolicyAssignmentPropertiesResponsePtrOutput() PolicyAssignmentPropertiesResponsePtrOutput {
	return i.ToPolicyAssignmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *policyAssignmentPropertiesResponsePtrType) ToPolicyAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAssignmentPropertiesResponsePtrOutput)
}

type PolicyAssignmentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAssignmentPropertiesResponse)(nil)).Elem()
}

func (o PolicyAssignmentPropertiesResponseOutput) ToPolicyAssignmentPropertiesResponseOutput() PolicyAssignmentPropertiesResponseOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseOutput) ToPolicyAssignmentPropertiesResponseOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseOutput) ToPolicyAssignmentPropertiesResponsePtrOutput() PolicyAssignmentPropertiesResponsePtrOutput {
	return o.ToPolicyAssignmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PolicyAssignmentPropertiesResponseOutput) ToPolicyAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyAssignmentPropertiesResponse) *PolicyAssignmentPropertiesResponse {
		return &v
	}).(PolicyAssignmentPropertiesResponsePtrOutput)
}

func (o PolicyAssignmentPropertiesResponseOutput) Policy() PolicyAssignmentPropertiesResponsePolicyPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponse) *PolicyAssignmentPropertiesResponsePolicy { return v.Policy }).(PolicyAssignmentPropertiesResponsePolicyPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseOutput) RoleDefinition() PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponse) *PolicyAssignmentPropertiesResponseRoleDefinition {
		return v.RoleDefinition
	}).(PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseOutput) Scope() PolicyAssignmentPropertiesResponseScopePtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponse) *PolicyAssignmentPropertiesResponseScope { return v.Scope }).(PolicyAssignmentPropertiesResponseScopePtrOutput)
}

type PolicyAssignmentPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignmentPropertiesResponse)(nil)).Elem()
}

func (o PolicyAssignmentPropertiesResponsePtrOutput) ToPolicyAssignmentPropertiesResponsePtrOutput() PolicyAssignmentPropertiesResponsePtrOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponsePtrOutput) ToPolicyAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponsePtrOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponsePtrOutput) Elem() PolicyAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponse) PolicyAssignmentPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PolicyAssignmentPropertiesResponse
		return ret
	}).(PolicyAssignmentPropertiesResponseOutput)
}

func (o PolicyAssignmentPropertiesResponsePtrOutput) Policy() PolicyAssignmentPropertiesResponsePolicyPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponse) *PolicyAssignmentPropertiesResponsePolicy {
		if v == nil {
			return nil
		}
		return v.Policy
	}).(PolicyAssignmentPropertiesResponsePolicyPtrOutput)
}

func (o PolicyAssignmentPropertiesResponsePtrOutput) RoleDefinition() PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponse) *PolicyAssignmentPropertiesResponseRoleDefinition {
		if v == nil {
			return nil
		}
		return v.RoleDefinition
	}).(PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput)
}

func (o PolicyAssignmentPropertiesResponsePtrOutput) Scope() PolicyAssignmentPropertiesResponseScopePtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponse) *PolicyAssignmentPropertiesResponseScope {
		if v == nil {
			return nil
		}
		return v.Scope
	}).(PolicyAssignmentPropertiesResponseScopePtrOutput)
}

type PolicyAssignmentPropertiesResponsePolicy struct {
	Id                   *string           `pulumi:"id"`
	LastModifiedBy       PrincipalResponse `pulumi:"lastModifiedBy"`
	LastModifiedDateTime *string           `pulumi:"lastModifiedDateTime"`
}





type PolicyAssignmentPropertiesResponsePolicyInput interface {
	pulumi.Input

	ToPolicyAssignmentPropertiesResponsePolicyOutput() PolicyAssignmentPropertiesResponsePolicyOutput
	ToPolicyAssignmentPropertiesResponsePolicyOutputWithContext(context.Context) PolicyAssignmentPropertiesResponsePolicyOutput
}

type PolicyAssignmentPropertiesResponsePolicyArgs struct {
	Id                   pulumi.StringPtrInput  `pulumi:"id"`
	LastModifiedBy       PrincipalResponseInput `pulumi:"lastModifiedBy"`
	LastModifiedDateTime pulumi.StringPtrInput  `pulumi:"lastModifiedDateTime"`
}

func (PolicyAssignmentPropertiesResponsePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAssignmentPropertiesResponsePolicy)(nil)).Elem()
}

func (i PolicyAssignmentPropertiesResponsePolicyArgs) ToPolicyAssignmentPropertiesResponsePolicyOutput() PolicyAssignmentPropertiesResponsePolicyOutput {
	return i.ToPolicyAssignmentPropertiesResponsePolicyOutputWithContext(context.Background())
}

func (i PolicyAssignmentPropertiesResponsePolicyArgs) ToPolicyAssignmentPropertiesResponsePolicyOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponsePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAssignmentPropertiesResponsePolicyOutput)
}

func (i PolicyAssignmentPropertiesResponsePolicyArgs) ToPolicyAssignmentPropertiesResponsePolicyPtrOutput() PolicyAssignmentPropertiesResponsePolicyPtrOutput {
	return i.ToPolicyAssignmentPropertiesResponsePolicyPtrOutputWithContext(context.Background())
}

func (i PolicyAssignmentPropertiesResponsePolicyArgs) ToPolicyAssignmentPropertiesResponsePolicyPtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponsePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAssignmentPropertiesResponsePolicyOutput).ToPolicyAssignmentPropertiesResponsePolicyPtrOutputWithContext(ctx)
}









type PolicyAssignmentPropertiesResponsePolicyPtrInput interface {
	pulumi.Input

	ToPolicyAssignmentPropertiesResponsePolicyPtrOutput() PolicyAssignmentPropertiesResponsePolicyPtrOutput
	ToPolicyAssignmentPropertiesResponsePolicyPtrOutputWithContext(context.Context) PolicyAssignmentPropertiesResponsePolicyPtrOutput
}

type policyAssignmentPropertiesResponsePolicyPtrType PolicyAssignmentPropertiesResponsePolicyArgs

func PolicyAssignmentPropertiesResponsePolicyPtr(v *PolicyAssignmentPropertiesResponsePolicyArgs) PolicyAssignmentPropertiesResponsePolicyPtrInput {
	return (*policyAssignmentPropertiesResponsePolicyPtrType)(v)
}

func (*policyAssignmentPropertiesResponsePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignmentPropertiesResponsePolicy)(nil)).Elem()
}

func (i *policyAssignmentPropertiesResponsePolicyPtrType) ToPolicyAssignmentPropertiesResponsePolicyPtrOutput() PolicyAssignmentPropertiesResponsePolicyPtrOutput {
	return i.ToPolicyAssignmentPropertiesResponsePolicyPtrOutputWithContext(context.Background())
}

func (i *policyAssignmentPropertiesResponsePolicyPtrType) ToPolicyAssignmentPropertiesResponsePolicyPtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponsePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAssignmentPropertiesResponsePolicyPtrOutput)
}

type PolicyAssignmentPropertiesResponsePolicyOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentPropertiesResponsePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAssignmentPropertiesResponsePolicy)(nil)).Elem()
}

func (o PolicyAssignmentPropertiesResponsePolicyOutput) ToPolicyAssignmentPropertiesResponsePolicyOutput() PolicyAssignmentPropertiesResponsePolicyOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponsePolicyOutput) ToPolicyAssignmentPropertiesResponsePolicyOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponsePolicyOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponsePolicyOutput) ToPolicyAssignmentPropertiesResponsePolicyPtrOutput() PolicyAssignmentPropertiesResponsePolicyPtrOutput {
	return o.ToPolicyAssignmentPropertiesResponsePolicyPtrOutputWithContext(context.Background())
}

func (o PolicyAssignmentPropertiesResponsePolicyOutput) ToPolicyAssignmentPropertiesResponsePolicyPtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponsePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyAssignmentPropertiesResponsePolicy) *PolicyAssignmentPropertiesResponsePolicy {
		return &v
	}).(PolicyAssignmentPropertiesResponsePolicyPtrOutput)
}

func (o PolicyAssignmentPropertiesResponsePolicyOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponsePolicy) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponsePolicyOutput) LastModifiedBy() PrincipalResponseOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponsePolicy) PrincipalResponse { return v.LastModifiedBy }).(PrincipalResponseOutput)
}

func (o PolicyAssignmentPropertiesResponsePolicyOutput) LastModifiedDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponsePolicy) *string { return v.LastModifiedDateTime }).(pulumi.StringPtrOutput)
}

type PolicyAssignmentPropertiesResponsePolicyPtrOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentPropertiesResponsePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignmentPropertiesResponsePolicy)(nil)).Elem()
}

func (o PolicyAssignmentPropertiesResponsePolicyPtrOutput) ToPolicyAssignmentPropertiesResponsePolicyPtrOutput() PolicyAssignmentPropertiesResponsePolicyPtrOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponsePolicyPtrOutput) ToPolicyAssignmentPropertiesResponsePolicyPtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponsePolicyPtrOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponsePolicyPtrOutput) Elem() PolicyAssignmentPropertiesResponsePolicyOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponsePolicy) PolicyAssignmentPropertiesResponsePolicy {
		if v != nil {
			return *v
		}
		var ret PolicyAssignmentPropertiesResponsePolicy
		return ret
	}).(PolicyAssignmentPropertiesResponsePolicyOutput)
}

func (o PolicyAssignmentPropertiesResponsePolicyPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponsePolicy) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponsePolicyPtrOutput) LastModifiedBy() PrincipalResponsePtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponsePolicy) *PrincipalResponse {
		if v == nil {
			return nil
		}
		return &v.LastModifiedBy
	}).(PrincipalResponsePtrOutput)
}

func (o PolicyAssignmentPropertiesResponsePolicyPtrOutput) LastModifiedDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponsePolicy) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedDateTime
	}).(pulumi.StringPtrOutput)
}

type PolicyAssignmentPropertiesResponseRoleDefinition struct {
	DisplayName *string `pulumi:"displayName"`
	Id          *string `pulumi:"id"`
	Type        *string `pulumi:"type"`
}





type PolicyAssignmentPropertiesResponseRoleDefinitionInput interface {
	pulumi.Input

	ToPolicyAssignmentPropertiesResponseRoleDefinitionOutput() PolicyAssignmentPropertiesResponseRoleDefinitionOutput
	ToPolicyAssignmentPropertiesResponseRoleDefinitionOutputWithContext(context.Context) PolicyAssignmentPropertiesResponseRoleDefinitionOutput
}

type PolicyAssignmentPropertiesResponseRoleDefinitionArgs struct {
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	Id          pulumi.StringPtrInput `pulumi:"id"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (PolicyAssignmentPropertiesResponseRoleDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAssignmentPropertiesResponseRoleDefinition)(nil)).Elem()
}

func (i PolicyAssignmentPropertiesResponseRoleDefinitionArgs) ToPolicyAssignmentPropertiesResponseRoleDefinitionOutput() PolicyAssignmentPropertiesResponseRoleDefinitionOutput {
	return i.ToPolicyAssignmentPropertiesResponseRoleDefinitionOutputWithContext(context.Background())
}

func (i PolicyAssignmentPropertiesResponseRoleDefinitionArgs) ToPolicyAssignmentPropertiesResponseRoleDefinitionOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseRoleDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAssignmentPropertiesResponseRoleDefinitionOutput)
}

func (i PolicyAssignmentPropertiesResponseRoleDefinitionArgs) ToPolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput() PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput {
	return i.ToPolicyAssignmentPropertiesResponseRoleDefinitionPtrOutputWithContext(context.Background())
}

func (i PolicyAssignmentPropertiesResponseRoleDefinitionArgs) ToPolicyAssignmentPropertiesResponseRoleDefinitionPtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAssignmentPropertiesResponseRoleDefinitionOutput).ToPolicyAssignmentPropertiesResponseRoleDefinitionPtrOutputWithContext(ctx)
}









type PolicyAssignmentPropertiesResponseRoleDefinitionPtrInput interface {
	pulumi.Input

	ToPolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput() PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput
	ToPolicyAssignmentPropertiesResponseRoleDefinitionPtrOutputWithContext(context.Context) PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput
}

type policyAssignmentPropertiesResponseRoleDefinitionPtrType PolicyAssignmentPropertiesResponseRoleDefinitionArgs

func PolicyAssignmentPropertiesResponseRoleDefinitionPtr(v *PolicyAssignmentPropertiesResponseRoleDefinitionArgs) PolicyAssignmentPropertiesResponseRoleDefinitionPtrInput {
	return (*policyAssignmentPropertiesResponseRoleDefinitionPtrType)(v)
}

func (*policyAssignmentPropertiesResponseRoleDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignmentPropertiesResponseRoleDefinition)(nil)).Elem()
}

func (i *policyAssignmentPropertiesResponseRoleDefinitionPtrType) ToPolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput() PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput {
	return i.ToPolicyAssignmentPropertiesResponseRoleDefinitionPtrOutputWithContext(context.Background())
}

func (i *policyAssignmentPropertiesResponseRoleDefinitionPtrType) ToPolicyAssignmentPropertiesResponseRoleDefinitionPtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput)
}

type PolicyAssignmentPropertiesResponseRoleDefinitionOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentPropertiesResponseRoleDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAssignmentPropertiesResponseRoleDefinition)(nil)).Elem()
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionOutput) ToPolicyAssignmentPropertiesResponseRoleDefinitionOutput() PolicyAssignmentPropertiesResponseRoleDefinitionOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionOutput) ToPolicyAssignmentPropertiesResponseRoleDefinitionOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseRoleDefinitionOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionOutput) ToPolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput() PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput {
	return o.ToPolicyAssignmentPropertiesResponseRoleDefinitionPtrOutputWithContext(context.Background())
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionOutput) ToPolicyAssignmentPropertiesResponseRoleDefinitionPtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyAssignmentPropertiesResponseRoleDefinition) *PolicyAssignmentPropertiesResponseRoleDefinition {
		return &v
	}).(PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponseRoleDefinition) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponseRoleDefinition) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponseRoleDefinition) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignmentPropertiesResponseRoleDefinition)(nil)).Elem()
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput) ToPolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput() PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput) ToPolicyAssignmentPropertiesResponseRoleDefinitionPtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput) Elem() PolicyAssignmentPropertiesResponseRoleDefinitionOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponseRoleDefinition) PolicyAssignmentPropertiesResponseRoleDefinition {
		if v != nil {
			return *v
		}
		var ret PolicyAssignmentPropertiesResponseRoleDefinition
		return ret
	}).(PolicyAssignmentPropertiesResponseRoleDefinitionOutput)
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponseRoleDefinition) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponseRoleDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponseRoleDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type PolicyAssignmentPropertiesResponseScope struct {
	DisplayName *string `pulumi:"displayName"`
	Id          *string `pulumi:"id"`
	Type        *string `pulumi:"type"`
}





type PolicyAssignmentPropertiesResponseScopeInput interface {
	pulumi.Input

	ToPolicyAssignmentPropertiesResponseScopeOutput() PolicyAssignmentPropertiesResponseScopeOutput
	ToPolicyAssignmentPropertiesResponseScopeOutputWithContext(context.Context) PolicyAssignmentPropertiesResponseScopeOutput
}

type PolicyAssignmentPropertiesResponseScopeArgs struct {
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	Id          pulumi.StringPtrInput `pulumi:"id"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (PolicyAssignmentPropertiesResponseScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAssignmentPropertiesResponseScope)(nil)).Elem()
}

func (i PolicyAssignmentPropertiesResponseScopeArgs) ToPolicyAssignmentPropertiesResponseScopeOutput() PolicyAssignmentPropertiesResponseScopeOutput {
	return i.ToPolicyAssignmentPropertiesResponseScopeOutputWithContext(context.Background())
}

func (i PolicyAssignmentPropertiesResponseScopeArgs) ToPolicyAssignmentPropertiesResponseScopeOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAssignmentPropertiesResponseScopeOutput)
}

func (i PolicyAssignmentPropertiesResponseScopeArgs) ToPolicyAssignmentPropertiesResponseScopePtrOutput() PolicyAssignmentPropertiesResponseScopePtrOutput {
	return i.ToPolicyAssignmentPropertiesResponseScopePtrOutputWithContext(context.Background())
}

func (i PolicyAssignmentPropertiesResponseScopeArgs) ToPolicyAssignmentPropertiesResponseScopePtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAssignmentPropertiesResponseScopeOutput).ToPolicyAssignmentPropertiesResponseScopePtrOutputWithContext(ctx)
}









type PolicyAssignmentPropertiesResponseScopePtrInput interface {
	pulumi.Input

	ToPolicyAssignmentPropertiesResponseScopePtrOutput() PolicyAssignmentPropertiesResponseScopePtrOutput
	ToPolicyAssignmentPropertiesResponseScopePtrOutputWithContext(context.Context) PolicyAssignmentPropertiesResponseScopePtrOutput
}

type policyAssignmentPropertiesResponseScopePtrType PolicyAssignmentPropertiesResponseScopeArgs

func PolicyAssignmentPropertiesResponseScopePtr(v *PolicyAssignmentPropertiesResponseScopeArgs) PolicyAssignmentPropertiesResponseScopePtrInput {
	return (*policyAssignmentPropertiesResponseScopePtrType)(v)
}

func (*policyAssignmentPropertiesResponseScopePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignmentPropertiesResponseScope)(nil)).Elem()
}

func (i *policyAssignmentPropertiesResponseScopePtrType) ToPolicyAssignmentPropertiesResponseScopePtrOutput() PolicyAssignmentPropertiesResponseScopePtrOutput {
	return i.ToPolicyAssignmentPropertiesResponseScopePtrOutputWithContext(context.Background())
}

func (i *policyAssignmentPropertiesResponseScopePtrType) ToPolicyAssignmentPropertiesResponseScopePtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAssignmentPropertiesResponseScopePtrOutput)
}

type PolicyAssignmentPropertiesResponseScopeOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentPropertiesResponseScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAssignmentPropertiesResponseScope)(nil)).Elem()
}

func (o PolicyAssignmentPropertiesResponseScopeOutput) ToPolicyAssignmentPropertiesResponseScopeOutput() PolicyAssignmentPropertiesResponseScopeOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseScopeOutput) ToPolicyAssignmentPropertiesResponseScopeOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseScopeOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseScopeOutput) ToPolicyAssignmentPropertiesResponseScopePtrOutput() PolicyAssignmentPropertiesResponseScopePtrOutput {
	return o.ToPolicyAssignmentPropertiesResponseScopePtrOutputWithContext(context.Background())
}

func (o PolicyAssignmentPropertiesResponseScopeOutput) ToPolicyAssignmentPropertiesResponseScopePtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseScopePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyAssignmentPropertiesResponseScope) *PolicyAssignmentPropertiesResponseScope {
		return &v
	}).(PolicyAssignmentPropertiesResponseScopePtrOutput)
}

func (o PolicyAssignmentPropertiesResponseScopeOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponseScope) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseScopeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponseScope) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseScopeOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyAssignmentPropertiesResponseScope) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PolicyAssignmentPropertiesResponseScopePtrOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentPropertiesResponseScopePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignmentPropertiesResponseScope)(nil)).Elem()
}

func (o PolicyAssignmentPropertiesResponseScopePtrOutput) ToPolicyAssignmentPropertiesResponseScopePtrOutput() PolicyAssignmentPropertiesResponseScopePtrOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseScopePtrOutput) ToPolicyAssignmentPropertiesResponseScopePtrOutputWithContext(ctx context.Context) PolicyAssignmentPropertiesResponseScopePtrOutput {
	return o
}

func (o PolicyAssignmentPropertiesResponseScopePtrOutput) Elem() PolicyAssignmentPropertiesResponseScopeOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponseScope) PolicyAssignmentPropertiesResponseScope {
		if v != nil {
			return *v
		}
		var ret PolicyAssignmentPropertiesResponseScope
		return ret
	}).(PolicyAssignmentPropertiesResponseScopeOutput)
}

func (o PolicyAssignmentPropertiesResponseScopePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponseScope) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseScopePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponseScope) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentPropertiesResponseScopePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentPropertiesResponseScope) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type PrincipalResponse struct {
	DisplayName *string `pulumi:"displayName"`
	Email       *string `pulumi:"email"`
	Id          *string `pulumi:"id"`
	Type        *string `pulumi:"type"`
}





type PrincipalResponseInput interface {
	pulumi.Input

	ToPrincipalResponseOutput() PrincipalResponseOutput
	ToPrincipalResponseOutputWithContext(context.Context) PrincipalResponseOutput
}

type PrincipalResponseArgs struct {
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	Email       pulumi.StringPtrInput `pulumi:"email"`
	Id          pulumi.StringPtrInput `pulumi:"id"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (PrincipalResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrincipalResponse)(nil)).Elem()
}

func (i PrincipalResponseArgs) ToPrincipalResponseOutput() PrincipalResponseOutput {
	return i.ToPrincipalResponseOutputWithContext(context.Background())
}

func (i PrincipalResponseArgs) ToPrincipalResponseOutputWithContext(ctx context.Context) PrincipalResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrincipalResponseOutput)
}

func (i PrincipalResponseArgs) ToPrincipalResponsePtrOutput() PrincipalResponsePtrOutput {
	return i.ToPrincipalResponsePtrOutputWithContext(context.Background())
}

func (i PrincipalResponseArgs) ToPrincipalResponsePtrOutputWithContext(ctx context.Context) PrincipalResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrincipalResponseOutput).ToPrincipalResponsePtrOutputWithContext(ctx)
}









type PrincipalResponsePtrInput interface {
	pulumi.Input

	ToPrincipalResponsePtrOutput() PrincipalResponsePtrOutput
	ToPrincipalResponsePtrOutputWithContext(context.Context) PrincipalResponsePtrOutput
}

type principalResponsePtrType PrincipalResponseArgs

func PrincipalResponsePtr(v *PrincipalResponseArgs) PrincipalResponsePtrInput {
	return (*principalResponsePtrType)(v)
}

func (*principalResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrincipalResponse)(nil)).Elem()
}

func (i *principalResponsePtrType) ToPrincipalResponsePtrOutput() PrincipalResponsePtrOutput {
	return i.ToPrincipalResponsePtrOutputWithContext(context.Background())
}

func (i *principalResponsePtrType) ToPrincipalResponsePtrOutputWithContext(ctx context.Context) PrincipalResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrincipalResponsePtrOutput)
}

type PrincipalResponseOutput struct{ *pulumi.OutputState }

func (PrincipalResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrincipalResponse)(nil)).Elem()
}

func (o PrincipalResponseOutput) ToPrincipalResponseOutput() PrincipalResponseOutput {
	return o
}

func (o PrincipalResponseOutput) ToPrincipalResponseOutputWithContext(ctx context.Context) PrincipalResponseOutput {
	return o
}

func (o PrincipalResponseOutput) ToPrincipalResponsePtrOutput() PrincipalResponsePtrOutput {
	return o.ToPrincipalResponsePtrOutputWithContext(context.Background())
}

func (o PrincipalResponseOutput) ToPrincipalResponsePtrOutputWithContext(ctx context.Context) PrincipalResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrincipalResponse) *PrincipalResponse {
		return &v
	}).(PrincipalResponsePtrOutput)
}

func (o PrincipalResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrincipalResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PrincipalResponseOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrincipalResponse) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o PrincipalResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrincipalResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PrincipalResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrincipalResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PrincipalResponsePtrOutput struct{ *pulumi.OutputState }

func (PrincipalResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrincipalResponse)(nil)).Elem()
}

func (o PrincipalResponsePtrOutput) ToPrincipalResponsePtrOutput() PrincipalResponsePtrOutput {
	return o
}

func (o PrincipalResponsePtrOutput) ToPrincipalResponsePtrOutputWithContext(ctx context.Context) PrincipalResponsePtrOutput {
	return o
}

func (o PrincipalResponsePtrOutput) Elem() PrincipalResponseOutput {
	return o.ApplyT(func(v *PrincipalResponse) PrincipalResponse {
		if v != nil {
			return *v
		}
		var ret PrincipalResponse
		return ret
	}).(PrincipalResponseOutput)
}

func (o PrincipalResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrincipalResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o PrincipalResponsePtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrincipalResponse) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o PrincipalResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrincipalResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o PrincipalResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrincipalResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAssignmentPropertiesResponseInput)(nil)).Elem(), PolicyAssignmentPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAssignmentPropertiesResponsePtrInput)(nil)).Elem(), PolicyAssignmentPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAssignmentPropertiesResponsePolicyInput)(nil)).Elem(), PolicyAssignmentPropertiesResponsePolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAssignmentPropertiesResponsePolicyPtrInput)(nil)).Elem(), PolicyAssignmentPropertiesResponsePolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAssignmentPropertiesResponseRoleDefinitionInput)(nil)).Elem(), PolicyAssignmentPropertiesResponseRoleDefinitionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAssignmentPropertiesResponseRoleDefinitionPtrInput)(nil)).Elem(), PolicyAssignmentPropertiesResponseRoleDefinitionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAssignmentPropertiesResponseScopeInput)(nil)).Elem(), PolicyAssignmentPropertiesResponseScopeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyAssignmentPropertiesResponseScopePtrInput)(nil)).Elem(), PolicyAssignmentPropertiesResponseScopeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrincipalResponseInput)(nil)).Elem(), PrincipalResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrincipalResponsePtrInput)(nil)).Elem(), PrincipalResponseArgs{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponsePolicyOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponsePolicyPtrOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponseRoleDefinitionOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponseScopeOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponseScopePtrOutput{})
	pulumi.RegisterOutputType(PrincipalResponseOutput{})
	pulumi.RegisterOutputType(PrincipalResponsePtrOutput{})
}
