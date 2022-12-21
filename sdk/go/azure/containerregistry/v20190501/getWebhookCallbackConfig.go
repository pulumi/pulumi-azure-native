


package v20190501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetWebhookCallbackConfig(ctx *pulumi.Context, args *GetWebhookCallbackConfigArgs, opts ...pulumi.InvokeOption) (*GetWebhookCallbackConfigResult, error) {
	var rv GetWebhookCallbackConfigResult
	err := ctx.Invoke("azure-native:containerregistry/v20190501:getWebhookCallbackConfig", args, &rv, opts...)
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

func GetWebhookCallbackConfigOutput(ctx *pulumi.Context, args GetWebhookCallbackConfigOutputArgs, opts ...pulumi.InvokeOption) GetWebhookCallbackConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetWebhookCallbackConfigResult, error) {
			args := v.(GetWebhookCallbackConfigArgs)
			r, err := GetWebhookCallbackConfig(ctx, &args, opts...)
			var s GetWebhookCallbackConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetWebhookCallbackConfigResultOutput)
}

type GetWebhookCallbackConfigOutputArgs struct {
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WebhookName       pulumi.StringInput `pulumi:"webhookName"`
}

func (GetWebhookCallbackConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWebhookCallbackConfigArgs)(nil)).Elem()
}


type GetWebhookCallbackConfigResultOutput struct{ *pulumi.OutputState }

func (GetWebhookCallbackConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWebhookCallbackConfigResult)(nil)).Elem()
}

func (o GetWebhookCallbackConfigResultOutput) ToGetWebhookCallbackConfigResultOutput() GetWebhookCallbackConfigResultOutput {
	return o
}

func (o GetWebhookCallbackConfigResultOutput) ToGetWebhookCallbackConfigResultOutputWithContext(ctx context.Context) GetWebhookCallbackConfigResultOutput {
	return o
}

func (o GetWebhookCallbackConfigResultOutput) CustomHeaders() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetWebhookCallbackConfigResult) map[string]string { return v.CustomHeaders }).(pulumi.StringMapOutput)
}

func (o GetWebhookCallbackConfigResultOutput) ServiceUri() pulumi.StringOutput {
	return o.ApplyT(func(v GetWebhookCallbackConfigResult) string { return v.ServiceUri }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWebhookCallbackConfigResultOutput{})
}
