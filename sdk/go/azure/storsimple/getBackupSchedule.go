


package storsimple

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackupSchedule(ctx *pulumi.Context, args *LookupBackupScheduleArgs, opts ...pulumi.InvokeOption) (*LookupBackupScheduleResult, error) {
	var rv LookupBackupScheduleResult
	err := ctx.Invoke("azure-native:storsimple:getBackupSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupScheduleArgs struct {
	BackupPolicyName   string `pulumi:"backupPolicyName"`
	BackupScheduleName string `pulumi:"backupScheduleName"`
	DeviceName         string `pulumi:"deviceName"`
	ManagerName        string `pulumi:"managerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupBackupScheduleResult struct {
	BackupType         string                     `pulumi:"backupType"`
	Id                 string                     `pulumi:"id"`
	Kind               *string                    `pulumi:"kind"`
	LastSuccessfulRun  string                     `pulumi:"lastSuccessfulRun"`
	Name               string                     `pulumi:"name"`
	RetentionCount     float64                    `pulumi:"retentionCount"`
	ScheduleRecurrence ScheduleRecurrenceResponse `pulumi:"scheduleRecurrence"`
	ScheduleStatus     string                     `pulumi:"scheduleStatus"`
	StartTime          string                     `pulumi:"startTime"`
	Type               string                     `pulumi:"type"`
}

func LookupBackupScheduleOutput(ctx *pulumi.Context, args LookupBackupScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupBackupScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupScheduleResult, error) {
			args := v.(LookupBackupScheduleArgs)
			r, err := LookupBackupSchedule(ctx, &args, opts...)
			var s LookupBackupScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupScheduleResultOutput)
}

type LookupBackupScheduleOutputArgs struct {
	BackupPolicyName   pulumi.StringInput `pulumi:"backupPolicyName"`
	BackupScheduleName pulumi.StringInput `pulumi:"backupScheduleName"`
	DeviceName         pulumi.StringInput `pulumi:"deviceName"`
	ManagerName        pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBackupScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupScheduleArgs)(nil)).Elem()
}


type LookupBackupScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupBackupScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupScheduleResult)(nil)).Elem()
}

func (o LookupBackupScheduleResultOutput) ToLookupBackupScheduleResultOutput() LookupBackupScheduleResultOutput {
	return o
}

func (o LookupBackupScheduleResultOutput) ToLookupBackupScheduleResultOutputWithContext(ctx context.Context) LookupBackupScheduleResultOutput {
	return o
}

func (o LookupBackupScheduleResultOutput) BackupType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupScheduleResult) string { return v.BackupType }).(pulumi.StringOutput)
}

func (o LookupBackupScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBackupScheduleResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackupScheduleResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupBackupScheduleResultOutput) LastSuccessfulRun() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupScheduleResult) string { return v.LastSuccessfulRun }).(pulumi.StringOutput)
}

func (o LookupBackupScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBackupScheduleResultOutput) RetentionCount() pulumi.Float64Output {
	return o.ApplyT(func(v LookupBackupScheduleResult) float64 { return v.RetentionCount }).(pulumi.Float64Output)
}

func (o LookupBackupScheduleResultOutput) ScheduleRecurrence() ScheduleRecurrenceResponseOutput {
	return o.ApplyT(func(v LookupBackupScheduleResult) ScheduleRecurrenceResponse { return v.ScheduleRecurrence }).(ScheduleRecurrenceResponseOutput)
}

func (o LookupBackupScheduleResultOutput) ScheduleStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupScheduleResult) string { return v.ScheduleStatus }).(pulumi.StringOutput)
}

func (o LookupBackupScheduleResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupScheduleResult) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o LookupBackupScheduleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupScheduleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupScheduleResultOutput{})
}
