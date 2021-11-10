


package v20180531

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReportConfigByResourceGroupName struct {
	pulumi.CustomResourceState

	Definition   ReportConfigDefinitionResponseOutput   `pulumi:"definition"`
	DeliveryInfo ReportConfigDeliveryInfoResponseOutput `pulumi:"deliveryInfo"`
	Format       pulumi.StringPtrOutput                 `pulumi:"format"`
	Name         pulumi.StringOutput                    `pulumi:"name"`
	Schedule     ReportConfigScheduleResponsePtrOutput  `pulumi:"schedule"`
	Tags         pulumi.StringMapOutput                 `pulumi:"tags"`
	Type         pulumi.StringOutput                    `pulumi:"type"`
}


func NewReportConfigByResourceGroupName(ctx *pulumi.Context,
	name string, args *ReportConfigByResourceGroupNameArgs, opts ...pulumi.ResourceOption) (*ReportConfigByResourceGroupName, error) {
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
	var resource ReportConfigByResourceGroupName
	err := ctx.RegisterResource("azure-native:costmanagement/v20180531:ReportConfigByResourceGroupName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReportConfigByResourceGroupName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReportConfigByResourceGroupNameState, opts ...pulumi.ResourceOption) (*ReportConfigByResourceGroupName, error) {
	var resource ReportConfigByResourceGroupName
	err := ctx.ReadResource("azure-native:costmanagement/v20180531:ReportConfigByResourceGroupName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type reportConfigByResourceGroupNameState struct {
}

type ReportConfigByResourceGroupNameState struct {
}

func (ReportConfigByResourceGroupNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*reportConfigByResourceGroupNameState)(nil)).Elem()
}

type reportConfigByResourceGroupNameArgs struct {
	Definition        ReportConfigDefinition   `pulumi:"definition"`
	DeliveryInfo      ReportConfigDeliveryInfo `pulumi:"deliveryInfo"`
	Format            *string                  `pulumi:"format"`
	ReportConfigName  *string                  `pulumi:"reportConfigName"`
	ResourceGroupName string                   `pulumi:"resourceGroupName"`
	Schedule          *ReportConfigSchedule    `pulumi:"schedule"`
}


type ReportConfigByResourceGroupNameArgs struct {
	Definition        ReportConfigDefinitionInput
	DeliveryInfo      ReportConfigDeliveryInfoInput
	Format            pulumi.StringPtrInput
	ReportConfigName  pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Schedule          ReportConfigSchedulePtrInput
}

func (ReportConfigByResourceGroupNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reportConfigByResourceGroupNameArgs)(nil)).Elem()
}

type ReportConfigByResourceGroupNameInput interface {
	pulumi.Input

	ToReportConfigByResourceGroupNameOutput() ReportConfigByResourceGroupNameOutput
	ToReportConfigByResourceGroupNameOutputWithContext(ctx context.Context) ReportConfigByResourceGroupNameOutput
}

func (*ReportConfigByResourceGroupName) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigByResourceGroupName)(nil))
}

func (i *ReportConfigByResourceGroupName) ToReportConfigByResourceGroupNameOutput() ReportConfigByResourceGroupNameOutput {
	return i.ToReportConfigByResourceGroupNameOutputWithContext(context.Background())
}

func (i *ReportConfigByResourceGroupName) ToReportConfigByResourceGroupNameOutputWithContext(ctx context.Context) ReportConfigByResourceGroupNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigByResourceGroupNameOutput)
}

type ReportConfigByResourceGroupNameOutput struct{ *pulumi.OutputState }

func (ReportConfigByResourceGroupNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigByResourceGroupName)(nil))
}

func (o ReportConfigByResourceGroupNameOutput) ToReportConfigByResourceGroupNameOutput() ReportConfigByResourceGroupNameOutput {
	return o
}

func (o ReportConfigByResourceGroupNameOutput) ToReportConfigByResourceGroupNameOutputWithContext(ctx context.Context) ReportConfigByResourceGroupNameOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ReportConfigByResourceGroupNameOutput{})
}
