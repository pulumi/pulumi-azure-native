


package v20211201

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
	Identity                   IdentityInfoResponsePtrOutput                `pulumi:"identity"`
	InboundIpRules             InboundIpRuleResponseArrayOutput             `pulumi:"inboundIpRules"`
	InputSchema                pulumi.StringPtrOutput                       `pulumi:"inputSchema"`
	InputSchemaMapping         JsonInputSchemaMappingResponsePtrOutput      `pulumi:"inputSchemaMapping"`
	Location                   pulumi.StringOutput                          `pulumi:"location"`
	MetricResourceId           pulumi.StringOutput                          `pulumi:"metricResourceId"`
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringOutput                          `pulumi:"provisioningState"`
	PublicNetworkAccess        pulumi.StringPtrOutput                       `pulumi:"publicNetworkAccess"`
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
			Type: pulumi.String("azure-native:eventgrid:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20170615preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20170915preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180101:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180501preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180915preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190101:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190201preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190601:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200101preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200601:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20210601preview:Topic"),
		},
	})
	opts = append(opts, aliases)
	var resource Topic
	err := ctx.RegisterResource("azure-native:eventgrid/v20211201:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("azure-native:eventgrid/v20211201:Topic", name, id, state, &resource, opts...)
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
	Identity            *IdentityInfo           `pulumi:"identity"`
	InboundIpRules      []InboundIpRule         `pulumi:"inboundIpRules"`
	InputSchema         *string                 `pulumi:"inputSchema"`
	InputSchemaMapping  *JsonInputSchemaMapping `pulumi:"inputSchemaMapping"`
	Location            *string                 `pulumi:"location"`
	PublicNetworkAccess *string                 `pulumi:"publicNetworkAccess"`
	ResourceGroupName   string                  `pulumi:"resourceGroupName"`
	Tags                map[string]string       `pulumi:"tags"`
	TopicName           *string                 `pulumi:"topicName"`
}


type TopicArgs struct {
	DisableLocalAuth    pulumi.BoolPtrInput
	Identity            IdentityInfoPtrInput
	InboundIpRules      InboundIpRuleArrayInput
	InputSchema         pulumi.StringPtrInput
	InputSchemaMapping  JsonInputSchemaMappingPtrInput
	Location            pulumi.StringPtrInput
	PublicNetworkAccess pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
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
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (i *Topic) ToTopicOutput() TopicOutput {
	return i.ToTopicOutputWithContext(context.Background())
}

func (i *Topic) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicOutput)
}

type TopicOutput struct{ *pulumi.OutputState }

func (TopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
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
