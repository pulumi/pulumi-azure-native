


package managedidentity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureResourceResponse struct {
	Id                      string `pulumi:"id"`
	Name                    string `pulumi:"name"`
	ResourceGroup           string `pulumi:"resourceGroup"`
	SubscriptionDisplayName string `pulumi:"subscriptionDisplayName"`
	SubscriptionId          string `pulumi:"subscriptionId"`
	Type                    string `pulumi:"type"`
}

type AzureResourceResponseOutput struct{ *pulumi.OutputState }

func (AzureResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceResponse)(nil)).Elem()
}

func (o AzureResourceResponseOutput) ToAzureResourceResponseOutput() AzureResourceResponseOutput {
	return o
}

func (o AzureResourceResponseOutput) ToAzureResourceResponseOutputWithContext(ctx context.Context) AzureResourceResponseOutput {
	return o
}

func (o AzureResourceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o AzureResourceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o AzureResourceResponseOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceResponse) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o AzureResourceResponseOutput) SubscriptionDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceResponse) string { return v.SubscriptionDisplayName }).(pulumi.StringOutput)
}

func (o AzureResourceResponseOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceResponse) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o AzureResourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type AzureResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureResourceResponse)(nil)).Elem()
}

func (o AzureResourceResponseArrayOutput) ToAzureResourceResponseArrayOutput() AzureResourceResponseArrayOutput {
	return o
}

func (o AzureResourceResponseArrayOutput) ToAzureResourceResponseArrayOutputWithContext(ctx context.Context) AzureResourceResponseArrayOutput {
	return o
}

func (o AzureResourceResponseArrayOutput) Index(i pulumi.IntInput) AzureResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureResourceResponse {
		return vs[0].([]AzureResourceResponse)[vs[1].(int)]
	}).(AzureResourceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureResourceResponseOutput{})
	pulumi.RegisterOutputType(AzureResourceResponseArrayOutput{})
}
