


package costmanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReport(ctx *pulumi.Context, args *LookupReportArgs, opts ...pulumi.InvokeOption) (*LookupReportResult, error) {
	var rv LookupReportResult
	err := ctx.Invoke("azure-native:costmanagement:getReport", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReportArgs struct {
	ReportName string `pulumi:"reportName"`
}


type LookupReportResult struct {
	Definition   ReportDefinitionResponse   `pulumi:"definition"`
	DeliveryInfo ReportDeliveryInfoResponse `pulumi:"deliveryInfo"`
	Format       *string                    `pulumi:"format"`
	Id           string                     `pulumi:"id"`
	Name         string                     `pulumi:"name"`
	Schedule     *ReportScheduleResponse    `pulumi:"schedule"`
	Tags         map[string]string          `pulumi:"tags"`
	Type         string                     `pulumi:"type"`
}

func LookupReportOutput(ctx *pulumi.Context, args LookupReportOutputArgs, opts ...pulumi.InvokeOption) LookupReportResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReportResult, error) {
			args := v.(LookupReportArgs)
			r, err := LookupReport(ctx, &args, opts...)
			var s LookupReportResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReportResultOutput)
}

type LookupReportOutputArgs struct {
	ReportName pulumi.StringInput `pulumi:"reportName"`
}

func (LookupReportOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReportArgs)(nil)).Elem()
}


type LookupReportResultOutput struct{ *pulumi.OutputState }

func (LookupReportResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReportResult)(nil)).Elem()
}

func (o LookupReportResultOutput) ToLookupReportResultOutput() LookupReportResultOutput {
	return o
}

func (o LookupReportResultOutput) ToLookupReportResultOutputWithContext(ctx context.Context) LookupReportResultOutput {
	return o
}

func (o LookupReportResultOutput) Definition() ReportDefinitionResponseOutput {
	return o.ApplyT(func(v LookupReportResult) ReportDefinitionResponse { return v.Definition }).(ReportDefinitionResponseOutput)
}

func (o LookupReportResultOutput) DeliveryInfo() ReportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v LookupReportResult) ReportDeliveryInfoResponse { return v.DeliveryInfo }).(ReportDeliveryInfoResponseOutput)
}

func (o LookupReportResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReportResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o LookupReportResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReportResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReportResultOutput) Schedule() ReportScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupReportResult) *ReportScheduleResponse { return v.Schedule }).(ReportScheduleResponsePtrOutput)
}

func (o LookupReportResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupReportResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupReportResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReportResultOutput{})
}
