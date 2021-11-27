


package v20190101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Export struct {
	pulumi.CustomResourceState

	Definition   QueryDefinitionResponseOutput    `pulumi:"definition"`
	DeliveryInfo ExportDeliveryInfoResponseOutput `pulumi:"deliveryInfo"`
	Format       pulumi.StringPtrOutput           `pulumi:"format"`
	Name         pulumi.StringOutput              `pulumi:"name"`
	Schedule     ExportScheduleResponsePtrOutput  `pulumi:"schedule"`
	Tags         pulumi.StringMapOutput           `pulumi:"tags"`
	Type         pulumi.StringOutput              `pulumi:"type"`
}


func NewExport(ctx *pulumi.Context,
	name string, args *ExportArgs, opts ...pulumi.ResourceOption) (*Export, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	if args.DeliveryInfo == nil {
		return nil, errors.New("invalid value for required argument 'DeliveryInfo'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement:Export"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20190901:Export"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20191001:Export"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20191101:Export"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20200601:Export"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20201201preview:Export"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20210101:Export"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20211001:Export"),
		},
	})
	opts = append(opts, aliases)
	var resource Export
	err := ctx.RegisterResource("azure-native:costmanagement/v20190101:Export", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExportState, opts ...pulumi.ResourceOption) (*Export, error) {
	var resource Export
	err := ctx.ReadResource("azure-native:costmanagement/v20190101:Export", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type exportState struct {
}

type ExportState struct {
}

func (ExportState) ElementType() reflect.Type {
	return reflect.TypeOf((*exportState)(nil)).Elem()
}

type exportArgs struct {
	Definition   QueryDefinition    `pulumi:"definition"`
	DeliveryInfo ExportDeliveryInfo `pulumi:"deliveryInfo"`
	ExportName   *string            `pulumi:"exportName"`
	Format       *string            `pulumi:"format"`
	Schedule     *ExportSchedule    `pulumi:"schedule"`
	Scope        string             `pulumi:"scope"`
}


type ExportArgs struct {
	Definition   QueryDefinitionInput
	DeliveryInfo ExportDeliveryInfoInput
	ExportName   pulumi.StringPtrInput
	Format       pulumi.StringPtrInput
	Schedule     ExportSchedulePtrInput
	Scope        pulumi.StringInput
}

func (ExportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*exportArgs)(nil)).Elem()
}

type ExportInput interface {
	pulumi.Input

	ToExportOutput() ExportOutput
	ToExportOutputWithContext(ctx context.Context) ExportOutput
}

func (*Export) ElementType() reflect.Type {
	return reflect.TypeOf((*Export)(nil))
}

func (i *Export) ToExportOutput() ExportOutput {
	return i.ToExportOutputWithContext(context.Background())
}

func (i *Export) ToExportOutputWithContext(ctx context.Context) ExportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportOutput)
}

type ExportOutput struct{ *pulumi.OutputState }

func (ExportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Export)(nil))
}

func (o ExportOutput) ToExportOutput() ExportOutput {
	return o
}

func (o ExportOutput) ToExportOutputWithContext(ctx context.Context) ExportOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ExportOutput{})
}
