


package v20221116preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReport(ctx *pulumi.Context, args *LookupReportArgs, opts ...pulumi.InvokeOption) (*LookupReportResult, error) {
	var rv LookupReportResult
	err := ctx.Invoke("azure-native:appcomplianceautomation/v20221116preview:getReport", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReportArgs struct {
	ReportName string `pulumi:"reportName"`
}


type LookupReportResult struct {
	Id         string                   `pulumi:"id"`
	Name       string                   `pulumi:"name"`
	Properties ReportPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse       `pulumi:"systemData"`
	Type       string                   `pulumi:"type"`
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

func (o LookupReportResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReportResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReportResultOutput) Properties() ReportPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReportResult) ReportPropertiesResponse { return v.Properties }).(ReportPropertiesResponseOutput)
}

func (o LookupReportResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupReportResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupReportResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReportResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReportResultOutput{})
}
