


package v20180101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRule(ctx *pulumi.Context, args *LookupRuleArgs, opts ...pulumi.InvokeOption) (*LookupRuleResult, error) {
	var rv LookupRuleResult
	err := ctx.Invoke("azure-native:servicebus/v20180101preview:getRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRuleArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleName          string `pulumi:"ruleName"`
	SubscriptionName  string `pulumi:"subscriptionName"`
	TopicName         string `pulumi:"topicName"`
}


type LookupRuleResult struct {
	Action            *ActionResponse            `pulumi:"action"`
	CorrelationFilter *CorrelationFilterResponse `pulumi:"correlationFilter"`
	FilterType        *string                    `pulumi:"filterType"`
	Id                string                     `pulumi:"id"`
	Name              string                     `pulumi:"name"`
	SqlFilter         *SqlFilterResponse         `pulumi:"sqlFilter"`
	Type              string                     `pulumi:"type"`
}


func (val *LookupRuleResult) Defaults() *LookupRuleResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Action = tmp.Action.Defaults()

	tmp.CorrelationFilter = tmp.CorrelationFilter.Defaults()

	tmp.SqlFilter = tmp.SqlFilter.Defaults()

	return &tmp
}

func LookupRuleOutput(ctx *pulumi.Context, args LookupRuleOutputArgs, opts ...pulumi.InvokeOption) LookupRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRuleResult, error) {
			args := v.(LookupRuleArgs)
			r, err := LookupRule(ctx, &args, opts...)
			var s LookupRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRuleResultOutput)
}

type LookupRuleOutputArgs struct {
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleName          pulumi.StringInput `pulumi:"ruleName"`
	SubscriptionName  pulumi.StringInput `pulumi:"subscriptionName"`
	TopicName         pulumi.StringInput `pulumi:"topicName"`
}

func (LookupRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleArgs)(nil)).Elem()
}


type LookupRuleResultOutput struct{ *pulumi.OutputState }

func (LookupRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleResult)(nil)).Elem()
}

func (o LookupRuleResultOutput) ToLookupRuleResultOutput() LookupRuleResultOutput {
	return o
}

func (o LookupRuleResultOutput) ToLookupRuleResultOutputWithContext(ctx context.Context) LookupRuleResultOutput {
	return o
}

func (o LookupRuleResultOutput) Action() ActionResponsePtrOutput {
	return o.ApplyT(func(v LookupRuleResult) *ActionResponse { return v.Action }).(ActionResponsePtrOutput)
}

func (o LookupRuleResultOutput) CorrelationFilter() CorrelationFilterResponsePtrOutput {
	return o.ApplyT(func(v LookupRuleResult) *CorrelationFilterResponse { return v.CorrelationFilter }).(CorrelationFilterResponsePtrOutput)
}

func (o LookupRuleResultOutput) FilterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRuleResult) *string { return v.FilterType }).(pulumi.StringPtrOutput)
}

func (o LookupRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRuleResultOutput) SqlFilter() SqlFilterResponsePtrOutput {
	return o.ApplyT(func(v LookupRuleResult) *SqlFilterResponse { return v.SqlFilter }).(SqlFilterResponsePtrOutput)
}

func (o LookupRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRuleResultOutput{})
}
