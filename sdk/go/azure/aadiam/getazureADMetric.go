


package aadiam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetazureADMetric(ctx *pulumi.Context, args *GetazureADMetricArgs, opts ...pulumi.InvokeOption) (*GetazureADMetricResult, error) {
	var rv GetazureADMetricResult
	err := ctx.Invoke("azure-native:aadiam:getazureADMetric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetazureADMetricArgs struct {
	AzureADMetricsName string `pulumi:"azureADMetricsName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type GetazureADMetricResult struct {
	Id         string                                 `pulumi:"id"`
	Location   string                                 `pulumi:"location"`
	Name       string                                 `pulumi:"name"`
	Properties AzureADMetricsPropertiesFormatResponse `pulumi:"properties"`
	Tags       map[string]string                      `pulumi:"tags"`
	Type       string                                 `pulumi:"type"`
}

func GetazureADMetricOutput(ctx *pulumi.Context, args GetazureADMetricOutputArgs, opts ...pulumi.InvokeOption) GetazureADMetricResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetazureADMetricResult, error) {
			args := v.(GetazureADMetricArgs)
			r, err := GetazureADMetric(ctx, &args, opts...)
			var s GetazureADMetricResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetazureADMetricResultOutput)
}

type GetazureADMetricOutputArgs struct {
	AzureADMetricsName pulumi.StringInput `pulumi:"azureADMetricsName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetazureADMetricOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetazureADMetricArgs)(nil)).Elem()
}


type GetazureADMetricResultOutput struct{ *pulumi.OutputState }

func (GetazureADMetricResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetazureADMetricResult)(nil)).Elem()
}

func (o GetazureADMetricResultOutput) ToGetazureADMetricResultOutput() GetazureADMetricResultOutput {
	return o
}

func (o GetazureADMetricResultOutput) ToGetazureADMetricResultOutputWithContext(ctx context.Context) GetazureADMetricResultOutput {
	return o
}

func (o GetazureADMetricResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetazureADMetricResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetazureADMetricResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetazureADMetricResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetazureADMetricResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetazureADMetricResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetazureADMetricResultOutput) Properties() AzureADMetricsPropertiesFormatResponseOutput {
	return o.ApplyT(func(v GetazureADMetricResult) AzureADMetricsPropertiesFormatResponse { return v.Properties }).(AzureADMetricsPropertiesFormatResponseOutput)
}

func (o GetazureADMetricResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetazureADMetricResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetazureADMetricResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetazureADMetricResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetazureADMetricResultOutput{})
}
