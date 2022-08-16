


package v20160701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupWCFRelayAuthorizationRule(ctx *pulumi.Context, args *LookupWCFRelayAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupWCFRelayAuthorizationRuleResult, error) {
	var rv LookupWCFRelayAuthorizationRuleResult
	err := ctx.Invoke("azure-native:relay/v20160701:getWCFRelayAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWCFRelayAuthorizationRuleArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	RelayName             string `pulumi:"relayName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupWCFRelayAuthorizationRuleResult struct {
	Id     string   `pulumi:"id"`
	Name   string   `pulumi:"name"`
	Rights []string `pulumi:"rights"`
	Type   string   `pulumi:"type"`
}

func LookupWCFRelayAuthorizationRuleOutput(ctx *pulumi.Context, args LookupWCFRelayAuthorizationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupWCFRelayAuthorizationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWCFRelayAuthorizationRuleResult, error) {
			args := v.(LookupWCFRelayAuthorizationRuleArgs)
			r, err := LookupWCFRelayAuthorizationRule(ctx, &args, opts...)
			var s LookupWCFRelayAuthorizationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWCFRelayAuthorizationRuleResultOutput)
}

type LookupWCFRelayAuthorizationRuleOutputArgs struct {
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	NamespaceName         pulumi.StringInput `pulumi:"namespaceName"`
	RelayName             pulumi.StringInput `pulumi:"relayName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWCFRelayAuthorizationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWCFRelayAuthorizationRuleArgs)(nil)).Elem()
}


type LookupWCFRelayAuthorizationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupWCFRelayAuthorizationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWCFRelayAuthorizationRuleResult)(nil)).Elem()
}

func (o LookupWCFRelayAuthorizationRuleResultOutput) ToLookupWCFRelayAuthorizationRuleResultOutput() LookupWCFRelayAuthorizationRuleResultOutput {
	return o
}

func (o LookupWCFRelayAuthorizationRuleResultOutput) ToLookupWCFRelayAuthorizationRuleResultOutputWithContext(ctx context.Context) LookupWCFRelayAuthorizationRuleResultOutput {
	return o
}

func (o LookupWCFRelayAuthorizationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWCFRelayAuthorizationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWCFRelayAuthorizationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWCFRelayAuthorizationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWCFRelayAuthorizationRuleResultOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWCFRelayAuthorizationRuleResult) []string { return v.Rights }).(pulumi.StringArrayOutput)
}

func (o LookupWCFRelayAuthorizationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWCFRelayAuthorizationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWCFRelayAuthorizationRuleResultOutput{})
}
