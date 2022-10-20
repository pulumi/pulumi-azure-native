


package v20180930

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Disk struct {
	pulumi.CustomResourceState

	CreationData                 CreationDataResponseOutput                    `pulumi:"creationData"`
	DiskIOPSReadWrite            pulumi.Float64PtrOutput                       `pulumi:"diskIOPSReadWrite"`
	DiskMBpsReadWrite            pulumi.IntPtrOutput                           `pulumi:"diskMBpsReadWrite"`
	DiskSizeGB                   pulumi.IntPtrOutput                           `pulumi:"diskSizeGB"`
	DiskState                    pulumi.StringOutput                           `pulumi:"diskState"`
	EncryptionSettingsCollection EncryptionSettingsCollectionResponsePtrOutput `pulumi:"encryptionSettingsCollection"`
	HyperVGeneration             pulumi.StringPtrOutput                        `pulumi:"hyperVGeneration"`
	Location                     pulumi.StringOutput                           `pulumi:"location"`
	ManagedBy                    pulumi.StringOutput                           `pulumi:"managedBy"`
	Name                         pulumi.StringOutput                           `pulumi:"name"`
	OsType                       pulumi.StringPtrOutput                        `pulumi:"osType"`
	ProvisioningState            pulumi.StringOutput                           `pulumi:"provisioningState"`
	Sku                          DiskSkuResponsePtrOutput                      `pulumi:"sku"`
	Tags                         pulumi.StringMapOutput                        `pulumi:"tags"`
	TimeCreated                  pulumi.StringOutput                           `pulumi:"timeCreated"`
	Type                         pulumi.StringOutput                           `pulumi:"type"`
	Zones                        pulumi.StringArrayOutput                      `pulumi:"zones"`
}


func NewDisk(ctx *pulumi.Context,
	name string, args *DiskArgs, opts ...pulumi.ResourceOption) (*Disk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CreationData == nil {
		return nil, errors.New("invalid value for required argument 'CreationData'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160430preview:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191101:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200501:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200630:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210801:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211201:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220302:Disk"),
		},
	})
	opts = append(opts, aliases)
	var resource Disk
	err := ctx.RegisterResource("azure-native:compute/v20180930:Disk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskState, opts ...pulumi.ResourceOption) (*Disk, error) {
	var resource Disk
	err := ctx.ReadResource("azure-native:compute/v20180930:Disk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type diskState struct {
}

type DiskState struct {
}

func (DiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskState)(nil)).Elem()
}

type diskArgs struct {
	CreationData                 CreationData                  `pulumi:"creationData"`
	DiskIOPSReadWrite            *float64                      `pulumi:"diskIOPSReadWrite"`
	DiskMBpsReadWrite            *int                          `pulumi:"diskMBpsReadWrite"`
	DiskName                     *string                       `pulumi:"diskName"`
	DiskSizeGB                   *int                          `pulumi:"diskSizeGB"`
	EncryptionSettingsCollection *EncryptionSettingsCollection `pulumi:"encryptionSettingsCollection"`
	HyperVGeneration             *string                       `pulumi:"hyperVGeneration"`
	Location                     *string                       `pulumi:"location"`
	OsType                       *OperatingSystemTypes         `pulumi:"osType"`
	ResourceGroupName            string                        `pulumi:"resourceGroupName"`
	Sku                          *DiskSku                      `pulumi:"sku"`
	Tags                         map[string]string             `pulumi:"tags"`
	Zones                        []string                      `pulumi:"zones"`
}


type DiskArgs struct {
	CreationData                 CreationDataInput
	DiskIOPSReadWrite            pulumi.Float64PtrInput
	DiskMBpsReadWrite            pulumi.IntPtrInput
	DiskName                     pulumi.StringPtrInput
	DiskSizeGB                   pulumi.IntPtrInput
	EncryptionSettingsCollection EncryptionSettingsCollectionPtrInput
	HyperVGeneration             pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	OsType                       OperatingSystemTypesPtrInput
	ResourceGroupName            pulumi.StringInput
	Sku                          DiskSkuPtrInput
	Tags                         pulumi.StringMapInput
	Zones                        pulumi.StringArrayInput
}

func (DiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskArgs)(nil)).Elem()
}

type DiskInput interface {
	pulumi.Input

	ToDiskOutput() DiskOutput
	ToDiskOutputWithContext(ctx context.Context) DiskOutput
}

func (*Disk) ElementType() reflect.Type {
	return reflect.TypeOf((**Disk)(nil)).Elem()
}

func (i *Disk) ToDiskOutput() DiskOutput {
	return i.ToDiskOutputWithContext(context.Background())
}

func (i *Disk) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskOutput)
}

type DiskOutput struct{ *pulumi.OutputState }

func (DiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Disk)(nil)).Elem()
}

func (o DiskOutput) ToDiskOutput() DiskOutput {
	return o
}

func (o DiskOutput) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return o
}

func (o DiskOutput) CreationData() CreationDataResponseOutput {
	return o.ApplyT(func(v *Disk) CreationDataResponseOutput { return v.CreationData }).(CreationDataResponseOutput)
}

func (o DiskOutput) DiskIOPSReadWrite() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.Float64PtrOutput { return v.DiskIOPSReadWrite }).(pulumi.Float64PtrOutput)
}

func (o DiskOutput) DiskMBpsReadWrite() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.IntPtrOutput { return v.DiskMBpsReadWrite }).(pulumi.IntPtrOutput)
}

func (o DiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.IntPtrOutput { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o DiskOutput) DiskState() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.DiskState }).(pulumi.StringOutput)
}

func (o DiskOutput) EncryptionSettingsCollection() EncryptionSettingsCollectionResponsePtrOutput {
	return o.ApplyT(func(v *Disk) EncryptionSettingsCollectionResponsePtrOutput { return v.EncryptionSettingsCollection }).(EncryptionSettingsCollectionResponsePtrOutput)
}

func (o DiskOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

func (o DiskOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DiskOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.ManagedBy }).(pulumi.StringOutput)
}

func (o DiskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DiskOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o DiskOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DiskOutput) Sku() DiskSkuResponsePtrOutput {
	return o.ApplyT(func(v *Disk) DiskSkuResponsePtrOutput { return v.Sku }).(DiskSkuResponsePtrOutput)
}

func (o DiskOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DiskOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.TimeCreated }).(pulumi.StringOutput)
}

func (o DiskOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o DiskOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(DiskOutput{})
}
