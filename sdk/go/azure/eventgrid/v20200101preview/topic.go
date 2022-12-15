


package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Topic struct {
	pulumi.CustomResourceState

	Endpoint           pulumi.StringOutput                     `pulumi:"endpoint"`
	InputSchema        pulumi.StringPtrOutput                  `pulumi:"inputSchema"`
	InputSchemaMapping JsonInputSchemaMappingResponsePtrOutput `pulumi:"inputSchemaMapping"`
	Location           pulumi.StringOutput                     `pulumi:"location"`
	MetricResourceId   pulumi.StringOutput                     `pulumi:"metricResourceId"`
	Name               pulumi.StringOutput                     `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput                     `pulumi:"provisioningState"`
	Tags               pulumi.StringMapOutput                  `pulumi:"tags"`
	Type               pulumi.StringOutput                     `pulumi:"type"`
}


func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.InputSchema) {
		args.InputSchema = pulumi.StringPtr("EventGridSchema")
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
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211201:Topic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20220615:Topic"),
		},
	})
	opts = append(opts, aliases)
	var resource Topic
	err := ctx.RegisterResource("azure-native:eventgrid/v20200101preview:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("azure-native:eventgrid/v20200101preview:Topic", name, id, state, &resource, opts...)
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
	InputSchema        *string                 `pulumi:"inputSchema"`
	InputSchemaMapping *JsonInputSchemaMapping `pulumi:"inputSchemaMapping"`
	Location           *string                 `pulumi:"location"`
	ResourceGroupName  string                  `pulumi:"resourceGroupName"`
	Tags               map[string]string       `pulumi:"tags"`
	TopicName          *string                 `pulumi:"topicName"`
}


type TopicArgs struct {
	InputSchema        pulumi.StringPtrInput
	InputSchemaMapping JsonInputSchemaMappingPtrInput
	Location           pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	Tags               pulumi.StringMapInput
	TopicName          pulumi.StringPtrInput
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

func (o TopicOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

func (o TopicOutput) InputSchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.InputSchema }).(pulumi.StringPtrOutput)
}

func (o TopicOutput) InputSchemaMapping() JsonInputSchemaMappingResponsePtrOutput {
	return o.ApplyT(func(v *Topic) JsonInputSchemaMappingResponsePtrOutput { return v.InputSchemaMapping }).(JsonInputSchemaMappingResponsePtrOutput)
}

func (o TopicOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o TopicOutput) MetricResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.MetricResourceId }).(pulumi.StringOutput)
}

func (o TopicOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TopicOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o TopicOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o TopicOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TopicOutput{})
}
