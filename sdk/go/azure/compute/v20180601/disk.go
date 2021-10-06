


package v20180601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Disk struct {
	pulumi.CustomResourceState

	CreationData       CreationDataResponseOutput          `pulumi:"creationData"`
	DiskIOPSReadWrite  pulumi.Float64PtrOutput             `pulumi:"diskIOPSReadWrite"`
	DiskMBpsReadWrite  pulumi.IntPtrOutput                 `pulumi:"diskMBpsReadWrite"`
	DiskSizeGB         pulumi.IntPtrOutput                 `pulumi:"diskSizeGB"`
	EncryptionSettings EncryptionSettingsResponsePtrOutput `pulumi:"encryptionSettings"`
	Location           pulumi.StringOutput                 `pulumi:"location"`
	ManagedBy          pulumi.StringOutput                 `pulumi:"managedBy"`
	Name               pulumi.StringOutput                 `pulumi:"name"`
	OsType             pulumi.StringPtrOutput              `pulumi:"osType"`
	ProvisioningState  pulumi.StringOutput                 `pulumi:"provisioningState"`
	Sku                DiskSkuResponsePtrOutput            `pulumi:"sku"`
	Tags               pulumi.StringMapOutput              `pulumi:"tags"`
	TimeCreated        pulumi.StringOutput                 `pulumi:"timeCreated"`
	Type               pulumi.StringOutput                 `pulumi:"type"`
	Zones              pulumi.StringArrayOutput            `pulumi:"zones"`
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
			Type: pulumi.String("azure-nextgen:compute/v20180601:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160430preview:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20160430preview:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20170330:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180401:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180930:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180930:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190301:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190701:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191101:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20191101:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200501:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20200501:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200630:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20200630:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20200930:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20201201:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:Disk"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20210401:Disk"),
		},
	})
	opts = append(opts, aliases)
	var resource Disk
	err := ctx.RegisterResource("azure-native:compute/v20180601:Disk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskState, opts ...pulumi.ResourceOption) (*Disk, error) {
	var resource Disk
	err := ctx.ReadResource("azure-native:compute/v20180601:Disk", name, id, state, &resource, opts...)
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
	CreationData       CreationData          `pulumi:"creationData"`
	DiskIOPSReadWrite  *float64              `pulumi:"diskIOPSReadWrite"`
	DiskMBpsReadWrite  *int                  `pulumi:"diskMBpsReadWrite"`
	DiskName           *string               `pulumi:"diskName"`
	DiskSizeGB         *int                  `pulumi:"diskSizeGB"`
	EncryptionSettings *EncryptionSettings   `pulumi:"encryptionSettings"`
	Location           *string               `pulumi:"location"`
	OsType             *OperatingSystemTypes `pulumi:"osType"`
	ResourceGroupName  string                `pulumi:"resourceGroupName"`
	Sku                *DiskSku              `pulumi:"sku"`
	Tags               map[string]string     `pulumi:"tags"`
	Zones              []string              `pulumi:"zones"`
}


type DiskArgs struct {
	CreationData       CreationDataInput
	DiskIOPSReadWrite  pulumi.Float64PtrInput
	DiskMBpsReadWrite  pulumi.IntPtrInput
	DiskName           pulumi.StringPtrInput
	DiskSizeGB         pulumi.IntPtrInput
	EncryptionSettings EncryptionSettingsPtrInput
	Location           pulumi.StringPtrInput
	OsType             OperatingSystemTypesPtrInput
	ResourceGroupName  pulumi.StringInput
	Sku                DiskSkuPtrInput
	Tags               pulumi.StringMapInput
	Zones              pulumi.StringArrayInput
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
	return reflect.TypeOf((*Disk)(nil))
}

func (i *Disk) ToDiskOutput() DiskOutput {
	return i.ToDiskOutputWithContext(context.Background())
}

func (i *Disk) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskOutput)
}

type DiskOutput struct{ *pulumi.OutputState }

func (DiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Disk)(nil))
}

func (o DiskOutput) ToDiskOutput() DiskOutput {
	return o
}

func (o DiskOutput) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DiskOutput{})
}
