


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CreateMode string

const (
	CreateModeDefault            = CreateMode("Default")
	CreateModeCreate             = CreateMode("Create")
	CreateModeUpdate             = CreateMode("Update")
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

type HighAvailabilityMode string

const (
	HighAvailabilityModeDisabled      = HighAvailabilityMode("Disabled")
	HighAvailabilityModeZoneRedundant = HighAvailabilityMode("ZoneRedundant")
)

func (HighAvailabilityMode) ElementType() reflect.Type {
	return reflect.TypeOf((*HighAvailabilityMode)(nil)).Elem()
}

func (e HighAvailabilityMode) ToHighAvailabilityModeOutput() HighAvailabilityModeOutput {
	return pulumi.ToOutput(e).(HighAvailabilityModeOutput)
}

func (e HighAvailabilityMode) ToHighAvailabilityModeOutputWithContext(ctx context.Context) HighAvailabilityModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HighAvailabilityModeOutput)
}

func (e HighAvailabilityMode) ToHighAvailabilityModePtrOutput() HighAvailabilityModePtrOutput {
	return e.ToHighAvailabilityModePtrOutputWithContext(context.Background())
}

func (e HighAvailabilityMode) ToHighAvailabilityModePtrOutputWithContext(ctx context.Context) HighAvailabilityModePtrOutput {
	return HighAvailabilityMode(e).ToHighAvailabilityModeOutputWithContext(ctx).ToHighAvailabilityModePtrOutputWithContext(ctx)
}

func (e HighAvailabilityMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HighAvailabilityMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HighAvailabilityMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HighAvailabilityMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HighAvailabilityModeOutput struct{ *pulumi.OutputState }

func (HighAvailabilityModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HighAvailabilityMode)(nil)).Elem()
}

func (o HighAvailabilityModeOutput) ToHighAvailabilityModeOutput() HighAvailabilityModeOutput {
	return o
}

func (o HighAvailabilityModeOutput) ToHighAvailabilityModeOutputWithContext(ctx context.Context) HighAvailabilityModeOutput {
	return o
}

func (o HighAvailabilityModeOutput) ToHighAvailabilityModePtrOutput() HighAvailabilityModePtrOutput {
	return o.ToHighAvailabilityModePtrOutputWithContext(context.Background())
}

func (o HighAvailabilityModeOutput) ToHighAvailabilityModePtrOutputWithContext(ctx context.Context) HighAvailabilityModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HighAvailabilityMode) *HighAvailabilityMode {
		return &v
	}).(HighAvailabilityModePtrOutput)
}

func (o HighAvailabilityModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HighAvailabilityModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HighAvailabilityMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HighAvailabilityModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HighAvailabilityModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HighAvailabilityMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HighAvailabilityModePtrOutput struct{ *pulumi.OutputState }

func (HighAvailabilityModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HighAvailabilityMode)(nil)).Elem()
}

func (o HighAvailabilityModePtrOutput) ToHighAvailabilityModePtrOutput() HighAvailabilityModePtrOutput {
	return o
}

func (o HighAvailabilityModePtrOutput) ToHighAvailabilityModePtrOutputWithContext(ctx context.Context) HighAvailabilityModePtrOutput {
	return o
}

func (o HighAvailabilityModePtrOutput) Elem() HighAvailabilityModeOutput {
	return o.ApplyT(func(v *HighAvailabilityMode) HighAvailabilityMode {
		if v != nil {
			return *v
		}
		var ret HighAvailabilityMode
		return ret
	}).(HighAvailabilityModeOutput)
}

func (o HighAvailabilityModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HighAvailabilityModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HighAvailabilityMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HighAvailabilityModeInput interface {
	pulumi.Input

	ToHighAvailabilityModeOutput() HighAvailabilityModeOutput
	ToHighAvailabilityModeOutputWithContext(context.Context) HighAvailabilityModeOutput
}

var highAvailabilityModePtrType = reflect.TypeOf((**HighAvailabilityMode)(nil)).Elem()

type HighAvailabilityModePtrInput interface {
	pulumi.Input

	ToHighAvailabilityModePtrOutput() HighAvailabilityModePtrOutput
	ToHighAvailabilityModePtrOutputWithContext(context.Context) HighAvailabilityModePtrOutput
}

type highAvailabilityModePtr string

func HighAvailabilityModePtr(v string) HighAvailabilityModePtrInput {
	return (*highAvailabilityModePtr)(&v)
}

func (*highAvailabilityModePtr) ElementType() reflect.Type {
	return highAvailabilityModePtrType
}

func (in *highAvailabilityModePtr) ToHighAvailabilityModePtrOutput() HighAvailabilityModePtrOutput {
	return pulumi.ToOutput(in).(HighAvailabilityModePtrOutput)
}

func (in *highAvailabilityModePtr) ToHighAvailabilityModePtrOutputWithContext(ctx context.Context) HighAvailabilityModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HighAvailabilityModePtrOutput)
}

type ServerVersion string

const (
	ServerVersion_13 = ServerVersion("13")
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
	pulumi.RegisterOutputType(HighAvailabilityModeOutput{})
	pulumi.RegisterOutputType(HighAvailabilityModePtrOutput{})
	pulumi.RegisterOutputType(ServerVersionOutput{})
	pulumi.RegisterOutputType(ServerVersionPtrOutput{})
	pulumi.RegisterOutputType(SkuTierOutput{})
	pulumi.RegisterOutputType(SkuTierPtrOutput{})
}
