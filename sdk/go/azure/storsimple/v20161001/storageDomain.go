


package v20161001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type StorageDomain struct {
	pulumi.CustomResourceState

	EncryptionKey               AsymmetricEncryptedSecretResponsePtrOutput `pulumi:"encryptionKey"`
	EncryptionStatus            pulumi.StringOutput                        `pulumi:"encryptionStatus"`
	Name                        pulumi.StringOutput                        `pulumi:"name"`
	StorageAccountCredentialIds pulumi.StringArrayOutput                   `pulumi:"storageAccountCredentialIds"`
	Type                        pulumi.StringOutput                        `pulumi:"type"`
}


func NewStorageDomain(ctx *pulumi.Context,
	name string, args *StorageDomainArgs, opts ...pulumi.ResourceOption) (*StorageDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EncryptionStatus == nil {
		return nil, errors.New("invalid value for required argument 'EncryptionStatus'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccountCredentialIds == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountCredentialIds'")
	}
	var resource StorageDomain
	err := ctx.RegisterResource("azure-native:storsimple/v20161001:StorageDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageDomainState, opts ...pulumi.ResourceOption) (*StorageDomain, error) {
	var resource StorageDomain
	err := ctx.ReadResource("azure-native:storsimple/v20161001:StorageDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type storageDomainState struct {
}

type StorageDomainState struct {
}

func (StorageDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageDomainState)(nil)).Elem()
}

type storageDomainArgs struct {
	EncryptionKey               *AsymmetricEncryptedSecret `pulumi:"encryptionKey"`
	EncryptionStatus            EncryptionStatus           `pulumi:"encryptionStatus"`
	ManagerName                 string                     `pulumi:"managerName"`
	ResourceGroupName           string                     `pulumi:"resourceGroupName"`
	StorageAccountCredentialIds []string                   `pulumi:"storageAccountCredentialIds"`
	StorageDomainName           *string                    `pulumi:"storageDomainName"`
}


type StorageDomainArgs struct {
	EncryptionKey               AsymmetricEncryptedSecretPtrInput
	EncryptionStatus            EncryptionStatusInput
	ManagerName                 pulumi.StringInput
	ResourceGroupName           pulumi.StringInput
	StorageAccountCredentialIds pulumi.StringArrayInput
	StorageDomainName           pulumi.StringPtrInput
}

func (StorageDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageDomainArgs)(nil)).Elem()
}

type StorageDomainInput interface {
	pulumi.Input

	ToStorageDomainOutput() StorageDomainOutput
	ToStorageDomainOutputWithContext(ctx context.Context) StorageDomainOutput
}

func (*StorageDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageDomain)(nil)).Elem()
}

func (i *StorageDomain) ToStorageDomainOutput() StorageDomainOutput {
	return i.ToStorageDomainOutputWithContext(context.Background())
}

func (i *StorageDomain) ToStorageDomainOutputWithContext(ctx context.Context) StorageDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageDomainOutput)
}

type StorageDomainOutput struct{ *pulumi.OutputState }

func (StorageDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageDomain)(nil)).Elem()
}

func (o StorageDomainOutput) ToStorageDomainOutput() StorageDomainOutput {
	return o
}

func (o StorageDomainOutput) ToStorageDomainOutputWithContext(ctx context.Context) StorageDomainOutput {
	return o
}

func (o StorageDomainOutput) EncryptionKey() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v *StorageDomain) AsymmetricEncryptedSecretResponsePtrOutput { return v.EncryptionKey }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

func (o StorageDomainOutput) EncryptionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageDomain) pulumi.StringOutput { return v.EncryptionStatus }).(pulumi.StringOutput)
}

func (o StorageDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageDomain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StorageDomainOutput) StorageAccountCredentialIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StorageDomain) pulumi.StringArrayOutput { return v.StorageAccountCredentialIds }).(pulumi.StringArrayOutput)
}

func (o StorageDomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageDomain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageDomainOutput{})
}
