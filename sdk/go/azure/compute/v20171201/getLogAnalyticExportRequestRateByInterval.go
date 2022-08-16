


package v20171201

import (
	"context"
	"reflect"

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

func GetLogAnalyticExportRequestRateByIntervalOutput(ctx *pulumi.Context, args GetLogAnalyticExportRequestRateByIntervalOutputArgs, opts ...pulumi.InvokeOption) GetLogAnalyticExportRequestRateByIntervalResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLogAnalyticExportRequestRateByIntervalResult, error) {
			args := v.(GetLogAnalyticExportRequestRateByIntervalArgs)
			r, err := GetLogAnalyticExportRequestRateByInterval(ctx, &args, opts...)
			var s GetLogAnalyticExportRequestRateByIntervalResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLogAnalyticExportRequestRateByIntervalResultOutput)
}

type GetLogAnalyticExportRequestRateByIntervalOutputArgs struct {
	BlobContainerSasUri   pulumi.StringInput  `pulumi:"blobContainerSasUri"`
	FromTime              pulumi.StringInput  `pulumi:"fromTime"`
	GroupByOperationName  pulumi.BoolPtrInput `pulumi:"groupByOperationName"`
	GroupByResourceName   pulumi.BoolPtrInput `pulumi:"groupByResourceName"`
	GroupByThrottlePolicy pulumi.BoolPtrInput `pulumi:"groupByThrottlePolicy"`
	IntervalLength        IntervalInMinsInput `pulumi:"intervalLength"`
	Location              pulumi.StringInput  `pulumi:"location"`
	ToTime                pulumi.StringInput  `pulumi:"toTime"`
}

func (GetLogAnalyticExportRequestRateByIntervalOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogAnalyticExportRequestRateByIntervalArgs)(nil)).Elem()
}


type GetLogAnalyticExportRequestRateByIntervalResultOutput struct{ *pulumi.OutputState }

func (GetLogAnalyticExportRequestRateByIntervalResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogAnalyticExportRequestRateByIntervalResult)(nil)).Elem()
}

func (o GetLogAnalyticExportRequestRateByIntervalResultOutput) ToGetLogAnalyticExportRequestRateByIntervalResultOutput() GetLogAnalyticExportRequestRateByIntervalResultOutput {
	return o
}

func (o GetLogAnalyticExportRequestRateByIntervalResultOutput) ToGetLogAnalyticExportRequestRateByIntervalResultOutputWithContext(ctx context.Context) GetLogAnalyticExportRequestRateByIntervalResultOutput {
	return o
}

func (o GetLogAnalyticExportRequestRateByIntervalResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogAnalyticExportRequestRateByIntervalResult) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o GetLogAnalyticExportRequestRateByIntervalResultOutput) Error() ApiErrorResponseOutput {
	return o.ApplyT(func(v GetLogAnalyticExportRequestRateByIntervalResult) ApiErrorResponse { return v.Error }).(ApiErrorResponseOutput)
}

func (o GetLogAnalyticExportRequestRateByIntervalResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogAnalyticExportRequestRateByIntervalResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetLogAnalyticExportRequestRateByIntervalResultOutput) Properties() LogAnalyticsOutputResponseOutput {
	return o.ApplyT(func(v GetLogAnalyticExportRequestRateByIntervalResult) LogAnalyticsOutputResponse {
		return v.Properties
	}).(LogAnalyticsOutputResponseOutput)
}

func (o GetLogAnalyticExportRequestRateByIntervalResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogAnalyticExportRequestRateByIntervalResult) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o GetLogAnalyticExportRequestRateByIntervalResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogAnalyticExportRequestRateByIntervalResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLogAnalyticExportRequestRateByIntervalResultOutput{})
}
