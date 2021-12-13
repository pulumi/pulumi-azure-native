


package v20210401preview

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
			Type: pulumi.String("azure-native:storagepool/v20210801:DiskPool"),
		},
	})
	opts = append(opts, aliases)
	var resource DiskPool
	err := ctx.RegisterResource("azure-native:storagepool/v20210401preview:DiskPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDiskPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskPoolState, opts ...pulumi.ResourceOption) (*DiskPool, error) {
	var resource DiskPool
	err := ctx.ReadResource("azure-native:storagepool/v20210401preview:DiskPool", name, id, state, &resource, opts...)
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

func init() {
	pulumi.RegisterOutputType(DiskPoolOutput{})
}
