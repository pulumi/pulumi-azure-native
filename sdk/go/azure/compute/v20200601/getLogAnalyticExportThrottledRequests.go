


package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetLogAnalyticExportThrottledRequests(ctx *pulumi.Context, args *GetLogAnalyticExportThrottledRequestsArgs, opts ...pulumi.InvokeOption) (*GetLogAnalyticExportThrottledRequestsResult, error) {
	var rv GetLogAnalyticExportThrottledRequestsResult
	err := ctx.Invoke("azure-native:compute/v20200601:getLogAnalyticExportThrottledRequests", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetLogAnalyticExportThrottledRequestsArgs struct {
	BlobContainerSasUri        string `pulumi:"blobContainerSasUri"`
	FromTime                   string `pulumi:"fromTime"`
	GroupByClientApplicationId *bool  `pulumi:"groupByClientApplicationId"`
	GroupByOperationName       *bool  `pulumi:"groupByOperationName"`
	GroupByResourceName        *bool  `pulumi:"groupByResourceName"`
	GroupByThrottlePolicy      *bool  `pulumi:"groupByThrottlePolicy"`
	GroupByUserAgent           *bool  `pulumi:"groupByUserAgent"`
	Location                   string `pulumi:"location"`
	ToTime                     string `pulumi:"toTime"`
}


type GetLogAnalyticExportThrottledRequestsResult struct {
	Properties LogAnalyticsOutputResponse `pulumi:"properties"`
}
