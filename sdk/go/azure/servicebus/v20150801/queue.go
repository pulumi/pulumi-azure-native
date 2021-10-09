


package v20150801

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
	EntityAvailabilityStatus            pulumi.StringPtrOutput            `pulumi:"entityAvailabilityStatus"`
	IsAnonymousAccessible               pulumi.BoolPtrOutput              `pulumi:"isAnonymousAccessible"`
	Location                            pulumi.StringPtrOutput            `pulumi:"location"`
	LockDuration                        pulumi.StringPtrOutput            `pulumi:"lockDuration"`
	MaxDeliveryCount                    pulumi.IntPtrOutput               `pulumi:"maxDeliveryCount"`
	MaxSizeInMegabytes                  pulumi.Float64PtrOutput           `pulumi:"maxSizeInMegabytes"`
	MessageCount                        pulumi.Float64Output              `pulumi:"messageCount"`
	Name                                pulumi.StringOutput               `pulumi:"name"`
	RequiresDuplicateDetection          pulumi.BoolPtrOutput              `pulumi:"requiresDuplicateDetection"`
	RequiresSession                     pulumi.BoolPtrOutput              `pulumi:"requiresSession"`
	SizeInBytes                         pulumi.Float64Output              `pulumi:"sizeInBytes"`
	Status                              pulumi.StringPtrOutput            `pulumi:"status"`
	SupportOrdering                     pulumi.BoolPtrOutput              `pulumi:"supportOrdering"`
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
			Type: pulumi.String("azure-nextgen:servicebus/v20150801:Queue"),
		},
		{
			Type: pulumi.String("azure-native:servicebus:Queue"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus:Queue"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20140901:Queue"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20140901:Queue"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:Queue"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20170401:Queue"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:Queue"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20180101preview:Queue"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210101preview:Queue"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20210101preview:Queue"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210601preview:Queue"),
		},
		{
			Type: pulumi.String("azure-nextgen:servicebus/v20210601preview:Queue"),
		},
	})
	opts = append(opts, aliases)
	var resource Queue
	err := ctx.RegisterResource("azure-native:servicebus/v20150801:Queue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueState, opts ...pulumi.ResourceOption) (*Queue, error) {
	var resource Queue
	err := ctx.ReadResource("azure-native:servicebus/v20150801:Queue", name, id, state, &resource, opts...)
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
	AutoDeleteOnIdle                    *string                   `pulumi:"autoDeleteOnIdle"`
	DeadLetteringOnMessageExpiration    *bool                     `pulumi:"deadLetteringOnMessageExpiration"`
	DefaultMessageTimeToLive            *string                   `pulumi:"defaultMessageTimeToLive"`
	DuplicateDetectionHistoryTimeWindow *string                   `pulumi:"duplicateDetectionHistoryTimeWindow"`
	EnableBatchedOperations             *bool                     `pulumi:"enableBatchedOperations"`
	EnableExpress                       *bool                     `pulumi:"enableExpress"`
	EnablePartitioning                  *bool                     `pulumi:"enablePartitioning"`
	EntityAvailabilityStatus            *EntityAvailabilityStatus `pulumi:"entityAvailabilityStatus"`
	IsAnonymousAccessible               *bool                     `pulumi:"isAnonymousAccessible"`
	Location                            *string                   `pulumi:"location"`
	LockDuration                        *string                   `pulumi:"lockDuration"`
	MaxDeliveryCount                    *int                      `pulumi:"maxDeliveryCount"`
	MaxSizeInMegabytes                  *float64                  `pulumi:"maxSizeInMegabytes"`
	Name                                *string                   `pulumi:"name"`
	NamespaceName                       string                    `pulumi:"namespaceName"`
	QueueName                           *string                   `pulumi:"queueName"`
	RequiresDuplicateDetection          *bool                     `pulumi:"requiresDuplicateDetection"`
	RequiresSession                     *bool                     `pulumi:"requiresSession"`
	ResourceGroupName                   string                    `pulumi:"resourceGroupName"`
	Status                              *EntityStatus             `pulumi:"status"`
	SupportOrdering                     *bool                     `pulumi:"supportOrdering"`
}


type QueueArgs struct {
	AutoDeleteOnIdle                    pulumi.StringPtrInput
	DeadLetteringOnMessageExpiration    pulumi.BoolPtrInput
	DefaultMessageTimeToLive            pulumi.StringPtrInput
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrInput
	EnableBatchedOperations             pulumi.BoolPtrInput
	EnableExpress                       pulumi.BoolPtrInput
	EnablePartitioning                  pulumi.BoolPtrInput
	EntityAvailabilityStatus            EntityAvailabilityStatusPtrInput
	IsAnonymousAccessible               pulumi.BoolPtrInput
	Location                            pulumi.StringPtrInput
	LockDuration                        pulumi.StringPtrInput
	MaxDeliveryCount                    pulumi.IntPtrInput
	MaxSizeInMegabytes                  pulumi.Float64PtrInput
	Name                                pulumi.StringPtrInput
	NamespaceName                       pulumi.StringInput
	QueueName                           pulumi.StringPtrInput
	RequiresDuplicateDetection          pulumi.BoolPtrInput
	RequiresSession                     pulumi.BoolPtrInput
	ResourceGroupName                   pulumi.StringInput
	Status                              EntityStatusPtrInput
	SupportOrdering                     pulumi.BoolPtrInput
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
	return reflect.TypeOf((*Queue)(nil))
}

func (i *Queue) ToQueueOutput() QueueOutput {
	return i.ToQueueOutputWithContext(context.Background())
}

func (i *Queue) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueOutput)
}

type QueueOutput struct{ *pulumi.OutputState }

func (QueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Queue)(nil))
}

func (o QueueOutput) ToQueueOutput() QueueOutput {
	return o
}

func (o QueueOutput) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(QueueOutput{})
}
