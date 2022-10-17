


package v20180101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupQueueAuthorizationRule(ctx *pulumi.Context, args *LookupQueueAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupQueueAuthorizationRuleResult, error) {
	var rv LookupQueueAuthorizationRuleResult
	err := ctx.Invoke("azure-native:servicebus/v20180101preview:getQueueAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueueAuthorizationRuleArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	QueueName             string `pulumi:"queueName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupQueueAuthorizationRuleResult struct {
	Id     string   `pulumi:"id"`
	Name   string   `pulumi:"name"`
	Rights []string `pulumi:"rights"`
	Type   string   `pulumi:"type"`
}

func LookupQueueAuthorizationRuleOutput(ctx *pulumi.Context, args LookupQueueAuthorizationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupQueueAuthorizationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQueueAuthorizationRuleResult, error) {
			args := v.(LookupQueueAuthorizationRuleArgs)
			r, err := LookupQueueAuthorizationRule(ctx, &args, opts...)
			var s LookupQueueAuthorizationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQueueAuthorizationRuleResultOutput)
}

type LookupQueueAuthorizationRuleOutputArgs struct {
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	NamespaceName         pulumi.StringInput `pulumi:"namespaceName"`
	QueueName             pulumi.StringInput `pulumi:"queueName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupQueueAuthorizationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueAuthorizationRuleArgs)(nil)).Elem()
}


type LookupQueueAuthorizationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupQueueAuthorizationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueAuthorizationRuleResult)(nil)).Elem()
}

func (o LookupQueueAuthorizationRuleResultOutput) ToLookupQueueAuthorizationRuleResultOutput() LookupQueueAuthorizationRuleResultOutput {
	return o
}

func (o LookupQueueAuthorizationRuleResultOutput) ToLookupQueueAuthorizationRuleResultOutputWithContext(ctx context.Context) LookupQueueAuthorizationRuleResultOutput {
	return o
}

func (o LookupQueueAuthorizationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupQueueAuthorizationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupQueueAuthorizationRuleResultOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) []string { return v.Rights }).(pulumi.StringArrayOutput)
}

func (o LookupQueueAuthorizationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueAuthorizationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQueueAuthorizationRuleResultOutput{})
}
