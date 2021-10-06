


package automation

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebhook(ctx *pulumi.Context, args *LookupWebhookArgs, opts ...pulumi.InvokeOption) (*LookupWebhookResult, error) {
	var rv LookupWebhookResult
	err := ctx.Invoke("azure-native:automation:getWebhook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
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
