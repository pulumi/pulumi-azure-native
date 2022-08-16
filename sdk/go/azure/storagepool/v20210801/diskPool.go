


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DiskPool struct {
	pulumi.CustomResourceState

	AdditionalCapabilities pulumi.StringArrayOutput     `pulumi:"additionalCapabilities"`
	AvailabilityZones      pulumi.StringArrayOutput     `pulumi:"availabilityZones"`
	Disks                  DiskResponseArrayOutput      `pulumi:"disks"`
	Location               pulumi.StringOutput          `pulumi:"location"`
	ManagedBy              pulumi.StringOutput          `pulumi:"managedBy"`
	ManagedByExtended      pulumi.StringArrayOutput     `pulumi:"managedByExtended"`
	Name                   pulumi.StringOutput          `pulumi:"name"`
	ProvisioningState      pulumi.StringOutput          `pulumi:"provisioningState"`
	Status                 pulumi.StringOutput          `pulumi:"status"`
	SubnetId               pulumi.StringOutput          `pulumi:"subnetId"`
	SystemData             SystemMetadataResponseOutput `pulumi:"systemData"`
	Tags                   pulumi.StringMapOutput       `pulumi:"tags"`
	Tier                   pulumi.StringPtrOutput       `pulumi:"tier"`
	Type                   pulumi.StringOutput          `pulumi:"type"`
}


func NewDiskPool(ctx *pulumi.Context,
	name string, args *DiskPoolArgs, opts ...pulumi.ResourceOption) (*DiskPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagepool:DiskPool"),
		},
		{
			Type: pulumi.String("azure-native:storagepool/v20200315preview:DiskPool"),
		},
		{
			Type: pulumi.String("azure-native:storagepool/v20210401preview:DiskPool"),
		},
	})
	opts = append(opts, aliases)
	var resource DiskPool
	err := ctx.RegisterResource("azure-native:storagepool/v20210801:DiskPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDiskPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskPoolState, opts ...pulumi.ResourceOption) (*DiskPool, error) {
	var resource DiskPool
	err := ctx.ReadResource("azure-native:storagepool/v20210801:DiskPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type diskPoolState struct {
}

type DiskPoolState struct {
}

func (DiskPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskPoolState)(nil)).Elem()
}

type diskPoolArgs struct {
	AdditionalCapabilities []string          `pulumi:"additionalCapabilities"`
	AvailabilityZones      []string          `pulumi:"availabilityZones"`
	DiskPoolName           *string           `pulumi:"diskPoolName"`
	Disks                  []Disk            `pulumi:"disks"`
	Location               *string           `pulumi:"location"`
	ManagedBy              *string           `pulumi:"managedBy"`
	ManagedByExtended      []string          `pulumi:"managedByExtended"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	Sku                    Sku               `pulumi:"sku"`
	SubnetId               string            `pulumi:"subnetId"`
	Tags                   map[string]string `pulumi:"tags"`
}


type DiskPoolArgs struct {
	AdditionalCapabilities pulumi.StringArrayInput
	AvailabilityZones      pulumi.StringArrayInput
	DiskPoolName           pulumi.StringPtrInput
	Disks                  DiskArrayInput
	Location               pulumi.StringPtrInput
	ManagedBy              pulumi.StringPtrInput
	ManagedByExtended      pulumi.StringArrayInput
	ResourceGroupName      pulumi.StringInput
	Sku                    SkuInput
	SubnetId               pulumi.StringInput
	Tags                   pulumi.StringMapInput
}

func (DiskPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskPoolArgs)(nil)).Elem()
}

type DiskPoolInput interface {
	pulumi.Input

	ToDiskPoolOutput() DiskPoolOutput
	ToDiskPoolOutputWithContext(ctx context.Context) DiskPoolOutput
}

func (*DiskPool) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskPool)(nil)).Elem()
}

func (i *DiskPool) ToDiskPoolOutput() DiskPoolOutput {
	return i.ToDiskPoolOutputWithContext(context.Background())
}

func (i *DiskPool) ToDiskPoolOutputWithContext(ctx context.Context) DiskPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskPoolOutput)
}

type DiskPoolOutput struct{ *pulumi.OutputState }

func (DiskPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskPool)(nil)).Elem()
}

func (o DiskPoolOutput) ToDiskPoolOutput() DiskPoolOutput {
	return o
}

func (o DiskPoolOutput) ToDiskPoolOutputWithContext(ctx context.Context) DiskPoolOutput {
	return o
}

func (o DiskPoolOutput) AdditionalCapabilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DiskPool) pulumi.StringArrayOutput { return v.AdditionalCapabilities }).(pulumi.StringArrayOutput)
}

func (o DiskPoolOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DiskPool) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o DiskPoolOutput) Disks() DiskResponseArrayOutput {
	return o.ApplyT(func(v *DiskPool) DiskResponseArrayOutput { return v.Disks }).(DiskResponseArrayOutput)
}

func (o DiskPoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskPool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DiskPoolOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskPool) pulumi.StringOutput { return v.ManagedBy }).(pulumi.StringOutput)
}

func (o DiskPoolOutput) ManagedByExtended() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DiskPool) pulumi.StringArrayOutput { return v.ManagedByExtended }).(pulumi.StringArrayOutput)
}

func (o DiskPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DiskPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DiskPoolOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskPool) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o DiskPoolOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskPool) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

func (o DiskPoolOutput) SystemData() SystemMetadataResponseOutput {
	return o.ApplyT(func(v *DiskPool) SystemMetadataResponseOutput { return v.SystemData }).(SystemMetadataResponseOutput)
}

func (o DiskPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DiskPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DiskPoolOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskPool) pulumi.StringPtrOutput { return v.Tier }).(pulumi.StringPtrOutput)
}

func (o DiskPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DiskPoolOutput{})
}
