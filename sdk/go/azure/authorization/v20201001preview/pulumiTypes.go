


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

type PolicyAssignmentPropertiesResponsePolicy struct {
	Id                   *string           `pulumi:"id"`
	LastModifiedBy       PrincipalResponse `pulumi:"lastModifiedBy"`
	LastModifiedDateTime *string           `pulumi:"lastModifiedDateTime"`
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
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponsePolicyOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponsePolicyPtrOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponseRoleDefinitionOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponseRoleDefinitionPtrOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponseScopeOutput{})
	pulumi.RegisterOutputType(PolicyAssignmentPropertiesResponseScopePtrOutput{})
	pulumi.RegisterOutputType(PrincipalResponseOutput{})
	pulumi.RegisterOutputType(PrincipalResponsePtrOutput{})
}
