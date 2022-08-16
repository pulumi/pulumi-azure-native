


package automation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebhook(ctx *pulumi.Context, args *LookupWebhookArgs, opts ...pulumi.InvokeOption) (*LookupWebhookResult, error) {
	var rv LookupWebhookResult
	err := ctx.Invoke("azure-native:automation:getWebhook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWebhookArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	WebhookName           string `pulumi:"webhookName"`
}


type LookupWebhookResult struct {
	CreationTime     *string                             `pulumi:"creationTime"`
	Description      *string                             `pulumi:"description"`
	ExpiryTime       *string                             `pulumi:"expiryTime"`
	Id               string                              `pulumi:"id"`
	IsEnabled        *bool                               `pulumi:"isEnabled"`
	LastInvokedTime  *string                             `pulumi:"lastInvokedTime"`
	LastModifiedBy   *string                             `pulumi:"lastModifiedBy"`
	LastModifiedTime *string                             `pulumi:"lastModifiedTime"`
	Name             string                              `pulumi:"name"`
	Parameters       map[string]string                   `pulumi:"parameters"`
	RunOn            *string                             `pulumi:"runOn"`
	Runbook          *RunbookAssociationPropertyResponse `pulumi:"runbook"`
	Type             string                              `pulumi:"type"`
	Uri              *string                             `pulumi:"uri"`
}


func (val *LookupWebhookResult) Defaults() *LookupWebhookResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsEnabled) {
		isEnabled_ := false
		tmp.IsEnabled = &isEnabled_
	}
	return &tmp
}

func LookupWebhookOutput(ctx *pulumi.Context, args LookupWebhookOutputArgs, opts ...pulumi.InvokeOption) LookupWebhookResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebhookResult, error) {
			args := v.(LookupWebhookArgs)
			r, err := LookupWebhook(ctx, &args, opts...)
			var s LookupWebhookResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebhookResultOutput)
}

type LookupWebhookOutputArgs struct {
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	WebhookName           pulumi.StringInput `pulumi:"webhookName"`
}

func (LookupWebhookOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebhookArgs)(nil)).Elem()
}


type LookupWebhookResultOutput struct{ *pulumi.OutputState }

func (LookupWebhookResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebhookResult)(nil)).Elem()
}

func (o LookupWebhookResultOutput) ToLookupWebhookResultOutput() LookupWebhookResultOutput {
	return o
}

func (o LookupWebhookResultOutput) ToLookupWebhookResultOutputWithContext(ctx context.Context) LookupWebhookResultOutput {
	return o
}

func (o LookupWebhookResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebhookResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupWebhookResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebhookResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupWebhookResultOutput) ExpiryTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebhookResult) *string { return v.ExpiryTime }).(pulumi.StringPtrOutput)
}

func (o LookupWebhookResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhookResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebhookResultOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebhookResult) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupWebhookResultOutput) LastInvokedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebhookResult) *string { return v.LastInvokedTime }).(pulumi.StringPtrOutput)
}

func (o LookupWebhookResultOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebhookResult) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o LookupWebhookResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebhookResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o LookupWebhookResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhookResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebhookResultOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebhookResult) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o LookupWebhookResultOutput) RunOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebhookResult) *string { return v.RunOn }).(pulumi.StringPtrOutput)
}

func (o LookupWebhookResultOutput) Runbook() RunbookAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupWebhookResult) *RunbookAssociationPropertyResponse { return v.Runbook }).(RunbookAssociationPropertyResponsePtrOutput)
}

func (o LookupWebhookResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhookResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWebhookResultOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebhookResult) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebhookResultOutput{})
}
