


package servicebus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamespaceAuthorizationRule(ctx *pulumi.Context, args *LookupNamespaceAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceAuthorizationRuleResult, error) {
	var rv LookupNamespaceAuthorizationRuleResult
	err := ctx.Invoke("azure-native:servicebus:getNamespaceAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceAuthorizationRuleArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupNamespaceAuthorizationRuleResult struct {
	Id     string   `pulumi:"id"`
	Name   string   `pulumi:"name"`
	Rights []string `pulumi:"rights"`
	Type   string   `pulumi:"type"`
}

func LookupNamespaceAuthorizationRuleOutput(ctx *pulumi.Context, args LookupNamespaceAuthorizationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNamespaceAuthorizationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNamespaceAuthorizationRuleResult, error) {
			args := v.(LookupNamespaceAuthorizationRuleArgs)
			r, err := LookupNamespaceAuthorizationRule(ctx, &args, opts...)
			var s LookupNamespaceAuthorizationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNamespaceAuthorizationRuleResultOutput)
}

type LookupNamespaceAuthorizationRuleOutputArgs struct {
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	NamespaceName         pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNamespaceAuthorizationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceAuthorizationRuleArgs)(nil)).Elem()
}


type LookupNamespaceAuthorizationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNamespaceAuthorizationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceAuthorizationRuleResult)(nil)).Elem()
}

func (o LookupNamespaceAuthorizationRuleResultOutput) ToLookupNamespaceAuthorizationRuleResultOutput() LookupNamespaceAuthorizationRuleResultOutput {
	return o
}

func (o LookupNamespaceAuthorizationRuleResultOutput) ToLookupNamespaceAuthorizationRuleResultOutputWithContext(ctx context.Context) LookupNamespaceAuthorizationRuleResultOutput {
	return o
}

func (o LookupNamespaceAuthorizationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNamespaceAuthorizationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNamespaceAuthorizationRuleResultOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) []string { return v.Rights }).(pulumi.StringArrayOutput)
}

func (o LookupNamespaceAuthorizationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceAuthorizationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamespaceAuthorizationRuleResultOutput{})
}
