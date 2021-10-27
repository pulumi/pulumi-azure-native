


package v20201201preview

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

type DoubleEncryptionStatus string

const (
	// Double encryption is disabled
	DoubleEncryptionStatusDisabled = DoubleEncryptionStatus("Disabled")
	// Double encryption is enabled
	DoubleEncryptionStatusEnabled = DoubleEncryptionStatus("Enabled")
)

func (DoubleEncryptionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*DoubleEncryptionStatus)(nil)).Elem()
}

func (e DoubleEncryptionStatus) ToDoubleEncryptionStatusOutput() DoubleEncryptionStatusOutput {
	return pulumi.ToOutput(e).(DoubleEncryptionStatusOutput)
}

func (e DoubleEncryptionStatus) ToDoubleEncryptionStatusOutputWithContext(ctx context.Context) DoubleEncryptionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DoubleEncryptionStatusOutput)
}

func (e DoubleEncryptionStatus) ToDoubleEncryptionStatusPtrOutput() DoubleEncryptionStatusPtrOutput {
	return e.ToDoubleEncryptionStatusPtrOutputWithContext(context.Background())
}

func (e DoubleEncryptionStatus) ToDoubleEncryptionStatusPtrOutputWithContext(ctx context.Context) DoubleEncryptionStatusPtrOutput {
	return DoubleEncryptionStatus(e).ToDoubleEncryptionStatusOutputWithContext(ctx).ToDoubleEncryptionStatusPtrOutputWithContext(ctx)
}

func (e DoubleEncryptionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DoubleEncryptionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DoubleEncryptionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DoubleEncryptionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DoubleEncryptionStatusOutput struct{ *pulumi.OutputState }

func (DoubleEncryptionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DoubleEncryptionStatus)(nil)).Elem()
}

func (o DoubleEncryptionStatusOutput) ToDoubleEncryptionStatusOutput() DoubleEncryptionStatusOutput {
	return o
}

func (o DoubleEncryptionStatusOutput) ToDoubleEncryptionStatusOutputWithContext(ctx context.Context) DoubleEncryptionStatusOutput {
	return o
}

func (o DoubleEncryptionStatusOutput) ToDoubleEncryptionStatusPtrOutput() DoubleEncryptionStatusPtrOutput {
	return o.ToDoubleEncryptionStatusPtrOutputWithContext(context.Background())
}

func (o DoubleEncryptionStatusOutput) ToDoubleEncryptionStatusPtrOutputWithContext(ctx context.Context) DoubleEncryptionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DoubleEncryptionStatus) *DoubleEncryptionStatus {
		return &v
	}).(DoubleEncryptionStatusPtrOutput)
}

func (o DoubleEncryptionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DoubleEncryptionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DoubleEncryptionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DoubleEncryptionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DoubleEncryptionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DoubleEncryptionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DoubleEncryptionStatusPtrOutput struct{ *pulumi.OutputState }

func (DoubleEncryptionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DoubleEncryptionStatus)(nil)).Elem()
}

func (o DoubleEncryptionStatusPtrOutput) ToDoubleEncryptionStatusPtrOutput() DoubleEncryptionStatusPtrOutput {
	return o
}

func (o DoubleEncryptionStatusPtrOutput) ToDoubleEncryptionStatusPtrOutputWithContext(ctx context.Context) DoubleEncryptionStatusPtrOutput {
	return o
}

func (o DoubleEncryptionStatusPtrOutput) Elem() DoubleEncryptionStatusOutput {
	return o.ApplyT(func(v *DoubleEncryptionStatus) DoubleEncryptionStatus {
		if v != nil {
			return *v
		}
		var ret DoubleEncryptionStatus
		return ret
	}).(DoubleEncryptionStatusOutput)
}

func (o DoubleEncryptionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DoubleEncryptionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DoubleEncryptionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DoubleEncryptionStatusInput interface {
	pulumi.Input

	ToDoubleEncryptionStatusOutput() DoubleEncryptionStatusOutput
	ToDoubleEncryptionStatusOutputWithContext(context.Context) DoubleEncryptionStatusOutput
}

var doubleEncryptionStatusPtrType = reflect.TypeOf((**DoubleEncryptionStatus)(nil)).Elem()

type DoubleEncryptionStatusPtrInput interface {
	pulumi.Input

	ToDoubleEncryptionStatusPtrOutput() DoubleEncryptionStatusPtrOutput
	ToDoubleEncryptionStatusPtrOutputWithContext(context.Context) DoubleEncryptionStatusPtrOutput
}

type doubleEncryptionStatusPtr string

func DoubleEncryptionStatusPtr(v string) DoubleEncryptionStatusPtrInput {
	return (*doubleEncryptionStatusPtr)(&v)
}

func (*doubleEncryptionStatusPtr) ElementType() reflect.Type {
	return doubleEncryptionStatusPtrType
}

func (in *doubleEncryptionStatusPtr) ToDoubleEncryptionStatusPtrOutput() DoubleEncryptionStatusPtrOutput {
	return pulumi.ToOutput(in).(DoubleEncryptionStatusPtrOutput)
}

func (in *doubleEncryptionStatusPtr) ToDoubleEncryptionStatusPtrOutputWithContext(ctx context.Context) DoubleEncryptionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DoubleEncryptionStatusPtrOutput)
}

type NotificationStageName string

const (
	// Notification at order item shipped from microsoft datacenter.
	NotificationStageNameShipped = NotificationStageName("Shipped")
	// Notification at order item delivered to customer.
	NotificationStageNameDelivered = NotificationStageName("Delivered")
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

type OrderItemType string

const (
	// Purchase OrderItem.
	OrderItemTypePurchase = OrderItemType("Purchase")
	// Rental OrderItem.
	OrderItemTypeRental = OrderItemType("Rental")
)

func (OrderItemType) ElementType() reflect.Type {
	return reflect.TypeOf((*OrderItemType)(nil)).Elem()
}

func (e OrderItemType) ToOrderItemTypeOutput() OrderItemTypeOutput {
	return pulumi.ToOutput(e).(OrderItemTypeOutput)
}

func (e OrderItemType) ToOrderItemTypeOutputWithContext(ctx context.Context) OrderItemTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OrderItemTypeOutput)
}

func (e OrderItemType) ToOrderItemTypePtrOutput() OrderItemTypePtrOutput {
	return e.ToOrderItemTypePtrOutputWithContext(context.Background())
}

func (e OrderItemType) ToOrderItemTypePtrOutputWithContext(ctx context.Context) OrderItemTypePtrOutput {
	return OrderItemType(e).ToOrderItemTypeOutputWithContext(ctx).ToOrderItemTypePtrOutputWithContext(ctx)
}

func (e OrderItemType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OrderItemType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OrderItemType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OrderItemType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OrderItemTypeOutput struct{ *pulumi.OutputState }

func (OrderItemTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrderItemType)(nil)).Elem()
}

func (o OrderItemTypeOutput) ToOrderItemTypeOutput() OrderItemTypeOutput {
	return o
}

func (o OrderItemTypeOutput) ToOrderItemTypeOutputWithContext(ctx context.Context) OrderItemTypeOutput {
	return o
}

func (o OrderItemTypeOutput) ToOrderItemTypePtrOutput() OrderItemTypePtrOutput {
	return o.ToOrderItemTypePtrOutputWithContext(context.Background())
}

func (o OrderItemTypeOutput) ToOrderItemTypePtrOutputWithContext(ctx context.Context) OrderItemTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OrderItemType) *OrderItemType {
		return &v
	}).(OrderItemTypePtrOutput)
}

func (o OrderItemTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OrderItemTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OrderItemType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OrderItemTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OrderItemTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OrderItemType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OrderItemTypePtrOutput struct{ *pulumi.OutputState }

func (OrderItemTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrderItemType)(nil)).Elem()
}

func (o OrderItemTypePtrOutput) ToOrderItemTypePtrOutput() OrderItemTypePtrOutput {
	return o
}

func (o OrderItemTypePtrOutput) ToOrderItemTypePtrOutputWithContext(ctx context.Context) OrderItemTypePtrOutput {
	return o
}

func (o OrderItemTypePtrOutput) Elem() OrderItemTypeOutput {
	return o.ApplyT(func(v *OrderItemType) OrderItemType {
		if v != nil {
			return *v
		}
		var ret OrderItemType
		return ret
	}).(OrderItemTypeOutput)
}

func (o OrderItemTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OrderItemTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OrderItemType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OrderItemTypeInput interface {
	pulumi.Input

	ToOrderItemTypeOutput() OrderItemTypeOutput
	ToOrderItemTypeOutputWithContext(context.Context) OrderItemTypeOutput
}

var orderItemTypePtrType = reflect.TypeOf((**OrderItemType)(nil)).Elem()

type OrderItemTypePtrInput interface {
	pulumi.Input

	ToOrderItemTypePtrOutput() OrderItemTypePtrOutput
	ToOrderItemTypePtrOutputWithContext(context.Context) OrderItemTypePtrOutput
}

type orderItemTypePtr string

func OrderItemTypePtr(v string) OrderItemTypePtrInput {
	return (*orderItemTypePtr)(&v)
}

func (*orderItemTypePtr) ElementType() reflect.Type {
	return orderItemTypePtrType
}

func (in *orderItemTypePtr) ToOrderItemTypePtrOutput() OrderItemTypePtrOutput {
	return pulumi.ToOutput(in).(OrderItemTypePtrOutput)
}

func (in *orderItemTypePtr) ToOrderItemTypePtrOutputWithContext(ctx context.Context) OrderItemTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OrderItemTypePtrOutput)
}

type SupportedFilterTypes string

const (
	// Ship to country
	SupportedFilterTypesShipToCountries = SupportedFilterTypes("ShipToCountries")
	// Double encryption status
	SupportedFilterTypesDoubleEncryptionStatus = SupportedFilterTypes("DoubleEncryptionStatus")
)

func (SupportedFilterTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportedFilterTypes)(nil)).Elem()
}

func (e SupportedFilterTypes) ToSupportedFilterTypesOutput() SupportedFilterTypesOutput {
	return pulumi.ToOutput(e).(SupportedFilterTypesOutput)
}

func (e SupportedFilterTypes) ToSupportedFilterTypesOutputWithContext(ctx context.Context) SupportedFilterTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SupportedFilterTypesOutput)
}

func (e SupportedFilterTypes) ToSupportedFilterTypesPtrOutput() SupportedFilterTypesPtrOutput {
	return e.ToSupportedFilterTypesPtrOutputWithContext(context.Background())
}

func (e SupportedFilterTypes) ToSupportedFilterTypesPtrOutputWithContext(ctx context.Context) SupportedFilterTypesPtrOutput {
	return SupportedFilterTypes(e).ToSupportedFilterTypesOutputWithContext(ctx).ToSupportedFilterTypesPtrOutputWithContext(ctx)
}

func (e SupportedFilterTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SupportedFilterTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SupportedFilterTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SupportedFilterTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SupportedFilterTypesOutput struct{ *pulumi.OutputState }

func (SupportedFilterTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportedFilterTypes)(nil)).Elem()
}

func (o SupportedFilterTypesOutput) ToSupportedFilterTypesOutput() SupportedFilterTypesOutput {
	return o
}

func (o SupportedFilterTypesOutput) ToSupportedFilterTypesOutputWithContext(ctx context.Context) SupportedFilterTypesOutput {
	return o
}

func (o SupportedFilterTypesOutput) ToSupportedFilterTypesPtrOutput() SupportedFilterTypesPtrOutput {
	return o.ToSupportedFilterTypesPtrOutputWithContext(context.Background())
}

func (o SupportedFilterTypesOutput) ToSupportedFilterTypesPtrOutputWithContext(ctx context.Context) SupportedFilterTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SupportedFilterTypes) *SupportedFilterTypes {
		return &v
	}).(SupportedFilterTypesPtrOutput)
}

func (o SupportedFilterTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SupportedFilterTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SupportedFilterTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SupportedFilterTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SupportedFilterTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SupportedFilterTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SupportedFilterTypesPtrOutput struct{ *pulumi.OutputState }

func (SupportedFilterTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SupportedFilterTypes)(nil)).Elem()
}

func (o SupportedFilterTypesPtrOutput) ToSupportedFilterTypesPtrOutput() SupportedFilterTypesPtrOutput {
	return o
}

func (o SupportedFilterTypesPtrOutput) ToSupportedFilterTypesPtrOutputWithContext(ctx context.Context) SupportedFilterTypesPtrOutput {
	return o
}

func (o SupportedFilterTypesPtrOutput) Elem() SupportedFilterTypesOutput {
	return o.ApplyT(func(v *SupportedFilterTypes) SupportedFilterTypes {
		if v != nil {
			return *v
		}
		var ret SupportedFilterTypes
		return ret
	}).(SupportedFilterTypesOutput)
}

func (o SupportedFilterTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SupportedFilterTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SupportedFilterTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SupportedFilterTypesInput interface {
	pulumi.Input

	ToSupportedFilterTypesOutput() SupportedFilterTypesOutput
	ToSupportedFilterTypesOutputWithContext(context.Context) SupportedFilterTypesOutput
}

var supportedFilterTypesPtrType = reflect.TypeOf((**SupportedFilterTypes)(nil)).Elem()

type SupportedFilterTypesPtrInput interface {
	pulumi.Input

	ToSupportedFilterTypesPtrOutput() SupportedFilterTypesPtrOutput
	ToSupportedFilterTypesPtrOutputWithContext(context.Context) SupportedFilterTypesPtrOutput
}

type supportedFilterTypesPtr string

func SupportedFilterTypesPtr(v string) SupportedFilterTypesPtrInput {
	return (*supportedFilterTypesPtr)(&v)
}

func (*supportedFilterTypesPtr) ElementType() reflect.Type {
	return supportedFilterTypesPtrType
}

func (in *supportedFilterTypesPtr) ToSupportedFilterTypesPtrOutput() SupportedFilterTypesPtrOutput {
	return pulumi.ToOutput(in).(SupportedFilterTypesPtrOutput)
}

func (in *supportedFilterTypesPtr) ToSupportedFilterTypesPtrOutputWithContext(ctx context.Context) SupportedFilterTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SupportedFilterTypesPtrOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*DoubleEncryptionStatusInput)(nil)).Elem(), DoubleEncryptionStatus("Disabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*DoubleEncryptionStatusPtrInput)(nil)).Elem(), DoubleEncryptionStatus("Disabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationStageNameInput)(nil)).Elem(), NotificationStageName("Shipped"))
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationStageNamePtrInput)(nil)).Elem(), NotificationStageName("Shipped"))
	pulumi.RegisterInputType(reflect.TypeOf((*OrderItemTypeInput)(nil)).Elem(), OrderItemType("Purchase"))
	pulumi.RegisterInputType(reflect.TypeOf((*OrderItemTypePtrInput)(nil)).Elem(), OrderItemType("Purchase"))
	pulumi.RegisterInputType(reflect.TypeOf((*SupportedFilterTypesInput)(nil)).Elem(), SupportedFilterTypes("ShipToCountries"))
	pulumi.RegisterInputType(reflect.TypeOf((*SupportedFilterTypesPtrInput)(nil)).Elem(), SupportedFilterTypes("ShipToCountries"))
	pulumi.RegisterInputType(reflect.TypeOf((*TransportShipmentTypesInput)(nil)).Elem(), TransportShipmentTypes("CustomerManaged"))
	pulumi.RegisterInputType(reflect.TypeOf((*TransportShipmentTypesPtrInput)(nil)).Elem(), TransportShipmentTypes("CustomerManaged"))
	pulumi.RegisterOutputType(AddressTypeOutput{})
	pulumi.RegisterOutputType(AddressTypePtrOutput{})
	pulumi.RegisterOutputType(DoubleEncryptionStatusOutput{})
	pulumi.RegisterOutputType(DoubleEncryptionStatusPtrOutput{})
	pulumi.RegisterOutputType(NotificationStageNameOutput{})
	pulumi.RegisterOutputType(NotificationStageNamePtrOutput{})
	pulumi.RegisterOutputType(OrderItemTypeOutput{})
	pulumi.RegisterOutputType(OrderItemTypePtrOutput{})
	pulumi.RegisterOutputType(SupportedFilterTypesOutput{})
	pulumi.RegisterOutputType(SupportedFilterTypesPtrOutput{})
	pulumi.RegisterOutputType(TransportShipmentTypesOutput{})
	pulumi.RegisterOutputType(TransportShipmentTypesPtrOutput{})
}
