


package v20170301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
	CreateModeOnlineSecondary                = CreateMode("OnlineSecondary")
	CreateModePointInTimeRestore             = CreateMode("PointInTimeRestore")
	CreateModeRestore                        = CreateMode("Restore")
	CreateModeRecovery                       = CreateMode("Recovery")
	CreateModeRestoreExternalBackup          = CreateMode("RestoreExternalBackup")
	CreateModeRestoreExternalBackupSecondary = CreateMode("RestoreExternalBackupSecondary")
	CreateModeRestoreLongTermRetentionBackup = CreateMode("RestoreLongTermRetentionBackup")
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
	pulumi.RegisterOutputType(BlobAuditingPolicyStateOutput{})
	pulumi.RegisterOutputType(BlobAuditingPolicyStatePtrOutput{})
	pulumi.RegisterOutputType(CatalogCollationTypeOutput{})
	pulumi.RegisterOutputType(CatalogCollationTypePtrOutput{})
	pulumi.RegisterOutputType(CreateModeOutput{})
	pulumi.RegisterOutputType(CreateModePtrOutput{})
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
	pulumi.RegisterOutputType(SampleNameOutput{})
	pulumi.RegisterOutputType(SampleNamePtrOutput{})
	pulumi.RegisterOutputType(SecurityAlertPolicyStateOutput{})
	pulumi.RegisterOutputType(SecurityAlertPolicyStatePtrOutput{})
	pulumi.RegisterOutputType(SensitivityLabelRankOutput{})
	pulumi.RegisterOutputType(SensitivityLabelRankPtrOutput{})
}
