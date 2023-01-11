


package v20171001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoleAssignment(ctx *pulumi.Context, args *LookupRoleAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupRoleAssignmentResult, error) {
	var rv LookupRoleAssignmentResult
	err := ctx.Invoke("azure-native:authorization/v20171001preview:getRoleAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoleAssignmentArgs struct {
	RoleAssignmentName string `pulumi:"roleAssignmentName"`
	Scope              string `pulumi:"scope"`
}


type LookupRoleAssignmentResult struct {
	CanDelegate      *bool   `pulumi:"canDelegate"`
	Id               string  `pulumi:"id"`
	Name             string  `pulumi:"name"`
	PrincipalId      *string `pulumi:"principalId"`
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	Scope            *string `pulumi:"scope"`
	Type             string  `pulumi:"type"`
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
	RoleAssignmentName pulumi.StringInput `pulumi:"roleAssignmentName"`
	Scope              pulumi.StringInput `pulumi:"scope"`
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

func (o LookupRoleAssignmentResultOutput) CanDelegate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *bool { return v.CanDelegate }).(pulumi.BoolPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoleAssignmentResultOutput{})
}
