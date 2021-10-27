


package v20210801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetWebhookCallbackConfig(ctx *pulumi.Context, args *GetWebhookCallbackConfigArgs, opts ...pulumi.InvokeOption) (*GetWebhookCallbackConfigResult, error) {
	var rv GetWebhookCallbackConfigResult
	err := ctx.Invoke("azure-native:containerregistry/v20210801preview:getWebhookCallbackConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetWebhookCallbackConfigArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WebhookName       string `pulumi:"webhookName"`
}


type GetWebhookCallbackConfigResult struct {
	CustomHeaders map[string]string `pulumi:"customHeaders"`
	ServiceUri    string            `pulumi:"serviceUri"`
}
