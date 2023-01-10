


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListActiveSecurityAdminRule(ctx *pulumi.Context, args *ListActiveSecurityAdminRuleArgs, opts ...pulumi.InvokeOption) (*ListActiveSecurityAdminRuleResult, error) {
	var rv ListActiveSecurityAdminRuleResult
	err := ctx.Invoke("azure-native:network:listActiveSecurityAdminRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListActiveSecurityAdminRuleArgs struct {
	NetworkManagerName string   `pulumi:"networkManagerName"`
	Regions            []string `pulumi:"regions"`
	ResourceGroupName  string   `pulumi:"resourceGroupName"`
	SkipToken          *string  `pulumi:"skipToken"`
}


type ListActiveSecurityAdminRuleResult struct {
	SkipToken *string       `pulumi:"skipToken"`
	Value     []interface{} `pulumi:"value"`
}

func ListActiveSecurityAdminRuleOutput(ctx *pulumi.Context, args ListActiveSecurityAdminRuleOutputArgs, opts ...pulumi.InvokeOption) ListActiveSecurityAdminRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListActiveSecurityAdminRuleResult, error) {
			args := v.(ListActiveSecurityAdminRuleArgs)
			r, err := ListActiveSecurityAdminRule(ctx, &args, opts...)
			var s ListActiveSecurityAdminRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListActiveSecurityAdminRuleResultOutput)
}

type ListActiveSecurityAdminRuleOutputArgs struct {
	NetworkManagerName pulumi.StringInput      `pulumi:"networkManagerName"`
	Regions            pulumi.StringArrayInput `pulumi:"regions"`
	ResourceGroupName  pulumi.StringInput      `pulumi:"resourceGroupName"`
	SkipToken          pulumi.StringPtrInput   `pulumi:"skipToken"`
}

func (ListActiveSecurityAdminRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveSecurityAdminRuleArgs)(nil)).Elem()
}


type ListActiveSecurityAdminRuleResultOutput struct{ *pulumi.OutputState }

func (ListActiveSecurityAdminRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveSecurityAdminRuleResult)(nil)).Elem()
}

func (o ListActiveSecurityAdminRuleResultOutput) ToListActiveSecurityAdminRuleResultOutput() ListActiveSecurityAdminRuleResultOutput {
	return o
}

func (o ListActiveSecurityAdminRuleResultOutput) ToListActiveSecurityAdminRuleResultOutputWithContext(ctx context.Context) ListActiveSecurityAdminRuleResultOutput {
	return o
}

func (o ListActiveSecurityAdminRuleResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListActiveSecurityAdminRuleResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

func (o ListActiveSecurityAdminRuleResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v ListActiveSecurityAdminRuleResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListActiveSecurityAdminRuleResultOutput{})
}
