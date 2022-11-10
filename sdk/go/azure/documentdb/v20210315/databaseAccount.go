


package v20210315

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabaseAccount struct {
	pulumi.CustomResourceState

	ApiProperties                      ApiPropertiesResponsePtrOutput               `pulumi:"apiProperties"`
	BackupPolicy                       pulumi.AnyOutput                             `pulumi:"backupPolicy"`
	Capabilities                       CapabilityResponseArrayOutput                `pulumi:"capabilities"`
	ConnectorOffer                     pulumi.StringPtrOutput                       `pulumi:"connectorOffer"`
	ConsistencyPolicy                  ConsistencyPolicyResponsePtrOutput           `pulumi:"consistencyPolicy"`
	Cors                               CorsPolicyResponseArrayOutput                `pulumi:"cors"`
	DatabaseAccountOfferType           pulumi.StringOutput                          `pulumi:"databaseAccountOfferType"`
	DefaultIdentity                    pulumi.StringPtrOutput                       `pulumi:"defaultIdentity"`
	DisableKeyBasedMetadataWriteAccess pulumi.BoolPtrOutput                         `pulumi:"disableKeyBasedMetadataWriteAccess"`
	DocumentEndpoint                   pulumi.StringOutput                          `pulumi:"documentEndpoint"`
	EnableAnalyticalStorage            pulumi.BoolPtrOutput                         `pulumi:"enableAnalyticalStorage"`
	EnableAutomaticFailover            pulumi.BoolPtrOutput                         `pulumi:"enableAutomaticFailover"`
	EnableCassandraConnector           pulumi.BoolPtrOutput                         `pulumi:"enableCassandraConnector"`
	EnableFreeTier                     pulumi.BoolPtrOutput                         `pulumi:"enableFreeTier"`
	EnableMultipleWriteLocations       pulumi.BoolPtrOutput                         `pulumi:"enableMultipleWriteLocations"`
	FailoverPolicies                   FailoverPolicyResponseArrayOutput            `pulumi:"failoverPolicies"`
	Identity                           ManagedServiceIdentityResponsePtrOutput      `pulumi:"identity"`
	IpRules                            IpAddressOrRangeResponseArrayOutput          `pulumi:"ipRules"`
	IsVirtualNetworkFilterEnabled      pulumi.BoolPtrOutput                         `pulumi:"isVirtualNetworkFilterEnabled"`
	KeyVaultKeyUri                     pulumi.StringPtrOutput                       `pulumi:"keyVaultKeyUri"`
	Kind                               pulumi.StringPtrOutput                       `pulumi:"kind"`
	Location                           pulumi.StringPtrOutput                       `pulumi:"location"`
	Locations                          LocationResponseArrayOutput                  `pulumi:"locations"`
	Name                               pulumi.StringOutput                          `pulumi:"name"`
	NetworkAclBypass                   pulumi.StringPtrOutput                       `pulumi:"networkAclBypass"`
	NetworkAclBypassResourceIds        pulumi.StringArrayOutput                     `pulumi:"networkAclBypassResourceIds"`
	PrivateEndpointConnections         PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	ProvisioningState                  pulumi.StringOutput                          `pulumi:"provisioningState"`
	PublicNetworkAccess                pulumi.StringPtrOutput                       `pulumi:"publicNetworkAccess"`
	ReadLocations                      LocationResponseArrayOutput                  `pulumi:"readLocations"`
	Tags                               pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                               pulumi.StringOutput                          `pulumi:"type"`
	VirtualNetworkRules                VirtualNetworkRuleResponseArrayOutput        `pulumi:"virtualNetworkRules"`
	WriteLocations                     LocationResponseArrayOutput                  `pulumi:"writeLocations"`
}


func NewDatabaseAccount(ctx *pulumi.Context,
	name string, args *DatabaseAccountArgs, opts ...pulumi.ResourceOption) (*DatabaseAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseAccountOfferType == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseAccountOfferType'")
	}
	if args.Locations == nil {
		return nil, errors.New("invalid value for required argument 'Locations'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.Kind) {
		args.Kind = pulumi.StringPtr("GlobalDocumentDB")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:DatabaseAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccount
	err := ctx.RegisterResource("azure-native:documentdb/v20210315:DatabaseAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabaseAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountState, opts ...pulumi.ResourceOption) (*DatabaseAccount, error) {
	var resource DatabaseAccount
	err := ctx.ReadResource("azure-native:documentdb/v20210315:DatabaseAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databaseAccountState struct {
}

type DatabaseAccountState struct {
}

func (DatabaseAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountState)(nil)).Elem()
}

type databaseAccountArgs struct {
	AccountName                        *string                  `pulumi:"accountName"`
	ApiProperties                      *ApiProperties           `pulumi:"apiProperties"`
	BackupPolicy                       interface{}              `pulumi:"backupPolicy"`
	Capabilities                       []Capability             `pulumi:"capabilities"`
	ConnectorOffer                     *string                  `pulumi:"connectorOffer"`
	ConsistencyPolicy                  *ConsistencyPolicy       `pulumi:"consistencyPolicy"`
	Cors                               []CorsPolicy             `pulumi:"cors"`
	DatabaseAccountOfferType           DatabaseAccountOfferType `pulumi:"databaseAccountOfferType"`
	DefaultIdentity                    *string                  `pulumi:"defaultIdentity"`
	DisableKeyBasedMetadataWriteAccess *bool                    `pulumi:"disableKeyBasedMetadataWriteAccess"`
	EnableAnalyticalStorage            *bool                    `pulumi:"enableAnalyticalStorage"`
	EnableAutomaticFailover            *bool                    `pulumi:"enableAutomaticFailover"`
	EnableCassandraConnector           *bool                    `pulumi:"enableCassandraConnector"`
	EnableFreeTier                     *bool                    `pulumi:"enableFreeTier"`
	EnableMultipleWriteLocations       *bool                    `pulumi:"enableMultipleWriteLocations"`
	Identity                           *ManagedServiceIdentity  `pulumi:"identity"`
	IpRules                            []IpAddressOrRange       `pulumi:"ipRules"`
	IsVirtualNetworkFilterEnabled      *bool                    `pulumi:"isVirtualNetworkFilterEnabled"`
	KeyVaultKeyUri                     *string                  `pulumi:"keyVaultKeyUri"`
	Kind                               *string                  `pulumi:"kind"`
	Location                           *string                  `pulumi:"location"`
	Locations                          []Location               `pulumi:"locations"`
	NetworkAclBypass                   *NetworkAclBypass        `pulumi:"networkAclBypass"`
	NetworkAclBypassResourceIds        []string                 `pulumi:"networkAclBypassResourceIds"`
	PublicNetworkAccess                *string                  `pulumi:"publicNetworkAccess"`
	ResourceGroupName                  string                   `pulumi:"resourceGroupName"`
	Tags                               map[string]string        `pulumi:"tags"`
	VirtualNetworkRules                []VirtualNetworkRule     `pulumi:"virtualNetworkRules"`
}


type DatabaseAccountArgs struct {
	AccountName                        pulumi.StringPtrInput
	ApiProperties                      ApiPropertiesPtrInput
	BackupPolicy                       pulumi.Input
	Capabilities                       CapabilityArrayInput
	ConnectorOffer                     pulumi.StringPtrInput
	ConsistencyPolicy                  ConsistencyPolicyPtrInput
	Cors                               CorsPolicyArrayInput
	DatabaseAccountOfferType           DatabaseAccountOfferTypeInput
	DefaultIdentity                    pulumi.StringPtrInput
	DisableKeyBasedMetadataWriteAccess pulumi.BoolPtrInput
	EnableAnalyticalStorage            pulumi.BoolPtrInput
	EnableAutomaticFailover            pulumi.BoolPtrInput
	EnableCassandraConnector           pulumi.BoolPtrInput
	EnableFreeTier                     pulumi.BoolPtrInput
	EnableMultipleWriteLocations       pulumi.BoolPtrInput
	Identity                           ManagedServiceIdentityPtrInput
	IpRules                            IpAddressOrRangeArrayInput
	IsVirtualNetworkFilterEnabled      pulumi.BoolPtrInput
	KeyVaultKeyUri                     pulumi.StringPtrInput
	Kind                               pulumi.StringPtrInput
	Location                           pulumi.StringPtrInput
	Locations                          LocationArrayInput
	NetworkAclBypass                   NetworkAclBypassPtrInput
	NetworkAclBypassResourceIds        pulumi.StringArrayInput
	PublicNetworkAccess                pulumi.StringPtrInput
	ResourceGroupName                  pulumi.StringInput
	Tags                               pulumi.StringMapInput
	VirtualNetworkRules                VirtualNetworkRuleArrayInput
}

func (DatabaseAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountArgs)(nil)).Elem()
}

type DatabaseAccountInput interface {
	pulumi.Input

	ToDatabaseAccountOutput() DatabaseAccountOutput
	ToDatabaseAccountOutputWithContext(ctx context.Context) DatabaseAccountOutput
}

func (*DatabaseAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccount)(nil)).Elem()
}

func (i *DatabaseAccount) ToDatabaseAccountOutput() DatabaseAccountOutput {
	return i.ToDatabaseAccountOutputWithContext(context.Background())
}

func (i *DatabaseAccount) ToDatabaseAccountOutputWithContext(ctx context.Context) DatabaseAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountOutput)
}

type DatabaseAccountOutput struct{ *pulumi.OutputState }

func (DatabaseAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccount)(nil)).Elem()
}

func (o DatabaseAccountOutput) ToDatabaseAccountOutput() DatabaseAccountOutput {
	return o
}

func (o DatabaseAccountOutput) ToDatabaseAccountOutputWithContext(ctx context.Context) DatabaseAccountOutput {
	return o
}

func (o DatabaseAccountOutput) ApiProperties() ApiPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) ApiPropertiesResponsePtrOutput { return v.ApiProperties }).(ApiPropertiesResponsePtrOutput)
}

func (o DatabaseAccountOutput) BackupPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.AnyOutput { return v.BackupPolicy }).(pulumi.AnyOutput)
}

func (o DatabaseAccountOutput) Capabilities() CapabilityResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) CapabilityResponseArrayOutput { return v.Capabilities }).(CapabilityResponseArrayOutput)
}

func (o DatabaseAccountOutput) ConnectorOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.ConnectorOffer }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountOutput) ConsistencyPolicy() ConsistencyPolicyResponsePtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) ConsistencyPolicyResponsePtrOutput { return v.ConsistencyPolicy }).(ConsistencyPolicyResponsePtrOutput)
}

func (o DatabaseAccountOutput) Cors() CorsPolicyResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) CorsPolicyResponseArrayOutput { return v.Cors }).(CorsPolicyResponseArrayOutput)
}

func (o DatabaseAccountOutput) DatabaseAccountOfferType() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringOutput { return v.DatabaseAccountOfferType }).(pulumi.StringOutput)
}

func (o DatabaseAccountOutput) DefaultIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.DefaultIdentity }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountOutput) DisableKeyBasedMetadataWriteAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.DisableKeyBasedMetadataWriteAccess }).(pulumi.BoolPtrOutput)
}

func (o DatabaseAccountOutput) DocumentEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringOutput { return v.DocumentEndpoint }).(pulumi.StringOutput)
}

func (o DatabaseAccountOutput) EnableAnalyticalStorage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.EnableAnalyticalStorage }).(pulumi.BoolPtrOutput)
}

func (o DatabaseAccountOutput) EnableAutomaticFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.EnableAutomaticFailover }).(pulumi.BoolPtrOutput)
}

func (o DatabaseAccountOutput) EnableCassandraConnector() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.EnableCassandraConnector }).(pulumi.BoolPtrOutput)
}

func (o DatabaseAccountOutput) EnableFreeTier() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.EnableFreeTier }).(pulumi.BoolPtrOutput)
}

func (o DatabaseAccountOutput) EnableMultipleWriteLocations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.EnableMultipleWriteLocations }).(pulumi.BoolPtrOutput)
}

func (o DatabaseAccountOutput) FailoverPolicies() FailoverPolicyResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) FailoverPolicyResponseArrayOutput { return v.FailoverPolicies }).(FailoverPolicyResponseArrayOutput)
}

func (o DatabaseAccountOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o DatabaseAccountOutput) IpRules() IpAddressOrRangeResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) IpAddressOrRangeResponseArrayOutput { return v.IpRules }).(IpAddressOrRangeResponseArrayOutput)
}

func (o DatabaseAccountOutput) IsVirtualNetworkFilterEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.IsVirtualNetworkFilterEnabled }).(pulumi.BoolPtrOutput)
}

func (o DatabaseAccountOutput) KeyVaultKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.KeyVaultKeyUri }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountOutput) Locations() LocationResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) LocationResponseArrayOutput { return v.Locations }).(LocationResponseArrayOutput)
}

func (o DatabaseAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseAccountOutput) NetworkAclBypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.NetworkAclBypass }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountOutput) NetworkAclBypassResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringArrayOutput { return v.NetworkAclBypassResourceIds }).(pulumi.StringArrayOutput)
}

func (o DatabaseAccountOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o DatabaseAccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DatabaseAccountOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountOutput) ReadLocations() LocationResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) LocationResponseArrayOutput { return v.ReadLocations }).(LocationResponseArrayOutput)
}

func (o DatabaseAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DatabaseAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o DatabaseAccountOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) VirtualNetworkRuleResponseArrayOutput { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

func (o DatabaseAccountOutput) WriteLocations() LocationResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) LocationResponseArrayOutput { return v.WriteLocations }).(LocationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseAccountOutput{})
}
