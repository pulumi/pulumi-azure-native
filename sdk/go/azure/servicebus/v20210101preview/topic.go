


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Topic struct {
	pulumi.CustomResourceState

	AccessedAt                          pulumi.StringOutput               `pulumi:"accessedAt"`
	AutoDeleteOnIdle                    pulumi.StringPtrOutput            `pulumi:"autoDeleteOnIdle"`
	CountDetails                        MessageCountDetailsResponseOutput `pulumi:"countDetails"`
	CreatedAt                           pulumi.StringOutput               `pulumi:"createdAt"`
	DefaultMessageTimeToLive            pulumi.StringPtrOutput            `pulumi:"defaultMessageTimeToLive"`
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrOutput            `pulumi:"duplicateDetectionHistoryTimeWindow"`
	EnableBatchedOperations             pulumi.BoolPtrOutput              `pulumi:"enableBatchedOperations"`
	EnableExpress                       pulumi.BoolPtrOutput              `pulumi:"enableExpress"`
	EnablePartitioning                  pulumi.BoolPtrOutput              `pulumi:"enablePartitioning"`
	MaxSizeInMegabytes                  pulumi.IntPtrOutput               `pulumi:"maxSizeInMegabytes"`
	Name                                pulumi.StringOutput               `pulumi:"name"`
	RequiresDuplicateDetection          pulumi.BoolPtrOutput              `pulumi:"requiresDuplicateDetection"`
	SizeInBytes                         pulumi.Float64Output              `pulumi:"sizeInBytes"`
	Status                              pulumi.StringPtrOutput            `pulumi:"status"`
	SubscriptionCount                   pulumi.IntOutput                  `pulumi:"subscriptionCount"`
	SupportOrdering                     pulumi.BoolPtrOutput              `pulumi:"supportOrdering"`
	SystemData                          SystemDataResponseOutput          `pulumi:"systemData"`
	Type                                pulumi.StringOutput               `pulumi:"type"`
	UpdatedAt                           pulumi.StringOutput               `pulumi:"updatedAt"`
}


func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicebus:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20140901:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20150801:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210601preview:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20211101:Topic"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20220101preview:Topic"),
		},
	})
	opts = append(opts, aliases)
	var resource Topic
	err := ctx.RegisterResource("azure-native:servicebus/v20210101preview:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("azure-native:servicebus/v20210101preview:Topic", name, id, state, &resource, opts...)
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
	AutoDeleteOnIdle                    *string       `pulumi:"autoDeleteOnIdle"`
	DefaultMessageTimeToLive            *string       `pulumi:"defaultMessageTimeToLive"`
	DuplicateDetectionHistoryTimeWindow *string       `pulumi:"duplicateDetectionHistoryTimeWindow"`
	EnableBatchedOperations             *bool         `pulumi:"enableBatchedOperations"`
	EnableExpress                       *bool         `pulumi:"enableExpress"`
	EnablePartitioning                  *bool         `pulumi:"enablePartitioning"`
	MaxSizeInMegabytes                  *int          `pulumi:"maxSizeInMegabytes"`
	NamespaceName                       string        `pulumi:"namespaceName"`
	RequiresDuplicateDetection          *bool         `pulumi:"requiresDuplicateDetection"`
	ResourceGroupName                   string        `pulumi:"resourceGroupName"`
	Status                              *EntityStatus `pulumi:"status"`
	SupportOrdering                     *bool         `pulumi:"supportOrdering"`
	TopicName                           *string       `pulumi:"topicName"`
}


type TopicArgs struct {
	AutoDeleteOnIdle                    pulumi.StringPtrInput
	DefaultMessageTimeToLive            pulumi.StringPtrInput
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrInput
	EnableBatchedOperations             pulumi.BoolPtrInput
	EnableExpress                       pulumi.BoolPtrInput
	EnablePartitioning                  pulumi.BoolPtrInput
	MaxSizeInMegabytes                  pulumi.IntPtrInput
	NamespaceName                       pulumi.StringInput
	RequiresDuplicateDetection          pulumi.BoolPtrInput
	ResourceGroupName                   pulumi.StringInput
	Status                              EntityStatusPtrInput
	SupportOrdering                     pulumi.BoolPtrInput
	TopicName                           pulumi.StringPtrInput
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

func (o TopicOutput) AccessedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.AccessedAt }).(pulumi.StringOutput)
}

func (o TopicOutput) AutoDeleteOnIdle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.AutoDeleteOnIdle }).(pulumi.StringPtrOutput)
}

func (o TopicOutput) CountDetails() MessageCountDetailsResponseOutput {
	return o.ApplyT(func(v *Topic) MessageCountDetailsResponseOutput { return v.CountDetails }).(MessageCountDetailsResponseOutput)
}

func (o TopicOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o TopicOutput) DefaultMessageTimeToLive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.DefaultMessageTimeToLive }).(pulumi.StringPtrOutput)
}

func (o TopicOutput) DuplicateDetectionHistoryTimeWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.DuplicateDetectionHistoryTimeWindow }).(pulumi.StringPtrOutput)
}

func (o TopicOutput) EnableBatchedOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.BoolPtrOutput { return v.EnableBatchedOperations }).(pulumi.BoolPtrOutput)
}

func (o TopicOutput) EnableExpress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.BoolPtrOutput { return v.EnableExpress }).(pulumi.BoolPtrOutput)
}

func (o TopicOutput) EnablePartitioning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.BoolPtrOutput { return v.EnablePartitioning }).(pulumi.BoolPtrOutput)
}

func (o TopicOutput) MaxSizeInMegabytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntPtrOutput { return v.MaxSizeInMegabytes }).(pulumi.IntPtrOutput)
}

func (o TopicOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TopicOutput) RequiresDuplicateDetection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.BoolPtrOutput { return v.RequiresDuplicateDetection }).(pulumi.BoolPtrOutput)
}

func (o TopicOutput) SizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *Topic) pulumi.Float64Output { return v.SizeInBytes }).(pulumi.Float64Output)
}

func (o TopicOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o TopicOutput) SubscriptionCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntOutput { return v.SubscriptionCount }).(pulumi.IntOutput)
}

func (o TopicOutput) SupportOrdering() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.BoolPtrOutput { return v.SupportOrdering }).(pulumi.BoolPtrOutput)
}

func (o TopicOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Topic) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o TopicOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o TopicOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TopicOutput{})
}
