


package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetEntityInsights(ctx *pulumi.Context, args *GetEntityInsightsArgs, opts ...pulumi.InvokeOption) (*GetEntityInsightsResult, error) {
	var rv GetEntityInsightsResult
	err := ctx.Invoke("azure-native:securityinsights/v20221201preview:getEntityInsights", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetEntityInsightsArgs struct {
	AddDefaultExtendedTimeRange *bool    `pulumi:"addDefaultExtendedTimeRange"`
	EndTime                     string   `pulumi:"endTime"`
	EntityId                    string   `pulumi:"entityId"`
	InsightQueryIds             []string `pulumi:"insightQueryIds"`
	ResourceGroupName           string   `pulumi:"resourceGroupName"`
	StartTime                   string   `pulumi:"startTime"`
	WorkspaceName               string   `pulumi:"workspaceName"`
}


type GetEntityInsightsResult struct {
	MetaData *GetInsightsResultsMetadataResponse `pulumi:"metaData"`
	Value    []EntityInsightItemResponse         `pulumi:"value"`
}

func GetEntityInsightsOutput(ctx *pulumi.Context, args GetEntityInsightsOutputArgs, opts ...pulumi.InvokeOption) GetEntityInsightsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEntityInsightsResult, error) {
			args := v.(GetEntityInsightsArgs)
			r, err := GetEntityInsights(ctx, &args, opts...)
			var s GetEntityInsightsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEntityInsightsResultOutput)
}

type GetEntityInsightsOutputArgs struct {
	AddDefaultExtendedTimeRange pulumi.BoolPtrInput     `pulumi:"addDefaultExtendedTimeRange"`
	EndTime                     pulumi.StringInput      `pulumi:"endTime"`
	EntityId                    pulumi.StringInput      `pulumi:"entityId"`
	InsightQueryIds             pulumi.StringArrayInput `pulumi:"insightQueryIds"`
	ResourceGroupName           pulumi.StringInput      `pulumi:"resourceGroupName"`
	StartTime                   pulumi.StringInput      `pulumi:"startTime"`
	WorkspaceName               pulumi.StringInput      `pulumi:"workspaceName"`
}

func (GetEntityInsightsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEntityInsightsArgs)(nil)).Elem()
}


type GetEntityInsightsResultOutput struct{ *pulumi.OutputState }

func (GetEntityInsightsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEntityInsightsResult)(nil)).Elem()
}

func (o GetEntityInsightsResultOutput) ToGetEntityInsightsResultOutput() GetEntityInsightsResultOutput {
	return o
}

func (o GetEntityInsightsResultOutput) ToGetEntityInsightsResultOutputWithContext(ctx context.Context) GetEntityInsightsResultOutput {
	return o
}

func (o GetEntityInsightsResultOutput) MetaData() GetInsightsResultsMetadataResponsePtrOutput {
	return o.ApplyT(func(v GetEntityInsightsResult) *GetInsightsResultsMetadataResponse { return v.MetaData }).(GetInsightsResultsMetadataResponsePtrOutput)
}

func (o GetEntityInsightsResultOutput) Value() EntityInsightItemResponseArrayOutput {
	return o.ApplyT(func(v GetEntityInsightsResult) []EntityInsightItemResponse { return v.Value }).(EntityInsightItemResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEntityInsightsResultOutput{})
}
