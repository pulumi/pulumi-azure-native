


package v20210401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Snapshot struct {
	pulumi.CustomResourceState

	CompletionPercent            pulumi.Float64PtrOutput                       `pulumi:"completionPercent"`
	CreationData                 CreationDataResponseOutput                    `pulumi:"creationData"`
	DiskAccessId                 pulumi.StringPtrOutput                        `pulumi:"diskAccessId"`
	DiskSizeBytes                pulumi.Float64Output                          `pulumi:"diskSizeBytes"`
	DiskSizeGB                   pulumi.IntPtrOutput                           `pulumi:"diskSizeGB"`
	DiskState                    pulumi.StringOutput                           `pulumi:"diskState"`
	Encryption                   EncryptionResponsePtrOutput                   `pulumi:"encryption"`
	EncryptionSettingsCollection EncryptionSettingsCollectionResponsePtrOutput `pulumi:"encryptionSettingsCollection"`
	ExtendedLocation             ExtendedLocationResponsePtrOutput             `pulumi:"extendedLocation"`
	HyperVGeneration             pulumi.StringPtrOutput                        `pulumi:"hyperVGeneration"`
	Incremental                  pulumi.BoolPtrOutput                          `pulumi:"incremental"`
	Location                     pulumi.StringOutput                           `pulumi:"location"`
	ManagedBy                    pulumi.StringOutput                           `pulumi:"managedBy"`
	Name                         pulumi.StringOutput                           `pulumi:"name"`
	NetworkAccessPolicy          pulumi.StringPtrOutput                        `pulumi:"networkAccessPolicy"`
	OsType                       pulumi.StringPtrOutput                        `pulumi:"osType"`
	ProvisioningState            pulumi.StringOutput                           `pulumi:"provisioningState"`
	PublicNetworkAccess          pulumi.StringPtrOutput                        `pulumi:"publicNetworkAccess"`
	PurchasePlan                 PurchasePlanResponsePtrOutput                 `pulumi:"purchasePlan"`
	Sku                          SnapshotSkuResponsePtrOutput                  `pulumi:"sku"`
	SupportedCapabilities        SupportedCapabilitiesResponsePtrOutput        `pulumi:"supportedCapabilities"`
	SupportsHibernation          pulumi.BoolPtrOutput                          `pulumi:"supportsHibernation"`
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
			Type: pulumi.String("azure-native:compute/v20191101:Snapshot"),
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
	})
	opts = append(opts, aliases)
	var resource Snapshot
	err := ctx.RegisterResource("azure-native:compute/v20210401:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("azure-native:compute/v20210401:Snapshot", name, id, state, &resource, opts...)
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
	CompletionPercent            *float64                      `pulumi:"completionPercent"`
	CreationData                 CreationData                  `pulumi:"creationData"`
	DiskAccessId                 *string                       `pulumi:"diskAccessId"`
	DiskSizeGB                   *int                          `pulumi:"diskSizeGB"`
	Encryption                   *Encryption                   `pulumi:"encryption"`
	EncryptionSettingsCollection *EncryptionSettingsCollection `pulumi:"encryptionSettingsCollection"`
	ExtendedLocation             *ExtendedLocation             `pulumi:"extendedLocation"`
	HyperVGeneration             *string                       `pulumi:"hyperVGeneration"`
	Incremental                  *bool                         `pulumi:"incremental"`
	Location                     *string                       `pulumi:"location"`
	NetworkAccessPolicy          *string                       `pulumi:"networkAccessPolicy"`
	OsType                       *OperatingSystemTypes         `pulumi:"osType"`
	PublicNetworkAccess          *string                       `pulumi:"publicNetworkAccess"`
	PurchasePlan                 *PurchasePlan                 `pulumi:"purchasePlan"`
	ResourceGroupName            string                        `pulumi:"resourceGroupName"`
	Sku                          *SnapshotSku                  `pulumi:"sku"`
	SnapshotName                 *string                       `pulumi:"snapshotName"`
	SupportedCapabilities        *SupportedCapabilities        `pulumi:"supportedCapabilities"`
	SupportsHibernation          *bool                         `pulumi:"supportsHibernation"`
	Tags                         map[string]string             `pulumi:"tags"`
}


type SnapshotArgs struct {
	CompletionPercent            pulumi.Float64PtrInput
	CreationData                 CreationDataInput
	DiskAccessId                 pulumi.StringPtrInput
	DiskSizeGB                   pulumi.IntPtrInput
	Encryption                   EncryptionPtrInput
	EncryptionSettingsCollection EncryptionSettingsCollectionPtrInput
	ExtendedLocation             ExtendedLocationPtrInput
	HyperVGeneration             pulumi.StringPtrInput
	Incremental                  pulumi.BoolPtrInput
	Location                     pulumi.StringPtrInput
	NetworkAccessPolicy          pulumi.StringPtrInput
	OsType                       OperatingSystemTypesPtrInput
	PublicNetworkAccess          pulumi.StringPtrInput
	PurchasePlan                 PurchasePlanPtrInput
	ResourceGroupName            pulumi.StringInput
	Sku                          SnapshotSkuPtrInput
	SnapshotName                 pulumi.StringPtrInput
	SupportedCapabilities        SupportedCapabilitiesPtrInput
	SupportsHibernation          pulumi.BoolPtrInput
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
	return reflect.TypeOf((*Snapshot)(nil))
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Snapshot)(nil))
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SnapshotOutput{})
}
