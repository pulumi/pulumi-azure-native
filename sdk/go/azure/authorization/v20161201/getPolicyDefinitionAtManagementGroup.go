


package v20161201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupPolicyDefinitionAtManagementGroup(ctx *pulumi.Context, args *LookupPolicyDefinitionAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupPolicyDefinitionAtManagementGroupResult, error) {
	var rv LookupPolicyDefinitionAtManagementGroupResult
	err := ctx.Invoke("azure-native:authorization/v20161201:getPolicyDefinitionAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyDefinitionAtManagementGroupArgs struct {
	ManagementGroupId    string `pulumi:"managementGroupId"`
	PolicyDefinitionName string `pulumi:"policyDefinitionName"`
}


type LookupPolicyDefinitionAtManagementGroupResult struct {
	Description *string     `pulumi:"description"`
	DisplayName *string     `pulumi:"displayName"`
	Id          string      `pulumi:"id"`
	Metadata    interface{} `pulumi:"metadata"`
	Mode        *string     `pulumi:"mode"`
	Name        string      `pulumi:"name"`
	Parameters  interface{} `pulumi:"parameters"`
	PolicyRule  interface{} `pulumi:"policyRule"`
	PolicyType  *string     `pulumi:"policyType"`
}

func LookupPolicyDefinitionAtManagementGroupOutput(ctx *pulumi.Context, args LookupPolicyDefinitionAtManagementGroupOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyDefinitionAtManagementGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyDefinitionAtManagementGroupResult, error) {
			args := v.(LookupPolicyDefinitionAtManagementGroupArgs)
			r, err := LookupPolicyDefinitionAtManagementGroup(ctx, &args, opts...)
			var s LookupPolicyDefinitionAtManagementGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyDefinitionAtManagementGroupResultOutput)
}

type LookupPolicyDefinitionAtManagementGroupOutputArgs struct {
	ManagementGroupId    pulumi.StringInput `pulumi:"managementGroupId"`
	PolicyDefinitionName pulumi.StringInput `pulumi:"policyDefinitionName"`
}

func (LookupPolicyDefinitionAtManagementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyDefinitionAtManagementGroupArgs)(nil)).Elem()
}


type LookupPolicyDefinitionAtManagementGroupResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyDefinitionAtManagementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyDefinitionAtManagementGroupResult)(nil)).Elem()
}

func (o LookupPolicyDefinitionAtManagementGroupResultOutput) ToLookupPolicyDefinitionAtManagementGroupResultOutput() LookupPolicyDefinitionAtManagementGroupResultOutput {
	return o
}

func (o LookupPolicyDefinitionAtManagementGroupResultOutput) ToLookupPolicyDefinitionAtManagementGroupResultOutputWithContext(ctx context.Context) LookupPolicyDefinitionAtManagementGroupResultOutput {
	return o
}

func (o LookupPolicyDefinitionAtManagementGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyDefinitionAtManagementGroupResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyDefinitionAtManagementGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPolicyDefinitionAtManagementGroupResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupPolicyDefinitionAtManagementGroupResultOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o LookupPolicyDefinitionAtManagementGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPolicyDefinitionAtManagementGroupResultOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o LookupPolicyDefinitionAtManagementGroupResultOutput) PolicyRule() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) interface{} { return v.PolicyRule }).(pulumi.AnyOutput)
}

func (o LookupPolicyDefinitionAtManagementGroupResultOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyDefinitionAtManagementGroupResult) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyDefinitionAtManagementGroupResultOutput{})
}
