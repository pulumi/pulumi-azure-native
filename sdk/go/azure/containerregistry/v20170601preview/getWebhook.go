


package v20170601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebhook(ctx *pulumi.Context, args *LookupWebhookArgs, opts ...pulumi.InvokeOption) (*LookupWebhookResult, error) {
	var rv LookupWebhookResult
	err := ctx.Invoke("azure-native:containerregistry/v20170601preview:getWebhook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebhookArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WebhookName       string `pulumi:"webhookName"`
}


type LookupWebhookResult struct {
	Actions           []string          `pulumi:"actions"`
	Id                string            `pulumi:"id"`
	Location          string            `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	Scope             *string           `pulumi:"scope"`
	Status            *string           `pulumi:"status"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}
