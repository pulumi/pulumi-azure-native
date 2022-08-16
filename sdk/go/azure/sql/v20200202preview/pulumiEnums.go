


package v20200202preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdministratorType string

const (
	AdministratorTypeActiveDirectory = AdministratorType("ActiveDirectory")
)

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

type DatabaseLicenseType string

const (
	DatabaseLicenseTypeLicenseIncluded = DatabaseLicenseType("LicenseIncluded")
	DatabaseLicenseTypeBasePrice       = DatabaseLicenseType("BasePrice")
)

type DatabaseReadScale string

const (
	DatabaseReadScaleEnabled  = DatabaseReadScale("Enabled")
	DatabaseReadScaleDisabled = DatabaseReadScale("Disabled")
)

type ElasticPoolLicenseType string

const (
	ElasticPoolLicenseTypeLicenseIncluded = ElasticPoolLicenseType("LicenseIncluded")
	ElasticPoolLicenseTypeBasePrice       = ElasticPoolLicenseType("BasePrice")
)

type IdentityType string

const (
	IdentityTypeNone                         = IdentityType("None")
	IdentityTypeSystemAssigned               = IdentityType("SystemAssigned")
	IdentityTypeUserAssigned                 = IdentityType("UserAssigned")
	IdentityType_SystemAssigned_UserAssigned = IdentityType("SystemAssigned,UserAssigned")
)

type InstancePoolLicenseType string

const (
	InstancePoolLicenseTypeLicenseIncluded = InstancePoolLicenseType("LicenseIncluded")
	InstancePoolLicenseTypeBasePrice       = InstancePoolLicenseType("BasePrice")
)

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

type JobStepActionType string

const (
	JobStepActionTypeTSql = JobStepActionType("TSql")
)

type JobStepOutputTypeEnum string

const (
	JobStepOutputTypeEnumSqlDatabase = JobStepOutputTypeEnum("SqlDatabase")
)

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

type ManagedDatabaseCreateMode string

const (
	ManagedDatabaseCreateModeDefault                        = ManagedDatabaseCreateMode("Default")
	ManagedDatabaseCreateModeRestoreExternalBackup          = ManagedDatabaseCreateMode("RestoreExternalBackup")
	ManagedDatabaseCreateModePointInTimeRestore             = ManagedDatabaseCreateMode("PointInTimeRestore")
	ManagedDatabaseCreateModeRecovery                       = ManagedDatabaseCreateMode("Recovery")
	ManagedDatabaseCreateModeRestoreLongTermRetentionBackup = ManagedDatabaseCreateMode("RestoreLongTermRetentionBackup")
)

type ManagedInstanceAdministratorType string

const (
	ManagedInstanceAdministratorTypeActiveDirectory = ManagedInstanceAdministratorType("ActiveDirectory")
)

type ManagedInstanceLicenseType string

const (
	ManagedInstanceLicenseTypeLicenseIncluded = ManagedInstanceLicenseType("LicenseIncluded")
	ManagedInstanceLicenseTypeBasePrice       = ManagedInstanceLicenseType("BasePrice")
)

type ManagedInstanceProxyOverride string

const (
	ManagedInstanceProxyOverrideProxy    = ManagedInstanceProxyOverride("Proxy")
	ManagedInstanceProxyOverrideRedirect = ManagedInstanceProxyOverride("Redirect")
	ManagedInstanceProxyOverrideDefault  = ManagedInstanceProxyOverride("Default")
)

type ManagedServerCreateMode string

const (
	ManagedServerCreateModeDefault            = ManagedServerCreateMode("Default")
	ManagedServerCreateModePointInTimeRestore = ManagedServerCreateMode("PointInTimeRestore")
)

type PrivateLinkServiceConnectionStateStatus string

const (
	PrivateLinkServiceConnectionStateStatusApproved     = PrivateLinkServiceConnectionStateStatus("Approved")
	PrivateLinkServiceConnectionStateStatusPending      = PrivateLinkServiceConnectionStateStatus("Pending")
	PrivateLinkServiceConnectionStateStatusRejected     = PrivateLinkServiceConnectionStateStatus("Rejected")
	PrivateLinkServiceConnectionStateStatusDisconnected = PrivateLinkServiceConnectionStateStatus("Disconnected")
)

type ReadOnlyEndpointFailoverPolicy string

const (
	ReadOnlyEndpointFailoverPolicyDisabled = ReadOnlyEndpointFailoverPolicy("Disabled")
	ReadOnlyEndpointFailoverPolicyEnabled  = ReadOnlyEndpointFailoverPolicy("Enabled")
)

type ReadWriteEndpointFailoverPolicy string

const (
	ReadWriteEndpointFailoverPolicyManual    = ReadWriteEndpointFailoverPolicy("Manual")
	ReadWriteEndpointFailoverPolicyAutomatic = ReadWriteEndpointFailoverPolicy("Automatic")
)

type SampleName string

const (
	SampleNameAdventureWorksLT       = SampleName("AdventureWorksLT")
	SampleNameWideWorldImportersStd  = SampleName("WideWorldImportersStd")
	SampleNameWideWorldImportersFull = SampleName("WideWorldImportersFull")
)

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

type ServerPublicNetworkAccess string

const (
	ServerPublicNetworkAccessEnabled  = ServerPublicNetworkAccess("Enabled")
	ServerPublicNetworkAccessDisabled = ServerPublicNetworkAccess("Disabled")
)

type StorageAccountType string

const (
	StorageAccountTypeGRS = StorageAccountType("GRS")
	StorageAccountTypeLRS = StorageAccountType("LRS")
	StorageAccountTypeZRS = StorageAccountType("ZRS")
)

type SyncConflictResolutionPolicy string

const (
	SyncConflictResolutionPolicyHubWin    = SyncConflictResolutionPolicy("HubWin")
	SyncConflictResolutionPolicyMemberWin = SyncConflictResolutionPolicy("MemberWin")
)

type SyncDirection string

const (
	SyncDirectionBidirectional     = SyncDirection("Bidirectional")
	SyncDirectionOneWayMemberToHub = SyncDirection("OneWayMemberToHub")
	SyncDirectionOneWayHubToMember = SyncDirection("OneWayHubToMember")
)

type SyncMemberDbType string

const (
	SyncMemberDbTypeAzureSqlDatabase  = SyncMemberDbType("AzureSqlDatabase")
	SyncMemberDbTypeSqlServerDatabase = SyncMemberDbType("SqlServerDatabase")
)

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
	pulumi.RegisterOutputType(AutoExecuteStatusOutput{})
	pulumi.RegisterOutputType(AutoExecuteStatusPtrOutput{})
	pulumi.RegisterOutputType(BlobAuditingPolicyStateOutput{})
	pulumi.RegisterOutputType(BlobAuditingPolicyStatePtrOutput{})
	pulumi.RegisterOutputType(JobScheduleTypeOutput{})
	pulumi.RegisterOutputType(JobScheduleTypePtrOutput{})
	pulumi.RegisterOutputType(JobTargetGroupMembershipTypeOutput{})
	pulumi.RegisterOutputType(JobTargetGroupMembershipTypePtrOutput{})
	pulumi.RegisterOutputType(SecurityAlertsPolicyStateOutput{})
	pulumi.RegisterOutputType(SecurityAlertsPolicyStatePtrOutput{})
	pulumi.RegisterOutputType(SensitivityLabelRankOutput{})
	pulumi.RegisterOutputType(SensitivityLabelRankPtrOutput{})
	pulumi.RegisterOutputType(TransparentDataEncryptionStateEnumOutput{})
	pulumi.RegisterOutputType(TransparentDataEncryptionStateEnumPtrOutput{})
}
