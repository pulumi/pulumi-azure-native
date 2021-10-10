


package aadiam

import (
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
