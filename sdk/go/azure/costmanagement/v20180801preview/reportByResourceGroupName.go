


package v20180801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReportByResourceGroupName struct {
	pulumi.CustomResourceState

	Definition   ReportDefinitionResponseOutput   `pulumi:"definition"`
	DeliveryInfo ReportDeliveryInfoResponseOutput `pulumi:"deliveryInfo"`
	Format       pulumi.StringPtrOutput           `pulumi:"format"`
	Name         pulumi.StringOutput              `pulumi:"name"`
	Schedule     ReportScheduleResponsePtrOutput  `pulumi:"schedule"`
	Tags         pulumi.StringMapOutput           `pulumi:"tags"`
	Type         pulumi.StringOutput              `pulumi:"type"`
}


func NewReportByResourceGroupName(ctx *pulumi.Context,
	name string, args *ReportByResourceGroupNameArgs, opts ...pulumi.ResourceOption) (*ReportByResourceGroupName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	if args.DeliveryInfo == nil {
		return nil, errors.New("invalid value for required argument 'DeliveryInfo'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement:ReportByResourceGroupName"),
		},
	})
	opts = append(opts, aliases)
	var resource ReportByResourceGroupName
	err := ctx.RegisterResource("azure-native:costmanagement/v20180801preview:ReportByResourceGroupName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReportByResourceGroupName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReportByResourceGroupNameState, opts ...pulumi.ResourceOption) (*ReportByResourceGroupName, error) {
	var resource ReportByResourceGroupName
	err := ctx.ReadResource("azure-native:costmanagement/v20180801preview:ReportByResourceGroupName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type reportByResourceGroupNameState struct {
}

type ReportByResourceGroupNameState struct {
}

func (ReportByResourceGroupNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*reportByResourceGroupNameState)(nil)).Elem()
}

type reportByResourceGroupNameArgs struct {
	Definition        ReportDefinition   `pulumi:"definition"`
	DeliveryInfo      ReportDeliveryInfo `pulumi:"deliveryInfo"`
	Format            *string            `pulumi:"format"`
	ReportName        *string            `pulumi:"reportName"`
	ResourceGroupName string             `pulumi:"resourceGroupName"`
	Schedule          *ReportSchedule    `pulumi:"schedule"`
}


type ReportByResourceGroupNameArgs struct {
	Definition        ReportDefinitionInput
	DeliveryInfo      ReportDeliveryInfoInput
	Format            pulumi.StringPtrInput
	ReportName        pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Schedule          ReportSchedulePtrInput
}

func (ReportByResourceGroupNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reportByResourceGroupNameArgs)(nil)).Elem()
}

type ReportByResourceGroupNameInput interface {
	pulumi.Input

	ToReportByResourceGroupNameOutput() ReportByResourceGroupNameOutput
	ToReportByResourceGroupNameOutputWithContext(ctx context.Context) ReportByResourceGroupNameOutput
}

func (*ReportByResourceGroupName) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportByResourceGroupName)(nil)).Elem()
}

func (i *ReportByResourceGroupName) ToReportByResourceGroupNameOutput() ReportByResourceGroupNameOutput {
	return i.ToReportByResourceGroupNameOutputWithContext(context.Background())
}

func (i *ReportByResourceGroupName) ToReportByResourceGroupNameOutputWithContext(ctx context.Context) ReportByResourceGroupNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportByResourceGroupNameOutput)
}

type ReportByResourceGroupNameOutput struct{ *pulumi.OutputState }

func (ReportByResourceGroupNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportByResourceGroupName)(nil)).Elem()
}

func (o ReportByResourceGroupNameOutput) ToReportByResourceGroupNameOutput() ReportByResourceGroupNameOutput {
	return o
}

func (o ReportByResourceGroupNameOutput) ToReportByResourceGroupNameOutputWithContext(ctx context.Context) ReportByResourceGroupNameOutput {
	return o
}

func (o ReportByResourceGroupNameOutput) Definition() ReportDefinitionResponseOutput {
	return o.ApplyT(func(v *ReportByResourceGroupName) ReportDefinitionResponseOutput { return v.Definition }).(ReportDefinitionResponseOutput)
}

func (o ReportByResourceGroupNameOutput) DeliveryInfo() ReportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v *ReportByResourceGroupName) ReportDeliveryInfoResponseOutput { return v.DeliveryInfo }).(ReportDeliveryInfoResponseOutput)
}

func (o ReportByResourceGroupNameOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportByResourceGroupName) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

func (o ReportByResourceGroupNameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportByResourceGroupName) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ReportByResourceGroupNameOutput) Schedule() ReportScheduleResponsePtrOutput {
	return o.ApplyT(func(v *ReportByResourceGroupName) ReportScheduleResponsePtrOutput { return v.Schedule }).(ReportScheduleResponsePtrOutput)
}

func (o ReportByResourceGroupNameOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ReportByResourceGroupName) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ReportByResourceGroupNameOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportByResourceGroupName) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReportByResourceGroupNameOutput{})
}
