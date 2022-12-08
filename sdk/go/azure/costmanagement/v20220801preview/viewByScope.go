


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ViewByScope struct {
	pulumi.CustomResourceState

	Accumulated               pulumi.StringPtrOutput                  `pulumi:"accumulated"`
	Chart                     pulumi.StringPtrOutput                  `pulumi:"chart"`
	CreatedOn                 pulumi.StringOutput                     `pulumi:"createdOn"`
	Currency                  pulumi.StringOutput                     `pulumi:"currency"`
	DataSet                   ReportConfigDatasetResponsePtrOutput    `pulumi:"dataSet"`
	DateRange                 pulumi.StringOutput                     `pulumi:"dateRange"`
	DisplayName               pulumi.StringPtrOutput                  `pulumi:"displayName"`
	ETag                      pulumi.StringPtrOutput                  `pulumi:"eTag"`
	IncludeMonetaryCommitment pulumi.BoolPtrOutput                    `pulumi:"includeMonetaryCommitment"`
	Kpis                      KpiPropertiesResponseArrayOutput        `pulumi:"kpis"`
	Metric                    pulumi.StringPtrOutput                  `pulumi:"metric"`
	ModifiedOn                pulumi.StringOutput                     `pulumi:"modifiedOn"`
	Name                      pulumi.StringOutput                     `pulumi:"name"`
	Pivots                    PivotPropertiesResponseArrayOutput      `pulumi:"pivots"`
	Scope                     pulumi.StringPtrOutput                  `pulumi:"scope"`
	TimePeriod                ReportConfigTimePeriodResponsePtrOutput `pulumi:"timePeriod"`
	Timeframe                 pulumi.StringOutput                     `pulumi:"timeframe"`
	Type                      pulumi.StringOutput                     `pulumi:"type"`
}


func NewViewByScope(ctx *pulumi.Context,
	name string, args *ViewByScopeArgs, opts ...pulumi.ResourceOption) (*ViewByScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.Timeframe == nil {
		return nil, errors.New("invalid value for required argument 'Timeframe'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20190401preview:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20191101:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20200601:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20211001:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20221001:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20221001preview:ViewByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20221005preview:ViewByScope"),
		},
	})
	opts = append(opts, aliases)
	var resource ViewByScope
	err := ctx.RegisterResource("azure-native:costmanagement/v20220801preview:ViewByScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetViewByScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ViewByScopeState, opts ...pulumi.ResourceOption) (*ViewByScope, error) {
	var resource ViewByScope
	err := ctx.ReadResource("azure-native:costmanagement/v20220801preview:ViewByScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type viewByScopeState struct {
}

type ViewByScopeState struct {
}

func (ViewByScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*viewByScopeState)(nil)).Elem()
}

type viewByScopeArgs struct {
	Accumulated               *string                 `pulumi:"accumulated"`
	Chart                     *string                 `pulumi:"chart"`
	DataSet                   *ReportConfigDataset    `pulumi:"dataSet"`
	DisplayName               *string                 `pulumi:"displayName"`
	ETag                      *string                 `pulumi:"eTag"`
	IncludeMonetaryCommitment *bool                   `pulumi:"includeMonetaryCommitment"`
	Kpis                      []KpiProperties         `pulumi:"kpis"`
	Metric                    *string                 `pulumi:"metric"`
	Pivots                    []PivotProperties       `pulumi:"pivots"`
	Scope                     string                  `pulumi:"scope"`
	TimePeriod                *ReportConfigTimePeriod `pulumi:"timePeriod"`
	Timeframe                 string                  `pulumi:"timeframe"`
	Type                      string                  `pulumi:"type"`
	ViewName                  *string                 `pulumi:"viewName"`
}


type ViewByScopeArgs struct {
	Accumulated               pulumi.StringPtrInput
	Chart                     pulumi.StringPtrInput
	DataSet                   ReportConfigDatasetPtrInput
	DisplayName               pulumi.StringPtrInput
	ETag                      pulumi.StringPtrInput
	IncludeMonetaryCommitment pulumi.BoolPtrInput
	Kpis                      KpiPropertiesArrayInput
	Metric                    pulumi.StringPtrInput
	Pivots                    PivotPropertiesArrayInput
	Scope                     pulumi.StringInput
	TimePeriod                ReportConfigTimePeriodPtrInput
	Timeframe                 pulumi.StringInput
	Type                      pulumi.StringInput
	ViewName                  pulumi.StringPtrInput
}

func (ViewByScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*viewByScopeArgs)(nil)).Elem()
}

type ViewByScopeInput interface {
	pulumi.Input

	ToViewByScopeOutput() ViewByScopeOutput
	ToViewByScopeOutputWithContext(ctx context.Context) ViewByScopeOutput
}

func (*ViewByScope) ElementType() reflect.Type {
	return reflect.TypeOf((**ViewByScope)(nil)).Elem()
}

func (i *ViewByScope) ToViewByScopeOutput() ViewByScopeOutput {
	return i.ToViewByScopeOutputWithContext(context.Background())
}

func (i *ViewByScope) ToViewByScopeOutputWithContext(ctx context.Context) ViewByScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewByScopeOutput)
}

type ViewByScopeOutput struct{ *pulumi.OutputState }

func (ViewByScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ViewByScope)(nil)).Elem()
}

func (o ViewByScopeOutput) ToViewByScopeOutput() ViewByScopeOutput {
	return o
}

func (o ViewByScopeOutput) ToViewByScopeOutputWithContext(ctx context.Context) ViewByScopeOutput {
	return o
}

func (o ViewByScopeOutput) Accumulated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringPtrOutput { return v.Accumulated }).(pulumi.StringPtrOutput)
}

func (o ViewByScopeOutput) Chart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringPtrOutput { return v.Chart }).(pulumi.StringPtrOutput)
}

func (o ViewByScopeOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o ViewByScopeOutput) Currency() pulumi.StringOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringOutput { return v.Currency }).(pulumi.StringOutput)
}

func (o ViewByScopeOutput) DataSet() ReportConfigDatasetResponsePtrOutput {
	return o.ApplyT(func(v *ViewByScope) ReportConfigDatasetResponsePtrOutput { return v.DataSet }).(ReportConfigDatasetResponsePtrOutput)
}

func (o ViewByScopeOutput) DateRange() pulumi.StringOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringOutput { return v.DateRange }).(pulumi.StringOutput)
}

func (o ViewByScopeOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ViewByScopeOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o ViewByScopeOutput) IncludeMonetaryCommitment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.BoolPtrOutput { return v.IncludeMonetaryCommitment }).(pulumi.BoolPtrOutput)
}

func (o ViewByScopeOutput) Kpis() KpiPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *ViewByScope) KpiPropertiesResponseArrayOutput { return v.Kpis }).(KpiPropertiesResponseArrayOutput)
}

func (o ViewByScopeOutput) Metric() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringPtrOutput { return v.Metric }).(pulumi.StringPtrOutput)
}

func (o ViewByScopeOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringOutput { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o ViewByScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ViewByScopeOutput) Pivots() PivotPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *ViewByScope) PivotPropertiesResponseArrayOutput { return v.Pivots }).(PivotPropertiesResponseArrayOutput)
}

func (o ViewByScopeOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o ViewByScopeOutput) TimePeriod() ReportConfigTimePeriodResponsePtrOutput {
	return o.ApplyT(func(v *ViewByScope) ReportConfigTimePeriodResponsePtrOutput { return v.TimePeriod }).(ReportConfigTimePeriodResponsePtrOutput)
}

func (o ViewByScopeOutput) Timeframe() pulumi.StringOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringOutput { return v.Timeframe }).(pulumi.StringOutput)
}

func (o ViewByScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ViewByScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ViewByScopeOutput{})
}
