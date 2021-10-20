


package v20180601preview

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

type CatalogCollationType string

const (
	CatalogCollationType_DATABASE_DEFAULT             = CatalogCollationType("DATABASE_DEFAULT")
	CatalogCollationType_SQL_Latin1_General_CP1_CI_AS = CatalogCollationType("SQL_Latin1_General_CP1_CI_AS")
)

func (CatalogCollationType) ElementType() reflect.Type {
	return reflect.TypeOf((*CatalogCollationType)(nil)).Elem()
}

func (e CatalogCollationType) ToCatalogCollationTypeOutput() CatalogCollationTypeOutput {
	return pulumi.ToOutput(e).(CatalogCollationTypeOutput)
}

func (e CatalogCollationType) ToCatalogCollationTypeOutputWithContext(ctx context.Context) CatalogCollationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CatalogCollationTypeOutput)
}

func (e CatalogCollationType) ToCatalogCollationTypePtrOutput() CatalogCollationTypePtrOutput {
	return e.ToCatalogCollationTypePtrOutputWithContext(context.Background())
}

func (e CatalogCollationType) ToCatalogCollationTypePtrOutputWithContext(ctx context.Context) CatalogCollationTypePtrOutput {
	return CatalogCollationType(e).ToCatalogCollationTypeOutputWithContext(ctx).ToCatalogCollationTypePtrOutputWithContext(ctx)
}

func (e CatalogCollationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CatalogCollationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CatalogCollationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CatalogCollationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CatalogCollationTypeOutput struct{ *pulumi.OutputState }

func (CatalogCollationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CatalogCollationType)(nil)).Elem()
}

func (o CatalogCollationTypeOutput) ToCatalogCollationTypeOutput() CatalogCollationTypeOutput {
	return o
}

func (o CatalogCollationTypeOutput) ToCatalogCollationTypeOutputWithContext(ctx context.Context) CatalogCollationTypeOutput {
	return o
}

func (o CatalogCollationTypeOutput) ToCatalogCollationTypePtrOutput() CatalogCollationTypePtrOutput {
	return o.ToCatalogCollationTypePtrOutputWithContext(context.Background())
}

func (o CatalogCollationTypeOutput) ToCatalogCollationTypePtrOutputWithContext(ctx context.Context) CatalogCollationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CatalogCollationType) *CatalogCollationType {
		return &v
	}).(CatalogCollationTypePtrOutput)
}

func (o CatalogCollationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CatalogCollationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CatalogCollationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CatalogCollationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CatalogCollationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CatalogCollationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CatalogCollationTypePtrOutput struct{ *pulumi.OutputState }

func (CatalogCollationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CatalogCollationType)(nil)).Elem()
}

func (o CatalogCollationTypePtrOutput) ToCatalogCollationTypePtrOutput() CatalogCollationTypePtrOutput {
	return o
}

func (o CatalogCollationTypePtrOutput) ToCatalogCollationTypePtrOutputWithContext(ctx context.Context) CatalogCollationTypePtrOutput {
	return o
}

func (o CatalogCollationTypePtrOutput) Elem() CatalogCollationTypeOutput {
	return o.ApplyT(func(v *CatalogCollationType) CatalogCollationType {
		if v != nil {
			return *v
		}
		var ret CatalogCollationType
		return ret
	}).(CatalogCollationTypeOutput)
}

func (o CatalogCollationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CatalogCollationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CatalogCollationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CatalogCollationTypeInput interface {
	pulumi.Input

	ToCatalogCollationTypeOutput() CatalogCollationTypeOutput
	ToCatalogCollationTypeOutputWithContext(context.Context) CatalogCollationTypeOutput
}

var catalogCollationTypePtrType = reflect.TypeOf((**CatalogCollationType)(nil)).Elem()

type CatalogCollationTypePtrInput interface {
	pulumi.Input

	ToCatalogCollationTypePtrOutput() CatalogCollationTypePtrOutput
	ToCatalogCollationTypePtrOutputWithContext(context.Context) CatalogCollationTypePtrOutput
}

type catalogCollationTypePtr string

func CatalogCollationTypePtr(v string) CatalogCollationTypePtrInput {
	return (*catalogCollationTypePtr)(&v)
}

func (*catalogCollationTypePtr) ElementType() reflect.Type {
	return catalogCollationTypePtrType
}

func (in *catalogCollationTypePtr) ToCatalogCollationTypePtrOutput() CatalogCollationTypePtrOutput {
	return pulumi.ToOutput(in).(CatalogCollationTypePtrOutput)
}

func (in *catalogCollationTypePtr) ToCatalogCollationTypePtrOutputWithContext(ctx context.Context) CatalogCollationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CatalogCollationTypePtrOutput)
}

type IdentityType string

const (
	IdentityTypeNone                         = IdentityType("None")
	IdentityTypeSystemAssigned               = IdentityType("SystemAssigned")
	IdentityTypeUserAssigned                 = IdentityType("UserAssigned")
	IdentityType_SystemAssigned_UserAssigned = IdentityType("SystemAssigned,UserAssigned")
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

type InstancePoolLicenseType string

const (
	InstancePoolLicenseTypeLicenseIncluded = InstancePoolLicenseType("LicenseIncluded")
	InstancePoolLicenseTypeBasePrice       = InstancePoolLicenseType("BasePrice")
)

func (InstancePoolLicenseType) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancePoolLicenseType)(nil)).Elem()
}

func (e InstancePoolLicenseType) ToInstancePoolLicenseTypeOutput() InstancePoolLicenseTypeOutput {
	return pulumi.ToOutput(e).(InstancePoolLicenseTypeOutput)
}

func (e InstancePoolLicenseType) ToInstancePoolLicenseTypeOutputWithContext(ctx context.Context) InstancePoolLicenseTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InstancePoolLicenseTypeOutput)
}

func (e InstancePoolLicenseType) ToInstancePoolLicenseTypePtrOutput() InstancePoolLicenseTypePtrOutput {
	return e.ToInstancePoolLicenseTypePtrOutputWithContext(context.Background())
}

func (e InstancePoolLicenseType) ToInstancePoolLicenseTypePtrOutputWithContext(ctx context.Context) InstancePoolLicenseTypePtrOutput {
	return InstancePoolLicenseType(e).ToInstancePoolLicenseTypeOutputWithContext(ctx).ToInstancePoolLicenseTypePtrOutputWithContext(ctx)
}

func (e InstancePoolLicenseType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstancePoolLicenseType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstancePoolLicenseType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InstancePoolLicenseType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InstancePoolLicenseTypeOutput struct{ *pulumi.OutputState }

func (InstancePoolLicenseTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancePoolLicenseType)(nil)).Elem()
}

func (o InstancePoolLicenseTypeOutput) ToInstancePoolLicenseTypeOutput() InstancePoolLicenseTypeOutput {
	return o
}

func (o InstancePoolLicenseTypeOutput) ToInstancePoolLicenseTypeOutputWithContext(ctx context.Context) InstancePoolLicenseTypeOutput {
	return o
}

func (o InstancePoolLicenseTypeOutput) ToInstancePoolLicenseTypePtrOutput() InstancePoolLicenseTypePtrOutput {
	return o.ToInstancePoolLicenseTypePtrOutputWithContext(context.Background())
}

func (o InstancePoolLicenseTypeOutput) ToInstancePoolLicenseTypePtrOutputWithContext(ctx context.Context) InstancePoolLicenseTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstancePoolLicenseType) *InstancePoolLicenseType {
		return &v
	}).(InstancePoolLicenseTypePtrOutput)
}

func (o InstancePoolLicenseTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InstancePoolLicenseTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InstancePoolLicenseType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InstancePoolLicenseTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InstancePoolLicenseTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InstancePoolLicenseType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InstancePoolLicenseTypePtrOutput struct{ *pulumi.OutputState }

func (InstancePoolLicenseTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstancePoolLicenseType)(nil)).Elem()
}

func (o InstancePoolLicenseTypePtrOutput) ToInstancePoolLicenseTypePtrOutput() InstancePoolLicenseTypePtrOutput {
	return o
}

func (o InstancePoolLicenseTypePtrOutput) ToInstancePoolLicenseTypePtrOutputWithContext(ctx context.Context) InstancePoolLicenseTypePtrOutput {
	return o
}

func (o InstancePoolLicenseTypePtrOutput) Elem() InstancePoolLicenseTypeOutput {
	return o.ApplyT(func(v *InstancePoolLicenseType) InstancePoolLicenseType {
		if v != nil {
			return *v
		}
		var ret InstancePoolLicenseType
		return ret
	}).(InstancePoolLicenseTypeOutput)
}

func (o InstancePoolLicenseTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InstancePoolLicenseTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InstancePoolLicenseType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type InstancePoolLicenseTypeInput interface {
	pulumi.Input

	ToInstancePoolLicenseTypeOutput() InstancePoolLicenseTypeOutput
	ToInstancePoolLicenseTypeOutputWithContext(context.Context) InstancePoolLicenseTypeOutput
}

var instancePoolLicenseTypePtrType = reflect.TypeOf((**InstancePoolLicenseType)(nil)).Elem()

type InstancePoolLicenseTypePtrInput interface {
	pulumi.Input

	ToInstancePoolLicenseTypePtrOutput() InstancePoolLicenseTypePtrOutput
	ToInstancePoolLicenseTypePtrOutputWithContext(context.Context) InstancePoolLicenseTypePtrOutput
}

type instancePoolLicenseTypePtr string

func InstancePoolLicenseTypePtr(v string) InstancePoolLicenseTypePtrInput {
	return (*instancePoolLicenseTypePtr)(&v)
}

func (*instancePoolLicenseTypePtr) ElementType() reflect.Type {
	return instancePoolLicenseTypePtrType
}

func (in *instancePoolLicenseTypePtr) ToInstancePoolLicenseTypePtrOutput() InstancePoolLicenseTypePtrOutput {
	return pulumi.ToOutput(in).(InstancePoolLicenseTypePtrOutput)
}

func (in *instancePoolLicenseTypePtr) ToInstancePoolLicenseTypePtrOutputWithContext(ctx context.Context) InstancePoolLicenseTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InstancePoolLicenseTypePtrOutput)
}

type ManagedDatabaseCreateMode string

const (
	ManagedDatabaseCreateModeDefault                        = ManagedDatabaseCreateMode("Default")
	ManagedDatabaseCreateModeRestoreExternalBackup          = ManagedDatabaseCreateMode("RestoreExternalBackup")
	ManagedDatabaseCreateModePointInTimeRestore             = ManagedDatabaseCreateMode("PointInTimeRestore")
	ManagedDatabaseCreateModeRecovery                       = ManagedDatabaseCreateMode("Recovery")
	ManagedDatabaseCreateModeRestoreLongTermRetentionBackup = ManagedDatabaseCreateMode("RestoreLongTermRetentionBackup")
)

func (ManagedDatabaseCreateMode) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedDatabaseCreateMode)(nil)).Elem()
}

func (e ManagedDatabaseCreateMode) ToManagedDatabaseCreateModeOutput() ManagedDatabaseCreateModeOutput {
	return pulumi.ToOutput(e).(ManagedDatabaseCreateModeOutput)
}

func (e ManagedDatabaseCreateMode) ToManagedDatabaseCreateModeOutputWithContext(ctx context.Context) ManagedDatabaseCreateModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedDatabaseCreateModeOutput)
}

func (e ManagedDatabaseCreateMode) ToManagedDatabaseCreateModePtrOutput() ManagedDatabaseCreateModePtrOutput {
	return e.ToManagedDatabaseCreateModePtrOutputWithContext(context.Background())
}

func (e ManagedDatabaseCreateMode) ToManagedDatabaseCreateModePtrOutputWithContext(ctx context.Context) ManagedDatabaseCreateModePtrOutput {
	return ManagedDatabaseCreateMode(e).ToManagedDatabaseCreateModeOutputWithContext(ctx).ToManagedDatabaseCreateModePtrOutputWithContext(ctx)
}

func (e ManagedDatabaseCreateMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedDatabaseCreateMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedDatabaseCreateMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedDatabaseCreateMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedDatabaseCreateModeOutput struct{ *pulumi.OutputState }

func (ManagedDatabaseCreateModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedDatabaseCreateMode)(nil)).Elem()
}

func (o ManagedDatabaseCreateModeOutput) ToManagedDatabaseCreateModeOutput() ManagedDatabaseCreateModeOutput {
	return o
}

func (o ManagedDatabaseCreateModeOutput) ToManagedDatabaseCreateModeOutputWithContext(ctx context.Context) ManagedDatabaseCreateModeOutput {
	return o
}

func (o ManagedDatabaseCreateModeOutput) ToManagedDatabaseCreateModePtrOutput() ManagedDatabaseCreateModePtrOutput {
	return o.ToManagedDatabaseCreateModePtrOutputWithContext(context.Background())
}

func (o ManagedDatabaseCreateModeOutput) ToManagedDatabaseCreateModePtrOutputWithContext(ctx context.Context) ManagedDatabaseCreateModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedDatabaseCreateMode) *ManagedDatabaseCreateMode {
		return &v
	}).(ManagedDatabaseCreateModePtrOutput)
}

func (o ManagedDatabaseCreateModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedDatabaseCreateModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedDatabaseCreateMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedDatabaseCreateModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedDatabaseCreateModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedDatabaseCreateMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedDatabaseCreateModePtrOutput struct{ *pulumi.OutputState }

func (ManagedDatabaseCreateModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedDatabaseCreateMode)(nil)).Elem()
}

func (o ManagedDatabaseCreateModePtrOutput) ToManagedDatabaseCreateModePtrOutput() ManagedDatabaseCreateModePtrOutput {
	return o
}

func (o ManagedDatabaseCreateModePtrOutput) ToManagedDatabaseCreateModePtrOutputWithContext(ctx context.Context) ManagedDatabaseCreateModePtrOutput {
	return o
}

func (o ManagedDatabaseCreateModePtrOutput) Elem() ManagedDatabaseCreateModeOutput {
	return o.ApplyT(func(v *ManagedDatabaseCreateMode) ManagedDatabaseCreateMode {
		if v != nil {
			return *v
		}
		var ret ManagedDatabaseCreateMode
		return ret
	}).(ManagedDatabaseCreateModeOutput)
}

func (o ManagedDatabaseCreateModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedDatabaseCreateModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedDatabaseCreateMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ManagedDatabaseCreateModeInput interface {
	pulumi.Input

	ToManagedDatabaseCreateModeOutput() ManagedDatabaseCreateModeOutput
	ToManagedDatabaseCreateModeOutputWithContext(context.Context) ManagedDatabaseCreateModeOutput
}

var managedDatabaseCreateModePtrType = reflect.TypeOf((**ManagedDatabaseCreateMode)(nil)).Elem()

type ManagedDatabaseCreateModePtrInput interface {
	pulumi.Input

	ToManagedDatabaseCreateModePtrOutput() ManagedDatabaseCreateModePtrOutput
	ToManagedDatabaseCreateModePtrOutputWithContext(context.Context) ManagedDatabaseCreateModePtrOutput
}

type managedDatabaseCreateModePtr string

func ManagedDatabaseCreateModePtr(v string) ManagedDatabaseCreateModePtrInput {
	return (*managedDatabaseCreateModePtr)(&v)
}

func (*managedDatabaseCreateModePtr) ElementType() reflect.Type {
	return managedDatabaseCreateModePtrType
}

func (in *managedDatabaseCreateModePtr) ToManagedDatabaseCreateModePtrOutput() ManagedDatabaseCreateModePtrOutput {
	return pulumi.ToOutput(in).(ManagedDatabaseCreateModePtrOutput)
}

func (in *managedDatabaseCreateModePtr) ToManagedDatabaseCreateModePtrOutputWithContext(ctx context.Context) ManagedDatabaseCreateModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedDatabaseCreateModePtrOutput)
}

type ManagedInstanceLicenseType string

const (
	ManagedInstanceLicenseTypeLicenseIncluded = ManagedInstanceLicenseType("LicenseIncluded")
	ManagedInstanceLicenseTypeBasePrice       = ManagedInstanceLicenseType("BasePrice")
)

func (ManagedInstanceLicenseType) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstanceLicenseType)(nil)).Elem()
}

func (e ManagedInstanceLicenseType) ToManagedInstanceLicenseTypeOutput() ManagedInstanceLicenseTypeOutput {
	return pulumi.ToOutput(e).(ManagedInstanceLicenseTypeOutput)
}

func (e ManagedInstanceLicenseType) ToManagedInstanceLicenseTypeOutputWithContext(ctx context.Context) ManagedInstanceLicenseTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedInstanceLicenseTypeOutput)
}

func (e ManagedInstanceLicenseType) ToManagedInstanceLicenseTypePtrOutput() ManagedInstanceLicenseTypePtrOutput {
	return e.ToManagedInstanceLicenseTypePtrOutputWithContext(context.Background())
}

func (e ManagedInstanceLicenseType) ToManagedInstanceLicenseTypePtrOutputWithContext(ctx context.Context) ManagedInstanceLicenseTypePtrOutput {
	return ManagedInstanceLicenseType(e).ToManagedInstanceLicenseTypeOutputWithContext(ctx).ToManagedInstanceLicenseTypePtrOutputWithContext(ctx)
}

func (e ManagedInstanceLicenseType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedInstanceLicenseType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedInstanceLicenseType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedInstanceLicenseType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedInstanceLicenseTypeOutput struct{ *pulumi.OutputState }

func (ManagedInstanceLicenseTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstanceLicenseType)(nil)).Elem()
}

func (o ManagedInstanceLicenseTypeOutput) ToManagedInstanceLicenseTypeOutput() ManagedInstanceLicenseTypeOutput {
	return o
}

func (o ManagedInstanceLicenseTypeOutput) ToManagedInstanceLicenseTypeOutputWithContext(ctx context.Context) ManagedInstanceLicenseTypeOutput {
	return o
}

func (o ManagedInstanceLicenseTypeOutput) ToManagedInstanceLicenseTypePtrOutput() ManagedInstanceLicenseTypePtrOutput {
	return o.ToManagedInstanceLicenseTypePtrOutputWithContext(context.Background())
}

func (o ManagedInstanceLicenseTypeOutput) ToManagedInstanceLicenseTypePtrOutputWithContext(ctx context.Context) ManagedInstanceLicenseTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedInstanceLicenseType) *ManagedInstanceLicenseType {
		return &v
	}).(ManagedInstanceLicenseTypePtrOutput)
}

func (o ManagedInstanceLicenseTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedInstanceLicenseTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedInstanceLicenseType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedInstanceLicenseTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedInstanceLicenseTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedInstanceLicenseType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedInstanceLicenseTypePtrOutput struct{ *pulumi.OutputState }

func (ManagedInstanceLicenseTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceLicenseType)(nil)).Elem()
}

func (o ManagedInstanceLicenseTypePtrOutput) ToManagedInstanceLicenseTypePtrOutput() ManagedInstanceLicenseTypePtrOutput {
	return o
}

func (o ManagedInstanceLicenseTypePtrOutput) ToManagedInstanceLicenseTypePtrOutputWithContext(ctx context.Context) ManagedInstanceLicenseTypePtrOutput {
	return o
}

func (o ManagedInstanceLicenseTypePtrOutput) Elem() ManagedInstanceLicenseTypeOutput {
	return o.ApplyT(func(v *ManagedInstanceLicenseType) ManagedInstanceLicenseType {
		if v != nil {
			return *v
		}
		var ret ManagedInstanceLicenseType
		return ret
	}).(ManagedInstanceLicenseTypeOutput)
}

func (o ManagedInstanceLicenseTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedInstanceLicenseTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedInstanceLicenseType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ManagedInstanceLicenseTypeInput interface {
	pulumi.Input

	ToManagedInstanceLicenseTypeOutput() ManagedInstanceLicenseTypeOutput
	ToManagedInstanceLicenseTypeOutputWithContext(context.Context) ManagedInstanceLicenseTypeOutput
}

var managedInstanceLicenseTypePtrType = reflect.TypeOf((**ManagedInstanceLicenseType)(nil)).Elem()

type ManagedInstanceLicenseTypePtrInput interface {
	pulumi.Input

	ToManagedInstanceLicenseTypePtrOutput() ManagedInstanceLicenseTypePtrOutput
	ToManagedInstanceLicenseTypePtrOutputWithContext(context.Context) ManagedInstanceLicenseTypePtrOutput
}

type managedInstanceLicenseTypePtr string

func ManagedInstanceLicenseTypePtr(v string) ManagedInstanceLicenseTypePtrInput {
	return (*managedInstanceLicenseTypePtr)(&v)
}

func (*managedInstanceLicenseTypePtr) ElementType() reflect.Type {
	return managedInstanceLicenseTypePtrType
}

func (in *managedInstanceLicenseTypePtr) ToManagedInstanceLicenseTypePtrOutput() ManagedInstanceLicenseTypePtrOutput {
	return pulumi.ToOutput(in).(ManagedInstanceLicenseTypePtrOutput)
}

func (in *managedInstanceLicenseTypePtr) ToManagedInstanceLicenseTypePtrOutputWithContext(ctx context.Context) ManagedInstanceLicenseTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedInstanceLicenseTypePtrOutput)
}

type ManagedInstanceProxyOverride string

const (
	ManagedInstanceProxyOverrideProxy    = ManagedInstanceProxyOverride("Proxy")
	ManagedInstanceProxyOverrideRedirect = ManagedInstanceProxyOverride("Redirect")
	ManagedInstanceProxyOverrideDefault  = ManagedInstanceProxyOverride("Default")
)

func (ManagedInstanceProxyOverride) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstanceProxyOverride)(nil)).Elem()
}

func (e ManagedInstanceProxyOverride) ToManagedInstanceProxyOverrideOutput() ManagedInstanceProxyOverrideOutput {
	return pulumi.ToOutput(e).(ManagedInstanceProxyOverrideOutput)
}

func (e ManagedInstanceProxyOverride) ToManagedInstanceProxyOverrideOutputWithContext(ctx context.Context) ManagedInstanceProxyOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedInstanceProxyOverrideOutput)
}

func (e ManagedInstanceProxyOverride) ToManagedInstanceProxyOverridePtrOutput() ManagedInstanceProxyOverridePtrOutput {
	return e.ToManagedInstanceProxyOverridePtrOutputWithContext(context.Background())
}

func (e ManagedInstanceProxyOverride) ToManagedInstanceProxyOverridePtrOutputWithContext(ctx context.Context) ManagedInstanceProxyOverridePtrOutput {
	return ManagedInstanceProxyOverride(e).ToManagedInstanceProxyOverrideOutputWithContext(ctx).ToManagedInstanceProxyOverridePtrOutputWithContext(ctx)
}

func (e ManagedInstanceProxyOverride) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedInstanceProxyOverride) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedInstanceProxyOverride) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedInstanceProxyOverride) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedInstanceProxyOverrideOutput struct{ *pulumi.OutputState }

func (ManagedInstanceProxyOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstanceProxyOverride)(nil)).Elem()
}

func (o ManagedInstanceProxyOverrideOutput) ToManagedInstanceProxyOverrideOutput() ManagedInstanceProxyOverrideOutput {
	return o
}

func (o ManagedInstanceProxyOverrideOutput) ToManagedInstanceProxyOverrideOutputWithContext(ctx context.Context) ManagedInstanceProxyOverrideOutput {
	return o
}

func (o ManagedInstanceProxyOverrideOutput) ToManagedInstanceProxyOverridePtrOutput() ManagedInstanceProxyOverridePtrOutput {
	return o.ToManagedInstanceProxyOverridePtrOutputWithContext(context.Background())
}

func (o ManagedInstanceProxyOverrideOutput) ToManagedInstanceProxyOverridePtrOutputWithContext(ctx context.Context) ManagedInstanceProxyOverridePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedInstanceProxyOverride) *ManagedInstanceProxyOverride {
		return &v
	}).(ManagedInstanceProxyOverridePtrOutput)
}

func (o ManagedInstanceProxyOverrideOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedInstanceProxyOverrideOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedInstanceProxyOverride) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedInstanceProxyOverrideOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedInstanceProxyOverrideOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedInstanceProxyOverride) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedInstanceProxyOverridePtrOutput struct{ *pulumi.OutputState }

func (ManagedInstanceProxyOverridePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceProxyOverride)(nil)).Elem()
}

func (o ManagedInstanceProxyOverridePtrOutput) ToManagedInstanceProxyOverridePtrOutput() ManagedInstanceProxyOverridePtrOutput {
	return o
}

func (o ManagedInstanceProxyOverridePtrOutput) ToManagedInstanceProxyOverridePtrOutputWithContext(ctx context.Context) ManagedInstanceProxyOverridePtrOutput {
	return o
}

func (o ManagedInstanceProxyOverridePtrOutput) Elem() ManagedInstanceProxyOverrideOutput {
	return o.ApplyT(func(v *ManagedInstanceProxyOverride) ManagedInstanceProxyOverride {
		if v != nil {
			return *v
		}
		var ret ManagedInstanceProxyOverride
		return ret
	}).(ManagedInstanceProxyOverrideOutput)
}

func (o ManagedInstanceProxyOverridePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedInstanceProxyOverridePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedInstanceProxyOverride) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ManagedInstanceProxyOverrideInput interface {
	pulumi.Input

	ToManagedInstanceProxyOverrideOutput() ManagedInstanceProxyOverrideOutput
	ToManagedInstanceProxyOverrideOutputWithContext(context.Context) ManagedInstanceProxyOverrideOutput
}

var managedInstanceProxyOverridePtrType = reflect.TypeOf((**ManagedInstanceProxyOverride)(nil)).Elem()

type ManagedInstanceProxyOverridePtrInput interface {
	pulumi.Input

	ToManagedInstanceProxyOverridePtrOutput() ManagedInstanceProxyOverridePtrOutput
	ToManagedInstanceProxyOverridePtrOutputWithContext(context.Context) ManagedInstanceProxyOverridePtrOutput
}

type managedInstanceProxyOverridePtr string

func ManagedInstanceProxyOverridePtr(v string) ManagedInstanceProxyOverridePtrInput {
	return (*managedInstanceProxyOverridePtr)(&v)
}

func (*managedInstanceProxyOverridePtr) ElementType() reflect.Type {
	return managedInstanceProxyOverridePtrType
}

func (in *managedInstanceProxyOverridePtr) ToManagedInstanceProxyOverridePtrOutput() ManagedInstanceProxyOverridePtrOutput {
	return pulumi.ToOutput(in).(ManagedInstanceProxyOverridePtrOutput)
}

func (in *managedInstanceProxyOverridePtr) ToManagedInstanceProxyOverridePtrOutputWithContext(ctx context.Context) ManagedInstanceProxyOverridePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedInstanceProxyOverridePtrOutput)
}

type ManagedServerCreateMode string

const (
	ManagedServerCreateModeDefault            = ManagedServerCreateMode("Default")
	ManagedServerCreateModePointInTimeRestore = ManagedServerCreateMode("PointInTimeRestore")
)

func (ManagedServerCreateMode) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServerCreateMode)(nil)).Elem()
}

func (e ManagedServerCreateMode) ToManagedServerCreateModeOutput() ManagedServerCreateModeOutput {
	return pulumi.ToOutput(e).(ManagedServerCreateModeOutput)
}

func (e ManagedServerCreateMode) ToManagedServerCreateModeOutputWithContext(ctx context.Context) ManagedServerCreateModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedServerCreateModeOutput)
}

func (e ManagedServerCreateMode) ToManagedServerCreateModePtrOutput() ManagedServerCreateModePtrOutput {
	return e.ToManagedServerCreateModePtrOutputWithContext(context.Background())
}

func (e ManagedServerCreateMode) ToManagedServerCreateModePtrOutputWithContext(ctx context.Context) ManagedServerCreateModePtrOutput {
	return ManagedServerCreateMode(e).ToManagedServerCreateModeOutputWithContext(ctx).ToManagedServerCreateModePtrOutputWithContext(ctx)
}

func (e ManagedServerCreateMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedServerCreateMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedServerCreateMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedServerCreateMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedServerCreateModeOutput struct{ *pulumi.OutputState }

func (ManagedServerCreateModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServerCreateMode)(nil)).Elem()
}

func (o ManagedServerCreateModeOutput) ToManagedServerCreateModeOutput() ManagedServerCreateModeOutput {
	return o
}

func (o ManagedServerCreateModeOutput) ToManagedServerCreateModeOutputWithContext(ctx context.Context) ManagedServerCreateModeOutput {
	return o
}

func (o ManagedServerCreateModeOutput) ToManagedServerCreateModePtrOutput() ManagedServerCreateModePtrOutput {
	return o.ToManagedServerCreateModePtrOutputWithContext(context.Background())
}

func (o ManagedServerCreateModeOutput) ToManagedServerCreateModePtrOutputWithContext(ctx context.Context) ManagedServerCreateModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServerCreateMode) *ManagedServerCreateMode {
		return &v
	}).(ManagedServerCreateModePtrOutput)
}

func (o ManagedServerCreateModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedServerCreateModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedServerCreateMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedServerCreateModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedServerCreateModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedServerCreateMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedServerCreateModePtrOutput struct{ *pulumi.OutputState }

func (ManagedServerCreateModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServerCreateMode)(nil)).Elem()
}

func (o ManagedServerCreateModePtrOutput) ToManagedServerCreateModePtrOutput() ManagedServerCreateModePtrOutput {
	return o
}

func (o ManagedServerCreateModePtrOutput) ToManagedServerCreateModePtrOutputWithContext(ctx context.Context) ManagedServerCreateModePtrOutput {
	return o
}

func (o ManagedServerCreateModePtrOutput) Elem() ManagedServerCreateModeOutput {
	return o.ApplyT(func(v *ManagedServerCreateMode) ManagedServerCreateMode {
		if v != nil {
			return *v
		}
		var ret ManagedServerCreateMode
		return ret
	}).(ManagedServerCreateModeOutput)
}

func (o ManagedServerCreateModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedServerCreateModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedServerCreateMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ManagedServerCreateModeInput interface {
	pulumi.Input

	ToManagedServerCreateModeOutput() ManagedServerCreateModeOutput
	ToManagedServerCreateModeOutputWithContext(context.Context) ManagedServerCreateModeOutput
}

var managedServerCreateModePtrType = reflect.TypeOf((**ManagedServerCreateMode)(nil)).Elem()

type ManagedServerCreateModePtrInput interface {
	pulumi.Input

	ToManagedServerCreateModePtrOutput() ManagedServerCreateModePtrOutput
	ToManagedServerCreateModePtrOutputWithContext(context.Context) ManagedServerCreateModePtrOutput
}

type managedServerCreateModePtr string

func ManagedServerCreateModePtr(v string) ManagedServerCreateModePtrInput {
	return (*managedServerCreateModePtr)(&v)
}

func (*managedServerCreateModePtr) ElementType() reflect.Type {
	return managedServerCreateModePtrType
}

func (in *managedServerCreateModePtr) ToManagedServerCreateModePtrOutput() ManagedServerCreateModePtrOutput {
	return pulumi.ToOutput(in).(ManagedServerCreateModePtrOutput)
}

func (in *managedServerCreateModePtr) ToManagedServerCreateModePtrOutputWithContext(ctx context.Context) ManagedServerCreateModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedServerCreateModePtrOutput)
}

type SecurityAlertPolicyState string

const (
	SecurityAlertPolicyStateNew      = SecurityAlertPolicyState("New")
	SecurityAlertPolicyStateEnabled  = SecurityAlertPolicyState("Enabled")
	SecurityAlertPolicyStateDisabled = SecurityAlertPolicyState("Disabled")
)

func (SecurityAlertPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAlertPolicyState)(nil)).Elem()
}

func (e SecurityAlertPolicyState) ToSecurityAlertPolicyStateOutput() SecurityAlertPolicyStateOutput {
	return pulumi.ToOutput(e).(SecurityAlertPolicyStateOutput)
}

func (e SecurityAlertPolicyState) ToSecurityAlertPolicyStateOutputWithContext(ctx context.Context) SecurityAlertPolicyStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SecurityAlertPolicyStateOutput)
}

func (e SecurityAlertPolicyState) ToSecurityAlertPolicyStatePtrOutput() SecurityAlertPolicyStatePtrOutput {
	return e.ToSecurityAlertPolicyStatePtrOutputWithContext(context.Background())
}

func (e SecurityAlertPolicyState) ToSecurityAlertPolicyStatePtrOutputWithContext(ctx context.Context) SecurityAlertPolicyStatePtrOutput {
	return SecurityAlertPolicyState(e).ToSecurityAlertPolicyStateOutputWithContext(ctx).ToSecurityAlertPolicyStatePtrOutputWithContext(ctx)
}

func (e SecurityAlertPolicyState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityAlertPolicyState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityAlertPolicyState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecurityAlertPolicyState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SecurityAlertPolicyStateOutput struct{ *pulumi.OutputState }

func (SecurityAlertPolicyStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAlertPolicyState)(nil)).Elem()
}

func (o SecurityAlertPolicyStateOutput) ToSecurityAlertPolicyStateOutput() SecurityAlertPolicyStateOutput {
	return o
}

func (o SecurityAlertPolicyStateOutput) ToSecurityAlertPolicyStateOutputWithContext(ctx context.Context) SecurityAlertPolicyStateOutput {
	return o
}

func (o SecurityAlertPolicyStateOutput) ToSecurityAlertPolicyStatePtrOutput() SecurityAlertPolicyStatePtrOutput {
	return o.ToSecurityAlertPolicyStatePtrOutputWithContext(context.Background())
}

func (o SecurityAlertPolicyStateOutput) ToSecurityAlertPolicyStatePtrOutputWithContext(ctx context.Context) SecurityAlertPolicyStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityAlertPolicyState) *SecurityAlertPolicyState {
		return &v
	}).(SecurityAlertPolicyStatePtrOutput)
}

func (o SecurityAlertPolicyStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecurityAlertPolicyStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityAlertPolicyState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecurityAlertPolicyStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityAlertPolicyStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityAlertPolicyState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecurityAlertPolicyStatePtrOutput struct{ *pulumi.OutputState }

func (SecurityAlertPolicyStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAlertPolicyState)(nil)).Elem()
}

func (o SecurityAlertPolicyStatePtrOutput) ToSecurityAlertPolicyStatePtrOutput() SecurityAlertPolicyStatePtrOutput {
	return o
}

func (o SecurityAlertPolicyStatePtrOutput) ToSecurityAlertPolicyStatePtrOutputWithContext(ctx context.Context) SecurityAlertPolicyStatePtrOutput {
	return o
}

func (o SecurityAlertPolicyStatePtrOutput) Elem() SecurityAlertPolicyStateOutput {
	return o.ApplyT(func(v *SecurityAlertPolicyState) SecurityAlertPolicyState {
		if v != nil {
			return *v
		}
		var ret SecurityAlertPolicyState
		return ret
	}).(SecurityAlertPolicyStateOutput)
}

func (o SecurityAlertPolicyStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityAlertPolicyStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecurityAlertPolicyState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SecurityAlertPolicyStateInput interface {
	pulumi.Input

	ToSecurityAlertPolicyStateOutput() SecurityAlertPolicyStateOutput
	ToSecurityAlertPolicyStateOutputWithContext(context.Context) SecurityAlertPolicyStateOutput
}

var securityAlertPolicyStatePtrType = reflect.TypeOf((**SecurityAlertPolicyState)(nil)).Elem()

type SecurityAlertPolicyStatePtrInput interface {
	pulumi.Input

	ToSecurityAlertPolicyStatePtrOutput() SecurityAlertPolicyStatePtrOutput
	ToSecurityAlertPolicyStatePtrOutputWithContext(context.Context) SecurityAlertPolicyStatePtrOutput
}

type securityAlertPolicyStatePtr string

func SecurityAlertPolicyStatePtr(v string) SecurityAlertPolicyStatePtrInput {
	return (*securityAlertPolicyStatePtr)(&v)
}

func (*securityAlertPolicyStatePtr) ElementType() reflect.Type {
	return securityAlertPolicyStatePtrType
}

func (in *securityAlertPolicyStatePtr) ToSecurityAlertPolicyStatePtrOutput() SecurityAlertPolicyStatePtrOutput {
	return pulumi.ToOutput(in).(SecurityAlertPolicyStatePtrOutput)
}

func (in *securityAlertPolicyStatePtr) ToSecurityAlertPolicyStatePtrOutputWithContext(ctx context.Context) SecurityAlertPolicyStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SecurityAlertPolicyStatePtrOutput)
}

type SensitivityLabelRank string

const (
	SensitivityLabelRankNone     = SensitivityLabelRank("None")
	SensitivityLabelRankLow      = SensitivityLabelRank("Low")
	SensitivityLabelRankMedium   = SensitivityLabelRank("Medium")
	SensitivityLabelRankHigh     = SensitivityLabelRank("High")
	SensitivityLabelRankCritical = SensitivityLabelRank("Critical")
)

func (SensitivityLabelRank) ElementType() reflect.Type {
	return reflect.TypeOf((*SensitivityLabelRank)(nil)).Elem()
}

func (e SensitivityLabelRank) ToSensitivityLabelRankOutput() SensitivityLabelRankOutput {
	return pulumi.ToOutput(e).(SensitivityLabelRankOutput)
}

func (e SensitivityLabelRank) ToSensitivityLabelRankOutputWithContext(ctx context.Context) SensitivityLabelRankOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SensitivityLabelRankOutput)
}

func (e SensitivityLabelRank) ToSensitivityLabelRankPtrOutput() SensitivityLabelRankPtrOutput {
	return e.ToSensitivityLabelRankPtrOutputWithContext(context.Background())
}

func (e SensitivityLabelRank) ToSensitivityLabelRankPtrOutputWithContext(ctx context.Context) SensitivityLabelRankPtrOutput {
	return SensitivityLabelRank(e).ToSensitivityLabelRankOutputWithContext(ctx).ToSensitivityLabelRankPtrOutputWithContext(ctx)
}

func (e SensitivityLabelRank) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SensitivityLabelRank) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SensitivityLabelRank) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SensitivityLabelRank) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SensitivityLabelRankOutput struct{ *pulumi.OutputState }

func (SensitivityLabelRankOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SensitivityLabelRank)(nil)).Elem()
}

func (o SensitivityLabelRankOutput) ToSensitivityLabelRankOutput() SensitivityLabelRankOutput {
	return o
}

func (o SensitivityLabelRankOutput) ToSensitivityLabelRankOutputWithContext(ctx context.Context) SensitivityLabelRankOutput {
	return o
}

func (o SensitivityLabelRankOutput) ToSensitivityLabelRankPtrOutput() SensitivityLabelRankPtrOutput {
	return o.ToSensitivityLabelRankPtrOutputWithContext(context.Background())
}

func (o SensitivityLabelRankOutput) ToSensitivityLabelRankPtrOutputWithContext(ctx context.Context) SensitivityLabelRankPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SensitivityLabelRank) *SensitivityLabelRank {
		return &v
	}).(SensitivityLabelRankPtrOutput)
}

func (o SensitivityLabelRankOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SensitivityLabelRankOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SensitivityLabelRank) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SensitivityLabelRankOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SensitivityLabelRankOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SensitivityLabelRank) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SensitivityLabelRankPtrOutput struct{ *pulumi.OutputState }

func (SensitivityLabelRankPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SensitivityLabelRank)(nil)).Elem()
}

func (o SensitivityLabelRankPtrOutput) ToSensitivityLabelRankPtrOutput() SensitivityLabelRankPtrOutput {
	return o
}

func (o SensitivityLabelRankPtrOutput) ToSensitivityLabelRankPtrOutputWithContext(ctx context.Context) SensitivityLabelRankPtrOutput {
	return o
}

func (o SensitivityLabelRankPtrOutput) Elem() SensitivityLabelRankOutput {
	return o.ApplyT(func(v *SensitivityLabelRank) SensitivityLabelRank {
		if v != nil {
			return *v
		}
		var ret SensitivityLabelRank
		return ret
	}).(SensitivityLabelRankOutput)
}

func (o SensitivityLabelRankPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SensitivityLabelRankPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SensitivityLabelRank) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SensitivityLabelRankInput interface {
	pulumi.Input

	ToSensitivityLabelRankOutput() SensitivityLabelRankOutput
	ToSensitivityLabelRankOutputWithContext(context.Context) SensitivityLabelRankOutput
}

var sensitivityLabelRankPtrType = reflect.TypeOf((**SensitivityLabelRank)(nil)).Elem()

type SensitivityLabelRankPtrInput interface {
	pulumi.Input

	ToSensitivityLabelRankPtrOutput() SensitivityLabelRankPtrOutput
	ToSensitivityLabelRankPtrOutputWithContext(context.Context) SensitivityLabelRankPtrOutput
}

type sensitivityLabelRankPtr string

func SensitivityLabelRankPtr(v string) SensitivityLabelRankPtrInput {
	return (*sensitivityLabelRankPtr)(&v)
}

func (*sensitivityLabelRankPtr) ElementType() reflect.Type {
	return sensitivityLabelRankPtrType
}

func (in *sensitivityLabelRankPtr) ToSensitivityLabelRankPtrOutput() SensitivityLabelRankPtrOutput {
	return pulumi.ToOutput(in).(SensitivityLabelRankPtrOutput)
}

func (in *sensitivityLabelRankPtr) ToSensitivityLabelRankPtrOutputWithContext(ctx context.Context) SensitivityLabelRankPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SensitivityLabelRankPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AdministratorTypeOutput{})
	pulumi.RegisterOutputType(AdministratorTypePtrOutput{})
	pulumi.RegisterOutputType(CatalogCollationTypeOutput{})
	pulumi.RegisterOutputType(CatalogCollationTypePtrOutput{})
	pulumi.RegisterOutputType(IdentityTypeOutput{})
	pulumi.RegisterOutputType(IdentityTypePtrOutput{})
	pulumi.RegisterOutputType(InstancePoolLicenseTypeOutput{})
	pulumi.RegisterOutputType(InstancePoolLicenseTypePtrOutput{})
	pulumi.RegisterOutputType(ManagedDatabaseCreateModeOutput{})
	pulumi.RegisterOutputType(ManagedDatabaseCreateModePtrOutput{})
	pulumi.RegisterOutputType(ManagedInstanceLicenseTypeOutput{})
	pulumi.RegisterOutputType(ManagedInstanceLicenseTypePtrOutput{})
	pulumi.RegisterOutputType(ManagedInstanceProxyOverrideOutput{})
	pulumi.RegisterOutputType(ManagedInstanceProxyOverridePtrOutput{})
	pulumi.RegisterOutputType(ManagedServerCreateModeOutput{})
	pulumi.RegisterOutputType(ManagedServerCreateModePtrOutput{})
	pulumi.RegisterOutputType(SecurityAlertPolicyStateOutput{})
	pulumi.RegisterOutputType(SecurityAlertPolicyStatePtrOutput{})
	pulumi.RegisterOutputType(SensitivityLabelRankOutput{})
	pulumi.RegisterOutputType(SensitivityLabelRankPtrOutput{})
}
