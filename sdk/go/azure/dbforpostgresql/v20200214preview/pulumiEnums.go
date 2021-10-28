


package v20200214preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CreateMode string

const (
	CreateModeDefault            = CreateMode("Default")
	CreateModePointInTimeRestore = CreateMode("PointInTimeRestore")
)

func (CreateMode) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateMode)(nil)).Elem()
}

func (e CreateMode) ToCreateModeOutput() CreateModeOutput {
	return pulumi.ToOutput(e).(CreateModeOutput)
}

func (e CreateMode) ToCreateModeOutputWithContext(ctx context.Context) CreateModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CreateModeOutput)
}

func (e CreateMode) ToCreateModePtrOutput() CreateModePtrOutput {
	return e.ToCreateModePtrOutputWithContext(context.Background())
}

func (e CreateMode) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return CreateMode(e).ToCreateModeOutputWithContext(ctx).ToCreateModePtrOutputWithContext(ctx)
}

func (e CreateMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CreateMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CreateMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CreateMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CreateModeOutput struct{ *pulumi.OutputState }

func (CreateModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateMode)(nil)).Elem()
}

func (o CreateModeOutput) ToCreateModeOutput() CreateModeOutput {
	return o
}

func (o CreateModeOutput) ToCreateModeOutputWithContext(ctx context.Context) CreateModeOutput {
	return o
}

func (o CreateModeOutput) ToCreateModePtrOutput() CreateModePtrOutput {
	return o.ToCreateModePtrOutputWithContext(context.Background())
}

func (o CreateModeOutput) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreateMode) *CreateMode {
		return &v
	}).(CreateModePtrOutput)
}

func (o CreateModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CreateModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CreateMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CreateModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CreateModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CreateMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CreateModePtrOutput struct{ *pulumi.OutputState }

func (CreateModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateMode)(nil)).Elem()
}

func (o CreateModePtrOutput) ToCreateModePtrOutput() CreateModePtrOutput {
	return o
}

func (o CreateModePtrOutput) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return o
}

func (o CreateModePtrOutput) Elem() CreateModeOutput {
	return o.ApplyT(func(v *CreateMode) CreateMode {
		if v != nil {
			return *v
		}
		var ret CreateMode
		return ret
	}).(CreateModeOutput)
}

func (o CreateModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CreateModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CreateMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CreateModeInput interface {
	pulumi.Input

	ToCreateModeOutput() CreateModeOutput
	ToCreateModeOutputWithContext(context.Context) CreateModeOutput
}

var createModePtrType = reflect.TypeOf((**CreateMode)(nil)).Elem()

type CreateModePtrInput interface {
	pulumi.Input

	ToCreateModePtrOutput() CreateModePtrOutput
	ToCreateModePtrOutputWithContext(context.Context) CreateModePtrOutput
}

type createModePtr string

func CreateModePtr(v string) CreateModePtrInput {
	return (*createModePtr)(&v)
}

func (*createModePtr) ElementType() reflect.Type {
	return createModePtrType
}

func (in *createModePtr) ToCreateModePtrOutput() CreateModePtrOutput {
	return pulumi.ToOutput(in).(CreateModePtrOutput)
}

func (in *createModePtr) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CreateModePtrOutput)
}

type GeoRedundantBackupEnum string

const (
	GeoRedundantBackupEnumEnabled  = GeoRedundantBackupEnum("Enabled")
	GeoRedundantBackupEnumDisabled = GeoRedundantBackupEnum("Disabled")
)

func (GeoRedundantBackupEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoRedundantBackupEnum)(nil)).Elem()
}

func (e GeoRedundantBackupEnum) ToGeoRedundantBackupEnumOutput() GeoRedundantBackupEnumOutput {
	return pulumi.ToOutput(e).(GeoRedundantBackupEnumOutput)
}

func (e GeoRedundantBackupEnum) ToGeoRedundantBackupEnumOutputWithContext(ctx context.Context) GeoRedundantBackupEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GeoRedundantBackupEnumOutput)
}

func (e GeoRedundantBackupEnum) ToGeoRedundantBackupEnumPtrOutput() GeoRedundantBackupEnumPtrOutput {
	return e.ToGeoRedundantBackupEnumPtrOutputWithContext(context.Background())
}

func (e GeoRedundantBackupEnum) ToGeoRedundantBackupEnumPtrOutputWithContext(ctx context.Context) GeoRedundantBackupEnumPtrOutput {
	return GeoRedundantBackupEnum(e).ToGeoRedundantBackupEnumOutputWithContext(ctx).ToGeoRedundantBackupEnumPtrOutputWithContext(ctx)
}

func (e GeoRedundantBackupEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GeoRedundantBackupEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GeoRedundantBackupEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GeoRedundantBackupEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GeoRedundantBackupEnumOutput struct{ *pulumi.OutputState }

func (GeoRedundantBackupEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoRedundantBackupEnum)(nil)).Elem()
}

func (o GeoRedundantBackupEnumOutput) ToGeoRedundantBackupEnumOutput() GeoRedundantBackupEnumOutput {
	return o
}

func (o GeoRedundantBackupEnumOutput) ToGeoRedundantBackupEnumOutputWithContext(ctx context.Context) GeoRedundantBackupEnumOutput {
	return o
}

func (o GeoRedundantBackupEnumOutput) ToGeoRedundantBackupEnumPtrOutput() GeoRedundantBackupEnumPtrOutput {
	return o.ToGeoRedundantBackupEnumPtrOutputWithContext(context.Background())
}

func (o GeoRedundantBackupEnumOutput) ToGeoRedundantBackupEnumPtrOutputWithContext(ctx context.Context) GeoRedundantBackupEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GeoRedundantBackupEnum) *GeoRedundantBackupEnum {
		return &v
	}).(GeoRedundantBackupEnumPtrOutput)
}

func (o GeoRedundantBackupEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GeoRedundantBackupEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GeoRedundantBackupEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GeoRedundantBackupEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GeoRedundantBackupEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GeoRedundantBackupEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GeoRedundantBackupEnumPtrOutput struct{ *pulumi.OutputState }

func (GeoRedundantBackupEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GeoRedundantBackupEnum)(nil)).Elem()
}

func (o GeoRedundantBackupEnumPtrOutput) ToGeoRedundantBackupEnumPtrOutput() GeoRedundantBackupEnumPtrOutput {
	return o
}

func (o GeoRedundantBackupEnumPtrOutput) ToGeoRedundantBackupEnumPtrOutputWithContext(ctx context.Context) GeoRedundantBackupEnumPtrOutput {
	return o
}

func (o GeoRedundantBackupEnumPtrOutput) Elem() GeoRedundantBackupEnumOutput {
	return o.ApplyT(func(v *GeoRedundantBackupEnum) GeoRedundantBackupEnum {
		if v != nil {
			return *v
		}
		var ret GeoRedundantBackupEnum
		return ret
	}).(GeoRedundantBackupEnumOutput)
}

func (o GeoRedundantBackupEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GeoRedundantBackupEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GeoRedundantBackupEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type GeoRedundantBackupEnumInput interface {
	pulumi.Input

	ToGeoRedundantBackupEnumOutput() GeoRedundantBackupEnumOutput
	ToGeoRedundantBackupEnumOutputWithContext(context.Context) GeoRedundantBackupEnumOutput
}

var geoRedundantBackupEnumPtrType = reflect.TypeOf((**GeoRedundantBackupEnum)(nil)).Elem()

type GeoRedundantBackupEnumPtrInput interface {
	pulumi.Input

	ToGeoRedundantBackupEnumPtrOutput() GeoRedundantBackupEnumPtrOutput
	ToGeoRedundantBackupEnumPtrOutputWithContext(context.Context) GeoRedundantBackupEnumPtrOutput
}

type geoRedundantBackupEnumPtr string

func GeoRedundantBackupEnumPtr(v string) GeoRedundantBackupEnumPtrInput {
	return (*geoRedundantBackupEnumPtr)(&v)
}

func (*geoRedundantBackupEnumPtr) ElementType() reflect.Type {
	return geoRedundantBackupEnumPtrType
}

func (in *geoRedundantBackupEnumPtr) ToGeoRedundantBackupEnumPtrOutput() GeoRedundantBackupEnumPtrOutput {
	return pulumi.ToOutput(in).(GeoRedundantBackupEnumPtrOutput)
}

func (in *geoRedundantBackupEnumPtr) ToGeoRedundantBackupEnumPtrOutputWithContext(ctx context.Context) GeoRedundantBackupEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GeoRedundantBackupEnumPtrOutput)
}

type HAEnabledEnum string

const (
	HAEnabledEnumEnabled  = HAEnabledEnum("Enabled")
	HAEnabledEnumDisabled = HAEnabledEnum("Disabled")
)

func (HAEnabledEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*HAEnabledEnum)(nil)).Elem()
}

func (e HAEnabledEnum) ToHAEnabledEnumOutput() HAEnabledEnumOutput {
	return pulumi.ToOutput(e).(HAEnabledEnumOutput)
}

func (e HAEnabledEnum) ToHAEnabledEnumOutputWithContext(ctx context.Context) HAEnabledEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HAEnabledEnumOutput)
}

func (e HAEnabledEnum) ToHAEnabledEnumPtrOutput() HAEnabledEnumPtrOutput {
	return e.ToHAEnabledEnumPtrOutputWithContext(context.Background())
}

func (e HAEnabledEnum) ToHAEnabledEnumPtrOutputWithContext(ctx context.Context) HAEnabledEnumPtrOutput {
	return HAEnabledEnum(e).ToHAEnabledEnumOutputWithContext(ctx).ToHAEnabledEnumPtrOutputWithContext(ctx)
}

func (e HAEnabledEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HAEnabledEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HAEnabledEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HAEnabledEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HAEnabledEnumOutput struct{ *pulumi.OutputState }

func (HAEnabledEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HAEnabledEnum)(nil)).Elem()
}

func (o HAEnabledEnumOutput) ToHAEnabledEnumOutput() HAEnabledEnumOutput {
	return o
}

func (o HAEnabledEnumOutput) ToHAEnabledEnumOutputWithContext(ctx context.Context) HAEnabledEnumOutput {
	return o
}

func (o HAEnabledEnumOutput) ToHAEnabledEnumPtrOutput() HAEnabledEnumPtrOutput {
	return o.ToHAEnabledEnumPtrOutputWithContext(context.Background())
}

func (o HAEnabledEnumOutput) ToHAEnabledEnumPtrOutputWithContext(ctx context.Context) HAEnabledEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HAEnabledEnum) *HAEnabledEnum {
		return &v
	}).(HAEnabledEnumPtrOutput)
}

func (o HAEnabledEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HAEnabledEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HAEnabledEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HAEnabledEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HAEnabledEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HAEnabledEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HAEnabledEnumPtrOutput struct{ *pulumi.OutputState }

func (HAEnabledEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HAEnabledEnum)(nil)).Elem()
}

func (o HAEnabledEnumPtrOutput) ToHAEnabledEnumPtrOutput() HAEnabledEnumPtrOutput {
	return o
}

func (o HAEnabledEnumPtrOutput) ToHAEnabledEnumPtrOutputWithContext(ctx context.Context) HAEnabledEnumPtrOutput {
	return o
}

func (o HAEnabledEnumPtrOutput) Elem() HAEnabledEnumOutput {
	return o.ApplyT(func(v *HAEnabledEnum) HAEnabledEnum {
		if v != nil {
			return *v
		}
		var ret HAEnabledEnum
		return ret
	}).(HAEnabledEnumOutput)
}

func (o HAEnabledEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HAEnabledEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HAEnabledEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HAEnabledEnumInput interface {
	pulumi.Input

	ToHAEnabledEnumOutput() HAEnabledEnumOutput
	ToHAEnabledEnumOutputWithContext(context.Context) HAEnabledEnumOutput
}

var haenabledEnumPtrType = reflect.TypeOf((**HAEnabledEnum)(nil)).Elem()

type HAEnabledEnumPtrInput interface {
	pulumi.Input

	ToHAEnabledEnumPtrOutput() HAEnabledEnumPtrOutput
	ToHAEnabledEnumPtrOutputWithContext(context.Context) HAEnabledEnumPtrOutput
}

type haenabledEnumPtr string

func HAEnabledEnumPtr(v string) HAEnabledEnumPtrInput {
	return (*haenabledEnumPtr)(&v)
}

func (*haenabledEnumPtr) ElementType() reflect.Type {
	return haenabledEnumPtrType
}

func (in *haenabledEnumPtr) ToHAEnabledEnumPtrOutput() HAEnabledEnumPtrOutput {
	return pulumi.ToOutput(in).(HAEnabledEnumPtrOutput)
}

func (in *haenabledEnumPtr) ToHAEnabledEnumPtrOutputWithContext(ctx context.Context) HAEnabledEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HAEnabledEnumPtrOutput)
}

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

type ServerVersion string

const (
	ServerVersion_12 = ServerVersion("12")
	ServerVersion_11 = ServerVersion("11")
)

func (ServerVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerVersion)(nil)).Elem()
}

func (e ServerVersion) ToServerVersionOutput() ServerVersionOutput {
	return pulumi.ToOutput(e).(ServerVersionOutput)
}

func (e ServerVersion) ToServerVersionOutputWithContext(ctx context.Context) ServerVersionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServerVersionOutput)
}

func (e ServerVersion) ToServerVersionPtrOutput() ServerVersionPtrOutput {
	return e.ToServerVersionPtrOutputWithContext(context.Background())
}

func (e ServerVersion) ToServerVersionPtrOutputWithContext(ctx context.Context) ServerVersionPtrOutput {
	return ServerVersion(e).ToServerVersionOutputWithContext(ctx).ToServerVersionPtrOutputWithContext(ctx)
}

func (e ServerVersion) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServerVersion) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServerVersion) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServerVersion) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServerVersionOutput struct{ *pulumi.OutputState }

func (ServerVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerVersion)(nil)).Elem()
}

func (o ServerVersionOutput) ToServerVersionOutput() ServerVersionOutput {
	return o
}

func (o ServerVersionOutput) ToServerVersionOutputWithContext(ctx context.Context) ServerVersionOutput {
	return o
}

func (o ServerVersionOutput) ToServerVersionPtrOutput() ServerVersionPtrOutput {
	return o.ToServerVersionPtrOutputWithContext(context.Background())
}

func (o ServerVersionOutput) ToServerVersionPtrOutputWithContext(ctx context.Context) ServerVersionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerVersion) *ServerVersion {
		return &v
	}).(ServerVersionPtrOutput)
}

func (o ServerVersionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServerVersionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServerVersion) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServerVersionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServerVersionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServerVersion) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServerVersionPtrOutput struct{ *pulumi.OutputState }

func (ServerVersionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerVersion)(nil)).Elem()
}

func (o ServerVersionPtrOutput) ToServerVersionPtrOutput() ServerVersionPtrOutput {
	return o
}

func (o ServerVersionPtrOutput) ToServerVersionPtrOutputWithContext(ctx context.Context) ServerVersionPtrOutput {
	return o
}

func (o ServerVersionPtrOutput) Elem() ServerVersionOutput {
	return o.ApplyT(func(v *ServerVersion) ServerVersion {
		if v != nil {
			return *v
		}
		var ret ServerVersion
		return ret
	}).(ServerVersionOutput)
}

func (o ServerVersionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServerVersionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServerVersion) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServerVersionInput interface {
	pulumi.Input

	ToServerVersionOutput() ServerVersionOutput
	ToServerVersionOutputWithContext(context.Context) ServerVersionOutput
}

var serverVersionPtrType = reflect.TypeOf((**ServerVersion)(nil)).Elem()

type ServerVersionPtrInput interface {
	pulumi.Input

	ToServerVersionPtrOutput() ServerVersionPtrOutput
	ToServerVersionPtrOutputWithContext(context.Context) ServerVersionPtrOutput
}

type serverVersionPtr string

func ServerVersionPtr(v string) ServerVersionPtrInput {
	return (*serverVersionPtr)(&v)
}

func (*serverVersionPtr) ElementType() reflect.Type {
	return serverVersionPtrType
}

func (in *serverVersionPtr) ToServerVersionPtrOutput() ServerVersionPtrOutput {
	return pulumi.ToOutput(in).(ServerVersionPtrOutput)
}

func (in *serverVersionPtr) ToServerVersionPtrOutputWithContext(ctx context.Context) ServerVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServerVersionPtrOutput)
}

type SkuTier string

const (
	SkuTierBurstable       = SkuTier("Burstable")
	SkuTierGeneralPurpose  = SkuTier("GeneralPurpose")
	SkuTierMemoryOptimized = SkuTier("MemoryOptimized")
)

func (SkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuTier)(nil)).Elem()
}

func (e SkuTier) ToSkuTierOutput() SkuTierOutput {
	return pulumi.ToOutput(e).(SkuTierOutput)
}

func (e SkuTier) ToSkuTierOutputWithContext(ctx context.Context) SkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuTierOutput)
}

func (e SkuTier) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return e.ToSkuTierPtrOutputWithContext(context.Background())
}

func (e SkuTier) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return SkuTier(e).ToSkuTierOutputWithContext(ctx).ToSkuTierPtrOutputWithContext(ctx)
}

func (e SkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuTierOutput struct{ *pulumi.OutputState }

func (SkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuTier)(nil)).Elem()
}

func (o SkuTierOutput) ToSkuTierOutput() SkuTierOutput {
	return o
}

func (o SkuTierOutput) ToSkuTierOutputWithContext(ctx context.Context) SkuTierOutput {
	return o
}

func (o SkuTierOutput) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return o.ToSkuTierPtrOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuTier) *SkuTier {
		return &v
	}).(SkuTierPtrOutput)
}

func (o SkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuTierPtrOutput struct{ *pulumi.OutputState }

func (SkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuTier)(nil)).Elem()
}

func (o SkuTierPtrOutput) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return o
}

func (o SkuTierPtrOutput) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return o
}

func (o SkuTierPtrOutput) Elem() SkuTierOutput {
	return o.ApplyT(func(v *SkuTier) SkuTier {
		if v != nil {
			return *v
		}
		var ret SkuTier
		return ret
	}).(SkuTierOutput)
}

func (o SkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuTierInput interface {
	pulumi.Input

	ToSkuTierOutput() SkuTierOutput
	ToSkuTierOutputWithContext(context.Context) SkuTierOutput
}

var skuTierPtrType = reflect.TypeOf((**SkuTier)(nil)).Elem()

type SkuTierPtrInput interface {
	pulumi.Input

	ToSkuTierPtrOutput() SkuTierPtrOutput
	ToSkuTierPtrOutputWithContext(context.Context) SkuTierPtrOutput
}

type skuTierPtr string

func SkuTierPtr(v string) SkuTierPtrInput {
	return (*skuTierPtr)(&v)
}

func (*skuTierPtr) ElementType() reflect.Type {
	return skuTierPtrType
}

func (in *skuTierPtr) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return pulumi.ToOutput(in).(SkuTierPtrOutput)
}

func (in *skuTierPtr) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuTierPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CreateModeOutput{})
	pulumi.RegisterOutputType(CreateModePtrOutput{})
	pulumi.RegisterOutputType(GeoRedundantBackupEnumOutput{})
	pulumi.RegisterOutputType(GeoRedundantBackupEnumPtrOutput{})
	pulumi.RegisterOutputType(HAEnabledEnumOutput{})
	pulumi.RegisterOutputType(HAEnabledEnumPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(ServerVersionOutput{})
	pulumi.RegisterOutputType(ServerVersionPtrOutput{})
	pulumi.RegisterOutputType(SkuTierOutput{})
	pulumi.RegisterOutputType(SkuTierPtrOutput{})
}
