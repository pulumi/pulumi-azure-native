


package v20160301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupNotificationHubAuthorizationRule(ctx *pulumi.Context, args *LookupNotificationHubAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupNotificationHubAuthorizationRuleResult, error) {
	var rv LookupNotificationHubAuthorizationRuleResult
	err := ctx.Invoke("azure-native:notificationhubs/v20160301:getNotificationHubAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationHubAuthorizationRuleArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	NotificationHubName   string `pulumi:"notificationHubName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupNotificationHubAuthorizationRuleResult struct {
	Id       string            `pulumi:"id"`
	Location string            `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Rights   []string          `pulumi:"rights"`
	Sku      *SkuResponse      `pulumi:"sku"`
	Tags     map[string]string `pulumi:"tags"`
	Type     string            `pulumi:"type"`
}

func LookupNotificationHubAuthorizationRuleOutput(ctx *pulumi.Context, args LookupNotificationHubAuthorizationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNotificationHubAuthorizationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNotificationHubAuthorizationRuleResult, error) {
			args := v.(LookupNotificationHubAuthorizationRuleArgs)
			r, err := LookupNotificationHubAuthorizationRule(ctx, &args, opts...)
			var s LookupNotificationHubAuthorizationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNotificationHubAuthorizationRuleResultOutput)
}

type LookupNotificationHubAuthorizationRuleOutputArgs struct {
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	NamespaceName         pulumi.StringInput `pulumi:"namespaceName"`
	NotificationHubName   pulumi.StringInput `pulumi:"notificationHubName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNotificationHubAuthorizationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationHubAuthorizationRuleArgs)(nil)).Elem()
}


type LookupNotificationHubAuthorizationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNotificationHubAuthorizationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationHubAuthorizationRuleResult)(nil)).Elem()
}

func (o LookupNotificationHubAuthorizationRuleResultOutput) ToLookupNotificationHubAuthorizationRuleResultOutput() LookupNotificationHubAuthorizationRuleResultOutput {
	return o
}

func (o LookupNotificationHubAuthorizationRuleResultOutput) ToLookupNotificationHubAuthorizationRuleResultOutputWithContext(ctx context.Context) LookupNotificationHubAuthorizationRuleResultOutput {
	return o
}

func (o LookupNotificationHubAuthorizationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNotificationHubAuthorizationRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupNotificationHubAuthorizationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNotificationHubAuthorizationRuleResultOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) []string { return v.Rights }).(pulumi.StringArrayOutput)
}

func (o LookupNotificationHubAuthorizationRuleResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupNotificationHubAuthorizationRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNotificationHubAuthorizationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationHubAuthorizationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotificationHubAuthorizationRuleResultOutput{})
}
