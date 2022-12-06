


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicySetDefinition(ctx *pulumi.Context, args *LookupPolicySetDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupPolicySetDefinitionResult, error) {
	var rv LookupPolicySetDefinitionResult
	err := ctx.Invoke("azure-native:authorization/v20200901:getPolicySetDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicySetDefinitionArgs struct {
	PolicySetDefinitionName string `pulumi:"policySetDefinitionName"`
}


type LookupPolicySetDefinitionResult struct {
	Description            *string                                      `pulumi:"description"`
	DisplayName            *string                                      `pulumi:"displayName"`
	Id                     string                                       `pulumi:"id"`
	Metadata               interface{}                                  `pulumi:"metadata"`
	Name                   string                                       `pulumi:"name"`
	Parameters             map[string]ParameterDefinitionsValueResponse `pulumi:"parameters"`
	PolicyDefinitionGroups []PolicyDefinitionGroupResponse              `pulumi:"policyDefinitionGroups"`
	PolicyDefinitions      []PolicyDefinitionReferenceResponse          `pulumi:"policyDefinitions"`
	PolicyType             *string                                      `pulumi:"policyType"`
	Type                   string                                       `pulumi:"type"`
}

func LookupPolicySetDefinitionOutput(ctx *pulumi.Context, args LookupPolicySetDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupPolicySetDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicySetDefinitionResult, error) {
			args := v.(LookupPolicySetDefinitionArgs)
			r, err := LookupPolicySetDefinition(ctx, &args, opts...)
			var s LookupPolicySetDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicySetDefinitionResultOutput)
}

type LookupPolicySetDefinitionOutputArgs struct {
	PolicySetDefinitionName pulumi.StringInput `pulumi:"policySetDefinitionName"`
}

func (LookupPolicySetDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicySetDefinitionArgs)(nil)).Elem()
}


type LookupPolicySetDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupPolicySetDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicySetDefinitionResult)(nil)).Elem()
}

func (o LookupPolicySetDefinitionResultOutput) ToLookupPolicySetDefinitionResultOutput() LookupPolicySetDefinitionResultOutput {
	return o
}

func (o LookupPolicySetDefinitionResultOutput) ToLookupPolicySetDefinitionResultOutputWithContext(ctx context.Context) LookupPolicySetDefinitionResultOutput {
	return o
}

func (o LookupPolicySetDefinitionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPolicySetDefinitionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupPolicySetDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPolicySetDefinitionResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupPolicySetDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPolicySetDefinitionResultOutput) Parameters() ParameterDefinitionsValueResponseMapOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) map[string]ParameterDefinitionsValueResponse {
		return v.Parameters
	}).(ParameterDefinitionsValueResponseMapOutput)
}

func (o LookupPolicySetDefinitionResultOutput) PolicyDefinitionGroups() PolicyDefinitionGroupResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) []PolicyDefinitionGroupResponse {
		return v.PolicyDefinitionGroups
	}).(PolicyDefinitionGroupResponseArrayOutput)
}

func (o LookupPolicySetDefinitionResultOutput) PolicyDefinitions() PolicyDefinitionReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) []PolicyDefinitionReferenceResponse {
		return v.PolicyDefinitions
	}).(PolicyDefinitionReferenceResponseArrayOutput)
}

func (o LookupPolicySetDefinitionResultOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

func (o LookupPolicySetDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicySetDefinitionResultOutput{})
}
