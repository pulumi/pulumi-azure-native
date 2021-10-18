


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Topic struct {
	pulumi.CustomResourceState

	DisableLocalAuth           pulumi.BoolPtrOutput                         `pulumi:"disableLocalAuth"`
	Endpoint                   pulumi.StringOutput                          `pulumi:"endpoint"`
	ExtendedLocation           ExtendedLocationResponsePtrOutput            `pulumi:"extendedLocation"`
	Identity                   IdentityInfoResponsePtrOutput                `pulumi:"identity"`
	InboundIpRules             InboundIpRuleResponseArrayOutput             `pulumi:"inboundIpRules"`
	InputSchema                pulumi.StringPtrOutput                       `pulumi:"inputSchema"`
	InputSchemaMapping         JsonInputSchemaMappingResponsePtrOutput      `pulumi:"inputSchemaMapping"`
	Kind                       pulumi.StringPtrOutput                       `pulumi:"kind"`
	Location                   pulumi.StringOutput                          `pulumi:"location"`
	MetricResourceId           pulumi.StringOutput                          `pulumi:"metricResourceId"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringOutput                          `pulumi:"provisioningState"`
	PublicNetworkAccess        pulumi.StringPtrOutput                       `pulumi:"publicNetworkAccess"`
	Sku                        ResourceSkuResponsePtrOutput                 `pulumi:"sku"`
	SystemData                 SystemDataResponseOutput                     `pulumi:"systemData"`
	Tags                       pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                       pulumi.StringOutput                          `pulumi:"type"`
}


func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.DisableLocalAuth == nil {
		args.DisableLocalAuth = pulumi.BoolPtr(false)
	}
	if args.InputSchema == nil {
		args.InputSchema = pulumi.StringPtr("EventGridSchema")
	}
	if args.Kind == nil {
		args.Kind = pulumi.StringPtr("Azure")
	}
	if args.PublicNetworkAccess == nil {
		args.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20210601preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20170615preview:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20170615preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20170915preview:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20170915preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180101:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20180101:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180501preview:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20180501preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180915preview:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20180915preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190101:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20190101:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190201preview:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20190201preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190601:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20190601:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200101preview:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20200101preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20200401preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200601:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20200601:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20201015preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211201:Topic"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20211201:Topic"),
		},
	})
	opts = append(opts, aliases)
	var resource Topic
	err := ctx.RegisterResource("azure-native:eventgrid/v20210601preview:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("azure-native:eventgrid/v20210601preview:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type topicState struct {
}

type TopicState struct {
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	DisableLocalAuth    *bool                   `pulumi:"disableLocalAuth"`
	ExtendedLocation    *ExtendedLocation       `pulumi:"extendedLocation"`
	Identity            *IdentityInfo           `pulumi:"identity"`
	InboundIpRules      []InboundIpRule         `pulumi:"inboundIpRules"`
	InputSchema         *string                 `pulumi:"inputSchema"`
	InputSchemaMapping  *JsonInputSchemaMapping `pulumi:"inputSchemaMapping"`
	Kind                *string                 `pulumi:"kind"`
	Location            *string                 `pulumi:"location"`
	PublicNetworkAccess *string                 `pulumi:"publicNetworkAccess"`
	ResourceGroupName   string                  `pulumi:"resourceGroupName"`
	Sku                 *ResourceSku            `pulumi:"sku"`
	Tags                map[string]string       `pulumi:"tags"`
	TopicName           *string                 `pulumi:"topicName"`
}


type TopicArgs struct {
	DisableLocalAuth    pulumi.BoolPtrInput
	ExtendedLocation    ExtendedLocationPtrInput
	Identity            IdentityInfoPtrInput
	InboundIpRules      InboundIpRuleArrayInput
	InputSchema         pulumi.StringPtrInput
	InputSchemaMapping  JsonInputSchemaMappingPtrInput
	Kind                pulumi.StringPtrInput
	Location            pulumi.StringPtrInput
	PublicNetworkAccess pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	Sku                 ResourceSkuPtrInput
	Tags                pulumi.StringMapInput
	TopicName           pulumi.StringPtrInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}

type TopicInput interface {
	pulumi.Input

	ToTopicOutput() TopicOutput
	ToTopicOutputWithContext(ctx context.Context) TopicOutput
}

func (*Topic) ElementType() reflect.Type {
	return reflect.TypeOf((*Topic)(nil))
}

func (i *Topic) ToTopicOutput() TopicOutput {
	return i.ToTopicOutputWithContext(context.Background())
}

func (i *Topic) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicOutput)
}

type TopicOutput struct{ *pulumi.OutputState }

func (TopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Topic)(nil))
}

func (o TopicOutput) ToTopicOutput() TopicOutput {
	return o
}

func (o TopicOutput) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TopicOutput{})
}
