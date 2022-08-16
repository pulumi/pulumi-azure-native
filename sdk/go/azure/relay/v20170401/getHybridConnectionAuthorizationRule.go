


package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHybridConnectionAuthorizationRule(ctx *pulumi.Context, args *LookupHybridConnectionAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupHybridConnectionAuthorizationRuleResult, error) {
	var rv LookupHybridConnectionAuthorizationRuleResult
	err := ctx.Invoke("azure-native:relay/v20170401:getHybridConnectionAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHybridConnectionAuthorizationRuleArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	HybridConnectionName  string `pulumi:"hybridConnectionName"`
	NamespaceName         string `pulumi:"namespaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupHybridConnectionAuthorizationRuleResult struct {
	Id     string   `pulumi:"id"`
	Name   string   `pulumi:"name"`
	Rights []string `pulumi:"rights"`
	Type   string   `pulumi:"type"`
}

func LookupHybridConnectionAuthorizationRuleOutput(ctx *pulumi.Context, args LookupHybridConnectionAuthorizationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupHybridConnectionAuthorizationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHybridConnectionAuthorizationRuleResult, error) {
			args := v.(LookupHybridConnectionAuthorizationRuleArgs)
			r, err := LookupHybridConnectionAuthorizationRule(ctx, &args, opts...)
			var s LookupHybridConnectionAuthorizationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHybridConnectionAuthorizationRuleResultOutput)
}

type LookupHybridConnectionAuthorizationRuleOutputArgs struct {
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	HybridConnectionName  pulumi.StringInput `pulumi:"hybridConnectionName"`
	NamespaceName         pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupHybridConnectionAuthorizationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridConnectionAuthorizationRuleArgs)(nil)).Elem()
}


type LookupHybridConnectionAuthorizationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupHybridConnectionAuthorizationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridConnectionAuthorizationRuleResult)(nil)).Elem()
}

func (o LookupHybridConnectionAuthorizationRuleResultOutput) ToLookupHybridConnectionAuthorizationRuleResultOutput() LookupHybridConnectionAuthorizationRuleResultOutput {
	return o
}

func (o LookupHybridConnectionAuthorizationRuleResultOutput) ToLookupHybridConnectionAuthorizationRuleResultOutputWithContext(ctx context.Context) LookupHybridConnectionAuthorizationRuleResultOutput {
	return o
}

func (o LookupHybridConnectionAuthorizationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridConnectionAuthorizationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHybridConnectionAuthorizationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridConnectionAuthorizationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupHybridConnectionAuthorizationRuleResultOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHybridConnectionAuthorizationRuleResult) []string { return v.Rights }).(pulumi.StringArrayOutput)
}

func (o LookupHybridConnectionAuthorizationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridConnectionAuthorizationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHybridConnectionAuthorizationRuleResultOutput{})
}
