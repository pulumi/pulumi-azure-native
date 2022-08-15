


package v20161201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupPolicyAssignment(ctx *pulumi.Context, args *LookupPolicyAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupPolicyAssignmentResult, error) {
	var rv LookupPolicyAssignmentResult
	err := ctx.Invoke("azure-native:authorization/v20161201:getPolicyAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyAssignmentArgs struct {
	PolicyAssignmentName string `pulumi:"policyAssignmentName"`
	Scope                string `pulumi:"scope"`
}


type LookupPolicyAssignmentResult struct {
	Description        *string     `pulumi:"description"`
	DisplayName        *string     `pulumi:"displayName"`
	Id                 string      `pulumi:"id"`
	Name               *string     `pulumi:"name"`
	Parameters         interface{} `pulumi:"parameters"`
	PolicyDefinitionId *string     `pulumi:"policyDefinitionId"`
	Scope              *string     `pulumi:"scope"`
	Type               *string     `pulumi:"type"`
}

func LookupPolicyAssignmentOutput(ctx *pulumi.Context, args LookupPolicyAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyAssignmentResult, error) {
			args := v.(LookupPolicyAssignmentArgs)
			r, err := LookupPolicyAssignment(ctx, &args, opts...)
			var s LookupPolicyAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyAssignmentResultOutput)
}

type LookupPolicyAssignmentOutputArgs struct {
	PolicyAssignmentName pulumi.StringInput `pulumi:"policyAssignmentName"`
	Scope                pulumi.StringInput `pulumi:"scope"`
}

func (LookupPolicyAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyAssignmentArgs)(nil)).Elem()
}


type LookupPolicyAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyAssignmentResult)(nil)).Elem()
}

func (o LookupPolicyAssignmentResultOutput) ToLookupPolicyAssignmentResultOutput() LookupPolicyAssignmentResultOutput {
	return o
}

func (o LookupPolicyAssignmentResultOutput) ToLookupPolicyAssignmentResultOutputWithContext(ctx context.Context) LookupPolicyAssignmentResultOutput {
	return o
}

func (o LookupPolicyAssignmentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyAssignmentResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPolicyAssignmentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyAssignmentResultOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o LookupPolicyAssignmentResultOutput) PolicyDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *string { return v.PolicyDefinitionId }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyAssignmentResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyAssignmentResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyAssignmentResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyAssignmentResultOutput{})
}
