


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupEventHubAuthorizationRule(ctx *pulumi.Context, args *LookupEventHubAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupEventHubAuthorizationRuleResult, error) {
	var rv LookupEventHubAuthorizationRuleResult
	err := ctx.Invoke("azure-native:eventhub/v20150801:getEventHubAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventHubAuthorizationRuleArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	EventHubName          string `pulumi:"eventHubName"`
	NamespaceName         string `pulumi:"namespaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupEventHubAuthorizationRuleResult struct {
	Id       string   `pulumi:"id"`
	Location *string  `pulumi:"location"`
	Name     string   `pulumi:"name"`
	Rights   []string `pulumi:"rights"`
	Type     string   `pulumi:"type"`
}

func LookupEventHubAuthorizationRuleOutput(ctx *pulumi.Context, args LookupEventHubAuthorizationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupEventHubAuthorizationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventHubAuthorizationRuleResult, error) {
			args := v.(LookupEventHubAuthorizationRuleArgs)
			r, err := LookupEventHubAuthorizationRule(ctx, &args, opts...)
			var s LookupEventHubAuthorizationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventHubAuthorizationRuleResultOutput)
}

type LookupEventHubAuthorizationRuleOutputArgs struct {
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	EventHubName          pulumi.StringInput `pulumi:"eventHubName"`
	NamespaceName         pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEventHubAuthorizationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventHubAuthorizationRuleArgs)(nil)).Elem()
}


type LookupEventHubAuthorizationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupEventHubAuthorizationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventHubAuthorizationRuleResult)(nil)).Elem()
}

func (o LookupEventHubAuthorizationRuleResultOutput) ToLookupEventHubAuthorizationRuleResultOutput() LookupEventHubAuthorizationRuleResultOutput {
	return o
}

func (o LookupEventHubAuthorizationRuleResultOutput) ToLookupEventHubAuthorizationRuleResultOutputWithContext(ctx context.Context) LookupEventHubAuthorizationRuleResultOutput {
	return o
}

func (o LookupEventHubAuthorizationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubAuthorizationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEventHubAuthorizationRuleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventHubAuthorizationRuleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupEventHubAuthorizationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubAuthorizationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEventHubAuthorizationRuleResultOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEventHubAuthorizationRuleResult) []string { return v.Rights }).(pulumi.StringArrayOutput)
}

func (o LookupEventHubAuthorizationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventHubAuthorizationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventHubAuthorizationRuleResultOutput{})
}
