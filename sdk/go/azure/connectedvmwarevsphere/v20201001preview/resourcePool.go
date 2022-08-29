


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourcePool struct {
	pulumi.CustomResourceState

	CpuLimitMHz        pulumi.Float64Output              `pulumi:"cpuLimitMHz"`
	CpuReservationMHz  pulumi.Float64Output              `pulumi:"cpuReservationMHz"`
	CpuSharesLevel     pulumi.StringOutput               `pulumi:"cpuSharesLevel"`
	CustomResourceName pulumi.StringOutput               `pulumi:"customResourceName"`
	ExtendedLocation   ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	InventoryItemId    pulumi.StringPtrOutput            `pulumi:"inventoryItemId"`
	Kind               pulumi.StringPtrOutput            `pulumi:"kind"`
	Location           pulumi.StringOutput               `pulumi:"location"`
	MemLimitMB         pulumi.Float64Output              `pulumi:"memLimitMB"`
	MemReservationMB   pulumi.Float64Output              `pulumi:"memReservationMB"`
	MemSharesLevel     pulumi.StringOutput               `pulumi:"memSharesLevel"`
	MoName             pulumi.StringOutput               `pulumi:"moName"`
	MoRefId            pulumi.StringPtrOutput            `pulumi:"moRefId"`
	Name               pulumi.StringOutput               `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput               `pulumi:"provisioningState"`
	Statuses           ResourceStatusResponseArrayOutput `pulumi:"statuses"`
	SystemData         SystemDataResponseOutput          `pulumi:"systemData"`
	Tags               pulumi.StringMapOutput            `pulumi:"tags"`
	Type               pulumi.StringOutput               `pulumi:"type"`
	Uuid               pulumi.StringOutput               `pulumi:"uuid"`
	VCenterId          pulumi.StringPtrOutput            `pulumi:"vCenterId"`
}


func NewResourcePool(ctx *pulumi.Context,
	name string, args *ResourcePoolArgs, opts ...pulumi.ResourceOption) (*ResourcePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere:ResourcePool"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220110preview:ResourcePool"),
		},
	})
	opts = append(opts, aliases)
	var resource ResourcePool
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere/v20201001preview:ResourcePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetResourcePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourcePoolState, opts ...pulumi.ResourceOption) (*ResourcePool, error) {
	var resource ResourcePool
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere/v20201001preview:ResourcePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type resourcePoolState struct {
}

type ResourcePoolState struct {
}

func (ResourcePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePoolState)(nil)).Elem()
}

type resourcePoolArgs struct {
	ExtendedLocation  *ExtendedLocation `pulumi:"extendedLocation"`
	InventoryItemId   *string           `pulumi:"inventoryItemId"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	MoRefId           *string           `pulumi:"moRefId"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourcePoolName  *string           `pulumi:"resourcePoolName"`
	Tags              map[string]string `pulumi:"tags"`
	VCenterId         *string           `pulumi:"vCenterId"`
}


type ResourcePoolArgs struct {
	ExtendedLocation  ExtendedLocationPtrInput
	InventoryItemId   pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	MoRefId           pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourcePoolName  pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	VCenterId         pulumi.StringPtrInput
}

func (ResourcePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePoolArgs)(nil)).Elem()
}

type ResourcePoolInput interface {
	pulumi.Input

	ToResourcePoolOutput() ResourcePoolOutput
	ToResourcePoolOutputWithContext(ctx context.Context) ResourcePoolOutput
}

func (*ResourcePool) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePool)(nil)).Elem()
}

func (i *ResourcePool) ToResourcePoolOutput() ResourcePoolOutput {
	return i.ToResourcePoolOutputWithContext(context.Background())
}

func (i *ResourcePool) ToResourcePoolOutputWithContext(ctx context.Context) ResourcePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePoolOutput)
}

type ResourcePoolOutput struct{ *pulumi.OutputState }

func (ResourcePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePool)(nil)).Elem()
}

func (o ResourcePoolOutput) ToResourcePoolOutput() ResourcePoolOutput {
	return o
}

func (o ResourcePoolOutput) ToResourcePoolOutputWithContext(ctx context.Context) ResourcePoolOutput {
	return o
}

func (o ResourcePoolOutput) CpuLimitMHz() pulumi.Float64Output {
	return o.ApplyT(func(v *ResourcePool) pulumi.Float64Output { return v.CpuLimitMHz }).(pulumi.Float64Output)
}

func (o ResourcePoolOutput) CpuReservationMHz() pulumi.Float64Output {
	return o.ApplyT(func(v *ResourcePool) pulumi.Float64Output { return v.CpuReservationMHz }).(pulumi.Float64Output)
}

func (o ResourcePoolOutput) CpuSharesLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.CpuSharesLevel }).(pulumi.StringOutput)
}

func (o ResourcePoolOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.CustomResourceName }).(pulumi.StringOutput)
}

func (o ResourcePoolOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *ResourcePool) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o ResourcePoolOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringPtrOutput { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o ResourcePoolOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ResourcePoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ResourcePoolOutput) MemLimitMB() pulumi.Float64Output {
	return o.ApplyT(func(v *ResourcePool) pulumi.Float64Output { return v.MemLimitMB }).(pulumi.Float64Output)
}

func (o ResourcePoolOutput) MemReservationMB() pulumi.Float64Output {
	return o.ApplyT(func(v *ResourcePool) pulumi.Float64Output { return v.MemReservationMB }).(pulumi.Float64Output)
}

func (o ResourcePoolOutput) MemSharesLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.MemSharesLevel }).(pulumi.StringOutput)
}

func (o ResourcePoolOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.MoName }).(pulumi.StringOutput)
}

func (o ResourcePoolOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringPtrOutput { return v.MoRefId }).(pulumi.StringPtrOutput)
}

func (o ResourcePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ResourcePoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ResourcePoolOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v *ResourcePool) ResourceStatusResponseArrayOutput { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

func (o ResourcePoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ResourcePool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ResourcePoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ResourcePoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ResourcePoolOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func (o ResourcePoolOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePool) pulumi.StringPtrOutput { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourcePoolOutput{})
}
