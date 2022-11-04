


package v20200601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IpAllocation struct {
	pulumi.CustomResourceState

	AllocationTags   pulumi.StringMapOutput    `pulumi:"allocationTags"`
	Etag             pulumi.StringOutput       `pulumi:"etag"`
	IpamAllocationId pulumi.StringPtrOutput    `pulumi:"ipamAllocationId"`
	Location         pulumi.StringPtrOutput    `pulumi:"location"`
	Name             pulumi.StringOutput       `pulumi:"name"`
	Prefix           pulumi.StringPtrOutput    `pulumi:"prefix"`
	PrefixLength     pulumi.IntPtrOutput       `pulumi:"prefixLength"`
	PrefixType       pulumi.StringPtrOutput    `pulumi:"prefixType"`
	Subnet           SubResourceResponseOutput `pulumi:"subnet"`
	Tags             pulumi.StringMapOutput    `pulumi:"tags"`
	Type             pulumi.StringOutput       `pulumi:"type"`
	VirtualNetwork   SubResourceResponseOutput `pulumi:"virtualNetwork"`
}


func NewIpAllocation(ctx *pulumi.Context,
	name string, args *IpAllocationArgs, opts ...pulumi.ResourceOption) (*IpAllocation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.PrefixLength) {
		args.PrefixLength = pulumi.IntPtr(0)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:IpAllocation"),
		},
	})
	opts = append(opts, aliases)
	var resource IpAllocation
	err := ctx.RegisterResource("azure-native:network/v20200601:IpAllocation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIpAllocation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpAllocationState, opts ...pulumi.ResourceOption) (*IpAllocation, error) {
	var resource IpAllocation
	err := ctx.ReadResource("azure-native:network/v20200601:IpAllocation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type ipAllocationState struct {
}

type IpAllocationState struct {
}

func (IpAllocationState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipAllocationState)(nil)).Elem()
}

type ipAllocationArgs struct {
	AllocationTags    map[string]string `pulumi:"allocationTags"`
	Id                *string           `pulumi:"id"`
	IpAllocationName  *string           `pulumi:"ipAllocationName"`
	IpamAllocationId  *string           `pulumi:"ipamAllocationId"`
	Location          *string           `pulumi:"location"`
	Prefix            *string           `pulumi:"prefix"`
	PrefixLength      *int              `pulumi:"prefixLength"`
	PrefixType        *string           `pulumi:"prefixType"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
}


type IpAllocationArgs struct {
	AllocationTags    pulumi.StringMapInput
	Id                pulumi.StringPtrInput
	IpAllocationName  pulumi.StringPtrInput
	IpamAllocationId  pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Prefix            pulumi.StringPtrInput
	PrefixLength      pulumi.IntPtrInput
	PrefixType        pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
}

func (IpAllocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipAllocationArgs)(nil)).Elem()
}

type IpAllocationInput interface {
	pulumi.Input

	ToIpAllocationOutput() IpAllocationOutput
	ToIpAllocationOutputWithContext(ctx context.Context) IpAllocationOutput
}

func (*IpAllocation) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAllocation)(nil)).Elem()
}

func (i *IpAllocation) ToIpAllocationOutput() IpAllocationOutput {
	return i.ToIpAllocationOutputWithContext(context.Background())
}

func (i *IpAllocation) ToIpAllocationOutputWithContext(ctx context.Context) IpAllocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAllocationOutput)
}

type IpAllocationOutput struct{ *pulumi.OutputState }

func (IpAllocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAllocation)(nil)).Elem()
}

func (o IpAllocationOutput) ToIpAllocationOutput() IpAllocationOutput {
	return o
}

func (o IpAllocationOutput) ToIpAllocationOutputWithContext(ctx context.Context) IpAllocationOutput {
	return o
}

func (o IpAllocationOutput) AllocationTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.StringMapOutput { return v.AllocationTags }).(pulumi.StringMapOutput)
}

func (o IpAllocationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o IpAllocationOutput) IpamAllocationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.StringPtrOutput { return v.IpamAllocationId }).(pulumi.StringPtrOutput)
}

func (o IpAllocationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o IpAllocationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IpAllocationOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.StringPtrOutput { return v.Prefix }).(pulumi.StringPtrOutput)
}

func (o IpAllocationOutput) PrefixLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.IntPtrOutput { return v.PrefixLength }).(pulumi.IntPtrOutput)
}

func (o IpAllocationOutput) PrefixType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.StringPtrOutput { return v.PrefixType }).(pulumi.StringPtrOutput)
}

func (o IpAllocationOutput) Subnet() SubResourceResponseOutput {
	return o.ApplyT(func(v *IpAllocation) SubResourceResponseOutput { return v.Subnet }).(SubResourceResponseOutput)
}

func (o IpAllocationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o IpAllocationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o IpAllocationOutput) VirtualNetwork() SubResourceResponseOutput {
	return o.ApplyT(func(v *IpAllocation) SubResourceResponseOutput { return v.VirtualNetwork }).(SubResourceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(IpAllocationOutput{})
}
