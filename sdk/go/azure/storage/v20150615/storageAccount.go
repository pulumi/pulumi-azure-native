


package v20150615

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type StorageAccount struct {
	pulumi.CustomResourceState

	AccountType         pulumi.StringPtrOutput        `pulumi:"accountType"`
	CreationTime        pulumi.StringPtrOutput        `pulumi:"creationTime"`
	CustomDomain        CustomDomainResponsePtrOutput `pulumi:"customDomain"`
	LastGeoFailoverTime pulumi.StringPtrOutput        `pulumi:"lastGeoFailoverTime"`
	Location            pulumi.StringPtrOutput        `pulumi:"location"`
	Name                pulumi.StringOutput           `pulumi:"name"`
	PrimaryEndpoints    EndpointsResponsePtrOutput    `pulumi:"primaryEndpoints"`
	PrimaryLocation     pulumi.StringPtrOutput        `pulumi:"primaryLocation"`
	ProvisioningState   pulumi.StringPtrOutput        `pulumi:"provisioningState"`
	SecondaryEndpoints  EndpointsResponsePtrOutput    `pulumi:"secondaryEndpoints"`
	SecondaryLocation   pulumi.StringPtrOutput        `pulumi:"secondaryLocation"`
	StatusOfPrimary     pulumi.StringPtrOutput        `pulumi:"statusOfPrimary"`
	StatusOfSecondary   pulumi.StringPtrOutput        `pulumi:"statusOfSecondary"`
	Tags                pulumi.StringMapOutput        `pulumi:"tags"`
	Type                pulumi.StringOutput           `pulumi:"type"`
}


func NewStorageAccount(ctx *pulumi.Context,
	name string, args *StorageAccountArgs, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountType == nil {
		return nil, errors.New("invalid value for required argument 'AccountType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20150501preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20160101:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20160501:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20161201:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20170601:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20171001:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180201:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180301preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180701:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20181101:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190401:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220901:StorageAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageAccount
	err := ctx.RegisterResource("azure-native:storage/v20150615:StorageAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAccountState, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	var resource StorageAccount
	err := ctx.ReadResource("azure-native:storage/v20150615:StorageAccount", name, id, state, &resource, opts...)
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
	AccountName       *string           `pulumi:"accountName"`
	AccountType       AccountType       `pulumi:"accountType"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type StorageAccountArgs struct {
	AccountName       pulumi.StringPtrInput
	AccountType       AccountTypeInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
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
	return reflect.TypeOf((**StorageAccount)(nil)).Elem()
}

func (i *StorageAccount) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i *StorageAccount) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}

type StorageAccountOutput struct{ *pulumi.OutputState }

func (StorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccount)(nil)).Elem()
}

func (o StorageAccountOutput) ToStorageAccountOutput() StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) AccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.AccountType }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) CustomDomain() CustomDomainResponsePtrOutput {
	return o.ApplyT(func(v *StorageAccount) CustomDomainResponsePtrOutput { return v.CustomDomain }).(CustomDomainResponsePtrOutput)
}

func (o StorageAccountOutput) LastGeoFailoverTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.LastGeoFailoverTime }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) PrimaryEndpoints() EndpointsResponsePtrOutput {
	return o.ApplyT(func(v *StorageAccount) EndpointsResponsePtrOutput { return v.PrimaryEndpoints }).(EndpointsResponsePtrOutput)
}

func (o StorageAccountOutput) PrimaryLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.PrimaryLocation }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) SecondaryEndpoints() EndpointsResponsePtrOutput {
	return o.ApplyT(func(v *StorageAccount) EndpointsResponsePtrOutput { return v.SecondaryEndpoints }).(EndpointsResponsePtrOutput)
}

func (o StorageAccountOutput) SecondaryLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.SecondaryLocation }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) StatusOfPrimary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.StatusOfPrimary }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) StatusOfSecondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.StatusOfSecondary }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o StorageAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageAccountOutput{})
}
