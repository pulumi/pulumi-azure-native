


package v20221005preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type View struct {
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


func NewView(ctx *pulumi.Context,
	name string, args *ViewArgs, opts ...pulumi.ResourceOption) (*View, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Timeframe == nil {
		return nil, errors.New("invalid value for required argument 'Timeframe'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement:View"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20190401preview:View"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20191101:View"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20200601:View"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20211001:View"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20220801preview:View"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20221001preview:View"),
		},
	})
	opts = append(opts, aliases)
	var resource View
	err := ctx.RegisterResource("azure-native:costmanagement/v20221005preview:View", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetView(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ViewState, opts ...pulumi.ResourceOption) (*View, error) {
	var resource View
	err := ctx.ReadResource("azure-native:costmanagement/v20221005preview:View", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type viewState struct {
}

type ViewState struct {
}

func (ViewState) ElementType() reflect.Type {
	return reflect.TypeOf((*viewState)(nil)).Elem()
}

type viewArgs struct {
	Accumulated               *string                 `pulumi:"accumulated"`
	Chart                     *string                 `pulumi:"chart"`
	DataSet                   *ReportConfigDataset    `pulumi:"dataSet"`
	DisplayName               *string                 `pulumi:"displayName"`
	ETag                      *string                 `pulumi:"eTag"`
	IncludeMonetaryCommitment *bool                   `pulumi:"includeMonetaryCommitment"`
	Kpis                      []KpiProperties         `pulumi:"kpis"`
	Metric                    *string                 `pulumi:"metric"`
	Pivots                    []PivotProperties       `pulumi:"pivots"`
	Scope                     *string                 `pulumi:"scope"`
	TimePeriod                *ReportConfigTimePeriod `pulumi:"timePeriod"`
	Timeframe                 string                  `pulumi:"timeframe"`
	Type                      string                  `pulumi:"type"`
	ViewName                  *string                 `pulumi:"viewName"`
}


type ViewArgs struct {
	Accumulated               pulumi.StringPtrInput
	Chart                     pulumi.StringPtrInput
	DataSet                   ReportConfigDatasetPtrInput
	DisplayName               pulumi.StringPtrInput
	ETag                      pulumi.StringPtrInput
	IncludeMonetaryCommitment pulumi.BoolPtrInput
	Kpis                      KpiPropertiesArrayInput
	Metric                    pulumi.StringPtrInput
	Pivots                    PivotPropertiesArrayInput
	Scope                     pulumi.StringPtrInput
	TimePeriod                ReportConfigTimePeriodPtrInput
	Timeframe                 pulumi.StringInput
	Type                      pulumi.StringInput
	ViewName                  pulumi.StringPtrInput
}

func (ViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*viewArgs)(nil)).Elem()
}

type ViewInput interface {
	pulumi.Input

	ToViewOutput() ViewOutput
	ToViewOutputWithContext(ctx context.Context) ViewOutput
}

func (*View) ElementType() reflect.Type {
	return reflect.TypeOf((**View)(nil)).Elem()
}

func (i *View) ToViewOutput() ViewOutput {
	return i.ToViewOutputWithContext(context.Background())
}

func (i *View) ToViewOutputWithContext(ctx context.Context) ViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewOutput)
}

type ViewOutput struct{ *pulumi.OutputState }

func (ViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**View)(nil)).Elem()
}

func (o ViewOutput) ToViewOutput() ViewOutput {
	return o
}

func (o ViewOutput) ToViewOutputWithContext(ctx context.Context) ViewOutput {
	return o
}

func (o ViewOutput) Accumulated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *View) pulumi.StringPtrOutput { return v.Accumulated }).(pulumi.StringPtrOutput)
}

func (o ViewOutput) Chart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *View) pulumi.StringPtrOutput { return v.Chart }).(pulumi.StringPtrOutput)
}

func (o ViewOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o ViewOutput) Currency() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.Currency }).(pulumi.StringOutput)
}

func (o ViewOutput) DataSet() ReportConfigDatasetResponsePtrOutput {
	return o.ApplyT(func(v *View) ReportConfigDatasetResponsePtrOutput { return v.DataSet }).(ReportConfigDatasetResponsePtrOutput)
}

func (o ViewOutput) DateRange() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.DateRange }).(pulumi.StringOutput)
}

func (o ViewOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *View) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ViewOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *View) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o ViewOutput) IncludeMonetaryCommitment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *View) pulumi.BoolPtrOutput { return v.IncludeMonetaryCommitment }).(pulumi.BoolPtrOutput)
}

func (o ViewOutput) Kpis() KpiPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *View) KpiPropertiesResponseArrayOutput { return v.Kpis }).(KpiPropertiesResponseArrayOutput)
}

func (o ViewOutput) Metric() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *View) pulumi.StringPtrOutput { return v.Metric }).(pulumi.StringPtrOutput)
}

func (o ViewOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.ModifiedOn }).(pulumi.StringOutput)
}

func (o ViewOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ViewOutput) Pivots() PivotPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *View) PivotPropertiesResponseArrayOutput { return v.Pivots }).(PivotPropertiesResponseArrayOutput)
}

func (o ViewOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *View) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o ViewOutput) TimePeriod() ReportConfigTimePeriodResponsePtrOutput {
	return o.ApplyT(func(v *View) ReportConfigTimePeriodResponsePtrOutput { return v.TimePeriod }).(ReportConfigTimePeriodResponsePtrOutput)
}

func (o ViewOutput) Timeframe() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.Timeframe }).(pulumi.StringOutput)
}

func (o ViewOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ViewOutput{})
}
