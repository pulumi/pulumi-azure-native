


package servicebus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamespaceVirtualNetworkRule(ctx *pulumi.Context, args *LookupNamespaceVirtualNetworkRuleArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceVirtualNetworkRuleResult, error) {
	var rv LookupNamespaceVirtualNetworkRuleResult
	err := ctx.Invoke("azure-native:servicebus:getNamespaceVirtualNetworkRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceVirtualNetworkRuleArgs struct {
	NamespaceName          string `pulumi:"namespaceName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	VirtualNetworkRuleName string `pulumi:"virtualNetworkRuleName"`
}


type LookupNamespaceVirtualNetworkRuleResult struct {
	Id                     string  `pulumi:"id"`
	Name                   string  `pulumi:"name"`
	Type                   string  `pulumi:"type"`
	VirtualNetworkSubnetId *string `pulumi:"virtualNetworkSubnetId"`
}

func LookupNamespaceVirtualNetworkRuleOutput(ctx *pulumi.Context, args LookupNamespaceVirtualNetworkRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNamespaceVirtualNetworkRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNamespaceVirtualNetworkRuleResult, error) {
			args := v.(LookupNamespaceVirtualNetworkRuleArgs)
			r, err := LookupNamespaceVirtualNetworkRule(ctx, &args, opts...)
			var s LookupNamespaceVirtualNetworkRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNamespaceVirtualNetworkRuleResultOutput)
}

type LookupNamespaceVirtualNetworkRuleOutputArgs struct {
	NamespaceName          pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualNetworkRuleName pulumi.StringInput `pulumi:"virtualNetworkRuleName"`
}

func (LookupNamespaceVirtualNetworkRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceVirtualNetworkRuleArgs)(nil)).Elem()
}


type LookupNamespaceVirtualNetworkRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNamespaceVirtualNetworkRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceVirtualNetworkRuleResult)(nil)).Elem()
}

func (o LookupNamespaceVirtualNetworkRuleResultOutput) ToLookupNamespaceVirtualNetworkRuleResultOutput() LookupNamespaceVirtualNetworkRuleResultOutput {
	return o
}

func (o LookupNamespaceVirtualNetworkRuleResultOutput) ToLookupNamespaceVirtualNetworkRuleResultOutputWithContext(ctx context.Context) LookupNamespaceVirtualNetworkRuleResultOutput {
	return o
}

func (o LookupNamespaceVirtualNetworkRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceVirtualNetworkRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNamespaceVirtualNetworkRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceVirtualNetworkRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNamespaceVirtualNetworkRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceVirtualNetworkRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupNamespaceVirtualNetworkRuleResultOutput) VirtualNetworkSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceVirtualNetworkRuleResult) *string { return v.VirtualNetworkSubnetId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamespaceVirtualNetworkRuleResultOutput{})
}
