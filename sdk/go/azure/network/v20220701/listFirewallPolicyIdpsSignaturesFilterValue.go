


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListFirewallPolicyIdpsSignaturesFilterValue(ctx *pulumi.Context, args *ListFirewallPolicyIdpsSignaturesFilterValueArgs, opts ...pulumi.InvokeOption) (*ListFirewallPolicyIdpsSignaturesFilterValueResult, error) {
	var rv ListFirewallPolicyIdpsSignaturesFilterValueResult
	err := ctx.Invoke("azure-native:network/v20220701:listFirewallPolicyIdpsSignaturesFilterValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListFirewallPolicyIdpsSignaturesFilterValueArgs struct {
	FilterName         *string `pulumi:"filterName"`
	FirewallPolicyName string  `pulumi:"firewallPolicyName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
}


type ListFirewallPolicyIdpsSignaturesFilterValueResult struct {
	FilterValues []string `pulumi:"filterValues"`
}

func ListFirewallPolicyIdpsSignaturesFilterValueOutput(ctx *pulumi.Context, args ListFirewallPolicyIdpsSignaturesFilterValueOutputArgs, opts ...pulumi.InvokeOption) ListFirewallPolicyIdpsSignaturesFilterValueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListFirewallPolicyIdpsSignaturesFilterValueResult, error) {
			args := v.(ListFirewallPolicyIdpsSignaturesFilterValueArgs)
			r, err := ListFirewallPolicyIdpsSignaturesFilterValue(ctx, &args, opts...)
			var s ListFirewallPolicyIdpsSignaturesFilterValueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListFirewallPolicyIdpsSignaturesFilterValueResultOutput)
}

type ListFirewallPolicyIdpsSignaturesFilterValueOutputArgs struct {
	FilterName         pulumi.StringPtrInput `pulumi:"filterName"`
	FirewallPolicyName pulumi.StringInput    `pulumi:"firewallPolicyName"`
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (ListFirewallPolicyIdpsSignaturesFilterValueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFirewallPolicyIdpsSignaturesFilterValueArgs)(nil)).Elem()
}


type ListFirewallPolicyIdpsSignaturesFilterValueResultOutput struct{ *pulumi.OutputState }

func (ListFirewallPolicyIdpsSignaturesFilterValueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFirewallPolicyIdpsSignaturesFilterValueResult)(nil)).Elem()
}

func (o ListFirewallPolicyIdpsSignaturesFilterValueResultOutput) ToListFirewallPolicyIdpsSignaturesFilterValueResultOutput() ListFirewallPolicyIdpsSignaturesFilterValueResultOutput {
	return o
}

func (o ListFirewallPolicyIdpsSignaturesFilterValueResultOutput) ToListFirewallPolicyIdpsSignaturesFilterValueResultOutputWithContext(ctx context.Context) ListFirewallPolicyIdpsSignaturesFilterValueResultOutput {
	return o
}

func (o ListFirewallPolicyIdpsSignaturesFilterValueResultOutput) FilterValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListFirewallPolicyIdpsSignaturesFilterValueResult) []string { return v.FilterValues }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListFirewallPolicyIdpsSignaturesFilterValueResultOutput{})
}
