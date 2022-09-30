


package v20220401

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
	Etag           pulumi.StringPtrOutput                    `pulumi:"etag"`
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
		{
			Type: pulumi.String("azure-native:insights/v20210801:Workbook"),
		},
	})
	opts = append(opts, aliases)
	var resource Workbook
	err := ctx.RegisterResource("azure-native:insights/v20220401:Workbook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkbook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkbookState, opts ...pulumi.ResourceOption) (*Workbook, error) {
	var resource Workbook
	err := ctx.ReadResource("azure-native:insights/v20220401:Workbook", name, id, state, &resource, opts...)
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

func (o WorkbookOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

func (o WorkbookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WorkbookOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o WorkbookOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o WorkbookOutput) Identity() WorkbookResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v *Workbook) WorkbookResourceResponseIdentityPtrOutput { return v.Identity }).(WorkbookResourceResponseIdentityPtrOutput)
}

func (o WorkbookOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WorkbookOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o WorkbookOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkbookOutput) Revision() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.Revision }).(pulumi.StringOutput)
}

func (o WorkbookOutput) SerializedData() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.SerializedData }).(pulumi.StringOutput)
}

func (o WorkbookOutput) SourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringPtrOutput { return v.SourceId }).(pulumi.StringPtrOutput)
}

func (o WorkbookOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringPtrOutput { return v.StorageUri }).(pulumi.StringPtrOutput)
}

func (o WorkbookOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Workbook) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
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

func init() {
	pulumi.RegisterOutputType(WorkbookOutput{})
}
