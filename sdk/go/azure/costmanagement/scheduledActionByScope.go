


package costmanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScheduledActionByScope struct {
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


func NewScheduledActionByScope(ctx *pulumi.Context,
	name string, args *ScheduledActionByScopeArgs, opts ...pulumi.ResourceOption) (*ScheduledActionByScope, error) {
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
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	if args.ViewId == nil {
		return nil, errors.New("invalid value for required argument 'ViewId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement/v20220401preview:ScheduledActionByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20220601preview:ScheduledActionByScope"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20221001:ScheduledActionByScope"),
		},
	})
	opts = append(opts, aliases)
	var resource ScheduledActionByScope
	err := ctx.RegisterResource("azure-native:costmanagement:ScheduledActionByScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScheduledActionByScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledActionByScopeState, opts ...pulumi.ResourceOption) (*ScheduledActionByScope, error) {
	var resource ScheduledActionByScope
	err := ctx.ReadResource("azure-native:costmanagement:ScheduledActionByScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type scheduledActionByScopeState struct {
}

type ScheduledActionByScopeState struct {
}

func (ScheduledActionByScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledActionByScopeState)(nil)).Elem()
}

type scheduledActionByScopeArgs struct {
	DisplayName     string                 `pulumi:"displayName"`
	FileDestination *FileDestination       `pulumi:"fileDestination"`
	Kind            *string                `pulumi:"kind"`
	Name            *string                `pulumi:"name"`
	Notification    NotificationProperties `pulumi:"notification"`
	Schedule        ScheduleProperties     `pulumi:"schedule"`
	Scope           string                 `pulumi:"scope"`
	Status          string                 `pulumi:"status"`
	ViewId          string                 `pulumi:"viewId"`
}


type ScheduledActionByScopeArgs struct {
	DisplayName     pulumi.StringInput
	FileDestination FileDestinationPtrInput
	Kind            pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	Notification    NotificationPropertiesInput
	Schedule        SchedulePropertiesInput
	Scope           pulumi.StringInput
	Status          pulumi.StringInput
	ViewId          pulumi.StringInput
}

func (ScheduledActionByScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledActionByScopeArgs)(nil)).Elem()
}

type ScheduledActionByScopeInput interface {
	pulumi.Input

	ToScheduledActionByScopeOutput() ScheduledActionByScopeOutput
	ToScheduledActionByScopeOutputWithContext(ctx context.Context) ScheduledActionByScopeOutput
}

func (*ScheduledActionByScope) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledActionByScope)(nil)).Elem()
}

func (i *ScheduledActionByScope) ToScheduledActionByScopeOutput() ScheduledActionByScopeOutput {
	return i.ToScheduledActionByScopeOutputWithContext(context.Background())
}

func (i *ScheduledActionByScope) ToScheduledActionByScopeOutputWithContext(ctx context.Context) ScheduledActionByScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledActionByScopeOutput)
}

type ScheduledActionByScopeOutput struct{ *pulumi.OutputState }

func (ScheduledActionByScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledActionByScope)(nil)).Elem()
}

func (o ScheduledActionByScopeOutput) ToScheduledActionByScopeOutput() ScheduledActionByScopeOutput {
	return o
}

func (o ScheduledActionByScopeOutput) ToScheduledActionByScopeOutputWithContext(ctx context.Context) ScheduledActionByScopeOutput {
	return o
}

func (o ScheduledActionByScopeOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ScheduledActionByScopeOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) pulumi.StringOutput { return v.ETag }).(pulumi.StringOutput)
}

func (o ScheduledActionByScopeOutput) FileDestination() FileDestinationResponsePtrOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) FileDestinationResponsePtrOutput { return v.FileDestination }).(FileDestinationResponsePtrOutput)
}

func (o ScheduledActionByScopeOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ScheduledActionByScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ScheduledActionByScopeOutput) Notification() NotificationPropertiesResponseOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) NotificationPropertiesResponseOutput { return v.Notification }).(NotificationPropertiesResponseOutput)
}

func (o ScheduledActionByScopeOutput) Schedule() SchedulePropertiesResponseOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) SchedulePropertiesResponseOutput { return v.Schedule }).(SchedulePropertiesResponseOutput)
}

func (o ScheduledActionByScopeOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o ScheduledActionByScopeOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o ScheduledActionByScopeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ScheduledActionByScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ScheduledActionByScopeOutput) ViewId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledActionByScope) pulumi.StringOutput { return v.ViewId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduledActionByScopeOutput{})
}
