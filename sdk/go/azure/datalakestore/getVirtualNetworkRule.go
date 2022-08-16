


package datalakestore

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkRule(ctx *pulumi.Context, args *LookupVirtualNetworkRuleArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkRuleResult, error) {
	var rv LookupVirtualNetworkRuleResult
	err := ctx.Invoke("azure-native:datalakestore:getVirtualNetworkRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkRuleArgs struct {
	AccountName            string `pulumi:"accountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	VirtualNetworkRuleName string `pulumi:"virtualNetworkRuleName"`
}


type LookupVirtualNetworkRuleResult struct {
	Id       string `pulumi:"id"`
	Name     string `pulumi:"name"`
	SubnetId string `pulumi:"subnetId"`
	Type     string `pulumi:"type"`
}

func LookupVirtualNetworkRuleOutput(ctx *pulumi.Context, args LookupVirtualNetworkRuleOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNetworkRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNetworkRuleResult, error) {
			args := v.(LookupVirtualNetworkRuleArgs)
			r, err := LookupVirtualNetworkRule(ctx, &args, opts...)
			var s LookupVirtualNetworkRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNetworkRuleResultOutput)
}

type LookupVirtualNetworkRuleOutputArgs struct {
	AccountName            pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualNetworkRuleName pulumi.StringInput `pulumi:"virtualNetworkRuleName"`
}

func (LookupVirtualNetworkRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkRuleArgs)(nil)).Elem()
}


type LookupVirtualNetworkRuleResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNetworkRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkRuleResult)(nil)).Elem()
}

func (o LookupVirtualNetworkRuleResultOutput) ToLookupVirtualNetworkRuleResultOutput() LookupVirtualNetworkRuleResultOutput {
	return o
}

func (o LookupVirtualNetworkRuleResultOutput) ToLookupVirtualNetworkRuleResultOutputWithContext(ctx context.Context) LookupVirtualNetworkRuleResultOutput {
	return o
}

func (o LookupVirtualNetworkRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkRuleResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRuleResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkRuleResultOutput{})
}
