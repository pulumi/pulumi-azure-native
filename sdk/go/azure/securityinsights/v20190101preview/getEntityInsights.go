


package v20190101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetEntityInsights(ctx *pulumi.Context, args *GetEntityInsightsArgs, opts ...pulumi.InvokeOption) (*GetEntityInsightsResult, error) {
	var rv GetEntityInsightsResult
	err := ctx.Invoke("azure-native:securityinsights/v20190101preview:getEntityInsights", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetEntityInsightsArgs struct {
	AddDefaultExtendedTimeRange         *bool    `pulumi:"addDefaultExtendedTimeRange"`
	EndTime                             string   `pulumi:"endTime"`
	EntityId                            string   `pulumi:"entityId"`
	InsightQueryIds                     []string `pulumi:"insightQueryIds"`
	OperationalInsightsResourceProvider string   `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string   `pulumi:"resourceGroupName"`
	StartTime                           string   `pulumi:"startTime"`
	WorkspaceName                       string   `pulumi:"workspaceName"`
}


type GetEntityInsightsResult struct {
	MetaData *GetInsightsResultsMetadataResponse `pulumi:"metaData"`
	Value    []EntityInsightItemResponse         `pulumi:"value"`
}
