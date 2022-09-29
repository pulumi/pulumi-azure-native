


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupViewByScope(ctx *pulumi.Context, args *LookupViewByScopeArgs, opts ...pulumi.InvokeOption) (*LookupViewByScopeResult, error) {
	var rv LookupViewByScopeResult
	err := ctx.Invoke("azure-native:costmanagement/v20221001preview:getViewByScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupViewByScopeArgs struct {
	Scope    string `pulumi:"scope"`
	ViewName string `pulumi:"viewName"`
}


type LookupViewByScopeResult struct {
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

func LookupViewByScopeOutput(ctx *pulumi.Context, args LookupViewByScopeOutputArgs, opts ...pulumi.InvokeOption) LookupViewByScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupViewByScopeResult, error) {
			args := v.(LookupViewByScopeArgs)
			r, err := LookupViewByScope(ctx, &args, opts...)
			var s LookupViewByScopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupViewByScopeResultOutput)
}

type LookupViewByScopeOutputArgs struct {
	Scope    pulumi.StringInput `pulumi:"scope"`
	ViewName pulumi.StringInput `pulumi:"viewName"`
}

func (LookupViewByScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupViewByScopeArgs)(nil)).Elem()
}


type LookupViewByScopeResultOutput struct{ *pulumi.OutputState }

func (LookupViewByScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupViewByScopeResult)(nil)).Elem()
}

func (o LookupViewByScopeResultOutput) ToLookupViewByScopeResultOutput() LookupViewByScopeResultOutput {
	return o
}

func (o LookupViewByScopeResultOutput) ToLookupViewByScopeResultOutputWithContext(ctx context.Context) LookupViewByScopeResultOutput {
	return o
}

func (o LookupViewByScopeResultOutput) Accumulated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *string { return v.Accumulated }).(pulumi.StringPtrOutput)
}

func (o LookupViewByScopeResultOutput) Chart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *string { return v.Chart }).(pulumi.StringPtrOutput)
}

func (o LookupViewByScopeResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o LookupViewByScopeResultOutput) Currency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) string { return v.Currency }).(pulumi.StringOutput)
}

func (o LookupViewByScopeResultOutput) DataSet() ReportConfigDatasetResponsePtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *ReportConfigDatasetResponse { return v.DataSet }).(ReportConfigDatasetResponsePtrOutput)
}

func (o LookupViewByScopeResultOutput) DateRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) string { return v.DateRange }).(pulumi.StringOutput)
}

func (o LookupViewByScopeResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupViewByScopeResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupViewByScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupViewByScopeResultOutput) IncludeMonetaryCommitment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *bool { return v.IncludeMonetaryCommitment }).(pulumi.BoolPtrOutput)
}

func (o LookupViewByScopeResultOutput) Kpis() KpiPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) []KpiPropertiesResponse { return v.Kpis }).(KpiPropertiesResponseArrayOutput)
}

func (o LookupViewByScopeResultOutput) Metric() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *string { return v.Metric }).(pulumi.StringPtrOutput)
}

func (o LookupViewByScopeResultOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o LookupViewByScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupViewByScopeResultOutput) Pivots() PivotPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) []PivotPropertiesResponse { return v.Pivots }).(PivotPropertiesResponseArrayOutput)
}

func (o LookupViewByScopeResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o LookupViewByScopeResultOutput) TimePeriod() ReportConfigTimePeriodResponsePtrOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) *ReportConfigTimePeriodResponse { return v.TimePeriod }).(ReportConfigTimePeriodResponsePtrOutput)
}

func (o LookupViewByScopeResultOutput) Timeframe() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) string { return v.Timeframe }).(pulumi.StringOutput)
}

func (o LookupViewByScopeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewByScopeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupViewByScopeResultOutput{})
}
