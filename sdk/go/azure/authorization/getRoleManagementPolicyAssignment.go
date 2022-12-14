


package authorization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoleManagementPolicyAssignment(ctx *pulumi.Context, args *LookupRoleManagementPolicyAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupRoleManagementPolicyAssignmentResult, error) {
	var rv LookupRoleManagementPolicyAssignmentResult
	err := ctx.Invoke("azure-native:authorization:getRoleManagementPolicyAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoleManagementPolicyAssignmentArgs struct {
	RoleManagementPolicyAssignmentName string `pulumi:"roleManagementPolicyAssignmentName"`
	Scope                              string `pulumi:"scope"`
}


type LookupRoleManagementPolicyAssignmentResult struct {
	EffectiveRules             []interface{}                      `pulumi:"effectiveRules"`
	Id                         string                             `pulumi:"id"`
	Name                       string                             `pulumi:"name"`
	PolicyAssignmentProperties PolicyAssignmentPropertiesResponse `pulumi:"policyAssignmentProperties"`
	PolicyId                   *string                            `pulumi:"policyId"`
	RoleDefinitionId           *string                            `pulumi:"roleDefinitionId"`
	Scope                      *string                            `pulumi:"scope"`
	Type                       string                             `pulumi:"type"`
}

func LookupRoleManagementPolicyAssignmentOutput(ctx *pulumi.Context, args LookupRoleManagementPolicyAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupRoleManagementPolicyAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoleManagementPolicyAssignmentResult, error) {
			args := v.(LookupRoleManagementPolicyAssignmentArgs)
			r, err := LookupRoleManagementPolicyAssignment(ctx, &args, opts...)
			var s LookupRoleManagementPolicyAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoleManagementPolicyAssignmentResultOutput)
}

type LookupRoleManagementPolicyAssignmentOutputArgs struct {
	RoleManagementPolicyAssignmentName pulumi.StringInput `pulumi:"roleManagementPolicyAssignmentName"`
	Scope                              pulumi.StringInput `pulumi:"scope"`
}

func (LookupRoleManagementPolicyAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleManagementPolicyAssignmentArgs)(nil)).Elem()
}


type LookupRoleManagementPolicyAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupRoleManagementPolicyAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleManagementPolicyAssignmentResult)(nil)).Elem()
}

func (o LookupRoleManagementPolicyAssignmentResultOutput) ToLookupRoleManagementPolicyAssignmentResultOutput() LookupRoleManagementPolicyAssignmentResultOutput {
	return o
}

func (o LookupRoleManagementPolicyAssignmentResultOutput) ToLookupRoleManagementPolicyAssignmentResultOutputWithContext(ctx context.Context) LookupRoleManagementPolicyAssignmentResultOutput {
	return o
}

func (o LookupRoleManagementPolicyAssignmentResultOutput) EffectiveRules() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupRoleManagementPolicyAssignmentResult) []interface{} { return v.EffectiveRules }).(pulumi.ArrayOutput)
}

func (o LookupRoleManagementPolicyAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleManagementPolicyAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRoleManagementPolicyAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleManagementPolicyAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRoleManagementPolicyAssignmentResultOutput) PolicyAssignmentProperties() PolicyAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v LookupRoleManagementPolicyAssignmentResult) PolicyAssignmentPropertiesResponse {
		return v.PolicyAssignmentProperties
	}).(PolicyAssignmentPropertiesResponseOutput)
}

func (o LookupRoleManagementPolicyAssignmentResultOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleManagementPolicyAssignmentResult) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o LookupRoleManagementPolicyAssignmentResultOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleManagementPolicyAssignmentResult) *string { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

func (o LookupRoleManagementPolicyAssignmentResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleManagementPolicyAssignmentResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o LookupRoleManagementPolicyAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleManagementPolicyAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoleManagementPolicyAssignmentResultOutput{})
}
