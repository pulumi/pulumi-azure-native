


package v20170401

import (
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
