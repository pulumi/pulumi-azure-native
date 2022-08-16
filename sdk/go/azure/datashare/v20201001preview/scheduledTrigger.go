


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScheduledTrigger struct {
	pulumi.CustomResourceState

	CreatedAt           pulumi.StringOutput      `pulumi:"createdAt"`
	Kind                pulumi.StringOutput      `pulumi:"kind"`
	Name                pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState   pulumi.StringOutput      `pulumi:"provisioningState"`
	RecurrenceInterval  pulumi.StringOutput      `pulumi:"recurrenceInterval"`
	SynchronizationMode pulumi.StringPtrOutput   `pulumi:"synchronizationMode"`
	SynchronizationTime pulumi.StringOutput      `pulumi:"synchronizationTime"`
	SystemData          SystemDataResponseOutput `pulumi:"systemData"`
	TriggerStatus       pulumi.StringOutput      `pulumi:"triggerStatus"`
	Type                pulumi.StringOutput      `pulumi:"type"`
	UserName            pulumi.StringOutput      `pulumi:"userName"`
}


func NewScheduledTrigger(ctx *pulumi.Context,
	name string, args *ScheduledTriggerArgs, opts ...pulumi.ResourceOption) (*ScheduledTrigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.RecurrenceInterval == nil {
		return nil, errors.New("invalid value for required argument 'RecurrenceInterval'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareSubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'ShareSubscriptionName'")
	}
	if args.SynchronizationTime == nil {
		return nil, errors.New("invalid value for required argument 'SynchronizationTime'")
	}
	args.Kind = pulumi.String("ScheduleBased")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:ScheduledTrigger"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:ScheduledTrigger"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:ScheduledTrigger"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:ScheduledTrigger"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:ScheduledTrigger"),
		},
	})
	opts = append(opts, aliases)
	var resource ScheduledTrigger
	err := ctx.RegisterResource("azure-native:datashare/v20201001preview:ScheduledTrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScheduledTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledTriggerState, opts ...pulumi.ResourceOption) (*ScheduledTrigger, error) {
	var resource ScheduledTrigger
	err := ctx.ReadResource("azure-native:datashare/v20201001preview:ScheduledTrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type scheduledTriggerState struct {
}

type ScheduledTriggerState struct {
}

func (ScheduledTriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledTriggerState)(nil)).Elem()
}

type scheduledTriggerArgs struct {
	AccountName           string  `pulumi:"accountName"`
	Kind                  string  `pulumi:"kind"`
	RecurrenceInterval    string  `pulumi:"recurrenceInterval"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName string  `pulumi:"shareSubscriptionName"`
	SynchronizationMode   *string `pulumi:"synchronizationMode"`
	SynchronizationTime   string  `pulumi:"synchronizationTime"`
	TriggerName           *string `pulumi:"triggerName"`
}


type ScheduledTriggerArgs struct {
	AccountName           pulumi.StringInput
	Kind                  pulumi.StringInput
	RecurrenceInterval    pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	ShareSubscriptionName pulumi.StringInput
	SynchronizationMode   pulumi.StringPtrInput
	SynchronizationTime   pulumi.StringInput
	TriggerName           pulumi.StringPtrInput
}

func (ScheduledTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledTriggerArgs)(nil)).Elem()
}

type ScheduledTriggerInput interface {
	pulumi.Input

	ToScheduledTriggerOutput() ScheduledTriggerOutput
	ToScheduledTriggerOutputWithContext(ctx context.Context) ScheduledTriggerOutput
}

func (*ScheduledTrigger) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledTrigger)(nil)).Elem()
}

func (i *ScheduledTrigger) ToScheduledTriggerOutput() ScheduledTriggerOutput {
	return i.ToScheduledTriggerOutputWithContext(context.Background())
}

func (i *ScheduledTrigger) ToScheduledTriggerOutputWithContext(ctx context.Context) ScheduledTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledTriggerOutput)
}

type ScheduledTriggerOutput struct{ *pulumi.OutputState }

func (ScheduledTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledTrigger)(nil)).Elem()
}

func (o ScheduledTriggerOutput) ToScheduledTriggerOutput() ScheduledTriggerOutput {
	return o
}

func (o ScheduledTriggerOutput) ToScheduledTriggerOutputWithContext(ctx context.Context) ScheduledTriggerOutput {
	return o
}

func (o ScheduledTriggerOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledTrigger) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o ScheduledTriggerOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledTrigger) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ScheduledTriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledTrigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ScheduledTriggerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledTrigger) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ScheduledTriggerOutput) RecurrenceInterval() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledTrigger) pulumi.StringOutput { return v.RecurrenceInterval }).(pulumi.StringOutput)
}

func (o ScheduledTriggerOutput) SynchronizationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledTrigger) pulumi.StringPtrOutput { return v.SynchronizationMode }).(pulumi.StringPtrOutput)
}

func (o ScheduledTriggerOutput) SynchronizationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledTrigger) pulumi.StringOutput { return v.SynchronizationTime }).(pulumi.StringOutput)
}

func (o ScheduledTriggerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ScheduledTrigger) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ScheduledTriggerOutput) TriggerStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledTrigger) pulumi.StringOutput { return v.TriggerStatus }).(pulumi.StringOutput)
}

func (o ScheduledTriggerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledTrigger) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ScheduledTriggerOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledTrigger) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduledTriggerOutput{})
}
