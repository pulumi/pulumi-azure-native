


package v20171201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetLogAnalyticExportRequestRateByInterval(ctx *pulumi.Context, args *GetLogAnalyticExportRequestRateByIntervalArgs, opts ...pulumi.InvokeOption) (*GetLogAnalyticExportRequestRateByIntervalResult, error) {
	var rv GetLogAnalyticExportRequestRateByIntervalResult
	err := ctx.Invoke("azure-native:compute/v20171201:getLogAnalyticExportRequestRateByInterval", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetLogAnalyticExportRequestRateByIntervalArgs struct {
	BlobContainerSasUri   string         `pulumi:"blobContainerSasUri"`
	FromTime              string         `pulumi:"fromTime"`
	GroupByOperationName  *bool          `pulumi:"groupByOperationName"`
	GroupByResourceName   *bool          `pulumi:"groupByResourceName"`
	GroupByThrottlePolicy *bool          `pulumi:"groupByThrottlePolicy"`
	IntervalLength        IntervalInMins `pulumi:"intervalLength"`
	Location              string         `pulumi:"location"`
	ToTime                string         `pulumi:"toTime"`
}


type GetLogAnalyticExportRequestRateByIntervalResult struct {
	EndTime    string                     `pulumi:"endTime"`
	Error      ApiErrorResponse           `pulumi:"error"`
	Name       string                     `pulumi:"name"`
	Properties LogAnalyticsOutputResponse `pulumi:"properties"`
	StartTime  string                     `pulumi:"startTime"`
	Status     string                     `pulumi:"status"`
}
