


package v20161001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StorageAccountCredential struct {
	pulumi.CustomResourceState

	AccessKey AsymmetricEncryptedSecretResponsePtrOutput `pulumi:"accessKey"`
	CloudType pulumi.StringOutput                        `pulumi:"cloudType"`
	EnableSSL pulumi.StringOutput                        `pulumi:"enableSSL"`
	EndPoint  pulumi.StringOutput                        `pulumi:"endPoint"`
	Location  pulumi.StringPtrOutput                     `pulumi:"location"`
	Login     pulumi.StringOutput                        `pulumi:"login"`
	Name      pulumi.StringOutput                        `pulumi:"name"`
	Type      pulumi.StringOutput                        `pulumi:"type"`
}


func NewStorageAccountCredential(ctx *pulumi.Context,
	name string, args *StorageAccountCredentialArgs, opts ...pulumi.ResourceOption) (*StorageAccountCredential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudType == nil {
		return nil, errors.New("invalid value for required argument 'CloudType'")
	}
	if args.EnableSSL == nil {
		return nil, errors.New("invalid value for required argument 'EnableSSL'")
	}
	if args.EndPoint == nil {
		return nil, errors.New("invalid value for required argument 'EndPoint'")
	}
	if args.Login == nil {
		return nil, errors.New("invalid value for required argument 'Login'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:storsimple/v20161001:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:storsimple:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-nextgen:storsimple:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:storsimple/v20170601:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-nextgen:storsimple/v20170601:StorageAccountCredential"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageAccountCredential
	err := ctx.RegisterResource("azure-native:storsimple/v20161001:StorageAccountCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageAccountCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAccountCredentialState, opts ...pulumi.ResourceOption) (*StorageAccountCredential, error) {
	var resource StorageAccountCredential
	err := ctx.ReadResource("azure-native:storsimple/v20161001:StorageAccountCredential", name, id, state, &resource, opts...)
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
	AccessKey         *AsymmetricEncryptedSecret `pulumi:"accessKey"`
	CloudType         CloudType                  `pulumi:"cloudType"`
	CredentialName    *string                    `pulumi:"credentialName"`
	EnableSSL         SslStatus                  `pulumi:"enableSSL"`
	EndPoint          string                     `pulumi:"endPoint"`
	Location          *string                    `pulumi:"location"`
	Login             string                     `pulumi:"login"`
	ManagerName       string                     `pulumi:"managerName"`
	ResourceGroupName string                     `pulumi:"resourceGroupName"`
}


type StorageAccountCredentialArgs struct {
	AccessKey         AsymmetricEncryptedSecretPtrInput
	CloudType         CloudTypeInput
	CredentialName    pulumi.StringPtrInput
	EnableSSL         SslStatusInput
	EndPoint          pulumi.StringInput
	Location          pulumi.StringPtrInput
	Login             pulumi.StringInput
	ManagerName       pulumi.StringInput
	ResourceGroupName pulumi.StringInput
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
