


package v20211201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Domain struct {
	pulumi.CustomResourceState

	AutoCreateTopicWithFirstSubscription pulumi.BoolPtrOutput                         `pulumi:"autoCreateTopicWithFirstSubscription"`
	AutoDeleteTopicWithLastSubscription  pulumi.BoolPtrOutput                         `pulumi:"autoDeleteTopicWithLastSubscription"`
	DisableLocalAuth                     pulumi.BoolPtrOutput                         `pulumi:"disableLocalAuth"`
	Endpoint                             pulumi.StringOutput                          `pulumi:"endpoint"`
	Identity                             IdentityInfoResponsePtrOutput                `pulumi:"identity"`
	InboundIpRules                       InboundIpRuleResponseArrayOutput             `pulumi:"inboundIpRules"`
	InputSchema                          pulumi.StringPtrOutput                       `pulumi:"inputSchema"`
	InputSchemaMapping                   JsonInputSchemaMappingResponsePtrOutput      `pulumi:"inputSchemaMapping"`
	Location                             pulumi.StringOutput                          `pulumi:"location"`
	MetricResourceId                     pulumi.StringOutput                          `pulumi:"metricResourceId"`
	Name                                 pulumi.StringOutput                          `pulumi:"name"`
	PrivateEndpointConnections           PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	ProvisioningState                    pulumi.StringOutput                          `pulumi:"provisioningState"`
	PublicNetworkAccess                  pulumi.StringPtrOutput                       `pulumi:"publicNetworkAccess"`
	SystemData                           SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags                                 pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                                 pulumi.StringOutput                          `pulumi:"type"`
}


func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.AutoCreateTopicWithFirstSubscription) {
		args.AutoCreateTopicWithFirstSubscription = pulumi.BoolPtr(true)
	}
	if isZero(args.AutoDeleteTopicWithLastSubscription) {
		args.AutoDeleteTopicWithLastSubscription = pulumi.BoolPtr(true)
	}
	if isZero(args.DisableLocalAuth) {
		args.DisableLocalAuth = pulumi.BoolPtr(false)
	}
	if isZero(args.InputSchema) {
		args.InputSchema = pulumi.StringPtr("EventGridSchema")
	}
	if isZero(args.PublicNetworkAccess) {
		args.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180915preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190201preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190601:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200101preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200601:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20210601preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20220615:Domain"),
		},
	})
	opts = append(opts, aliases)
	var resource Domain
	err := ctx.RegisterResource("azure-native:eventgrid/v20211201:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("azure-native:eventgrid/v20211201:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type domainState struct {
}

type DomainState struct {
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	AutoCreateTopicWithFirstSubscription *bool                   `pulumi:"autoCreateTopicWithFirstSubscription"`
	AutoDeleteTopicWithLastSubscription  *bool                   `pulumi:"autoDeleteTopicWithLastSubscription"`
	DisableLocalAuth                     *bool                   `pulumi:"disableLocalAuth"`
	DomainName                           *string                 `pulumi:"domainName"`
	Identity                             *IdentityInfo           `pulumi:"identity"`
	InboundIpRules                       []InboundIpRule         `pulumi:"inboundIpRules"`
	InputSchema                          *string                 `pulumi:"inputSchema"`
	InputSchemaMapping                   *JsonInputSchemaMapping `pulumi:"inputSchemaMapping"`
	Location                             *string                 `pulumi:"location"`
	PublicNetworkAccess                  *string                 `pulumi:"publicNetworkAccess"`
	ResourceGroupName                    string                  `pulumi:"resourceGroupName"`
	Tags                                 map[string]string       `pulumi:"tags"`
}


type DomainArgs struct {
	AutoCreateTopicWithFirstSubscription pulumi.BoolPtrInput
	AutoDeleteTopicWithLastSubscription  pulumi.BoolPtrInput
	DisableLocalAuth                     pulumi.BoolPtrInput
	DomainName                           pulumi.StringPtrInput
	Identity                             IdentityInfoPtrInput
	InboundIpRules                       InboundIpRuleArrayInput
	InputSchema                          pulumi.StringPtrInput
	InputSchemaMapping                   JsonInputSchemaMappingPtrInput
	Location                             pulumi.StringPtrInput
	PublicNetworkAccess                  pulumi.StringPtrInput
	ResourceGroupName                    pulumi.StringInput
	Tags                                 pulumi.StringMapInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

type DomainOutput struct{ *pulumi.OutputState }

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

func (o DomainOutput) AutoCreateTopicWithFirstSubscription() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.BoolPtrOutput { return v.AutoCreateTopicWithFirstSubscription }).(pulumi.BoolPtrOutput)
}

func (o DomainOutput) AutoDeleteTopicWithLastSubscription() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.BoolPtrOutput { return v.AutoDeleteTopicWithLastSubscription }).(pulumi.BoolPtrOutput)
}

func (o DomainOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.BoolPtrOutput { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o DomainOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

func (o DomainOutput) Identity() IdentityInfoResponsePtrOutput {
	return o.ApplyT(func(v *Domain) IdentityInfoResponsePtrOutput { return v.Identity }).(IdentityInfoResponsePtrOutput)
}

func (o DomainOutput) InboundIpRules() InboundIpRuleResponseArrayOutput {
	return o.ApplyT(func(v *Domain) InboundIpRuleResponseArrayOutput { return v.InboundIpRules }).(InboundIpRuleResponseArrayOutput)
}

func (o DomainOutput) InputSchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.InputSchema }).(pulumi.StringPtrOutput)
}

func (o DomainOutput) InputSchemaMapping() JsonInputSchemaMappingResponsePtrOutput {
	return o.ApplyT(func(v *Domain) JsonInputSchemaMappingResponsePtrOutput { return v.InputSchemaMapping }).(JsonInputSchemaMappingResponsePtrOutput)
}

func (o DomainOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DomainOutput) MetricResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.MetricResourceId }).(pulumi.StringOutput)
}

func (o DomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DomainOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Domain) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o DomainOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DomainOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o DomainOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Domain) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DomainOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainOutput{})
}
