


package v20170701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BackupManagementType string

const (
	BackupManagementTypeInvalid           = BackupManagementType("Invalid")
	BackupManagementTypeAzureIaasVM       = BackupManagementType("AzureIaasVM")
	BackupManagementTypeMAB               = BackupManagementType("MAB")
	BackupManagementTypeDPM               = BackupManagementType("DPM")
	BackupManagementTypeAzureBackupServer = BackupManagementType("AzureBackupServer")
	BackupManagementTypeAzureSql          = BackupManagementType("AzureSql")
	BackupManagementTypeAzureStorage      = BackupManagementType("AzureStorage")
	BackupManagementTypeAzureWorkload     = BackupManagementType("AzureWorkload")
	BackupManagementTypeDefaultBackup     = BackupManagementType("DefaultBackup")
)

func (BackupManagementType) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupManagementType)(nil)).Elem()
}

func (e BackupManagementType) ToBackupManagementTypeOutput() BackupManagementTypeOutput {
	return pulumi.ToOutput(e).(BackupManagementTypeOutput)
}

func (e BackupManagementType) ToBackupManagementTypeOutputWithContext(ctx context.Context) BackupManagementTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BackupManagementTypeOutput)
}

func (e BackupManagementType) ToBackupManagementTypePtrOutput() BackupManagementTypePtrOutput {
	return e.ToBackupManagementTypePtrOutputWithContext(context.Background())
}

func (e BackupManagementType) ToBackupManagementTypePtrOutputWithContext(ctx context.Context) BackupManagementTypePtrOutput {
	return BackupManagementType(e).ToBackupManagementTypeOutputWithContext(ctx).ToBackupManagementTypePtrOutputWithContext(ctx)
}

func (e BackupManagementType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackupManagementType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BackupManagementType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BackupManagementType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BackupManagementTypeOutput struct{ *pulumi.OutputState }

func (BackupManagementTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupManagementType)(nil)).Elem()
}

func (o BackupManagementTypeOutput) ToBackupManagementTypeOutput() BackupManagementTypeOutput {
	return o
}

func (o BackupManagementTypeOutput) ToBackupManagementTypeOutputWithContext(ctx context.Context) BackupManagementTypeOutput {
	return o
}

func (o BackupManagementTypeOutput) ToBackupManagementTypePtrOutput() BackupManagementTypePtrOutput {
	return o.ToBackupManagementTypePtrOutputWithContext(context.Background())
}

func (o BackupManagementTypeOutput) ToBackupManagementTypePtrOutputWithContext(ctx context.Context) BackupManagementTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupManagementType) *BackupManagementType {
		return &v
	}).(BackupManagementTypePtrOutput)
}

func (o BackupManagementTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BackupManagementTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BackupManagementType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BackupManagementTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BackupManagementTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BackupManagementType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BackupManagementTypePtrOutput struct{ *pulumi.OutputState }

func (BackupManagementTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupManagementType)(nil)).Elem()
}

func (o BackupManagementTypePtrOutput) ToBackupManagementTypePtrOutput() BackupManagementTypePtrOutput {
	return o
}

func (o BackupManagementTypePtrOutput) ToBackupManagementTypePtrOutputWithContext(ctx context.Context) BackupManagementTypePtrOutput {
	return o
}

func (o BackupManagementTypePtrOutput) Elem() BackupManagementTypeOutput {
	return o.ApplyT(func(v *BackupManagementType) BackupManagementType {
		if v != nil {
			return *v
		}
		var ret BackupManagementType
		return ret
	}).(BackupManagementTypeOutput)
}

func (o BackupManagementTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BackupManagementTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BackupManagementType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BackupManagementTypeInput interface {
	pulumi.Input

	ToBackupManagementTypeOutput() BackupManagementTypeOutput
	ToBackupManagementTypeOutputWithContext(context.Context) BackupManagementTypeOutput
}

var backupManagementTypePtrType = reflect.TypeOf((**BackupManagementType)(nil)).Elem()

type BackupManagementTypePtrInput interface {
	pulumi.Input

	ToBackupManagementTypePtrOutput() BackupManagementTypePtrOutput
	ToBackupManagementTypePtrOutputWithContext(context.Context) BackupManagementTypePtrOutput
}

type backupManagementTypePtr string

func BackupManagementTypePtr(v string) BackupManagementTypePtrInput {
	return (*backupManagementTypePtr)(&v)
}

func (*backupManagementTypePtr) ElementType() reflect.Type {
	return backupManagementTypePtrType
}

func (in *backupManagementTypePtr) ToBackupManagementTypePtrOutput() BackupManagementTypePtrOutput {
	return pulumi.ToOutput(in).(BackupManagementTypePtrOutput)
}

func (in *backupManagementTypePtr) ToBackupManagementTypePtrOutputWithContext(ctx context.Context) BackupManagementTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BackupManagementTypePtrOutput)
}

type ProtectionStatus string

const (
	ProtectionStatusInvalid          = ProtectionStatus("Invalid")
	ProtectionStatusNotProtected     = ProtectionStatus("NotProtected")
	ProtectionStatusProtecting       = ProtectionStatus("Protecting")
	ProtectionStatusProtected        = ProtectionStatus("Protected")
	ProtectionStatusProtectionFailed = ProtectionStatus("ProtectionFailed")
)

func (ProtectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionStatus)(nil)).Elem()
}

func (e ProtectionStatus) ToProtectionStatusOutput() ProtectionStatusOutput {
	return pulumi.ToOutput(e).(ProtectionStatusOutput)
}

func (e ProtectionStatus) ToProtectionStatusOutputWithContext(ctx context.Context) ProtectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProtectionStatusOutput)
}

func (e ProtectionStatus) ToProtectionStatusPtrOutput() ProtectionStatusPtrOutput {
	return e.ToProtectionStatusPtrOutputWithContext(context.Background())
}

func (e ProtectionStatus) ToProtectionStatusPtrOutputWithContext(ctx context.Context) ProtectionStatusPtrOutput {
	return ProtectionStatus(e).ToProtectionStatusOutputWithContext(ctx).ToProtectionStatusPtrOutputWithContext(ctx)
}

func (e ProtectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProtectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProtectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProtectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProtectionStatusOutput struct{ *pulumi.OutputState }

func (ProtectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionStatus)(nil)).Elem()
}

func (o ProtectionStatusOutput) ToProtectionStatusOutput() ProtectionStatusOutput {
	return o
}

func (o ProtectionStatusOutput) ToProtectionStatusOutputWithContext(ctx context.Context) ProtectionStatusOutput {
	return o
}

func (o ProtectionStatusOutput) ToProtectionStatusPtrOutput() ProtectionStatusPtrOutput {
	return o.ToProtectionStatusPtrOutputWithContext(context.Background())
}

func (o ProtectionStatusOutput) ToProtectionStatusPtrOutputWithContext(ctx context.Context) ProtectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProtectionStatus) *ProtectionStatus {
		return &v
	}).(ProtectionStatusPtrOutput)
}

func (o ProtectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProtectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProtectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProtectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProtectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProtectionStatusPtrOutput struct{ *pulumi.OutputState }

func (ProtectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionStatus)(nil)).Elem()
}

func (o ProtectionStatusPtrOutput) ToProtectionStatusPtrOutput() ProtectionStatusPtrOutput {
	return o
}

func (o ProtectionStatusPtrOutput) ToProtectionStatusPtrOutputWithContext(ctx context.Context) ProtectionStatusPtrOutput {
	return o
}

func (o ProtectionStatusPtrOutput) Elem() ProtectionStatusOutput {
	return o.ApplyT(func(v *ProtectionStatus) ProtectionStatus {
		if v != nil {
			return *v
		}
		var ret ProtectionStatus
		return ret
	}).(ProtectionStatusOutput)
}

func (o ProtectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProtectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProtectionStatusInput interface {
	pulumi.Input

	ToProtectionStatusOutput() ProtectionStatusOutput
	ToProtectionStatusOutputWithContext(context.Context) ProtectionStatusOutput
}

var protectionStatusPtrType = reflect.TypeOf((**ProtectionStatus)(nil)).Elem()

type ProtectionStatusPtrInput interface {
	pulumi.Input

	ToProtectionStatusPtrOutput() ProtectionStatusPtrOutput
	ToProtectionStatusPtrOutputWithContext(context.Context) ProtectionStatusPtrOutput
}

type protectionStatusPtr string

func ProtectionStatusPtr(v string) ProtectionStatusPtrInput {
	return (*protectionStatusPtr)(&v)
}

func (*protectionStatusPtr) ElementType() reflect.Type {
	return protectionStatusPtrType
}

func (in *protectionStatusPtr) ToProtectionStatusPtrOutput() ProtectionStatusPtrOutput {
	return pulumi.ToOutput(in).(ProtectionStatusPtrOutput)
}

func (in *protectionStatusPtr) ToProtectionStatusPtrOutputWithContext(ctx context.Context) ProtectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProtectionStatusPtrOutput)
}

type WorkloadItemType string

const (
	WorkloadItemTypeInvalid         = WorkloadItemType("Invalid")
	WorkloadItemTypeSQLInstance     = WorkloadItemType("SQLInstance")
	WorkloadItemTypeSQLDataBase     = WorkloadItemType("SQLDataBase")
	WorkloadItemTypeSAPHanaSystem   = WorkloadItemType("SAPHanaSystem")
	WorkloadItemTypeSAPHanaDatabase = WorkloadItemType("SAPHanaDatabase")
	WorkloadItemTypeSAPAseSystem    = WorkloadItemType("SAPAseSystem")
	WorkloadItemTypeSAPAseDatabase  = WorkloadItemType("SAPAseDatabase")
)

func (WorkloadItemType) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadItemType)(nil)).Elem()
}

func (e WorkloadItemType) ToWorkloadItemTypeOutput() WorkloadItemTypeOutput {
	return pulumi.ToOutput(e).(WorkloadItemTypeOutput)
}

func (e WorkloadItemType) ToWorkloadItemTypeOutputWithContext(ctx context.Context) WorkloadItemTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WorkloadItemTypeOutput)
}

func (e WorkloadItemType) ToWorkloadItemTypePtrOutput() WorkloadItemTypePtrOutput {
	return e.ToWorkloadItemTypePtrOutputWithContext(context.Background())
}

func (e WorkloadItemType) ToWorkloadItemTypePtrOutputWithContext(ctx context.Context) WorkloadItemTypePtrOutput {
	return WorkloadItemType(e).ToWorkloadItemTypeOutputWithContext(ctx).ToWorkloadItemTypePtrOutputWithContext(ctx)
}

func (e WorkloadItemType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkloadItemType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkloadItemType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WorkloadItemType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WorkloadItemTypeOutput struct{ *pulumi.OutputState }

func (WorkloadItemTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadItemType)(nil)).Elem()
}

func (o WorkloadItemTypeOutput) ToWorkloadItemTypeOutput() WorkloadItemTypeOutput {
	return o
}

func (o WorkloadItemTypeOutput) ToWorkloadItemTypeOutputWithContext(ctx context.Context) WorkloadItemTypeOutput {
	return o
}

func (o WorkloadItemTypeOutput) ToWorkloadItemTypePtrOutput() WorkloadItemTypePtrOutput {
	return o.ToWorkloadItemTypePtrOutputWithContext(context.Background())
}

func (o WorkloadItemTypeOutput) ToWorkloadItemTypePtrOutputWithContext(ctx context.Context) WorkloadItemTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkloadItemType) *WorkloadItemType {
		return &v
	}).(WorkloadItemTypePtrOutput)
}

func (o WorkloadItemTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WorkloadItemTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkloadItemType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WorkloadItemTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkloadItemTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkloadItemType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkloadItemTypePtrOutput struct{ *pulumi.OutputState }

func (WorkloadItemTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadItemType)(nil)).Elem()
}

func (o WorkloadItemTypePtrOutput) ToWorkloadItemTypePtrOutput() WorkloadItemTypePtrOutput {
	return o
}

func (o WorkloadItemTypePtrOutput) ToWorkloadItemTypePtrOutputWithContext(ctx context.Context) WorkloadItemTypePtrOutput {
	return o
}

func (o WorkloadItemTypePtrOutput) Elem() WorkloadItemTypeOutput {
	return o.ApplyT(func(v *WorkloadItemType) WorkloadItemType {
		if v != nil {
			return *v
		}
		var ret WorkloadItemType
		return ret
	}).(WorkloadItemTypeOutput)
}

func (o WorkloadItemTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkloadItemTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WorkloadItemType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WorkloadItemTypeInput interface {
	pulumi.Input

	ToWorkloadItemTypeOutput() WorkloadItemTypeOutput
	ToWorkloadItemTypeOutputWithContext(context.Context) WorkloadItemTypeOutput
}

var workloadItemTypePtrType = reflect.TypeOf((**WorkloadItemType)(nil)).Elem()

type WorkloadItemTypePtrInput interface {
	pulumi.Input

	ToWorkloadItemTypePtrOutput() WorkloadItemTypePtrOutput
	ToWorkloadItemTypePtrOutputWithContext(context.Context) WorkloadItemTypePtrOutput
}

type workloadItemTypePtr string

func WorkloadItemTypePtr(v string) WorkloadItemTypePtrInput {
	return (*workloadItemTypePtr)(&v)
}

func (*workloadItemTypePtr) ElementType() reflect.Type {
	return workloadItemTypePtrType
}

func (in *workloadItemTypePtr) ToWorkloadItemTypePtrOutput() WorkloadItemTypePtrOutput {
	return pulumi.ToOutput(in).(WorkloadItemTypePtrOutput)
}

func (in *workloadItemTypePtr) ToWorkloadItemTypePtrOutputWithContext(ctx context.Context) WorkloadItemTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WorkloadItemTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupManagementTypeInput)(nil)).Elem(), BackupManagementType("Invalid"))
	pulumi.RegisterInputType(reflect.TypeOf((*BackupManagementTypePtrInput)(nil)).Elem(), BackupManagementType("Invalid"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProtectionStatusInput)(nil)).Elem(), ProtectionStatus("Invalid"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProtectionStatusPtrInput)(nil)).Elem(), ProtectionStatus("Invalid"))
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadItemTypeInput)(nil)).Elem(), WorkloadItemType("Invalid"))
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadItemTypePtrInput)(nil)).Elem(), WorkloadItemType("Invalid"))
	pulumi.RegisterOutputType(BackupManagementTypeOutput{})
	pulumi.RegisterOutputType(BackupManagementTypePtrOutput{})
	pulumi.RegisterOutputType(ProtectionStatusOutput{})
	pulumi.RegisterOutputType(ProtectionStatusPtrOutput{})
	pulumi.RegisterOutputType(WorkloadItemTypeOutput{})
	pulumi.RegisterOutputType(WorkloadItemTypePtrOutput{})
}
