


package v20180101

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
	// DataBox orders.
	ClassDiscriminatorDataBox = ClassDiscriminator("DataBox")
	// DataBoxDisk orders.
	ClassDiscriminatorDataBoxDisk = ClassDiscriminator("DataBoxDisk")
	// DataBoxHeavy orders.
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
	// Unknown type.
	DataDestinationTypeUnknownType = DataDestinationType("UnknownType")
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
	// DataBox.
	SkuNameDataBox = SkuName("DataBox")
	// DataBoxDisk.
	SkuNameDataBoxDisk = SkuName("DataBoxDisk")
	// DataBoxHeavy.
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

func init() {
	pulumi.RegisterOutputType(AddressTypeOutput{})
	pulumi.RegisterOutputType(AddressTypePtrOutput{})
	pulumi.RegisterOutputType(ClassDiscriminatorOutput{})
	pulumi.RegisterOutputType(ClassDiscriminatorPtrOutput{})
	pulumi.RegisterOutputType(DataDestinationTypeOutput{})
	pulumi.RegisterOutputType(DataDestinationTypePtrOutput{})
	pulumi.RegisterOutputType(NotificationStageNameOutput{})
	pulumi.RegisterOutputType(NotificationStageNamePtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
}
