


package v20180617preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workbook struct {
	pulumi.CustomResourceState

	Category       pulumi.StringOutput    `pulumi:"category"`
	DisplayName    pulumi.StringOutput    `pulumi:"displayName"`
	Kind           pulumi.StringPtrOutput `pulumi:"kind"`
	Location       pulumi.StringOutput    `pulumi:"location"`
	Name           pulumi.StringOutput    `pulumi:"name"`
	SerializedData pulumi.StringOutput    `pulumi:"serializedData"`
	SourceId       pulumi.StringPtrOutput `pulumi:"sourceId"`
	Tags           pulumi.StringMapOutput `pulumi:"tags"`
	TimeModified   pulumi.StringOutput    `pulumi:"timeModified"`
	Type           pulumi.StringOutput    `pulumi:"type"`
	UserId         pulumi.StringOutput    `pulumi:"userId"`
	Version        pulumi.StringPtrOutput `pulumi:"version"`
}


func NewWorkbook(ctx *pulumi.Context,
	name string, args *WorkbookArgs, opts ...pulumi.ResourceOption) (*Workbook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SerializedData == nil {
		return nil, errors.New("invalid value for required argument 'SerializedData'")
	}
	if args.SourceId == nil {
		return nil, errors.New("invalid value for required argument 'SourceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:Workbook"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20150501:Workbook"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20201020:Workbook"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210308:Workbook"),
		},
	})
	opts = append(opts, aliases)
	var resource Workbook
	err := ctx.RegisterResource("azure-native:insights/v20180617preview:Workbook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkbook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkbookState, opts ...pulumi.ResourceOption) (*Workbook, error) {
	var resource Workbook
	err := ctx.ReadResource("azure-native:insights/v20180617preview:Workbook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workbookState struct {
}

type WorkbookState struct {
}

func (WorkbookState) ElementType() reflect.Type {
	return reflect.TypeOf((*workbookState)(nil)).Elem()
}

type workbookArgs struct {
	Category          string            `pulumi:"category"`
	DisplayName       string            `pulumi:"displayName"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      *string           `pulumi:"resourceName"`
	SerializedData    string            `pulumi:"serializedData"`
	SourceId          string            `pulumi:"sourceId"`
	Tags              map[string]string `pulumi:"tags"`
	Version           *string           `pulumi:"version"`
}


type WorkbookArgs struct {
	Category          pulumi.StringInput
	DisplayName       pulumi.StringInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	SerializedData    pulumi.StringInput
	SourceId          pulumi.StringInput
	Tags              pulumi.StringMapInput
	Version           pulumi.StringPtrInput
}

func (WorkbookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workbookArgs)(nil)).Elem()
}

type WorkbookInput interface {
	pulumi.Input

	ToWorkbookOutput() WorkbookOutput
	ToWorkbookOutputWithContext(ctx context.Context) WorkbookOutput
}

func (*Workbook) ElementType() reflect.Type {
	return reflect.TypeOf((*Workbook)(nil))
}

func (i *Workbook) ToWorkbookOutput() WorkbookOutput {
	return i.ToWorkbookOutputWithContext(context.Background())
}

func (i *Workbook) ToWorkbookOutputWithContext(ctx context.Context) WorkbookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookOutput)
}

type WorkbookOutput struct{ *pulumi.OutputState }

func (WorkbookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Workbook)(nil))
}

func (o WorkbookOutput) ToWorkbookOutput() WorkbookOutput {
	return o
}

func (o WorkbookOutput) ToWorkbookOutputWithContext(ctx context.Context) WorkbookOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkbookOutput{})
}
