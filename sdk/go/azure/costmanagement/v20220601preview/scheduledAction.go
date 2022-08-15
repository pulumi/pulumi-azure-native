


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScheduledAction struct {
	pulumi.CustomResourceState

	DisplayName     pulumi.StringOutput                  `pulumi:"displayName"`
	ETag            pulumi.StringOutput                  `pulumi:"eTag"`
	FileDestination FileDestinationResponsePtrOutput     `pulumi:"fileDestination"`
	Kind            pulumi.StringPtrOutput               `pulumi:"kind"`
	Name            pulumi.StringOutput                  `pulumi:"name"`
	Notification    NotificationPropertiesResponseOutput `pulumi:"notification"`
	Schedule        SchedulePropertiesResponseOutput     `pulumi:"schedule"`
	Scope           pulumi.StringPtrOutput               `pulumi:"scope"`
	Status          pulumi.StringOutput                  `pulumi:"status"`
	SystemData      SystemDataResponseOutput             `pulumi:"systemData"`
	Type            pulumi.StringOutput                  `pulumi:"type"`
	ViewId          pulumi.StringOutput                  `pulumi:"viewId"`
}


func NewScheduledAction(ctx *pulumi.Context,
	name string, args *ScheduledActionArgs, opts ...pulumi.ResourceOption) (*ScheduledAction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Notification == nil {
		return nil, errors.New("invalid value for required argument 'Notification'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	if args.ViewId == nil {
		return nil, errors.New("invalid value for required argument 'ViewId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement:ScheduledAction"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20220401preview:ScheduledAction"),
		},
	})
	opts = append(opts, aliases)
	var resource ScheduledAction
	err := ctx.RegisterResource("azure-native:costmanagement/v20220601preview:ScheduledAction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScheduledAction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledActionState, opts ...pulumi.ResourceOption) (*ScheduledAction, error) {
	var resource ScheduledAction
	err := ctx.ReadResource("azure-native:costmanagement/v20220601preview:ScheduledAction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type scheduledActionState struct {
}

type ScheduledActionState struct {
}

func (ScheduledActionState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledActionState)(nil)).Elem()
}

type scheduledActionArgs struct {
	DisplayName     string                 `pulumi:"displayName"`
	FileDestination *FileDestination       `pulumi:"fileDestination"`
	Kind            *string                `pulumi:"kind"`
	Name            *string                `pulumi:"name"`
	Notification    NotificationProperties `pulumi:"notification"`
	Schedule        ScheduleProperties     `pulumi:"schedule"`
	Scope           *string                `pulumi:"scope"`
	Status          string                 `pulumi:"status"`
	ViewId          string                 `pulumi:"viewId"`
}


type ScheduledActionArgs struct {
	DisplayName     pulumi.StringInput
	FileDestination FileDestinationPtrInput
	Kind            pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	Notification    NotificationPropertiesInput
	Schedule        SchedulePropertiesInput
	Scope           pulumi.StringPtrInput
	Status          pulumi.StringInput
	ViewId          pulumi.StringInput
}

func (ScheduledActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledActionArgs)(nil)).Elem()
}

type ScheduledActionInput interface {
	pulumi.Input

	ToScheduledActionOutput() ScheduledActionOutput
	ToScheduledActionOutputWithContext(ctx context.Context) ScheduledActionOutput
}

func (*ScheduledAction) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledAction)(nil)).Elem()
}

func (i *ScheduledAction) ToScheduledActionOutput() ScheduledActionOutput {
	return i.ToScheduledActionOutputWithContext(context.Background())
}

func (i *ScheduledAction) ToScheduledActionOutputWithContext(ctx context.Context) ScheduledActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledActionOutput)
}

type ScheduledActionOutput struct{ *pulumi.OutputState }

func (ScheduledActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledAction)(nil)).Elem()
}

func (o ScheduledActionOutput) ToScheduledActionOutput() ScheduledActionOutput {
	return o
}

func (o ScheduledActionOutput) ToScheduledActionOutputWithContext(ctx context.Context) ScheduledActionOutput {
	return o
}

func (o ScheduledActionOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAction) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ScheduledActionOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAction) pulumi.StringOutput { return v.ETag }).(pulumi.StringOutput)
}

func (o ScheduledActionOutput) FileDestination() FileDestinationResponsePtrOutput {
	return o.ApplyT(func(v *ScheduledAction) FileDestinationResponsePtrOutput { return v.FileDestination }).(FileDestinationResponsePtrOutput)
}

func (o ScheduledActionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledAction) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ScheduledActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAction) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ScheduledActionOutput) Notification() NotificationPropertiesResponseOutput {
	return o.ApplyT(func(v *ScheduledAction) NotificationPropertiesResponseOutput { return v.Notification }).(NotificationPropertiesResponseOutput)
}

func (o ScheduledActionOutput) Schedule() SchedulePropertiesResponseOutput {
	return o.ApplyT(func(v *ScheduledAction) SchedulePropertiesResponseOutput { return v.Schedule }).(SchedulePropertiesResponseOutput)
}

func (o ScheduledActionOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledAction) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o ScheduledActionOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAction) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o ScheduledActionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ScheduledAction) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ScheduledActionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAction) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ScheduledActionOutput) ViewId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAction) pulumi.StringOutput { return v.ViewId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduledActionOutput{})
}
