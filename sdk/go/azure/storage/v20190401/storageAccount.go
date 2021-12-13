


package v20190401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StorageAccount struct {
	pulumi.CustomResourceState

	AccessTier                            pulumi.StringOutput                                    `pulumi:"accessTier"`
	AllowBlobPublicAccess                 pulumi.BoolPtrOutput                                   `pulumi:"allowBlobPublicAccess"`
	AllowSharedKeyAccess                  pulumi.BoolPtrOutput                                   `pulumi:"allowSharedKeyAccess"`
	AzureFilesIdentityBasedAuthentication AzureFilesIdentityBasedAuthenticationResponsePtrOutput `pulumi:"azureFilesIdentityBasedAuthentication"`
	CreationTime                          pulumi.StringOutput                                    `pulumi:"creationTime"`
	CustomDomain                          CustomDomainResponseOutput                             `pulumi:"customDomain"`
	EnableHttpsTrafficOnly                pulumi.BoolPtrOutput                                   `pulumi:"enableHttpsTrafficOnly"`
	Encryption                            EncryptionResponseOutput                               `pulumi:"encryption"`
	FailoverInProgress                    pulumi.BoolOutput                                      `pulumi:"failoverInProgress"`
	GeoReplicationStats                   GeoReplicationStatsResponseOutput                      `pulumi:"geoReplicationStats"`
	Identity                              IdentityResponsePtrOutput                              `pulumi:"identity"`
	IsHnsEnabled                          pulumi.BoolPtrOutput                                   `pulumi:"isHnsEnabled"`
	Kind                                  pulumi.StringOutput                                    `pulumi:"kind"`
	LargeFileSharesState                  pulumi.StringPtrOutput                                 `pulumi:"largeFileSharesState"`
	LastGeoFailoverTime                   pulumi.StringOutput                                    `pulumi:"lastGeoFailoverTime"`
	Location                              pulumi.StringOutput                                    `pulumi:"location"`
	MinimumTlsVersion                     pulumi.StringPtrOutput                                 `pulumi:"minimumTlsVersion"`
	Name                                  pulumi.StringOutput                                    `pulumi:"name"`
	NetworkRuleSet                        NetworkRuleSetResponseOutput                           `pulumi:"networkRuleSet"`
	PrimaryEndpoints                      EndpointsResponseOutput                                `pulumi:"primaryEndpoints"`
	PrimaryLocation                       pulumi.StringOutput                                    `pulumi:"primaryLocation"`
	ProvisioningState                     pulumi.StringOutput                                    `pulumi:"provisioningState"`
	SecondaryEndpoints                    EndpointsResponseOutput                                `pulumi:"secondaryEndpoints"`
	SecondaryLocation                     pulumi.StringOutput                                    `pulumi:"secondaryLocation"`
	Sku                                   SkuResponseOutput                                      `pulumi:"sku"`
	StatusOfPrimary                       pulumi.StringOutput                                    `pulumi:"statusOfPrimary"`
	StatusOfSecondary                     pulumi.StringOutput                                    `pulumi:"statusOfSecondary"`
	Tags                                  pulumi.StringMapOutput                                 `pulumi:"tags"`
	Type                                  pulumi.StringOutput                                    `pulumi:"type"`
}


func NewStorageAccount(ctx *pulumi.Context,
	name string, args *StorageAccountArgs, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.Encryption != nil {
		args.Encryption = args.Encryption.ToEncryptionPtrOutput().ApplyT(func(v *Encryption) *Encryption { return v.Defaults() }).(EncryptionPtrOutput)
	}
	if args.NetworkRuleSet != nil {
		args.NetworkRuleSet = args.NetworkRuleSet.ToNetworkRuleSetPtrOutput().ApplyT(func(v *NetworkRuleSet) *NetworkRuleSet { return v.Defaults() }).(NetworkRuleSetPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20150501preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20150615:StorageAccount"),
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
	})
	opts = append(opts, aliases)
	var resource StorageAccount
	err := ctx.RegisterResource("azure-native:storage/v20190401:StorageAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAccountState, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	var resource StorageAccount
	err := ctx.ReadResource("azure-native:storage/v20190401:StorageAccount", name, id, state, &resource, opts...)
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
	AccessTier                            *AccessTier                            `pulumi:"accessTier"`
	AccountName                           *string                                `pulumi:"accountName"`
	AllowBlobPublicAccess                 *bool                                  `pulumi:"allowBlobPublicAccess"`
	AllowSharedKeyAccess                  *bool                                  `pulumi:"allowSharedKeyAccess"`
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthentication `pulumi:"azureFilesIdentityBasedAuthentication"`
	CustomDomain                          *CustomDomain                          `pulumi:"customDomain"`
	EnableHttpsTrafficOnly                *bool                                  `pulumi:"enableHttpsTrafficOnly"`
	Encryption                            *Encryption                            `pulumi:"encryption"`
	Identity                              *Identity                              `pulumi:"identity"`
	IsHnsEnabled                          *bool                                  `pulumi:"isHnsEnabled"`
	Kind                                  string                                 `pulumi:"kind"`
	LargeFileSharesState                  *string                                `pulumi:"largeFileSharesState"`
	Location                              *string                                `pulumi:"location"`
	MinimumTlsVersion                     *string                                `pulumi:"minimumTlsVersion"`
	NetworkRuleSet                        *NetworkRuleSet                        `pulumi:"networkRuleSet"`
	ResourceGroupName                     string                                 `pulumi:"resourceGroupName"`
	Sku                                   Sku                                    `pulumi:"sku"`
	Tags                                  map[string]string                      `pulumi:"tags"`
}


type StorageAccountArgs struct {
	AccessTier                            AccessTierPtrInput
	AccountName                           pulumi.StringPtrInput
	AllowBlobPublicAccess                 pulumi.BoolPtrInput
	AllowSharedKeyAccess                  pulumi.BoolPtrInput
	AzureFilesIdentityBasedAuthentication AzureFilesIdentityBasedAuthenticationPtrInput
	CustomDomain                          CustomDomainPtrInput
	EnableHttpsTrafficOnly                pulumi.BoolPtrInput
	Encryption                            EncryptionPtrInput
	Identity                              IdentityPtrInput
	IsHnsEnabled                          pulumi.BoolPtrInput
	Kind                                  pulumi.StringInput
	LargeFileSharesState                  pulumi.StringPtrInput
	Location                              pulumi.StringPtrInput
	MinimumTlsVersion                     pulumi.StringPtrInput
	NetworkRuleSet                        NetworkRuleSetPtrInput
	ResourceGroupName                     pulumi.StringInput
	Sku                                   SkuInput
	Tags                                  pulumi.StringMapInput
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

func init() {
	pulumi.RegisterOutputType(StorageAccountOutput{})
}
