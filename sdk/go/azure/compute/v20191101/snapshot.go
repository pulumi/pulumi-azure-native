


package v20191101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Snapshot struct {
	pulumi.CustomResourceState

	CreationData                 CreationDataResponseOutput                    `pulumi:"creationData"`
	DiskSizeBytes                pulumi.Float64Output                          `pulumi:"diskSizeBytes"`
	DiskSizeGB                   pulumi.IntPtrOutput                           `pulumi:"diskSizeGB"`
	Encryption                   EncryptionResponsePtrOutput                   `pulumi:"encryption"`
	EncryptionSettingsCollection EncryptionSettingsCollectionResponsePtrOutput `pulumi:"encryptionSettingsCollection"`
	HyperVGeneration             pulumi.StringPtrOutput                        `pulumi:"hyperVGeneration"`
	Incremental                  pulumi.BoolPtrOutput                          `pulumi:"incremental"`
	Location                     pulumi.StringOutput                           `pulumi:"location"`
	ManagedBy                    pulumi.StringOutput                           `pulumi:"managedBy"`
	Name                         pulumi.StringOutput                           `pulumi:"name"`
	OsType                       pulumi.StringPtrOutput                        `pulumi:"osType"`
	ProvisioningState            pulumi.StringOutput                           `pulumi:"provisioningState"`
	Sku                          SnapshotSkuResponsePtrOutput                  `pulumi:"sku"`
	Tags                         pulumi.StringMapOutput                        `pulumi:"tags"`
	TimeCreated                  pulumi.StringOutput                           `pulumi:"timeCreated"`
	Type                         pulumi.StringOutput                           `pulumi:"type"`
	UniqueId                     pulumi.StringOutput                           `pulumi:"uniqueId"`
}


func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
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
			Type: pulumi.String("azure-native:compute:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160430preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180930:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200501:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200630:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210801:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211201:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220302:Snapshot"),
		},
	})
	opts = append(opts, aliases)
	var resource Snapshot
	err := ctx.RegisterResource("azure-native:compute/v20191101:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("azure-native:compute/v20191101:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type snapshotState struct {
}

type SnapshotState struct {
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	CreationData                 CreationData                  `pulumi:"creationData"`
	DiskSizeGB                   *int                          `pulumi:"diskSizeGB"`
	Encryption                   *Encryption                   `pulumi:"encryption"`
	EncryptionSettingsCollection *EncryptionSettingsCollection `pulumi:"encryptionSettingsCollection"`
	HyperVGeneration             *string                       `pulumi:"hyperVGeneration"`
	Incremental                  *bool                         `pulumi:"incremental"`
	Location                     *string                       `pulumi:"location"`
	OsType                       *OperatingSystemTypes         `pulumi:"osType"`
	ResourceGroupName            string                        `pulumi:"resourceGroupName"`
	Sku                          *SnapshotSku                  `pulumi:"sku"`
	SnapshotName                 *string                       `pulumi:"snapshotName"`
	Tags                         map[string]string             `pulumi:"tags"`
}


type SnapshotArgs struct {
	CreationData                 CreationDataInput
	DiskSizeGB                   pulumi.IntPtrInput
	Encryption                   EncryptionPtrInput
	EncryptionSettingsCollection EncryptionSettingsCollectionPtrInput
	HyperVGeneration             pulumi.StringPtrInput
	Incremental                  pulumi.BoolPtrInput
	Location                     pulumi.StringPtrInput
	OsType                       OperatingSystemTypesPtrInput
	ResourceGroupName            pulumi.StringInput
	Sku                          SnapshotSkuPtrInput
	SnapshotName                 pulumi.StringPtrInput
	Tags                         pulumi.StringMapInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

func (o SnapshotOutput) CreationData() CreationDataResponseOutput {
	return o.ApplyT(func(v *Snapshot) CreationDataResponseOutput { return v.CreationData }).(CreationDataResponseOutput)
}

func (o SnapshotOutput) DiskSizeBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *Snapshot) pulumi.Float64Output { return v.DiskSizeBytes }).(pulumi.Float64Output)
}

func (o SnapshotOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.IntPtrOutput { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o SnapshotOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v *Snapshot) EncryptionResponsePtrOutput { return v.Encryption }).(EncryptionResponsePtrOutput)
}

func (o SnapshotOutput) EncryptionSettingsCollection() EncryptionSettingsCollectionResponsePtrOutput {
	return o.ApplyT(func(v *Snapshot) EncryptionSettingsCollectionResponsePtrOutput { return v.EncryptionSettingsCollection }).(EncryptionSettingsCollectionResponsePtrOutput)
}

func (o SnapshotOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

func (o SnapshotOutput) Incremental() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.BoolPtrOutput { return v.Incremental }).(pulumi.BoolPtrOutput)
}

func (o SnapshotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SnapshotOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.ManagedBy }).(pulumi.StringOutput)
}

func (o SnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SnapshotOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o SnapshotOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SnapshotOutput) Sku() SnapshotSkuResponsePtrOutput {
	return o.ApplyT(func(v *Snapshot) SnapshotSkuResponsePtrOutput { return v.Sku }).(SnapshotSkuResponsePtrOutput)
}

func (o SnapshotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SnapshotOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.TimeCreated }).(pulumi.StringOutput)
}

func (o SnapshotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SnapshotOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.UniqueId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SnapshotOutput{})
}
