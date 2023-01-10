


package v20221116preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Report struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput            `pulumi:"name"`
	Properties ReportPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput       `pulumi:"systemData"`
	Type       pulumi.StringOutput            `pulumi:"type"`
}


func NewReport(ctx *pulumi.Context,
	name string, args *ReportArgs, opts ...pulumi.ResourceOption) (*Report, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appcomplianceautomation:Report"),
		},
	})
	opts = append(opts, aliases)
	var resource Report
	err := ctx.RegisterResource("azure-native:appcomplianceautomation/v20221116preview:Report", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReportState, opts ...pulumi.ResourceOption) (*Report, error) {
	var resource Report
	err := ctx.ReadResource("azure-native:appcomplianceautomation/v20221116preview:Report", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type reportState struct {
}

type ReportState struct {
}

func (ReportState) ElementType() reflect.Type {
	return reflect.TypeOf((*reportState)(nil)).Elem()
}

type reportArgs struct {
	Properties ReportProperties `pulumi:"properties"`
	ReportName *string          `pulumi:"reportName"`
}


type ReportArgs struct {
	Properties ReportPropertiesInput
	ReportName pulumi.StringPtrInput
}

func (ReportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reportArgs)(nil)).Elem()
}

type ReportInput interface {
	pulumi.Input

	ToReportOutput() ReportOutput
	ToReportOutputWithContext(ctx context.Context) ReportOutput
}

func (*Report) ElementType() reflect.Type {
	return reflect.TypeOf((**Report)(nil)).Elem()
}

func (i *Report) ToReportOutput() ReportOutput {
	return i.ToReportOutputWithContext(context.Background())
}

func (i *Report) ToReportOutputWithContext(ctx context.Context) ReportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportOutput)
}

type ReportOutput struct{ *pulumi.OutputState }

func (ReportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Report)(nil)).Elem()
}

func (o ReportOutput) ToReportOutput() ReportOutput {
	return o
}

func (o ReportOutput) ToReportOutputWithContext(ctx context.Context) ReportOutput {
	return o
}

func (o ReportOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Report) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ReportOutput) Properties() ReportPropertiesResponseOutput {
	return o.ApplyT(func(v *Report) ReportPropertiesResponseOutput { return v.Properties }).(ReportPropertiesResponseOutput)
}

func (o ReportOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Report) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ReportOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Report) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReportOutput{})
}
