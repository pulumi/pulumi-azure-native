


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListActiveSecurityAdminRules(ctx *pulumi.Context, args *ListActiveSecurityAdminRulesArgs, opts ...pulumi.InvokeOption) (*ListActiveSecurityAdminRulesResult, error) {
	var rv ListActiveSecurityAdminRulesResult
	err := ctx.Invoke("azure-native:network/v20220401preview:listActiveSecurityAdminRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListActiveSecurityAdminRulesArgs struct {
	NetworkManagerName string   `pulumi:"networkManagerName"`
	Regions            []string `pulumi:"regions"`
	ResourceGroupName  string   `pulumi:"resourceGroupName"`
	SkipToken          *string  `pulumi:"skipToken"`
}


type ListActiveSecurityAdminRulesResult struct {
	SkipToken *string       `pulumi:"skipToken"`
	Value     []interface{} `pulumi:"value"`
}

func ListActiveSecurityAdminRulesOutput(ctx *pulumi.Context, args ListActiveSecurityAdminRulesOutputArgs, opts ...pulumi.InvokeOption) ListActiveSecurityAdminRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListActiveSecurityAdminRulesResult, error) {
			args := v.(ListActiveSecurityAdminRulesArgs)
			r, err := ListActiveSecurityAdminRules(ctx, &args, opts...)
			var s ListActiveSecurityAdminRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListActiveSecurityAdminRulesResultOutput)
}

type ListActiveSecurityAdminRulesOutputArgs struct {
	NetworkManagerName pulumi.StringInput      `pulumi:"networkManagerName"`
	Regions            pulumi.StringArrayInput `pulumi:"regions"`
	ResourceGroupName  pulumi.StringInput      `pulumi:"resourceGroupName"`
	SkipToken          pulumi.StringPtrInput   `pulumi:"skipToken"`
}

func (ListActiveSecurityAdminRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveSecurityAdminRulesArgs)(nil)).Elem()
}


type ListActiveSecurityAdminRulesResultOutput struct{ *pulumi.OutputState }

func (ListActiveSecurityAdminRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveSecurityAdminRulesResult)(nil)).Elem()
}

func (o ListActiveSecurityAdminRulesResultOutput) ToListActiveSecurityAdminRulesResultOutput() ListActiveSecurityAdminRulesResultOutput {
	return o
}

func (o ListActiveSecurityAdminRulesResultOutput) ToListActiveSecurityAdminRulesResultOutputWithContext(ctx context.Context) ListActiveSecurityAdminRulesResultOutput {
	return o
}

func (o ListActiveSecurityAdminRulesResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListActiveSecurityAdminRulesResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

func (o ListActiveSecurityAdminRulesResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v ListActiveSecurityAdminRulesResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListActiveSecurityAdminRulesResultOutput{})
}
