


package v20190301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupActionGroup(ctx *pulumi.Context, args *LookupActionGroupArgs, opts ...pulumi.InvokeOption) (*LookupActionGroupResult, error) {
	var rv LookupActionGroupResult
	err := ctx.Invoke("azure-native:insights/v20190301:getActionGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupActionGroupArgs struct {
	ActionGroupName   string `pulumi:"actionGroupName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupActionGroupResult struct {
	ArmRoleReceivers           []ArmRoleReceiverResponse           `pulumi:"armRoleReceivers"`
	AutomationRunbookReceivers []AutomationRunbookReceiverResponse `pulumi:"automationRunbookReceivers"`
	AzureAppPushReceivers      []AzureAppPushReceiverResponse      `pulumi:"azureAppPushReceivers"`
	AzureFunctionReceivers     []AzureFunctionReceiverResponse     `pulumi:"azureFunctionReceivers"`
	EmailReceivers             []EmailReceiverResponse             `pulumi:"emailReceivers"`
	Enabled                    bool                                `pulumi:"enabled"`
	GroupShortName             string                              `pulumi:"groupShortName"`
	Id                         string                              `pulumi:"id"`
	Identity                   string                              `pulumi:"identity"`
	ItsmReceivers              []ItsmReceiverResponse              `pulumi:"itsmReceivers"`
	Kind                       string                              `pulumi:"kind"`
	Location                   string                              `pulumi:"location"`
	LogicAppReceivers          []LogicAppReceiverResponse          `pulumi:"logicAppReceivers"`
	Name                       string                              `pulumi:"name"`
	SmsReceivers               []SmsReceiverResponse               `pulumi:"smsReceivers"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
	VoiceReceivers             []VoiceReceiverResponse             `pulumi:"voiceReceivers"`
	WebhookReceivers           []WebhookReceiverResponse           `pulumi:"webhookReceivers"`
}
