


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListActiveSecurityUserRule(ctx *pulumi.Context, args *ListActiveSecurityUserRuleArgs, opts ...pulumi.InvokeOption) (*ListActiveSecurityUserRuleResult, error) {
	var rv ListActiveSecurityUserRuleResult
	err := ctx.Invoke("azure-native:network:listActiveSecurityUserRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListActiveSecurityUserRuleArgs struct {
	NetworkManagerName string   `pulumi:"networkManagerName"`
	Regions            []string `pulumi:"regions"`
	ResourceGroupName  string   `pulumi:"resourceGroupName"`
	SkipToken          *string  `pulumi:"skipToken"`
}


type ListActiveSecurityUserRuleResult struct {
	SkipToken *string       `pulumi:"skipToken"`
	Value     []interface{} `pulumi:"value"`
}

func ListActiveSecurityUserRuleOutput(ctx *pulumi.Context, args ListActiveSecurityUserRuleOutputArgs, opts ...pulumi.InvokeOption) ListActiveSecurityUserRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListActiveSecurityUserRuleResult, error) {
			args := v.(ListActiveSecurityUserRuleArgs)
			r, err := ListActiveSecurityUserRule(ctx, &args, opts...)
			var s ListActiveSecurityUserRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListActiveSecurityUserRuleResultOutput)
}

type ListActiveSecurityUserRuleOutputArgs struct {
	NetworkManagerName pulumi.StringInput      `pulumi:"networkManagerName"`
	Regions            pulumi.StringArrayInput `pulumi:"regions"`
	ResourceGroupName  pulumi.StringInput      `pulumi:"resourceGroupName"`
	SkipToken          pulumi.StringPtrInput   `pulumi:"skipToken"`
}

func (ListActiveSecurityUserRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveSecurityUserRuleArgs)(nil)).Elem()
}


type ListActiveSecurityUserRuleResultOutput struct{ *pulumi.OutputState }

func (ListActiveSecurityUserRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveSecurityUserRuleResult)(nil)).Elem()
}

func (o ListActiveSecurityUserRuleResultOutput) ToListActiveSecurityUserRuleResultOutput() ListActiveSecurityUserRuleResultOutput {
	return o
}

func (o ListActiveSecurityUserRuleResultOutput) ToListActiveSecurityUserRuleResultOutputWithContext(ctx context.Context) ListActiveSecurityUserRuleResultOutput {
	return o
}

func (o ListActiveSecurityUserRuleResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListActiveSecurityUserRuleResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

func (o ListActiveSecurityUserRuleResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v ListActiveSecurityUserRuleResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListActiveSecurityUserRuleResultOutput{})
}
