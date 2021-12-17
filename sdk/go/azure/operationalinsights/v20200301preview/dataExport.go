


package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataExport struct {
	pulumi.CustomResourceState

	CreatedDate      pulumi.StringPtrOutput   `pulumi:"createdDate"`
	DataExportId     pulumi.StringPtrOutput   `pulumi:"dataExportId"`
	Enable           pulumi.BoolPtrOutput     `pulumi:"enable"`
	EventHubName     pulumi.StringPtrOutput   `pulumi:"eventHubName"`
	LastModifiedDate pulumi.StringPtrOutput   `pulumi:"lastModifiedDate"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	ResourceId       pulumi.StringOutput      `pulumi:"resourceId"`
	TableNames       pulumi.StringArrayOutput `pulumi:"tableNames"`
	Type             pulumi.StringOutput      `pulumi:"type"`
}


func NewDataExport(ctx *pulumi.Context,
	name string, args *DataExportArgs, opts ...pulumi.ResourceOption) (*DataExport, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.TableNames == nil {
		return nil, errors.New("invalid value for required argument 'TableNames'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights:DataExport"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20190801preview:DataExport"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200801:DataExport"),
		},
	})
	opts = append(opts, aliases)
	var resource DataExport
	err := ctx.RegisterResource("azure-native:operationalinsights/v20200301preview:DataExport", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataExport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataExportState, opts ...pulumi.ResourceOption) (*DataExport, error) {
	var resource DataExport
	err := ctx.ReadResource("azure-native:operationalinsights/v20200301preview:DataExport", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dataExportState struct {
}

type DataExportState struct {
}

func (DataExportState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataExportState)(nil)).Elem()
}

type dataExportArgs struct {
	CreatedDate       *string  `pulumi:"createdDate"`
	DataExportId      *string  `pulumi:"dataExportId"`
	DataExportName    *string  `pulumi:"dataExportName"`
	Enable            *bool    `pulumi:"enable"`
	EventHubName      *string  `pulumi:"eventHubName"`
	LastModifiedDate  *string  `pulumi:"lastModifiedDate"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	ResourceId        string   `pulumi:"resourceId"`
	TableNames        []string `pulumi:"tableNames"`
	WorkspaceName     string   `pulumi:"workspaceName"`
}


type DataExportArgs struct {
	CreatedDate       pulumi.StringPtrInput
	DataExportId      pulumi.StringPtrInput
	DataExportName    pulumi.StringPtrInput
	Enable            pulumi.BoolPtrInput
	EventHubName      pulumi.StringPtrInput
	LastModifiedDate  pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceId        pulumi.StringInput
	TableNames        pulumi.StringArrayInput
	WorkspaceName     pulumi.StringInput
}

func (DataExportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataExportArgs)(nil)).Elem()
}

type DataExportInput interface {
	pulumi.Input

	ToDataExportOutput() DataExportOutput
	ToDataExportOutputWithContext(ctx context.Context) DataExportOutput
}

func (*DataExport) ElementType() reflect.Type {
	return reflect.TypeOf((**DataExport)(nil)).Elem()
}

func (i *DataExport) ToDataExportOutput() DataExportOutput {
	return i.ToDataExportOutputWithContext(context.Background())
}

func (i *DataExport) ToDataExportOutputWithContext(ctx context.Context) DataExportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataExportOutput)
}

type DataExportOutput struct{ *pulumi.OutputState }

func (DataExportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataExport)(nil)).Elem()
}

func (o DataExportOutput) ToDataExportOutput() DataExportOutput {
	return o
}

func (o DataExportOutput) ToDataExportOutputWithContext(ctx context.Context) DataExportOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DataExportOutput{})
}
