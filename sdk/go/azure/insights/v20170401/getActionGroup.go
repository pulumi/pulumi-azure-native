


package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupActionGroup(ctx *pulumi.Context, args *LookupActionGroupArgs, opts ...pulumi.InvokeOption) (*LookupActionGroupResult, error) {
	var rv LookupActionGroupResult
	err := ctx.Invoke("azure-native:insights/v20170401:getActionGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupActionGroupArgs struct {
	ActionGroupName   string `pulumi:"actionGroupName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupActionGroupResult struct {
	AutomationRunbookReceivers []AutomationRunbookReceiverResponse `pulumi:"automationRunbookReceivers"`
	AzureAppPushReceivers      []AzureAppPushReceiverResponse      `pulumi:"azureAppPushReceivers"`
	EmailReceivers             []EmailReceiverResponse             `pulumi:"emailReceivers"`
	Enabled                    bool                                `pulumi:"enabled"`
	GroupShortName             string                              `pulumi:"groupShortName"`
	Id                         string                              `pulumi:"id"`
	ItsmReceivers              []ItsmReceiverResponse              `pulumi:"itsmReceivers"`
	Location                   string                              `pulumi:"location"`
	Name                       string                              `pulumi:"name"`
	SmsReceivers               []SmsReceiverResponse               `pulumi:"smsReceivers"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
	WebhookReceivers           []WebhookReceiverResponse           `pulumi:"webhookReceivers"`
}


func (val *LookupActionGroupResult) Defaults() *LookupActionGroupResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Enabled) {
		tmp.Enabled = true
	}
	return &tmp
}

func LookupActionGroupOutput(ctx *pulumi.Context, args LookupActionGroupOutputArgs, opts ...pulumi.InvokeOption) LookupActionGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupActionGroupResult, error) {
			args := v.(LookupActionGroupArgs)
			r, err := LookupActionGroup(ctx, &args, opts...)
			var s LookupActionGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupActionGroupResultOutput)
}

type LookupActionGroupOutputArgs struct {
	ActionGroupName   pulumi.StringInput `pulumi:"actionGroupName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupActionGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActionGroupArgs)(nil)).Elem()
}


type LookupActionGroupResultOutput struct{ *pulumi.OutputState }

func (LookupActionGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActionGroupResult)(nil)).Elem()
}

func (o LookupActionGroupResultOutput) ToLookupActionGroupResultOutput() LookupActionGroupResultOutput {
	return o
}

func (o LookupActionGroupResultOutput) ToLookupActionGroupResultOutputWithContext(ctx context.Context) LookupActionGroupResultOutput {
	return o
}

func (o LookupActionGroupResultOutput) AutomationRunbookReceivers() AutomationRunbookReceiverResponseArrayOutput {
	return o.ApplyT(func(v LookupActionGroupResult) []AutomationRunbookReceiverResponse {
		return v.AutomationRunbookReceivers
	}).(AutomationRunbookReceiverResponseArrayOutput)
}

func (o LookupActionGroupResultOutput) AzureAppPushReceivers() AzureAppPushReceiverResponseArrayOutput {
	return o.ApplyT(func(v LookupActionGroupResult) []AzureAppPushReceiverResponse { return v.AzureAppPushReceivers }).(AzureAppPushReceiverResponseArrayOutput)
}

func (o LookupActionGroupResultOutput) EmailReceivers() EmailReceiverResponseArrayOutput {
	return o.ApplyT(func(v LookupActionGroupResult) []EmailReceiverResponse { return v.EmailReceivers }).(EmailReceiverResponseArrayOutput)
}

func (o LookupActionGroupResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupActionGroupResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupActionGroupResultOutput) GroupShortName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionGroupResult) string { return v.GroupShortName }).(pulumi.StringOutput)
}

func (o LookupActionGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupActionGroupResultOutput) ItsmReceivers() ItsmReceiverResponseArrayOutput {
	return o.ApplyT(func(v LookupActionGroupResult) []ItsmReceiverResponse { return v.ItsmReceivers }).(ItsmReceiverResponseArrayOutput)
}

func (o LookupActionGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupActionGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupActionGroupResultOutput) SmsReceivers() SmsReceiverResponseArrayOutput {
	return o.ApplyT(func(v LookupActionGroupResult) []SmsReceiverResponse { return v.SmsReceivers }).(SmsReceiverResponseArrayOutput)
}

func (o LookupActionGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupActionGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupActionGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActionGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupActionGroupResultOutput) WebhookReceivers() WebhookReceiverResponseArrayOutput {
	return o.ApplyT(func(v LookupActionGroupResult) []WebhookReceiverResponse { return v.WebhookReceivers }).(WebhookReceiverResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupActionGroupResultOutput{})
}
