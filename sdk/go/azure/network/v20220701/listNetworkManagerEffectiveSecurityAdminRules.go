


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListNetworkManagerEffectiveSecurityAdminRules(ctx *pulumi.Context, args *ListNetworkManagerEffectiveSecurityAdminRulesArgs, opts ...pulumi.InvokeOption) (*ListNetworkManagerEffectiveSecurityAdminRulesResult, error) {
	var rv ListNetworkManagerEffectiveSecurityAdminRulesResult
	err := ctx.Invoke("azure-native:network/v20220701:listNetworkManagerEffectiveSecurityAdminRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNetworkManagerEffectiveSecurityAdminRulesArgs struct {
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	SkipToken          *string `pulumi:"skipToken"`
	Top                *int    `pulumi:"top"`
	VirtualNetworkName string  `pulumi:"virtualNetworkName"`
}


type ListNetworkManagerEffectiveSecurityAdminRulesResult struct {
	SkipToken *string       `pulumi:"skipToken"`
	Value     []interface{} `pulumi:"value"`
}

func ListNetworkManagerEffectiveSecurityAdminRulesOutput(ctx *pulumi.Context, args ListNetworkManagerEffectiveSecurityAdminRulesOutputArgs, opts ...pulumi.InvokeOption) ListNetworkManagerEffectiveSecurityAdminRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListNetworkManagerEffectiveSecurityAdminRulesResult, error) {
			args := v.(ListNetworkManagerEffectiveSecurityAdminRulesArgs)
			r, err := ListNetworkManagerEffectiveSecurityAdminRules(ctx, &args, opts...)
			var s ListNetworkManagerEffectiveSecurityAdminRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListNetworkManagerEffectiveSecurityAdminRulesResultOutput)
}

type ListNetworkManagerEffectiveSecurityAdminRulesOutputArgs struct {
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
	SkipToken          pulumi.StringPtrInput `pulumi:"skipToken"`
	Top                pulumi.IntPtrInput    `pulumi:"top"`
	VirtualNetworkName pulumi.StringInput    `pulumi:"virtualNetworkName"`
}

func (ListNetworkManagerEffectiveSecurityAdminRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNetworkManagerEffectiveSecurityAdminRulesArgs)(nil)).Elem()
}


type ListNetworkManagerEffectiveSecurityAdminRulesResultOutput struct{ *pulumi.OutputState }

func (ListNetworkManagerEffectiveSecurityAdminRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNetworkManagerEffectiveSecurityAdminRulesResult)(nil)).Elem()
}

func (o ListNetworkManagerEffectiveSecurityAdminRulesResultOutput) ToListNetworkManagerEffectiveSecurityAdminRulesResultOutput() ListNetworkManagerEffectiveSecurityAdminRulesResultOutput {
	return o
}

func (o ListNetworkManagerEffectiveSecurityAdminRulesResultOutput) ToListNetworkManagerEffectiveSecurityAdminRulesResultOutputWithContext(ctx context.Context) ListNetworkManagerEffectiveSecurityAdminRulesResultOutput {
	return o
}

func (o ListNetworkManagerEffectiveSecurityAdminRulesResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListNetworkManagerEffectiveSecurityAdminRulesResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

func (o ListNetworkManagerEffectiveSecurityAdminRulesResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v ListNetworkManagerEffectiveSecurityAdminRulesResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListNetworkManagerEffectiveSecurityAdminRulesResultOutput{})
}
