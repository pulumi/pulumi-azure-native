


package v20210801preview

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

type AutoExecuteStatus string

const (
	AutoExecuteStatusEnabled  = AutoExecuteStatus("Enabled")
	AutoExecuteStatusDisabled = AutoExecuteStatus("Disabled")
	AutoExecuteStatusDefault  = AutoExecuteStatus("Default")
)

func (AutoExecuteStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoExecuteStatus)(nil)).Elem()
}

func (e AutoExecuteStatus) ToAutoExecuteStatusOutput() AutoExecuteStatusOutput {
	return pulumi.ToOutput(e).(AutoExecuteStatusOutput)
}

func (e AutoExecuteStatus) ToAutoExecuteStatusOutputWithContext(ctx context.Context) AutoExecuteStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutoExecuteStatusOutput)
}

func (e AutoExecuteStatus) ToAutoExecuteStatusPtrOutput() AutoExecuteStatusPtrOutput {
	return e.ToAutoExecuteStatusPtrOutputWithContext(context.Background())
}

func (e AutoExecuteStatus) ToAutoExecuteStatusPtrOutputWithContext(ctx context.Context) AutoExecuteStatusPtrOutput {
	return AutoExecuteStatus(e).ToAutoExecuteStatusOutputWithContext(ctx).ToAutoExecuteStatusPtrOutputWithContext(ctx)
}

func (e AutoExecuteStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoExecuteStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoExecuteStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutoExecuteStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutoExecuteStatusOutput struct{ *pulumi.OutputState }

func (AutoExecuteStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoExecuteStatus)(nil)).Elem()
}

func (o AutoExecuteStatusOutput) ToAutoExecuteStatusOutput() AutoExecuteStatusOutput {
	return o
}

func (o AutoExecuteStatusOutput) ToAutoExecuteStatusOutputWithContext(ctx context.Context) AutoExecuteStatusOutput {
	return o
}

func (o AutoExecuteStatusOutput) ToAutoExecuteStatusPtrOutput() AutoExecuteStatusPtrOutput {
	return o.ToAutoExecuteStatusPtrOutputWithContext(context.Background())
}

func (o AutoExecuteStatusOutput) ToAutoExecuteStatusPtrOutputWithContext(ctx context.Context) AutoExecuteStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoExecuteStatus) *AutoExecuteStatus {
		return &v
	}).(AutoExecuteStatusPtrOutput)
}

func (o AutoExecuteStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutoExecuteStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoExecuteStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutoExecuteStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoExecuteStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoExecuteStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutoExecuteStatusPtrOutput struct{ *pulumi.OutputState }

func (AutoExecuteStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoExecuteStatus)(nil)).Elem()
}

func (o AutoExecuteStatusPtrOutput) ToAutoExecuteStatusPtrOutput() AutoExecuteStatusPtrOutput {
	return o
}

func (o AutoExecuteStatusPtrOutput) ToAutoExecuteStatusPtrOutputWithContext(ctx context.Context) AutoExecuteStatusPtrOutput {
	return o
}

func (o AutoExecuteStatusPtrOutput) Elem() AutoExecuteStatusOutput {
	return o.ApplyT(func(v *AutoExecuteStatus) AutoExecuteStatus {
		if v != nil {
			return *v
		}
		var ret AutoExecuteStatus
		return ret
	}).(AutoExecuteStatusOutput)
}

func (o AutoExecuteStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoExecuteStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutoExecuteStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AutoExecuteStatusInput interface {
	pulumi.Input

	ToAutoExecuteStatusOutput() AutoExecuteStatusOutput
	ToAutoExecuteStatusOutputWithContext(context.Context) AutoExecuteStatusOutput
}

var autoExecuteStatusPtrType = reflect.TypeOf((**AutoExecuteStatus)(nil)).Elem()

type AutoExecuteStatusPtrInput interface {
	pulumi.Input

	ToAutoExecuteStatusPtrOutput() AutoExecuteStatusPtrOutput
	ToAutoExecuteStatusPtrOutputWithContext(context.Context) AutoExecuteStatusPtrOutput
}

type autoExecuteStatusPtr string

func AutoExecuteStatusPtr(v string) AutoExecuteStatusPtrInput {
	return (*autoExecuteStatusPtr)(&v)
}

func (*autoExecuteStatusPtr) ElementType() reflect.Type {
	return autoExecuteStatusPtrType
}

func (in *autoExecuteStatusPtr) ToAutoExecuteStatusPtrOutput() AutoExecuteStatusPtrOutput {
	return pulumi.ToOutput(in).(AutoExecuteStatusPtrOutput)
}

func (in *autoExecuteStatusPtr) ToAutoExecuteStatusPtrOutputWithContext(ctx context.Context) AutoExecuteStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutoExecuteStatusPtrOutput)
}

type BackupStorageRedundancy string

const (
	BackupStorageRedundancyGeo     = BackupStorageRedundancy("Geo")
	BackupStorageRedundancyLocal   = BackupStorageRedundancy("Local")
	BackupStorageRedundancyZone    = BackupStorageRedundancy("Zone")
	BackupStorageRedundancyGeoZone = BackupStorageRedundancy("GeoZone")
)

func (BackupStorageRedundancy) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupStorageRedundancy)(nil)).Elem()
}

func (e BackupStorageRedundancy) ToBackupStorageRedundancyOutput() BackupStorageRedundancyOutput {
	return pulumi.ToOutput(e).(BackupStorageRedundancyOutput)
}

func (e BackupStorageRedundancy) ToBackupStorageRedundancyOutputWithContext(ctx context.Context) BackupStorageRedundancyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BackupStorageRedundancyOutput)
}

func (e BackupStorageRedundancy) ToBackupStorageRedundancyPtrOutput() BackupStorageRedundancyPtrOutput {
	return e.ToBackupStorageRedundancyPtrOutputWithContext(context.Background())
}

func (e BackupStorageRedundancy) ToBackupStorageRedundancyPtrOutputWithContext(ctx context.Context) BackupStorageRedundancyPtrOutput {
	return BackupStorageRedundancy(e).ToBackupStorageRedundancyOutputWithContext(ctx).ToBackupStorageRedundancyPtrOutputWithContext(ctx)
}

func (e BackupStorageRedundancy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackupStorageRedundancy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackupStorageRedundancy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BackupStorageRedundancy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BackupStorageRedundancyOutput struct{ *pulumi.OutputState }

func (BackupStorageRedundancyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupStorageRedundancy)(nil)).Elem()
}

func (o BackupStorageRedundancyOutput) ToBackupStorageRedundancyOutput() BackupStorageRedundancyOutput {
	return o
}

func (o BackupStorageRedundancyOutput) ToBackupStorageRedundancyOutputWithContext(ctx context.Context) BackupStorageRedundancyOutput {
	return o
}

func (o BackupStorageRedundancyOutput) ToBackupStorageRedundancyPtrOutput() BackupStorageRedundancyPtrOutput {
	return o.ToBackupStorageRedundancyPtrOutputWithContext(context.Background())
}

func (o BackupStorageRedundancyOutput) ToBackupStorageRedundancyPtrOutputWithContext(ctx context.Context) BackupStorageRedundancyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupStorageRedundancy) *BackupStorageRedundancy {
		return &v
	}).(BackupStorageRedundancyPtrOutput)
}

func (o BackupStorageRedundancyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BackupStorageRedundancyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BackupStorageRedundancy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BackupStorageRedundancyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BackupStorageRedundancyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BackupStorageRedundancy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BackupStorageRedundancyPtrOutput struct{ *pulumi.OutputState }

func (BackupStorageRedundancyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupStorageRedundancy)(nil)).Elem()
}

func (o BackupStorageRedundancyPtrOutput) ToBackupStorageRedundancyPtrOutput() BackupStorageRedundancyPtrOutput {
	return o
}

func (o BackupStorageRedundancyPtrOutput) ToBackupStorageRedundancyPtrOutputWithContext(ctx context.Context) BackupStorageRedundancyPtrOutput {
	return o
}

func (o BackupStorageRedundancyPtrOutput) Elem() BackupStorageRedundancyOutput {
	return o.ApplyT(func(v *BackupStorageRedundancy) BackupStorageRedundancy {
		if v != nil {
			return *v
		}
		var ret BackupStorageRedundancy
		return ret
	}).(BackupStorageRedundancyOutput)
}

func (o BackupStorageRedundancyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BackupStorageRedundancyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BackupStorageRedundancy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BackupStorageRedundancyInput interface {
	pulumi.Input

	ToBackupStorageRedundancyOutput() BackupStorageRedundancyOutput
	ToBackupStorageRedundancyOutputWithContext(context.Context) BackupStorageRedundancyOutput
}

var backupStorageRedundancyPtrType = reflect.TypeOf((**BackupStorageRedundancy)(nil)).Elem()

type BackupStorageRedundancyPtrInput interface {
	pulumi.Input

	ToBackupStorageRedundancyPtrOutput() BackupStorageRedundancyPtrOutput
	ToBackupStorageRedundancyPtrOutputWithContext(context.Context) BackupStorageRedundancyPtrOutput
}

type backupStorageRedundancyPtr string

func BackupStorageRedundancyPtr(v string) BackupStorageRedundancyPtrInput {
	return (*backupStorageRedundancyPtr)(&v)
}

func (*backupStorageRedundancyPtr) ElementType() reflect.Type {
	return backupStorageRedundancyPtrType
}

func (in *backupStorageRedundancyPtr) ToBackupStorageRedundancyPtrOutput() BackupStorageRedundancyPtrOutput {
	return pulumi.ToOutput(in).(BackupStorageRedundancyPtrOutput)
}

func (in *backupStorageRedundancyPtr) ToBackupStorageRedundancyPtrOutputWithContext(ctx context.Context) BackupStorageRedundancyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BackupStorageRedundancyPtrOutput)
}

type BlobAuditingPolicyState string

const (
	BlobAuditingPolicyStateEnabled  = BlobAuditingPolicyState("Enabled")
	BlobAuditingPolicyStateDisabled = BlobAuditingPolicyState("Disabled")
)

func (BlobAuditingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobAuditingPolicyState)(nil)).Elem()
}

func (e BlobAuditingPolicyState) ToBlobAuditingPolicyStateOutput() BlobAuditingPolicyStateOutput {
	return pulumi.ToOutput(e).(BlobAuditingPolicyStateOutput)
}

func (e BlobAuditingPolicyState) ToBlobAuditingPolicyStateOutputWithContext(ctx context.Context) BlobAuditingPolicyStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BlobAuditingPolicyStateOutput)
}

func (e BlobAuditingPolicyState) ToBlobAuditingPolicyStatePtrOutput() BlobAuditingPolicyStatePtrOutput {
	return e.ToBlobAuditingPolicyStatePtrOutputWithContext(context.Background())
}

func (e BlobAuditingPolicyState) ToBlobAuditingPolicyStatePtrOutputWithContext(ctx context.Context) BlobAuditingPolicyStatePtrOutput {
	return BlobAuditingPolicyState(e).ToBlobAuditingPolicyStateOutputWithContext(ctx).ToBlobAuditingPolicyStatePtrOutputWithContext(ctx)
}

func (e BlobAuditingPolicyState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlobAuditingPolicyState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlobAuditingPolicyState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BlobAuditingPolicyState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BlobAuditingPolicyStateOutput struct{ *pulumi.OutputState }

func (BlobAuditingPolicyStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobAuditingPolicyState)(nil)).Elem()
}

func (o BlobAuditingPolicyStateOutput) ToBlobAuditingPolicyStateOutput() BlobAuditingPolicyStateOutput {
	return o
}

func (o BlobAuditingPolicyStateOutput) ToBlobAuditingPolicyStateOutputWithContext(ctx context.Context) BlobAuditingPolicyStateOutput {
	return o
}

func (o BlobAuditingPolicyStateOutput) ToBlobAuditingPolicyStatePtrOutput() BlobAuditingPolicyStatePtrOutput {
	return o.ToBlobAuditingPolicyStatePtrOutputWithContext(context.Background())
}

func (o BlobAuditingPolicyStateOutput) ToBlobAuditingPolicyStatePtrOutputWithContext(ctx context.Context) BlobAuditingPolicyStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlobAuditingPolicyState) *BlobAuditingPolicyState {
		return &v
	}).(BlobAuditingPolicyStatePtrOutput)
}

func (o BlobAuditingPolicyStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BlobAuditingPolicyStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlobAuditingPolicyState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BlobAuditingPolicyStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlobAuditingPolicyStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlobAuditingPolicyState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BlobAuditingPolicyStatePtrOutput struct{ *pulumi.OutputState }

func (BlobAuditingPolicyStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobAuditingPolicyState)(nil)).Elem()
}

func (o BlobAuditingPolicyStatePtrOutput) ToBlobAuditingPolicyStatePtrOutput() BlobAuditingPolicyStatePtrOutput {
	return o
}

func (o BlobAuditingPolicyStatePtrOutput) ToBlobAuditingPolicyStatePtrOutputWithContext(ctx context.Context) BlobAuditingPolicyStatePtrOutput {
	return o
}

func (o BlobAuditingPolicyStatePtrOutput) Elem() BlobAuditingPolicyStateOutput {
	return o.ApplyT(func(v *BlobAuditingPolicyState) BlobAuditingPolicyState {
		if v != nil {
			return *v
		}
		var ret BlobAuditingPolicyState
		return ret
	}).(BlobAuditingPolicyStateOutput)
}

func (o BlobAuditingPolicyStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlobAuditingPolicyStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BlobAuditingPolicyState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BlobAuditingPolicyStateInput interface {
	pulumi.Input

	ToBlobAuditingPolicyStateOutput() BlobAuditingPolicyStateOutput
	ToBlobAuditingPolicyStateOutputWithContext(context.Context) BlobAuditingPolicyStateOutput
}

var blobAuditingPolicyStatePtrType = reflect.TypeOf((**BlobAuditingPolicyState)(nil)).Elem()

type BlobAuditingPolicyStatePtrInput interface {
	pulumi.Input

	ToBlobAuditingPolicyStatePtrOutput() BlobAuditingPolicyStatePtrOutput
	ToBlobAuditingPolicyStatePtrOutputWithContext(context.Context) BlobAuditingPolicyStatePtrOutput
}

type blobAuditingPolicyStatePtr string

func BlobAuditingPolicyStatePtr(v string) BlobAuditingPolicyStatePtrInput {
	return (*blobAuditingPolicyStatePtr)(&v)
}

func (*blobAuditingPolicyStatePtr) ElementType() reflect.Type {
	return blobAuditingPolicyStatePtrType
}

func (in *blobAuditingPolicyStatePtr) ToBlobAuditingPolicyStatePtrOutput() BlobAuditingPolicyStatePtrOutput {
	return pulumi.ToOutput(in).(BlobAuditingPolicyStatePtrOutput)
}

func (in *blobAuditingPolicyStatePtr) ToBlobAuditingPolicyStatePtrOutputWithContext(ctx context.Context) BlobAuditingPolicyStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BlobAuditingPolicyStatePtrOutput)
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

type CreateMode string

const (
	CreateModeDefault                        = CreateMode("Default")
	CreateModeCopy                           = CreateMode("Copy")
	CreateModeSecondary                      = CreateMode("Secondary")
	CreateModePointInTimeRestore             = CreateMode("PointInTimeRestore")
	CreateModeRestore                        = CreateMode("Restore")
	CreateModeRecovery                       = CreateMode("Recovery")
	CreateModeRestoreExternalBackup          = CreateMode("RestoreExternalBackup")
	CreateModeRestoreExternalBackupSecondary = CreateMode("RestoreExternalBackupSecondary")
	CreateModeRestoreLongTermRetentionBackup = CreateMode("RestoreLongTermRetentionBackup")
	CreateModeOnlineSecondary                = CreateMode("OnlineSecondary")
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

type DatabaseIdentityType string

const (
	DatabaseIdentityTypeNone         = DatabaseIdentityType("None")
	DatabaseIdentityTypeUserAssigned = DatabaseIdentityType("UserAssigned")
)

func (DatabaseIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseIdentityType)(nil)).Elem()
}

func (e DatabaseIdentityType) ToDatabaseIdentityTypeOutput() DatabaseIdentityTypeOutput {
	return pulumi.ToOutput(e).(DatabaseIdentityTypeOutput)
}

func (e DatabaseIdentityType) ToDatabaseIdentityTypeOutputWithContext(ctx context.Context) DatabaseIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DatabaseIdentityTypeOutput)
}

func (e DatabaseIdentityType) ToDatabaseIdentityTypePtrOutput() DatabaseIdentityTypePtrOutput {
	return e.ToDatabaseIdentityTypePtrOutputWithContext(context.Background())
}

func (e DatabaseIdentityType) ToDatabaseIdentityTypePtrOutputWithContext(ctx context.Context) DatabaseIdentityTypePtrOutput {
	return DatabaseIdentityType(e).ToDatabaseIdentityTypeOutputWithContext(ctx).ToDatabaseIdentityTypePtrOutputWithContext(ctx)
}

func (e DatabaseIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DatabaseIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DatabaseIdentityTypeOutput struct{ *pulumi.OutputState }

func (DatabaseIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseIdentityType)(nil)).Elem()
}

func (o DatabaseIdentityTypeOutput) ToDatabaseIdentityTypeOutput() DatabaseIdentityTypeOutput {
	return o
}

func (o DatabaseIdentityTypeOutput) ToDatabaseIdentityTypeOutputWithContext(ctx context.Context) DatabaseIdentityTypeOutput {
	return o
}

func (o DatabaseIdentityTypeOutput) ToDatabaseIdentityTypePtrOutput() DatabaseIdentityTypePtrOutput {
	return o.ToDatabaseIdentityTypePtrOutputWithContext(context.Background())
}

func (o DatabaseIdentityTypeOutput) ToDatabaseIdentityTypePtrOutputWithContext(ctx context.Context) DatabaseIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabaseIdentityType) *DatabaseIdentityType {
		return &v
	}).(DatabaseIdentityTypePtrOutput)
}

func (o DatabaseIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DatabaseIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DatabaseIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DatabaseIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (DatabaseIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseIdentityType)(nil)).Elem()
}

func (o DatabaseIdentityTypePtrOutput) ToDatabaseIdentityTypePtrOutput() DatabaseIdentityTypePtrOutput {
	return o
}

func (o DatabaseIdentityTypePtrOutput) ToDatabaseIdentityTypePtrOutputWithContext(ctx context.Context) DatabaseIdentityTypePtrOutput {
	return o
}

func (o DatabaseIdentityTypePtrOutput) Elem() DatabaseIdentityTypeOutput {
	return o.ApplyT(func(v *DatabaseIdentityType) DatabaseIdentityType {
		if v != nil {
			return *v
		}
		var ret DatabaseIdentityType
		return ret
	}).(DatabaseIdentityTypeOutput)
}

func (o DatabaseIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DatabaseIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DatabaseIdentityTypeInput interface {
	pulumi.Input

	ToDatabaseIdentityTypeOutput() DatabaseIdentityTypeOutput
	ToDatabaseIdentityTypeOutputWithContext(context.Context) DatabaseIdentityTypeOutput
}

var databaseIdentityTypePtrType = reflect.TypeOf((**DatabaseIdentityType)(nil)).Elem()

type DatabaseIdentityTypePtrInput interface {
	pulumi.Input

	ToDatabaseIdentityTypePtrOutput() DatabaseIdentityTypePtrOutput
	ToDatabaseIdentityTypePtrOutputWithContext(context.Context) DatabaseIdentityTypePtrOutput
}

type databaseIdentityTypePtr string

func DatabaseIdentityTypePtr(v string) DatabaseIdentityTypePtrInput {
	return (*databaseIdentityTypePtr)(&v)
}

func (*databaseIdentityTypePtr) ElementType() reflect.Type {
	return databaseIdentityTypePtrType
}

func (in *databaseIdentityTypePtr) ToDatabaseIdentityTypePtrOutput() DatabaseIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(DatabaseIdentityTypePtrOutput)
}

func (in *databaseIdentityTypePtr) ToDatabaseIdentityTypePtrOutputWithContext(ctx context.Context) DatabaseIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DatabaseIdentityTypePtrOutput)
}

type DatabaseLicenseType string

const (
	DatabaseLicenseTypeLicenseIncluded = DatabaseLicenseType("LicenseIncluded")
	DatabaseLicenseTypeBasePrice       = DatabaseLicenseType("BasePrice")
)

func (DatabaseLicenseType) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseLicenseType)(nil)).Elem()
}

func (e DatabaseLicenseType) ToDatabaseLicenseTypeOutput() DatabaseLicenseTypeOutput {
	return pulumi.ToOutput(e).(DatabaseLicenseTypeOutput)
}

func (e DatabaseLicenseType) ToDatabaseLicenseTypeOutputWithContext(ctx context.Context) DatabaseLicenseTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DatabaseLicenseTypeOutput)
}

func (e DatabaseLicenseType) ToDatabaseLicenseTypePtrOutput() DatabaseLicenseTypePtrOutput {
	return e.ToDatabaseLicenseTypePtrOutputWithContext(context.Background())
}

func (e DatabaseLicenseType) ToDatabaseLicenseTypePtrOutputWithContext(ctx context.Context) DatabaseLicenseTypePtrOutput {
	return DatabaseLicenseType(e).ToDatabaseLicenseTypeOutputWithContext(ctx).ToDatabaseLicenseTypePtrOutputWithContext(ctx)
}

func (e DatabaseLicenseType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseLicenseType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseLicenseType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DatabaseLicenseType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DatabaseLicenseTypeOutput struct{ *pulumi.OutputState }

func (DatabaseLicenseTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseLicenseType)(nil)).Elem()
}

func (o DatabaseLicenseTypeOutput) ToDatabaseLicenseTypeOutput() DatabaseLicenseTypeOutput {
	return o
}

func (o DatabaseLicenseTypeOutput) ToDatabaseLicenseTypeOutputWithContext(ctx context.Context) DatabaseLicenseTypeOutput {
	return o
}

func (o DatabaseLicenseTypeOutput) ToDatabaseLicenseTypePtrOutput() DatabaseLicenseTypePtrOutput {
	return o.ToDatabaseLicenseTypePtrOutputWithContext(context.Background())
}

func (o DatabaseLicenseTypeOutput) ToDatabaseLicenseTypePtrOutputWithContext(ctx context.Context) DatabaseLicenseTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabaseLicenseType) *DatabaseLicenseType {
		return &v
	}).(DatabaseLicenseTypePtrOutput)
}

func (o DatabaseLicenseTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DatabaseLicenseTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseLicenseType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DatabaseLicenseTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseLicenseTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseLicenseType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DatabaseLicenseTypePtrOutput struct{ *pulumi.OutputState }

func (DatabaseLicenseTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseLicenseType)(nil)).Elem()
}

func (o DatabaseLicenseTypePtrOutput) ToDatabaseLicenseTypePtrOutput() DatabaseLicenseTypePtrOutput {
	return o
}

func (o DatabaseLicenseTypePtrOutput) ToDatabaseLicenseTypePtrOutputWithContext(ctx context.Context) DatabaseLicenseTypePtrOutput {
	return o
}

func (o DatabaseLicenseTypePtrOutput) Elem() DatabaseLicenseTypeOutput {
	return o.ApplyT(func(v *DatabaseLicenseType) DatabaseLicenseType {
		if v != nil {
			return *v
		}
		var ret DatabaseLicenseType
		return ret
	}).(DatabaseLicenseTypeOutput)
}

func (o DatabaseLicenseTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseLicenseTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DatabaseLicenseType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DatabaseLicenseTypeInput interface {
	pulumi.Input

	ToDatabaseLicenseTypeOutput() DatabaseLicenseTypeOutput
	ToDatabaseLicenseTypeOutputWithContext(context.Context) DatabaseLicenseTypeOutput
}

var databaseLicenseTypePtrType = reflect.TypeOf((**DatabaseLicenseType)(nil)).Elem()

type DatabaseLicenseTypePtrInput interface {
	pulumi.Input

	ToDatabaseLicenseTypePtrOutput() DatabaseLicenseTypePtrOutput
	ToDatabaseLicenseTypePtrOutputWithContext(context.Context) DatabaseLicenseTypePtrOutput
}

type databaseLicenseTypePtr string

func DatabaseLicenseTypePtr(v string) DatabaseLicenseTypePtrInput {
	return (*databaseLicenseTypePtr)(&v)
}

func (*databaseLicenseTypePtr) ElementType() reflect.Type {
	return databaseLicenseTypePtrType
}

func (in *databaseLicenseTypePtr) ToDatabaseLicenseTypePtrOutput() DatabaseLicenseTypePtrOutput {
	return pulumi.ToOutput(in).(DatabaseLicenseTypePtrOutput)
}

func (in *databaseLicenseTypePtr) ToDatabaseLicenseTypePtrOutputWithContext(ctx context.Context) DatabaseLicenseTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DatabaseLicenseTypePtrOutput)
}

type DatabaseReadScale string

const (
	DatabaseReadScaleEnabled  = DatabaseReadScale("Enabled")
	DatabaseReadScaleDisabled = DatabaseReadScale("Disabled")
)

func (DatabaseReadScale) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseReadScale)(nil)).Elem()
}

func (e DatabaseReadScale) ToDatabaseReadScaleOutput() DatabaseReadScaleOutput {
	return pulumi.ToOutput(e).(DatabaseReadScaleOutput)
}

func (e DatabaseReadScale) ToDatabaseReadScaleOutputWithContext(ctx context.Context) DatabaseReadScaleOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DatabaseReadScaleOutput)
}

func (e DatabaseReadScale) ToDatabaseReadScalePtrOutput() DatabaseReadScalePtrOutput {
	return e.ToDatabaseReadScalePtrOutputWithContext(context.Background())
}

func (e DatabaseReadScale) ToDatabaseReadScalePtrOutputWithContext(ctx context.Context) DatabaseReadScalePtrOutput {
	return DatabaseReadScale(e).ToDatabaseReadScaleOutputWithContext(ctx).ToDatabaseReadScalePtrOutputWithContext(ctx)
}

func (e DatabaseReadScale) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseReadScale) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseReadScale) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DatabaseReadScale) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DatabaseReadScaleOutput struct{ *pulumi.OutputState }

func (DatabaseReadScaleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseReadScale)(nil)).Elem()
}

func (o DatabaseReadScaleOutput) ToDatabaseReadScaleOutput() DatabaseReadScaleOutput {
	return o
}

func (o DatabaseReadScaleOutput) ToDatabaseReadScaleOutputWithContext(ctx context.Context) DatabaseReadScaleOutput {
	return o
}

func (o DatabaseReadScaleOutput) ToDatabaseReadScalePtrOutput() DatabaseReadScalePtrOutput {
	return o.ToDatabaseReadScalePtrOutputWithContext(context.Background())
}

func (o DatabaseReadScaleOutput) ToDatabaseReadScalePtrOutputWithContext(ctx context.Context) DatabaseReadScalePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabaseReadScale) *DatabaseReadScale {
		return &v
	}).(DatabaseReadScalePtrOutput)
}

func (o DatabaseReadScaleOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DatabaseReadScaleOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseReadScale) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DatabaseReadScaleOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseReadScaleOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseReadScale) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DatabaseReadScalePtrOutput struct{ *pulumi.OutputState }

func (DatabaseReadScalePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseReadScale)(nil)).Elem()
}

func (o DatabaseReadScalePtrOutput) ToDatabaseReadScalePtrOutput() DatabaseReadScalePtrOutput {
	return o
}

func (o DatabaseReadScalePtrOutput) ToDatabaseReadScalePtrOutputWithContext(ctx context.Context) DatabaseReadScalePtrOutput {
	return o
}

func (o DatabaseReadScalePtrOutput) Elem() DatabaseReadScaleOutput {
	return o.ApplyT(func(v *DatabaseReadScale) DatabaseReadScale {
		if v != nil {
			return *v
		}
		var ret DatabaseReadScale
		return ret
	}).(DatabaseReadScaleOutput)
}

func (o DatabaseReadScalePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseReadScalePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DatabaseReadScale) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DatabaseReadScaleInput interface {
	pulumi.Input

	ToDatabaseReadScaleOutput() DatabaseReadScaleOutput
	ToDatabaseReadScaleOutputWithContext(context.Context) DatabaseReadScaleOutput
}

var databaseReadScalePtrType = reflect.TypeOf((**DatabaseReadScale)(nil)).Elem()

type DatabaseReadScalePtrInput interface {
	pulumi.Input

	ToDatabaseReadScalePtrOutput() DatabaseReadScalePtrOutput
	ToDatabaseReadScalePtrOutputWithContext(context.Context) DatabaseReadScalePtrOutput
}

type databaseReadScalePtr string

func DatabaseReadScalePtr(v string) DatabaseReadScalePtrInput {
	return (*databaseReadScalePtr)(&v)
}

func (*databaseReadScalePtr) ElementType() reflect.Type {
	return databaseReadScalePtrType
}

func (in *databaseReadScalePtr) ToDatabaseReadScalePtrOutput() DatabaseReadScalePtrOutput {
	return pulumi.ToOutput(in).(DatabaseReadScalePtrOutput)
}

func (in *databaseReadScalePtr) ToDatabaseReadScalePtrOutputWithContext(ctx context.Context) DatabaseReadScalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DatabaseReadScalePtrOutput)
}

type ElasticPoolLicenseType string

const (
	ElasticPoolLicenseTypeLicenseIncluded = ElasticPoolLicenseType("LicenseIncluded")
	ElasticPoolLicenseTypeBasePrice       = ElasticPoolLicenseType("BasePrice")
)

func (ElasticPoolLicenseType) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPoolLicenseType)(nil)).Elem()
}

func (e ElasticPoolLicenseType) ToElasticPoolLicenseTypeOutput() ElasticPoolLicenseTypeOutput {
	return pulumi.ToOutput(e).(ElasticPoolLicenseTypeOutput)
}

func (e ElasticPoolLicenseType) ToElasticPoolLicenseTypeOutputWithContext(ctx context.Context) ElasticPoolLicenseTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ElasticPoolLicenseTypeOutput)
}

func (e ElasticPoolLicenseType) ToElasticPoolLicenseTypePtrOutput() ElasticPoolLicenseTypePtrOutput {
	return e.ToElasticPoolLicenseTypePtrOutputWithContext(context.Background())
}

func (e ElasticPoolLicenseType) ToElasticPoolLicenseTypePtrOutputWithContext(ctx context.Context) ElasticPoolLicenseTypePtrOutput {
	return ElasticPoolLicenseType(e).ToElasticPoolLicenseTypeOutputWithContext(ctx).ToElasticPoolLicenseTypePtrOutputWithContext(ctx)
}

func (e ElasticPoolLicenseType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ElasticPoolLicenseType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ElasticPoolLicenseType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ElasticPoolLicenseType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ElasticPoolLicenseTypeOutput struct{ *pulumi.OutputState }

func (ElasticPoolLicenseTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElasticPoolLicenseType)(nil)).Elem()
}

func (o ElasticPoolLicenseTypeOutput) ToElasticPoolLicenseTypeOutput() ElasticPoolLicenseTypeOutput {
	return o
}

func (o ElasticPoolLicenseTypeOutput) ToElasticPoolLicenseTypeOutputWithContext(ctx context.Context) ElasticPoolLicenseTypeOutput {
	return o
}

func (o ElasticPoolLicenseTypeOutput) ToElasticPoolLicenseTypePtrOutput() ElasticPoolLicenseTypePtrOutput {
	return o.ToElasticPoolLicenseTypePtrOutputWithContext(context.Background())
}

func (o ElasticPoolLicenseTypeOutput) ToElasticPoolLicenseTypePtrOutputWithContext(ctx context.Context) ElasticPoolLicenseTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ElasticPoolLicenseType) *ElasticPoolLicenseType {
		return &v
	}).(ElasticPoolLicenseTypePtrOutput)
}

func (o ElasticPoolLicenseTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ElasticPoolLicenseTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ElasticPoolLicenseType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ElasticPoolLicenseTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ElasticPoolLicenseTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ElasticPoolLicenseType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ElasticPoolLicenseTypePtrOutput struct{ *pulumi.OutputState }

func (ElasticPoolLicenseTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticPoolLicenseType)(nil)).Elem()
}

func (o ElasticPoolLicenseTypePtrOutput) ToElasticPoolLicenseTypePtrOutput() ElasticPoolLicenseTypePtrOutput {
	return o
}

func (o ElasticPoolLicenseTypePtrOutput) ToElasticPoolLicenseTypePtrOutputWithContext(ctx context.Context) ElasticPoolLicenseTypePtrOutput {
	return o
}

func (o ElasticPoolLicenseTypePtrOutput) Elem() ElasticPoolLicenseTypeOutput {
	return o.ApplyT(func(v *ElasticPoolLicenseType) ElasticPoolLicenseType {
		if v != nil {
			return *v
		}
		var ret ElasticPoolLicenseType
		return ret
	}).(ElasticPoolLicenseTypeOutput)
}

func (o ElasticPoolLicenseTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ElasticPoolLicenseTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ElasticPoolLicenseType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ElasticPoolLicenseTypeInput interface {
	pulumi.Input

	ToElasticPoolLicenseTypeOutput() ElasticPoolLicenseTypeOutput
	ToElasticPoolLicenseTypeOutputWithContext(context.Context) ElasticPoolLicenseTypeOutput
}

var elasticPoolLicenseTypePtrType = reflect.TypeOf((**ElasticPoolLicenseType)(nil)).Elem()

type ElasticPoolLicenseTypePtrInput interface {
	pulumi.Input

	ToElasticPoolLicenseTypePtrOutput() ElasticPoolLicenseTypePtrOutput
	ToElasticPoolLicenseTypePtrOutputWithContext(context.Context) ElasticPoolLicenseTypePtrOutput
}

type elasticPoolLicenseTypePtr string

func ElasticPoolLicenseTypePtr(v string) ElasticPoolLicenseTypePtrInput {
	return (*elasticPoolLicenseTypePtr)(&v)
}

func (*elasticPoolLicenseTypePtr) ElementType() reflect.Type {
	return elasticPoolLicenseTypePtrType
}

func (in *elasticPoolLicenseTypePtr) ToElasticPoolLicenseTypePtrOutput() ElasticPoolLicenseTypePtrOutput {
	return pulumi.ToOutput(in).(ElasticPoolLicenseTypePtrOutput)
}

func (in *elasticPoolLicenseTypePtr) ToElasticPoolLicenseTypePtrOutputWithContext(ctx context.Context) ElasticPoolLicenseTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ElasticPoolLicenseTypePtrOutput)
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

type JobScheduleType string

const (
	JobScheduleTypeOnce      = JobScheduleType("Once")
	JobScheduleTypeRecurring = JobScheduleType("Recurring")
)

func (JobScheduleType) ElementType() reflect.Type {
	return reflect.TypeOf((*JobScheduleType)(nil)).Elem()
}

func (e JobScheduleType) ToJobScheduleTypeOutput() JobScheduleTypeOutput {
	return pulumi.ToOutput(e).(JobScheduleTypeOutput)
}

func (e JobScheduleType) ToJobScheduleTypeOutputWithContext(ctx context.Context) JobScheduleTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(JobScheduleTypeOutput)
}

func (e JobScheduleType) ToJobScheduleTypePtrOutput() JobScheduleTypePtrOutput {
	return e.ToJobScheduleTypePtrOutputWithContext(context.Background())
}

func (e JobScheduleType) ToJobScheduleTypePtrOutputWithContext(ctx context.Context) JobScheduleTypePtrOutput {
	return JobScheduleType(e).ToJobScheduleTypeOutputWithContext(ctx).ToJobScheduleTypePtrOutputWithContext(ctx)
}

func (e JobScheduleType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobScheduleType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobScheduleType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobScheduleType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type JobScheduleTypeOutput struct{ *pulumi.OutputState }

func (JobScheduleTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobScheduleType)(nil)).Elem()
}

func (o JobScheduleTypeOutput) ToJobScheduleTypeOutput() JobScheduleTypeOutput {
	return o
}

func (o JobScheduleTypeOutput) ToJobScheduleTypeOutputWithContext(ctx context.Context) JobScheduleTypeOutput {
	return o
}

func (o JobScheduleTypeOutput) ToJobScheduleTypePtrOutput() JobScheduleTypePtrOutput {
	return o.ToJobScheduleTypePtrOutputWithContext(context.Background())
}

func (o JobScheduleTypeOutput) ToJobScheduleTypePtrOutputWithContext(ctx context.Context) JobScheduleTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobScheduleType) *JobScheduleType {
		return &v
	}).(JobScheduleTypePtrOutput)
}

func (o JobScheduleTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o JobScheduleTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JobScheduleType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o JobScheduleTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JobScheduleTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JobScheduleType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type JobScheduleTypePtrOutput struct{ *pulumi.OutputState }

func (JobScheduleTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobScheduleType)(nil)).Elem()
}

func (o JobScheduleTypePtrOutput) ToJobScheduleTypePtrOutput() JobScheduleTypePtrOutput {
	return o
}

func (o JobScheduleTypePtrOutput) ToJobScheduleTypePtrOutputWithContext(ctx context.Context) JobScheduleTypePtrOutput {
	return o
}

func (o JobScheduleTypePtrOutput) Elem() JobScheduleTypeOutput {
	return o.ApplyT(func(v *JobScheduleType) JobScheduleType {
		if v != nil {
			return *v
		}
		var ret JobScheduleType
		return ret
	}).(JobScheduleTypeOutput)
}

func (o JobScheduleTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JobScheduleTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *JobScheduleType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type JobScheduleTypeInput interface {
	pulumi.Input

	ToJobScheduleTypeOutput() JobScheduleTypeOutput
	ToJobScheduleTypeOutputWithContext(context.Context) JobScheduleTypeOutput
}

var jobScheduleTypePtrType = reflect.TypeOf((**JobScheduleType)(nil)).Elem()

type JobScheduleTypePtrInput interface {
	pulumi.Input

	ToJobScheduleTypePtrOutput() JobScheduleTypePtrOutput
	ToJobScheduleTypePtrOutputWithContext(context.Context) JobScheduleTypePtrOutput
}

type jobScheduleTypePtr string

func JobScheduleTypePtr(v string) JobScheduleTypePtrInput {
	return (*jobScheduleTypePtr)(&v)
}

func (*jobScheduleTypePtr) ElementType() reflect.Type {
	return jobScheduleTypePtrType
}

func (in *jobScheduleTypePtr) ToJobScheduleTypePtrOutput() JobScheduleTypePtrOutput {
	return pulumi.ToOutput(in).(JobScheduleTypePtrOutput)
}

func (in *jobScheduleTypePtr) ToJobScheduleTypePtrOutputWithContext(ctx context.Context) JobScheduleTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(JobScheduleTypePtrOutput)
}

type JobStepActionSource string

const (
	JobStepActionSourceInline = JobStepActionSource("Inline")
)

func (JobStepActionSource) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepActionSource)(nil)).Elem()
}

func (e JobStepActionSource) ToJobStepActionSourceOutput() JobStepActionSourceOutput {
	return pulumi.ToOutput(e).(JobStepActionSourceOutput)
}

func (e JobStepActionSource) ToJobStepActionSourceOutputWithContext(ctx context.Context) JobStepActionSourceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(JobStepActionSourceOutput)
}

func (e JobStepActionSource) ToJobStepActionSourcePtrOutput() JobStepActionSourcePtrOutput {
	return e.ToJobStepActionSourcePtrOutputWithContext(context.Background())
}

func (e JobStepActionSource) ToJobStepActionSourcePtrOutputWithContext(ctx context.Context) JobStepActionSourcePtrOutput {
	return JobStepActionSource(e).ToJobStepActionSourceOutputWithContext(ctx).ToJobStepActionSourcePtrOutputWithContext(ctx)
}

func (e JobStepActionSource) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobStepActionSource) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobStepActionSource) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobStepActionSource) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type JobStepActionSourceOutput struct{ *pulumi.OutputState }

func (JobStepActionSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepActionSource)(nil)).Elem()
}

func (o JobStepActionSourceOutput) ToJobStepActionSourceOutput() JobStepActionSourceOutput {
	return o
}

func (o JobStepActionSourceOutput) ToJobStepActionSourceOutputWithContext(ctx context.Context) JobStepActionSourceOutput {
	return o
}

func (o JobStepActionSourceOutput) ToJobStepActionSourcePtrOutput() JobStepActionSourcePtrOutput {
	return o.ToJobStepActionSourcePtrOutputWithContext(context.Background())
}

func (o JobStepActionSourceOutput) ToJobStepActionSourcePtrOutputWithContext(ctx context.Context) JobStepActionSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStepActionSource) *JobStepActionSource {
		return &v
	}).(JobStepActionSourcePtrOutput)
}

func (o JobStepActionSourceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o JobStepActionSourceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JobStepActionSource) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o JobStepActionSourceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JobStepActionSourceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JobStepActionSource) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type JobStepActionSourcePtrOutput struct{ *pulumi.OutputState }

func (JobStepActionSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepActionSource)(nil)).Elem()
}

func (o JobStepActionSourcePtrOutput) ToJobStepActionSourcePtrOutput() JobStepActionSourcePtrOutput {
	return o
}

func (o JobStepActionSourcePtrOutput) ToJobStepActionSourcePtrOutputWithContext(ctx context.Context) JobStepActionSourcePtrOutput {
	return o
}

func (o JobStepActionSourcePtrOutput) Elem() JobStepActionSourceOutput {
	return o.ApplyT(func(v *JobStepActionSource) JobStepActionSource {
		if v != nil {
			return *v
		}
		var ret JobStepActionSource
		return ret
	}).(JobStepActionSourceOutput)
}

func (o JobStepActionSourcePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JobStepActionSourcePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *JobStepActionSource) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type JobStepActionSourceInput interface {
	pulumi.Input

	ToJobStepActionSourceOutput() JobStepActionSourceOutput
	ToJobStepActionSourceOutputWithContext(context.Context) JobStepActionSourceOutput
}

var jobStepActionSourcePtrType = reflect.TypeOf((**JobStepActionSource)(nil)).Elem()

type JobStepActionSourcePtrInput interface {
	pulumi.Input

	ToJobStepActionSourcePtrOutput() JobStepActionSourcePtrOutput
	ToJobStepActionSourcePtrOutputWithContext(context.Context) JobStepActionSourcePtrOutput
}

type jobStepActionSourcePtr string

func JobStepActionSourcePtr(v string) JobStepActionSourcePtrInput {
	return (*jobStepActionSourcePtr)(&v)
}

func (*jobStepActionSourcePtr) ElementType() reflect.Type {
	return jobStepActionSourcePtrType
}

func (in *jobStepActionSourcePtr) ToJobStepActionSourcePtrOutput() JobStepActionSourcePtrOutput {
	return pulumi.ToOutput(in).(JobStepActionSourcePtrOutput)
}

func (in *jobStepActionSourcePtr) ToJobStepActionSourcePtrOutputWithContext(ctx context.Context) JobStepActionSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(JobStepActionSourcePtrOutput)
}

type JobStepActionType string

const (
	JobStepActionTypeTSql = JobStepActionType("TSql")
)

func (JobStepActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepActionType)(nil)).Elem()
}

func (e JobStepActionType) ToJobStepActionTypeOutput() JobStepActionTypeOutput {
	return pulumi.ToOutput(e).(JobStepActionTypeOutput)
}

func (e JobStepActionType) ToJobStepActionTypeOutputWithContext(ctx context.Context) JobStepActionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(JobStepActionTypeOutput)
}

func (e JobStepActionType) ToJobStepActionTypePtrOutput() JobStepActionTypePtrOutput {
	return e.ToJobStepActionTypePtrOutputWithContext(context.Background())
}

func (e JobStepActionType) ToJobStepActionTypePtrOutputWithContext(ctx context.Context) JobStepActionTypePtrOutput {
	return JobStepActionType(e).ToJobStepActionTypeOutputWithContext(ctx).ToJobStepActionTypePtrOutputWithContext(ctx)
}

func (e JobStepActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobStepActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobStepActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobStepActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type JobStepActionTypeOutput struct{ *pulumi.OutputState }

func (JobStepActionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepActionType)(nil)).Elem()
}

func (o JobStepActionTypeOutput) ToJobStepActionTypeOutput() JobStepActionTypeOutput {
	return o
}

func (o JobStepActionTypeOutput) ToJobStepActionTypeOutputWithContext(ctx context.Context) JobStepActionTypeOutput {
	return o
}

func (o JobStepActionTypeOutput) ToJobStepActionTypePtrOutput() JobStepActionTypePtrOutput {
	return o.ToJobStepActionTypePtrOutputWithContext(context.Background())
}

func (o JobStepActionTypeOutput) ToJobStepActionTypePtrOutputWithContext(ctx context.Context) JobStepActionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStepActionType) *JobStepActionType {
		return &v
	}).(JobStepActionTypePtrOutput)
}

func (o JobStepActionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o JobStepActionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JobStepActionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o JobStepActionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JobStepActionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JobStepActionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type JobStepActionTypePtrOutput struct{ *pulumi.OutputState }

func (JobStepActionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepActionType)(nil)).Elem()
}

func (o JobStepActionTypePtrOutput) ToJobStepActionTypePtrOutput() JobStepActionTypePtrOutput {
	return o
}

func (o JobStepActionTypePtrOutput) ToJobStepActionTypePtrOutputWithContext(ctx context.Context) JobStepActionTypePtrOutput {
	return o
}

func (o JobStepActionTypePtrOutput) Elem() JobStepActionTypeOutput {
	return o.ApplyT(func(v *JobStepActionType) JobStepActionType {
		if v != nil {
			return *v
		}
		var ret JobStepActionType
		return ret
	}).(JobStepActionTypeOutput)
}

func (o JobStepActionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JobStepActionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *JobStepActionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type JobStepActionTypeInput interface {
	pulumi.Input

	ToJobStepActionTypeOutput() JobStepActionTypeOutput
	ToJobStepActionTypeOutputWithContext(context.Context) JobStepActionTypeOutput
}

var jobStepActionTypePtrType = reflect.TypeOf((**JobStepActionType)(nil)).Elem()

type JobStepActionTypePtrInput interface {
	pulumi.Input

	ToJobStepActionTypePtrOutput() JobStepActionTypePtrOutput
	ToJobStepActionTypePtrOutputWithContext(context.Context) JobStepActionTypePtrOutput
}

type jobStepActionTypePtr string

func JobStepActionTypePtr(v string) JobStepActionTypePtrInput {
	return (*jobStepActionTypePtr)(&v)
}

func (*jobStepActionTypePtr) ElementType() reflect.Type {
	return jobStepActionTypePtrType
}

func (in *jobStepActionTypePtr) ToJobStepActionTypePtrOutput() JobStepActionTypePtrOutput {
	return pulumi.ToOutput(in).(JobStepActionTypePtrOutput)
}

func (in *jobStepActionTypePtr) ToJobStepActionTypePtrOutputWithContext(ctx context.Context) JobStepActionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(JobStepActionTypePtrOutput)
}

type JobStepOutputTypeEnum string

const (
	JobStepOutputTypeEnumSqlDatabase = JobStepOutputTypeEnum("SqlDatabase")
)

func (JobStepOutputTypeEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepOutputTypeEnum)(nil)).Elem()
}

func (e JobStepOutputTypeEnum) ToJobStepOutputTypeEnumOutput() JobStepOutputTypeEnumOutput {
	return pulumi.ToOutput(e).(JobStepOutputTypeEnumOutput)
}

func (e JobStepOutputTypeEnum) ToJobStepOutputTypeEnumOutputWithContext(ctx context.Context) JobStepOutputTypeEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(JobStepOutputTypeEnumOutput)
}

func (e JobStepOutputTypeEnum) ToJobStepOutputTypeEnumPtrOutput() JobStepOutputTypeEnumPtrOutput {
	return e.ToJobStepOutputTypeEnumPtrOutputWithContext(context.Background())
}

func (e JobStepOutputTypeEnum) ToJobStepOutputTypeEnumPtrOutputWithContext(ctx context.Context) JobStepOutputTypeEnumPtrOutput {
	return JobStepOutputTypeEnum(e).ToJobStepOutputTypeEnumOutputWithContext(ctx).ToJobStepOutputTypeEnumPtrOutputWithContext(ctx)
}

func (e JobStepOutputTypeEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobStepOutputTypeEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobStepOutputTypeEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobStepOutputTypeEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type JobStepOutputTypeEnumOutput struct{ *pulumi.OutputState }

func (JobStepOutputTypeEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepOutputTypeEnum)(nil)).Elem()
}

func (o JobStepOutputTypeEnumOutput) ToJobStepOutputTypeEnumOutput() JobStepOutputTypeEnumOutput {
	return o
}

func (o JobStepOutputTypeEnumOutput) ToJobStepOutputTypeEnumOutputWithContext(ctx context.Context) JobStepOutputTypeEnumOutput {
	return o
}

func (o JobStepOutputTypeEnumOutput) ToJobStepOutputTypeEnumPtrOutput() JobStepOutputTypeEnumPtrOutput {
	return o.ToJobStepOutputTypeEnumPtrOutputWithContext(context.Background())
}

func (o JobStepOutputTypeEnumOutput) ToJobStepOutputTypeEnumPtrOutputWithContext(ctx context.Context) JobStepOutputTypeEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStepOutputTypeEnum) *JobStepOutputTypeEnum {
		return &v
	}).(JobStepOutputTypeEnumPtrOutput)
}

func (o JobStepOutputTypeEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o JobStepOutputTypeEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JobStepOutputTypeEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o JobStepOutputTypeEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JobStepOutputTypeEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JobStepOutputTypeEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type JobStepOutputTypeEnumPtrOutput struct{ *pulumi.OutputState }

func (JobStepOutputTypeEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepOutputTypeEnum)(nil)).Elem()
}

func (o JobStepOutputTypeEnumPtrOutput) ToJobStepOutputTypeEnumPtrOutput() JobStepOutputTypeEnumPtrOutput {
	return o
}

func (o JobStepOutputTypeEnumPtrOutput) ToJobStepOutputTypeEnumPtrOutputWithContext(ctx context.Context) JobStepOutputTypeEnumPtrOutput {
	return o
}

func (o JobStepOutputTypeEnumPtrOutput) Elem() JobStepOutputTypeEnumOutput {
	return o.ApplyT(func(v *JobStepOutputTypeEnum) JobStepOutputTypeEnum {
		if v != nil {
			return *v
		}
		var ret JobStepOutputTypeEnum
		return ret
	}).(JobStepOutputTypeEnumOutput)
}

func (o JobStepOutputTypeEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JobStepOutputTypeEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *JobStepOutputTypeEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type JobStepOutputTypeEnumInput interface {
	pulumi.Input

	ToJobStepOutputTypeEnumOutput() JobStepOutputTypeEnumOutput
	ToJobStepOutputTypeEnumOutputWithContext(context.Context) JobStepOutputTypeEnumOutput
}

var jobStepOutputTypeEnumPtrType = reflect.TypeOf((**JobStepOutputTypeEnum)(nil)).Elem()

type JobStepOutputTypeEnumPtrInput interface {
	pulumi.Input

	ToJobStepOutputTypeEnumPtrOutput() JobStepOutputTypeEnumPtrOutput
	ToJobStepOutputTypeEnumPtrOutputWithContext(context.Context) JobStepOutputTypeEnumPtrOutput
}

type jobStepOutputTypeEnumPtr string

func JobStepOutputTypeEnumPtr(v string) JobStepOutputTypeEnumPtrInput {
	return (*jobStepOutputTypeEnumPtr)(&v)
}

func (*jobStepOutputTypeEnumPtr) ElementType() reflect.Type {
	return jobStepOutputTypeEnumPtrType
}

func (in *jobStepOutputTypeEnumPtr) ToJobStepOutputTypeEnumPtrOutput() JobStepOutputTypeEnumPtrOutput {
	return pulumi.ToOutput(in).(JobStepOutputTypeEnumPtrOutput)
}

func (in *jobStepOutputTypeEnumPtr) ToJobStepOutputTypeEnumPtrOutputWithContext(ctx context.Context) JobStepOutputTypeEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(JobStepOutputTypeEnumPtrOutput)
}

type JobTargetGroupMembershipType string

const (
	JobTargetGroupMembershipTypeInclude = JobTargetGroupMembershipType("Include")
	JobTargetGroupMembershipTypeExclude = JobTargetGroupMembershipType("Exclude")
)

func (JobTargetGroupMembershipType) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTargetGroupMembershipType)(nil)).Elem()
}

func (e JobTargetGroupMembershipType) ToJobTargetGroupMembershipTypeOutput() JobTargetGroupMembershipTypeOutput {
	return pulumi.ToOutput(e).(JobTargetGroupMembershipTypeOutput)
}

func (e JobTargetGroupMembershipType) ToJobTargetGroupMembershipTypeOutputWithContext(ctx context.Context) JobTargetGroupMembershipTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(JobTargetGroupMembershipTypeOutput)
}

func (e JobTargetGroupMembershipType) ToJobTargetGroupMembershipTypePtrOutput() JobTargetGroupMembershipTypePtrOutput {
	return e.ToJobTargetGroupMembershipTypePtrOutputWithContext(context.Background())
}

func (e JobTargetGroupMembershipType) ToJobTargetGroupMembershipTypePtrOutputWithContext(ctx context.Context) JobTargetGroupMembershipTypePtrOutput {
	return JobTargetGroupMembershipType(e).ToJobTargetGroupMembershipTypeOutputWithContext(ctx).ToJobTargetGroupMembershipTypePtrOutputWithContext(ctx)
}

func (e JobTargetGroupMembershipType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobTargetGroupMembershipType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobTargetGroupMembershipType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobTargetGroupMembershipType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type JobTargetGroupMembershipTypeOutput struct{ *pulumi.OutputState }

func (JobTargetGroupMembershipTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTargetGroupMembershipType)(nil)).Elem()
}

func (o JobTargetGroupMembershipTypeOutput) ToJobTargetGroupMembershipTypeOutput() JobTargetGroupMembershipTypeOutput {
	return o
}

func (o JobTargetGroupMembershipTypeOutput) ToJobTargetGroupMembershipTypeOutputWithContext(ctx context.Context) JobTargetGroupMembershipTypeOutput {
	return o
}

func (o JobTargetGroupMembershipTypeOutput) ToJobTargetGroupMembershipTypePtrOutput() JobTargetGroupMembershipTypePtrOutput {
	return o.ToJobTargetGroupMembershipTypePtrOutputWithContext(context.Background())
}

func (o JobTargetGroupMembershipTypeOutput) ToJobTargetGroupMembershipTypePtrOutputWithContext(ctx context.Context) JobTargetGroupMembershipTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobTargetGroupMembershipType) *JobTargetGroupMembershipType {
		return &v
	}).(JobTargetGroupMembershipTypePtrOutput)
}

func (o JobTargetGroupMembershipTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o JobTargetGroupMembershipTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JobTargetGroupMembershipType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o JobTargetGroupMembershipTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JobTargetGroupMembershipTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JobTargetGroupMembershipType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type JobTargetGroupMembershipTypePtrOutput struct{ *pulumi.OutputState }

func (JobTargetGroupMembershipTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobTargetGroupMembershipType)(nil)).Elem()
}

func (o JobTargetGroupMembershipTypePtrOutput) ToJobTargetGroupMembershipTypePtrOutput() JobTargetGroupMembershipTypePtrOutput {
	return o
}

func (o JobTargetGroupMembershipTypePtrOutput) ToJobTargetGroupMembershipTypePtrOutputWithContext(ctx context.Context) JobTargetGroupMembershipTypePtrOutput {
	return o
}

func (o JobTargetGroupMembershipTypePtrOutput) Elem() JobTargetGroupMembershipTypeOutput {
	return o.ApplyT(func(v *JobTargetGroupMembershipType) JobTargetGroupMembershipType {
		if v != nil {
			return *v
		}
		var ret JobTargetGroupMembershipType
		return ret
	}).(JobTargetGroupMembershipTypeOutput)
}

func (o JobTargetGroupMembershipTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JobTargetGroupMembershipTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *JobTargetGroupMembershipType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type JobTargetGroupMembershipTypeInput interface {
	pulumi.Input

	ToJobTargetGroupMembershipTypeOutput() JobTargetGroupMembershipTypeOutput
	ToJobTargetGroupMembershipTypeOutputWithContext(context.Context) JobTargetGroupMembershipTypeOutput
}

var jobTargetGroupMembershipTypePtrType = reflect.TypeOf((**JobTargetGroupMembershipType)(nil)).Elem()

type JobTargetGroupMembershipTypePtrInput interface {
	pulumi.Input

	ToJobTargetGroupMembershipTypePtrOutput() JobTargetGroupMembershipTypePtrOutput
	ToJobTargetGroupMembershipTypePtrOutputWithContext(context.Context) JobTargetGroupMembershipTypePtrOutput
}

type jobTargetGroupMembershipTypePtr string

func JobTargetGroupMembershipTypePtr(v string) JobTargetGroupMembershipTypePtrInput {
	return (*jobTargetGroupMembershipTypePtr)(&v)
}

func (*jobTargetGroupMembershipTypePtr) ElementType() reflect.Type {
	return jobTargetGroupMembershipTypePtrType
}

func (in *jobTargetGroupMembershipTypePtr) ToJobTargetGroupMembershipTypePtrOutput() JobTargetGroupMembershipTypePtrOutput {
	return pulumi.ToOutput(in).(JobTargetGroupMembershipTypePtrOutput)
}

func (in *jobTargetGroupMembershipTypePtr) ToJobTargetGroupMembershipTypePtrOutputWithContext(ctx context.Context) JobTargetGroupMembershipTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(JobTargetGroupMembershipTypePtrOutput)
}

type JobTargetType string

const (
	JobTargetTypeTargetGroup    = JobTargetType("TargetGroup")
	JobTargetTypeSqlDatabase    = JobTargetType("SqlDatabase")
	JobTargetTypeSqlElasticPool = JobTargetType("SqlElasticPool")
	JobTargetTypeSqlShardMap    = JobTargetType("SqlShardMap")
	JobTargetTypeSqlServer      = JobTargetType("SqlServer")
)

func (JobTargetType) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTargetType)(nil)).Elem()
}

func (e JobTargetType) ToJobTargetTypeOutput() JobTargetTypeOutput {
	return pulumi.ToOutput(e).(JobTargetTypeOutput)
}

func (e JobTargetType) ToJobTargetTypeOutputWithContext(ctx context.Context) JobTargetTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(JobTargetTypeOutput)
}

func (e JobTargetType) ToJobTargetTypePtrOutput() JobTargetTypePtrOutput {
	return e.ToJobTargetTypePtrOutputWithContext(context.Background())
}

func (e JobTargetType) ToJobTargetTypePtrOutputWithContext(ctx context.Context) JobTargetTypePtrOutput {
	return JobTargetType(e).ToJobTargetTypeOutputWithContext(ctx).ToJobTargetTypePtrOutputWithContext(ctx)
}

func (e JobTargetType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobTargetType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobTargetType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobTargetType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type JobTargetTypeOutput struct{ *pulumi.OutputState }

func (JobTargetTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTargetType)(nil)).Elem()
}

func (o JobTargetTypeOutput) ToJobTargetTypeOutput() JobTargetTypeOutput {
	return o
}

func (o JobTargetTypeOutput) ToJobTargetTypeOutputWithContext(ctx context.Context) JobTargetTypeOutput {
	return o
}

func (o JobTargetTypeOutput) ToJobTargetTypePtrOutput() JobTargetTypePtrOutput {
	return o.ToJobTargetTypePtrOutputWithContext(context.Background())
}

func (o JobTargetTypeOutput) ToJobTargetTypePtrOutputWithContext(ctx context.Context) JobTargetTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobTargetType) *JobTargetType {
		return &v
	}).(JobTargetTypePtrOutput)
}

func (o JobTargetTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o JobTargetTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JobTargetType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o JobTargetTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JobTargetTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JobTargetType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type JobTargetTypePtrOutput struct{ *pulumi.OutputState }

func (JobTargetTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobTargetType)(nil)).Elem()
}

func (o JobTargetTypePtrOutput) ToJobTargetTypePtrOutput() JobTargetTypePtrOutput {
	return o
}

func (o JobTargetTypePtrOutput) ToJobTargetTypePtrOutputWithContext(ctx context.Context) JobTargetTypePtrOutput {
	return o
}

func (o JobTargetTypePtrOutput) Elem() JobTargetTypeOutput {
	return o.ApplyT(func(v *JobTargetType) JobTargetType {
		if v != nil {
			return *v
		}
		var ret JobTargetType
		return ret
	}).(JobTargetTypeOutput)
}

func (o JobTargetTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JobTargetTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *JobTargetType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type JobTargetTypeInput interface {
	pulumi.Input

	ToJobTargetTypeOutput() JobTargetTypeOutput
	ToJobTargetTypeOutputWithContext(context.Context) JobTargetTypeOutput
}

var jobTargetTypePtrType = reflect.TypeOf((**JobTargetType)(nil)).Elem()

type JobTargetTypePtrInput interface {
	pulumi.Input

	ToJobTargetTypePtrOutput() JobTargetTypePtrOutput
	ToJobTargetTypePtrOutputWithContext(context.Context) JobTargetTypePtrOutput
}

type jobTargetTypePtr string

func JobTargetTypePtr(v string) JobTargetTypePtrInput {
	return (*jobTargetTypePtr)(&v)
}

func (*jobTargetTypePtr) ElementType() reflect.Type {
	return jobTargetTypePtrType
}

func (in *jobTargetTypePtr) ToJobTargetTypePtrOutput() JobTargetTypePtrOutput {
	return pulumi.ToOutput(in).(JobTargetTypePtrOutput)
}

func (in *jobTargetTypePtr) ToJobTargetTypePtrOutputWithContext(ctx context.Context) JobTargetTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(JobTargetTypePtrOutput)
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

type ManagedInstanceAdministratorType string

const (
	ManagedInstanceAdministratorTypeActiveDirectory = ManagedInstanceAdministratorType("ActiveDirectory")
)

func (ManagedInstanceAdministratorType) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstanceAdministratorType)(nil)).Elem()
}

func (e ManagedInstanceAdministratorType) ToManagedInstanceAdministratorTypeOutput() ManagedInstanceAdministratorTypeOutput {
	return pulumi.ToOutput(e).(ManagedInstanceAdministratorTypeOutput)
}

func (e ManagedInstanceAdministratorType) ToManagedInstanceAdministratorTypeOutputWithContext(ctx context.Context) ManagedInstanceAdministratorTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedInstanceAdministratorTypeOutput)
}

func (e ManagedInstanceAdministratorType) ToManagedInstanceAdministratorTypePtrOutput() ManagedInstanceAdministratorTypePtrOutput {
	return e.ToManagedInstanceAdministratorTypePtrOutputWithContext(context.Background())
}

func (e ManagedInstanceAdministratorType) ToManagedInstanceAdministratorTypePtrOutputWithContext(ctx context.Context) ManagedInstanceAdministratorTypePtrOutput {
	return ManagedInstanceAdministratorType(e).ToManagedInstanceAdministratorTypeOutputWithContext(ctx).ToManagedInstanceAdministratorTypePtrOutputWithContext(ctx)
}

func (e ManagedInstanceAdministratorType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedInstanceAdministratorType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedInstanceAdministratorType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedInstanceAdministratorType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedInstanceAdministratorTypeOutput struct{ *pulumi.OutputState }

func (ManagedInstanceAdministratorTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedInstanceAdministratorType)(nil)).Elem()
}

func (o ManagedInstanceAdministratorTypeOutput) ToManagedInstanceAdministratorTypeOutput() ManagedInstanceAdministratorTypeOutput {
	return o
}

func (o ManagedInstanceAdministratorTypeOutput) ToManagedInstanceAdministratorTypeOutputWithContext(ctx context.Context) ManagedInstanceAdministratorTypeOutput {
	return o
}

func (o ManagedInstanceAdministratorTypeOutput) ToManagedInstanceAdministratorTypePtrOutput() ManagedInstanceAdministratorTypePtrOutput {
	return o.ToManagedInstanceAdministratorTypePtrOutputWithContext(context.Background())
}

func (o ManagedInstanceAdministratorTypeOutput) ToManagedInstanceAdministratorTypePtrOutputWithContext(ctx context.Context) ManagedInstanceAdministratorTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedInstanceAdministratorType) *ManagedInstanceAdministratorType {
		return &v
	}).(ManagedInstanceAdministratorTypePtrOutput)
}

func (o ManagedInstanceAdministratorTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedInstanceAdministratorTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedInstanceAdministratorType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedInstanceAdministratorTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedInstanceAdministratorTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedInstanceAdministratorType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedInstanceAdministratorTypePtrOutput struct{ *pulumi.OutputState }

func (ManagedInstanceAdministratorTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceAdministratorType)(nil)).Elem()
}

func (o ManagedInstanceAdministratorTypePtrOutput) ToManagedInstanceAdministratorTypePtrOutput() ManagedInstanceAdministratorTypePtrOutput {
	return o
}

func (o ManagedInstanceAdministratorTypePtrOutput) ToManagedInstanceAdministratorTypePtrOutputWithContext(ctx context.Context) ManagedInstanceAdministratorTypePtrOutput {
	return o
}

func (o ManagedInstanceAdministratorTypePtrOutput) Elem() ManagedInstanceAdministratorTypeOutput {
	return o.ApplyT(func(v *ManagedInstanceAdministratorType) ManagedInstanceAdministratorType {
		if v != nil {
			return *v
		}
		var ret ManagedInstanceAdministratorType
		return ret
	}).(ManagedInstanceAdministratorTypeOutput)
}

func (o ManagedInstanceAdministratorTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedInstanceAdministratorTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedInstanceAdministratorType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ManagedInstanceAdministratorTypeInput interface {
	pulumi.Input

	ToManagedInstanceAdministratorTypeOutput() ManagedInstanceAdministratorTypeOutput
	ToManagedInstanceAdministratorTypeOutputWithContext(context.Context) ManagedInstanceAdministratorTypeOutput
}

var managedInstanceAdministratorTypePtrType = reflect.TypeOf((**ManagedInstanceAdministratorType)(nil)).Elem()

type ManagedInstanceAdministratorTypePtrInput interface {
	pulumi.Input

	ToManagedInstanceAdministratorTypePtrOutput() ManagedInstanceAdministratorTypePtrOutput
	ToManagedInstanceAdministratorTypePtrOutputWithContext(context.Context) ManagedInstanceAdministratorTypePtrOutput
}

type managedInstanceAdministratorTypePtr string

func ManagedInstanceAdministratorTypePtr(v string) ManagedInstanceAdministratorTypePtrInput {
	return (*managedInstanceAdministratorTypePtr)(&v)
}

func (*managedInstanceAdministratorTypePtr) ElementType() reflect.Type {
	return managedInstanceAdministratorTypePtrType
}

func (in *managedInstanceAdministratorTypePtr) ToManagedInstanceAdministratorTypePtrOutput() ManagedInstanceAdministratorTypePtrOutput {
	return pulumi.ToOutput(in).(ManagedInstanceAdministratorTypePtrOutput)
}

func (in *managedInstanceAdministratorTypePtr) ToManagedInstanceAdministratorTypePtrOutputWithContext(ctx context.Context) ManagedInstanceAdministratorTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedInstanceAdministratorTypePtrOutput)
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

type PrincipalType string

const (
	PrincipalTypeUser        = PrincipalType("User")
	PrincipalTypeGroup       = PrincipalType("Group")
	PrincipalTypeApplication = PrincipalType("Application")
)

func (PrincipalType) ElementType() reflect.Type {
	return reflect.TypeOf((*PrincipalType)(nil)).Elem()
}

func (e PrincipalType) ToPrincipalTypeOutput() PrincipalTypeOutput {
	return pulumi.ToOutput(e).(PrincipalTypeOutput)
}

func (e PrincipalType) ToPrincipalTypeOutputWithContext(ctx context.Context) PrincipalTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrincipalTypeOutput)
}

func (e PrincipalType) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return e.ToPrincipalTypePtrOutputWithContext(context.Background())
}

func (e PrincipalType) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return PrincipalType(e).ToPrincipalTypeOutputWithContext(ctx).ToPrincipalTypePtrOutputWithContext(ctx)
}

func (e PrincipalType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrincipalType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrincipalType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrincipalType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrincipalTypeOutput struct{ *pulumi.OutputState }

func (PrincipalTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrincipalType)(nil)).Elem()
}

func (o PrincipalTypeOutput) ToPrincipalTypeOutput() PrincipalTypeOutput {
	return o
}

func (o PrincipalTypeOutput) ToPrincipalTypeOutputWithContext(ctx context.Context) PrincipalTypeOutput {
	return o
}

func (o PrincipalTypeOutput) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return o.ToPrincipalTypePtrOutputWithContext(context.Background())
}

func (o PrincipalTypeOutput) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrincipalType) *PrincipalType {
		return &v
	}).(PrincipalTypePtrOutput)
}

func (o PrincipalTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrincipalTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrincipalType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrincipalTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrincipalTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrincipalType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrincipalTypePtrOutput struct{ *pulumi.OutputState }

func (PrincipalTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrincipalType)(nil)).Elem()
}

func (o PrincipalTypePtrOutput) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return o
}

func (o PrincipalTypePtrOutput) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return o
}

func (o PrincipalTypePtrOutput) Elem() PrincipalTypeOutput {
	return o.ApplyT(func(v *PrincipalType) PrincipalType {
		if v != nil {
			return *v
		}
		var ret PrincipalType
		return ret
	}).(PrincipalTypeOutput)
}

func (o PrincipalTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrincipalTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrincipalType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrincipalTypeInput interface {
	pulumi.Input

	ToPrincipalTypeOutput() PrincipalTypeOutput
	ToPrincipalTypeOutputWithContext(context.Context) PrincipalTypeOutput
}

var principalTypePtrType = reflect.TypeOf((**PrincipalType)(nil)).Elem()

type PrincipalTypePtrInput interface {
	pulumi.Input

	ToPrincipalTypePtrOutput() PrincipalTypePtrOutput
	ToPrincipalTypePtrOutputWithContext(context.Context) PrincipalTypePtrOutput
}

type principalTypePtr string

func PrincipalTypePtr(v string) PrincipalTypePtrInput {
	return (*principalTypePtr)(&v)
}

func (*principalTypePtr) ElementType() reflect.Type {
	return principalTypePtrType
}

func (in *principalTypePtr) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return pulumi.ToOutput(in).(PrincipalTypePtrOutput)
}

func (in *principalTypePtr) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrincipalTypePtrOutput)
}

type PrivateLinkServiceConnectionStateStatus string

const (
	PrivateLinkServiceConnectionStateStatusApproved     = PrivateLinkServiceConnectionStateStatus("Approved")
	PrivateLinkServiceConnectionStateStatusPending      = PrivateLinkServiceConnectionStateStatus("Pending")
	PrivateLinkServiceConnectionStateStatusRejected     = PrivateLinkServiceConnectionStateStatus("Rejected")
	PrivateLinkServiceConnectionStateStatusDisconnected = PrivateLinkServiceConnectionStateStatus("Disconnected")
)

func (PrivateLinkServiceConnectionStateStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateStatus)(nil)).Elem()
}

func (e PrivateLinkServiceConnectionStateStatus) ToPrivateLinkServiceConnectionStateStatusOutput() PrivateLinkServiceConnectionStateStatusOutput {
	return pulumi.ToOutput(e).(PrivateLinkServiceConnectionStateStatusOutput)
}

func (e PrivateLinkServiceConnectionStateStatus) ToPrivateLinkServiceConnectionStateStatusOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrivateLinkServiceConnectionStateStatusOutput)
}

func (e PrivateLinkServiceConnectionStateStatus) ToPrivateLinkServiceConnectionStateStatusPtrOutput() PrivateLinkServiceConnectionStateStatusPtrOutput {
	return e.ToPrivateLinkServiceConnectionStateStatusPtrOutputWithContext(context.Background())
}

func (e PrivateLinkServiceConnectionStateStatus) ToPrivateLinkServiceConnectionStateStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateStatusPtrOutput {
	return PrivateLinkServiceConnectionStateStatus(e).ToPrivateLinkServiceConnectionStateStatusOutputWithContext(ctx).ToPrivateLinkServiceConnectionStateStatusPtrOutputWithContext(ctx)
}

func (e PrivateLinkServiceConnectionStateStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateLinkServiceConnectionStateStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateLinkServiceConnectionStateStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateLinkServiceConnectionStateStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrivateLinkServiceConnectionStateStatusOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateStatus)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateStatusOutput) ToPrivateLinkServiceConnectionStateStatusOutput() PrivateLinkServiceConnectionStateStatusOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateStatusOutput) ToPrivateLinkServiceConnectionStateStatusOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateStatusOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateStatusOutput) ToPrivateLinkServiceConnectionStateStatusPtrOutput() PrivateLinkServiceConnectionStateStatusPtrOutput {
	return o.ToPrivateLinkServiceConnectionStateStatusPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateStatusOutput) ToPrivateLinkServiceConnectionStateStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateStatus) *PrivateLinkServiceConnectionStateStatus {
		return &v
	}).(PrivateLinkServiceConnectionStateStatusPtrOutput)
}

func (o PrivateLinkServiceConnectionStateStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateLinkServiceConnectionStateStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStateStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateLinkServiceConnectionStateStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateStatusPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateStatus)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateStatusPtrOutput) ToPrivateLinkServiceConnectionStateStatusPtrOutput() PrivateLinkServiceConnectionStateStatusPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateStatusPtrOutput) ToPrivateLinkServiceConnectionStateStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateStatusPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateStatusPtrOutput) Elem() PrivateLinkServiceConnectionStateStatusOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateStatus) PrivateLinkServiceConnectionStateStatus {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateStatus
		return ret
	}).(PrivateLinkServiceConnectionStateStatusOutput)
}

func (o PrivateLinkServiceConnectionStateStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrivateLinkServiceConnectionStateStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrivateLinkServiceConnectionStateStatusInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateStatusOutput() PrivateLinkServiceConnectionStateStatusOutput
	ToPrivateLinkServiceConnectionStateStatusOutputWithContext(context.Context) PrivateLinkServiceConnectionStateStatusOutput
}

var privateLinkServiceConnectionStateStatusPtrType = reflect.TypeOf((**PrivateLinkServiceConnectionStateStatus)(nil)).Elem()

type PrivateLinkServiceConnectionStateStatusPtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateStatusPtrOutput() PrivateLinkServiceConnectionStateStatusPtrOutput
	ToPrivateLinkServiceConnectionStateStatusPtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStateStatusPtrOutput
}

type privateLinkServiceConnectionStateStatusPtr string

func PrivateLinkServiceConnectionStateStatusPtr(v string) PrivateLinkServiceConnectionStateStatusPtrInput {
	return (*privateLinkServiceConnectionStateStatusPtr)(&v)
}

func (*privateLinkServiceConnectionStateStatusPtr) ElementType() reflect.Type {
	return privateLinkServiceConnectionStateStatusPtrType
}

func (in *privateLinkServiceConnectionStateStatusPtr) ToPrivateLinkServiceConnectionStateStatusPtrOutput() PrivateLinkServiceConnectionStateStatusPtrOutput {
	return pulumi.ToOutput(in).(PrivateLinkServiceConnectionStateStatusPtrOutput)
}

func (in *privateLinkServiceConnectionStateStatusPtr) ToPrivateLinkServiceConnectionStateStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrivateLinkServiceConnectionStateStatusPtrOutput)
}

type ReadOnlyEndpointFailoverPolicy string

const (
	ReadOnlyEndpointFailoverPolicyDisabled = ReadOnlyEndpointFailoverPolicy("Disabled")
	ReadOnlyEndpointFailoverPolicyEnabled  = ReadOnlyEndpointFailoverPolicy("Enabled")
)

func (ReadOnlyEndpointFailoverPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ReadOnlyEndpointFailoverPolicy)(nil)).Elem()
}

func (e ReadOnlyEndpointFailoverPolicy) ToReadOnlyEndpointFailoverPolicyOutput() ReadOnlyEndpointFailoverPolicyOutput {
	return pulumi.ToOutput(e).(ReadOnlyEndpointFailoverPolicyOutput)
}

func (e ReadOnlyEndpointFailoverPolicy) ToReadOnlyEndpointFailoverPolicyOutputWithContext(ctx context.Context) ReadOnlyEndpointFailoverPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ReadOnlyEndpointFailoverPolicyOutput)
}

func (e ReadOnlyEndpointFailoverPolicy) ToReadOnlyEndpointFailoverPolicyPtrOutput() ReadOnlyEndpointFailoverPolicyPtrOutput {
	return e.ToReadOnlyEndpointFailoverPolicyPtrOutputWithContext(context.Background())
}

func (e ReadOnlyEndpointFailoverPolicy) ToReadOnlyEndpointFailoverPolicyPtrOutputWithContext(ctx context.Context) ReadOnlyEndpointFailoverPolicyPtrOutput {
	return ReadOnlyEndpointFailoverPolicy(e).ToReadOnlyEndpointFailoverPolicyOutputWithContext(ctx).ToReadOnlyEndpointFailoverPolicyPtrOutputWithContext(ctx)
}

func (e ReadOnlyEndpointFailoverPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReadOnlyEndpointFailoverPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReadOnlyEndpointFailoverPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ReadOnlyEndpointFailoverPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ReadOnlyEndpointFailoverPolicyOutput struct{ *pulumi.OutputState }

func (ReadOnlyEndpointFailoverPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReadOnlyEndpointFailoverPolicy)(nil)).Elem()
}

func (o ReadOnlyEndpointFailoverPolicyOutput) ToReadOnlyEndpointFailoverPolicyOutput() ReadOnlyEndpointFailoverPolicyOutput {
	return o
}

func (o ReadOnlyEndpointFailoverPolicyOutput) ToReadOnlyEndpointFailoverPolicyOutputWithContext(ctx context.Context) ReadOnlyEndpointFailoverPolicyOutput {
	return o
}

func (o ReadOnlyEndpointFailoverPolicyOutput) ToReadOnlyEndpointFailoverPolicyPtrOutput() ReadOnlyEndpointFailoverPolicyPtrOutput {
	return o.ToReadOnlyEndpointFailoverPolicyPtrOutputWithContext(context.Background())
}

func (o ReadOnlyEndpointFailoverPolicyOutput) ToReadOnlyEndpointFailoverPolicyPtrOutputWithContext(ctx context.Context) ReadOnlyEndpointFailoverPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReadOnlyEndpointFailoverPolicy) *ReadOnlyEndpointFailoverPolicy {
		return &v
	}).(ReadOnlyEndpointFailoverPolicyPtrOutput)
}

func (o ReadOnlyEndpointFailoverPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ReadOnlyEndpointFailoverPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReadOnlyEndpointFailoverPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ReadOnlyEndpointFailoverPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReadOnlyEndpointFailoverPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReadOnlyEndpointFailoverPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ReadOnlyEndpointFailoverPolicyPtrOutput struct{ *pulumi.OutputState }

func (ReadOnlyEndpointFailoverPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadOnlyEndpointFailoverPolicy)(nil)).Elem()
}

func (o ReadOnlyEndpointFailoverPolicyPtrOutput) ToReadOnlyEndpointFailoverPolicyPtrOutput() ReadOnlyEndpointFailoverPolicyPtrOutput {
	return o
}

func (o ReadOnlyEndpointFailoverPolicyPtrOutput) ToReadOnlyEndpointFailoverPolicyPtrOutputWithContext(ctx context.Context) ReadOnlyEndpointFailoverPolicyPtrOutput {
	return o
}

func (o ReadOnlyEndpointFailoverPolicyPtrOutput) Elem() ReadOnlyEndpointFailoverPolicyOutput {
	return o.ApplyT(func(v *ReadOnlyEndpointFailoverPolicy) ReadOnlyEndpointFailoverPolicy {
		if v != nil {
			return *v
		}
		var ret ReadOnlyEndpointFailoverPolicy
		return ret
	}).(ReadOnlyEndpointFailoverPolicyOutput)
}

func (o ReadOnlyEndpointFailoverPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReadOnlyEndpointFailoverPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ReadOnlyEndpointFailoverPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ReadOnlyEndpointFailoverPolicyInput interface {
	pulumi.Input

	ToReadOnlyEndpointFailoverPolicyOutput() ReadOnlyEndpointFailoverPolicyOutput
	ToReadOnlyEndpointFailoverPolicyOutputWithContext(context.Context) ReadOnlyEndpointFailoverPolicyOutput
}

var readOnlyEndpointFailoverPolicyPtrType = reflect.TypeOf((**ReadOnlyEndpointFailoverPolicy)(nil)).Elem()

type ReadOnlyEndpointFailoverPolicyPtrInput interface {
	pulumi.Input

	ToReadOnlyEndpointFailoverPolicyPtrOutput() ReadOnlyEndpointFailoverPolicyPtrOutput
	ToReadOnlyEndpointFailoverPolicyPtrOutputWithContext(context.Context) ReadOnlyEndpointFailoverPolicyPtrOutput
}

type readOnlyEndpointFailoverPolicyPtr string

func ReadOnlyEndpointFailoverPolicyPtr(v string) ReadOnlyEndpointFailoverPolicyPtrInput {
	return (*readOnlyEndpointFailoverPolicyPtr)(&v)
}

func (*readOnlyEndpointFailoverPolicyPtr) ElementType() reflect.Type {
	return readOnlyEndpointFailoverPolicyPtrType
}

func (in *readOnlyEndpointFailoverPolicyPtr) ToReadOnlyEndpointFailoverPolicyPtrOutput() ReadOnlyEndpointFailoverPolicyPtrOutput {
	return pulumi.ToOutput(in).(ReadOnlyEndpointFailoverPolicyPtrOutput)
}

func (in *readOnlyEndpointFailoverPolicyPtr) ToReadOnlyEndpointFailoverPolicyPtrOutputWithContext(ctx context.Context) ReadOnlyEndpointFailoverPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ReadOnlyEndpointFailoverPolicyPtrOutput)
}

type ReadWriteEndpointFailoverPolicy string

const (
	ReadWriteEndpointFailoverPolicyManual    = ReadWriteEndpointFailoverPolicy("Manual")
	ReadWriteEndpointFailoverPolicyAutomatic = ReadWriteEndpointFailoverPolicy("Automatic")
)

func (ReadWriteEndpointFailoverPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ReadWriteEndpointFailoverPolicy)(nil)).Elem()
}

func (e ReadWriteEndpointFailoverPolicy) ToReadWriteEndpointFailoverPolicyOutput() ReadWriteEndpointFailoverPolicyOutput {
	return pulumi.ToOutput(e).(ReadWriteEndpointFailoverPolicyOutput)
}

func (e ReadWriteEndpointFailoverPolicy) ToReadWriteEndpointFailoverPolicyOutputWithContext(ctx context.Context) ReadWriteEndpointFailoverPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ReadWriteEndpointFailoverPolicyOutput)
}

func (e ReadWriteEndpointFailoverPolicy) ToReadWriteEndpointFailoverPolicyPtrOutput() ReadWriteEndpointFailoverPolicyPtrOutput {
	return e.ToReadWriteEndpointFailoverPolicyPtrOutputWithContext(context.Background())
}

func (e ReadWriteEndpointFailoverPolicy) ToReadWriteEndpointFailoverPolicyPtrOutputWithContext(ctx context.Context) ReadWriteEndpointFailoverPolicyPtrOutput {
	return ReadWriteEndpointFailoverPolicy(e).ToReadWriteEndpointFailoverPolicyOutputWithContext(ctx).ToReadWriteEndpointFailoverPolicyPtrOutputWithContext(ctx)
}

func (e ReadWriteEndpointFailoverPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReadWriteEndpointFailoverPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReadWriteEndpointFailoverPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ReadWriteEndpointFailoverPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ReadWriteEndpointFailoverPolicyOutput struct{ *pulumi.OutputState }

func (ReadWriteEndpointFailoverPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReadWriteEndpointFailoverPolicy)(nil)).Elem()
}

func (o ReadWriteEndpointFailoverPolicyOutput) ToReadWriteEndpointFailoverPolicyOutput() ReadWriteEndpointFailoverPolicyOutput {
	return o
}

func (o ReadWriteEndpointFailoverPolicyOutput) ToReadWriteEndpointFailoverPolicyOutputWithContext(ctx context.Context) ReadWriteEndpointFailoverPolicyOutput {
	return o
}

func (o ReadWriteEndpointFailoverPolicyOutput) ToReadWriteEndpointFailoverPolicyPtrOutput() ReadWriteEndpointFailoverPolicyPtrOutput {
	return o.ToReadWriteEndpointFailoverPolicyPtrOutputWithContext(context.Background())
}

func (o ReadWriteEndpointFailoverPolicyOutput) ToReadWriteEndpointFailoverPolicyPtrOutputWithContext(ctx context.Context) ReadWriteEndpointFailoverPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReadWriteEndpointFailoverPolicy) *ReadWriteEndpointFailoverPolicy {
		return &v
	}).(ReadWriteEndpointFailoverPolicyPtrOutput)
}

func (o ReadWriteEndpointFailoverPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ReadWriteEndpointFailoverPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReadWriteEndpointFailoverPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ReadWriteEndpointFailoverPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReadWriteEndpointFailoverPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReadWriteEndpointFailoverPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ReadWriteEndpointFailoverPolicyPtrOutput struct{ *pulumi.OutputState }

func (ReadWriteEndpointFailoverPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadWriteEndpointFailoverPolicy)(nil)).Elem()
}

func (o ReadWriteEndpointFailoverPolicyPtrOutput) ToReadWriteEndpointFailoverPolicyPtrOutput() ReadWriteEndpointFailoverPolicyPtrOutput {
	return o
}

func (o ReadWriteEndpointFailoverPolicyPtrOutput) ToReadWriteEndpointFailoverPolicyPtrOutputWithContext(ctx context.Context) ReadWriteEndpointFailoverPolicyPtrOutput {
	return o
}

func (o ReadWriteEndpointFailoverPolicyPtrOutput) Elem() ReadWriteEndpointFailoverPolicyOutput {
	return o.ApplyT(func(v *ReadWriteEndpointFailoverPolicy) ReadWriteEndpointFailoverPolicy {
		if v != nil {
			return *v
		}
		var ret ReadWriteEndpointFailoverPolicy
		return ret
	}).(ReadWriteEndpointFailoverPolicyOutput)
}

func (o ReadWriteEndpointFailoverPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReadWriteEndpointFailoverPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ReadWriteEndpointFailoverPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ReadWriteEndpointFailoverPolicyInput interface {
	pulumi.Input

	ToReadWriteEndpointFailoverPolicyOutput() ReadWriteEndpointFailoverPolicyOutput
	ToReadWriteEndpointFailoverPolicyOutputWithContext(context.Context) ReadWriteEndpointFailoverPolicyOutput
}

var readWriteEndpointFailoverPolicyPtrType = reflect.TypeOf((**ReadWriteEndpointFailoverPolicy)(nil)).Elem()

type ReadWriteEndpointFailoverPolicyPtrInput interface {
	pulumi.Input

	ToReadWriteEndpointFailoverPolicyPtrOutput() ReadWriteEndpointFailoverPolicyPtrOutput
	ToReadWriteEndpointFailoverPolicyPtrOutputWithContext(context.Context) ReadWriteEndpointFailoverPolicyPtrOutput
}

type readWriteEndpointFailoverPolicyPtr string

func ReadWriteEndpointFailoverPolicyPtr(v string) ReadWriteEndpointFailoverPolicyPtrInput {
	return (*readWriteEndpointFailoverPolicyPtr)(&v)
}

func (*readWriteEndpointFailoverPolicyPtr) ElementType() reflect.Type {
	return readWriteEndpointFailoverPolicyPtrType
}

func (in *readWriteEndpointFailoverPolicyPtr) ToReadWriteEndpointFailoverPolicyPtrOutput() ReadWriteEndpointFailoverPolicyPtrOutput {
	return pulumi.ToOutput(in).(ReadWriteEndpointFailoverPolicyPtrOutput)
}

func (in *readWriteEndpointFailoverPolicyPtr) ToReadWriteEndpointFailoverPolicyPtrOutputWithContext(ctx context.Context) ReadWriteEndpointFailoverPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ReadWriteEndpointFailoverPolicyPtrOutput)
}

type ReplicationMode string

const (
	ReplicationModeAsync = ReplicationMode("Async")
	ReplicationModeSync  = ReplicationMode("Sync")
)

func (ReplicationMode) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationMode)(nil)).Elem()
}

func (e ReplicationMode) ToReplicationModeOutput() ReplicationModeOutput {
	return pulumi.ToOutput(e).(ReplicationModeOutput)
}

func (e ReplicationMode) ToReplicationModeOutputWithContext(ctx context.Context) ReplicationModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ReplicationModeOutput)
}

func (e ReplicationMode) ToReplicationModePtrOutput() ReplicationModePtrOutput {
	return e.ToReplicationModePtrOutputWithContext(context.Background())
}

func (e ReplicationMode) ToReplicationModePtrOutputWithContext(ctx context.Context) ReplicationModePtrOutput {
	return ReplicationMode(e).ToReplicationModeOutputWithContext(ctx).ToReplicationModePtrOutputWithContext(ctx)
}

func (e ReplicationMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReplicationMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReplicationMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ReplicationMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ReplicationModeOutput struct{ *pulumi.OutputState }

func (ReplicationModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationMode)(nil)).Elem()
}

func (o ReplicationModeOutput) ToReplicationModeOutput() ReplicationModeOutput {
	return o
}

func (o ReplicationModeOutput) ToReplicationModeOutputWithContext(ctx context.Context) ReplicationModeOutput {
	return o
}

func (o ReplicationModeOutput) ToReplicationModePtrOutput() ReplicationModePtrOutput {
	return o.ToReplicationModePtrOutputWithContext(context.Background())
}

func (o ReplicationModeOutput) ToReplicationModePtrOutputWithContext(ctx context.Context) ReplicationModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReplicationMode) *ReplicationMode {
		return &v
	}).(ReplicationModePtrOutput)
}

func (o ReplicationModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ReplicationModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReplicationMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ReplicationModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReplicationModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReplicationMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ReplicationModePtrOutput struct{ *pulumi.OutputState }

func (ReplicationModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationMode)(nil)).Elem()
}

func (o ReplicationModePtrOutput) ToReplicationModePtrOutput() ReplicationModePtrOutput {
	return o
}

func (o ReplicationModePtrOutput) ToReplicationModePtrOutputWithContext(ctx context.Context) ReplicationModePtrOutput {
	return o
}

func (o ReplicationModePtrOutput) Elem() ReplicationModeOutput {
	return o.ApplyT(func(v *ReplicationMode) ReplicationMode {
		if v != nil {
			return *v
		}
		var ret ReplicationMode
		return ret
	}).(ReplicationModeOutput)
}

func (o ReplicationModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReplicationModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ReplicationMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ReplicationModeInput interface {
	pulumi.Input

	ToReplicationModeOutput() ReplicationModeOutput
	ToReplicationModeOutputWithContext(context.Context) ReplicationModeOutput
}

var replicationModePtrType = reflect.TypeOf((**ReplicationMode)(nil)).Elem()

type ReplicationModePtrInput interface {
	pulumi.Input

	ToReplicationModePtrOutput() ReplicationModePtrOutput
	ToReplicationModePtrOutputWithContext(context.Context) ReplicationModePtrOutput
}

type replicationModePtr string

func ReplicationModePtr(v string) ReplicationModePtrInput {
	return (*replicationModePtr)(&v)
}

func (*replicationModePtr) ElementType() reflect.Type {
	return replicationModePtrType
}

func (in *replicationModePtr) ToReplicationModePtrOutput() ReplicationModePtrOutput {
	return pulumi.ToOutput(in).(ReplicationModePtrOutput)
}

func (in *replicationModePtr) ToReplicationModePtrOutputWithContext(ctx context.Context) ReplicationModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ReplicationModePtrOutput)
}

type SampleName string

const (
	SampleNameAdventureWorksLT       = SampleName("AdventureWorksLT")
	SampleNameWideWorldImportersStd  = SampleName("WideWorldImportersStd")
	SampleNameWideWorldImportersFull = SampleName("WideWorldImportersFull")
)

func (SampleName) ElementType() reflect.Type {
	return reflect.TypeOf((*SampleName)(nil)).Elem()
}

func (e SampleName) ToSampleNameOutput() SampleNameOutput {
	return pulumi.ToOutput(e).(SampleNameOutput)
}

func (e SampleName) ToSampleNameOutputWithContext(ctx context.Context) SampleNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SampleNameOutput)
}

func (e SampleName) ToSampleNamePtrOutput() SampleNamePtrOutput {
	return e.ToSampleNamePtrOutputWithContext(context.Background())
}

func (e SampleName) ToSampleNamePtrOutputWithContext(ctx context.Context) SampleNamePtrOutput {
	return SampleName(e).ToSampleNameOutputWithContext(ctx).ToSampleNamePtrOutputWithContext(ctx)
}

func (e SampleName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SampleName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SampleName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SampleName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SampleNameOutput struct{ *pulumi.OutputState }

func (SampleNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SampleName)(nil)).Elem()
}

func (o SampleNameOutput) ToSampleNameOutput() SampleNameOutput {
	return o
}

func (o SampleNameOutput) ToSampleNameOutputWithContext(ctx context.Context) SampleNameOutput {
	return o
}

func (o SampleNameOutput) ToSampleNamePtrOutput() SampleNamePtrOutput {
	return o.ToSampleNamePtrOutputWithContext(context.Background())
}

func (o SampleNameOutput) ToSampleNamePtrOutputWithContext(ctx context.Context) SampleNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SampleName) *SampleName {
		return &v
	}).(SampleNamePtrOutput)
}

func (o SampleNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SampleNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SampleName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SampleNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SampleNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SampleName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SampleNamePtrOutput struct{ *pulumi.OutputState }

func (SampleNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SampleName)(nil)).Elem()
}

func (o SampleNamePtrOutput) ToSampleNamePtrOutput() SampleNamePtrOutput {
	return o
}

func (o SampleNamePtrOutput) ToSampleNamePtrOutputWithContext(ctx context.Context) SampleNamePtrOutput {
	return o
}

func (o SampleNamePtrOutput) Elem() SampleNameOutput {
	return o.ApplyT(func(v *SampleName) SampleName {
		if v != nil {
			return *v
		}
		var ret SampleName
		return ret
	}).(SampleNameOutput)
}

func (o SampleNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SampleNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SampleName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SampleNameInput interface {
	pulumi.Input

	ToSampleNameOutput() SampleNameOutput
	ToSampleNameOutputWithContext(context.Context) SampleNameOutput
}

var sampleNamePtrType = reflect.TypeOf((**SampleName)(nil)).Elem()

type SampleNamePtrInput interface {
	pulumi.Input

	ToSampleNamePtrOutput() SampleNamePtrOutput
	ToSampleNamePtrOutputWithContext(context.Context) SampleNamePtrOutput
}

type sampleNamePtr string

func SampleNamePtr(v string) SampleNamePtrInput {
	return (*sampleNamePtr)(&v)
}

func (*sampleNamePtr) ElementType() reflect.Type {
	return sampleNamePtrType
}

func (in *sampleNamePtr) ToSampleNamePtrOutput() SampleNamePtrOutput {
	return pulumi.ToOutput(in).(SampleNamePtrOutput)
}

func (in *sampleNamePtr) ToSampleNamePtrOutputWithContext(ctx context.Context) SampleNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SampleNamePtrOutput)
}

type SecondaryType string

const (
	SecondaryTypeGeo   = SecondaryType("Geo")
	SecondaryTypeNamed = SecondaryType("Named")
)

func (SecondaryType) ElementType() reflect.Type {
	return reflect.TypeOf((*SecondaryType)(nil)).Elem()
}

func (e SecondaryType) ToSecondaryTypeOutput() SecondaryTypeOutput {
	return pulumi.ToOutput(e).(SecondaryTypeOutput)
}

func (e SecondaryType) ToSecondaryTypeOutputWithContext(ctx context.Context) SecondaryTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SecondaryTypeOutput)
}

func (e SecondaryType) ToSecondaryTypePtrOutput() SecondaryTypePtrOutput {
	return e.ToSecondaryTypePtrOutputWithContext(context.Background())
}

func (e SecondaryType) ToSecondaryTypePtrOutputWithContext(ctx context.Context) SecondaryTypePtrOutput {
	return SecondaryType(e).ToSecondaryTypeOutputWithContext(ctx).ToSecondaryTypePtrOutputWithContext(ctx)
}

func (e SecondaryType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecondaryType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecondaryType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecondaryType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SecondaryTypeOutput struct{ *pulumi.OutputState }

func (SecondaryTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecondaryType)(nil)).Elem()
}

func (o SecondaryTypeOutput) ToSecondaryTypeOutput() SecondaryTypeOutput {
	return o
}

func (o SecondaryTypeOutput) ToSecondaryTypeOutputWithContext(ctx context.Context) SecondaryTypeOutput {
	return o
}

func (o SecondaryTypeOutput) ToSecondaryTypePtrOutput() SecondaryTypePtrOutput {
	return o.ToSecondaryTypePtrOutputWithContext(context.Background())
}

func (o SecondaryTypeOutput) ToSecondaryTypePtrOutputWithContext(ctx context.Context) SecondaryTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecondaryType) *SecondaryType {
		return &v
	}).(SecondaryTypePtrOutput)
}

func (o SecondaryTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecondaryTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecondaryType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecondaryTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecondaryTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecondaryType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecondaryTypePtrOutput struct{ *pulumi.OutputState }

func (SecondaryTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecondaryType)(nil)).Elem()
}

func (o SecondaryTypePtrOutput) ToSecondaryTypePtrOutput() SecondaryTypePtrOutput {
	return o
}

func (o SecondaryTypePtrOutput) ToSecondaryTypePtrOutputWithContext(ctx context.Context) SecondaryTypePtrOutput {
	return o
}

func (o SecondaryTypePtrOutput) Elem() SecondaryTypeOutput {
	return o.ApplyT(func(v *SecondaryType) SecondaryType {
		if v != nil {
			return *v
		}
		var ret SecondaryType
		return ret
	}).(SecondaryTypeOutput)
}

func (o SecondaryTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecondaryTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecondaryType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SecondaryTypeInput interface {
	pulumi.Input

	ToSecondaryTypeOutput() SecondaryTypeOutput
	ToSecondaryTypeOutputWithContext(context.Context) SecondaryTypeOutput
}

var secondaryTypePtrType = reflect.TypeOf((**SecondaryType)(nil)).Elem()

type SecondaryTypePtrInput interface {
	pulumi.Input

	ToSecondaryTypePtrOutput() SecondaryTypePtrOutput
	ToSecondaryTypePtrOutputWithContext(context.Context) SecondaryTypePtrOutput
}

type secondaryTypePtr string

func SecondaryTypePtr(v string) SecondaryTypePtrInput {
	return (*secondaryTypePtr)(&v)
}

func (*secondaryTypePtr) ElementType() reflect.Type {
	return secondaryTypePtrType
}

func (in *secondaryTypePtr) ToSecondaryTypePtrOutput() SecondaryTypePtrOutput {
	return pulumi.ToOutput(in).(SecondaryTypePtrOutput)
}

func (in *secondaryTypePtr) ToSecondaryTypePtrOutputWithContext(ctx context.Context) SecondaryTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SecondaryTypePtrOutput)
}

type SecurityAlertsPolicyState string

const (
	SecurityAlertsPolicyStateEnabled  = SecurityAlertsPolicyState("Enabled")
	SecurityAlertsPolicyStateDisabled = SecurityAlertsPolicyState("Disabled")
)

func (SecurityAlertsPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAlertsPolicyState)(nil)).Elem()
}

func (e SecurityAlertsPolicyState) ToSecurityAlertsPolicyStateOutput() SecurityAlertsPolicyStateOutput {
	return pulumi.ToOutput(e).(SecurityAlertsPolicyStateOutput)
}

func (e SecurityAlertsPolicyState) ToSecurityAlertsPolicyStateOutputWithContext(ctx context.Context) SecurityAlertsPolicyStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SecurityAlertsPolicyStateOutput)
}

func (e SecurityAlertsPolicyState) ToSecurityAlertsPolicyStatePtrOutput() SecurityAlertsPolicyStatePtrOutput {
	return e.ToSecurityAlertsPolicyStatePtrOutputWithContext(context.Background())
}

func (e SecurityAlertsPolicyState) ToSecurityAlertsPolicyStatePtrOutputWithContext(ctx context.Context) SecurityAlertsPolicyStatePtrOutput {
	return SecurityAlertsPolicyState(e).ToSecurityAlertsPolicyStateOutputWithContext(ctx).ToSecurityAlertsPolicyStatePtrOutputWithContext(ctx)
}

func (e SecurityAlertsPolicyState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityAlertsPolicyState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityAlertsPolicyState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecurityAlertsPolicyState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SecurityAlertsPolicyStateOutput struct{ *pulumi.OutputState }

func (SecurityAlertsPolicyStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAlertsPolicyState)(nil)).Elem()
}

func (o SecurityAlertsPolicyStateOutput) ToSecurityAlertsPolicyStateOutput() SecurityAlertsPolicyStateOutput {
	return o
}

func (o SecurityAlertsPolicyStateOutput) ToSecurityAlertsPolicyStateOutputWithContext(ctx context.Context) SecurityAlertsPolicyStateOutput {
	return o
}

func (o SecurityAlertsPolicyStateOutput) ToSecurityAlertsPolicyStatePtrOutput() SecurityAlertsPolicyStatePtrOutput {
	return o.ToSecurityAlertsPolicyStatePtrOutputWithContext(context.Background())
}

func (o SecurityAlertsPolicyStateOutput) ToSecurityAlertsPolicyStatePtrOutputWithContext(ctx context.Context) SecurityAlertsPolicyStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityAlertsPolicyState) *SecurityAlertsPolicyState {
		return &v
	}).(SecurityAlertsPolicyStatePtrOutput)
}

func (o SecurityAlertsPolicyStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecurityAlertsPolicyStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityAlertsPolicyState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecurityAlertsPolicyStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityAlertsPolicyStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityAlertsPolicyState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecurityAlertsPolicyStatePtrOutput struct{ *pulumi.OutputState }

func (SecurityAlertsPolicyStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAlertsPolicyState)(nil)).Elem()
}

func (o SecurityAlertsPolicyStatePtrOutput) ToSecurityAlertsPolicyStatePtrOutput() SecurityAlertsPolicyStatePtrOutput {
	return o
}

func (o SecurityAlertsPolicyStatePtrOutput) ToSecurityAlertsPolicyStatePtrOutputWithContext(ctx context.Context) SecurityAlertsPolicyStatePtrOutput {
	return o
}

func (o SecurityAlertsPolicyStatePtrOutput) Elem() SecurityAlertsPolicyStateOutput {
	return o.ApplyT(func(v *SecurityAlertsPolicyState) SecurityAlertsPolicyState {
		if v != nil {
			return *v
		}
		var ret SecurityAlertsPolicyState
		return ret
	}).(SecurityAlertsPolicyStateOutput)
}

func (o SecurityAlertsPolicyStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityAlertsPolicyStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecurityAlertsPolicyState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SecurityAlertsPolicyStateInput interface {
	pulumi.Input

	ToSecurityAlertsPolicyStateOutput() SecurityAlertsPolicyStateOutput
	ToSecurityAlertsPolicyStateOutputWithContext(context.Context) SecurityAlertsPolicyStateOutput
}

var securityAlertsPolicyStatePtrType = reflect.TypeOf((**SecurityAlertsPolicyState)(nil)).Elem()

type SecurityAlertsPolicyStatePtrInput interface {
	pulumi.Input

	ToSecurityAlertsPolicyStatePtrOutput() SecurityAlertsPolicyStatePtrOutput
	ToSecurityAlertsPolicyStatePtrOutputWithContext(context.Context) SecurityAlertsPolicyStatePtrOutput
}

type securityAlertsPolicyStatePtr string

func SecurityAlertsPolicyStatePtr(v string) SecurityAlertsPolicyStatePtrInput {
	return (*securityAlertsPolicyStatePtr)(&v)
}

func (*securityAlertsPolicyStatePtr) ElementType() reflect.Type {
	return securityAlertsPolicyStatePtrType
}

func (in *securityAlertsPolicyStatePtr) ToSecurityAlertsPolicyStatePtrOutput() SecurityAlertsPolicyStatePtrOutput {
	return pulumi.ToOutput(in).(SecurityAlertsPolicyStatePtrOutput)
}

func (in *securityAlertsPolicyStatePtr) ToSecurityAlertsPolicyStatePtrOutputWithContext(ctx context.Context) SecurityAlertsPolicyStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SecurityAlertsPolicyStatePtrOutput)
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

type ServerKeyType string

const (
	ServerKeyTypeServiceManaged = ServerKeyType("ServiceManaged")
	ServerKeyTypeAzureKeyVault  = ServerKeyType("AzureKeyVault")
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

type ServerNetworkAccessFlag string

const (
	ServerNetworkAccessFlagEnabled  = ServerNetworkAccessFlag("Enabled")
	ServerNetworkAccessFlagDisabled = ServerNetworkAccessFlag("Disabled")
)

func (ServerNetworkAccessFlag) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerNetworkAccessFlag)(nil)).Elem()
}

func (e ServerNetworkAccessFlag) ToServerNetworkAccessFlagOutput() ServerNetworkAccessFlagOutput {
	return pulumi.ToOutput(e).(ServerNetworkAccessFlagOutput)
}

func (e ServerNetworkAccessFlag) ToServerNetworkAccessFlagOutputWithContext(ctx context.Context) ServerNetworkAccessFlagOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServerNetworkAccessFlagOutput)
}

func (e ServerNetworkAccessFlag) ToServerNetworkAccessFlagPtrOutput() ServerNetworkAccessFlagPtrOutput {
	return e.ToServerNetworkAccessFlagPtrOutputWithContext(context.Background())
}

func (e ServerNetworkAccessFlag) ToServerNetworkAccessFlagPtrOutputWithContext(ctx context.Context) ServerNetworkAccessFlagPtrOutput {
	return ServerNetworkAccessFlag(e).ToServerNetworkAccessFlagOutputWithContext(ctx).ToServerNetworkAccessFlagPtrOutputWithContext(ctx)
}

func (e ServerNetworkAccessFlag) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServerNetworkAccessFlag) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServerNetworkAccessFlag) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServerNetworkAccessFlag) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServerNetworkAccessFlagOutput struct{ *pulumi.OutputState }

func (ServerNetworkAccessFlagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerNetworkAccessFlag)(nil)).Elem()
}

func (o ServerNetworkAccessFlagOutput) ToServerNetworkAccessFlagOutput() ServerNetworkAccessFlagOutput {
	return o
}

func (o ServerNetworkAccessFlagOutput) ToServerNetworkAccessFlagOutputWithContext(ctx context.Context) ServerNetworkAccessFlagOutput {
	return o
}

func (o ServerNetworkAccessFlagOutput) ToServerNetworkAccessFlagPtrOutput() ServerNetworkAccessFlagPtrOutput {
	return o.ToServerNetworkAccessFlagPtrOutputWithContext(context.Background())
}

func (o ServerNetworkAccessFlagOutput) ToServerNetworkAccessFlagPtrOutputWithContext(ctx context.Context) ServerNetworkAccessFlagPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerNetworkAccessFlag) *ServerNetworkAccessFlag {
		return &v
	}).(ServerNetworkAccessFlagPtrOutput)
}

func (o ServerNetworkAccessFlagOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServerNetworkAccessFlagOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServerNetworkAccessFlag) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServerNetworkAccessFlagOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServerNetworkAccessFlagOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServerNetworkAccessFlag) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServerNetworkAccessFlagPtrOutput struct{ *pulumi.OutputState }

func (ServerNetworkAccessFlagPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerNetworkAccessFlag)(nil)).Elem()
}

func (o ServerNetworkAccessFlagPtrOutput) ToServerNetworkAccessFlagPtrOutput() ServerNetworkAccessFlagPtrOutput {
	return o
}

func (o ServerNetworkAccessFlagPtrOutput) ToServerNetworkAccessFlagPtrOutputWithContext(ctx context.Context) ServerNetworkAccessFlagPtrOutput {
	return o
}

func (o ServerNetworkAccessFlagPtrOutput) Elem() ServerNetworkAccessFlagOutput {
	return o.ApplyT(func(v *ServerNetworkAccessFlag) ServerNetworkAccessFlag {
		if v != nil {
			return *v
		}
		var ret ServerNetworkAccessFlag
		return ret
	}).(ServerNetworkAccessFlagOutput)
}

func (o ServerNetworkAccessFlagPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServerNetworkAccessFlagPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServerNetworkAccessFlag) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServerNetworkAccessFlagInput interface {
	pulumi.Input

	ToServerNetworkAccessFlagOutput() ServerNetworkAccessFlagOutput
	ToServerNetworkAccessFlagOutputWithContext(context.Context) ServerNetworkAccessFlagOutput
}

var serverNetworkAccessFlagPtrType = reflect.TypeOf((**ServerNetworkAccessFlag)(nil)).Elem()

type ServerNetworkAccessFlagPtrInput interface {
	pulumi.Input

	ToServerNetworkAccessFlagPtrOutput() ServerNetworkAccessFlagPtrOutput
	ToServerNetworkAccessFlagPtrOutputWithContext(context.Context) ServerNetworkAccessFlagPtrOutput
}

type serverNetworkAccessFlagPtr string

func ServerNetworkAccessFlagPtr(v string) ServerNetworkAccessFlagPtrInput {
	return (*serverNetworkAccessFlagPtr)(&v)
}

func (*serverNetworkAccessFlagPtr) ElementType() reflect.Type {
	return serverNetworkAccessFlagPtrType
}

func (in *serverNetworkAccessFlagPtr) ToServerNetworkAccessFlagPtrOutput() ServerNetworkAccessFlagPtrOutput {
	return pulumi.ToOutput(in).(ServerNetworkAccessFlagPtrOutput)
}

func (in *serverNetworkAccessFlagPtr) ToServerNetworkAccessFlagPtrOutputWithContext(ctx context.Context) ServerNetworkAccessFlagPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServerNetworkAccessFlagPtrOutput)
}

type ServicePrincipalType string

const (
	ServicePrincipalTypeNone           = ServicePrincipalType("None")
	ServicePrincipalTypeSystemAssigned = ServicePrincipalType("SystemAssigned")
)

func (ServicePrincipalType) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalType)(nil)).Elem()
}

func (e ServicePrincipalType) ToServicePrincipalTypeOutput() ServicePrincipalTypeOutput {
	return pulumi.ToOutput(e).(ServicePrincipalTypeOutput)
}

func (e ServicePrincipalType) ToServicePrincipalTypeOutputWithContext(ctx context.Context) ServicePrincipalTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServicePrincipalTypeOutput)
}

func (e ServicePrincipalType) ToServicePrincipalTypePtrOutput() ServicePrincipalTypePtrOutput {
	return e.ToServicePrincipalTypePtrOutputWithContext(context.Background())
}

func (e ServicePrincipalType) ToServicePrincipalTypePtrOutputWithContext(ctx context.Context) ServicePrincipalTypePtrOutput {
	return ServicePrincipalType(e).ToServicePrincipalTypeOutputWithContext(ctx).ToServicePrincipalTypePtrOutputWithContext(ctx)
}

func (e ServicePrincipalType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServicePrincipalType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServicePrincipalType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServicePrincipalType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServicePrincipalTypeOutput struct{ *pulumi.OutputState }

func (ServicePrincipalTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalType)(nil)).Elem()
}

func (o ServicePrincipalTypeOutput) ToServicePrincipalTypeOutput() ServicePrincipalTypeOutput {
	return o
}

func (o ServicePrincipalTypeOutput) ToServicePrincipalTypeOutputWithContext(ctx context.Context) ServicePrincipalTypeOutput {
	return o
}

func (o ServicePrincipalTypeOutput) ToServicePrincipalTypePtrOutput() ServicePrincipalTypePtrOutput {
	return o.ToServicePrincipalTypePtrOutputWithContext(context.Background())
}

func (o ServicePrincipalTypeOutput) ToServicePrincipalTypePtrOutputWithContext(ctx context.Context) ServicePrincipalTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicePrincipalType) *ServicePrincipalType {
		return &v
	}).(ServicePrincipalTypePtrOutput)
}

func (o ServicePrincipalTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServicePrincipalTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServicePrincipalType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServicePrincipalTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServicePrincipalTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServicePrincipalType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServicePrincipalTypePtrOutput struct{ *pulumi.OutputState }

func (ServicePrincipalTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalType)(nil)).Elem()
}

func (o ServicePrincipalTypePtrOutput) ToServicePrincipalTypePtrOutput() ServicePrincipalTypePtrOutput {
	return o
}

func (o ServicePrincipalTypePtrOutput) ToServicePrincipalTypePtrOutputWithContext(ctx context.Context) ServicePrincipalTypePtrOutput {
	return o
}

func (o ServicePrincipalTypePtrOutput) Elem() ServicePrincipalTypeOutput {
	return o.ApplyT(func(v *ServicePrincipalType) ServicePrincipalType {
		if v != nil {
			return *v
		}
		var ret ServicePrincipalType
		return ret
	}).(ServicePrincipalTypeOutput)
}

func (o ServicePrincipalTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServicePrincipalTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServicePrincipalType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServicePrincipalTypeInput interface {
	pulumi.Input

	ToServicePrincipalTypeOutput() ServicePrincipalTypeOutput
	ToServicePrincipalTypeOutputWithContext(context.Context) ServicePrincipalTypeOutput
}

var servicePrincipalTypePtrType = reflect.TypeOf((**ServicePrincipalType)(nil)).Elem()

type ServicePrincipalTypePtrInput interface {
	pulumi.Input

	ToServicePrincipalTypePtrOutput() ServicePrincipalTypePtrOutput
	ToServicePrincipalTypePtrOutputWithContext(context.Context) ServicePrincipalTypePtrOutput
}

type servicePrincipalTypePtr string

func ServicePrincipalTypePtr(v string) ServicePrincipalTypePtrInput {
	return (*servicePrincipalTypePtr)(&v)
}

func (*servicePrincipalTypePtr) ElementType() reflect.Type {
	return servicePrincipalTypePtrType
}

func (in *servicePrincipalTypePtr) ToServicePrincipalTypePtrOutput() ServicePrincipalTypePtrOutput {
	return pulumi.ToOutput(in).(ServicePrincipalTypePtrOutput)
}

func (in *servicePrincipalTypePtr) ToServicePrincipalTypePtrOutputWithContext(ctx context.Context) ServicePrincipalTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServicePrincipalTypePtrOutput)
}

type SyncConflictResolutionPolicy string

const (
	SyncConflictResolutionPolicyHubWin    = SyncConflictResolutionPolicy("HubWin")
	SyncConflictResolutionPolicyMemberWin = SyncConflictResolutionPolicy("MemberWin")
)

func (SyncConflictResolutionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncConflictResolutionPolicy)(nil)).Elem()
}

func (e SyncConflictResolutionPolicy) ToSyncConflictResolutionPolicyOutput() SyncConflictResolutionPolicyOutput {
	return pulumi.ToOutput(e).(SyncConflictResolutionPolicyOutput)
}

func (e SyncConflictResolutionPolicy) ToSyncConflictResolutionPolicyOutputWithContext(ctx context.Context) SyncConflictResolutionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SyncConflictResolutionPolicyOutput)
}

func (e SyncConflictResolutionPolicy) ToSyncConflictResolutionPolicyPtrOutput() SyncConflictResolutionPolicyPtrOutput {
	return e.ToSyncConflictResolutionPolicyPtrOutputWithContext(context.Background())
}

func (e SyncConflictResolutionPolicy) ToSyncConflictResolutionPolicyPtrOutputWithContext(ctx context.Context) SyncConflictResolutionPolicyPtrOutput {
	return SyncConflictResolutionPolicy(e).ToSyncConflictResolutionPolicyOutputWithContext(ctx).ToSyncConflictResolutionPolicyPtrOutputWithContext(ctx)
}

func (e SyncConflictResolutionPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SyncConflictResolutionPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SyncConflictResolutionPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SyncConflictResolutionPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SyncConflictResolutionPolicyOutput struct{ *pulumi.OutputState }

func (SyncConflictResolutionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncConflictResolutionPolicy)(nil)).Elem()
}

func (o SyncConflictResolutionPolicyOutput) ToSyncConflictResolutionPolicyOutput() SyncConflictResolutionPolicyOutput {
	return o
}

func (o SyncConflictResolutionPolicyOutput) ToSyncConflictResolutionPolicyOutputWithContext(ctx context.Context) SyncConflictResolutionPolicyOutput {
	return o
}

func (o SyncConflictResolutionPolicyOutput) ToSyncConflictResolutionPolicyPtrOutput() SyncConflictResolutionPolicyPtrOutput {
	return o.ToSyncConflictResolutionPolicyPtrOutputWithContext(context.Background())
}

func (o SyncConflictResolutionPolicyOutput) ToSyncConflictResolutionPolicyPtrOutputWithContext(ctx context.Context) SyncConflictResolutionPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncConflictResolutionPolicy) *SyncConflictResolutionPolicy {
		return &v
	}).(SyncConflictResolutionPolicyPtrOutput)
}

func (o SyncConflictResolutionPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SyncConflictResolutionPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncConflictResolutionPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SyncConflictResolutionPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncConflictResolutionPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncConflictResolutionPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SyncConflictResolutionPolicyPtrOutput struct{ *pulumi.OutputState }

func (SyncConflictResolutionPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncConflictResolutionPolicy)(nil)).Elem()
}

func (o SyncConflictResolutionPolicyPtrOutput) ToSyncConflictResolutionPolicyPtrOutput() SyncConflictResolutionPolicyPtrOutput {
	return o
}

func (o SyncConflictResolutionPolicyPtrOutput) ToSyncConflictResolutionPolicyPtrOutputWithContext(ctx context.Context) SyncConflictResolutionPolicyPtrOutput {
	return o
}

func (o SyncConflictResolutionPolicyPtrOutput) Elem() SyncConflictResolutionPolicyOutput {
	return o.ApplyT(func(v *SyncConflictResolutionPolicy) SyncConflictResolutionPolicy {
		if v != nil {
			return *v
		}
		var ret SyncConflictResolutionPolicy
		return ret
	}).(SyncConflictResolutionPolicyOutput)
}

func (o SyncConflictResolutionPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncConflictResolutionPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SyncConflictResolutionPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SyncConflictResolutionPolicyInput interface {
	pulumi.Input

	ToSyncConflictResolutionPolicyOutput() SyncConflictResolutionPolicyOutput
	ToSyncConflictResolutionPolicyOutputWithContext(context.Context) SyncConflictResolutionPolicyOutput
}

var syncConflictResolutionPolicyPtrType = reflect.TypeOf((**SyncConflictResolutionPolicy)(nil)).Elem()

type SyncConflictResolutionPolicyPtrInput interface {
	pulumi.Input

	ToSyncConflictResolutionPolicyPtrOutput() SyncConflictResolutionPolicyPtrOutput
	ToSyncConflictResolutionPolicyPtrOutputWithContext(context.Context) SyncConflictResolutionPolicyPtrOutput
}

type syncConflictResolutionPolicyPtr string

func SyncConflictResolutionPolicyPtr(v string) SyncConflictResolutionPolicyPtrInput {
	return (*syncConflictResolutionPolicyPtr)(&v)
}

func (*syncConflictResolutionPolicyPtr) ElementType() reflect.Type {
	return syncConflictResolutionPolicyPtrType
}

func (in *syncConflictResolutionPolicyPtr) ToSyncConflictResolutionPolicyPtrOutput() SyncConflictResolutionPolicyPtrOutput {
	return pulumi.ToOutput(in).(SyncConflictResolutionPolicyPtrOutput)
}

func (in *syncConflictResolutionPolicyPtr) ToSyncConflictResolutionPolicyPtrOutputWithContext(ctx context.Context) SyncConflictResolutionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SyncConflictResolutionPolicyPtrOutput)
}

type SyncDirection string

const (
	SyncDirectionBidirectional     = SyncDirection("Bidirectional")
	SyncDirectionOneWayMemberToHub = SyncDirection("OneWayMemberToHub")
	SyncDirectionOneWayHubToMember = SyncDirection("OneWayHubToMember")
)

func (SyncDirection) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncDirection)(nil)).Elem()
}

func (e SyncDirection) ToSyncDirectionOutput() SyncDirectionOutput {
	return pulumi.ToOutput(e).(SyncDirectionOutput)
}

func (e SyncDirection) ToSyncDirectionOutputWithContext(ctx context.Context) SyncDirectionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SyncDirectionOutput)
}

func (e SyncDirection) ToSyncDirectionPtrOutput() SyncDirectionPtrOutput {
	return e.ToSyncDirectionPtrOutputWithContext(context.Background())
}

func (e SyncDirection) ToSyncDirectionPtrOutputWithContext(ctx context.Context) SyncDirectionPtrOutput {
	return SyncDirection(e).ToSyncDirectionOutputWithContext(ctx).ToSyncDirectionPtrOutputWithContext(ctx)
}

func (e SyncDirection) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SyncDirection) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SyncDirection) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SyncDirection) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SyncDirectionOutput struct{ *pulumi.OutputState }

func (SyncDirectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncDirection)(nil)).Elem()
}

func (o SyncDirectionOutput) ToSyncDirectionOutput() SyncDirectionOutput {
	return o
}

func (o SyncDirectionOutput) ToSyncDirectionOutputWithContext(ctx context.Context) SyncDirectionOutput {
	return o
}

func (o SyncDirectionOutput) ToSyncDirectionPtrOutput() SyncDirectionPtrOutput {
	return o.ToSyncDirectionPtrOutputWithContext(context.Background())
}

func (o SyncDirectionOutput) ToSyncDirectionPtrOutputWithContext(ctx context.Context) SyncDirectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncDirection) *SyncDirection {
		return &v
	}).(SyncDirectionPtrOutput)
}

func (o SyncDirectionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SyncDirectionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncDirection) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SyncDirectionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncDirectionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncDirection) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SyncDirectionPtrOutput struct{ *pulumi.OutputState }

func (SyncDirectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncDirection)(nil)).Elem()
}

func (o SyncDirectionPtrOutput) ToSyncDirectionPtrOutput() SyncDirectionPtrOutput {
	return o
}

func (o SyncDirectionPtrOutput) ToSyncDirectionPtrOutputWithContext(ctx context.Context) SyncDirectionPtrOutput {
	return o
}

func (o SyncDirectionPtrOutput) Elem() SyncDirectionOutput {
	return o.ApplyT(func(v *SyncDirection) SyncDirection {
		if v != nil {
			return *v
		}
		var ret SyncDirection
		return ret
	}).(SyncDirectionOutput)
}

func (o SyncDirectionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncDirectionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SyncDirection) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SyncDirectionInput interface {
	pulumi.Input

	ToSyncDirectionOutput() SyncDirectionOutput
	ToSyncDirectionOutputWithContext(context.Context) SyncDirectionOutput
}

var syncDirectionPtrType = reflect.TypeOf((**SyncDirection)(nil)).Elem()

type SyncDirectionPtrInput interface {
	pulumi.Input

	ToSyncDirectionPtrOutput() SyncDirectionPtrOutput
	ToSyncDirectionPtrOutputWithContext(context.Context) SyncDirectionPtrOutput
}

type syncDirectionPtr string

func SyncDirectionPtr(v string) SyncDirectionPtrInput {
	return (*syncDirectionPtr)(&v)
}

func (*syncDirectionPtr) ElementType() reflect.Type {
	return syncDirectionPtrType
}

func (in *syncDirectionPtr) ToSyncDirectionPtrOutput() SyncDirectionPtrOutput {
	return pulumi.ToOutput(in).(SyncDirectionPtrOutput)
}

func (in *syncDirectionPtr) ToSyncDirectionPtrOutputWithContext(ctx context.Context) SyncDirectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SyncDirectionPtrOutput)
}

type SyncMemberDbType string

const (
	SyncMemberDbTypeAzureSqlDatabase  = SyncMemberDbType("AzureSqlDatabase")
	SyncMemberDbTypeSqlServerDatabase = SyncMemberDbType("SqlServerDatabase")
)

func (SyncMemberDbType) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncMemberDbType)(nil)).Elem()
}

func (e SyncMemberDbType) ToSyncMemberDbTypeOutput() SyncMemberDbTypeOutput {
	return pulumi.ToOutput(e).(SyncMemberDbTypeOutput)
}

func (e SyncMemberDbType) ToSyncMemberDbTypeOutputWithContext(ctx context.Context) SyncMemberDbTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SyncMemberDbTypeOutput)
}

func (e SyncMemberDbType) ToSyncMemberDbTypePtrOutput() SyncMemberDbTypePtrOutput {
	return e.ToSyncMemberDbTypePtrOutputWithContext(context.Background())
}

func (e SyncMemberDbType) ToSyncMemberDbTypePtrOutputWithContext(ctx context.Context) SyncMemberDbTypePtrOutput {
	return SyncMemberDbType(e).ToSyncMemberDbTypeOutputWithContext(ctx).ToSyncMemberDbTypePtrOutputWithContext(ctx)
}

func (e SyncMemberDbType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SyncMemberDbType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SyncMemberDbType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SyncMemberDbType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SyncMemberDbTypeOutput struct{ *pulumi.OutputState }

func (SyncMemberDbTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncMemberDbType)(nil)).Elem()
}

func (o SyncMemberDbTypeOutput) ToSyncMemberDbTypeOutput() SyncMemberDbTypeOutput {
	return o
}

func (o SyncMemberDbTypeOutput) ToSyncMemberDbTypeOutputWithContext(ctx context.Context) SyncMemberDbTypeOutput {
	return o
}

func (o SyncMemberDbTypeOutput) ToSyncMemberDbTypePtrOutput() SyncMemberDbTypePtrOutput {
	return o.ToSyncMemberDbTypePtrOutputWithContext(context.Background())
}

func (o SyncMemberDbTypeOutput) ToSyncMemberDbTypePtrOutputWithContext(ctx context.Context) SyncMemberDbTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncMemberDbType) *SyncMemberDbType {
		return &v
	}).(SyncMemberDbTypePtrOutput)
}

func (o SyncMemberDbTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SyncMemberDbTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncMemberDbType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SyncMemberDbTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncMemberDbTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncMemberDbType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SyncMemberDbTypePtrOutput struct{ *pulumi.OutputState }

func (SyncMemberDbTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncMemberDbType)(nil)).Elem()
}

func (o SyncMemberDbTypePtrOutput) ToSyncMemberDbTypePtrOutput() SyncMemberDbTypePtrOutput {
	return o
}

func (o SyncMemberDbTypePtrOutput) ToSyncMemberDbTypePtrOutputWithContext(ctx context.Context) SyncMemberDbTypePtrOutput {
	return o
}

func (o SyncMemberDbTypePtrOutput) Elem() SyncMemberDbTypeOutput {
	return o.ApplyT(func(v *SyncMemberDbType) SyncMemberDbType {
		if v != nil {
			return *v
		}
		var ret SyncMemberDbType
		return ret
	}).(SyncMemberDbTypeOutput)
}

func (o SyncMemberDbTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncMemberDbTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SyncMemberDbType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SyncMemberDbTypeInput interface {
	pulumi.Input

	ToSyncMemberDbTypeOutput() SyncMemberDbTypeOutput
	ToSyncMemberDbTypeOutputWithContext(context.Context) SyncMemberDbTypeOutput
}

var syncMemberDbTypePtrType = reflect.TypeOf((**SyncMemberDbType)(nil)).Elem()

type SyncMemberDbTypePtrInput interface {
	pulumi.Input

	ToSyncMemberDbTypePtrOutput() SyncMemberDbTypePtrOutput
	ToSyncMemberDbTypePtrOutputWithContext(context.Context) SyncMemberDbTypePtrOutput
}

type syncMemberDbTypePtr string

func SyncMemberDbTypePtr(v string) SyncMemberDbTypePtrInput {
	return (*syncMemberDbTypePtr)(&v)
}

func (*syncMemberDbTypePtr) ElementType() reflect.Type {
	return syncMemberDbTypePtrType
}

func (in *syncMemberDbTypePtr) ToSyncMemberDbTypePtrOutput() SyncMemberDbTypePtrOutput {
	return pulumi.ToOutput(in).(SyncMemberDbTypePtrOutput)
}

func (in *syncMemberDbTypePtr) ToSyncMemberDbTypePtrOutputWithContext(ctx context.Context) SyncMemberDbTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SyncMemberDbTypePtrOutput)
}

type TransparentDataEncryptionStateEnum string

const (
	TransparentDataEncryptionStateEnumEnabled  = TransparentDataEncryptionStateEnum("Enabled")
	TransparentDataEncryptionStateEnumDisabled = TransparentDataEncryptionStateEnum("Disabled")
)

func (TransparentDataEncryptionStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*TransparentDataEncryptionStateEnum)(nil)).Elem()
}

func (e TransparentDataEncryptionStateEnum) ToTransparentDataEncryptionStateEnumOutput() TransparentDataEncryptionStateEnumOutput {
	return pulumi.ToOutput(e).(TransparentDataEncryptionStateEnumOutput)
}

func (e TransparentDataEncryptionStateEnum) ToTransparentDataEncryptionStateEnumOutputWithContext(ctx context.Context) TransparentDataEncryptionStateEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TransparentDataEncryptionStateEnumOutput)
}

func (e TransparentDataEncryptionStateEnum) ToTransparentDataEncryptionStateEnumPtrOutput() TransparentDataEncryptionStateEnumPtrOutput {
	return e.ToTransparentDataEncryptionStateEnumPtrOutputWithContext(context.Background())
}

func (e TransparentDataEncryptionStateEnum) ToTransparentDataEncryptionStateEnumPtrOutputWithContext(ctx context.Context) TransparentDataEncryptionStateEnumPtrOutput {
	return TransparentDataEncryptionStateEnum(e).ToTransparentDataEncryptionStateEnumOutputWithContext(ctx).ToTransparentDataEncryptionStateEnumPtrOutputWithContext(ctx)
}

func (e TransparentDataEncryptionStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TransparentDataEncryptionStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TransparentDataEncryptionStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TransparentDataEncryptionStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TransparentDataEncryptionStateEnumOutput struct{ *pulumi.OutputState }

func (TransparentDataEncryptionStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransparentDataEncryptionStateEnum)(nil)).Elem()
}

func (o TransparentDataEncryptionStateEnumOutput) ToTransparentDataEncryptionStateEnumOutput() TransparentDataEncryptionStateEnumOutput {
	return o
}

func (o TransparentDataEncryptionStateEnumOutput) ToTransparentDataEncryptionStateEnumOutputWithContext(ctx context.Context) TransparentDataEncryptionStateEnumOutput {
	return o
}

func (o TransparentDataEncryptionStateEnumOutput) ToTransparentDataEncryptionStateEnumPtrOutput() TransparentDataEncryptionStateEnumPtrOutput {
	return o.ToTransparentDataEncryptionStateEnumPtrOutputWithContext(context.Background())
}

func (o TransparentDataEncryptionStateEnumOutput) ToTransparentDataEncryptionStateEnumPtrOutputWithContext(ctx context.Context) TransparentDataEncryptionStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TransparentDataEncryptionStateEnum) *TransparentDataEncryptionStateEnum {
		return &v
	}).(TransparentDataEncryptionStateEnumPtrOutput)
}

func (o TransparentDataEncryptionStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TransparentDataEncryptionStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TransparentDataEncryptionStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TransparentDataEncryptionStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TransparentDataEncryptionStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TransparentDataEncryptionStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TransparentDataEncryptionStateEnumPtrOutput struct{ *pulumi.OutputState }

func (TransparentDataEncryptionStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransparentDataEncryptionStateEnum)(nil)).Elem()
}

func (o TransparentDataEncryptionStateEnumPtrOutput) ToTransparentDataEncryptionStateEnumPtrOutput() TransparentDataEncryptionStateEnumPtrOutput {
	return o
}

func (o TransparentDataEncryptionStateEnumPtrOutput) ToTransparentDataEncryptionStateEnumPtrOutputWithContext(ctx context.Context) TransparentDataEncryptionStateEnumPtrOutput {
	return o
}

func (o TransparentDataEncryptionStateEnumPtrOutput) Elem() TransparentDataEncryptionStateEnumOutput {
	return o.ApplyT(func(v *TransparentDataEncryptionStateEnum) TransparentDataEncryptionStateEnum {
		if v != nil {
			return *v
		}
		var ret TransparentDataEncryptionStateEnum
		return ret
	}).(TransparentDataEncryptionStateEnumOutput)
}

func (o TransparentDataEncryptionStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TransparentDataEncryptionStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TransparentDataEncryptionStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TransparentDataEncryptionStateEnumInput interface {
	pulumi.Input

	ToTransparentDataEncryptionStateEnumOutput() TransparentDataEncryptionStateEnumOutput
	ToTransparentDataEncryptionStateEnumOutputWithContext(context.Context) TransparentDataEncryptionStateEnumOutput
}

var transparentDataEncryptionStateEnumPtrType = reflect.TypeOf((**TransparentDataEncryptionStateEnum)(nil)).Elem()

type TransparentDataEncryptionStateEnumPtrInput interface {
	pulumi.Input

	ToTransparentDataEncryptionStateEnumPtrOutput() TransparentDataEncryptionStateEnumPtrOutput
	ToTransparentDataEncryptionStateEnumPtrOutputWithContext(context.Context) TransparentDataEncryptionStateEnumPtrOutput
}

type transparentDataEncryptionStateEnumPtr string

func TransparentDataEncryptionStateEnumPtr(v string) TransparentDataEncryptionStateEnumPtrInput {
	return (*transparentDataEncryptionStateEnumPtr)(&v)
}

func (*transparentDataEncryptionStateEnumPtr) ElementType() reflect.Type {
	return transparentDataEncryptionStateEnumPtrType
}

func (in *transparentDataEncryptionStateEnumPtr) ToTransparentDataEncryptionStateEnumPtrOutput() TransparentDataEncryptionStateEnumPtrOutput {
	return pulumi.ToOutput(in).(TransparentDataEncryptionStateEnumPtrOutput)
}

func (in *transparentDataEncryptionStateEnumPtr) ToTransparentDataEncryptionStateEnumPtrOutputWithContext(ctx context.Context) TransparentDataEncryptionStateEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TransparentDataEncryptionStateEnumPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AdministratorTypeOutput{})
	pulumi.RegisterOutputType(AdministratorTypePtrOutput{})
	pulumi.RegisterOutputType(AutoExecuteStatusOutput{})
	pulumi.RegisterOutputType(AutoExecuteStatusPtrOutput{})
	pulumi.RegisterOutputType(BackupStorageRedundancyOutput{})
	pulumi.RegisterOutputType(BackupStorageRedundancyPtrOutput{})
	pulumi.RegisterOutputType(BlobAuditingPolicyStateOutput{})
	pulumi.RegisterOutputType(BlobAuditingPolicyStatePtrOutput{})
	pulumi.RegisterOutputType(CatalogCollationTypeOutput{})
	pulumi.RegisterOutputType(CatalogCollationTypePtrOutput{})
	pulumi.RegisterOutputType(CreateModeOutput{})
	pulumi.RegisterOutputType(CreateModePtrOutput{})
	pulumi.RegisterOutputType(DatabaseIdentityTypeOutput{})
	pulumi.RegisterOutputType(DatabaseIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(DatabaseLicenseTypeOutput{})
	pulumi.RegisterOutputType(DatabaseLicenseTypePtrOutput{})
	pulumi.RegisterOutputType(DatabaseReadScaleOutput{})
	pulumi.RegisterOutputType(DatabaseReadScalePtrOutput{})
	pulumi.RegisterOutputType(ElasticPoolLicenseTypeOutput{})
	pulumi.RegisterOutputType(ElasticPoolLicenseTypePtrOutput{})
	pulumi.RegisterOutputType(IdentityTypeOutput{})
	pulumi.RegisterOutputType(IdentityTypePtrOutput{})
	pulumi.RegisterOutputType(InstancePoolLicenseTypeOutput{})
	pulumi.RegisterOutputType(InstancePoolLicenseTypePtrOutput{})
	pulumi.RegisterOutputType(JobScheduleTypeOutput{})
	pulumi.RegisterOutputType(JobScheduleTypePtrOutput{})
	pulumi.RegisterOutputType(JobStepActionSourceOutput{})
	pulumi.RegisterOutputType(JobStepActionSourcePtrOutput{})
	pulumi.RegisterOutputType(JobStepActionTypeOutput{})
	pulumi.RegisterOutputType(JobStepActionTypePtrOutput{})
	pulumi.RegisterOutputType(JobStepOutputTypeEnumOutput{})
	pulumi.RegisterOutputType(JobStepOutputTypeEnumPtrOutput{})
	pulumi.RegisterOutputType(JobTargetGroupMembershipTypeOutput{})
	pulumi.RegisterOutputType(JobTargetGroupMembershipTypePtrOutput{})
	pulumi.RegisterOutputType(JobTargetTypeOutput{})
	pulumi.RegisterOutputType(JobTargetTypePtrOutput{})
	pulumi.RegisterOutputType(ManagedDatabaseCreateModeOutput{})
	pulumi.RegisterOutputType(ManagedDatabaseCreateModePtrOutput{})
	pulumi.RegisterOutputType(ManagedInstanceAdministratorTypeOutput{})
	pulumi.RegisterOutputType(ManagedInstanceAdministratorTypePtrOutput{})
	pulumi.RegisterOutputType(ManagedInstanceLicenseTypeOutput{})
	pulumi.RegisterOutputType(ManagedInstanceLicenseTypePtrOutput{})
	pulumi.RegisterOutputType(ManagedInstanceProxyOverrideOutput{})
	pulumi.RegisterOutputType(ManagedInstanceProxyOverridePtrOutput{})
	pulumi.RegisterOutputType(ManagedServerCreateModeOutput{})
	pulumi.RegisterOutputType(ManagedServerCreateModePtrOutput{})
	pulumi.RegisterOutputType(PrincipalTypeOutput{})
	pulumi.RegisterOutputType(PrincipalTypePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateStatusOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateStatusPtrOutput{})
	pulumi.RegisterOutputType(ReadOnlyEndpointFailoverPolicyOutput{})
	pulumi.RegisterOutputType(ReadOnlyEndpointFailoverPolicyPtrOutput{})
	pulumi.RegisterOutputType(ReadWriteEndpointFailoverPolicyOutput{})
	pulumi.RegisterOutputType(ReadWriteEndpointFailoverPolicyPtrOutput{})
	pulumi.RegisterOutputType(ReplicationModeOutput{})
	pulumi.RegisterOutputType(ReplicationModePtrOutput{})
	pulumi.RegisterOutputType(SampleNameOutput{})
	pulumi.RegisterOutputType(SampleNamePtrOutput{})
	pulumi.RegisterOutputType(SecondaryTypeOutput{})
	pulumi.RegisterOutputType(SecondaryTypePtrOutput{})
	pulumi.RegisterOutputType(SecurityAlertsPolicyStateOutput{})
	pulumi.RegisterOutputType(SecurityAlertsPolicyStatePtrOutput{})
	pulumi.RegisterOutputType(SensitivityLabelRankOutput{})
	pulumi.RegisterOutputType(SensitivityLabelRankPtrOutput{})
	pulumi.RegisterOutputType(ServerKeyTypeOutput{})
	pulumi.RegisterOutputType(ServerKeyTypePtrOutput{})
	pulumi.RegisterOutputType(ServerNetworkAccessFlagOutput{})
	pulumi.RegisterOutputType(ServerNetworkAccessFlagPtrOutput{})
	pulumi.RegisterOutputType(ServicePrincipalTypeOutput{})
	pulumi.RegisterOutputType(ServicePrincipalTypePtrOutput{})
	pulumi.RegisterOutputType(SyncConflictResolutionPolicyOutput{})
	pulumi.RegisterOutputType(SyncConflictResolutionPolicyPtrOutput{})
	pulumi.RegisterOutputType(SyncDirectionOutput{})
	pulumi.RegisterOutputType(SyncDirectionPtrOutput{})
	pulumi.RegisterOutputType(SyncMemberDbTypeOutput{})
	pulumi.RegisterOutputType(SyncMemberDbTypePtrOutput{})
	pulumi.RegisterOutputType(TransparentDataEncryptionStateEnumOutput{})
	pulumi.RegisterOutputType(TransparentDataEncryptionStateEnumPtrOutput{})
}
