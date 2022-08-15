


package v20150501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workbook struct {
	pulumi.CustomResourceState

	Category         pulumi.StringOutput    `pulumi:"category"`
	Kind             pulumi.StringPtrOutput `pulumi:"kind"`
	Location         pulumi.StringPtrOutput `pulumi:"location"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	SerializedData   pulumi.StringOutput    `pulumi:"serializedData"`
	SharedTypeKind   pulumi.StringOutput    `pulumi:"sharedTypeKind"`
	SourceResourceId pulumi.StringPtrOutput `pulumi:"sourceResourceId"`
	Tags             pulumi.StringMapOutput `pulumi:"tags"`
	TimeModified     pulumi.StringOutput    `pulumi:"timeModified"`
	Type             pulumi.StringOutput    `pulumi:"type"`
	UserId           pulumi.StringOutput    `pulumi:"userId"`
	Version          pulumi.StringPtrOutput `pulumi:"version"`
	WorkbookId       pulumi.StringOutput    `pulumi:"workbookId"`
}


func NewWorkbook(ctx *pulumi.Context,
	name string, args *WorkbookArgs, opts ...pulumi.ResourceOption) (*Workbook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SerializedData == nil {
		return nil, errors.New("invalid value for required argument 'SerializedData'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	if args.WorkbookId == nil {
		return nil, errors.New("invalid value for required argument 'WorkbookId'")
	}
	if isZero(args.SharedTypeKind) {
		args.SharedTypeKind = pulumi.String("shared")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:Workbook"),
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
		{
			Type: pulumi.String("azure-native:insights/v20210801:Workbook"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20220401:Workbook"),
		},
	})
	opts = append(opts, aliases)
	var resource Workbook
	err := ctx.RegisterResource("azure-native:insights/v20150501:Workbook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkbook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkbookState, opts ...pulumi.ResourceOption) (*Workbook, error) {
	var resource Workbook
	err := ctx.ReadResource("azure-native:insights/v20150501:Workbook", name, id, state, &resource, opts...)
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
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      *string           `pulumi:"resourceName"`
	SerializedData    string            `pulumi:"serializedData"`
	SharedTypeKind    string            `pulumi:"sharedTypeKind"`
	SourceResourceId  *string           `pulumi:"sourceResourceId"`
	Tags              map[string]string `pulumi:"tags"`
	UserId            string            `pulumi:"userId"`
	Version           *string           `pulumi:"version"`
	WorkbookId        string            `pulumi:"workbookId"`
}


type WorkbookArgs struct {
	Category          pulumi.StringInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	SerializedData    pulumi.StringInput
	SharedTypeKind    pulumi.StringInput
	SourceResourceId  pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	UserId            pulumi.StringInput
	Version           pulumi.StringPtrInput
	WorkbookId        pulumi.StringInput
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

func (o WorkbookOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

func (o WorkbookOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WorkbookOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o WorkbookOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkbookOutput) SerializedData() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.SerializedData }).(pulumi.StringOutput)
}

func (o WorkbookOutput) SharedTypeKind() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.SharedTypeKind }).(pulumi.StringOutput)
}

func (o WorkbookOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringPtrOutput { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o WorkbookOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o WorkbookOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.TimeModified }).(pulumi.StringOutput)
}

func (o WorkbookOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o WorkbookOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

func (o WorkbookOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func (o WorkbookOutput) WorkbookId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.WorkbookId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkbookOutput{})
}
