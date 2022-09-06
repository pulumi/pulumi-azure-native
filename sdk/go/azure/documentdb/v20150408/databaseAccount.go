


package v20150408

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type DatabaseAccount struct {
	pulumi.CustomResourceState

	Capabilities                  CapabilityResponseArrayOutput         `pulumi:"capabilities"`
	ConnectorOffer                pulumi.StringPtrOutput                `pulumi:"connectorOffer"`
	ConsistencyPolicy             ConsistencyPolicyResponsePtrOutput    `pulumi:"consistencyPolicy"`
	DatabaseAccountOfferType      pulumi.StringOutput                   `pulumi:"databaseAccountOfferType"`
	DocumentEndpoint              pulumi.StringOutput                   `pulumi:"documentEndpoint"`
	EnableAutomaticFailover       pulumi.BoolPtrOutput                  `pulumi:"enableAutomaticFailover"`
	EnableCassandraConnector      pulumi.BoolPtrOutput                  `pulumi:"enableCassandraConnector"`
	EnableMultipleWriteLocations  pulumi.BoolPtrOutput                  `pulumi:"enableMultipleWriteLocations"`
	FailoverPolicies              FailoverPolicyResponseArrayOutput     `pulumi:"failoverPolicies"`
	IpRangeFilter                 pulumi.StringPtrOutput                `pulumi:"ipRangeFilter"`
	IsVirtualNetworkFilterEnabled pulumi.BoolPtrOutput                  `pulumi:"isVirtualNetworkFilterEnabled"`
	Kind                          pulumi.StringPtrOutput                `pulumi:"kind"`
	Location                      pulumi.StringPtrOutput                `pulumi:"location"`
	Name                          pulumi.StringOutput                   `pulumi:"name"`
	ProvisioningState             pulumi.StringOutput                   `pulumi:"provisioningState"`
	ReadLocations                 LocationResponseArrayOutput           `pulumi:"readLocations"`
	Tags                          pulumi.StringMapOutput                `pulumi:"tags"`
	Type                          pulumi.StringOutput                   `pulumi:"type"`
	VirtualNetworkRules           VirtualNetworkRuleResponseArrayOutput `pulumi:"virtualNetworkRules"`
	WriteLocations                LocationResponseArrayOutput           `pulumi:"writeLocations"`
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
			Type: pulumi.String("azure-native:documentdb/v20210315:DatabaseAccount"),
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
	})
	opts = append(opts, aliases)
	var resource DatabaseAccount
	err := ctx.RegisterResource("azure-native:documentdb/v20150408:DatabaseAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabaseAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountState, opts ...pulumi.ResourceOption) (*DatabaseAccount, error) {
	var resource DatabaseAccount
	err := ctx.ReadResource("azure-native:documentdb/v20150408:DatabaseAccount", name, id, state, &resource, opts...)
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
	AccountName                   *string                  `pulumi:"accountName"`
	Capabilities                  []Capability             `pulumi:"capabilities"`
	ConnectorOffer                *string                  `pulumi:"connectorOffer"`
	ConsistencyPolicy             *ConsistencyPolicy       `pulumi:"consistencyPolicy"`
	DatabaseAccountOfferType      DatabaseAccountOfferType `pulumi:"databaseAccountOfferType"`
	EnableAutomaticFailover       *bool                    `pulumi:"enableAutomaticFailover"`
	EnableCassandraConnector      *bool                    `pulumi:"enableCassandraConnector"`
	EnableMultipleWriteLocations  *bool                    `pulumi:"enableMultipleWriteLocations"`
	IpRangeFilter                 *string                  `pulumi:"ipRangeFilter"`
	IsVirtualNetworkFilterEnabled *bool                    `pulumi:"isVirtualNetworkFilterEnabled"`
	Kind                          *string                  `pulumi:"kind"`
	Location                      *string                  `pulumi:"location"`
	Locations                     []Location               `pulumi:"locations"`
	ResourceGroupName             string                   `pulumi:"resourceGroupName"`
	Tags                          map[string]string        `pulumi:"tags"`
	VirtualNetworkRules           []VirtualNetworkRule     `pulumi:"virtualNetworkRules"`
}


type DatabaseAccountArgs struct {
	AccountName                   pulumi.StringPtrInput
	Capabilities                  CapabilityArrayInput
	ConnectorOffer                pulumi.StringPtrInput
	ConsistencyPolicy             ConsistencyPolicyPtrInput
	DatabaseAccountOfferType      DatabaseAccountOfferTypeInput
	EnableAutomaticFailover       pulumi.BoolPtrInput
	EnableCassandraConnector      pulumi.BoolPtrInput
	EnableMultipleWriteLocations  pulumi.BoolPtrInput
	IpRangeFilter                 pulumi.StringPtrInput
	IsVirtualNetworkFilterEnabled pulumi.BoolPtrInput
	Kind                          pulumi.StringPtrInput
	Location                      pulumi.StringPtrInput
	Locations                     LocationArrayInput
	ResourceGroupName             pulumi.StringInput
	Tags                          pulumi.StringMapInput
	VirtualNetworkRules           VirtualNetworkRuleArrayInput
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

func (o DatabaseAccountOutput) Capabilities() CapabilityResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) CapabilityResponseArrayOutput { return v.Capabilities }).(CapabilityResponseArrayOutput)
}

func (o DatabaseAccountOutput) ConnectorOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.ConnectorOffer }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountOutput) ConsistencyPolicy() ConsistencyPolicyResponsePtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) ConsistencyPolicyResponsePtrOutput { return v.ConsistencyPolicy }).(ConsistencyPolicyResponsePtrOutput)
}

func (o DatabaseAccountOutput) DatabaseAccountOfferType() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringOutput { return v.DatabaseAccountOfferType }).(pulumi.StringOutput)
}

func (o DatabaseAccountOutput) DocumentEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringOutput { return v.DocumentEndpoint }).(pulumi.StringOutput)
}

func (o DatabaseAccountOutput) EnableAutomaticFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.EnableAutomaticFailover }).(pulumi.BoolPtrOutput)
}

func (o DatabaseAccountOutput) EnableCassandraConnector() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.EnableCassandraConnector }).(pulumi.BoolPtrOutput)
}

func (o DatabaseAccountOutput) EnableMultipleWriteLocations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.EnableMultipleWriteLocations }).(pulumi.BoolPtrOutput)
}

func (o DatabaseAccountOutput) FailoverPolicies() FailoverPolicyResponseArrayOutput {
	return o.ApplyT(func(v *DatabaseAccount) FailoverPolicyResponseArrayOutput { return v.FailoverPolicies }).(FailoverPolicyResponseArrayOutput)
}

func (o DatabaseAccountOutput) IpRangeFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.IpRangeFilter }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountOutput) IsVirtualNetworkFilterEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.BoolPtrOutput { return v.IsVirtualNetworkFilterEnabled }).(pulumi.BoolPtrOutput)
}

func (o DatabaseAccountOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DatabaseAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseAccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAccount) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
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
