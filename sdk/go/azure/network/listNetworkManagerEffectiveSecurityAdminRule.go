


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListNetworkManagerEffectiveSecurityAdminRule(ctx *pulumi.Context, args *ListNetworkManagerEffectiveSecurityAdminRuleArgs, opts ...pulumi.InvokeOption) (*ListNetworkManagerEffectiveSecurityAdminRuleResult, error) {
	var rv ListNetworkManagerEffectiveSecurityAdminRuleResult
	err := ctx.Invoke("azure-native:network:listNetworkManagerEffectiveSecurityAdminRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNetworkManagerEffectiveSecurityAdminRuleArgs struct {
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	SkipToken          *string `pulumi:"skipToken"`
	VirtualNetworkName string  `pulumi:"virtualNetworkName"`
}


type ListNetworkManagerEffectiveSecurityAdminRuleResult struct {
	SkipToken *string       `pulumi:"skipToken"`
	Value     []interface{} `pulumi:"value"`
}

func ListNetworkManagerEffectiveSecurityAdminRuleOutput(ctx *pulumi.Context, args ListNetworkManagerEffectiveSecurityAdminRuleOutputArgs, opts ...pulumi.InvokeOption) ListNetworkManagerEffectiveSecurityAdminRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListNetworkManagerEffectiveSecurityAdminRuleResult, error) {
			args := v.(ListNetworkManagerEffectiveSecurityAdminRuleArgs)
			r, err := ListNetworkManagerEffectiveSecurityAdminRule(ctx, &args, opts...)
			var s ListNetworkManagerEffectiveSecurityAdminRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListNetworkManagerEffectiveSecurityAdminRuleResultOutput)
}

type ListNetworkManagerEffectiveSecurityAdminRuleOutputArgs struct {
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
	SkipToken          pulumi.StringPtrInput `pulumi:"skipToken"`
	VirtualNetworkName pulumi.StringInput    `pulumi:"virtualNetworkName"`
}

func (ListNetworkManagerEffectiveSecurityAdminRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNetworkManagerEffectiveSecurityAdminRuleArgs)(nil)).Elem()
}


type ListNetworkManagerEffectiveSecurityAdminRuleResultOutput struct{ *pulumi.OutputState }

func (ListNetworkManagerEffectiveSecurityAdminRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNetworkManagerEffectiveSecurityAdminRuleResult)(nil)).Elem()
}

func (o ListNetworkManagerEffectiveSecurityAdminRuleResultOutput) ToListNetworkManagerEffectiveSecurityAdminRuleResultOutput() ListNetworkManagerEffectiveSecurityAdminRuleResultOutput {
	return o
}

func (o ListNetworkManagerEffectiveSecurityAdminRuleResultOutput) ToListNetworkManagerEffectiveSecurityAdminRuleResultOutputWithContext(ctx context.Context) ListNetworkManagerEffectiveSecurityAdminRuleResultOutput {
	return o
}

func (o ListNetworkManagerEffectiveSecurityAdminRuleResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListNetworkManagerEffectiveSecurityAdminRuleResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

func (o ListNetworkManagerEffectiveSecurityAdminRuleResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v ListNetworkManagerEffectiveSecurityAdminRuleResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListNetworkManagerEffectiveSecurityAdminRuleResultOutput{})
}
