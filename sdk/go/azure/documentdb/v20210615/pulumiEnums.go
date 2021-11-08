


package v20210615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AnalyticalStorageSchemaType string

const (
	AnalyticalStorageSchemaTypeWellDefined  = AnalyticalStorageSchemaType("WellDefined")
	AnalyticalStorageSchemaTypeFullFidelity = AnalyticalStorageSchemaType("FullFidelity")
)

func (AnalyticalStorageSchemaType) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticalStorageSchemaType)(nil)).Elem()
}

func (e AnalyticalStorageSchemaType) ToAnalyticalStorageSchemaTypeOutput() AnalyticalStorageSchemaTypeOutput {
	return pulumi.ToOutput(e).(AnalyticalStorageSchemaTypeOutput)
}

func (e AnalyticalStorageSchemaType) ToAnalyticalStorageSchemaTypeOutputWithContext(ctx context.Context) AnalyticalStorageSchemaTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AnalyticalStorageSchemaTypeOutput)
}

func (e AnalyticalStorageSchemaType) ToAnalyticalStorageSchemaTypePtrOutput() AnalyticalStorageSchemaTypePtrOutput {
	return e.ToAnalyticalStorageSchemaTypePtrOutputWithContext(context.Background())
}

func (e AnalyticalStorageSchemaType) ToAnalyticalStorageSchemaTypePtrOutputWithContext(ctx context.Context) AnalyticalStorageSchemaTypePtrOutput {
	return AnalyticalStorageSchemaType(e).ToAnalyticalStorageSchemaTypeOutputWithContext(ctx).ToAnalyticalStorageSchemaTypePtrOutputWithContext(ctx)
}

func (e AnalyticalStorageSchemaType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AnalyticalStorageSchemaType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AnalyticalStorageSchemaType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AnalyticalStorageSchemaType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AnalyticalStorageSchemaTypeOutput struct{ *pulumi.OutputState }

func (AnalyticalStorageSchemaTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticalStorageSchemaType)(nil)).Elem()
}

func (o AnalyticalStorageSchemaTypeOutput) ToAnalyticalStorageSchemaTypeOutput() AnalyticalStorageSchemaTypeOutput {
	return o
}

func (o AnalyticalStorageSchemaTypeOutput) ToAnalyticalStorageSchemaTypeOutputWithContext(ctx context.Context) AnalyticalStorageSchemaTypeOutput {
	return o
}

func (o AnalyticalStorageSchemaTypeOutput) ToAnalyticalStorageSchemaTypePtrOutput() AnalyticalStorageSchemaTypePtrOutput {
	return o.ToAnalyticalStorageSchemaTypePtrOutputWithContext(context.Background())
}

func (o AnalyticalStorageSchemaTypeOutput) ToAnalyticalStorageSchemaTypePtrOutputWithContext(ctx context.Context) AnalyticalStorageSchemaTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AnalyticalStorageSchemaType) *AnalyticalStorageSchemaType {
		return &v
	}).(AnalyticalStorageSchemaTypePtrOutput)
}

func (o AnalyticalStorageSchemaTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AnalyticalStorageSchemaTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AnalyticalStorageSchemaType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AnalyticalStorageSchemaTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AnalyticalStorageSchemaTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AnalyticalStorageSchemaType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AnalyticalStorageSchemaTypePtrOutput struct{ *pulumi.OutputState }

func (AnalyticalStorageSchemaTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticalStorageSchemaType)(nil)).Elem()
}

func (o AnalyticalStorageSchemaTypePtrOutput) ToAnalyticalStorageSchemaTypePtrOutput() AnalyticalStorageSchemaTypePtrOutput {
	return o
}

func (o AnalyticalStorageSchemaTypePtrOutput) ToAnalyticalStorageSchemaTypePtrOutputWithContext(ctx context.Context) AnalyticalStorageSchemaTypePtrOutput {
	return o
}

func (o AnalyticalStorageSchemaTypePtrOutput) Elem() AnalyticalStorageSchemaTypeOutput {
	return o.ApplyT(func(v *AnalyticalStorageSchemaType) AnalyticalStorageSchemaType {
		if v != nil {
			return *v
		}
		var ret AnalyticalStorageSchemaType
		return ret
	}).(AnalyticalStorageSchemaTypeOutput)
}

func (o AnalyticalStorageSchemaTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AnalyticalStorageSchemaTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AnalyticalStorageSchemaType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AnalyticalStorageSchemaTypeInput interface {
	pulumi.Input

	ToAnalyticalStorageSchemaTypeOutput() AnalyticalStorageSchemaTypeOutput
	ToAnalyticalStorageSchemaTypeOutputWithContext(context.Context) AnalyticalStorageSchemaTypeOutput
}

var analyticalStorageSchemaTypePtrType = reflect.TypeOf((**AnalyticalStorageSchemaType)(nil)).Elem()

type AnalyticalStorageSchemaTypePtrInput interface {
	pulumi.Input

	ToAnalyticalStorageSchemaTypePtrOutput() AnalyticalStorageSchemaTypePtrOutput
	ToAnalyticalStorageSchemaTypePtrOutputWithContext(context.Context) AnalyticalStorageSchemaTypePtrOutput
}

type analyticalStorageSchemaTypePtr string

func AnalyticalStorageSchemaTypePtr(v string) AnalyticalStorageSchemaTypePtrInput {
	return (*analyticalStorageSchemaTypePtr)(&v)
}

func (*analyticalStorageSchemaTypePtr) ElementType() reflect.Type {
	return analyticalStorageSchemaTypePtrType
}

func (in *analyticalStorageSchemaTypePtr) ToAnalyticalStorageSchemaTypePtrOutput() AnalyticalStorageSchemaTypePtrOutput {
	return pulumi.ToOutput(in).(AnalyticalStorageSchemaTypePtrOutput)
}

func (in *analyticalStorageSchemaTypePtr) ToAnalyticalStorageSchemaTypePtrOutputWithContext(ctx context.Context) AnalyticalStorageSchemaTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AnalyticalStorageSchemaTypePtrOutput)
}

type BackupPolicyMigrationStatus string

const (
	BackupPolicyMigrationStatusInvalid    = BackupPolicyMigrationStatus("Invalid")
	BackupPolicyMigrationStatusInProgress = BackupPolicyMigrationStatus("InProgress")
	BackupPolicyMigrationStatusCompleted  = BackupPolicyMigrationStatus("Completed")
	BackupPolicyMigrationStatusFailed     = BackupPolicyMigrationStatus("Failed")
)

func (BackupPolicyMigrationStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyMigrationStatus)(nil)).Elem()
}

func (e BackupPolicyMigrationStatus) ToBackupPolicyMigrationStatusOutput() BackupPolicyMigrationStatusOutput {
	return pulumi.ToOutput(e).(BackupPolicyMigrationStatusOutput)
}

func (e BackupPolicyMigrationStatus) ToBackupPolicyMigrationStatusOutputWithContext(ctx context.Context) BackupPolicyMigrationStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BackupPolicyMigrationStatusOutput)
}

func (e BackupPolicyMigrationStatus) ToBackupPolicyMigrationStatusPtrOutput() BackupPolicyMigrationStatusPtrOutput {
	return e.ToBackupPolicyMigrationStatusPtrOutputWithContext(context.Background())
}

func (e BackupPolicyMigrationStatus) ToBackupPolicyMigrationStatusPtrOutputWithContext(ctx context.Context) BackupPolicyMigrationStatusPtrOutput {
	return BackupPolicyMigrationStatus(e).ToBackupPolicyMigrationStatusOutputWithContext(ctx).ToBackupPolicyMigrationStatusPtrOutputWithContext(ctx)
}

func (e BackupPolicyMigrationStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackupPolicyMigrationStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackupPolicyMigrationStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BackupPolicyMigrationStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BackupPolicyMigrationStatusOutput struct{ *pulumi.OutputState }

func (BackupPolicyMigrationStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyMigrationStatus)(nil)).Elem()
}

func (o BackupPolicyMigrationStatusOutput) ToBackupPolicyMigrationStatusOutput() BackupPolicyMigrationStatusOutput {
	return o
}

func (o BackupPolicyMigrationStatusOutput) ToBackupPolicyMigrationStatusOutputWithContext(ctx context.Context) BackupPolicyMigrationStatusOutput {
	return o
}

func (o BackupPolicyMigrationStatusOutput) ToBackupPolicyMigrationStatusPtrOutput() BackupPolicyMigrationStatusPtrOutput {
	return o.ToBackupPolicyMigrationStatusPtrOutputWithContext(context.Background())
}

func (o BackupPolicyMigrationStatusOutput) ToBackupPolicyMigrationStatusPtrOutputWithContext(ctx context.Context) BackupPolicyMigrationStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupPolicyMigrationStatus) *BackupPolicyMigrationStatus {
		return &v
	}).(BackupPolicyMigrationStatusPtrOutput)
}

func (o BackupPolicyMigrationStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BackupPolicyMigrationStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BackupPolicyMigrationStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BackupPolicyMigrationStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BackupPolicyMigrationStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BackupPolicyMigrationStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BackupPolicyMigrationStatusPtrOutput struct{ *pulumi.OutputState }

func (BackupPolicyMigrationStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicyMigrationStatus)(nil)).Elem()
}

func (o BackupPolicyMigrationStatusPtrOutput) ToBackupPolicyMigrationStatusPtrOutput() BackupPolicyMigrationStatusPtrOutput {
	return o
}

func (o BackupPolicyMigrationStatusPtrOutput) ToBackupPolicyMigrationStatusPtrOutputWithContext(ctx context.Context) BackupPolicyMigrationStatusPtrOutput {
	return o
}

func (o BackupPolicyMigrationStatusPtrOutput) Elem() BackupPolicyMigrationStatusOutput {
	return o.ApplyT(func(v *BackupPolicyMigrationStatus) BackupPolicyMigrationStatus {
		if v != nil {
			return *v
		}
		var ret BackupPolicyMigrationStatus
		return ret
	}).(BackupPolicyMigrationStatusOutput)
}

func (o BackupPolicyMigrationStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BackupPolicyMigrationStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BackupPolicyMigrationStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BackupPolicyMigrationStatusInput interface {
	pulumi.Input

	ToBackupPolicyMigrationStatusOutput() BackupPolicyMigrationStatusOutput
	ToBackupPolicyMigrationStatusOutputWithContext(context.Context) BackupPolicyMigrationStatusOutput
}

var backupPolicyMigrationStatusPtrType = reflect.TypeOf((**BackupPolicyMigrationStatus)(nil)).Elem()

type BackupPolicyMigrationStatusPtrInput interface {
	pulumi.Input

	ToBackupPolicyMigrationStatusPtrOutput() BackupPolicyMigrationStatusPtrOutput
	ToBackupPolicyMigrationStatusPtrOutputWithContext(context.Context) BackupPolicyMigrationStatusPtrOutput
}

type backupPolicyMigrationStatusPtr string

func BackupPolicyMigrationStatusPtr(v string) BackupPolicyMigrationStatusPtrInput {
	return (*backupPolicyMigrationStatusPtr)(&v)
}

func (*backupPolicyMigrationStatusPtr) ElementType() reflect.Type {
	return backupPolicyMigrationStatusPtrType
}

func (in *backupPolicyMigrationStatusPtr) ToBackupPolicyMigrationStatusPtrOutput() BackupPolicyMigrationStatusPtrOutput {
	return pulumi.ToOutput(in).(BackupPolicyMigrationStatusPtrOutput)
}

func (in *backupPolicyMigrationStatusPtr) ToBackupPolicyMigrationStatusPtrOutputWithContext(ctx context.Context) BackupPolicyMigrationStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BackupPolicyMigrationStatusPtrOutput)
}

type BackupPolicyType string

const (
	BackupPolicyTypePeriodic   = BackupPolicyType("Periodic")
	BackupPolicyTypeContinuous = BackupPolicyType("Continuous")
)

func (BackupPolicyType) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyType)(nil)).Elem()
}

func (e BackupPolicyType) ToBackupPolicyTypeOutput() BackupPolicyTypeOutput {
	return pulumi.ToOutput(e).(BackupPolicyTypeOutput)
}

func (e BackupPolicyType) ToBackupPolicyTypeOutputWithContext(ctx context.Context) BackupPolicyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BackupPolicyTypeOutput)
}

func (e BackupPolicyType) ToBackupPolicyTypePtrOutput() BackupPolicyTypePtrOutput {
	return e.ToBackupPolicyTypePtrOutputWithContext(context.Background())
}

func (e BackupPolicyType) ToBackupPolicyTypePtrOutputWithContext(ctx context.Context) BackupPolicyTypePtrOutput {
	return BackupPolicyType(e).ToBackupPolicyTypeOutputWithContext(ctx).ToBackupPolicyTypePtrOutputWithContext(ctx)
}

func (e BackupPolicyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackupPolicyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackupPolicyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BackupPolicyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BackupPolicyTypeOutput struct{ *pulumi.OutputState }

func (BackupPolicyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyType)(nil)).Elem()
}

func (o BackupPolicyTypeOutput) ToBackupPolicyTypeOutput() BackupPolicyTypeOutput {
	return o
}

func (o BackupPolicyTypeOutput) ToBackupPolicyTypeOutputWithContext(ctx context.Context) BackupPolicyTypeOutput {
	return o
}

func (o BackupPolicyTypeOutput) ToBackupPolicyTypePtrOutput() BackupPolicyTypePtrOutput {
	return o.ToBackupPolicyTypePtrOutputWithContext(context.Background())
}

func (o BackupPolicyTypeOutput) ToBackupPolicyTypePtrOutputWithContext(ctx context.Context) BackupPolicyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupPolicyType) *BackupPolicyType {
		return &v
	}).(BackupPolicyTypePtrOutput)
}

func (o BackupPolicyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BackupPolicyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BackupPolicyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BackupPolicyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BackupPolicyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BackupPolicyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BackupPolicyTypePtrOutput struct{ *pulumi.OutputState }

func (BackupPolicyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicyType)(nil)).Elem()
}

func (o BackupPolicyTypePtrOutput) ToBackupPolicyTypePtrOutput() BackupPolicyTypePtrOutput {
	return o
}

func (o BackupPolicyTypePtrOutput) ToBackupPolicyTypePtrOutputWithContext(ctx context.Context) BackupPolicyTypePtrOutput {
	return o
}

func (o BackupPolicyTypePtrOutput) Elem() BackupPolicyTypeOutput {
	return o.ApplyT(func(v *BackupPolicyType) BackupPolicyType {
		if v != nil {
			return *v
		}
		var ret BackupPolicyType
		return ret
	}).(BackupPolicyTypeOutput)
}

func (o BackupPolicyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BackupPolicyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BackupPolicyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BackupPolicyTypeInput interface {
	pulumi.Input

	ToBackupPolicyTypeOutput() BackupPolicyTypeOutput
	ToBackupPolicyTypeOutputWithContext(context.Context) BackupPolicyTypeOutput
}

var backupPolicyTypePtrType = reflect.TypeOf((**BackupPolicyType)(nil)).Elem()

type BackupPolicyTypePtrInput interface {
	pulumi.Input

	ToBackupPolicyTypePtrOutput() BackupPolicyTypePtrOutput
	ToBackupPolicyTypePtrOutputWithContext(context.Context) BackupPolicyTypePtrOutput
}

type backupPolicyTypePtr string

func BackupPolicyTypePtr(v string) BackupPolicyTypePtrInput {
	return (*backupPolicyTypePtr)(&v)
}

func (*backupPolicyTypePtr) ElementType() reflect.Type {
	return backupPolicyTypePtrType
}

func (in *backupPolicyTypePtr) ToBackupPolicyTypePtrOutput() BackupPolicyTypePtrOutput {
	return pulumi.ToOutput(in).(BackupPolicyTypePtrOutput)
}

func (in *backupPolicyTypePtr) ToBackupPolicyTypePtrOutputWithContext(ctx context.Context) BackupPolicyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BackupPolicyTypePtrOutput)
}

type CompositePathSortOrder string

const (
	CompositePathSortOrderAscending  = CompositePathSortOrder("ascending")
	CompositePathSortOrderDescending = CompositePathSortOrder("descending")
)

func (CompositePathSortOrder) ElementType() reflect.Type {
	return reflect.TypeOf((*CompositePathSortOrder)(nil)).Elem()
}

func (e CompositePathSortOrder) ToCompositePathSortOrderOutput() CompositePathSortOrderOutput {
	return pulumi.ToOutput(e).(CompositePathSortOrderOutput)
}

func (e CompositePathSortOrder) ToCompositePathSortOrderOutputWithContext(ctx context.Context) CompositePathSortOrderOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CompositePathSortOrderOutput)
}

func (e CompositePathSortOrder) ToCompositePathSortOrderPtrOutput() CompositePathSortOrderPtrOutput {
	return e.ToCompositePathSortOrderPtrOutputWithContext(context.Background())
}

func (e CompositePathSortOrder) ToCompositePathSortOrderPtrOutputWithContext(ctx context.Context) CompositePathSortOrderPtrOutput {
	return CompositePathSortOrder(e).ToCompositePathSortOrderOutputWithContext(ctx).ToCompositePathSortOrderPtrOutputWithContext(ctx)
}

func (e CompositePathSortOrder) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CompositePathSortOrder) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CompositePathSortOrder) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CompositePathSortOrder) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CompositePathSortOrderOutput struct{ *pulumi.OutputState }

func (CompositePathSortOrderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CompositePathSortOrder)(nil)).Elem()
}

func (o CompositePathSortOrderOutput) ToCompositePathSortOrderOutput() CompositePathSortOrderOutput {
	return o
}

func (o CompositePathSortOrderOutput) ToCompositePathSortOrderOutputWithContext(ctx context.Context) CompositePathSortOrderOutput {
	return o
}

func (o CompositePathSortOrderOutput) ToCompositePathSortOrderPtrOutput() CompositePathSortOrderPtrOutput {
	return o.ToCompositePathSortOrderPtrOutputWithContext(context.Background())
}

func (o CompositePathSortOrderOutput) ToCompositePathSortOrderPtrOutputWithContext(ctx context.Context) CompositePathSortOrderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CompositePathSortOrder) *CompositePathSortOrder {
		return &v
	}).(CompositePathSortOrderPtrOutput)
}

func (o CompositePathSortOrderOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CompositePathSortOrderOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CompositePathSortOrder) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CompositePathSortOrderOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CompositePathSortOrderOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CompositePathSortOrder) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CompositePathSortOrderPtrOutput struct{ *pulumi.OutputState }

func (CompositePathSortOrderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CompositePathSortOrder)(nil)).Elem()
}

func (o CompositePathSortOrderPtrOutput) ToCompositePathSortOrderPtrOutput() CompositePathSortOrderPtrOutput {
	return o
}

func (o CompositePathSortOrderPtrOutput) ToCompositePathSortOrderPtrOutputWithContext(ctx context.Context) CompositePathSortOrderPtrOutput {
	return o
}

func (o CompositePathSortOrderPtrOutput) Elem() CompositePathSortOrderOutput {
	return o.ApplyT(func(v *CompositePathSortOrder) CompositePathSortOrder {
		if v != nil {
			return *v
		}
		var ret CompositePathSortOrder
		return ret
	}).(CompositePathSortOrderOutput)
}

func (o CompositePathSortOrderPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CompositePathSortOrderPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CompositePathSortOrder) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CompositePathSortOrderInput interface {
	pulumi.Input

	ToCompositePathSortOrderOutput() CompositePathSortOrderOutput
	ToCompositePathSortOrderOutputWithContext(context.Context) CompositePathSortOrderOutput
}

var compositePathSortOrderPtrType = reflect.TypeOf((**CompositePathSortOrder)(nil)).Elem()

type CompositePathSortOrderPtrInput interface {
	pulumi.Input

	ToCompositePathSortOrderPtrOutput() CompositePathSortOrderPtrOutput
	ToCompositePathSortOrderPtrOutputWithContext(context.Context) CompositePathSortOrderPtrOutput
}

type compositePathSortOrderPtr string

func CompositePathSortOrderPtr(v string) CompositePathSortOrderPtrInput {
	return (*compositePathSortOrderPtr)(&v)
}

func (*compositePathSortOrderPtr) ElementType() reflect.Type {
	return compositePathSortOrderPtrType
}

func (in *compositePathSortOrderPtr) ToCompositePathSortOrderPtrOutput() CompositePathSortOrderPtrOutput {
	return pulumi.ToOutput(in).(CompositePathSortOrderPtrOutput)
}

func (in *compositePathSortOrderPtr) ToCompositePathSortOrderPtrOutputWithContext(ctx context.Context) CompositePathSortOrderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CompositePathSortOrderPtrOutput)
}

type ConflictResolutionMode string

const (
	ConflictResolutionModeLastWriterWins = ConflictResolutionMode("LastWriterWins")
	ConflictResolutionModeCustom         = ConflictResolutionMode("Custom")
)

func (ConflictResolutionMode) ElementType() reflect.Type {
	return reflect.TypeOf((*ConflictResolutionMode)(nil)).Elem()
}

func (e ConflictResolutionMode) ToConflictResolutionModeOutput() ConflictResolutionModeOutput {
	return pulumi.ToOutput(e).(ConflictResolutionModeOutput)
}

func (e ConflictResolutionMode) ToConflictResolutionModeOutputWithContext(ctx context.Context) ConflictResolutionModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConflictResolutionModeOutput)
}

func (e ConflictResolutionMode) ToConflictResolutionModePtrOutput() ConflictResolutionModePtrOutput {
	return e.ToConflictResolutionModePtrOutputWithContext(context.Background())
}

func (e ConflictResolutionMode) ToConflictResolutionModePtrOutputWithContext(ctx context.Context) ConflictResolutionModePtrOutput {
	return ConflictResolutionMode(e).ToConflictResolutionModeOutputWithContext(ctx).ToConflictResolutionModePtrOutputWithContext(ctx)
}

func (e ConflictResolutionMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConflictResolutionMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConflictResolutionMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConflictResolutionMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConflictResolutionModeOutput struct{ *pulumi.OutputState }

func (ConflictResolutionModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConflictResolutionMode)(nil)).Elem()
}

func (o ConflictResolutionModeOutput) ToConflictResolutionModeOutput() ConflictResolutionModeOutput {
	return o
}

func (o ConflictResolutionModeOutput) ToConflictResolutionModeOutputWithContext(ctx context.Context) ConflictResolutionModeOutput {
	return o
}

func (o ConflictResolutionModeOutput) ToConflictResolutionModePtrOutput() ConflictResolutionModePtrOutput {
	return o.ToConflictResolutionModePtrOutputWithContext(context.Background())
}

func (o ConflictResolutionModeOutput) ToConflictResolutionModePtrOutputWithContext(ctx context.Context) ConflictResolutionModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConflictResolutionMode) *ConflictResolutionMode {
		return &v
	}).(ConflictResolutionModePtrOutput)
}

func (o ConflictResolutionModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConflictResolutionModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConflictResolutionMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConflictResolutionModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConflictResolutionModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConflictResolutionMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConflictResolutionModePtrOutput struct{ *pulumi.OutputState }

func (ConflictResolutionModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConflictResolutionMode)(nil)).Elem()
}

func (o ConflictResolutionModePtrOutput) ToConflictResolutionModePtrOutput() ConflictResolutionModePtrOutput {
	return o
}

func (o ConflictResolutionModePtrOutput) ToConflictResolutionModePtrOutputWithContext(ctx context.Context) ConflictResolutionModePtrOutput {
	return o
}

func (o ConflictResolutionModePtrOutput) Elem() ConflictResolutionModeOutput {
	return o.ApplyT(func(v *ConflictResolutionMode) ConflictResolutionMode {
		if v != nil {
			return *v
		}
		var ret ConflictResolutionMode
		return ret
	}).(ConflictResolutionModeOutput)
}

func (o ConflictResolutionModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConflictResolutionModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConflictResolutionMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConflictResolutionModeInput interface {
	pulumi.Input

	ToConflictResolutionModeOutput() ConflictResolutionModeOutput
	ToConflictResolutionModeOutputWithContext(context.Context) ConflictResolutionModeOutput
}

var conflictResolutionModePtrType = reflect.TypeOf((**ConflictResolutionMode)(nil)).Elem()

type ConflictResolutionModePtrInput interface {
	pulumi.Input

	ToConflictResolutionModePtrOutput() ConflictResolutionModePtrOutput
	ToConflictResolutionModePtrOutputWithContext(context.Context) ConflictResolutionModePtrOutput
}

type conflictResolutionModePtr string

func ConflictResolutionModePtr(v string) ConflictResolutionModePtrInput {
	return (*conflictResolutionModePtr)(&v)
}

func (*conflictResolutionModePtr) ElementType() reflect.Type {
	return conflictResolutionModePtrType
}

func (in *conflictResolutionModePtr) ToConflictResolutionModePtrOutput() ConflictResolutionModePtrOutput {
	return pulumi.ToOutput(in).(ConflictResolutionModePtrOutput)
}

func (in *conflictResolutionModePtr) ToConflictResolutionModePtrOutputWithContext(ctx context.Context) ConflictResolutionModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConflictResolutionModePtrOutput)
}

type ConnectorOffer string

const (
	ConnectorOfferSmall = ConnectorOffer("Small")
)

func (ConnectorOffer) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorOffer)(nil)).Elem()
}

func (e ConnectorOffer) ToConnectorOfferOutput() ConnectorOfferOutput {
	return pulumi.ToOutput(e).(ConnectorOfferOutput)
}

func (e ConnectorOffer) ToConnectorOfferOutputWithContext(ctx context.Context) ConnectorOfferOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConnectorOfferOutput)
}

func (e ConnectorOffer) ToConnectorOfferPtrOutput() ConnectorOfferPtrOutput {
	return e.ToConnectorOfferPtrOutputWithContext(context.Background())
}

func (e ConnectorOffer) ToConnectorOfferPtrOutputWithContext(ctx context.Context) ConnectorOfferPtrOutput {
	return ConnectorOffer(e).ToConnectorOfferOutputWithContext(ctx).ToConnectorOfferPtrOutputWithContext(ctx)
}

func (e ConnectorOffer) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectorOffer) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectorOffer) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConnectorOffer) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConnectorOfferOutput struct{ *pulumi.OutputState }

func (ConnectorOfferOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorOffer)(nil)).Elem()
}

func (o ConnectorOfferOutput) ToConnectorOfferOutput() ConnectorOfferOutput {
	return o
}

func (o ConnectorOfferOutput) ToConnectorOfferOutputWithContext(ctx context.Context) ConnectorOfferOutput {
	return o
}

func (o ConnectorOfferOutput) ToConnectorOfferPtrOutput() ConnectorOfferPtrOutput {
	return o.ToConnectorOfferPtrOutputWithContext(context.Background())
}

func (o ConnectorOfferOutput) ToConnectorOfferPtrOutputWithContext(ctx context.Context) ConnectorOfferPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectorOffer) *ConnectorOffer {
		return &v
	}).(ConnectorOfferPtrOutput)
}

func (o ConnectorOfferOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConnectorOfferOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectorOffer) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConnectorOfferOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectorOfferOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectorOffer) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConnectorOfferPtrOutput struct{ *pulumi.OutputState }

func (ConnectorOfferPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorOffer)(nil)).Elem()
}

func (o ConnectorOfferPtrOutput) ToConnectorOfferPtrOutput() ConnectorOfferPtrOutput {
	return o
}

func (o ConnectorOfferPtrOutput) ToConnectorOfferPtrOutputWithContext(ctx context.Context) ConnectorOfferPtrOutput {
	return o
}

func (o ConnectorOfferPtrOutput) Elem() ConnectorOfferOutput {
	return o.ApplyT(func(v *ConnectorOffer) ConnectorOffer {
		if v != nil {
			return *v
		}
		var ret ConnectorOffer
		return ret
	}).(ConnectorOfferOutput)
}

func (o ConnectorOfferPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectorOfferPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConnectorOffer) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConnectorOfferInput interface {
	pulumi.Input

	ToConnectorOfferOutput() ConnectorOfferOutput
	ToConnectorOfferOutputWithContext(context.Context) ConnectorOfferOutput
}

var connectorOfferPtrType = reflect.TypeOf((**ConnectorOffer)(nil)).Elem()

type ConnectorOfferPtrInput interface {
	pulumi.Input

	ToConnectorOfferPtrOutput() ConnectorOfferPtrOutput
	ToConnectorOfferPtrOutputWithContext(context.Context) ConnectorOfferPtrOutput
}

type connectorOfferPtr string

func ConnectorOfferPtr(v string) ConnectorOfferPtrInput {
	return (*connectorOfferPtr)(&v)
}

func (*connectorOfferPtr) ElementType() reflect.Type {
	return connectorOfferPtrType
}

func (in *connectorOfferPtr) ToConnectorOfferPtrOutput() ConnectorOfferPtrOutput {
	return pulumi.ToOutput(in).(ConnectorOfferPtrOutput)
}

func (in *connectorOfferPtr) ToConnectorOfferPtrOutputWithContext(ctx context.Context) ConnectorOfferPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConnectorOfferPtrOutput)
}

type CreateMode string

const (
	CreateModeDefault = CreateMode("Default")
	CreateModeRestore = CreateMode("Restore")
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

type DataType string

const (
	DataTypeString       = DataType("String")
	DataTypeNumber       = DataType("Number")
	DataTypePoint        = DataType("Point")
	DataTypePolygon      = DataType("Polygon")
	DataTypeLineString   = DataType("LineString")
	DataTypeMultiPolygon = DataType("MultiPolygon")
)

func (DataType) ElementType() reflect.Type {
	return reflect.TypeOf((*DataType)(nil)).Elem()
}

func (e DataType) ToDataTypeOutput() DataTypeOutput {
	return pulumi.ToOutput(e).(DataTypeOutput)
}

func (e DataType) ToDataTypeOutputWithContext(ctx context.Context) DataTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataTypeOutput)
}

func (e DataType) ToDataTypePtrOutput() DataTypePtrOutput {
	return e.ToDataTypePtrOutputWithContext(context.Background())
}

func (e DataType) ToDataTypePtrOutputWithContext(ctx context.Context) DataTypePtrOutput {
	return DataType(e).ToDataTypeOutputWithContext(ctx).ToDataTypePtrOutputWithContext(ctx)
}

func (e DataType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataTypeOutput struct{ *pulumi.OutputState }

func (DataTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataType)(nil)).Elem()
}

func (o DataTypeOutput) ToDataTypeOutput() DataTypeOutput {
	return o
}

func (o DataTypeOutput) ToDataTypeOutputWithContext(ctx context.Context) DataTypeOutput {
	return o
}

func (o DataTypeOutput) ToDataTypePtrOutput() DataTypePtrOutput {
	return o.ToDataTypePtrOutputWithContext(context.Background())
}

func (o DataTypeOutput) ToDataTypePtrOutputWithContext(ctx context.Context) DataTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataType) *DataType {
		return &v
	}).(DataTypePtrOutput)
}

func (o DataTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataTypePtrOutput struct{ *pulumi.OutputState }

func (DataTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataType)(nil)).Elem()
}

func (o DataTypePtrOutput) ToDataTypePtrOutput() DataTypePtrOutput {
	return o
}

func (o DataTypePtrOutput) ToDataTypePtrOutputWithContext(ctx context.Context) DataTypePtrOutput {
	return o
}

func (o DataTypePtrOutput) Elem() DataTypeOutput {
	return o.ApplyT(func(v *DataType) DataType {
		if v != nil {
			return *v
		}
		var ret DataType
		return ret
	}).(DataTypeOutput)
}

func (o DataTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DataTypeInput interface {
	pulumi.Input

	ToDataTypeOutput() DataTypeOutput
	ToDataTypeOutputWithContext(context.Context) DataTypeOutput
}

var dataTypePtrType = reflect.TypeOf((**DataType)(nil)).Elem()

type DataTypePtrInput interface {
	pulumi.Input

	ToDataTypePtrOutput() DataTypePtrOutput
	ToDataTypePtrOutputWithContext(context.Context) DataTypePtrOutput
}

type dataTypePtr string

func DataTypePtr(v string) DataTypePtrInput {
	return (*dataTypePtr)(&v)
}

func (*dataTypePtr) ElementType() reflect.Type {
	return dataTypePtrType
}

func (in *dataTypePtr) ToDataTypePtrOutput() DataTypePtrOutput {
	return pulumi.ToOutput(in).(DataTypePtrOutput)
}

func (in *dataTypePtr) ToDataTypePtrOutputWithContext(ctx context.Context) DataTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataTypePtrOutput)
}

type DatabaseAccountKind string

const (
	DatabaseAccountKindGlobalDocumentDB = DatabaseAccountKind("GlobalDocumentDB")
	DatabaseAccountKindMongoDB          = DatabaseAccountKind("MongoDB")
	DatabaseAccountKindParse            = DatabaseAccountKind("Parse")
)

func (DatabaseAccountKind) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseAccountKind)(nil)).Elem()
}

func (e DatabaseAccountKind) ToDatabaseAccountKindOutput() DatabaseAccountKindOutput {
	return pulumi.ToOutput(e).(DatabaseAccountKindOutput)
}

func (e DatabaseAccountKind) ToDatabaseAccountKindOutputWithContext(ctx context.Context) DatabaseAccountKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DatabaseAccountKindOutput)
}

func (e DatabaseAccountKind) ToDatabaseAccountKindPtrOutput() DatabaseAccountKindPtrOutput {
	return e.ToDatabaseAccountKindPtrOutputWithContext(context.Background())
}

func (e DatabaseAccountKind) ToDatabaseAccountKindPtrOutputWithContext(ctx context.Context) DatabaseAccountKindPtrOutput {
	return DatabaseAccountKind(e).ToDatabaseAccountKindOutputWithContext(ctx).ToDatabaseAccountKindPtrOutputWithContext(ctx)
}

func (e DatabaseAccountKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseAccountKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseAccountKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DatabaseAccountKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DatabaseAccountKindOutput struct{ *pulumi.OutputState }

func (DatabaseAccountKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseAccountKind)(nil)).Elem()
}

func (o DatabaseAccountKindOutput) ToDatabaseAccountKindOutput() DatabaseAccountKindOutput {
	return o
}

func (o DatabaseAccountKindOutput) ToDatabaseAccountKindOutputWithContext(ctx context.Context) DatabaseAccountKindOutput {
	return o
}

func (o DatabaseAccountKindOutput) ToDatabaseAccountKindPtrOutput() DatabaseAccountKindPtrOutput {
	return o.ToDatabaseAccountKindPtrOutputWithContext(context.Background())
}

func (o DatabaseAccountKindOutput) ToDatabaseAccountKindPtrOutputWithContext(ctx context.Context) DatabaseAccountKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabaseAccountKind) *DatabaseAccountKind {
		return &v
	}).(DatabaseAccountKindPtrOutput)
}

func (o DatabaseAccountKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DatabaseAccountKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseAccountKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DatabaseAccountKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseAccountKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseAccountKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DatabaseAccountKindPtrOutput struct{ *pulumi.OutputState }

func (DatabaseAccountKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountKind)(nil)).Elem()
}

func (o DatabaseAccountKindPtrOutput) ToDatabaseAccountKindPtrOutput() DatabaseAccountKindPtrOutput {
	return o
}

func (o DatabaseAccountKindPtrOutput) ToDatabaseAccountKindPtrOutputWithContext(ctx context.Context) DatabaseAccountKindPtrOutput {
	return o
}

func (o DatabaseAccountKindPtrOutput) Elem() DatabaseAccountKindOutput {
	return o.ApplyT(func(v *DatabaseAccountKind) DatabaseAccountKind {
		if v != nil {
			return *v
		}
		var ret DatabaseAccountKind
		return ret
	}).(DatabaseAccountKindOutput)
}

func (o DatabaseAccountKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseAccountKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DatabaseAccountKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DatabaseAccountKindInput interface {
	pulumi.Input

	ToDatabaseAccountKindOutput() DatabaseAccountKindOutput
	ToDatabaseAccountKindOutputWithContext(context.Context) DatabaseAccountKindOutput
}

var databaseAccountKindPtrType = reflect.TypeOf((**DatabaseAccountKind)(nil)).Elem()

type DatabaseAccountKindPtrInput interface {
	pulumi.Input

	ToDatabaseAccountKindPtrOutput() DatabaseAccountKindPtrOutput
	ToDatabaseAccountKindPtrOutputWithContext(context.Context) DatabaseAccountKindPtrOutput
}

type databaseAccountKindPtr string

func DatabaseAccountKindPtr(v string) DatabaseAccountKindPtrInput {
	return (*databaseAccountKindPtr)(&v)
}

func (*databaseAccountKindPtr) ElementType() reflect.Type {
	return databaseAccountKindPtrType
}

func (in *databaseAccountKindPtr) ToDatabaseAccountKindPtrOutput() DatabaseAccountKindPtrOutput {
	return pulumi.ToOutput(in).(DatabaseAccountKindPtrOutput)
}

func (in *databaseAccountKindPtr) ToDatabaseAccountKindPtrOutputWithContext(ctx context.Context) DatabaseAccountKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DatabaseAccountKindPtrOutput)
}

type DatabaseAccountOfferType string

const (
	DatabaseAccountOfferTypeStandard = DatabaseAccountOfferType("Standard")
)

func (DatabaseAccountOfferType) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseAccountOfferType)(nil)).Elem()
}

func (e DatabaseAccountOfferType) ToDatabaseAccountOfferTypeOutput() DatabaseAccountOfferTypeOutput {
	return pulumi.ToOutput(e).(DatabaseAccountOfferTypeOutput)
}

func (e DatabaseAccountOfferType) ToDatabaseAccountOfferTypeOutputWithContext(ctx context.Context) DatabaseAccountOfferTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DatabaseAccountOfferTypeOutput)
}

func (e DatabaseAccountOfferType) ToDatabaseAccountOfferTypePtrOutput() DatabaseAccountOfferTypePtrOutput {
	return e.ToDatabaseAccountOfferTypePtrOutputWithContext(context.Background())
}

func (e DatabaseAccountOfferType) ToDatabaseAccountOfferTypePtrOutputWithContext(ctx context.Context) DatabaseAccountOfferTypePtrOutput {
	return DatabaseAccountOfferType(e).ToDatabaseAccountOfferTypeOutputWithContext(ctx).ToDatabaseAccountOfferTypePtrOutputWithContext(ctx)
}

func (e DatabaseAccountOfferType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseAccountOfferType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseAccountOfferType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DatabaseAccountOfferType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DatabaseAccountOfferTypeOutput struct{ *pulumi.OutputState }

func (DatabaseAccountOfferTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseAccountOfferType)(nil)).Elem()
}

func (o DatabaseAccountOfferTypeOutput) ToDatabaseAccountOfferTypeOutput() DatabaseAccountOfferTypeOutput {
	return o
}

func (o DatabaseAccountOfferTypeOutput) ToDatabaseAccountOfferTypeOutputWithContext(ctx context.Context) DatabaseAccountOfferTypeOutput {
	return o
}

func (o DatabaseAccountOfferTypeOutput) ToDatabaseAccountOfferTypePtrOutput() DatabaseAccountOfferTypePtrOutput {
	return o.ToDatabaseAccountOfferTypePtrOutputWithContext(context.Background())
}

func (o DatabaseAccountOfferTypeOutput) ToDatabaseAccountOfferTypePtrOutputWithContext(ctx context.Context) DatabaseAccountOfferTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabaseAccountOfferType) *DatabaseAccountOfferType {
		return &v
	}).(DatabaseAccountOfferTypePtrOutput)
}

func (o DatabaseAccountOfferTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DatabaseAccountOfferTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseAccountOfferType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DatabaseAccountOfferTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseAccountOfferTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseAccountOfferType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DatabaseAccountOfferTypePtrOutput struct{ *pulumi.OutputState }

func (DatabaseAccountOfferTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountOfferType)(nil)).Elem()
}

func (o DatabaseAccountOfferTypePtrOutput) ToDatabaseAccountOfferTypePtrOutput() DatabaseAccountOfferTypePtrOutput {
	return o
}

func (o DatabaseAccountOfferTypePtrOutput) ToDatabaseAccountOfferTypePtrOutputWithContext(ctx context.Context) DatabaseAccountOfferTypePtrOutput {
	return o
}

func (o DatabaseAccountOfferTypePtrOutput) Elem() DatabaseAccountOfferTypeOutput {
	return o.ApplyT(func(v *DatabaseAccountOfferType) DatabaseAccountOfferType {
		if v != nil {
			return *v
		}
		var ret DatabaseAccountOfferType
		return ret
	}).(DatabaseAccountOfferTypeOutput)
}

func (o DatabaseAccountOfferTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseAccountOfferTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DatabaseAccountOfferType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DatabaseAccountOfferTypeInput interface {
	pulumi.Input

	ToDatabaseAccountOfferTypeOutput() DatabaseAccountOfferTypeOutput
	ToDatabaseAccountOfferTypeOutputWithContext(context.Context) DatabaseAccountOfferTypeOutput
}

var databaseAccountOfferTypePtrType = reflect.TypeOf((**DatabaseAccountOfferType)(nil)).Elem()

type DatabaseAccountOfferTypePtrInput interface {
	pulumi.Input

	ToDatabaseAccountOfferTypePtrOutput() DatabaseAccountOfferTypePtrOutput
	ToDatabaseAccountOfferTypePtrOutputWithContext(context.Context) DatabaseAccountOfferTypePtrOutput
}

type databaseAccountOfferTypePtr string

func DatabaseAccountOfferTypePtr(v string) DatabaseAccountOfferTypePtrInput {
	return (*databaseAccountOfferTypePtr)(&v)
}

func (*databaseAccountOfferTypePtr) ElementType() reflect.Type {
	return databaseAccountOfferTypePtrType
}

func (in *databaseAccountOfferTypePtr) ToDatabaseAccountOfferTypePtrOutput() DatabaseAccountOfferTypePtrOutput {
	return pulumi.ToOutput(in).(DatabaseAccountOfferTypePtrOutput)
}

func (in *databaseAccountOfferTypePtr) ToDatabaseAccountOfferTypePtrOutputWithContext(ctx context.Context) DatabaseAccountOfferTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DatabaseAccountOfferTypePtrOutput)
}

type DefaultConsistencyLevel string

const (
	DefaultConsistencyLevelEventual         = DefaultConsistencyLevel("Eventual")
	DefaultConsistencyLevelSession          = DefaultConsistencyLevel("Session")
	DefaultConsistencyLevelBoundedStaleness = DefaultConsistencyLevel("BoundedStaleness")
	DefaultConsistencyLevelStrong           = DefaultConsistencyLevel("Strong")
	DefaultConsistencyLevelConsistentPrefix = DefaultConsistencyLevel("ConsistentPrefix")
)

func (DefaultConsistencyLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultConsistencyLevel)(nil)).Elem()
}

func (e DefaultConsistencyLevel) ToDefaultConsistencyLevelOutput() DefaultConsistencyLevelOutput {
	return pulumi.ToOutput(e).(DefaultConsistencyLevelOutput)
}

func (e DefaultConsistencyLevel) ToDefaultConsistencyLevelOutputWithContext(ctx context.Context) DefaultConsistencyLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DefaultConsistencyLevelOutput)
}

func (e DefaultConsistencyLevel) ToDefaultConsistencyLevelPtrOutput() DefaultConsistencyLevelPtrOutput {
	return e.ToDefaultConsistencyLevelPtrOutputWithContext(context.Background())
}

func (e DefaultConsistencyLevel) ToDefaultConsistencyLevelPtrOutputWithContext(ctx context.Context) DefaultConsistencyLevelPtrOutput {
	return DefaultConsistencyLevel(e).ToDefaultConsistencyLevelOutputWithContext(ctx).ToDefaultConsistencyLevelPtrOutputWithContext(ctx)
}

func (e DefaultConsistencyLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultConsistencyLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultConsistencyLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DefaultConsistencyLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DefaultConsistencyLevelOutput struct{ *pulumi.OutputState }

func (DefaultConsistencyLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultConsistencyLevel)(nil)).Elem()
}

func (o DefaultConsistencyLevelOutput) ToDefaultConsistencyLevelOutput() DefaultConsistencyLevelOutput {
	return o
}

func (o DefaultConsistencyLevelOutput) ToDefaultConsistencyLevelOutputWithContext(ctx context.Context) DefaultConsistencyLevelOutput {
	return o
}

func (o DefaultConsistencyLevelOutput) ToDefaultConsistencyLevelPtrOutput() DefaultConsistencyLevelPtrOutput {
	return o.ToDefaultConsistencyLevelPtrOutputWithContext(context.Background())
}

func (o DefaultConsistencyLevelOutput) ToDefaultConsistencyLevelPtrOutputWithContext(ctx context.Context) DefaultConsistencyLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefaultConsistencyLevel) *DefaultConsistencyLevel {
		return &v
	}).(DefaultConsistencyLevelPtrOutput)
}

func (o DefaultConsistencyLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DefaultConsistencyLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefaultConsistencyLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DefaultConsistencyLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefaultConsistencyLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefaultConsistencyLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DefaultConsistencyLevelPtrOutput struct{ *pulumi.OutputState }

func (DefaultConsistencyLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultConsistencyLevel)(nil)).Elem()
}

func (o DefaultConsistencyLevelPtrOutput) ToDefaultConsistencyLevelPtrOutput() DefaultConsistencyLevelPtrOutput {
	return o
}

func (o DefaultConsistencyLevelPtrOutput) ToDefaultConsistencyLevelPtrOutputWithContext(ctx context.Context) DefaultConsistencyLevelPtrOutput {
	return o
}

func (o DefaultConsistencyLevelPtrOutput) Elem() DefaultConsistencyLevelOutput {
	return o.ApplyT(func(v *DefaultConsistencyLevel) DefaultConsistencyLevel {
		if v != nil {
			return *v
		}
		var ret DefaultConsistencyLevel
		return ret
	}).(DefaultConsistencyLevelOutput)
}

func (o DefaultConsistencyLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefaultConsistencyLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DefaultConsistencyLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DefaultConsistencyLevelInput interface {
	pulumi.Input

	ToDefaultConsistencyLevelOutput() DefaultConsistencyLevelOutput
	ToDefaultConsistencyLevelOutputWithContext(context.Context) DefaultConsistencyLevelOutput
}

var defaultConsistencyLevelPtrType = reflect.TypeOf((**DefaultConsistencyLevel)(nil)).Elem()

type DefaultConsistencyLevelPtrInput interface {
	pulumi.Input

	ToDefaultConsistencyLevelPtrOutput() DefaultConsistencyLevelPtrOutput
	ToDefaultConsistencyLevelPtrOutputWithContext(context.Context) DefaultConsistencyLevelPtrOutput
}

type defaultConsistencyLevelPtr string

func DefaultConsistencyLevelPtr(v string) DefaultConsistencyLevelPtrInput {
	return (*defaultConsistencyLevelPtr)(&v)
}

func (*defaultConsistencyLevelPtr) ElementType() reflect.Type {
	return defaultConsistencyLevelPtrType
}

func (in *defaultConsistencyLevelPtr) ToDefaultConsistencyLevelPtrOutput() DefaultConsistencyLevelPtrOutput {
	return pulumi.ToOutput(in).(DefaultConsistencyLevelPtrOutput)
}

func (in *defaultConsistencyLevelPtr) ToDefaultConsistencyLevelPtrOutputWithContext(ctx context.Context) DefaultConsistencyLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DefaultConsistencyLevelPtrOutput)
}

type IndexKind string

const (
	IndexKindHash    = IndexKind("Hash")
	IndexKindRange   = IndexKind("Range")
	IndexKindSpatial = IndexKind("Spatial")
)

func (IndexKind) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexKind)(nil)).Elem()
}

func (e IndexKind) ToIndexKindOutput() IndexKindOutput {
	return pulumi.ToOutput(e).(IndexKindOutput)
}

func (e IndexKind) ToIndexKindOutputWithContext(ctx context.Context) IndexKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IndexKindOutput)
}

func (e IndexKind) ToIndexKindPtrOutput() IndexKindPtrOutput {
	return e.ToIndexKindPtrOutputWithContext(context.Background())
}

func (e IndexKind) ToIndexKindPtrOutputWithContext(ctx context.Context) IndexKindPtrOutput {
	return IndexKind(e).ToIndexKindOutputWithContext(ctx).ToIndexKindPtrOutputWithContext(ctx)
}

func (e IndexKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IndexKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IndexKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IndexKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IndexKindOutput struct{ *pulumi.OutputState }

func (IndexKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexKind)(nil)).Elem()
}

func (o IndexKindOutput) ToIndexKindOutput() IndexKindOutput {
	return o
}

func (o IndexKindOutput) ToIndexKindOutputWithContext(ctx context.Context) IndexKindOutput {
	return o
}

func (o IndexKindOutput) ToIndexKindPtrOutput() IndexKindPtrOutput {
	return o.ToIndexKindPtrOutputWithContext(context.Background())
}

func (o IndexKindOutput) ToIndexKindPtrOutputWithContext(ctx context.Context) IndexKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IndexKind) *IndexKind {
		return &v
	}).(IndexKindPtrOutput)
}

func (o IndexKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IndexKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IndexKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IndexKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IndexKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IndexKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IndexKindPtrOutput struct{ *pulumi.OutputState }

func (IndexKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IndexKind)(nil)).Elem()
}

func (o IndexKindPtrOutput) ToIndexKindPtrOutput() IndexKindPtrOutput {
	return o
}

func (o IndexKindPtrOutput) ToIndexKindPtrOutputWithContext(ctx context.Context) IndexKindPtrOutput {
	return o
}

func (o IndexKindPtrOutput) Elem() IndexKindOutput {
	return o.ApplyT(func(v *IndexKind) IndexKind {
		if v != nil {
			return *v
		}
		var ret IndexKind
		return ret
	}).(IndexKindOutput)
}

func (o IndexKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IndexKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IndexKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IndexKindInput interface {
	pulumi.Input

	ToIndexKindOutput() IndexKindOutput
	ToIndexKindOutputWithContext(context.Context) IndexKindOutput
}

var indexKindPtrType = reflect.TypeOf((**IndexKind)(nil)).Elem()

type IndexKindPtrInput interface {
	pulumi.Input

	ToIndexKindPtrOutput() IndexKindPtrOutput
	ToIndexKindPtrOutputWithContext(context.Context) IndexKindPtrOutput
}

type indexKindPtr string

func IndexKindPtr(v string) IndexKindPtrInput {
	return (*indexKindPtr)(&v)
}

func (*indexKindPtr) ElementType() reflect.Type {
	return indexKindPtrType
}

func (in *indexKindPtr) ToIndexKindPtrOutput() IndexKindPtrOutput {
	return pulumi.ToOutput(in).(IndexKindPtrOutput)
}

func (in *indexKindPtr) ToIndexKindPtrOutputWithContext(ctx context.Context) IndexKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IndexKindPtrOutput)
}

type IndexingMode string

const (
	IndexingModeConsistent = IndexingMode("consistent")
	IndexingModeLazy       = IndexingMode("lazy")
	IndexingModeNone       = IndexingMode("none")
)

func (IndexingMode) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexingMode)(nil)).Elem()
}

func (e IndexingMode) ToIndexingModeOutput() IndexingModeOutput {
	return pulumi.ToOutput(e).(IndexingModeOutput)
}

func (e IndexingMode) ToIndexingModeOutputWithContext(ctx context.Context) IndexingModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IndexingModeOutput)
}

func (e IndexingMode) ToIndexingModePtrOutput() IndexingModePtrOutput {
	return e.ToIndexingModePtrOutputWithContext(context.Background())
}

func (e IndexingMode) ToIndexingModePtrOutputWithContext(ctx context.Context) IndexingModePtrOutput {
	return IndexingMode(e).ToIndexingModeOutputWithContext(ctx).ToIndexingModePtrOutputWithContext(ctx)
}

func (e IndexingMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IndexingMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IndexingMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IndexingMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IndexingModeOutput struct{ *pulumi.OutputState }

func (IndexingModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexingMode)(nil)).Elem()
}

func (o IndexingModeOutput) ToIndexingModeOutput() IndexingModeOutput {
	return o
}

func (o IndexingModeOutput) ToIndexingModeOutputWithContext(ctx context.Context) IndexingModeOutput {
	return o
}

func (o IndexingModeOutput) ToIndexingModePtrOutput() IndexingModePtrOutput {
	return o.ToIndexingModePtrOutputWithContext(context.Background())
}

func (o IndexingModeOutput) ToIndexingModePtrOutputWithContext(ctx context.Context) IndexingModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IndexingMode) *IndexingMode {
		return &v
	}).(IndexingModePtrOutput)
}

func (o IndexingModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IndexingModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IndexingMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IndexingModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IndexingModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IndexingMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IndexingModePtrOutput struct{ *pulumi.OutputState }

func (IndexingModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IndexingMode)(nil)).Elem()
}

func (o IndexingModePtrOutput) ToIndexingModePtrOutput() IndexingModePtrOutput {
	return o
}

func (o IndexingModePtrOutput) ToIndexingModePtrOutputWithContext(ctx context.Context) IndexingModePtrOutput {
	return o
}

func (o IndexingModePtrOutput) Elem() IndexingModeOutput {
	return o.ApplyT(func(v *IndexingMode) IndexingMode {
		if v != nil {
			return *v
		}
		var ret IndexingMode
		return ret
	}).(IndexingModeOutput)
}

func (o IndexingModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IndexingModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IndexingMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IndexingModeInput interface {
	pulumi.Input

	ToIndexingModeOutput() IndexingModeOutput
	ToIndexingModeOutputWithContext(context.Context) IndexingModeOutput
}

var indexingModePtrType = reflect.TypeOf((**IndexingMode)(nil)).Elem()

type IndexingModePtrInput interface {
	pulumi.Input

	ToIndexingModePtrOutput() IndexingModePtrOutput
	ToIndexingModePtrOutputWithContext(context.Context) IndexingModePtrOutput
}

type indexingModePtr string

func IndexingModePtr(v string) IndexingModePtrInput {
	return (*indexingModePtr)(&v)
}

func (*indexingModePtr) ElementType() reflect.Type {
	return indexingModePtrType
}

func (in *indexingModePtr) ToIndexingModePtrOutput() IndexingModePtrOutput {
	return pulumi.ToOutput(in).(IndexingModePtrOutput)
}

func (in *indexingModePtr) ToIndexingModePtrOutputWithContext(ctx context.Context) IndexingModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IndexingModePtrOutput)
}

type NetworkAclBypass string

const (
	NetworkAclBypassNone          = NetworkAclBypass("None")
	NetworkAclBypassAzureServices = NetworkAclBypass("AzureServices")
)

func (NetworkAclBypass) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAclBypass)(nil)).Elem()
}

func (e NetworkAclBypass) ToNetworkAclBypassOutput() NetworkAclBypassOutput {
	return pulumi.ToOutput(e).(NetworkAclBypassOutput)
}

func (e NetworkAclBypass) ToNetworkAclBypassOutputWithContext(ctx context.Context) NetworkAclBypassOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkAclBypassOutput)
}

func (e NetworkAclBypass) ToNetworkAclBypassPtrOutput() NetworkAclBypassPtrOutput {
	return e.ToNetworkAclBypassPtrOutputWithContext(context.Background())
}

func (e NetworkAclBypass) ToNetworkAclBypassPtrOutputWithContext(ctx context.Context) NetworkAclBypassPtrOutput {
	return NetworkAclBypass(e).ToNetworkAclBypassOutputWithContext(ctx).ToNetworkAclBypassPtrOutputWithContext(ctx)
}

func (e NetworkAclBypass) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkAclBypass) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkAclBypass) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkAclBypass) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkAclBypassOutput struct{ *pulumi.OutputState }

func (NetworkAclBypassOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAclBypass)(nil)).Elem()
}

func (o NetworkAclBypassOutput) ToNetworkAclBypassOutput() NetworkAclBypassOutput {
	return o
}

func (o NetworkAclBypassOutput) ToNetworkAclBypassOutputWithContext(ctx context.Context) NetworkAclBypassOutput {
	return o
}

func (o NetworkAclBypassOutput) ToNetworkAclBypassPtrOutput() NetworkAclBypassPtrOutput {
	return o.ToNetworkAclBypassPtrOutputWithContext(context.Background())
}

func (o NetworkAclBypassOutput) ToNetworkAclBypassPtrOutputWithContext(ctx context.Context) NetworkAclBypassPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkAclBypass) *NetworkAclBypass {
		return &v
	}).(NetworkAclBypassPtrOutput)
}

func (o NetworkAclBypassOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkAclBypassOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkAclBypass) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkAclBypassOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkAclBypassOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkAclBypass) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkAclBypassPtrOutput struct{ *pulumi.OutputState }

func (NetworkAclBypassPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAclBypass)(nil)).Elem()
}

func (o NetworkAclBypassPtrOutput) ToNetworkAclBypassPtrOutput() NetworkAclBypassPtrOutput {
	return o
}

func (o NetworkAclBypassPtrOutput) ToNetworkAclBypassPtrOutputWithContext(ctx context.Context) NetworkAclBypassPtrOutput {
	return o
}

func (o NetworkAclBypassPtrOutput) Elem() NetworkAclBypassOutput {
	return o.ApplyT(func(v *NetworkAclBypass) NetworkAclBypass {
		if v != nil {
			return *v
		}
		var ret NetworkAclBypass
		return ret
	}).(NetworkAclBypassOutput)
}

func (o NetworkAclBypassPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkAclBypassPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkAclBypass) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NetworkAclBypassInput interface {
	pulumi.Input

	ToNetworkAclBypassOutput() NetworkAclBypassOutput
	ToNetworkAclBypassOutputWithContext(context.Context) NetworkAclBypassOutput
}

var networkAclBypassPtrType = reflect.TypeOf((**NetworkAclBypass)(nil)).Elem()

type NetworkAclBypassPtrInput interface {
	pulumi.Input

	ToNetworkAclBypassPtrOutput() NetworkAclBypassPtrOutput
	ToNetworkAclBypassPtrOutputWithContext(context.Context) NetworkAclBypassPtrOutput
}

type networkAclBypassPtr string

func NetworkAclBypassPtr(v string) NetworkAclBypassPtrInput {
	return (*networkAclBypassPtr)(&v)
}

func (*networkAclBypassPtr) ElementType() reflect.Type {
	return networkAclBypassPtrType
}

func (in *networkAclBypassPtr) ToNetworkAclBypassPtrOutput() NetworkAclBypassPtrOutput {
	return pulumi.ToOutput(in).(NetworkAclBypassPtrOutput)
}

func (in *networkAclBypassPtr) ToNetworkAclBypassPtrOutputWithContext(ctx context.Context) NetworkAclBypassPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkAclBypassPtrOutput)
}

type PartitionKind string

const (
	PartitionKindHash      = PartitionKind("Hash")
	PartitionKindRange     = PartitionKind("Range")
	PartitionKindMultiHash = PartitionKind("MultiHash")
)

func (PartitionKind) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionKind)(nil)).Elem()
}

func (e PartitionKind) ToPartitionKindOutput() PartitionKindOutput {
	return pulumi.ToOutput(e).(PartitionKindOutput)
}

func (e PartitionKind) ToPartitionKindOutputWithContext(ctx context.Context) PartitionKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PartitionKindOutput)
}

func (e PartitionKind) ToPartitionKindPtrOutput() PartitionKindPtrOutput {
	return e.ToPartitionKindPtrOutputWithContext(context.Background())
}

func (e PartitionKind) ToPartitionKindPtrOutputWithContext(ctx context.Context) PartitionKindPtrOutput {
	return PartitionKind(e).ToPartitionKindOutputWithContext(ctx).ToPartitionKindPtrOutputWithContext(ctx)
}

func (e PartitionKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PartitionKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PartitionKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PartitionKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PartitionKindOutput struct{ *pulumi.OutputState }

func (PartitionKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionKind)(nil)).Elem()
}

func (o PartitionKindOutput) ToPartitionKindOutput() PartitionKindOutput {
	return o
}

func (o PartitionKindOutput) ToPartitionKindOutputWithContext(ctx context.Context) PartitionKindOutput {
	return o
}

func (o PartitionKindOutput) ToPartitionKindPtrOutput() PartitionKindPtrOutput {
	return o.ToPartitionKindPtrOutputWithContext(context.Background())
}

func (o PartitionKindOutput) ToPartitionKindPtrOutputWithContext(ctx context.Context) PartitionKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PartitionKind) *PartitionKind {
		return &v
	}).(PartitionKindPtrOutput)
}

func (o PartitionKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PartitionKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PartitionKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PartitionKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PartitionKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PartitionKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PartitionKindPtrOutput struct{ *pulumi.OutputState }

func (PartitionKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartitionKind)(nil)).Elem()
}

func (o PartitionKindPtrOutput) ToPartitionKindPtrOutput() PartitionKindPtrOutput {
	return o
}

func (o PartitionKindPtrOutput) ToPartitionKindPtrOutputWithContext(ctx context.Context) PartitionKindPtrOutput {
	return o
}

func (o PartitionKindPtrOutput) Elem() PartitionKindOutput {
	return o.ApplyT(func(v *PartitionKind) PartitionKind {
		if v != nil {
			return *v
		}
		var ret PartitionKind
		return ret
	}).(PartitionKindOutput)
}

func (o PartitionKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PartitionKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PartitionKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PartitionKindInput interface {
	pulumi.Input

	ToPartitionKindOutput() PartitionKindOutput
	ToPartitionKindOutputWithContext(context.Context) PartitionKindOutput
}

var partitionKindPtrType = reflect.TypeOf((**PartitionKind)(nil)).Elem()

type PartitionKindPtrInput interface {
	pulumi.Input

	ToPartitionKindPtrOutput() PartitionKindPtrOutput
	ToPartitionKindPtrOutputWithContext(context.Context) PartitionKindPtrOutput
}

type partitionKindPtr string

func PartitionKindPtr(v string) PartitionKindPtrInput {
	return (*partitionKindPtr)(&v)
}

func (*partitionKindPtr) ElementType() reflect.Type {
	return partitionKindPtrType
}

func (in *partitionKindPtr) ToPartitionKindPtrOutput() PartitionKindPtrOutput {
	return pulumi.ToOutput(in).(PartitionKindPtrOutput)
}

func (in *partitionKindPtr) ToPartitionKindPtrOutputWithContext(ctx context.Context) PartitionKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PartitionKindPtrOutput)
}

type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

func (PublicNetworkAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccess)(nil)).Elem()
}

func (e PublicNetworkAccess) ToPublicNetworkAccessOutput() PublicNetworkAccessOutput {
	return pulumi.ToOutput(e).(PublicNetworkAccessOutput)
}

func (e PublicNetworkAccess) ToPublicNetworkAccessOutputWithContext(ctx context.Context) PublicNetworkAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicNetworkAccessOutput)
}

func (e PublicNetworkAccess) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return e.ToPublicNetworkAccessPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccess) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return PublicNetworkAccess(e).ToPublicNetworkAccessOutputWithContext(ctx).ToPublicNetworkAccessPtrOutputWithContext(ctx)
}

func (e PublicNetworkAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicNetworkAccessOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccess)(nil)).Elem()
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessOutput() PublicNetworkAccessOutput {
	return o
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessOutputWithContext(ctx context.Context) PublicNetworkAccessOutput {
	return o
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return o.ToPublicNetworkAccessPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicNetworkAccess) *PublicNetworkAccess {
		return &v
	}).(PublicNetworkAccessPtrOutput)
}

func (o PublicNetworkAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicNetworkAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicNetworkAccessPtrOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicNetworkAccess)(nil)).Elem()
}

func (o PublicNetworkAccessPtrOutput) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return o
}

func (o PublicNetworkAccessPtrOutput) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return o
}

func (o PublicNetworkAccessPtrOutput) Elem() PublicNetworkAccessOutput {
	return o.ApplyT(func(v *PublicNetworkAccess) PublicNetworkAccess {
		if v != nil {
			return *v
		}
		var ret PublicNetworkAccess
		return ret
	}).(PublicNetworkAccessOutput)
}

func (o PublicNetworkAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicNetworkAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PublicNetworkAccessInput interface {
	pulumi.Input

	ToPublicNetworkAccessOutput() PublicNetworkAccessOutput
	ToPublicNetworkAccessOutputWithContext(context.Context) PublicNetworkAccessOutput
}

var publicNetworkAccessPtrType = reflect.TypeOf((**PublicNetworkAccess)(nil)).Elem()

type PublicNetworkAccessPtrInput interface {
	pulumi.Input

	ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput
	ToPublicNetworkAccessPtrOutputWithContext(context.Context) PublicNetworkAccessPtrOutput
}

type publicNetworkAccessPtr string

func PublicNetworkAccessPtr(v string) PublicNetworkAccessPtrInput {
	return (*publicNetworkAccessPtr)(&v)
}

func (*publicNetworkAccessPtr) ElementType() reflect.Type {
	return publicNetworkAccessPtrType
}

func (in *publicNetworkAccessPtr) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return pulumi.ToOutput(in).(PublicNetworkAccessPtrOutput)
}

func (in *publicNetworkAccessPtr) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicNetworkAccessPtrOutput)
}

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned,UserAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
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

type RestoreMode string

const (
	RestoreModePointInTime = RestoreMode("PointInTime")
)

func (RestoreMode) ElementType() reflect.Type {
	return reflect.TypeOf((*RestoreMode)(nil)).Elem()
}

func (e RestoreMode) ToRestoreModeOutput() RestoreModeOutput {
	return pulumi.ToOutput(e).(RestoreModeOutput)
}

func (e RestoreMode) ToRestoreModeOutputWithContext(ctx context.Context) RestoreModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RestoreModeOutput)
}

func (e RestoreMode) ToRestoreModePtrOutput() RestoreModePtrOutput {
	return e.ToRestoreModePtrOutputWithContext(context.Background())
}

func (e RestoreMode) ToRestoreModePtrOutputWithContext(ctx context.Context) RestoreModePtrOutput {
	return RestoreMode(e).ToRestoreModeOutputWithContext(ctx).ToRestoreModePtrOutputWithContext(ctx)
}

func (e RestoreMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RestoreMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RestoreMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RestoreMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RestoreModeOutput struct{ *pulumi.OutputState }

func (RestoreModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestoreMode)(nil)).Elem()
}

func (o RestoreModeOutput) ToRestoreModeOutput() RestoreModeOutput {
	return o
}

func (o RestoreModeOutput) ToRestoreModeOutputWithContext(ctx context.Context) RestoreModeOutput {
	return o
}

func (o RestoreModeOutput) ToRestoreModePtrOutput() RestoreModePtrOutput {
	return o.ToRestoreModePtrOutputWithContext(context.Background())
}

func (o RestoreModeOutput) ToRestoreModePtrOutputWithContext(ctx context.Context) RestoreModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RestoreMode) *RestoreMode {
		return &v
	}).(RestoreModePtrOutput)
}

func (o RestoreModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RestoreModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RestoreMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RestoreModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RestoreModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RestoreMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RestoreModePtrOutput struct{ *pulumi.OutputState }

func (RestoreModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestoreMode)(nil)).Elem()
}

func (o RestoreModePtrOutput) ToRestoreModePtrOutput() RestoreModePtrOutput {
	return o
}

func (o RestoreModePtrOutput) ToRestoreModePtrOutputWithContext(ctx context.Context) RestoreModePtrOutput {
	return o
}

func (o RestoreModePtrOutput) Elem() RestoreModeOutput {
	return o.ApplyT(func(v *RestoreMode) RestoreMode {
		if v != nil {
			return *v
		}
		var ret RestoreMode
		return ret
	}).(RestoreModeOutput)
}

func (o RestoreModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RestoreModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RestoreMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RestoreModeInput interface {
	pulumi.Input

	ToRestoreModeOutput() RestoreModeOutput
	ToRestoreModeOutputWithContext(context.Context) RestoreModeOutput
}

var restoreModePtrType = reflect.TypeOf((**RestoreMode)(nil)).Elem()

type RestoreModePtrInput interface {
	pulumi.Input

	ToRestoreModePtrOutput() RestoreModePtrOutput
	ToRestoreModePtrOutputWithContext(context.Context) RestoreModePtrOutput
}

type restoreModePtr string

func RestoreModePtr(v string) RestoreModePtrInput {
	return (*restoreModePtr)(&v)
}

func (*restoreModePtr) ElementType() reflect.Type {
	return restoreModePtrType
}

func (in *restoreModePtr) ToRestoreModePtrOutput() RestoreModePtrOutput {
	return pulumi.ToOutput(in).(RestoreModePtrOutput)
}

func (in *restoreModePtr) ToRestoreModePtrOutputWithContext(ctx context.Context) RestoreModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RestoreModePtrOutput)
}

type RoleDefinitionType string

const (
	RoleDefinitionTypeBuiltInRole = RoleDefinitionType("BuiltInRole")
	RoleDefinitionTypeCustomRole  = RoleDefinitionType("CustomRole")
)

func (RoleDefinitionType) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleDefinitionType)(nil)).Elem()
}

func (e RoleDefinitionType) ToRoleDefinitionTypeOutput() RoleDefinitionTypeOutput {
	return pulumi.ToOutput(e).(RoleDefinitionTypeOutput)
}

func (e RoleDefinitionType) ToRoleDefinitionTypeOutputWithContext(ctx context.Context) RoleDefinitionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RoleDefinitionTypeOutput)
}

func (e RoleDefinitionType) ToRoleDefinitionTypePtrOutput() RoleDefinitionTypePtrOutput {
	return e.ToRoleDefinitionTypePtrOutputWithContext(context.Background())
}

func (e RoleDefinitionType) ToRoleDefinitionTypePtrOutputWithContext(ctx context.Context) RoleDefinitionTypePtrOutput {
	return RoleDefinitionType(e).ToRoleDefinitionTypeOutputWithContext(ctx).ToRoleDefinitionTypePtrOutputWithContext(ctx)
}

func (e RoleDefinitionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RoleDefinitionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RoleDefinitionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RoleDefinitionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RoleDefinitionTypeOutput struct{ *pulumi.OutputState }

func (RoleDefinitionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleDefinitionType)(nil)).Elem()
}

func (o RoleDefinitionTypeOutput) ToRoleDefinitionTypeOutput() RoleDefinitionTypeOutput {
	return o
}

func (o RoleDefinitionTypeOutput) ToRoleDefinitionTypeOutputWithContext(ctx context.Context) RoleDefinitionTypeOutput {
	return o
}

func (o RoleDefinitionTypeOutput) ToRoleDefinitionTypePtrOutput() RoleDefinitionTypePtrOutput {
	return o.ToRoleDefinitionTypePtrOutputWithContext(context.Background())
}

func (o RoleDefinitionTypeOutput) ToRoleDefinitionTypePtrOutputWithContext(ctx context.Context) RoleDefinitionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoleDefinitionType) *RoleDefinitionType {
		return &v
	}).(RoleDefinitionTypePtrOutput)
}

func (o RoleDefinitionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RoleDefinitionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RoleDefinitionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RoleDefinitionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RoleDefinitionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RoleDefinitionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RoleDefinitionTypePtrOutput struct{ *pulumi.OutputState }

func (RoleDefinitionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleDefinitionType)(nil)).Elem()
}

func (o RoleDefinitionTypePtrOutput) ToRoleDefinitionTypePtrOutput() RoleDefinitionTypePtrOutput {
	return o
}

func (o RoleDefinitionTypePtrOutput) ToRoleDefinitionTypePtrOutputWithContext(ctx context.Context) RoleDefinitionTypePtrOutput {
	return o
}

func (o RoleDefinitionTypePtrOutput) Elem() RoleDefinitionTypeOutput {
	return o.ApplyT(func(v *RoleDefinitionType) RoleDefinitionType {
		if v != nil {
			return *v
		}
		var ret RoleDefinitionType
		return ret
	}).(RoleDefinitionTypeOutput)
}

func (o RoleDefinitionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RoleDefinitionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RoleDefinitionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RoleDefinitionTypeInput interface {
	pulumi.Input

	ToRoleDefinitionTypeOutput() RoleDefinitionTypeOutput
	ToRoleDefinitionTypeOutputWithContext(context.Context) RoleDefinitionTypeOutput
}

var roleDefinitionTypePtrType = reflect.TypeOf((**RoleDefinitionType)(nil)).Elem()

type RoleDefinitionTypePtrInput interface {
	pulumi.Input

	ToRoleDefinitionTypePtrOutput() RoleDefinitionTypePtrOutput
	ToRoleDefinitionTypePtrOutputWithContext(context.Context) RoleDefinitionTypePtrOutput
}

type roleDefinitionTypePtr string

func RoleDefinitionTypePtr(v string) RoleDefinitionTypePtrInput {
	return (*roleDefinitionTypePtr)(&v)
}

func (*roleDefinitionTypePtr) ElementType() reflect.Type {
	return roleDefinitionTypePtrType
}

func (in *roleDefinitionTypePtr) ToRoleDefinitionTypePtrOutput() RoleDefinitionTypePtrOutput {
	return pulumi.ToOutput(in).(RoleDefinitionTypePtrOutput)
}

func (in *roleDefinitionTypePtr) ToRoleDefinitionTypePtrOutputWithContext(ctx context.Context) RoleDefinitionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RoleDefinitionTypePtrOutput)
}

type ServerVersion string

const (
	ServerVersion_3_2 = ServerVersion("3.2")
	ServerVersion_3_6 = ServerVersion("3.6")
	ServerVersion_4_0 = ServerVersion("4.0")
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

type SpatialType string

const (
	SpatialTypePoint        = SpatialType("Point")
	SpatialTypeLineString   = SpatialType("LineString")
	SpatialTypePolygon      = SpatialType("Polygon")
	SpatialTypeMultiPolygon = SpatialType("MultiPolygon")
)

func (SpatialType) ElementType() reflect.Type {
	return reflect.TypeOf((*SpatialType)(nil)).Elem()
}

func (e SpatialType) ToSpatialTypeOutput() SpatialTypeOutput {
	return pulumi.ToOutput(e).(SpatialTypeOutput)
}

func (e SpatialType) ToSpatialTypeOutputWithContext(ctx context.Context) SpatialTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SpatialTypeOutput)
}

func (e SpatialType) ToSpatialTypePtrOutput() SpatialTypePtrOutput {
	return e.ToSpatialTypePtrOutputWithContext(context.Background())
}

func (e SpatialType) ToSpatialTypePtrOutputWithContext(ctx context.Context) SpatialTypePtrOutput {
	return SpatialType(e).ToSpatialTypeOutputWithContext(ctx).ToSpatialTypePtrOutputWithContext(ctx)
}

func (e SpatialType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SpatialType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SpatialType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SpatialType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SpatialTypeOutput struct{ *pulumi.OutputState }

func (SpatialTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SpatialType)(nil)).Elem()
}

func (o SpatialTypeOutput) ToSpatialTypeOutput() SpatialTypeOutput {
	return o
}

func (o SpatialTypeOutput) ToSpatialTypeOutputWithContext(ctx context.Context) SpatialTypeOutput {
	return o
}

func (o SpatialTypeOutput) ToSpatialTypePtrOutput() SpatialTypePtrOutput {
	return o.ToSpatialTypePtrOutputWithContext(context.Background())
}

func (o SpatialTypeOutput) ToSpatialTypePtrOutputWithContext(ctx context.Context) SpatialTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SpatialType) *SpatialType {
		return &v
	}).(SpatialTypePtrOutput)
}

func (o SpatialTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SpatialTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SpatialType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SpatialTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SpatialTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SpatialType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SpatialTypePtrOutput struct{ *pulumi.OutputState }

func (SpatialTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpatialType)(nil)).Elem()
}

func (o SpatialTypePtrOutput) ToSpatialTypePtrOutput() SpatialTypePtrOutput {
	return o
}

func (o SpatialTypePtrOutput) ToSpatialTypePtrOutputWithContext(ctx context.Context) SpatialTypePtrOutput {
	return o
}

func (o SpatialTypePtrOutput) Elem() SpatialTypeOutput {
	return o.ApplyT(func(v *SpatialType) SpatialType {
		if v != nil {
			return *v
		}
		var ret SpatialType
		return ret
	}).(SpatialTypeOutput)
}

func (o SpatialTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SpatialTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SpatialType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SpatialTypeInput interface {
	pulumi.Input

	ToSpatialTypeOutput() SpatialTypeOutput
	ToSpatialTypeOutputWithContext(context.Context) SpatialTypeOutput
}

var spatialTypePtrType = reflect.TypeOf((**SpatialType)(nil)).Elem()

type SpatialTypePtrInput interface {
	pulumi.Input

	ToSpatialTypePtrOutput() SpatialTypePtrOutput
	ToSpatialTypePtrOutputWithContext(context.Context) SpatialTypePtrOutput
}

type spatialTypePtr string

func SpatialTypePtr(v string) SpatialTypePtrInput {
	return (*spatialTypePtr)(&v)
}

func (*spatialTypePtr) ElementType() reflect.Type {
	return spatialTypePtrType
}

func (in *spatialTypePtr) ToSpatialTypePtrOutput() SpatialTypePtrOutput {
	return pulumi.ToOutput(in).(SpatialTypePtrOutput)
}

func (in *spatialTypePtr) ToSpatialTypePtrOutputWithContext(ctx context.Context) SpatialTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SpatialTypePtrOutput)
}

type TriggerOperation string

const (
	TriggerOperationAll     = TriggerOperation("All")
	TriggerOperationCreate  = TriggerOperation("Create")
	TriggerOperationUpdate  = TriggerOperation("Update")
	TriggerOperationDelete  = TriggerOperation("Delete")
	TriggerOperationReplace = TriggerOperation("Replace")
)

func (TriggerOperation) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerOperation)(nil)).Elem()
}

func (e TriggerOperation) ToTriggerOperationOutput() TriggerOperationOutput {
	return pulumi.ToOutput(e).(TriggerOperationOutput)
}

func (e TriggerOperation) ToTriggerOperationOutputWithContext(ctx context.Context) TriggerOperationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TriggerOperationOutput)
}

func (e TriggerOperation) ToTriggerOperationPtrOutput() TriggerOperationPtrOutput {
	return e.ToTriggerOperationPtrOutputWithContext(context.Background())
}

func (e TriggerOperation) ToTriggerOperationPtrOutputWithContext(ctx context.Context) TriggerOperationPtrOutput {
	return TriggerOperation(e).ToTriggerOperationOutputWithContext(ctx).ToTriggerOperationPtrOutputWithContext(ctx)
}

func (e TriggerOperation) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggerOperation) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggerOperation) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TriggerOperation) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TriggerOperationOutput struct{ *pulumi.OutputState }

func (TriggerOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerOperation)(nil)).Elem()
}

func (o TriggerOperationOutput) ToTriggerOperationOutput() TriggerOperationOutput {
	return o
}

func (o TriggerOperationOutput) ToTriggerOperationOutputWithContext(ctx context.Context) TriggerOperationOutput {
	return o
}

func (o TriggerOperationOutput) ToTriggerOperationPtrOutput() TriggerOperationPtrOutput {
	return o.ToTriggerOperationPtrOutputWithContext(context.Background())
}

func (o TriggerOperationOutput) ToTriggerOperationPtrOutputWithContext(ctx context.Context) TriggerOperationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TriggerOperation) *TriggerOperation {
		return &v
	}).(TriggerOperationPtrOutput)
}

func (o TriggerOperationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TriggerOperationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TriggerOperation) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TriggerOperationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TriggerOperationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TriggerOperation) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TriggerOperationPtrOutput struct{ *pulumi.OutputState }

func (TriggerOperationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerOperation)(nil)).Elem()
}

func (o TriggerOperationPtrOutput) ToTriggerOperationPtrOutput() TriggerOperationPtrOutput {
	return o
}

func (o TriggerOperationPtrOutput) ToTriggerOperationPtrOutputWithContext(ctx context.Context) TriggerOperationPtrOutput {
	return o
}

func (o TriggerOperationPtrOutput) Elem() TriggerOperationOutput {
	return o.ApplyT(func(v *TriggerOperation) TriggerOperation {
		if v != nil {
			return *v
		}
		var ret TriggerOperation
		return ret
	}).(TriggerOperationOutput)
}

func (o TriggerOperationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TriggerOperationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TriggerOperation) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TriggerOperationInput interface {
	pulumi.Input

	ToTriggerOperationOutput() TriggerOperationOutput
	ToTriggerOperationOutputWithContext(context.Context) TriggerOperationOutput
}

var triggerOperationPtrType = reflect.TypeOf((**TriggerOperation)(nil)).Elem()

type TriggerOperationPtrInput interface {
	pulumi.Input

	ToTriggerOperationPtrOutput() TriggerOperationPtrOutput
	ToTriggerOperationPtrOutputWithContext(context.Context) TriggerOperationPtrOutput
}

type triggerOperationPtr string

func TriggerOperationPtr(v string) TriggerOperationPtrInput {
	return (*triggerOperationPtr)(&v)
}

func (*triggerOperationPtr) ElementType() reflect.Type {
	return triggerOperationPtrType
}

func (in *triggerOperationPtr) ToTriggerOperationPtrOutput() TriggerOperationPtrOutput {
	return pulumi.ToOutput(in).(TriggerOperationPtrOutput)
}

func (in *triggerOperationPtr) ToTriggerOperationPtrOutputWithContext(ctx context.Context) TriggerOperationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TriggerOperationPtrOutput)
}

type TriggerType string

const (
	TriggerTypePre  = TriggerType("Pre")
	TriggerTypePost = TriggerType("Post")
)

func (TriggerType) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerType)(nil)).Elem()
}

func (e TriggerType) ToTriggerTypeOutput() TriggerTypeOutput {
	return pulumi.ToOutput(e).(TriggerTypeOutput)
}

func (e TriggerType) ToTriggerTypeOutputWithContext(ctx context.Context) TriggerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TriggerTypeOutput)
}

func (e TriggerType) ToTriggerTypePtrOutput() TriggerTypePtrOutput {
	return e.ToTriggerTypePtrOutputWithContext(context.Background())
}

func (e TriggerType) ToTriggerTypePtrOutputWithContext(ctx context.Context) TriggerTypePtrOutput {
	return TriggerType(e).ToTriggerTypeOutputWithContext(ctx).ToTriggerTypePtrOutputWithContext(ctx)
}

func (e TriggerType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggerType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggerType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TriggerType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TriggerTypeOutput struct{ *pulumi.OutputState }

func (TriggerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerType)(nil)).Elem()
}

func (o TriggerTypeOutput) ToTriggerTypeOutput() TriggerTypeOutput {
	return o
}

func (o TriggerTypeOutput) ToTriggerTypeOutputWithContext(ctx context.Context) TriggerTypeOutput {
	return o
}

func (o TriggerTypeOutput) ToTriggerTypePtrOutput() TriggerTypePtrOutput {
	return o.ToTriggerTypePtrOutputWithContext(context.Background())
}

func (o TriggerTypeOutput) ToTriggerTypePtrOutputWithContext(ctx context.Context) TriggerTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TriggerType) *TriggerType {
		return &v
	}).(TriggerTypePtrOutput)
}

func (o TriggerTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TriggerTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TriggerType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TriggerTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TriggerTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TriggerType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TriggerTypePtrOutput struct{ *pulumi.OutputState }

func (TriggerTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerType)(nil)).Elem()
}

func (o TriggerTypePtrOutput) ToTriggerTypePtrOutput() TriggerTypePtrOutput {
	return o
}

func (o TriggerTypePtrOutput) ToTriggerTypePtrOutputWithContext(ctx context.Context) TriggerTypePtrOutput {
	return o
}

func (o TriggerTypePtrOutput) Elem() TriggerTypeOutput {
	return o.ApplyT(func(v *TriggerType) TriggerType {
		if v != nil {
			return *v
		}
		var ret TriggerType
		return ret
	}).(TriggerTypeOutput)
}

func (o TriggerTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TriggerTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TriggerType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TriggerTypeInput interface {
	pulumi.Input

	ToTriggerTypeOutput() TriggerTypeOutput
	ToTriggerTypeOutputWithContext(context.Context) TriggerTypeOutput
}

var triggerTypePtrType = reflect.TypeOf((**TriggerType)(nil)).Elem()

type TriggerTypePtrInput interface {
	pulumi.Input

	ToTriggerTypePtrOutput() TriggerTypePtrOutput
	ToTriggerTypePtrOutputWithContext(context.Context) TriggerTypePtrOutput
}

type triggerTypePtr string

func TriggerTypePtr(v string) TriggerTypePtrInput {
	return (*triggerTypePtr)(&v)
}

func (*triggerTypePtr) ElementType() reflect.Type {
	return triggerTypePtrType
}

func (in *triggerTypePtr) ToTriggerTypePtrOutput() TriggerTypePtrOutput {
	return pulumi.ToOutput(in).(TriggerTypePtrOutput)
}

func (in *triggerTypePtr) ToTriggerTypePtrOutputWithContext(ctx context.Context) TriggerTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TriggerTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AnalyticalStorageSchemaTypeOutput{})
	pulumi.RegisterOutputType(AnalyticalStorageSchemaTypePtrOutput{})
	pulumi.RegisterOutputType(BackupPolicyMigrationStatusOutput{})
	pulumi.RegisterOutputType(BackupPolicyMigrationStatusPtrOutput{})
	pulumi.RegisterOutputType(BackupPolicyTypeOutput{})
	pulumi.RegisterOutputType(BackupPolicyTypePtrOutput{})
	pulumi.RegisterOutputType(CompositePathSortOrderOutput{})
	pulumi.RegisterOutputType(CompositePathSortOrderPtrOutput{})
	pulumi.RegisterOutputType(ConflictResolutionModeOutput{})
	pulumi.RegisterOutputType(ConflictResolutionModePtrOutput{})
	pulumi.RegisterOutputType(ConnectorOfferOutput{})
	pulumi.RegisterOutputType(ConnectorOfferPtrOutput{})
	pulumi.RegisterOutputType(CreateModeOutput{})
	pulumi.RegisterOutputType(CreateModePtrOutput{})
	pulumi.RegisterOutputType(DataTypeOutput{})
	pulumi.RegisterOutputType(DataTypePtrOutput{})
	pulumi.RegisterOutputType(DatabaseAccountKindOutput{})
	pulumi.RegisterOutputType(DatabaseAccountKindPtrOutput{})
	pulumi.RegisterOutputType(DatabaseAccountOfferTypeOutput{})
	pulumi.RegisterOutputType(DatabaseAccountOfferTypePtrOutput{})
	pulumi.RegisterOutputType(DefaultConsistencyLevelOutput{})
	pulumi.RegisterOutputType(DefaultConsistencyLevelPtrOutput{})
	pulumi.RegisterOutputType(IndexKindOutput{})
	pulumi.RegisterOutputType(IndexKindPtrOutput{})
	pulumi.RegisterOutputType(IndexingModeOutput{})
	pulumi.RegisterOutputType(IndexingModePtrOutput{})
	pulumi.RegisterOutputType(NetworkAclBypassOutput{})
	pulumi.RegisterOutputType(NetworkAclBypassPtrOutput{})
	pulumi.RegisterOutputType(PartitionKindOutput{})
	pulumi.RegisterOutputType(PartitionKindPtrOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(RestoreModeOutput{})
	pulumi.RegisterOutputType(RestoreModePtrOutput{})
	pulumi.RegisterOutputType(RoleDefinitionTypeOutput{})
	pulumi.RegisterOutputType(RoleDefinitionTypePtrOutput{})
	pulumi.RegisterOutputType(ServerVersionOutput{})
	pulumi.RegisterOutputType(ServerVersionPtrOutput{})
	pulumi.RegisterOutputType(SpatialTypeOutput{})
	pulumi.RegisterOutputType(SpatialTypePtrOutput{})
	pulumi.RegisterOutputType(TriggerOperationOutput{})
	pulumi.RegisterOutputType(TriggerOperationPtrOutput{})
	pulumi.RegisterOutputType(TriggerTypeOutput{})
	pulumi.RegisterOutputType(TriggerTypePtrOutput{})
}
