


package storsimple

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BackupSchedule struct {
	pulumi.CustomResourceState

	BackupType         pulumi.StringOutput              `pulumi:"backupType"`
	Kind               pulumi.StringPtrOutput           `pulumi:"kind"`
	LastSuccessfulRun  pulumi.StringOutput              `pulumi:"lastSuccessfulRun"`
	Name               pulumi.StringOutput              `pulumi:"name"`
	RetentionCount     pulumi.Float64Output             `pulumi:"retentionCount"`
	ScheduleRecurrence ScheduleRecurrenceResponseOutput `pulumi:"scheduleRecurrence"`
	ScheduleStatus     pulumi.StringOutput              `pulumi:"scheduleStatus"`
	StartTime          pulumi.StringOutput              `pulumi:"startTime"`
	Type               pulumi.StringOutput              `pulumi:"type"`
}


func NewBackupSchedule(ctx *pulumi.Context,
	name string, args *BackupScheduleArgs, opts ...pulumi.ResourceOption) (*BackupSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'BackupPolicyName'")
	}
	if args.BackupType == nil {
		return nil, errors.New("invalid value for required argument 'BackupType'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RetentionCount == nil {
		return nil, errors.New("invalid value for required argument 'RetentionCount'")
	}
	if args.ScheduleRecurrence == nil {
		return nil, errors.New("invalid value for required argument 'ScheduleRecurrence'")
	}
	if args.ScheduleStatus == nil {
		return nil, errors.New("invalid value for required argument 'ScheduleStatus'")
	}
	if args.StartTime == nil {
		return nil, errors.New("invalid value for required argument 'StartTime'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storsimple/v20170601:BackupSchedule"),
		},
	})
	opts = append(opts, aliases)
	var resource BackupSchedule
	err := ctx.RegisterResource("azure-native:storsimple:BackupSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBackupSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupScheduleState, opts ...pulumi.ResourceOption) (*BackupSchedule, error) {
	var resource BackupSchedule
	err := ctx.ReadResource("azure-native:storsimple:BackupSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type backupScheduleState struct {
}

type BackupScheduleState struct {
}

func (BackupScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupScheduleState)(nil)).Elem()
}

type backupScheduleArgs struct {
	BackupPolicyName   string             `pulumi:"backupPolicyName"`
	BackupScheduleName *string            `pulumi:"backupScheduleName"`
	BackupType         BackupType         `pulumi:"backupType"`
	DeviceName         string             `pulumi:"deviceName"`
	Kind               *Kind              `pulumi:"kind"`
	ManagerName        string             `pulumi:"managerName"`
	ResourceGroupName  string             `pulumi:"resourceGroupName"`
	RetentionCount     float64            `pulumi:"retentionCount"`
	ScheduleRecurrence ScheduleRecurrence `pulumi:"scheduleRecurrence"`
	ScheduleStatus     ScheduleStatus     `pulumi:"scheduleStatus"`
	StartTime          string             `pulumi:"startTime"`
}


type BackupScheduleArgs struct {
	BackupPolicyName   pulumi.StringInput
	BackupScheduleName pulumi.StringPtrInput
	BackupType         BackupTypeInput
	DeviceName         pulumi.StringInput
	Kind               KindPtrInput
	ManagerName        pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	RetentionCount     pulumi.Float64Input
	ScheduleRecurrence ScheduleRecurrenceInput
	ScheduleStatus     ScheduleStatusInput
	StartTime          pulumi.StringInput
}

func (BackupScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupScheduleArgs)(nil)).Elem()
}

type BackupScheduleInput interface {
	pulumi.Input

	ToBackupScheduleOutput() BackupScheduleOutput
	ToBackupScheduleOutputWithContext(ctx context.Context) BackupScheduleOutput
}

func (*BackupSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupSchedule)(nil)).Elem()
}

func (i *BackupSchedule) ToBackupScheduleOutput() BackupScheduleOutput {
	return i.ToBackupScheduleOutputWithContext(context.Background())
}

func (i *BackupSchedule) ToBackupScheduleOutputWithContext(ctx context.Context) BackupScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupScheduleOutput)
}

type BackupScheduleOutput struct{ *pulumi.OutputState }

func (BackupScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupSchedule)(nil)).Elem()
}

func (o BackupScheduleOutput) ToBackupScheduleOutput() BackupScheduleOutput {
	return o
}

func (o BackupScheduleOutput) ToBackupScheduleOutputWithContext(ctx context.Context) BackupScheduleOutput {
	return o
}

func (o BackupScheduleOutput) BackupType() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupSchedule) pulumi.StringOutput { return v.BackupType }).(pulumi.StringOutput)
}

func (o BackupScheduleOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupSchedule) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o BackupScheduleOutput) LastSuccessfulRun() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupSchedule) pulumi.StringOutput { return v.LastSuccessfulRun }).(pulumi.StringOutput)
}

func (o BackupScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupSchedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BackupScheduleOutput) RetentionCount() pulumi.Float64Output {
	return o.ApplyT(func(v *BackupSchedule) pulumi.Float64Output { return v.RetentionCount }).(pulumi.Float64Output)
}

func (o BackupScheduleOutput) ScheduleRecurrence() ScheduleRecurrenceResponseOutput {
	return o.ApplyT(func(v *BackupSchedule) ScheduleRecurrenceResponseOutput { return v.ScheduleRecurrence }).(ScheduleRecurrenceResponseOutput)
}

func (o BackupScheduleOutput) ScheduleStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupSchedule) pulumi.StringOutput { return v.ScheduleStatus }).(pulumi.StringOutput)
}

func (o BackupScheduleOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupSchedule) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

func (o BackupScheduleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupSchedule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupScheduleOutput{})
}
