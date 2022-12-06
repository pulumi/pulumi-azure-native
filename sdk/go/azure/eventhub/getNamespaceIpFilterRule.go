


package eventhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamespaceIpFilterRule(ctx *pulumi.Context, args *LookupNamespaceIpFilterRuleArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceIpFilterRuleResult, error) {
	var rv LookupNamespaceIpFilterRuleResult
	err := ctx.Invoke("azure-native:eventhub:getNamespaceIpFilterRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceIpFilterRuleArgs struct {
	IpFilterRuleName  string `pulumi:"ipFilterRuleName"`
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupNamespaceIpFilterRuleResult struct {
	Action     *string `pulumi:"action"`
	FilterName *string `pulumi:"filterName"`
	Id         string  `pulumi:"id"`
	IpMask     *string `pulumi:"ipMask"`
	Name       string  `pulumi:"name"`
	Type       string  `pulumi:"type"`
}

func LookupNamespaceIpFilterRuleOutput(ctx *pulumi.Context, args LookupNamespaceIpFilterRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNamespaceIpFilterRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNamespaceIpFilterRuleResult, error) {
			args := v.(LookupNamespaceIpFilterRuleArgs)
			r, err := LookupNamespaceIpFilterRule(ctx, &args, opts...)
			var s LookupNamespaceIpFilterRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNamespaceIpFilterRuleResultOutput)
}

type LookupNamespaceIpFilterRuleOutputArgs struct {
	IpFilterRuleName  pulumi.StringInput `pulumi:"ipFilterRuleName"`
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNamespaceIpFilterRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceIpFilterRuleArgs)(nil)).Elem()
}


type LookupNamespaceIpFilterRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNamespaceIpFilterRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceIpFilterRuleResult)(nil)).Elem()
}

func (o LookupNamespaceIpFilterRuleResultOutput) ToLookupNamespaceIpFilterRuleResultOutput() LookupNamespaceIpFilterRuleResultOutput {
	return o
}

func (o LookupNamespaceIpFilterRuleResultOutput) ToLookupNamespaceIpFilterRuleResultOutputWithContext(ctx context.Context) LookupNamespaceIpFilterRuleResultOutput {
	return o
}

func (o LookupNamespaceIpFilterRuleResultOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceIpFilterRuleResult) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceIpFilterRuleResultOutput) FilterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceIpFilterRuleResult) *string { return v.FilterName }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceIpFilterRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceIpFilterRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNamespaceIpFilterRuleResultOutput) IpMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceIpFilterRuleResult) *string { return v.IpMask }).(pulumi.StringPtrOutput)
}

func (o LookupNamespaceIpFilterRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceIpFilterRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNamespaceIpFilterRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceIpFilterRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamespaceIpFilterRuleResultOutput{})
}
