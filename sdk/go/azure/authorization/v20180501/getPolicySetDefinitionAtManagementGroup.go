


package v20180501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicySetDefinitionAtManagementGroup(ctx *pulumi.Context, args *LookupPolicySetDefinitionAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupPolicySetDefinitionAtManagementGroupResult, error) {
	var rv LookupPolicySetDefinitionAtManagementGroupResult
	err := ctx.Invoke("azure-native:authorization/v20180501:getPolicySetDefinitionAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicySetDefinitionAtManagementGroupArgs struct {
	ManagementGroupId       string `pulumi:"managementGroupId"`
	PolicySetDefinitionName string `pulumi:"policySetDefinitionName"`
}


type LookupPolicySetDefinitionAtManagementGroupResult struct {
	Description       *string                             `pulumi:"description"`
	DisplayName       *string                             `pulumi:"displayName"`
	Id                string                              `pulumi:"id"`
	Metadata          interface{}                         `pulumi:"metadata"`
	Name              string                              `pulumi:"name"`
	Parameters        interface{}                         `pulumi:"parameters"`
	PolicyDefinitions []PolicyDefinitionReferenceResponse `pulumi:"policyDefinitions"`
	PolicyType        *string                             `pulumi:"policyType"`
	Type              string                              `pulumi:"type"`
}

func LookupPolicySetDefinitionAtManagementGroupOutput(ctx *pulumi.Context, args LookupPolicySetDefinitionAtManagementGroupOutputArgs, opts ...pulumi.InvokeOption) LookupPolicySetDefinitionAtManagementGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicySetDefinitionAtManagementGroupResult, error) {
			args := v.(LookupPolicySetDefinitionAtManagementGroupArgs)
			r, err := LookupPolicySetDefinitionAtManagementGroup(ctx, &args, opts...)
			var s LookupPolicySetDefinitionAtManagementGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicySetDefinitionAtManagementGroupResultOutput)
}

type LookupPolicySetDefinitionAtManagementGroupOutputArgs struct {
	ManagementGroupId       pulumi.StringInput `pulumi:"managementGroupId"`
	PolicySetDefinitionName pulumi.StringInput `pulumi:"policySetDefinitionName"`
}

func (LookupPolicySetDefinitionAtManagementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicySetDefinitionAtManagementGroupArgs)(nil)).Elem()
}


type LookupPolicySetDefinitionAtManagementGroupResultOutput struct{ *pulumi.OutputState }

func (LookupPolicySetDefinitionAtManagementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicySetDefinitionAtManagementGroupResult)(nil)).Elem()
}

func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) ToLookupPolicySetDefinitionAtManagementGroupResultOutput() LookupPolicySetDefinitionAtManagementGroupResultOutput {
	return o
}

func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) ToLookupPolicySetDefinitionAtManagementGroupResultOutputWithContext(ctx context.Context) LookupPolicySetDefinitionAtManagementGroupResultOutput {
	return o
}

func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) PolicyDefinitions() PolicyDefinitionReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) []PolicyDefinitionReferenceResponse {
		return v.PolicyDefinitions
	}).(PolicyDefinitionReferenceResponseArrayOutput)
}

func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

func (o LookupPolicySetDefinitionAtManagementGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetDefinitionAtManagementGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicySetDefinitionAtManagementGroupResultOutput{})
}
