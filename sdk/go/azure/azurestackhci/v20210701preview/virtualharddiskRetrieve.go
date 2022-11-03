


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualharddiskRetrieve struct {
	pulumi.CustomResourceState

	BlockSizeBytes      pulumi.IntPtrOutput                 `pulumi:"blockSizeBytes"`
	DiskSizeBytes       pulumi.Float64PtrOutput             `pulumi:"diskSizeBytes"`
	Dynamic             pulumi.BoolPtrOutput                `pulumi:"dynamic"`
	ExtendedLocation    ExtendedLocationResponsePtrOutput   `pulumi:"extendedLocation"`
	Location            pulumi.StringOutput                 `pulumi:"location"`
	LogicalSectorBytes  pulumi.IntPtrOutput                 `pulumi:"logicalSectorBytes"`
	Name                pulumi.StringOutput                 `pulumi:"name"`
	PhysicalSectorBytes pulumi.IntPtrOutput                 `pulumi:"physicalSectorBytes"`
	ProvisioningState   pulumi.StringOutput                 `pulumi:"provisioningState"`
	ResourceName        pulumi.StringPtrOutput              `pulumi:"resourceName"`
	Status              VirtualHardDiskStatusResponseOutput `pulumi:"status"`
	SystemData          SystemDataResponseOutput            `pulumi:"systemData"`
	Tags                pulumi.StringMapOutput              `pulumi:"tags"`
	Type                pulumi.StringOutput                 `pulumi:"type"`
}


func NewVirtualharddiskRetrieve(ctx *pulumi.Context,
	name string, args *VirtualharddiskRetrieveArgs, opts ...pulumi.ResourceOption) (*VirtualharddiskRetrieve, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource VirtualharddiskRetrieve
	err := ctx.RegisterResource("azure-native:azurestackhci/v20210701preview:virtualharddiskRetrieve", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualharddiskRetrieve(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualharddiskRetrieveState, opts ...pulumi.ResourceOption) (*VirtualharddiskRetrieve, error) {
	var resource VirtualharddiskRetrieve
	err := ctx.ReadResource("azure-native:azurestackhci/v20210701preview:virtualharddiskRetrieve", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualharddiskRetrieveState struct {
}

type VirtualharddiskRetrieveState struct {
}

func (VirtualharddiskRetrieveState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualharddiskRetrieveState)(nil)).Elem()
}

type virtualharddiskRetrieveArgs struct {
	BlockSizeBytes       *int              `pulumi:"blockSizeBytes"`
	DiskSizeBytes        *float64          `pulumi:"diskSizeBytes"`
	Dynamic              *bool             `pulumi:"dynamic"`
	ExtendedLocation     *ExtendedLocation `pulumi:"extendedLocation"`
	Location             *string           `pulumi:"location"`
	LogicalSectorBytes   *int              `pulumi:"logicalSectorBytes"`
	PhysicalSectorBytes  *int              `pulumi:"physicalSectorBytes"`
	ResourceGroupName    string            `pulumi:"resourceGroupName"`
	ResourceName         *string           `pulumi:"resourceName"`
	Tags                 map[string]string `pulumi:"tags"`
	VirtualharddisksName *string           `pulumi:"virtualharddisksName"`
}


type VirtualharddiskRetrieveArgs struct {
	BlockSizeBytes       pulumi.IntPtrInput
	DiskSizeBytes        pulumi.Float64PtrInput
	Dynamic              pulumi.BoolPtrInput
	ExtendedLocation     ExtendedLocationPtrInput
	Location             pulumi.StringPtrInput
	LogicalSectorBytes   pulumi.IntPtrInput
	PhysicalSectorBytes  pulumi.IntPtrInput
	ResourceGroupName    pulumi.StringInput
	ResourceName         pulumi.StringPtrInput
	Tags                 pulumi.StringMapInput
	VirtualharddisksName pulumi.StringPtrInput
}

func (VirtualharddiskRetrieveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualharddiskRetrieveArgs)(nil)).Elem()
}

type VirtualharddiskRetrieveInput interface {
	pulumi.Input

	ToVirtualharddiskRetrieveOutput() VirtualharddiskRetrieveOutput
	ToVirtualharddiskRetrieveOutputWithContext(ctx context.Context) VirtualharddiskRetrieveOutput
}

func (*VirtualharddiskRetrieve) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualharddiskRetrieve)(nil)).Elem()
}

func (i *VirtualharddiskRetrieve) ToVirtualharddiskRetrieveOutput() VirtualharddiskRetrieveOutput {
	return i.ToVirtualharddiskRetrieveOutputWithContext(context.Background())
}

func (i *VirtualharddiskRetrieve) ToVirtualharddiskRetrieveOutputWithContext(ctx context.Context) VirtualharddiskRetrieveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualharddiskRetrieveOutput)
}

type VirtualharddiskRetrieveOutput struct{ *pulumi.OutputState }

func (VirtualharddiskRetrieveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualharddiskRetrieve)(nil)).Elem()
}

func (o VirtualharddiskRetrieveOutput) ToVirtualharddiskRetrieveOutput() VirtualharddiskRetrieveOutput {
	return o
}

func (o VirtualharddiskRetrieveOutput) ToVirtualharddiskRetrieveOutputWithContext(ctx context.Context) VirtualharddiskRetrieveOutput {
	return o
}

func (o VirtualharddiskRetrieveOutput) BlockSizeBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualharddiskRetrieve) pulumi.IntPtrOutput { return v.BlockSizeBytes }).(pulumi.IntPtrOutput)
}

func (o VirtualharddiskRetrieveOutput) DiskSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualharddiskRetrieve) pulumi.Float64PtrOutput { return v.DiskSizeBytes }).(pulumi.Float64PtrOutput)
}

func (o VirtualharddiskRetrieveOutput) Dynamic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualharddiskRetrieve) pulumi.BoolPtrOutput { return v.Dynamic }).(pulumi.BoolPtrOutput)
}

func (o VirtualharddiskRetrieveOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualharddiskRetrieve) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o VirtualharddiskRetrieveOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualharddiskRetrieve) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualharddiskRetrieveOutput) LogicalSectorBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualharddiskRetrieve) pulumi.IntPtrOutput { return v.LogicalSectorBytes }).(pulumi.IntPtrOutput)
}

func (o VirtualharddiskRetrieveOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualharddiskRetrieve) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualharddiskRetrieveOutput) PhysicalSectorBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualharddiskRetrieve) pulumi.IntPtrOutput { return v.PhysicalSectorBytes }).(pulumi.IntPtrOutput)
}

func (o VirtualharddiskRetrieveOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualharddiskRetrieve) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualharddiskRetrieveOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualharddiskRetrieve) pulumi.StringPtrOutput { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o VirtualharddiskRetrieveOutput) Status() VirtualHardDiskStatusResponseOutput {
	return o.ApplyT(func(v *VirtualharddiskRetrieve) VirtualHardDiskStatusResponseOutput { return v.Status }).(VirtualHardDiskStatusResponseOutput)
}

func (o VirtualharddiskRetrieveOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VirtualharddiskRetrieve) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VirtualharddiskRetrieveOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualharddiskRetrieve) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualharddiskRetrieveOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualharddiskRetrieve) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualharddiskRetrieveOutput{})
}
