


package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Queue struct {
	pulumi.CustomResourceState

	AccessedAt                          pulumi.StringOutput               `pulumi:"accessedAt"`
	AutoDeleteOnIdle                    pulumi.StringPtrOutput            `pulumi:"autoDeleteOnIdle"`
	CountDetails                        MessageCountDetailsResponseOutput `pulumi:"countDetails"`
	CreatedAt                           pulumi.StringOutput               `pulumi:"createdAt"`
	DeadLetteringOnMessageExpiration    pulumi.BoolPtrOutput              `pulumi:"deadLetteringOnMessageExpiration"`
	DefaultMessageTimeToLive            pulumi.StringPtrOutput            `pulumi:"defaultMessageTimeToLive"`
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrOutput            `pulumi:"duplicateDetectionHistoryTimeWindow"`
	EnableBatchedOperations             pulumi.BoolPtrOutput              `pulumi:"enableBatchedOperations"`
	EnableExpress                       pulumi.BoolPtrOutput              `pulumi:"enableExpress"`
	EnablePartitioning                  pulumi.BoolPtrOutput              `pulumi:"enablePartitioning"`
	ForwardDeadLetteredMessagesTo       pulumi.StringPtrOutput            `pulumi:"forwardDeadLetteredMessagesTo"`
	ForwardTo                           pulumi.StringPtrOutput            `pulumi:"forwardTo"`
	Location                            pulumi.StringOutput               `pulumi:"location"`
	LockDuration                        pulumi.StringPtrOutput            `pulumi:"lockDuration"`
	MaxDeliveryCount                    pulumi.IntPtrOutput               `pulumi:"maxDeliveryCount"`
	MaxMessageSizeInKilobytes           pulumi.Float64PtrOutput           `pulumi:"maxMessageSizeInKilobytes"`
	MaxSizeInMegabytes                  pulumi.IntPtrOutput               `pulumi:"maxSizeInMegabytes"`
	MessageCount                        pulumi.Float64Output              `pulumi:"messageCount"`
	Name                                pulumi.StringOutput               `pulumi:"name"`
	RequiresDuplicateDetection          pulumi.BoolPtrOutput              `pulumi:"requiresDuplicateDetection"`
	RequiresSession                     pulumi.BoolPtrOutput              `pulumi:"requiresSession"`
	SizeInBytes                         pulumi.Float64Output              `pulumi:"sizeInBytes"`
	Status                              pulumi.StringPtrOutput            `pulumi:"status"`
	SystemData                          SystemDataResponseOutput          `pulumi:"systemData"`
	Type                                pulumi.StringOutput               `pulumi:"type"`
	UpdatedAt                           pulumi.StringOutput               `pulumi:"updatedAt"`
}


func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOption) (*Queue, error) {
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
			Type: pulumi.String("azure-native:servicebus:Queue"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20140901:Queue"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20150801:Queue"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:Queue"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:Queue"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210101preview:Queue"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210601preview:Queue"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20220101preview:Queue"),
		},
	})
	opts = append(opts, aliases)
	var resource Queue
	err := ctx.RegisterResource("azure-native:servicebus/v20211101:Queue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueState, opts ...pulumi.ResourceOption) (*Queue, error) {
	var resource Queue
	err := ctx.ReadResource("azure-native:servicebus/v20211101:Queue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type queueState struct {
}

type QueueState struct {
}

func (QueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueState)(nil)).Elem()
}

type queueArgs struct {
	AutoDeleteOnIdle                    *string       `pulumi:"autoDeleteOnIdle"`
	DeadLetteringOnMessageExpiration    *bool         `pulumi:"deadLetteringOnMessageExpiration"`
	DefaultMessageTimeToLive            *string       `pulumi:"defaultMessageTimeToLive"`
	DuplicateDetectionHistoryTimeWindow *string       `pulumi:"duplicateDetectionHistoryTimeWindow"`
	EnableBatchedOperations             *bool         `pulumi:"enableBatchedOperations"`
	EnableExpress                       *bool         `pulumi:"enableExpress"`
	EnablePartitioning                  *bool         `pulumi:"enablePartitioning"`
	ForwardDeadLetteredMessagesTo       *string       `pulumi:"forwardDeadLetteredMessagesTo"`
	ForwardTo                           *string       `pulumi:"forwardTo"`
	LockDuration                        *string       `pulumi:"lockDuration"`
	MaxDeliveryCount                    *int          `pulumi:"maxDeliveryCount"`
	MaxMessageSizeInKilobytes           *float64      `pulumi:"maxMessageSizeInKilobytes"`
	MaxSizeInMegabytes                  *int          `pulumi:"maxSizeInMegabytes"`
	NamespaceName                       string        `pulumi:"namespaceName"`
	QueueName                           *string       `pulumi:"queueName"`
	RequiresDuplicateDetection          *bool         `pulumi:"requiresDuplicateDetection"`
	RequiresSession                     *bool         `pulumi:"requiresSession"`
	ResourceGroupName                   string        `pulumi:"resourceGroupName"`
	Status                              *EntityStatus `pulumi:"status"`
}


type QueueArgs struct {
	AutoDeleteOnIdle                    pulumi.StringPtrInput
	DeadLetteringOnMessageExpiration    pulumi.BoolPtrInput
	DefaultMessageTimeToLive            pulumi.StringPtrInput
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrInput
	EnableBatchedOperations             pulumi.BoolPtrInput
	EnableExpress                       pulumi.BoolPtrInput
	EnablePartitioning                  pulumi.BoolPtrInput
	ForwardDeadLetteredMessagesTo       pulumi.StringPtrInput
	ForwardTo                           pulumi.StringPtrInput
	LockDuration                        pulumi.StringPtrInput
	MaxDeliveryCount                    pulumi.IntPtrInput
	MaxMessageSizeInKilobytes           pulumi.Float64PtrInput
	MaxSizeInMegabytes                  pulumi.IntPtrInput
	NamespaceName                       pulumi.StringInput
	QueueName                           pulumi.StringPtrInput
	RequiresDuplicateDetection          pulumi.BoolPtrInput
	RequiresSession                     pulumi.BoolPtrInput
	ResourceGroupName                   pulumi.StringInput
	Status                              EntityStatusPtrInput
}

func (QueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueArgs)(nil)).Elem()
}

type QueueInput interface {
	pulumi.Input

	ToQueueOutput() QueueOutput
	ToQueueOutputWithContext(ctx context.Context) QueueOutput
}

func (*Queue) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil)).Elem()
}

func (i *Queue) ToQueueOutput() QueueOutput {
	return i.ToQueueOutputWithContext(context.Background())
}

func (i *Queue) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueOutput)
}

type QueueOutput struct{ *pulumi.OutputState }

func (QueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil)).Elem()
}

func (o QueueOutput) ToQueueOutput() QueueOutput {
	return o
}

func (o QueueOutput) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return o
}

func (o QueueOutput) AccessedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.AccessedAt }).(pulumi.StringOutput)
}

func (o QueueOutput) AutoDeleteOnIdle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.AutoDeleteOnIdle }).(pulumi.StringPtrOutput)
}

func (o QueueOutput) CountDetails() MessageCountDetailsResponseOutput {
	return o.ApplyT(func(v *Queue) MessageCountDetailsResponseOutput { return v.CountDetails }).(MessageCountDetailsResponseOutput)
}

func (o QueueOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o QueueOutput) DeadLetteringOnMessageExpiration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.BoolPtrOutput { return v.DeadLetteringOnMessageExpiration }).(pulumi.BoolPtrOutput)
}

func (o QueueOutput) DefaultMessageTimeToLive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.DefaultMessageTimeToLive }).(pulumi.StringPtrOutput)
}

func (o QueueOutput) DuplicateDetectionHistoryTimeWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.DuplicateDetectionHistoryTimeWindow }).(pulumi.StringPtrOutput)
}

func (o QueueOutput) EnableBatchedOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.BoolPtrOutput { return v.EnableBatchedOperations }).(pulumi.BoolPtrOutput)
}

func (o QueueOutput) EnableExpress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.BoolPtrOutput { return v.EnableExpress }).(pulumi.BoolPtrOutput)
}

func (o QueueOutput) EnablePartitioning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.BoolPtrOutput { return v.EnablePartitioning }).(pulumi.BoolPtrOutput)
}

func (o QueueOutput) ForwardDeadLetteredMessagesTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.ForwardDeadLetteredMessagesTo }).(pulumi.StringPtrOutput)
}

func (o QueueOutput) ForwardTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.ForwardTo }).(pulumi.StringPtrOutput)
}

func (o QueueOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o QueueOutput) LockDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.LockDuration }).(pulumi.StringPtrOutput)
}

func (o QueueOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.IntPtrOutput { return v.MaxDeliveryCount }).(pulumi.IntPtrOutput)
}

func (o QueueOutput) MaxMessageSizeInKilobytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.Float64PtrOutput { return v.MaxMessageSizeInKilobytes }).(pulumi.Float64PtrOutput)
}

func (o QueueOutput) MaxSizeInMegabytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.IntPtrOutput { return v.MaxSizeInMegabytes }).(pulumi.IntPtrOutput)
}

func (o QueueOutput) MessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v *Queue) pulumi.Float64Output { return v.MessageCount }).(pulumi.Float64Output)
}

func (o QueueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o QueueOutput) RequiresDuplicateDetection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.BoolPtrOutput { return v.RequiresDuplicateDetection }).(pulumi.BoolPtrOutput)
}

func (o QueueOutput) RequiresSession() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.BoolPtrOutput { return v.RequiresSession }).(pulumi.BoolPtrOutput)
}

func (o QueueOutput) SizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *Queue) pulumi.Float64Output { return v.SizeInBytes }).(pulumi.Float64Output)
}

func (o QueueOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o QueueOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Queue) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o QueueOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o QueueOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(QueueOutput{})
}
