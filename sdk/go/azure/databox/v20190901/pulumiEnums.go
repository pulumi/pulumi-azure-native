


package v20190901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AddressType string

const (
	// Address type not known.
	AddressTypeNone = AddressType("None")
	// Residential Address.
	AddressTypeResidential = AddressType("Residential")
	// Commercial Address.
	AddressTypeCommercial = AddressType("Commercial")
)

func (AddressType) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressType)(nil)).Elem()
}

func (e AddressType) ToAddressTypeOutput() AddressTypeOutput {
	return pulumi.ToOutput(e).(AddressTypeOutput)
}

func (e AddressType) ToAddressTypeOutputWithContext(ctx context.Context) AddressTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AddressTypeOutput)
}

func (e AddressType) ToAddressTypePtrOutput() AddressTypePtrOutput {
	return e.ToAddressTypePtrOutputWithContext(context.Background())
}

func (e AddressType) ToAddressTypePtrOutputWithContext(ctx context.Context) AddressTypePtrOutput {
	return AddressType(e).ToAddressTypeOutputWithContext(ctx).ToAddressTypePtrOutputWithContext(ctx)
}

func (e AddressType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AddressType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AddressType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AddressType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AddressTypeOutput struct{ *pulumi.OutputState }

func (AddressTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressType)(nil)).Elem()
}

func (o AddressTypeOutput) ToAddressTypeOutput() AddressTypeOutput {
	return o
}

func (o AddressTypeOutput) ToAddressTypeOutputWithContext(ctx context.Context) AddressTypeOutput {
	return o
}

func (o AddressTypeOutput) ToAddressTypePtrOutput() AddressTypePtrOutput {
	return o.ToAddressTypePtrOutputWithContext(context.Background())
}

func (o AddressTypeOutput) ToAddressTypePtrOutputWithContext(ctx context.Context) AddressTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AddressType) *AddressType {
		return &v
	}).(AddressTypePtrOutput)
}

func (o AddressTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AddressTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AddressType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AddressTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AddressTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AddressType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AddressTypePtrOutput struct{ *pulumi.OutputState }

func (AddressTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressType)(nil)).Elem()
}

func (o AddressTypePtrOutput) ToAddressTypePtrOutput() AddressTypePtrOutput {
	return o
}

func (o AddressTypePtrOutput) ToAddressTypePtrOutputWithContext(ctx context.Context) AddressTypePtrOutput {
	return o
}

func (o AddressTypePtrOutput) Elem() AddressTypeOutput {
	return o.ApplyT(func(v *AddressType) AddressType {
		if v != nil {
			return *v
		}
		var ret AddressType
		return ret
	}).(AddressTypeOutput)
}

func (o AddressTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AddressTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AddressType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AddressTypeInput interface {
	pulumi.Input

	ToAddressTypeOutput() AddressTypeOutput
	ToAddressTypeOutputWithContext(context.Context) AddressTypeOutput
}

var addressTypePtrType = reflect.TypeOf((**AddressType)(nil)).Elem()

type AddressTypePtrInput interface {
	pulumi.Input

	ToAddressTypePtrOutput() AddressTypePtrOutput
	ToAddressTypePtrOutputWithContext(context.Context) AddressTypePtrOutput
}

type addressTypePtr string

func AddressTypePtr(v string) AddressTypePtrInput {
	return (*addressTypePtr)(&v)
}

func (*addressTypePtr) ElementType() reflect.Type {
	return addressTypePtrType
}

func (in *addressTypePtr) ToAddressTypePtrOutput() AddressTypePtrOutput {
	return pulumi.ToOutput(in).(AddressTypePtrOutput)
}

func (in *addressTypePtr) ToAddressTypePtrOutputWithContext(ctx context.Context) AddressTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AddressTypePtrOutput)
}

type ClassDiscriminator string

const (
	// Databox orders.
	ClassDiscriminatorDataBox = ClassDiscriminator("DataBox")
	// DataboxDisk orders.
	ClassDiscriminatorDataBoxDisk = ClassDiscriminator("DataBoxDisk")
	// DataboxHeavy orders.
	ClassDiscriminatorDataBoxHeavy = ClassDiscriminator("DataBoxHeavy")
)

func (ClassDiscriminator) ElementType() reflect.Type {
	return reflect.TypeOf((*ClassDiscriminator)(nil)).Elem()
}

func (e ClassDiscriminator) ToClassDiscriminatorOutput() ClassDiscriminatorOutput {
	return pulumi.ToOutput(e).(ClassDiscriminatorOutput)
}

func (e ClassDiscriminator) ToClassDiscriminatorOutputWithContext(ctx context.Context) ClassDiscriminatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClassDiscriminatorOutput)
}

func (e ClassDiscriminator) ToClassDiscriminatorPtrOutput() ClassDiscriminatorPtrOutput {
	return e.ToClassDiscriminatorPtrOutputWithContext(context.Background())
}

func (e ClassDiscriminator) ToClassDiscriminatorPtrOutputWithContext(ctx context.Context) ClassDiscriminatorPtrOutput {
	return ClassDiscriminator(e).ToClassDiscriminatorOutputWithContext(ctx).ToClassDiscriminatorPtrOutputWithContext(ctx)
}

func (e ClassDiscriminator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClassDiscriminator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClassDiscriminator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClassDiscriminator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClassDiscriminatorOutput struct{ *pulumi.OutputState }

func (ClassDiscriminatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClassDiscriminator)(nil)).Elem()
}

func (o ClassDiscriminatorOutput) ToClassDiscriminatorOutput() ClassDiscriminatorOutput {
	return o
}

func (o ClassDiscriminatorOutput) ToClassDiscriminatorOutputWithContext(ctx context.Context) ClassDiscriminatorOutput {
	return o
}

func (o ClassDiscriminatorOutput) ToClassDiscriminatorPtrOutput() ClassDiscriminatorPtrOutput {
	return o.ToClassDiscriminatorPtrOutputWithContext(context.Background())
}

func (o ClassDiscriminatorOutput) ToClassDiscriminatorPtrOutputWithContext(ctx context.Context) ClassDiscriminatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClassDiscriminator) *ClassDiscriminator {
		return &v
	}).(ClassDiscriminatorPtrOutput)
}

func (o ClassDiscriminatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClassDiscriminatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClassDiscriminator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClassDiscriminatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClassDiscriminatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClassDiscriminator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClassDiscriminatorPtrOutput struct{ *pulumi.OutputState }

func (ClassDiscriminatorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClassDiscriminator)(nil)).Elem()
}

func (o ClassDiscriminatorPtrOutput) ToClassDiscriminatorPtrOutput() ClassDiscriminatorPtrOutput {
	return o
}

func (o ClassDiscriminatorPtrOutput) ToClassDiscriminatorPtrOutputWithContext(ctx context.Context) ClassDiscriminatorPtrOutput {
	return o
}

func (o ClassDiscriminatorPtrOutput) Elem() ClassDiscriminatorOutput {
	return o.ApplyT(func(v *ClassDiscriminator) ClassDiscriminator {
		if v != nil {
			return *v
		}
		var ret ClassDiscriminator
		return ret
	}).(ClassDiscriminatorOutput)
}

func (o ClassDiscriminatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClassDiscriminatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClassDiscriminator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ClassDiscriminatorInput interface {
	pulumi.Input

	ToClassDiscriminatorOutput() ClassDiscriminatorOutput
	ToClassDiscriminatorOutputWithContext(context.Context) ClassDiscriminatorOutput
}

var classDiscriminatorPtrType = reflect.TypeOf((**ClassDiscriminator)(nil)).Elem()

type ClassDiscriminatorPtrInput interface {
	pulumi.Input

	ToClassDiscriminatorPtrOutput() ClassDiscriminatorPtrOutput
	ToClassDiscriminatorPtrOutputWithContext(context.Context) ClassDiscriminatorPtrOutput
}

type classDiscriminatorPtr string

func ClassDiscriminatorPtr(v string) ClassDiscriminatorPtrInput {
	return (*classDiscriminatorPtr)(&v)
}

func (*classDiscriminatorPtr) ElementType() reflect.Type {
	return classDiscriminatorPtrType
}

func (in *classDiscriminatorPtr) ToClassDiscriminatorPtrOutput() ClassDiscriminatorPtrOutput {
	return pulumi.ToOutput(in).(ClassDiscriminatorPtrOutput)
}

func (in *classDiscriminatorPtr) ToClassDiscriminatorPtrOutputWithContext(ctx context.Context) ClassDiscriminatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClassDiscriminatorPtrOutput)
}

type DataDestinationType string

const (
	// Storage Accounts .
	DataDestinationTypeStorageAccount = DataDestinationType("StorageAccount")
	// Azure Managed disk storage.
	DataDestinationTypeManagedDisk = DataDestinationType("ManagedDisk")
)

func (DataDestinationType) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDestinationType)(nil)).Elem()
}

func (e DataDestinationType) ToDataDestinationTypeOutput() DataDestinationTypeOutput {
	return pulumi.ToOutput(e).(DataDestinationTypeOutput)
}

func (e DataDestinationType) ToDataDestinationTypeOutputWithContext(ctx context.Context) DataDestinationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataDestinationTypeOutput)
}

func (e DataDestinationType) ToDataDestinationTypePtrOutput() DataDestinationTypePtrOutput {
	return e.ToDataDestinationTypePtrOutputWithContext(context.Background())
}

func (e DataDestinationType) ToDataDestinationTypePtrOutputWithContext(ctx context.Context) DataDestinationTypePtrOutput {
	return DataDestinationType(e).ToDataDestinationTypeOutputWithContext(ctx).ToDataDestinationTypePtrOutputWithContext(ctx)
}

func (e DataDestinationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataDestinationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataDestinationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataDestinationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataDestinationTypeOutput struct{ *pulumi.OutputState }

func (DataDestinationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDestinationType)(nil)).Elem()
}

func (o DataDestinationTypeOutput) ToDataDestinationTypeOutput() DataDestinationTypeOutput {
	return o
}

func (o DataDestinationTypeOutput) ToDataDestinationTypeOutputWithContext(ctx context.Context) DataDestinationTypeOutput {
	return o
}

func (o DataDestinationTypeOutput) ToDataDestinationTypePtrOutput() DataDestinationTypePtrOutput {
	return o.ToDataDestinationTypePtrOutputWithContext(context.Background())
}

func (o DataDestinationTypeOutput) ToDataDestinationTypePtrOutputWithContext(ctx context.Context) DataDestinationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataDestinationType) *DataDestinationType {
		return &v
	}).(DataDestinationTypePtrOutput)
}

func (o DataDestinationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataDestinationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataDestinationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataDestinationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataDestinationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataDestinationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataDestinationTypePtrOutput struct{ *pulumi.OutputState }

func (DataDestinationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataDestinationType)(nil)).Elem()
}

func (o DataDestinationTypePtrOutput) ToDataDestinationTypePtrOutput() DataDestinationTypePtrOutput {
	return o
}

func (o DataDestinationTypePtrOutput) ToDataDestinationTypePtrOutputWithContext(ctx context.Context) DataDestinationTypePtrOutput {
	return o
}

func (o DataDestinationTypePtrOutput) Elem() DataDestinationTypeOutput {
	return o.ApplyT(func(v *DataDestinationType) DataDestinationType {
		if v != nil {
			return *v
		}
		var ret DataDestinationType
		return ret
	}).(DataDestinationTypeOutput)
}

func (o DataDestinationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataDestinationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataDestinationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DataDestinationTypeInput interface {
	pulumi.Input

	ToDataDestinationTypeOutput() DataDestinationTypeOutput
	ToDataDestinationTypeOutputWithContext(context.Context) DataDestinationTypeOutput
}

var dataDestinationTypePtrType = reflect.TypeOf((**DataDestinationType)(nil)).Elem()

type DataDestinationTypePtrInput interface {
	pulumi.Input

	ToDataDestinationTypePtrOutput() DataDestinationTypePtrOutput
	ToDataDestinationTypePtrOutputWithContext(context.Context) DataDestinationTypePtrOutput
}

type dataDestinationTypePtr string

func DataDestinationTypePtr(v string) DataDestinationTypePtrInput {
	return (*dataDestinationTypePtr)(&v)
}

func (*dataDestinationTypePtr) ElementType() reflect.Type {
	return dataDestinationTypePtrType
}

func (in *dataDestinationTypePtr) ToDataDestinationTypePtrOutput() DataDestinationTypePtrOutput {
	return pulumi.ToOutput(in).(DataDestinationTypePtrOutput)
}

func (in *dataDestinationTypePtr) ToDataDestinationTypePtrOutputWithContext(ctx context.Context) DataDestinationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataDestinationTypePtrOutput)
}

type JobDeliveryType string

const (
	// Non Scheduled job.
	JobDeliveryTypeNonScheduled = JobDeliveryType("NonScheduled")
	// Scheduled job.
	JobDeliveryTypeScheduled = JobDeliveryType("Scheduled")
)

func (JobDeliveryType) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDeliveryType)(nil)).Elem()
}

func (e JobDeliveryType) ToJobDeliveryTypeOutput() JobDeliveryTypeOutput {
	return pulumi.ToOutput(e).(JobDeliveryTypeOutput)
}

func (e JobDeliveryType) ToJobDeliveryTypeOutputWithContext(ctx context.Context) JobDeliveryTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(JobDeliveryTypeOutput)
}

func (e JobDeliveryType) ToJobDeliveryTypePtrOutput() JobDeliveryTypePtrOutput {
	return e.ToJobDeliveryTypePtrOutputWithContext(context.Background())
}

func (e JobDeliveryType) ToJobDeliveryTypePtrOutputWithContext(ctx context.Context) JobDeliveryTypePtrOutput {
	return JobDeliveryType(e).ToJobDeliveryTypeOutputWithContext(ctx).ToJobDeliveryTypePtrOutputWithContext(ctx)
}

func (e JobDeliveryType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobDeliveryType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobDeliveryType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobDeliveryType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type JobDeliveryTypeOutput struct{ *pulumi.OutputState }

func (JobDeliveryTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDeliveryType)(nil)).Elem()
}

func (o JobDeliveryTypeOutput) ToJobDeliveryTypeOutput() JobDeliveryTypeOutput {
	return o
}

func (o JobDeliveryTypeOutput) ToJobDeliveryTypeOutputWithContext(ctx context.Context) JobDeliveryTypeOutput {
	return o
}

func (o JobDeliveryTypeOutput) ToJobDeliveryTypePtrOutput() JobDeliveryTypePtrOutput {
	return o.ToJobDeliveryTypePtrOutputWithContext(context.Background())
}

func (o JobDeliveryTypeOutput) ToJobDeliveryTypePtrOutputWithContext(ctx context.Context) JobDeliveryTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobDeliveryType) *JobDeliveryType {
		return &v
	}).(JobDeliveryTypePtrOutput)
}

func (o JobDeliveryTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o JobDeliveryTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JobDeliveryType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o JobDeliveryTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JobDeliveryTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JobDeliveryType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type JobDeliveryTypePtrOutput struct{ *pulumi.OutputState }

func (JobDeliveryTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDeliveryType)(nil)).Elem()
}

func (o JobDeliveryTypePtrOutput) ToJobDeliveryTypePtrOutput() JobDeliveryTypePtrOutput {
	return o
}

func (o JobDeliveryTypePtrOutput) ToJobDeliveryTypePtrOutputWithContext(ctx context.Context) JobDeliveryTypePtrOutput {
	return o
}

func (o JobDeliveryTypePtrOutput) Elem() JobDeliveryTypeOutput {
	return o.ApplyT(func(v *JobDeliveryType) JobDeliveryType {
		if v != nil {
			return *v
		}
		var ret JobDeliveryType
		return ret
	}).(JobDeliveryTypeOutput)
}

func (o JobDeliveryTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JobDeliveryTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *JobDeliveryType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type JobDeliveryTypeInput interface {
	pulumi.Input

	ToJobDeliveryTypeOutput() JobDeliveryTypeOutput
	ToJobDeliveryTypeOutputWithContext(context.Context) JobDeliveryTypeOutput
}

var jobDeliveryTypePtrType = reflect.TypeOf((**JobDeliveryType)(nil)).Elem()

type JobDeliveryTypePtrInput interface {
	pulumi.Input

	ToJobDeliveryTypePtrOutput() JobDeliveryTypePtrOutput
	ToJobDeliveryTypePtrOutputWithContext(context.Context) JobDeliveryTypePtrOutput
}

type jobDeliveryTypePtr string

func JobDeliveryTypePtr(v string) JobDeliveryTypePtrInput {
	return (*jobDeliveryTypePtr)(&v)
}

func (*jobDeliveryTypePtr) ElementType() reflect.Type {
	return jobDeliveryTypePtrType
}

func (in *jobDeliveryTypePtr) ToJobDeliveryTypePtrOutput() JobDeliveryTypePtrOutput {
	return pulumi.ToOutput(in).(JobDeliveryTypePtrOutput)
}

func (in *jobDeliveryTypePtr) ToJobDeliveryTypePtrOutputWithContext(ctx context.Context) JobDeliveryTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(JobDeliveryTypePtrOutput)
}

type NotificationStageName string

const (
	// Notification at device prepared stage.
	NotificationStageNameDevicePrepared = NotificationStageName("DevicePrepared")
	// Notification at device dispatched stage.
	NotificationStageNameDispatched = NotificationStageName("Dispatched")
	// Notification at device delivered stage.
	NotificationStageNameDelivered = NotificationStageName("Delivered")
	// Notification at device picked up from user stage.
	NotificationStageNamePickedUp = NotificationStageName("PickedUp")
	// Notification at device received at azure datacenter stage.
	NotificationStageNameAtAzureDC = NotificationStageName("AtAzureDC")
	// Notification at data copy started stage.
	NotificationStageNameDataCopy = NotificationStageName("DataCopy")
)

func (NotificationStageName) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationStageName)(nil)).Elem()
}

func (e NotificationStageName) ToNotificationStageNameOutput() NotificationStageNameOutput {
	return pulumi.ToOutput(e).(NotificationStageNameOutput)
}

func (e NotificationStageName) ToNotificationStageNameOutputWithContext(ctx context.Context) NotificationStageNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NotificationStageNameOutput)
}

func (e NotificationStageName) ToNotificationStageNamePtrOutput() NotificationStageNamePtrOutput {
	return e.ToNotificationStageNamePtrOutputWithContext(context.Background())
}

func (e NotificationStageName) ToNotificationStageNamePtrOutputWithContext(ctx context.Context) NotificationStageNamePtrOutput {
	return NotificationStageName(e).ToNotificationStageNameOutputWithContext(ctx).ToNotificationStageNamePtrOutputWithContext(ctx)
}

func (e NotificationStageName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NotificationStageName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NotificationStageName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NotificationStageName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NotificationStageNameOutput struct{ *pulumi.OutputState }

func (NotificationStageNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationStageName)(nil)).Elem()
}

func (o NotificationStageNameOutput) ToNotificationStageNameOutput() NotificationStageNameOutput {
	return o
}

func (o NotificationStageNameOutput) ToNotificationStageNameOutputWithContext(ctx context.Context) NotificationStageNameOutput {
	return o
}

func (o NotificationStageNameOutput) ToNotificationStageNamePtrOutput() NotificationStageNamePtrOutput {
	return o.ToNotificationStageNamePtrOutputWithContext(context.Background())
}

func (o NotificationStageNameOutput) ToNotificationStageNamePtrOutputWithContext(ctx context.Context) NotificationStageNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotificationStageName) *NotificationStageName {
		return &v
	}).(NotificationStageNamePtrOutput)
}

func (o NotificationStageNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NotificationStageNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NotificationStageName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NotificationStageNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NotificationStageNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NotificationStageName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NotificationStageNamePtrOutput struct{ *pulumi.OutputState }

func (NotificationStageNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationStageName)(nil)).Elem()
}

func (o NotificationStageNamePtrOutput) ToNotificationStageNamePtrOutput() NotificationStageNamePtrOutput {
	return o
}

func (o NotificationStageNamePtrOutput) ToNotificationStageNamePtrOutputWithContext(ctx context.Context) NotificationStageNamePtrOutput {
	return o
}

func (o NotificationStageNamePtrOutput) Elem() NotificationStageNameOutput {
	return o.ApplyT(func(v *NotificationStageName) NotificationStageName {
		if v != nil {
			return *v
		}
		var ret NotificationStageName
		return ret
	}).(NotificationStageNameOutput)
}

func (o NotificationStageNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NotificationStageNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NotificationStageName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NotificationStageNameInput interface {
	pulumi.Input

	ToNotificationStageNameOutput() NotificationStageNameOutput
	ToNotificationStageNameOutputWithContext(context.Context) NotificationStageNameOutput
}

var notificationStageNamePtrType = reflect.TypeOf((**NotificationStageName)(nil)).Elem()

type NotificationStageNamePtrInput interface {
	pulumi.Input

	ToNotificationStageNamePtrOutput() NotificationStageNamePtrOutput
	ToNotificationStageNamePtrOutputWithContext(context.Context) NotificationStageNamePtrOutput
}

type notificationStageNamePtr string

func NotificationStageNamePtr(v string) NotificationStageNamePtrInput {
	return (*notificationStageNamePtr)(&v)
}

func (*notificationStageNamePtr) ElementType() reflect.Type {
	return notificationStageNamePtrType
}

func (in *notificationStageNamePtr) ToNotificationStageNamePtrOutput() NotificationStageNamePtrOutput {
	return pulumi.ToOutput(in).(NotificationStageNamePtrOutput)
}

func (in *notificationStageNamePtr) ToNotificationStageNamePtrOutputWithContext(ctx context.Context) NotificationStageNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NotificationStageNamePtrOutput)
}

type SkuName string

const (
	// Databox.
	SkuNameDataBox = SkuName("DataBox")
	// DataboxDisk.
	SkuNameDataBoxDisk = SkuName("DataBoxDisk")
	// DataboxHeavy.
	SkuNameDataBoxHeavy = SkuName("DataBoxHeavy")
)

func (SkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (e SkuName) ToSkuNameOutput() SkuNameOutput {
	return pulumi.ToOutput(e).(SkuNameOutput)
}

func (e SkuName) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuNameOutput)
}

func (e SkuName) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return e.ToSkuNamePtrOutputWithContext(context.Background())
}

func (e SkuName) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return SkuName(e).ToSkuNameOutputWithContext(ctx).ToSkuNamePtrOutputWithContext(ctx)
}

func (e SkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuNameOutput struct{ *pulumi.OutputState }

func (SkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (o SkuNameOutput) ToSkuNameOutput() SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o.ToSkuNamePtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuName) *SkuName {
		return &v
	}).(SkuNamePtrOutput)
}

func (o SkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuNamePtrOutput struct{ *pulumi.OutputState }

func (SkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuName)(nil)).Elem()
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) Elem() SkuNameOutput {
	return o.ApplyT(func(v *SkuName) SkuName {
		if v != nil {
			return *v
		}
		var ret SkuName
		return ret
	}).(SkuNameOutput)
}

func (o SkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuNameInput interface {
	pulumi.Input

	ToSkuNameOutput() SkuNameOutput
	ToSkuNameOutputWithContext(context.Context) SkuNameOutput
}

var skuNamePtrType = reflect.TypeOf((**SkuName)(nil)).Elem()

type SkuNamePtrInput interface {
	pulumi.Input

	ToSkuNamePtrOutput() SkuNamePtrOutput
	ToSkuNamePtrOutputWithContext(context.Context) SkuNamePtrOutput
}

type skuNamePtr string

func SkuNamePtr(v string) SkuNamePtrInput {
	return (*skuNamePtr)(&v)
}

func (*skuNamePtr) ElementType() reflect.Type {
	return skuNamePtrType
}

func (in *skuNamePtr) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return pulumi.ToOutput(in).(SkuNamePtrOutput)
}

func (in *skuNamePtr) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuNamePtrOutput)
}

type TransportShipmentTypes string

const (
	// Shipment Logistics is handled by the customer.
	TransportShipmentTypesCustomerManaged = TransportShipmentTypes("CustomerManaged")
	// Shipment Logistics is handled by Microsoft.
	TransportShipmentTypesMicrosoftManaged = TransportShipmentTypes("MicrosoftManaged")
)

func (TransportShipmentTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*TransportShipmentTypes)(nil)).Elem()
}

func (e TransportShipmentTypes) ToTransportShipmentTypesOutput() TransportShipmentTypesOutput {
	return pulumi.ToOutput(e).(TransportShipmentTypesOutput)
}

func (e TransportShipmentTypes) ToTransportShipmentTypesOutputWithContext(ctx context.Context) TransportShipmentTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TransportShipmentTypesOutput)
}

func (e TransportShipmentTypes) ToTransportShipmentTypesPtrOutput() TransportShipmentTypesPtrOutput {
	return e.ToTransportShipmentTypesPtrOutputWithContext(context.Background())
}

func (e TransportShipmentTypes) ToTransportShipmentTypesPtrOutputWithContext(ctx context.Context) TransportShipmentTypesPtrOutput {
	return TransportShipmentTypes(e).ToTransportShipmentTypesOutputWithContext(ctx).ToTransportShipmentTypesPtrOutputWithContext(ctx)
}

func (e TransportShipmentTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TransportShipmentTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TransportShipmentTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TransportShipmentTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TransportShipmentTypesOutput struct{ *pulumi.OutputState }

func (TransportShipmentTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransportShipmentTypes)(nil)).Elem()
}

func (o TransportShipmentTypesOutput) ToTransportShipmentTypesOutput() TransportShipmentTypesOutput {
	return o
}

func (o TransportShipmentTypesOutput) ToTransportShipmentTypesOutputWithContext(ctx context.Context) TransportShipmentTypesOutput {
	return o
}

func (o TransportShipmentTypesOutput) ToTransportShipmentTypesPtrOutput() TransportShipmentTypesPtrOutput {
	return o.ToTransportShipmentTypesPtrOutputWithContext(context.Background())
}

func (o TransportShipmentTypesOutput) ToTransportShipmentTypesPtrOutputWithContext(ctx context.Context) TransportShipmentTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TransportShipmentTypes) *TransportShipmentTypes {
		return &v
	}).(TransportShipmentTypesPtrOutput)
}

func (o TransportShipmentTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TransportShipmentTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TransportShipmentTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TransportShipmentTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TransportShipmentTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TransportShipmentTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TransportShipmentTypesPtrOutput struct{ *pulumi.OutputState }

func (TransportShipmentTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransportShipmentTypes)(nil)).Elem()
}

func (o TransportShipmentTypesPtrOutput) ToTransportShipmentTypesPtrOutput() TransportShipmentTypesPtrOutput {
	return o
}

func (o TransportShipmentTypesPtrOutput) ToTransportShipmentTypesPtrOutputWithContext(ctx context.Context) TransportShipmentTypesPtrOutput {
	return o
}

func (o TransportShipmentTypesPtrOutput) Elem() TransportShipmentTypesOutput {
	return o.ApplyT(func(v *TransportShipmentTypes) TransportShipmentTypes {
		if v != nil {
			return *v
		}
		var ret TransportShipmentTypes
		return ret
	}).(TransportShipmentTypesOutput)
}

func (o TransportShipmentTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TransportShipmentTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TransportShipmentTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TransportShipmentTypesInput interface {
	pulumi.Input

	ToTransportShipmentTypesOutput() TransportShipmentTypesOutput
	ToTransportShipmentTypesOutputWithContext(context.Context) TransportShipmentTypesOutput
}

var transportShipmentTypesPtrType = reflect.TypeOf((**TransportShipmentTypes)(nil)).Elem()

type TransportShipmentTypesPtrInput interface {
	pulumi.Input

	ToTransportShipmentTypesPtrOutput() TransportShipmentTypesPtrOutput
	ToTransportShipmentTypesPtrOutputWithContext(context.Context) TransportShipmentTypesPtrOutput
}

type transportShipmentTypesPtr string

func TransportShipmentTypesPtr(v string) TransportShipmentTypesPtrInput {
	return (*transportShipmentTypesPtr)(&v)
}

func (*transportShipmentTypesPtr) ElementType() reflect.Type {
	return transportShipmentTypesPtrType
}

func (in *transportShipmentTypesPtr) ToTransportShipmentTypesPtrOutput() TransportShipmentTypesPtrOutput {
	return pulumi.ToOutput(in).(TransportShipmentTypesPtrOutput)
}

func (in *transportShipmentTypesPtr) ToTransportShipmentTypesPtrOutputWithContext(ctx context.Context) TransportShipmentTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TransportShipmentTypesPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AddressTypeInput)(nil)).Elem(), AddressType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*AddressTypePtrInput)(nil)).Elem(), AddressType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*ClassDiscriminatorInput)(nil)).Elem(), ClassDiscriminator("DataBox"))
	pulumi.RegisterInputType(reflect.TypeOf((*ClassDiscriminatorPtrInput)(nil)).Elem(), ClassDiscriminator("DataBox"))
	pulumi.RegisterInputType(reflect.TypeOf((*DataDestinationTypeInput)(nil)).Elem(), DataDestinationType("StorageAccount"))
	pulumi.RegisterInputType(reflect.TypeOf((*DataDestinationTypePtrInput)(nil)).Elem(), DataDestinationType("StorageAccount"))
	pulumi.RegisterInputType(reflect.TypeOf((*JobDeliveryTypeInput)(nil)).Elem(), JobDeliveryType("NonScheduled"))
	pulumi.RegisterInputType(reflect.TypeOf((*JobDeliveryTypePtrInput)(nil)).Elem(), JobDeliveryType("NonScheduled"))
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationStageNameInput)(nil)).Elem(), NotificationStageName("DevicePrepared"))
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationStageNamePtrInput)(nil)).Elem(), NotificationStageName("DevicePrepared"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuNameInput)(nil)).Elem(), SkuName("DataBox"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuNamePtrInput)(nil)).Elem(), SkuName("DataBox"))
	pulumi.RegisterInputType(reflect.TypeOf((*TransportShipmentTypesInput)(nil)).Elem(), TransportShipmentTypes("CustomerManaged"))
	pulumi.RegisterInputType(reflect.TypeOf((*TransportShipmentTypesPtrInput)(nil)).Elem(), TransportShipmentTypes("CustomerManaged"))
	pulumi.RegisterOutputType(AddressTypeOutput{})
	pulumi.RegisterOutputType(AddressTypePtrOutput{})
	pulumi.RegisterOutputType(ClassDiscriminatorOutput{})
	pulumi.RegisterOutputType(ClassDiscriminatorPtrOutput{})
	pulumi.RegisterOutputType(DataDestinationTypeOutput{})
	pulumi.RegisterOutputType(DataDestinationTypePtrOutput{})
	pulumi.RegisterOutputType(JobDeliveryTypeOutput{})
	pulumi.RegisterOutputType(JobDeliveryTypePtrOutput{})
	pulumi.RegisterOutputType(NotificationStageNameOutput{})
	pulumi.RegisterOutputType(NotificationStageNamePtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
	pulumi.RegisterOutputType(TransportShipmentTypesOutput{})
	pulumi.RegisterOutputType(TransportShipmentTypesPtrOutput{})
}
