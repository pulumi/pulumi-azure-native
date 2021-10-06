


package v20200501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StorageAccountCredential struct {
	pulumi.CustomResourceState

	AccountKey       AsymmetricEncryptedSecretResponsePtrOutput `pulumi:"accountKey"`
	AccountType      pulumi.StringOutput                        `pulumi:"accountType"`
	Alias            pulumi.StringOutput                        `pulumi:"alias"`
	BlobDomainName   pulumi.StringPtrOutput                     `pulumi:"blobDomainName"`
	ConnectionString pulumi.StringPtrOutput                     `pulumi:"connectionString"`
	Name             pulumi.StringOutput                        `pulumi:"name"`
	SslStatus        pulumi.StringOutput                        `pulumi:"sslStatus"`
	StorageAccountId pulumi.StringPtrOutput                     `pulumi:"storageAccountId"`
	Type             pulumi.StringOutput                        `pulumi:"type"`
	UserName         pulumi.StringPtrOutput                     `pulumi:"userName"`
}


func NewStorageAccountCredential(ctx *pulumi.Context,
	name string, args *StorageAccountCredentialArgs, opts ...pulumi.ResourceOption) (*StorageAccountCredential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountType == nil {
		return nil, errors.New("invalid value for required argument 'AccountType'")
	}
	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SslStatus == nil {
		return nil, errors.New("invalid value for required argument 'SslStatus'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20200501preview:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20190301:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20190701:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20190801:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20200901:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20200901preview:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20201201:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20210201:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20210201preview:StorageAccountCredential"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageAccountCredential
	err := ctx.RegisterResource("azure-native:databoxedge/v20200501preview:StorageAccountCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageAccountCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAccountCredentialState, opts ...pulumi.ResourceOption) (*StorageAccountCredential, error) {
	var resource StorageAccountCredential
	err := ctx.ReadResource("azure-native:databoxedge/v20200501preview:StorageAccountCredential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type storageAccountCredentialState struct {
}

type StorageAccountCredentialState struct {
}

func (StorageAccountCredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountCredentialState)(nil)).Elem()
}

type storageAccountCredentialArgs struct {
	AccountKey        *AsymmetricEncryptedSecret `pulumi:"accountKey"`
	AccountType       string                     `pulumi:"accountType"`
	Alias             string                     `pulumi:"alias"`
	BlobDomainName    *string                    `pulumi:"blobDomainName"`
	ConnectionString  *string                    `pulumi:"connectionString"`
	DeviceName        string                     `pulumi:"deviceName"`
	Name              *string                    `pulumi:"name"`
	ResourceGroupName string                     `pulumi:"resourceGroupName"`
	SslStatus         string                     `pulumi:"sslStatus"`
	StorageAccountId  *string                    `pulumi:"storageAccountId"`
	UserName          *string                    `pulumi:"userName"`
}


type StorageAccountCredentialArgs struct {
	AccountKey        AsymmetricEncryptedSecretPtrInput
	AccountType       pulumi.StringInput
	Alias             pulumi.StringInput
	BlobDomainName    pulumi.StringPtrInput
	ConnectionString  pulumi.StringPtrInput
	DeviceName        pulumi.StringInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	SslStatus         pulumi.StringInput
	StorageAccountId  pulumi.StringPtrInput
	UserName          pulumi.StringPtrInput
}

func (StorageAccountCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountCredentialArgs)(nil)).Elem()
}

type StorageAccountCredentialInput interface {
	pulumi.Input

	ToStorageAccountCredentialOutput() StorageAccountCredentialOutput
	ToStorageAccountCredentialOutputWithContext(ctx context.Context) StorageAccountCredentialOutput
}

func (*StorageAccountCredential) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountCredential)(nil))
}

func (i *StorageAccountCredential) ToStorageAccountCredentialOutput() StorageAccountCredentialOutput {
	return i.ToStorageAccountCredentialOutputWithContext(context.Background())
}

func (i *StorageAccountCredential) ToStorageAccountCredentialOutputWithContext(ctx context.Context) StorageAccountCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountCredentialOutput)
}

type StorageAccountCredentialOutput struct{ *pulumi.OutputState }

func (StorageAccountCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountCredential)(nil))
}

func (o StorageAccountCredentialOutput) ToStorageAccountCredentialOutput() StorageAccountCredentialOutput {
	return o
}

func (o StorageAccountCredentialOutput) ToStorageAccountCredentialOutputWithContext(ctx context.Context) StorageAccountCredentialOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StorageAccountCredentialOutput{})
}
