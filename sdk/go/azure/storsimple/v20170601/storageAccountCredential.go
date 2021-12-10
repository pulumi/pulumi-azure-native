


package v20170601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StorageAccountCredential struct {
	pulumi.CustomResourceState

	AccessKey    AsymmetricEncryptedSecretResponsePtrOutput `pulumi:"accessKey"`
	EndPoint     pulumi.StringOutput                        `pulumi:"endPoint"`
	Kind         pulumi.StringPtrOutput                     `pulumi:"kind"`
	Name         pulumi.StringOutput                        `pulumi:"name"`
	SslStatus    pulumi.StringOutput                        `pulumi:"sslStatus"`
	Type         pulumi.StringOutput                        `pulumi:"type"`
	VolumesCount pulumi.IntOutput                           `pulumi:"volumesCount"`
}


func NewStorageAccountCredential(ctx *pulumi.Context,
	name string, args *StorageAccountCredentialArgs, opts ...pulumi.ResourceOption) (*StorageAccountCredential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndPoint == nil {
		return nil, errors.New("invalid value for required argument 'EndPoint'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SslStatus == nil {
		return nil, errors.New("invalid value for required argument 'SslStatus'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storsimple:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:storsimple/v20161001:StorageAccountCredential"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageAccountCredential
	err := ctx.RegisterResource("azure-native:storsimple/v20170601:StorageAccountCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageAccountCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAccountCredentialState, opts ...pulumi.ResourceOption) (*StorageAccountCredential, error) {
	var resource StorageAccountCredential
	err := ctx.ReadResource("azure-native:storsimple/v20170601:StorageAccountCredential", name, id, state, &resource, opts...)
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
	AccessKey                    *AsymmetricEncryptedSecret `pulumi:"accessKey"`
	EndPoint                     string                     `pulumi:"endPoint"`
	Kind                         *Kind                      `pulumi:"kind"`
	ManagerName                  string                     `pulumi:"managerName"`
	ResourceGroupName            string                     `pulumi:"resourceGroupName"`
	SslStatus                    SslStatus                  `pulumi:"sslStatus"`
	StorageAccountCredentialName *string                    `pulumi:"storageAccountCredentialName"`
}


type StorageAccountCredentialArgs struct {
	AccessKey                    AsymmetricEncryptedSecretPtrInput
	EndPoint                     pulumi.StringInput
	Kind                         KindPtrInput
	ManagerName                  pulumi.StringInput
	ResourceGroupName            pulumi.StringInput
	SslStatus                    SslStatusInput
	StorageAccountCredentialName pulumi.StringPtrInput
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
