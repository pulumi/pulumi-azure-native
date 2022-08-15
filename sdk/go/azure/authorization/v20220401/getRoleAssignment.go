


package v20220401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoleAssignment(ctx *pulumi.Context, args *LookupRoleAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupRoleAssignmentResult, error) {
	var rv LookupRoleAssignmentResult
	err := ctx.Invoke("azure-native:authorization/v20220401:getRoleAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRoleAssignmentArgs struct {
	RoleAssignmentName string  `pulumi:"roleAssignmentName"`
	Scope              string  `pulumi:"scope"`
	TenantId           *string `pulumi:"tenantId"`
}


type LookupRoleAssignmentResult struct {
	Condition                          *string `pulumi:"condition"`
	ConditionVersion                   *string `pulumi:"conditionVersion"`
	CreatedBy                          string  `pulumi:"createdBy"`
	CreatedOn                          string  `pulumi:"createdOn"`
	DelegatedManagedIdentityResourceId *string `pulumi:"delegatedManagedIdentityResourceId"`
	Description                        *string `pulumi:"description"`
	Id                                 string  `pulumi:"id"`
	Name                               string  `pulumi:"name"`
	PrincipalId                        string  `pulumi:"principalId"`
	PrincipalType                      *string `pulumi:"principalType"`
	RoleDefinitionId                   string  `pulumi:"roleDefinitionId"`
	Scope                              string  `pulumi:"scope"`
	Type                               string  `pulumi:"type"`
	UpdatedBy                          string  `pulumi:"updatedBy"`
	UpdatedOn                          string  `pulumi:"updatedOn"`
}


func (val *LookupRoleAssignmentResult) Defaults() *LookupRoleAssignmentResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PrincipalType) {
		principalType_ := "User"
		tmp.PrincipalType = &principalType_
	}
	return &tmp
}

func LookupRoleAssignmentOutput(ctx *pulumi.Context, args LookupRoleAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupRoleAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoleAssignmentResult, error) {
			args := v.(LookupRoleAssignmentArgs)
			r, err := LookupRoleAssignment(ctx, &args, opts...)
			var s LookupRoleAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoleAssignmentResultOutput)
}

type LookupRoleAssignmentOutputArgs struct {
	RoleAssignmentName pulumi.StringInput    `pulumi:"roleAssignmentName"`
	Scope              pulumi.StringInput    `pulumi:"scope"`
	TenantId           pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (LookupRoleAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleAssignmentArgs)(nil)).Elem()
}


type LookupRoleAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupRoleAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleAssignmentResult)(nil)).Elem()
}

func (o LookupRoleAssignmentResultOutput) ToLookupRoleAssignmentResultOutput() LookupRoleAssignmentResultOutput {
	return o
}

func (o LookupRoleAssignmentResultOutput) ToLookupRoleAssignmentResultOutputWithContext(ctx context.Context) LookupRoleAssignmentResultOutput {
	return o
}

func (o LookupRoleAssignmentResultOutput) Condition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.Condition }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) ConditionVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.ConditionVersion }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.CreatedBy }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) DelegatedManagedIdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.DelegatedManagedIdentityResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) PrincipalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.PrincipalType }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.Scope }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) UpdatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.UpdatedBy }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) UpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.UpdatedOn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoleAssignmentResultOutput{})
}
