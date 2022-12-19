


package v20221001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupView(ctx *pulumi.Context, args *LookupViewArgs, opts ...pulumi.InvokeOption) (*LookupViewResult, error) {
	var rv LookupViewResult
	err := ctx.Invoke("azure-native:costmanagement/v20221001:getView", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupViewArgs struct {
	ViewName string `pulumi:"viewName"`
}


type LookupViewResult struct {
	Accumulated               *string                         `pulumi:"accumulated"`
	Chart                     *string                         `pulumi:"chart"`
	CreatedOn                 string                          `pulumi:"createdOn"`
	Currency                  string                          `pulumi:"currency"`
	DataSet                   *ReportConfigDatasetResponse    `pulumi:"dataSet"`
	DateRange                 string                          `pulumi:"dateRange"`
	DisplayName               *string                         `pulumi:"displayName"`
	ETag                      *string                         `pulumi:"eTag"`
	Id                        string                          `pulumi:"id"`
	IncludeMonetaryCommitment *bool                           `pulumi:"includeMonetaryCommitment"`
	Kpis                      []KpiPropertiesResponse         `pulumi:"kpis"`
	Metric                    *string                         `pulumi:"metric"`
	ModifiedOn                string                          `pulumi:"modifiedOn"`
	Name                      string                          `pulumi:"name"`
	Pivots                    []PivotPropertiesResponse       `pulumi:"pivots"`
	Scope                     *string                         `pulumi:"scope"`
	TimePeriod                *ReportConfigTimePeriodResponse `pulumi:"timePeriod"`
	Timeframe                 string                          `pulumi:"timeframe"`
	Type                      string                          `pulumi:"type"`
}

func LookupViewOutput(ctx *pulumi.Context, args LookupViewOutputArgs, opts ...pulumi.InvokeOption) LookupViewResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupViewResult, error) {
			args := v.(LookupViewArgs)
			r, err := LookupView(ctx, &args, opts...)
			var s LookupViewResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupViewResultOutput)
}

type LookupViewOutputArgs struct {
	ViewName pulumi.StringInput `pulumi:"viewName"`
}

func (LookupViewOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupViewArgs)(nil)).Elem()
}


type LookupViewResultOutput struct{ *pulumi.OutputState }

func (LookupViewResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupViewResult)(nil)).Elem()
}

func (o LookupViewResultOutput) ToLookupViewResultOutput() LookupViewResultOutput {
	return o
}

func (o LookupViewResultOutput) ToLookupViewResultOutputWithContext(ctx context.Context) LookupViewResultOutput {
	return o
}

func (o LookupViewResultOutput) Accumulated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.Accumulated }).(pulumi.StringPtrOutput)
}

func (o LookupViewResultOutput) Chart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.Chart }).(pulumi.StringPtrOutput)
}

func (o LookupViewResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o LookupViewResultOutput) Currency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Currency }).(pulumi.StringOutput)
}

func (o LookupViewResultOutput) DataSet() ReportConfigDatasetResponsePtrOutput {
	return o.ApplyT(func(v LookupViewResult) *ReportConfigDatasetResponse { return v.DataSet }).(ReportConfigDatasetResponsePtrOutput)
}

func (o LookupViewResultOutput) DateRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.DateRange }).(pulumi.StringOutput)
}

func (o LookupViewResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupViewResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupViewResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupViewResultOutput) IncludeMonetaryCommitment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *bool { return v.IncludeMonetaryCommitment }).(pulumi.BoolPtrOutput)
}

func (o LookupViewResultOutput) Kpis() KpiPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupViewResult) []KpiPropertiesResponse { return v.Kpis }).(KpiPropertiesResponseArrayOutput)
}

func (o LookupViewResultOutput) Metric() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.Metric }).(pulumi.StringPtrOutput)
}

func (o LookupViewResultOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o LookupViewResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupViewResultOutput) Pivots() PivotPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupViewResult) []PivotPropertiesResponse { return v.Pivots }).(PivotPropertiesResponseArrayOutput)
}

func (o LookupViewResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o LookupViewResultOutput) TimePeriod() ReportConfigTimePeriodResponsePtrOutput {
	return o.ApplyT(func(v LookupViewResult) *ReportConfigTimePeriodResponse { return v.TimePeriod }).(ReportConfigTimePeriodResponsePtrOutput)
}

func (o LookupViewResultOutput) Timeframe() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Timeframe }).(pulumi.StringOutput)
}

func (o LookupViewResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupViewResultOutput{})
}
