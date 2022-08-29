


package aadiam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureADMetric struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput                          `pulumi:"location"`
	Name       pulumi.StringOutput                          `pulumi:"name"`
	Properties AzureADMetricsPropertiesFormatResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                       `pulumi:"tags"`
	Type       pulumi.StringOutput                          `pulumi:"type"`
}


func NewAzureADMetric(ctx *pulumi.Context,
	name string, args *AzureADMetricArgs, opts ...pulumi.ResourceOption) (*AzureADMetric, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:aadiam/v20200701preview:azureADMetric"),
		},
	})
	opts = append(opts, aliases)
	var resource AzureADMetric
	err := ctx.RegisterResource("azure-native:aadiam:azureADMetric", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAzureADMetric(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureADMetricState, opts ...pulumi.ResourceOption) (*AzureADMetric, error) {
	var resource AzureADMetric
	err := ctx.ReadResource("azure-native:aadiam:azureADMetric", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type azureADMetricState struct {
}

type AzureADMetricState struct {
}

func (AzureADMetricState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureADMetricState)(nil)).Elem()
}

type azureADMetricArgs struct {
	AzureADMetricsName *string           `pulumi:"azureADMetricsName"`
	Location           *string           `pulumi:"location"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	Tags               map[string]string `pulumi:"tags"`
}


type AzureADMetricArgs struct {
	AzureADMetricsName pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	Tags               pulumi.StringMapInput
}

func (AzureADMetricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureADMetricArgs)(nil)).Elem()
}

type AzureADMetricInput interface {
	pulumi.Input

	ToAzureADMetricOutput() AzureADMetricOutput
	ToAzureADMetricOutputWithContext(ctx context.Context) AzureADMetricOutput
}

func (*AzureADMetric) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureADMetric)(nil)).Elem()
}

func (i *AzureADMetric) ToAzureADMetricOutput() AzureADMetricOutput {
	return i.ToAzureADMetricOutputWithContext(context.Background())
}

func (i *AzureADMetric) ToAzureADMetricOutputWithContext(ctx context.Context) AzureADMetricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureADMetricOutput)
}

type AzureADMetricOutput struct{ *pulumi.OutputState }

func (AzureADMetricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureADMetric)(nil)).Elem()
}

func (o AzureADMetricOutput) ToAzureADMetricOutput() AzureADMetricOutput {
	return o
}

func (o AzureADMetricOutput) ToAzureADMetricOutputWithContext(ctx context.Context) AzureADMetricOutput {
	return o
}

func (o AzureADMetricOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureADMetric) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AzureADMetricOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureADMetric) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AzureADMetricOutput) Properties() AzureADMetricsPropertiesFormatResponseOutput {
	return o.ApplyT(func(v *AzureADMetric) AzureADMetricsPropertiesFormatResponseOutput { return v.Properties }).(AzureADMetricsPropertiesFormatResponseOutput)
}

func (o AzureADMetricOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AzureADMetric) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AzureADMetricOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureADMetric) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureADMetricOutput{})
}
