


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConsumerGroup struct {
	pulumi.CustomResourceState

	CreatedAt    pulumi.StringOutput      `pulumi:"createdAt"`
	Name         pulumi.StringOutput      `pulumi:"name"`
	SystemData   SystemDataResponseOutput `pulumi:"systemData"`
	Type         pulumi.StringOutput      `pulumi:"type"`
	UpdatedAt    pulumi.StringOutput      `pulumi:"updatedAt"`
	UserMetadata pulumi.StringPtrOutput   `pulumi:"userMetadata"`
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
			Type: pulumi.String("azure-native:eventhub/v20140901:ConsumerGroup"),
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
			Type: pulumi.String("azure-native:eventhub/v20211101:ConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20220101preview:ConsumerGroup"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20221001preview:ConsumerGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ConsumerGroup
	err := ctx.RegisterResource("azure-native:eventhub/v20210601preview:ConsumerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConsumerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsumerGroupState, opts ...pulumi.ResourceOption) (*ConsumerGroup, error) {
	var resource ConsumerGroup
	err := ctx.ReadResource("azure-native:eventhub/v20210601preview:ConsumerGroup", name, id, state, &resource, opts...)
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
	NamespaceName     string  `pulumi:"namespaceName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	UserMetadata      *string `pulumi:"userMetadata"`
}


type ConsumerGroupArgs struct {
	ConsumerGroupName pulumi.StringPtrInput
	EventHubName      pulumi.StringInput
	NamespaceName     pulumi.StringInput
	ResourceGroupName pulumi.StringInput
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

func (o ConsumerGroupOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerGroup) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o ConsumerGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConsumerGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConsumerGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConsumerGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ConsumerGroupOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ConsumerGroup) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o ConsumerGroupOutput) UserMetadata() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsumerGroup) pulumi.StringPtrOutput { return v.UserMetadata }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ConsumerGroupOutput{})
}
