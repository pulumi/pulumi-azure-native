


package dbforpostgresql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdministratorType string

const (
	AdministratorTypeActiveDirectory = AdministratorType("ActiveDirectory")
)

func (AdministratorType) ElementType() reflect.Type {
	return reflect.TypeOf((*AdministratorType)(nil)).Elem()
}

func (e AdministratorType) ToAdministratorTypeOutput() AdministratorTypeOutput {
	return pulumi.ToOutput(e).(AdministratorTypeOutput)
}

func (e AdministratorType) ToAdministratorTypeOutputWithContext(ctx context.Context) AdministratorTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AdministratorTypeOutput)
}

func (e AdministratorType) ToAdministratorTypePtrOutput() AdministratorTypePtrOutput {
	return e.ToAdministratorTypePtrOutputWithContext(context.Background())
}

func (e AdministratorType) ToAdministratorTypePtrOutputWithContext(ctx context.Context) AdministratorTypePtrOutput {
	return AdministratorType(e).ToAdministratorTypeOutputWithContext(ctx).ToAdministratorTypePtrOutputWithContext(ctx)
}

func (e AdministratorType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AdministratorType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AdministratorType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AdministratorType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AdministratorTypeOutput struct{ *pulumi.OutputState }

func (AdministratorTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdministratorType)(nil)).Elem()
}

func (o AdministratorTypeOutput) ToAdministratorTypeOutput() AdministratorTypeOutput {
	return o
}

func (o AdministratorTypeOutput) ToAdministratorTypeOutputWithContext(ctx context.Context) AdministratorTypeOutput {
	return o
}

func (o AdministratorTypeOutput) ToAdministratorTypePtrOutput() AdministratorTypePtrOutput {
	return o.ToAdministratorTypePtrOutputWithContext(context.Background())
}

func (o AdministratorTypeOutput) ToAdministratorTypePtrOutputWithContext(ctx context.Context) AdministratorTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AdministratorType) *AdministratorType {
		return &v
	}).(AdministratorTypePtrOutput)
}

func (o AdministratorTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AdministratorTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AdministratorType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AdministratorTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AdministratorTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AdministratorType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AdministratorTypePtrOutput struct{ *pulumi.OutputState }

func (AdministratorTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdministratorType)(nil)).Elem()
}

func (o AdministratorTypePtrOutput) ToAdministratorTypePtrOutput() AdministratorTypePtrOutput {
	return o
}

func (o AdministratorTypePtrOutput) ToAdministratorTypePtrOutputWithContext(ctx context.Context) AdministratorTypePtrOutput {
	return o
}

func (o AdministratorTypePtrOutput) Elem() AdministratorTypeOutput {
	return o.ApplyT(func(v *AdministratorType) AdministratorType {
		if v != nil {
			return *v
		}
		var ret AdministratorType
		return ret
	}).(AdministratorTypeOutput)
}

func (o AdministratorTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AdministratorTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AdministratorType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AdministratorTypeInput interface {
	pulumi.Input

	ToAdministratorTypeOutput() AdministratorTypeOutput
	ToAdministratorTypeOutputWithContext(context.Context) AdministratorTypeOutput
}

var administratorTypePtrType = reflect.TypeOf((**AdministratorType)(nil)).Elem()

type AdministratorTypePtrInput interface {
	pulumi.Input

	ToAdministratorTypePtrOutput() AdministratorTypePtrOutput
	ToAdministratorTypePtrOutputWithContext(context.Context) AdministratorTypePtrOutput
}

type administratorTypePtr string

func AdministratorTypePtr(v string) AdministratorTypePtrInput {
	return (*administratorTypePtr)(&v)
}

func (*administratorTypePtr) ElementType() reflect.Type {
	return administratorTypePtrType
}

func (in *administratorTypePtr) ToAdministratorTypePtrOutput() AdministratorTypePtrOutput {
	return pulumi.ToOutput(in).(AdministratorTypePtrOutput)
}

func (in *administratorTypePtr) ToAdministratorTypePtrOutputWithContext(ctx context.Context) AdministratorTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AdministratorTypePtrOutput)
}

type CreateMode string

const (
	CreateModeDefault            = CreateMode("Default")
	CreateModePointInTimeRestore = CreateMode("PointInTimeRestore")
	CreateModeGeoRestore         = CreateMode("GeoRestore")
	CreateModeReplica            = CreateMode("Replica")
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

type GeoRedundantBackup string

const (
	GeoRedundantBackupEnabled  = GeoRedundantBackup("Enabled")
	GeoRedundantBackupDisabled = GeoRedundantBackup("Disabled")
)

func (GeoRedundantBackup) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoRedundantBackup)(nil)).Elem()
}

func (e GeoRedundantBackup) ToGeoRedundantBackupOutput() GeoRedundantBackupOutput {
	return pulumi.ToOutput(e).(GeoRedundantBackupOutput)
}

func (e GeoRedundantBackup) ToGeoRedundantBackupOutputWithContext(ctx context.Context) GeoRedundantBackupOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GeoRedundantBackupOutput)
}

func (e GeoRedundantBackup) ToGeoRedundantBackupPtrOutput() GeoRedundantBackupPtrOutput {
	return e.ToGeoRedundantBackupPtrOutputWithContext(context.Background())
}

func (e GeoRedundantBackup) ToGeoRedundantBackupPtrOutputWithContext(ctx context.Context) GeoRedundantBackupPtrOutput {
	return GeoRedundantBackup(e).ToGeoRedundantBackupOutputWithContext(ctx).ToGeoRedundantBackupPtrOutputWithContext(ctx)
}

func (e GeoRedundantBackup) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GeoRedundantBackup) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GeoRedundantBackup) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GeoRedundantBackup) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GeoRedundantBackupOutput struct{ *pulumi.OutputState }

func (GeoRedundantBackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoRedundantBackup)(nil)).Elem()
}

func (o GeoRedundantBackupOutput) ToGeoRedundantBackupOutput() GeoRedundantBackupOutput {
	return o
}

func (o GeoRedundantBackupOutput) ToGeoRedundantBackupOutputWithContext(ctx context.Context) GeoRedundantBackupOutput {
	return o
}

func (o GeoRedundantBackupOutput) ToGeoRedundantBackupPtrOutput() GeoRedundantBackupPtrOutput {
	return o.ToGeoRedundantBackupPtrOutputWithContext(context.Background())
}

func (o GeoRedundantBackupOutput) ToGeoRedundantBackupPtrOutputWithContext(ctx context.Context) GeoRedundantBackupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GeoRedundantBackup) *GeoRedundantBackup {
		return &v
	}).(GeoRedundantBackupPtrOutput)
}

func (o GeoRedundantBackupOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GeoRedundantBackupOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GeoRedundantBackup) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GeoRedundantBackupOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GeoRedundantBackupOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GeoRedundantBackup) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GeoRedundantBackupPtrOutput struct{ *pulumi.OutputState }

func (GeoRedundantBackupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GeoRedundantBackup)(nil)).Elem()
}

func (o GeoRedundantBackupPtrOutput) ToGeoRedundantBackupPtrOutput() GeoRedundantBackupPtrOutput {
	return o
}

func (o GeoRedundantBackupPtrOutput) ToGeoRedundantBackupPtrOutputWithContext(ctx context.Context) GeoRedundantBackupPtrOutput {
	return o
}

func (o GeoRedundantBackupPtrOutput) Elem() GeoRedundantBackupOutput {
	return o.ApplyT(func(v *GeoRedundantBackup) GeoRedundantBackup {
		if v != nil {
			return *v
		}
		var ret GeoRedundantBackup
		return ret
	}).(GeoRedundantBackupOutput)
}

func (o GeoRedundantBackupPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GeoRedundantBackupPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GeoRedundantBackup) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type GeoRedundantBackupInput interface {
	pulumi.Input

	ToGeoRedundantBackupOutput() GeoRedundantBackupOutput
	ToGeoRedundantBackupOutputWithContext(context.Context) GeoRedundantBackupOutput
}

var geoRedundantBackupPtrType = reflect.TypeOf((**GeoRedundantBackup)(nil)).Elem()

type GeoRedundantBackupPtrInput interface {
	pulumi.Input

	ToGeoRedundantBackupPtrOutput() GeoRedundantBackupPtrOutput
	ToGeoRedundantBackupPtrOutputWithContext(context.Context) GeoRedundantBackupPtrOutput
}

type geoRedundantBackupPtr string

func GeoRedundantBackupPtr(v string) GeoRedundantBackupPtrInput {
	return (*geoRedundantBackupPtr)(&v)
}

func (*geoRedundantBackupPtr) ElementType() reflect.Type {
	return geoRedundantBackupPtrType
}

func (in *geoRedundantBackupPtr) ToGeoRedundantBackupPtrOutput() GeoRedundantBackupPtrOutput {
	return pulumi.ToOutput(in).(GeoRedundantBackupPtrOutput)
}

func (in *geoRedundantBackupPtr) ToGeoRedundantBackupPtrOutputWithContext(ctx context.Context) GeoRedundantBackupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GeoRedundantBackupPtrOutput)
}

type IdentityType string

const (
	IdentityTypeSystemAssigned = IdentityType("SystemAssigned")
)

func (IdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (e IdentityType) ToIdentityTypeOutput() IdentityTypeOutput {
	return pulumi.ToOutput(e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return e.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (e IdentityType) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return IdentityType(e).ToIdentityTypeOutputWithContext(ctx).ToIdentityTypePtrOutputWithContext(ctx)
}

func (e IdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IdentityTypeOutput struct{ *pulumi.OutputState }

func (IdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (o IdentityTypeOutput) ToIdentityTypeOutput() IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityType) *IdentityType {
		return &v
	}).(IdentityTypePtrOutput)
}

func (o IdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IdentityTypePtrOutput struct{ *pulumi.OutputState }

func (IdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityType)(nil)).Elem()
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) Elem() IdentityTypeOutput {
	return o.ApplyT(func(v *IdentityType) IdentityType {
		if v != nil {
			return *v
		}
		var ret IdentityType
		return ret
	}).(IdentityTypeOutput)
}

func (o IdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IdentityTypeInput interface {
	pulumi.Input

	ToIdentityTypeOutput() IdentityTypeOutput
	ToIdentityTypeOutputWithContext(context.Context) IdentityTypeOutput
}

var identityTypePtrType = reflect.TypeOf((**IdentityType)(nil)).Elem()

type IdentityTypePtrInput interface {
	pulumi.Input

	ToIdentityTypePtrOutput() IdentityTypePtrOutput
	ToIdentityTypePtrOutputWithContext(context.Context) IdentityTypePtrOutput
}

type identityTypePtr string

func IdentityTypePtr(v string) IdentityTypePtrInput {
	return (*identityTypePtr)(&v)
}

func (*identityTypePtr) ElementType() reflect.Type {
	return identityTypePtrType
}

func (in *identityTypePtr) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return pulumi.ToOutput(in).(IdentityTypePtrOutput)
}

func (in *identityTypePtr) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IdentityTypePtrOutput)
}

type InfrastructureEncryption string

const (
	// Default value for single layer of encryption for data at rest.
	InfrastructureEncryptionEnabled = InfrastructureEncryption("Enabled")
	// Additional (2nd) layer of encryption for data at rest
	InfrastructureEncryptionDisabled = InfrastructureEncryption("Disabled")
)

func (InfrastructureEncryption) ElementType() reflect.Type {
	return reflect.TypeOf((*InfrastructureEncryption)(nil)).Elem()
}

func (e InfrastructureEncryption) ToInfrastructureEncryptionOutput() InfrastructureEncryptionOutput {
	return pulumi.ToOutput(e).(InfrastructureEncryptionOutput)
}

func (e InfrastructureEncryption) ToInfrastructureEncryptionOutputWithContext(ctx context.Context) InfrastructureEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InfrastructureEncryptionOutput)
}

func (e InfrastructureEncryption) ToInfrastructureEncryptionPtrOutput() InfrastructureEncryptionPtrOutput {
	return e.ToInfrastructureEncryptionPtrOutputWithContext(context.Background())
}

func (e InfrastructureEncryption) ToInfrastructureEncryptionPtrOutputWithContext(ctx context.Context) InfrastructureEncryptionPtrOutput {
	return InfrastructureEncryption(e).ToInfrastructureEncryptionOutputWithContext(ctx).ToInfrastructureEncryptionPtrOutputWithContext(ctx)
}

func (e InfrastructureEncryption) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InfrastructureEncryption) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InfrastructureEncryption) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InfrastructureEncryption) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InfrastructureEncryptionOutput struct{ *pulumi.OutputState }

func (InfrastructureEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InfrastructureEncryption)(nil)).Elem()
}

func (o InfrastructureEncryptionOutput) ToInfrastructureEncryptionOutput() InfrastructureEncryptionOutput {
	return o
}

func (o InfrastructureEncryptionOutput) ToInfrastructureEncryptionOutputWithContext(ctx context.Context) InfrastructureEncryptionOutput {
	return o
}

func (o InfrastructureEncryptionOutput) ToInfrastructureEncryptionPtrOutput() InfrastructureEncryptionPtrOutput {
	return o.ToInfrastructureEncryptionPtrOutputWithContext(context.Background())
}

func (o InfrastructureEncryptionOutput) ToInfrastructureEncryptionPtrOutputWithContext(ctx context.Context) InfrastructureEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InfrastructureEncryption) *InfrastructureEncryption {
		return &v
	}).(InfrastructureEncryptionPtrOutput)
}

func (o InfrastructureEncryptionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InfrastructureEncryptionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InfrastructureEncryption) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InfrastructureEncryptionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InfrastructureEncryptionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InfrastructureEncryption) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InfrastructureEncryptionPtrOutput struct{ *pulumi.OutputState }

func (InfrastructureEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InfrastructureEncryption)(nil)).Elem()
}

func (o InfrastructureEncryptionPtrOutput) ToInfrastructureEncryptionPtrOutput() InfrastructureEncryptionPtrOutput {
	return o
}

func (o InfrastructureEncryptionPtrOutput) ToInfrastructureEncryptionPtrOutputWithContext(ctx context.Context) InfrastructureEncryptionPtrOutput {
	return o
}

func (o InfrastructureEncryptionPtrOutput) Elem() InfrastructureEncryptionOutput {
	return o.ApplyT(func(v *InfrastructureEncryption) InfrastructureEncryption {
		if v != nil {
			return *v
		}
		var ret InfrastructureEncryption
		return ret
	}).(InfrastructureEncryptionOutput)
}

func (o InfrastructureEncryptionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InfrastructureEncryptionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InfrastructureEncryption) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type InfrastructureEncryptionInput interface {
	pulumi.Input

	ToInfrastructureEncryptionOutput() InfrastructureEncryptionOutput
	ToInfrastructureEncryptionOutputWithContext(context.Context) InfrastructureEncryptionOutput
}

var infrastructureEncryptionPtrType = reflect.TypeOf((**InfrastructureEncryption)(nil)).Elem()

type InfrastructureEncryptionPtrInput interface {
	pulumi.Input

	ToInfrastructureEncryptionPtrOutput() InfrastructureEncryptionPtrOutput
	ToInfrastructureEncryptionPtrOutputWithContext(context.Context) InfrastructureEncryptionPtrOutput
}

type infrastructureEncryptionPtr string

func InfrastructureEncryptionPtr(v string) InfrastructureEncryptionPtrInput {
	return (*infrastructureEncryptionPtr)(&v)
}

func (*infrastructureEncryptionPtr) ElementType() reflect.Type {
	return infrastructureEncryptionPtrType
}

func (in *infrastructureEncryptionPtr) ToInfrastructureEncryptionPtrOutput() InfrastructureEncryptionPtrOutput {
	return pulumi.ToOutput(in).(InfrastructureEncryptionPtrOutput)
}

func (in *infrastructureEncryptionPtr) ToInfrastructureEncryptionPtrOutputWithContext(ctx context.Context) InfrastructureEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InfrastructureEncryptionPtrOutput)
}

type MinimalTlsVersionEnum string

const (
	MinimalTlsVersionEnum_TLS1_0                = MinimalTlsVersionEnum("TLS1_0")
	MinimalTlsVersionEnum_TLS1_1                = MinimalTlsVersionEnum("TLS1_1")
	MinimalTlsVersionEnum_TLS1_2                = MinimalTlsVersionEnum("TLS1_2")
	MinimalTlsVersionEnumTLSEnforcementDisabled = MinimalTlsVersionEnum("TLSEnforcementDisabled")
)

func (MinimalTlsVersionEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*MinimalTlsVersionEnum)(nil)).Elem()
}

func (e MinimalTlsVersionEnum) ToMinimalTlsVersionEnumOutput() MinimalTlsVersionEnumOutput {
	return pulumi.ToOutput(e).(MinimalTlsVersionEnumOutput)
}

func (e MinimalTlsVersionEnum) ToMinimalTlsVersionEnumOutputWithContext(ctx context.Context) MinimalTlsVersionEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MinimalTlsVersionEnumOutput)
}

func (e MinimalTlsVersionEnum) ToMinimalTlsVersionEnumPtrOutput() MinimalTlsVersionEnumPtrOutput {
	return e.ToMinimalTlsVersionEnumPtrOutputWithContext(context.Background())
}

func (e MinimalTlsVersionEnum) ToMinimalTlsVersionEnumPtrOutputWithContext(ctx context.Context) MinimalTlsVersionEnumPtrOutput {
	return MinimalTlsVersionEnum(e).ToMinimalTlsVersionEnumOutputWithContext(ctx).ToMinimalTlsVersionEnumPtrOutputWithContext(ctx)
}

func (e MinimalTlsVersionEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MinimalTlsVersionEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MinimalTlsVersionEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MinimalTlsVersionEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MinimalTlsVersionEnumOutput struct{ *pulumi.OutputState }

func (MinimalTlsVersionEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MinimalTlsVersionEnum)(nil)).Elem()
}

func (o MinimalTlsVersionEnumOutput) ToMinimalTlsVersionEnumOutput() MinimalTlsVersionEnumOutput {
	return o
}

func (o MinimalTlsVersionEnumOutput) ToMinimalTlsVersionEnumOutputWithContext(ctx context.Context) MinimalTlsVersionEnumOutput {
	return o
}

func (o MinimalTlsVersionEnumOutput) ToMinimalTlsVersionEnumPtrOutput() MinimalTlsVersionEnumPtrOutput {
	return o.ToMinimalTlsVersionEnumPtrOutputWithContext(context.Background())
}

func (o MinimalTlsVersionEnumOutput) ToMinimalTlsVersionEnumPtrOutputWithContext(ctx context.Context) MinimalTlsVersionEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MinimalTlsVersionEnum) *MinimalTlsVersionEnum {
		return &v
	}).(MinimalTlsVersionEnumPtrOutput)
}

func (o MinimalTlsVersionEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MinimalTlsVersionEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MinimalTlsVersionEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MinimalTlsVersionEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MinimalTlsVersionEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MinimalTlsVersionEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MinimalTlsVersionEnumPtrOutput struct{ *pulumi.OutputState }

func (MinimalTlsVersionEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MinimalTlsVersionEnum)(nil)).Elem()
}

func (o MinimalTlsVersionEnumPtrOutput) ToMinimalTlsVersionEnumPtrOutput() MinimalTlsVersionEnumPtrOutput {
	return o
}

func (o MinimalTlsVersionEnumPtrOutput) ToMinimalTlsVersionEnumPtrOutputWithContext(ctx context.Context) MinimalTlsVersionEnumPtrOutput {
	return o
}

func (o MinimalTlsVersionEnumPtrOutput) Elem() MinimalTlsVersionEnumOutput {
	return o.ApplyT(func(v *MinimalTlsVersionEnum) MinimalTlsVersionEnum {
		if v != nil {
			return *v
		}
		var ret MinimalTlsVersionEnum
		return ret
	}).(MinimalTlsVersionEnumOutput)
}

func (o MinimalTlsVersionEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MinimalTlsVersionEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MinimalTlsVersionEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MinimalTlsVersionEnumInput interface {
	pulumi.Input

	ToMinimalTlsVersionEnumOutput() MinimalTlsVersionEnumOutput
	ToMinimalTlsVersionEnumOutputWithContext(context.Context) MinimalTlsVersionEnumOutput
}

var minimalTlsVersionEnumPtrType = reflect.TypeOf((**MinimalTlsVersionEnum)(nil)).Elem()

type MinimalTlsVersionEnumPtrInput interface {
	pulumi.Input

	ToMinimalTlsVersionEnumPtrOutput() MinimalTlsVersionEnumPtrOutput
	ToMinimalTlsVersionEnumPtrOutputWithContext(context.Context) MinimalTlsVersionEnumPtrOutput
}

type minimalTlsVersionEnumPtr string

func MinimalTlsVersionEnumPtr(v string) MinimalTlsVersionEnumPtrInput {
	return (*minimalTlsVersionEnumPtr)(&v)
}

func (*minimalTlsVersionEnumPtr) ElementType() reflect.Type {
	return minimalTlsVersionEnumPtrType
}

func (in *minimalTlsVersionEnumPtr) ToMinimalTlsVersionEnumPtrOutput() MinimalTlsVersionEnumPtrOutput {
	return pulumi.ToOutput(in).(MinimalTlsVersionEnumPtrOutput)
}

func (in *minimalTlsVersionEnumPtr) ToMinimalTlsVersionEnumPtrOutputWithContext(ctx context.Context) MinimalTlsVersionEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MinimalTlsVersionEnumPtrOutput)
}

type PublicNetworkAccessEnum string

const (
	PublicNetworkAccessEnumEnabled  = PublicNetworkAccessEnum("Enabled")
	PublicNetworkAccessEnumDisabled = PublicNetworkAccessEnum("Disabled")
)

func (PublicNetworkAccessEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccessEnum)(nil)).Elem()
}

func (e PublicNetworkAccessEnum) ToPublicNetworkAccessEnumOutput() PublicNetworkAccessEnumOutput {
	return pulumi.ToOutput(e).(PublicNetworkAccessEnumOutput)
}

func (e PublicNetworkAccessEnum) ToPublicNetworkAccessEnumOutputWithContext(ctx context.Context) PublicNetworkAccessEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicNetworkAccessEnumOutput)
}

func (e PublicNetworkAccessEnum) ToPublicNetworkAccessEnumPtrOutput() PublicNetworkAccessEnumPtrOutput {
	return e.ToPublicNetworkAccessEnumPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccessEnum) ToPublicNetworkAccessEnumPtrOutputWithContext(ctx context.Context) PublicNetworkAccessEnumPtrOutput {
	return PublicNetworkAccessEnum(e).ToPublicNetworkAccessEnumOutputWithContext(ctx).ToPublicNetworkAccessEnumPtrOutputWithContext(ctx)
}

func (e PublicNetworkAccessEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccessEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccessEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccessEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicNetworkAccessEnumOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccessEnum)(nil)).Elem()
}

func (o PublicNetworkAccessEnumOutput) ToPublicNetworkAccessEnumOutput() PublicNetworkAccessEnumOutput {
	return o
}

func (o PublicNetworkAccessEnumOutput) ToPublicNetworkAccessEnumOutputWithContext(ctx context.Context) PublicNetworkAccessEnumOutput {
	return o
}

func (o PublicNetworkAccessEnumOutput) ToPublicNetworkAccessEnumPtrOutput() PublicNetworkAccessEnumPtrOutput {
	return o.ToPublicNetworkAccessEnumPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessEnumOutput) ToPublicNetworkAccessEnumPtrOutputWithContext(ctx context.Context) PublicNetworkAccessEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicNetworkAccessEnum) *PublicNetworkAccessEnum {
		return &v
	}).(PublicNetworkAccessEnumPtrOutput)
}

func (o PublicNetworkAccessEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicNetworkAccessEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccessEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicNetworkAccessEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccessEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicNetworkAccessEnumPtrOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicNetworkAccessEnum)(nil)).Elem()
}

func (o PublicNetworkAccessEnumPtrOutput) ToPublicNetworkAccessEnumPtrOutput() PublicNetworkAccessEnumPtrOutput {
	return o
}

func (o PublicNetworkAccessEnumPtrOutput) ToPublicNetworkAccessEnumPtrOutputWithContext(ctx context.Context) PublicNetworkAccessEnumPtrOutput {
	return o
}

func (o PublicNetworkAccessEnumPtrOutput) Elem() PublicNetworkAccessEnumOutput {
	return o.ApplyT(func(v *PublicNetworkAccessEnum) PublicNetworkAccessEnum {
		if v != nil {
			return *v
		}
		var ret PublicNetworkAccessEnum
		return ret
	}).(PublicNetworkAccessEnumOutput)
}

func (o PublicNetworkAccessEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicNetworkAccessEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PublicNetworkAccessEnumInput interface {
	pulumi.Input

	ToPublicNetworkAccessEnumOutput() PublicNetworkAccessEnumOutput
	ToPublicNetworkAccessEnumOutputWithContext(context.Context) PublicNetworkAccessEnumOutput
}

var publicNetworkAccessEnumPtrType = reflect.TypeOf((**PublicNetworkAccessEnum)(nil)).Elem()

type PublicNetworkAccessEnumPtrInput interface {
	pulumi.Input

	ToPublicNetworkAccessEnumPtrOutput() PublicNetworkAccessEnumPtrOutput
	ToPublicNetworkAccessEnumPtrOutputWithContext(context.Context) PublicNetworkAccessEnumPtrOutput
}

type publicNetworkAccessEnumPtr string

func PublicNetworkAccessEnumPtr(v string) PublicNetworkAccessEnumPtrInput {
	return (*publicNetworkAccessEnumPtr)(&v)
}

func (*publicNetworkAccessEnumPtr) ElementType() reflect.Type {
	return publicNetworkAccessEnumPtrType
}

func (in *publicNetworkAccessEnumPtr) ToPublicNetworkAccessEnumPtrOutput() PublicNetworkAccessEnumPtrOutput {
	return pulumi.ToOutput(in).(PublicNetworkAccessEnumPtrOutput)
}

func (in *publicNetworkAccessEnumPtr) ToPublicNetworkAccessEnumPtrOutputWithContext(ctx context.Context) PublicNetworkAccessEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicNetworkAccessEnumPtrOutput)
}

type ServerKeyType string

const (
	ServerKeyTypeAzureKeyVault = ServerKeyType("AzureKeyVault")
)

func (ServerKeyType) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerKeyType)(nil)).Elem()
}

func (e ServerKeyType) ToServerKeyTypeOutput() ServerKeyTypeOutput {
	return pulumi.ToOutput(e).(ServerKeyTypeOutput)
}

func (e ServerKeyType) ToServerKeyTypeOutputWithContext(ctx context.Context) ServerKeyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServerKeyTypeOutput)
}

func (e ServerKeyType) ToServerKeyTypePtrOutput() ServerKeyTypePtrOutput {
	return e.ToServerKeyTypePtrOutputWithContext(context.Background())
}

func (e ServerKeyType) ToServerKeyTypePtrOutputWithContext(ctx context.Context) ServerKeyTypePtrOutput {
	return ServerKeyType(e).ToServerKeyTypeOutputWithContext(ctx).ToServerKeyTypePtrOutputWithContext(ctx)
}

func (e ServerKeyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServerKeyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServerKeyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServerKeyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServerKeyTypeOutput struct{ *pulumi.OutputState }

func (ServerKeyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerKeyType)(nil)).Elem()
}

func (o ServerKeyTypeOutput) ToServerKeyTypeOutput() ServerKeyTypeOutput {
	return o
}

func (o ServerKeyTypeOutput) ToServerKeyTypeOutputWithContext(ctx context.Context) ServerKeyTypeOutput {
	return o
}

func (o ServerKeyTypeOutput) ToServerKeyTypePtrOutput() ServerKeyTypePtrOutput {
	return o.ToServerKeyTypePtrOutputWithContext(context.Background())
}

func (o ServerKeyTypeOutput) ToServerKeyTypePtrOutputWithContext(ctx context.Context) ServerKeyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerKeyType) *ServerKeyType {
		return &v
	}).(ServerKeyTypePtrOutput)
}

func (o ServerKeyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServerKeyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServerKeyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServerKeyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServerKeyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServerKeyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServerKeyTypePtrOutput struct{ *pulumi.OutputState }

func (ServerKeyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerKeyType)(nil)).Elem()
}

func (o ServerKeyTypePtrOutput) ToServerKeyTypePtrOutput() ServerKeyTypePtrOutput {
	return o
}

func (o ServerKeyTypePtrOutput) ToServerKeyTypePtrOutputWithContext(ctx context.Context) ServerKeyTypePtrOutput {
	return o
}

func (o ServerKeyTypePtrOutput) Elem() ServerKeyTypeOutput {
	return o.ApplyT(func(v *ServerKeyType) ServerKeyType {
		if v != nil {
			return *v
		}
		var ret ServerKeyType
		return ret
	}).(ServerKeyTypeOutput)
}

func (o ServerKeyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServerKeyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServerKeyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServerKeyTypeInput interface {
	pulumi.Input

	ToServerKeyTypeOutput() ServerKeyTypeOutput
	ToServerKeyTypeOutputWithContext(context.Context) ServerKeyTypeOutput
}

var serverKeyTypePtrType = reflect.TypeOf((**ServerKeyType)(nil)).Elem()

type ServerKeyTypePtrInput interface {
	pulumi.Input

	ToServerKeyTypePtrOutput() ServerKeyTypePtrOutput
	ToServerKeyTypePtrOutputWithContext(context.Context) ServerKeyTypePtrOutput
}

type serverKeyTypePtr string

func ServerKeyTypePtr(v string) ServerKeyTypePtrInput {
	return (*serverKeyTypePtr)(&v)
}

func (*serverKeyTypePtr) ElementType() reflect.Type {
	return serverKeyTypePtrType
}

func (in *serverKeyTypePtr) ToServerKeyTypePtrOutput() ServerKeyTypePtrOutput {
	return pulumi.ToOutput(in).(ServerKeyTypePtrOutput)
}

func (in *serverKeyTypePtr) ToServerKeyTypePtrOutputWithContext(ctx context.Context) ServerKeyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServerKeyTypePtrOutput)
}

type ServerSecurityAlertPolicyStateEnum string

const (
	ServerSecurityAlertPolicyStateEnumEnabled  = ServerSecurityAlertPolicyStateEnum("Enabled")
	ServerSecurityAlertPolicyStateEnumDisabled = ServerSecurityAlertPolicyStateEnum("Disabled")
)

func (ServerSecurityAlertPolicyStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerSecurityAlertPolicyStateEnum)(nil)).Elem()
}

func (e ServerSecurityAlertPolicyStateEnum) ToServerSecurityAlertPolicyStateEnumOutput() ServerSecurityAlertPolicyStateEnumOutput {
	return pulumi.ToOutput(e).(ServerSecurityAlertPolicyStateEnumOutput)
}

func (e ServerSecurityAlertPolicyStateEnum) ToServerSecurityAlertPolicyStateEnumOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyStateEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServerSecurityAlertPolicyStateEnumOutput)
}

func (e ServerSecurityAlertPolicyStateEnum) ToServerSecurityAlertPolicyStateEnumPtrOutput() ServerSecurityAlertPolicyStateEnumPtrOutput {
	return e.ToServerSecurityAlertPolicyStateEnumPtrOutputWithContext(context.Background())
}

func (e ServerSecurityAlertPolicyStateEnum) ToServerSecurityAlertPolicyStateEnumPtrOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyStateEnumPtrOutput {
	return ServerSecurityAlertPolicyStateEnum(e).ToServerSecurityAlertPolicyStateEnumOutputWithContext(ctx).ToServerSecurityAlertPolicyStateEnumPtrOutputWithContext(ctx)
}

func (e ServerSecurityAlertPolicyStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServerSecurityAlertPolicyStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServerSecurityAlertPolicyStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServerSecurityAlertPolicyStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServerSecurityAlertPolicyStateEnumOutput struct{ *pulumi.OutputState }

func (ServerSecurityAlertPolicyStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerSecurityAlertPolicyStateEnum)(nil)).Elem()
}

func (o ServerSecurityAlertPolicyStateEnumOutput) ToServerSecurityAlertPolicyStateEnumOutput() ServerSecurityAlertPolicyStateEnumOutput {
	return o
}

func (o ServerSecurityAlertPolicyStateEnumOutput) ToServerSecurityAlertPolicyStateEnumOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyStateEnumOutput {
	return o
}

func (o ServerSecurityAlertPolicyStateEnumOutput) ToServerSecurityAlertPolicyStateEnumPtrOutput() ServerSecurityAlertPolicyStateEnumPtrOutput {
	return o.ToServerSecurityAlertPolicyStateEnumPtrOutputWithContext(context.Background())
}

func (o ServerSecurityAlertPolicyStateEnumOutput) ToServerSecurityAlertPolicyStateEnumPtrOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerSecurityAlertPolicyStateEnum) *ServerSecurityAlertPolicyStateEnum {
		return &v
	}).(ServerSecurityAlertPolicyStateEnumPtrOutput)
}

func (o ServerSecurityAlertPolicyStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServerSecurityAlertPolicyStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServerSecurityAlertPolicyStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServerSecurityAlertPolicyStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServerSecurityAlertPolicyStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServerSecurityAlertPolicyStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServerSecurityAlertPolicyStateEnumPtrOutput struct{ *pulumi.OutputState }

func (ServerSecurityAlertPolicyStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerSecurityAlertPolicyStateEnum)(nil)).Elem()
}

func (o ServerSecurityAlertPolicyStateEnumPtrOutput) ToServerSecurityAlertPolicyStateEnumPtrOutput() ServerSecurityAlertPolicyStateEnumPtrOutput {
	return o
}

func (o ServerSecurityAlertPolicyStateEnumPtrOutput) ToServerSecurityAlertPolicyStateEnumPtrOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyStateEnumPtrOutput {
	return o
}

func (o ServerSecurityAlertPolicyStateEnumPtrOutput) Elem() ServerSecurityAlertPolicyStateEnumOutput {
	return o.ApplyT(func(v *ServerSecurityAlertPolicyStateEnum) ServerSecurityAlertPolicyStateEnum {
		if v != nil {
			return *v
		}
		var ret ServerSecurityAlertPolicyStateEnum
		return ret
	}).(ServerSecurityAlertPolicyStateEnumOutput)
}

func (o ServerSecurityAlertPolicyStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServerSecurityAlertPolicyStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServerSecurityAlertPolicyStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServerSecurityAlertPolicyStateEnumInput interface {
	pulumi.Input

	ToServerSecurityAlertPolicyStateEnumOutput() ServerSecurityAlertPolicyStateEnumOutput
	ToServerSecurityAlertPolicyStateEnumOutputWithContext(context.Context) ServerSecurityAlertPolicyStateEnumOutput
}

var serverSecurityAlertPolicyStateEnumPtrType = reflect.TypeOf((**ServerSecurityAlertPolicyStateEnum)(nil)).Elem()

type ServerSecurityAlertPolicyStateEnumPtrInput interface {
	pulumi.Input

	ToServerSecurityAlertPolicyStateEnumPtrOutput() ServerSecurityAlertPolicyStateEnumPtrOutput
	ToServerSecurityAlertPolicyStateEnumPtrOutputWithContext(context.Context) ServerSecurityAlertPolicyStateEnumPtrOutput
}

type serverSecurityAlertPolicyStateEnumPtr string

func ServerSecurityAlertPolicyStateEnumPtr(v string) ServerSecurityAlertPolicyStateEnumPtrInput {
	return (*serverSecurityAlertPolicyStateEnumPtr)(&v)
}

func (*serverSecurityAlertPolicyStateEnumPtr) ElementType() reflect.Type {
	return serverSecurityAlertPolicyStateEnumPtrType
}

func (in *serverSecurityAlertPolicyStateEnumPtr) ToServerSecurityAlertPolicyStateEnumPtrOutput() ServerSecurityAlertPolicyStateEnumPtrOutput {
	return pulumi.ToOutput(in).(ServerSecurityAlertPolicyStateEnumPtrOutput)
}

func (in *serverSecurityAlertPolicyStateEnumPtr) ToServerSecurityAlertPolicyStateEnumPtrOutputWithContext(ctx context.Context) ServerSecurityAlertPolicyStateEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServerSecurityAlertPolicyStateEnumPtrOutput)
}

type ServerVersion string

const (
	ServerVersion_9_5  = ServerVersion("9.5")
	ServerVersion_9_6  = ServerVersion("9.6")
	ServerVersion_10   = ServerVersion("10")
	ServerVersion_10_0 = ServerVersion("10.0")
	ServerVersion_10_2 = ServerVersion("10.2")
	ServerVersion_11   = ServerVersion("11")
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
	SkuTierBasic           = SkuTier("Basic")
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

type SslEnforcementEnum string

const (
	SslEnforcementEnumEnabled  = SslEnforcementEnum("Enabled")
	SslEnforcementEnumDisabled = SslEnforcementEnum("Disabled")
)

func (SslEnforcementEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*SslEnforcementEnum)(nil)).Elem()
}

func (e SslEnforcementEnum) ToSslEnforcementEnumOutput() SslEnforcementEnumOutput {
	return pulumi.ToOutput(e).(SslEnforcementEnumOutput)
}

func (e SslEnforcementEnum) ToSslEnforcementEnumOutputWithContext(ctx context.Context) SslEnforcementEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SslEnforcementEnumOutput)
}

func (e SslEnforcementEnum) ToSslEnforcementEnumPtrOutput() SslEnforcementEnumPtrOutput {
	return e.ToSslEnforcementEnumPtrOutputWithContext(context.Background())
}

func (e SslEnforcementEnum) ToSslEnforcementEnumPtrOutputWithContext(ctx context.Context) SslEnforcementEnumPtrOutput {
	return SslEnforcementEnum(e).ToSslEnforcementEnumOutputWithContext(ctx).ToSslEnforcementEnumPtrOutputWithContext(ctx)
}

func (e SslEnforcementEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslEnforcementEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslEnforcementEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SslEnforcementEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SslEnforcementEnumOutput struct{ *pulumi.OutputState }

func (SslEnforcementEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslEnforcementEnum)(nil)).Elem()
}

func (o SslEnforcementEnumOutput) ToSslEnforcementEnumOutput() SslEnforcementEnumOutput {
	return o
}

func (o SslEnforcementEnumOutput) ToSslEnforcementEnumOutputWithContext(ctx context.Context) SslEnforcementEnumOutput {
	return o
}

func (o SslEnforcementEnumOutput) ToSslEnforcementEnumPtrOutput() SslEnforcementEnumPtrOutput {
	return o.ToSslEnforcementEnumPtrOutputWithContext(context.Background())
}

func (o SslEnforcementEnumOutput) ToSslEnforcementEnumPtrOutputWithContext(ctx context.Context) SslEnforcementEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SslEnforcementEnum) *SslEnforcementEnum {
		return &v
	}).(SslEnforcementEnumPtrOutput)
}

func (o SslEnforcementEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SslEnforcementEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslEnforcementEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SslEnforcementEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslEnforcementEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslEnforcementEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SslEnforcementEnumPtrOutput struct{ *pulumi.OutputState }

func (SslEnforcementEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslEnforcementEnum)(nil)).Elem()
}

func (o SslEnforcementEnumPtrOutput) ToSslEnforcementEnumPtrOutput() SslEnforcementEnumPtrOutput {
	return o
}

func (o SslEnforcementEnumPtrOutput) ToSslEnforcementEnumPtrOutputWithContext(ctx context.Context) SslEnforcementEnumPtrOutput {
	return o
}

func (o SslEnforcementEnumPtrOutput) Elem() SslEnforcementEnumOutput {
	return o.ApplyT(func(v *SslEnforcementEnum) SslEnforcementEnum {
		if v != nil {
			return *v
		}
		var ret SslEnforcementEnum
		return ret
	}).(SslEnforcementEnumOutput)
}

func (o SslEnforcementEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslEnforcementEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SslEnforcementEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SslEnforcementEnumInput interface {
	pulumi.Input

	ToSslEnforcementEnumOutput() SslEnforcementEnumOutput
	ToSslEnforcementEnumOutputWithContext(context.Context) SslEnforcementEnumOutput
}

var sslEnforcementEnumPtrType = reflect.TypeOf((**SslEnforcementEnum)(nil)).Elem()

type SslEnforcementEnumPtrInput interface {
	pulumi.Input

	ToSslEnforcementEnumPtrOutput() SslEnforcementEnumPtrOutput
	ToSslEnforcementEnumPtrOutputWithContext(context.Context) SslEnforcementEnumPtrOutput
}

type sslEnforcementEnumPtr string

func SslEnforcementEnumPtr(v string) SslEnforcementEnumPtrInput {
	return (*sslEnforcementEnumPtr)(&v)
}

func (*sslEnforcementEnumPtr) ElementType() reflect.Type {
	return sslEnforcementEnumPtrType
}

func (in *sslEnforcementEnumPtr) ToSslEnforcementEnumPtrOutput() SslEnforcementEnumPtrOutput {
	return pulumi.ToOutput(in).(SslEnforcementEnumPtrOutput)
}

func (in *sslEnforcementEnumPtr) ToSslEnforcementEnumPtrOutputWithContext(ctx context.Context) SslEnforcementEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SslEnforcementEnumPtrOutput)
}

type StorageAutogrow string

const (
	StorageAutogrowEnabled  = StorageAutogrow("Enabled")
	StorageAutogrowDisabled = StorageAutogrow("Disabled")
)

func (StorageAutogrow) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAutogrow)(nil)).Elem()
}

func (e StorageAutogrow) ToStorageAutogrowOutput() StorageAutogrowOutput {
	return pulumi.ToOutput(e).(StorageAutogrowOutput)
}

func (e StorageAutogrow) ToStorageAutogrowOutputWithContext(ctx context.Context) StorageAutogrowOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StorageAutogrowOutput)
}

func (e StorageAutogrow) ToStorageAutogrowPtrOutput() StorageAutogrowPtrOutput {
	return e.ToStorageAutogrowPtrOutputWithContext(context.Background())
}

func (e StorageAutogrow) ToStorageAutogrowPtrOutputWithContext(ctx context.Context) StorageAutogrowPtrOutput {
	return StorageAutogrow(e).ToStorageAutogrowOutputWithContext(ctx).ToStorageAutogrowPtrOutputWithContext(ctx)
}

func (e StorageAutogrow) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageAutogrow) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageAutogrow) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StorageAutogrow) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StorageAutogrowOutput struct{ *pulumi.OutputState }

func (StorageAutogrowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAutogrow)(nil)).Elem()
}

func (o StorageAutogrowOutput) ToStorageAutogrowOutput() StorageAutogrowOutput {
	return o
}

func (o StorageAutogrowOutput) ToStorageAutogrowOutputWithContext(ctx context.Context) StorageAutogrowOutput {
	return o
}

func (o StorageAutogrowOutput) ToStorageAutogrowPtrOutput() StorageAutogrowPtrOutput {
	return o.ToStorageAutogrowPtrOutputWithContext(context.Background())
}

func (o StorageAutogrowOutput) ToStorageAutogrowPtrOutputWithContext(ctx context.Context) StorageAutogrowPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAutogrow) *StorageAutogrow {
		return &v
	}).(StorageAutogrowPtrOutput)
}

func (o StorageAutogrowOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StorageAutogrowOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageAutogrow) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StorageAutogrowOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageAutogrowOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageAutogrow) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StorageAutogrowPtrOutput struct{ *pulumi.OutputState }

func (StorageAutogrowPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAutogrow)(nil)).Elem()
}

func (o StorageAutogrowPtrOutput) ToStorageAutogrowPtrOutput() StorageAutogrowPtrOutput {
	return o
}

func (o StorageAutogrowPtrOutput) ToStorageAutogrowPtrOutputWithContext(ctx context.Context) StorageAutogrowPtrOutput {
	return o
}

func (o StorageAutogrowPtrOutput) Elem() StorageAutogrowOutput {
	return o.ApplyT(func(v *StorageAutogrow) StorageAutogrow {
		if v != nil {
			return *v
		}
		var ret StorageAutogrow
		return ret
	}).(StorageAutogrowOutput)
}

func (o StorageAutogrowPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageAutogrowPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StorageAutogrow) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StorageAutogrowInput interface {
	pulumi.Input

	ToStorageAutogrowOutput() StorageAutogrowOutput
	ToStorageAutogrowOutputWithContext(context.Context) StorageAutogrowOutput
}

var storageAutogrowPtrType = reflect.TypeOf((**StorageAutogrow)(nil)).Elem()

type StorageAutogrowPtrInput interface {
	pulumi.Input

	ToStorageAutogrowPtrOutput() StorageAutogrowPtrOutput
	ToStorageAutogrowPtrOutputWithContext(context.Context) StorageAutogrowPtrOutput
}

type storageAutogrowPtr string

func StorageAutogrowPtr(v string) StorageAutogrowPtrInput {
	return (*storageAutogrowPtr)(&v)
}

func (*storageAutogrowPtr) ElementType() reflect.Type {
	return storageAutogrowPtrType
}

func (in *storageAutogrowPtr) ToStorageAutogrowPtrOutput() StorageAutogrowPtrOutput {
	return pulumi.ToOutput(in).(StorageAutogrowPtrOutput)
}

func (in *storageAutogrowPtr) ToStorageAutogrowPtrOutputWithContext(ctx context.Context) StorageAutogrowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StorageAutogrowPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AdministratorTypeOutput{})
	pulumi.RegisterOutputType(AdministratorTypePtrOutput{})
	pulumi.RegisterOutputType(CreateModeOutput{})
	pulumi.RegisterOutputType(CreateModePtrOutput{})
	pulumi.RegisterOutputType(GeoRedundantBackupOutput{})
	pulumi.RegisterOutputType(GeoRedundantBackupPtrOutput{})
	pulumi.RegisterOutputType(IdentityTypeOutput{})
	pulumi.RegisterOutputType(IdentityTypePtrOutput{})
	pulumi.RegisterOutputType(InfrastructureEncryptionOutput{})
	pulumi.RegisterOutputType(InfrastructureEncryptionPtrOutput{})
	pulumi.RegisterOutputType(MinimalTlsVersionEnumOutput{})
	pulumi.RegisterOutputType(MinimalTlsVersionEnumPtrOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessEnumOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessEnumPtrOutput{})
	pulumi.RegisterOutputType(ServerKeyTypeOutput{})
	pulumi.RegisterOutputType(ServerKeyTypePtrOutput{})
	pulumi.RegisterOutputType(ServerSecurityAlertPolicyStateEnumOutput{})
	pulumi.RegisterOutputType(ServerSecurityAlertPolicyStateEnumPtrOutput{})
	pulumi.RegisterOutputType(ServerVersionOutput{})
	pulumi.RegisterOutputType(ServerVersionPtrOutput{})
	pulumi.RegisterOutputType(SkuTierOutput{})
	pulumi.RegisterOutputType(SkuTierPtrOutput{})
	pulumi.RegisterOutputType(SslEnforcementEnumOutput{})
	pulumi.RegisterOutputType(SslEnforcementEnumPtrOutput{})
	pulumi.RegisterOutputType(StorageAutogrowOutput{})
	pulumi.RegisterOutputType(StorageAutogrowPtrOutput{})
}
