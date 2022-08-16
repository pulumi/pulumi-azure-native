


package v20200630

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Disk struct {
	pulumi.CustomResourceState

	CreationData                 CreationDataResponseOutput                    `pulumi:"creationData"`
	DiskAccessId                 pulumi.StringPtrOutput                        `pulumi:"diskAccessId"`
	DiskIOPSReadOnly             pulumi.Float64PtrOutput                       `pulumi:"diskIOPSReadOnly"`
	DiskIOPSReadWrite            pulumi.Float64PtrOutput                       `pulumi:"diskIOPSReadWrite"`
	DiskMBpsReadOnly             pulumi.Float64PtrOutput                       `pulumi:"diskMBpsReadOnly"`
	DiskMBpsReadWrite            pulumi.Float64PtrOutput                       `pulumi:"diskMBpsReadWrite"`
	DiskSizeBytes                pulumi.Float64Output                          `pulumi:"diskSizeBytes"`
	DiskSizeGB                   pulumi.IntPtrOutput                           `pulumi:"diskSizeGB"`
	DiskState                    pulumi.StringOutput                           `pulumi:"diskState"`
	Encryption                   EncryptionResponsePtrOutput                   `pulumi:"encryption"`
	EncryptionSettingsCollection EncryptionSettingsCollectionResponsePtrOutput `pulumi:"encryptionSettingsCollection"`
	HyperVGeneration             pulumi.StringPtrOutput                        `pulumi:"hyperVGeneration"`
	Location                     pulumi.StringOutput                           `pulumi:"location"`
	ManagedBy                    pulumi.StringOutput                           `pulumi:"managedBy"`
	ManagedByExtended            pulumi.StringArrayOutput                      `pulumi:"managedByExtended"`
	MaxShares                    pulumi.IntPtrOutput                           `pulumi:"maxShares"`
	Name                         pulumi.StringOutput                           `pulumi:"name"`
	NetworkAccessPolicy          pulumi.StringPtrOutput                        `pulumi:"networkAccessPolicy"`
	OsType                       pulumi.StringPtrOutput                        `pulumi:"osType"`
	ProvisioningState            pulumi.StringOutput                           `pulumi:"provisioningState"`
	ShareInfo                    ShareInfoElementResponseArrayOutput           `pulumi:"shareInfo"`
	Sku                          DiskSkuResponsePtrOutput                      `pulumi:"sku"`
	Tags                         pulumi.StringMapOutput                        `pulumi:"tags"`
	Tier                         pulumi.StringPtrOutput                        `pulumi:"tier"`
	TimeCreated                  pulumi.StringOutput                           `pulumi:"timeCreated"`
	Type                         pulumi.StringOutput                           `pulumi:"type"`
	UniqueId                     pulumi.StringOutput                           `pulumi:"uniqueId"`
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
			Type: pulumi.String("azure-native:compute/v20180930:Disk"),
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
	err := ctx.RegisterResource("azure-native:compute/v20200630:Disk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskState, opts ...pulumi.ResourceOption) (*Disk, error) {
	var resource Disk
	err := ctx.ReadResource("azure-native:compute/v20200630:Disk", name, id, state, &resource, opts...)
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
	DiskAccessId                 *string                       `pulumi:"diskAccessId"`
	DiskIOPSReadOnly             *float64                      `pulumi:"diskIOPSReadOnly"`
	DiskIOPSReadWrite            *float64                      `pulumi:"diskIOPSReadWrite"`
	DiskMBpsReadOnly             *float64                      `pulumi:"diskMBpsReadOnly"`
	DiskMBpsReadWrite            *float64                      `pulumi:"diskMBpsReadWrite"`
	DiskName                     *string                       `pulumi:"diskName"`
	DiskSizeGB                   *int                          `pulumi:"diskSizeGB"`
	Encryption                   *Encryption                   `pulumi:"encryption"`
	EncryptionSettingsCollection *EncryptionSettingsCollection `pulumi:"encryptionSettingsCollection"`
	HyperVGeneration             *string                       `pulumi:"hyperVGeneration"`
	Location                     *string                       `pulumi:"location"`
	MaxShares                    *int                          `pulumi:"maxShares"`
	NetworkAccessPolicy          *string                       `pulumi:"networkAccessPolicy"`
	OsType                       *OperatingSystemTypes         `pulumi:"osType"`
	ResourceGroupName            string                        `pulumi:"resourceGroupName"`
	Sku                          *DiskSku                      `pulumi:"sku"`
	Tags                         map[string]string             `pulumi:"tags"`
	Tier                         *string                       `pulumi:"tier"`
	Zones                        []string                      `pulumi:"zones"`
}


type DiskArgs struct {
	CreationData                 CreationDataInput
	DiskAccessId                 pulumi.StringPtrInput
	DiskIOPSReadOnly             pulumi.Float64PtrInput
	DiskIOPSReadWrite            pulumi.Float64PtrInput
	DiskMBpsReadOnly             pulumi.Float64PtrInput
	DiskMBpsReadWrite            pulumi.Float64PtrInput
	DiskName                     pulumi.StringPtrInput
	DiskSizeGB                   pulumi.IntPtrInput
	Encryption                   EncryptionPtrInput
	EncryptionSettingsCollection EncryptionSettingsCollectionPtrInput
	HyperVGeneration             pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	MaxShares                    pulumi.IntPtrInput
	NetworkAccessPolicy          pulumi.StringPtrInput
	OsType                       OperatingSystemTypesPtrInput
	ResourceGroupName            pulumi.StringInput
	Sku                          DiskSkuPtrInput
	Tags                         pulumi.StringMapInput
	Tier                         pulumi.StringPtrInput
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

func (o DiskOutput) DiskAccessId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.DiskAccessId }).(pulumi.StringPtrOutput)
}

func (o DiskOutput) DiskIOPSReadOnly() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.Float64PtrOutput { return v.DiskIOPSReadOnly }).(pulumi.Float64PtrOutput)
}

func (o DiskOutput) DiskIOPSReadWrite() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.Float64PtrOutput { return v.DiskIOPSReadWrite }).(pulumi.Float64PtrOutput)
}

func (o DiskOutput) DiskMBpsReadOnly() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.Float64PtrOutput { return v.DiskMBpsReadOnly }).(pulumi.Float64PtrOutput)
}

func (o DiskOutput) DiskMBpsReadWrite() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.Float64PtrOutput { return v.DiskMBpsReadWrite }).(pulumi.Float64PtrOutput)
}

func (o DiskOutput) DiskSizeBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *Disk) pulumi.Float64Output { return v.DiskSizeBytes }).(pulumi.Float64Output)
}

func (o DiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.IntPtrOutput { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o DiskOutput) DiskState() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.DiskState }).(pulumi.StringOutput)
}

func (o DiskOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v *Disk) EncryptionResponsePtrOutput { return v.Encryption }).(EncryptionResponsePtrOutput)
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

func (o DiskOutput) ManagedByExtended() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringArrayOutput { return v.ManagedByExtended }).(pulumi.StringArrayOutput)
}

func (o DiskOutput) MaxShares() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.IntPtrOutput { return v.MaxShares }).(pulumi.IntPtrOutput)
}

func (o DiskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DiskOutput) NetworkAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.NetworkAccessPolicy }).(pulumi.StringPtrOutput)
}

func (o DiskOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o DiskOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DiskOutput) ShareInfo() ShareInfoElementResponseArrayOutput {
	return o.ApplyT(func(v *Disk) ShareInfoElementResponseArrayOutput { return v.ShareInfo }).(ShareInfoElementResponseArrayOutput)
}

func (o DiskOutput) Sku() DiskSkuResponsePtrOutput {
	return o.ApplyT(func(v *Disk) DiskSkuResponsePtrOutput { return v.Sku }).(DiskSkuResponsePtrOutput)
}

func (o DiskOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DiskOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.Tier }).(pulumi.StringPtrOutput)
}

func (o DiskOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.TimeCreated }).(pulumi.StringOutput)
}

func (o DiskOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o DiskOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.UniqueId }).(pulumi.StringOutput)
}

func (o DiskOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(DiskOutput{})
}
