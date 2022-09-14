


package mobilenetwork

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Slice struct {
	pulumi.CustomResourceState

	CreatedAt          pulumi.StringPtrOutput   `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrOutput   `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrOutput   `pulumi:"createdByType"`
	Description        pulumi.StringPtrOutput   `pulumi:"description"`
	LastModifiedAt     pulumi.StringPtrOutput   `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrOutput   `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrOutput   `pulumi:"lastModifiedByType"`
	Location           pulumi.StringOutput      `pulumi:"location"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput      `pulumi:"provisioningState"`
	Snssai             SnssaiResponseOutput     `pulumi:"snssai"`
	SystemData         SystemDataResponseOutput `pulumi:"systemData"`
	Tags               pulumi.StringMapOutput   `pulumi:"tags"`
	Type               pulumi.StringOutput      `pulumi:"type"`
}


func NewSlice(ctx *pulumi.Context,
	name string, args *SliceArgs, opts ...pulumi.ResourceOption) (*Slice, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MobileNetworkName == nil {
		return nil, errors.New("invalid value for required argument 'MobileNetworkName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Snssai == nil {
		return nil, errors.New("invalid value for required argument 'Snssai'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220301preview:Slice"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220401preview:Slice"),
		},
	})
	opts = append(opts, aliases)
	var resource Slice
	err := ctx.RegisterResource("azure-native:mobilenetwork:Slice", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSlice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SliceState, opts ...pulumi.ResourceOption) (*Slice, error) {
	var resource Slice
	err := ctx.ReadResource("azure-native:mobilenetwork:Slice", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sliceState struct {
}

type SliceState struct {
}

func (SliceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sliceState)(nil)).Elem()
}

type sliceArgs struct {
	CreatedAt          *string           `pulumi:"createdAt"`
	CreatedBy          *string           `pulumi:"createdBy"`
	CreatedByType      *string           `pulumi:"createdByType"`
	Description        *string           `pulumi:"description"`
	LastModifiedAt     *string           `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string           `pulumi:"lastModifiedBy"`
	LastModifiedByType *string           `pulumi:"lastModifiedByType"`
	Location           *string           `pulumi:"location"`
	MobileNetworkName  string            `pulumi:"mobileNetworkName"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	SliceName          *string           `pulumi:"sliceName"`
	Snssai             Snssai            `pulumi:"snssai"`
	Tags               map[string]string `pulumi:"tags"`
}


type SliceArgs struct {
	CreatedAt          pulumi.StringPtrInput
	CreatedBy          pulumi.StringPtrInput
	CreatedByType      pulumi.StringPtrInput
	Description        pulumi.StringPtrInput
	LastModifiedAt     pulumi.StringPtrInput
	LastModifiedBy     pulumi.StringPtrInput
	LastModifiedByType pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	MobileNetworkName  pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	SliceName          pulumi.StringPtrInput
	Snssai             SnssaiInput
	Tags               pulumi.StringMapInput
}

func (SliceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sliceArgs)(nil)).Elem()
}

type SliceInput interface {
	pulumi.Input

	ToSliceOutput() SliceOutput
	ToSliceOutputWithContext(ctx context.Context) SliceOutput
}

func (*Slice) ElementType() reflect.Type {
	return reflect.TypeOf((**Slice)(nil)).Elem()
}

func (i *Slice) ToSliceOutput() SliceOutput {
	return i.ToSliceOutputWithContext(context.Background())
}

func (i *Slice) ToSliceOutputWithContext(ctx context.Context) SliceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SliceOutput)
}

type SliceOutput struct{ *pulumi.OutputState }

func (SliceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Slice)(nil)).Elem()
}

func (o SliceOutput) ToSliceOutput() SliceOutput {
	return o
}

func (o SliceOutput) ToSliceOutputWithContext(ctx context.Context) SliceOutput {
	return o
}

func (o SliceOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SliceOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SliceOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringPtrOutput { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SliceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SliceOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringPtrOutput { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SliceOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SliceOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringPtrOutput { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o SliceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SliceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SliceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SliceOutput) Snssai() SnssaiResponseOutput {
	return o.ApplyT(func(v *Slice) SnssaiResponseOutput { return v.Snssai }).(SnssaiResponseOutput)
}

func (o SliceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Slice) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SliceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SliceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SliceOutput{})
}
