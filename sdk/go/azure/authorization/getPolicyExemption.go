


package authorization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicyExemption(ctx *pulumi.Context, args *LookupPolicyExemptionArgs, opts ...pulumi.InvokeOption) (*LookupPolicyExemptionResult, error) {
	var rv LookupPolicyExemptionResult
	err := ctx.Invoke("azure-native:authorization:getPolicyExemption", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyExemptionArgs struct {
	PolicyExemptionName string `pulumi:"policyExemptionName"`
	Scope               string `pulumi:"scope"`
}


type LookupPolicyExemptionResult struct {
	Description                  *string            `pulumi:"description"`
	DisplayName                  *string            `pulumi:"displayName"`
	ExemptionCategory            string             `pulumi:"exemptionCategory"`
	ExpiresOn                    *string            `pulumi:"expiresOn"`
	Id                           string             `pulumi:"id"`
	Metadata                     interface{}        `pulumi:"metadata"`
	Name                         string             `pulumi:"name"`
	PolicyAssignmentId           string             `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceIds []string           `pulumi:"policyDefinitionReferenceIds"`
	SystemData                   SystemDataResponse `pulumi:"systemData"`
	Type                         string             `pulumi:"type"`
}

func LookupPolicyExemptionOutput(ctx *pulumi.Context, args LookupPolicyExemptionOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyExemptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyExemptionResult, error) {
			args := v.(LookupPolicyExemptionArgs)
			r, err := LookupPolicyExemption(ctx, &args, opts...)
			var s LookupPolicyExemptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyExemptionResultOutput)
}

type LookupPolicyExemptionOutputArgs struct {
	PolicyExemptionName pulumi.StringInput `pulumi:"policyExemptionName"`
	Scope               pulumi.StringInput `pulumi:"scope"`
}

func (LookupPolicyExemptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyExemptionArgs)(nil)).Elem()
}


type LookupPolicyExemptionResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyExemptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyExemptionResult)(nil)).Elem()
}

func (o LookupPolicyExemptionResultOutput) ToLookupPolicyExemptionResultOutput() LookupPolicyExemptionResultOutput {
	return o
}

func (o LookupPolicyExemptionResultOutput) ToLookupPolicyExemptionResultOutputWithContext(ctx context.Context) LookupPolicyExemptionResultOutput {
	return o
}

func (o LookupPolicyExemptionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyExemptionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyExemptionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyExemptionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyExemptionResultOutput) ExemptionCategory() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyExemptionResult) string { return v.ExemptionCategory }).(pulumi.StringOutput)
}

func (o LookupPolicyExemptionResultOutput) ExpiresOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyExemptionResult) *string { return v.ExpiresOn }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyExemptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyExemptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPolicyExemptionResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyExemptionResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupPolicyExemptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyExemptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPolicyExemptionResultOutput) PolicyAssignmentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyExemptionResult) string { return v.PolicyAssignmentId }).(pulumi.StringOutput)
}

func (o LookupPolicyExemptionResultOutput) PolicyDefinitionReferenceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyExemptionResult) []string { return v.PolicyDefinitionReferenceIds }).(pulumi.StringArrayOutput)
}

func (o LookupPolicyExemptionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPolicyExemptionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPolicyExemptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyExemptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyExemptionResultOutput{})
}
