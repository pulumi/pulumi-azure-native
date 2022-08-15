


package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListActiveSecurityUserRules(ctx *pulumi.Context, args *ListActiveSecurityUserRulesArgs, opts ...pulumi.InvokeOption) (*ListActiveSecurityUserRulesResult, error) {
	var rv ListActiveSecurityUserRulesResult
	err := ctx.Invoke("azure-native:network/v20220201preview:listActiveSecurityUserRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListActiveSecurityUserRulesArgs struct {
	NetworkManagerName string   `pulumi:"networkManagerName"`
	Regions            []string `pulumi:"regions"`
	ResourceGroupName  string   `pulumi:"resourceGroupName"`
	SkipToken          *string  `pulumi:"skipToken"`
}


type ListActiveSecurityUserRulesResult struct {
	SkipToken *string       `pulumi:"skipToken"`
	Value     []interface{} `pulumi:"value"`
}

func ListActiveSecurityUserRulesOutput(ctx *pulumi.Context, args ListActiveSecurityUserRulesOutputArgs, opts ...pulumi.InvokeOption) ListActiveSecurityUserRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListActiveSecurityUserRulesResult, error) {
			args := v.(ListActiveSecurityUserRulesArgs)
			r, err := ListActiveSecurityUserRules(ctx, &args, opts...)
			var s ListActiveSecurityUserRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListActiveSecurityUserRulesResultOutput)
}

type ListActiveSecurityUserRulesOutputArgs struct {
	NetworkManagerName pulumi.StringInput      `pulumi:"networkManagerName"`
	Regions            pulumi.StringArrayInput `pulumi:"regions"`
	ResourceGroupName  pulumi.StringInput      `pulumi:"resourceGroupName"`
	SkipToken          pulumi.StringPtrInput   `pulumi:"skipToken"`
}

func (ListActiveSecurityUserRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveSecurityUserRulesArgs)(nil)).Elem()
}


type ListActiveSecurityUserRulesResultOutput struct{ *pulumi.OutputState }

func (ListActiveSecurityUserRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveSecurityUserRulesResult)(nil)).Elem()
}

func (o ListActiveSecurityUserRulesResultOutput) ToListActiveSecurityUserRulesResultOutput() ListActiveSecurityUserRulesResultOutput {
	return o
}

func (o ListActiveSecurityUserRulesResultOutput) ToListActiveSecurityUserRulesResultOutputWithContext(ctx context.Context) ListActiveSecurityUserRulesResultOutput {
	return o
}

func (o ListActiveSecurityUserRulesResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListActiveSecurityUserRulesResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

func (o ListActiveSecurityUserRulesResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v ListActiveSecurityUserRulesResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListActiveSecurityUserRulesResultOutput{})
}
