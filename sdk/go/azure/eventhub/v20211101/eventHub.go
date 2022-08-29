


package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EventHub struct {
	pulumi.CustomResourceState

	CaptureDescription     CaptureDescriptionResponsePtrOutput `pulumi:"captureDescription"`
	CreatedAt              pulumi.StringOutput                 `pulumi:"createdAt"`
	Location               pulumi.StringOutput                 `pulumi:"location"`
	MessageRetentionInDays pulumi.Float64PtrOutput             `pulumi:"messageRetentionInDays"`
	Name                   pulumi.StringOutput                 `pulumi:"name"`
	PartitionCount         pulumi.Float64PtrOutput             `pulumi:"partitionCount"`
	PartitionIds           pulumi.StringArrayOutput            `pulumi:"partitionIds"`
	Status                 pulumi.StringPtrOutput              `pulumi:"status"`
	SystemData             SystemDataResponseOutput            `pulumi:"systemData"`
	Type                   pulumi.StringOutput                 `pulumi:"type"`
	UpdatedAt              pulumi.StringOutput                 `pulumi:"updatedAt"`
}


func NewEventHub(ctx *pulumi.Context,
	name string, args *EventHubArgs, opts ...pulumi.ResourceOption) (*EventHub, error) {
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
			Type: pulumi.String("azure-native:eventhub:EventHub"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20140901:EventHub"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20150801:EventHub"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20170401:EventHub"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20180101preview:EventHub"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210101preview:EventHub"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210601preview:EventHub"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20220101preview:EventHub"),
		},
	})
	opts = append(opts, aliases)
	var resource EventHub
	err := ctx.RegisterResource("azure-native:eventhub/v20211101:EventHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEventHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubState, opts ...pulumi.ResourceOption) (*EventHub, error) {
	var resource EventHub
	err := ctx.ReadResource("azure-native:eventhub/v20211101:EventHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type eventHubState struct {
}

type EventHubState struct {
}

func (EventHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubState)(nil)).Elem()
}

type eventHubArgs struct {
	CaptureDescription     *CaptureDescription `pulumi:"captureDescription"`
	EventHubName           *string             `pulumi:"eventHubName"`
	MessageRetentionInDays *float64            `pulumi:"messageRetentionInDays"`
	NamespaceName          string              `pulumi:"namespaceName"`
	PartitionCount         *float64            `pulumi:"partitionCount"`
	ResourceGroupName      string              `pulumi:"resourceGroupName"`
	Status                 *EntityStatus       `pulumi:"status"`
}


type EventHubArgs struct {
	CaptureDescription     CaptureDescriptionPtrInput
	EventHubName           pulumi.StringPtrInput
	MessageRetentionInDays pulumi.Float64PtrInput
	NamespaceName          pulumi.StringInput
	PartitionCount         pulumi.Float64PtrInput
	ResourceGroupName      pulumi.StringInput
	Status                 EntityStatusPtrInput
}

func (EventHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubArgs)(nil)).Elem()
}

type EventHubInput interface {
	pulumi.Input

	ToEventHubOutput() EventHubOutput
	ToEventHubOutputWithContext(ctx context.Context) EventHubOutput
}

func (*EventHub) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHub)(nil)).Elem()
}

func (i *EventHub) ToEventHubOutput() EventHubOutput {
	return i.ToEventHubOutputWithContext(context.Background())
}

func (i *EventHub) ToEventHubOutputWithContext(ctx context.Context) EventHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubOutput)
}

type EventHubOutput struct{ *pulumi.OutputState }

func (EventHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHub)(nil)).Elem()
}

func (o EventHubOutput) ToEventHubOutput() EventHubOutput {
	return o
}

func (o EventHubOutput) ToEventHubOutputWithContext(ctx context.Context) EventHubOutput {
	return o
}

func (o EventHubOutput) CaptureDescription() CaptureDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *EventHub) CaptureDescriptionResponsePtrOutput { return v.CaptureDescription }).(CaptureDescriptionResponsePtrOutput)
}

func (o EventHubOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHub) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o EventHubOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHub) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o EventHubOutput) MessageRetentionInDays() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EventHub) pulumi.Float64PtrOutput { return v.MessageRetentionInDays }).(pulumi.Float64PtrOutput)
}

func (o EventHubOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHub) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EventHubOutput) PartitionCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EventHub) pulumi.Float64PtrOutput { return v.PartitionCount }).(pulumi.Float64PtrOutput)
}

func (o EventHubOutput) PartitionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventHub) pulumi.StringArrayOutput { return v.PartitionIds }).(pulumi.StringArrayOutput)
}

func (o EventHubOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHub) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o EventHubOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EventHub) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o EventHubOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHub) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o EventHubOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHub) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EventHubOutput{})
}
