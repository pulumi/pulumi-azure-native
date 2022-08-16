


package v20200301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicyDefinition(ctx *pulumi.Context, args *LookupPolicyDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupPolicyDefinitionResult, error) {
	var rv LookupPolicyDefinitionResult
	err := ctx.Invoke("azure-native:authorization/v20200301:getPolicyDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyDefinitionArgs struct {
	PolicyDefinitionName string `pulumi:"policyDefinitionName"`
}


type LookupPolicyDefinitionResult struct {
	Description *string                                      `pulumi:"description"`
	DisplayName *string                                      `pulumi:"displayName"`
	Id          string                                       `pulumi:"id"`
	Metadata    interface{}                                  `pulumi:"metadata"`
	Mode        *string                                      `pulumi:"mode"`
	Name        string                                       `pulumi:"name"`
	Parameters  map[string]ParameterDefinitionsValueResponse `pulumi:"parameters"`
	PolicyRule  interface{}                                  `pulumi:"policyRule"`
	PolicyType  *string                                      `pulumi:"policyType"`
	Type        string                                       `pulumi:"type"`
}

func LookupPolicyDefinitionOutput(ctx *pulumi.Context, args LookupPolicyDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyDefinitionResult, error) {
			args := v.(LookupPolicyDefinitionArgs)
			r, err := LookupPolicyDefinition(ctx, &args, opts...)
			var s LookupPolicyDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyDefinitionResultOutput)
}

type LookupPolicyDefinitionOutputArgs struct {
	PolicyDefinitionName pulumi.StringInput `pulumi:"policyDefinitionName"`
}

func (LookupPolicyDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyDefinitionArgs)(nil)).Elem()
}


type LookupPolicyDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyDefinitionResult)(nil)).Elem()
}

func (o LookupPolicyDefinitionResultOutput) ToLookupPolicyDefinitionResultOutput() LookupPolicyDefinitionResultOutput {
	return o
}

func (o LookupPolicyDefinitionResultOutput) ToLookupPolicyDefinitionResultOutputWithContext(ctx context.Context) LookupPolicyDefinitionResultOutput {
	return o
}

func (o LookupPolicyDefinitionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyDefinitionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPolicyDefinitionResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupPolicyDefinitionResultOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPolicyDefinitionResultOutput) Parameters() ParameterDefinitionsValueResponseMapOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) map[string]ParameterDefinitionsValueResponse { return v.Parameters }).(ParameterDefinitionsValueResponseMapOutput)
}

func (o LookupPolicyDefinitionResultOutput) PolicyRule() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) interface{} { return v.PolicyRule }).(pulumi.AnyOutput)
}

func (o LookupPolicyDefinitionResultOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyDefinitionResultOutput{})
}
