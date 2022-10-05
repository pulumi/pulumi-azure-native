


package v20210201preview

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
	SystemData       SystemDataResponseOutput                   `pulumi:"systemData"`
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
			Type: pulumi.String("azure-native:databoxedge:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:StorageAccountCredential"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageAccountCredential
	err := ctx.RegisterResource("azure-native:databoxedge/v20210201preview:StorageAccountCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageAccountCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAccountCredentialState, opts ...pulumi.ResourceOption) (*StorageAccountCredential, error) {
	var resource StorageAccountCredential
	err := ctx.ReadResource("azure-native:databoxedge/v20210201preview:StorageAccountCredential", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((**StorageAccountCredential)(nil)).Elem()
}

func (i *StorageAccountCredential) ToStorageAccountCredentialOutput() StorageAccountCredentialOutput {
	return i.ToStorageAccountCredentialOutputWithContext(context.Background())
}

func (i *StorageAccountCredential) ToStorageAccountCredentialOutputWithContext(ctx context.Context) StorageAccountCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountCredentialOutput)
}

type StorageAccountCredentialOutput struct{ *pulumi.OutputState }

func (StorageAccountCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountCredential)(nil)).Elem()
}

func (o StorageAccountCredentialOutput) ToStorageAccountCredentialOutput() StorageAccountCredentialOutput {
	return o
}

func (o StorageAccountCredentialOutput) ToStorageAccountCredentialOutputWithContext(ctx context.Context) StorageAccountCredentialOutput {
	return o
}

func (o StorageAccountCredentialOutput) AccountKey() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v *StorageAccountCredential) AsymmetricEncryptedSecretResponsePtrOutput { return v.AccountKey }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

func (o StorageAccountCredentialOutput) AccountType() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringOutput { return v.AccountType }).(pulumi.StringOutput)
}

func (o StorageAccountCredentialOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

func (o StorageAccountCredentialOutput) BlobDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringPtrOutput { return v.BlobDomainName }).(pulumi.StringPtrOutput)
}

func (o StorageAccountCredentialOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringPtrOutput { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o StorageAccountCredentialOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StorageAccountCredentialOutput) SslStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringOutput { return v.SslStatus }).(pulumi.StringOutput)
}

func (o StorageAccountCredentialOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringPtrOutput { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

func (o StorageAccountCredentialOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StorageAccountCredential) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o StorageAccountCredentialOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o StorageAccountCredentialOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringPtrOutput { return v.UserName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageAccountCredentialOutput{})
}
