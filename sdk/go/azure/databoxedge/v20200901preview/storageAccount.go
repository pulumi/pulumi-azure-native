


package v20200901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StorageAccount struct {
	pulumi.CustomResourceState

	BlobEndpoint               pulumi.StringOutput      `pulumi:"blobEndpoint"`
	ContainerCount             pulumi.IntOutput         `pulumi:"containerCount"`
	DataPolicy                 pulumi.StringOutput      `pulumi:"dataPolicy"`
	Description                pulumi.StringPtrOutput   `pulumi:"description"`
	Name                       pulumi.StringOutput      `pulumi:"name"`
	StorageAccountCredentialId pulumi.StringPtrOutput   `pulumi:"storageAccountCredentialId"`
	StorageAccountStatus       pulumi.StringPtrOutput   `pulumi:"storageAccountStatus"`
	SystemData                 SystemDataResponseOutput `pulumi:"systemData"`
	Type                       pulumi.StringOutput      `pulumi:"type"`
}


func NewStorageAccount(ctx *pulumi.Context,
	name string, args *StorageAccountArgs, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataPolicy == nil {
		return nil, errors.New("invalid value for required argument 'DataPolicy'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20200901preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20190801:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20200501preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20200901:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20201201:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20210201:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20210201preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20210601:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20210601preview:StorageAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageAccount
	err := ctx.RegisterResource("azure-native:databoxedge/v20200901preview:StorageAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAccountState, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	var resource StorageAccount
	err := ctx.ReadResource("azure-native:databoxedge/v20200901preview:StorageAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type storageAccountState struct {
}

type StorageAccountState struct {
}

func (StorageAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountState)(nil)).Elem()
}

type storageAccountArgs struct {
	DataPolicy                 string  `pulumi:"dataPolicy"`
	Description                *string `pulumi:"description"`
	DeviceName                 string  `pulumi:"deviceName"`
	ResourceGroupName          string  `pulumi:"resourceGroupName"`
	StorageAccountCredentialId *string `pulumi:"storageAccountCredentialId"`
	StorageAccountName         *string `pulumi:"storageAccountName"`
	StorageAccountStatus       *string `pulumi:"storageAccountStatus"`
}


type StorageAccountArgs struct {
	DataPolicy                 pulumi.StringInput
	Description                pulumi.StringPtrInput
	DeviceName                 pulumi.StringInput
	ResourceGroupName          pulumi.StringInput
	StorageAccountCredentialId pulumi.StringPtrInput
	StorageAccountName         pulumi.StringPtrInput
	StorageAccountStatus       pulumi.StringPtrInput
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountArgs)(nil)).Elem()
}

type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput
}

func (*StorageAccount) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil))
}

func (i *StorageAccount) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i *StorageAccount) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}

type StorageAccountOutput struct{ *pulumi.OutputState }

func (StorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil))
}

func (o StorageAccountOutput) ToStorageAccountOutput() StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountInput)(nil)).Elem(), &StorageAccount{})
	pulumi.RegisterOutputType(StorageAccountOutput{})
}
