


package v20210801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkRule(ctx *pulumi.Context, args *LookupVirtualNetworkRuleArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkRuleResult, error) {
	var rv LookupVirtualNetworkRuleResult
	err := ctx.Invoke("azure-native:sql/v20210801preview:getVirtualNetworkRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkRuleArgs struct {
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ServerName             string `pulumi:"serverName"`
	VirtualNetworkRuleName string `pulumi:"virtualNetworkRuleName"`
}


type LookupVirtualNetworkRuleResult struct {
	Id                               string `pulumi:"id"`
	IgnoreMissingVnetServiceEndpoint *bool  `pulumi:"ignoreMissingVnetServiceEndpoint"`
	Name                             string `pulumi:"name"`
	State                            string `pulumi:"state"`
	Type                             string `pulumi:"type"`
	VirtualNetworkSubnetId           string `pulumi:"virtualNetworkSubnetId"`
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
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName             pulumi.StringInput `pulumi:"serverName"`
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

func (o LookupVirtualNetworkRuleResultOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRuleResult) *bool { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNetworkRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkRuleResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRuleResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkRuleResultOutput) VirtualNetworkSubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRuleResult) string { return v.VirtualNetworkSubnetId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkRuleResultOutput{})
}
