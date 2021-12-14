


package v20140901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConsumerGroup struct {
	pulumi.CustomResourceState

	CreatedAt    pulumi.StringOutput    `pulumi:"createdAt"`
	EventHubPath pulumi.StringOutput    `pulumi:"eventHubPath"`
	Location     pulumi.StringPtrOutput `pulumi:"location"`
	Name         pulumi.StringOutput    `pulumi:"name"`
	Type         pulumi.StringOutput    `pulumi:"type"`
	UpdatedAt    pulumi.StringOutput    `pulumi:"updatedAt"`
	UserMetadata pulumi.StringPtrOutput `pulumi:"userMetadata"`
}


func NewConsumerGroup(ctx *pulumi.Context,
	name string, args *ConsumerGroupArgs, opts ...pulumi.ResourceOption) (*ConsumerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventHubName == nil {
		return nil, errors.New("invalid value for required argument 'EventHubName'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventhub:ConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20150801:ConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20170401:ConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20180101preview:ConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210101preview:ConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210601preview:ConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20211101:ConsumerGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ConsumerGroup
	err := ctx.RegisterResource("azure-native:eventhub/v20140901:ConsumerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConsumerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsumerGroupState, opts ...pulumi.ResourceOption) (*ConsumerGroup, error) {
	var resource ConsumerGroup
	err := ctx.ReadResource("azure-native:eventhub/v20140901:ConsumerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type consumerGroupState struct {
}

type ConsumerGroupState struct {
}

func (ConsumerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerGroupState)(nil)).Elem()
}

type consumerGroupArgs struct {
	ConsumerGroupName *string `pulumi:"consumerGroupName"`
	EventHubName      string  `pulumi:"eventHubName"`
	Location          *string `pulumi:"location"`
	Name              *string `pulumi:"name"`
	NamespaceName     string  `pulumi:"namespaceName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	Type              *string `pulumi:"type"`
	UserMetadata      *string `pulumi:"userMetadata"`
}


type ConsumerGroupArgs struct {
	ConsumerGroupName pulumi.StringPtrInput
	EventHubName      pulumi.StringInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	NamespaceName     pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Type              pulumi.StringPtrInput
	UserMetadata      pulumi.StringPtrInput
}

func (ConsumerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerGroupArgs)(nil)).Elem()
}

type ConsumerGroupInput interface {
	pulumi.Input

	ToConsumerGroupOutput() ConsumerGroupOutput
	ToConsumerGroupOutputWithContext(ctx context.Context) ConsumerGroupOutput
}

func (*ConsumerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsumerGroup)(nil)).Elem()
}

func (i *ConsumerGroup) ToConsumerGroupOutput() ConsumerGroupOutput {
	return i.ToConsumerGroupOutputWithContext(context.Background())
}

func (i *ConsumerGroup) ToConsumerGroupOutputWithContext(ctx context.Context) ConsumerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerGroupOutput)
}

type ConsumerGroupOutput struct{ *pulumi.OutputState }

func (ConsumerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsumerGroup)(nil)).Elem()
}

func (o ConsumerGroupOutput) ToConsumerGroupOutput() ConsumerGroupOutput {
	return o
}

func (o ConsumerGroupOutput) ToConsumerGroupOutputWithContext(ctx context.Context) ConsumerGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConsumerGroupOutput{})
}
