


package v20180401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Snapshot struct {
	pulumi.CustomResourceState

	CreationData       CreationDataResponseOutput          `pulumi:"creationData"`
	DiskSizeGB         pulumi.IntPtrOutput                 `pulumi:"diskSizeGB"`
	EncryptionSettings EncryptionSettingsResponsePtrOutput `pulumi:"encryptionSettings"`
	Location           pulumi.StringOutput                 `pulumi:"location"`
	ManagedBy          pulumi.StringOutput                 `pulumi:"managedBy"`
	Name               pulumi.StringOutput                 `pulumi:"name"`
	OsType             pulumi.StringPtrOutput              `pulumi:"osType"`
	ProvisioningState  pulumi.StringOutput                 `pulumi:"provisioningState"`
	Sku                SnapshotSkuResponsePtrOutput        `pulumi:"sku"`
	Tags               pulumi.StringMapOutput              `pulumi:"tags"`
	TimeCreated        pulumi.StringOutput                 `pulumi:"timeCreated"`
	Type               pulumi.StringOutput                 `pulumi:"type"`
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
			Type: pulumi.String("azure-nextgen:compute/v20180401:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160430preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20160430preview:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20170330:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180601:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180930:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180930:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190301:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190701:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191101:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20191101:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200501:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20200501:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200630:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20200630:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20200930:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20201201:Snapshot"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:Snapshot"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20210401:Snapshot"),
		},
	})
	opts = append(opts, aliases)
	var resource Snapshot
	err := ctx.RegisterResource("azure-native:compute/v20180401:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("azure-native:compute/v20180401:Snapshot", name, id, state, &resource, opts...)
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
	CreationData       CreationData          `pulumi:"creationData"`
	DiskSizeGB         *int                  `pulumi:"diskSizeGB"`
	EncryptionSettings *EncryptionSettings   `pulumi:"encryptionSettings"`
	Location           *string               `pulumi:"location"`
	OsType             *OperatingSystemTypes `pulumi:"osType"`
	ResourceGroupName  string                `pulumi:"resourceGroupName"`
	Sku                *SnapshotSku          `pulumi:"sku"`
	SnapshotName       *string               `pulumi:"snapshotName"`
	Tags               map[string]string     `pulumi:"tags"`
}


type SnapshotArgs struct {
	CreationData       CreationDataInput
	DiskSizeGB         pulumi.IntPtrInput
	EncryptionSettings EncryptionSettingsPtrInput
	Location           pulumi.StringPtrInput
	OsType             OperatingSystemTypesPtrInput
	ResourceGroupName  pulumi.StringInput
	Sku                SnapshotSkuPtrInput
	SnapshotName       pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
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
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotInput)(nil)).Elem(), &Snapshot{})
	pulumi.RegisterOutputType(SnapshotOutput{})
}
