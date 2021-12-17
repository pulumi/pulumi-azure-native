


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workbook struct {
	pulumi.CustomResourceState

	Category       pulumi.StringOutput                       `pulumi:"category"`
	Description    pulumi.StringPtrOutput                    `pulumi:"description"`
	DisplayName    pulumi.StringOutput                       `pulumi:"displayName"`
	Etag           pulumi.StringMapOutput                    `pulumi:"etag"`
	Identity       WorkbookResourceResponseIdentityPtrOutput `pulumi:"identity"`
	Kind           pulumi.StringPtrOutput                    `pulumi:"kind"`
	Location       pulumi.StringOutput                       `pulumi:"location"`
	Name           pulumi.StringOutput                       `pulumi:"name"`
	Revision       pulumi.StringOutput                       `pulumi:"revision"`
	SerializedData pulumi.StringOutput                       `pulumi:"serializedData"`
	SourceId       pulumi.StringPtrOutput                    `pulumi:"sourceId"`
	StorageUri     pulumi.StringPtrOutput                    `pulumi:"storageUri"`
	SystemData     SystemDataResponseOutput                  `pulumi:"systemData"`
	Tags           pulumi.StringMapOutput                    `pulumi:"tags"`
	TimeModified   pulumi.StringOutput                       `pulumi:"timeModified"`
	Type           pulumi.StringOutput                       `pulumi:"type"`
	UserId         pulumi.StringOutput                       `pulumi:"userId"`
	Version        pulumi.StringPtrOutput                    `pulumi:"version"`
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:Workbook"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20150501:Workbook"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20180617preview:Workbook"),
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
	err := ctx.RegisterResource("azure-native:insights/v20210801:Workbook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkbook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkbookState, opts ...pulumi.ResourceOption) (*Workbook, error) {
	var resource Workbook
	err := ctx.ReadResource("azure-native:insights/v20210801:Workbook", name, id, state, &resource, opts...)
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
	Category          string                    `pulumi:"category"`
	Description       *string                   `pulumi:"description"`
	DisplayName       string                    `pulumi:"displayName"`
	Etag              map[string]string         `pulumi:"etag"`
	Identity          *WorkbookResourceIdentity `pulumi:"identity"`
	Kind              *string                   `pulumi:"kind"`
	Location          *string                   `pulumi:"location"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	ResourceName      *string                   `pulumi:"resourceName"`
	SerializedData    string                    `pulumi:"serializedData"`
	SourceId          *string                   `pulumi:"sourceId"`
	StorageUri        *string                   `pulumi:"storageUri"`
	Tags              map[string]string         `pulumi:"tags"`
	Version           *string                   `pulumi:"version"`
}


type WorkbookArgs struct {
	Category          pulumi.StringInput
	Description       pulumi.StringPtrInput
	DisplayName       pulumi.StringInput
	Etag              pulumi.StringMapInput
	Identity          WorkbookResourceIdentityPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	SerializedData    pulumi.StringInput
	SourceId          pulumi.StringPtrInput
	StorageUri        pulumi.StringPtrInput
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
	return reflect.TypeOf((**Workbook)(nil)).Elem()
}

func (i *Workbook) ToWorkbookOutput() WorkbookOutput {
	return i.ToWorkbookOutputWithContext(context.Background())
}

func (i *Workbook) ToWorkbookOutputWithContext(ctx context.Context) WorkbookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookOutput)
}

type WorkbookOutput struct{ *pulumi.OutputState }

func (WorkbookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workbook)(nil)).Elem()
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
