


package v20200701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureADMetricsPropertiesFormatResponse struct {
	ProvisioningState string `pulumi:"provisioningState"`
}





type AzureADMetricsPropertiesFormatResponseInput interface {
	pulumi.Input

	ToAzureADMetricsPropertiesFormatResponseOutput() AzureADMetricsPropertiesFormatResponseOutput
	ToAzureADMetricsPropertiesFormatResponseOutputWithContext(context.Context) AzureADMetricsPropertiesFormatResponseOutput
}

type AzureADMetricsPropertiesFormatResponseArgs struct {
	ProvisioningState pulumi.StringInput `pulumi:"provisioningState"`
}

func (AzureADMetricsPropertiesFormatResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureADMetricsPropertiesFormatResponse)(nil)).Elem()
}

func (i AzureADMetricsPropertiesFormatResponseArgs) ToAzureADMetricsPropertiesFormatResponseOutput() AzureADMetricsPropertiesFormatResponseOutput {
	return i.ToAzureADMetricsPropertiesFormatResponseOutputWithContext(context.Background())
}

func (i AzureADMetricsPropertiesFormatResponseArgs) ToAzureADMetricsPropertiesFormatResponseOutputWithContext(ctx context.Context) AzureADMetricsPropertiesFormatResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureADMetricsPropertiesFormatResponseOutput)
}

func (i AzureADMetricsPropertiesFormatResponseArgs) ToAzureADMetricsPropertiesFormatResponsePtrOutput() AzureADMetricsPropertiesFormatResponsePtrOutput {
	return i.ToAzureADMetricsPropertiesFormatResponsePtrOutputWithContext(context.Background())
}

func (i AzureADMetricsPropertiesFormatResponseArgs) ToAzureADMetricsPropertiesFormatResponsePtrOutputWithContext(ctx context.Context) AzureADMetricsPropertiesFormatResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureADMetricsPropertiesFormatResponseOutput).ToAzureADMetricsPropertiesFormatResponsePtrOutputWithContext(ctx)
}









type AzureADMetricsPropertiesFormatResponsePtrInput interface {
	pulumi.Input

	ToAzureADMetricsPropertiesFormatResponsePtrOutput() AzureADMetricsPropertiesFormatResponsePtrOutput
	ToAzureADMetricsPropertiesFormatResponsePtrOutputWithContext(context.Context) AzureADMetricsPropertiesFormatResponsePtrOutput
}

type azureADMetricsPropertiesFormatResponsePtrType AzureADMetricsPropertiesFormatResponseArgs

func AzureADMetricsPropertiesFormatResponsePtr(v *AzureADMetricsPropertiesFormatResponseArgs) AzureADMetricsPropertiesFormatResponsePtrInput {
	return (*azureADMetricsPropertiesFormatResponsePtrType)(v)
}

func (*azureADMetricsPropertiesFormatResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureADMetricsPropertiesFormatResponse)(nil)).Elem()
}

func (i *azureADMetricsPropertiesFormatResponsePtrType) ToAzureADMetricsPropertiesFormatResponsePtrOutput() AzureADMetricsPropertiesFormatResponsePtrOutput {
	return i.ToAzureADMetricsPropertiesFormatResponsePtrOutputWithContext(context.Background())
}

func (i *azureADMetricsPropertiesFormatResponsePtrType) ToAzureADMetricsPropertiesFormatResponsePtrOutputWithContext(ctx context.Context) AzureADMetricsPropertiesFormatResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureADMetricsPropertiesFormatResponsePtrOutput)
}

type AzureADMetricsPropertiesFormatResponseOutput struct{ *pulumi.OutputState }

func (AzureADMetricsPropertiesFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureADMetricsPropertiesFormatResponse)(nil)).Elem()
}

func (o AzureADMetricsPropertiesFormatResponseOutput) ToAzureADMetricsPropertiesFormatResponseOutput() AzureADMetricsPropertiesFormatResponseOutput {
	return o
}

func (o AzureADMetricsPropertiesFormatResponseOutput) ToAzureADMetricsPropertiesFormatResponseOutputWithContext(ctx context.Context) AzureADMetricsPropertiesFormatResponseOutput {
	return o
}

func (o AzureADMetricsPropertiesFormatResponseOutput) ToAzureADMetricsPropertiesFormatResponsePtrOutput() AzureADMetricsPropertiesFormatResponsePtrOutput {
	return o.ToAzureADMetricsPropertiesFormatResponsePtrOutputWithContext(context.Background())
}

func (o AzureADMetricsPropertiesFormatResponseOutput) ToAzureADMetricsPropertiesFormatResponsePtrOutputWithContext(ctx context.Context) AzureADMetricsPropertiesFormatResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureADMetricsPropertiesFormatResponse) *AzureADMetricsPropertiesFormatResponse {
		return &v
	}).(AzureADMetricsPropertiesFormatResponsePtrOutput)
}

func (o AzureADMetricsPropertiesFormatResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AzureADMetricsPropertiesFormatResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type AzureADMetricsPropertiesFormatResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureADMetricsPropertiesFormatResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureADMetricsPropertiesFormatResponse)(nil)).Elem()
}

func (o AzureADMetricsPropertiesFormatResponsePtrOutput) ToAzureADMetricsPropertiesFormatResponsePtrOutput() AzureADMetricsPropertiesFormatResponsePtrOutput {
	return o
}

func (o AzureADMetricsPropertiesFormatResponsePtrOutput) ToAzureADMetricsPropertiesFormatResponsePtrOutputWithContext(ctx context.Context) AzureADMetricsPropertiesFormatResponsePtrOutput {
	return o
}

func (o AzureADMetricsPropertiesFormatResponsePtrOutput) Elem() AzureADMetricsPropertiesFormatResponseOutput {
	return o.ApplyT(func(v *AzureADMetricsPropertiesFormatResponse) AzureADMetricsPropertiesFormatResponse {
		if v != nil {
			return *v
		}
		var ret AzureADMetricsPropertiesFormatResponse
		return ret
	}).(AzureADMetricsPropertiesFormatResponseOutput)
}

func (o AzureADMetricsPropertiesFormatResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureADMetricsPropertiesFormatResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureADMetricsPropertiesFormatResponseOutput{})
	pulumi.RegisterOutputType(AzureADMetricsPropertiesFormatResponsePtrOutput{})
}
