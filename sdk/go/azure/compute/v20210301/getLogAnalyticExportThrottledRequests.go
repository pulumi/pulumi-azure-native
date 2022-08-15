


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetLogAnalyticExportThrottledRequests(ctx *pulumi.Context, args *GetLogAnalyticExportThrottledRequestsArgs, opts ...pulumi.InvokeOption) (*GetLogAnalyticExportThrottledRequestsResult, error) {
	var rv GetLogAnalyticExportThrottledRequestsResult
	err := ctx.Invoke("azure-native:compute/v20210301:getLogAnalyticExportThrottledRequests", args, &rv, opts...)
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

func GetLogAnalyticExportThrottledRequestsOutput(ctx *pulumi.Context, args GetLogAnalyticExportThrottledRequestsOutputArgs, opts ...pulumi.InvokeOption) GetLogAnalyticExportThrottledRequestsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLogAnalyticExportThrottledRequestsResult, error) {
			args := v.(GetLogAnalyticExportThrottledRequestsArgs)
			r, err := GetLogAnalyticExportThrottledRequests(ctx, &args, opts...)
			var s GetLogAnalyticExportThrottledRequestsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLogAnalyticExportThrottledRequestsResultOutput)
}

type GetLogAnalyticExportThrottledRequestsOutputArgs struct {
	BlobContainerSasUri        pulumi.StringInput  `pulumi:"blobContainerSasUri"`
	FromTime                   pulumi.StringInput  `pulumi:"fromTime"`
	GroupByClientApplicationId pulumi.BoolPtrInput `pulumi:"groupByClientApplicationId"`
	GroupByOperationName       pulumi.BoolPtrInput `pulumi:"groupByOperationName"`
	GroupByResourceName        pulumi.BoolPtrInput `pulumi:"groupByResourceName"`
	GroupByThrottlePolicy      pulumi.BoolPtrInput `pulumi:"groupByThrottlePolicy"`
	GroupByUserAgent           pulumi.BoolPtrInput `pulumi:"groupByUserAgent"`
	Location                   pulumi.StringInput  `pulumi:"location"`
	ToTime                     pulumi.StringInput  `pulumi:"toTime"`
}

func (GetLogAnalyticExportThrottledRequestsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogAnalyticExportThrottledRequestsArgs)(nil)).Elem()
}


type GetLogAnalyticExportThrottledRequestsResultOutput struct{ *pulumi.OutputState }

func (GetLogAnalyticExportThrottledRequestsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogAnalyticExportThrottledRequestsResult)(nil)).Elem()
}

func (o GetLogAnalyticExportThrottledRequestsResultOutput) ToGetLogAnalyticExportThrottledRequestsResultOutput() GetLogAnalyticExportThrottledRequestsResultOutput {
	return o
}

func (o GetLogAnalyticExportThrottledRequestsResultOutput) ToGetLogAnalyticExportThrottledRequestsResultOutputWithContext(ctx context.Context) GetLogAnalyticExportThrottledRequestsResultOutput {
	return o
}

func (o GetLogAnalyticExportThrottledRequestsResultOutput) Properties() LogAnalyticsOutputResponseOutput {
	return o.ApplyT(func(v GetLogAnalyticExportThrottledRequestsResult) LogAnalyticsOutputResponse { return v.Properties }).(LogAnalyticsOutputResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLogAnalyticExportThrottledRequestsResultOutput{})
}
