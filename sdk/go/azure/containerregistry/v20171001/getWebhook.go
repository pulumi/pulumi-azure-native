


package v20171001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebhook(ctx *pulumi.Context, args *LookupWebhookArgs, opts ...pulumi.InvokeOption) (*LookupWebhookResult, error) {
	var rv LookupWebhookResult
	err := ctx.Invoke("azure-native:containerregistry/v20171001:getWebhook", args, &rv, opts...)
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
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WebhookName       pulumi.StringInput `pulumi:"webhookName"`
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

func (o LookupWebhookResultOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebhookResult) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

func (o LookupWebhookResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhookResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebhookResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhookResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupWebhookResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhookResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebhookResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhookResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWebhookResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebhookResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o LookupWebhookResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebhookResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupWebhookResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebhookResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWebhookResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebhookResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebhookResultOutput{})
}
