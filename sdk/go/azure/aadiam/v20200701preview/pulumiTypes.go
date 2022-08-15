


package v20200701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureADMetricsPropertiesFormatResponse struct {
	ProvisioningState string `pulumi:"provisioningState"`
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

func (o AzureADMetricsPropertiesFormatResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AzureADMetricsPropertiesFormatResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureADMetricsPropertiesFormatResponseOutput{})
}
