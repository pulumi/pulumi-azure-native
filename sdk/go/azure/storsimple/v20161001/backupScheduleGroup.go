


package v20161001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BackupScheduleGroup struct {
	pulumi.CustomResourceState

	Name      pulumi.StringOutput `pulumi:"name"`
	StartTime TimeResponseOutput  `pulumi:"startTime"`
	Type      pulumi.StringOutput `pulumi:"type"`
}


func NewBackupScheduleGroup(ctx *pulumi.Context,
	name string, args *BackupScheduleGroupArgs, opts ...pulumi.ResourceOption) (*BackupScheduleGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	if args.StartTime == nil {
		return nil, errors.New("invalid value for required argument 'StartTime'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:storsimple/v20161001:BackupScheduleGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource BackupScheduleGroup
	err := ctx.RegisterResource("azure-native:storsimple/v20161001:BackupScheduleGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBackupScheduleGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupScheduleGroupState, opts ...pulumi.ResourceOption) (*BackupScheduleGroup, error) {
	var resource BackupScheduleGroup
	err := ctx.ReadResource("azure-native:storsimple/v20161001:BackupScheduleGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type backupScheduleGroupState struct {
}

type BackupScheduleGroupState struct {
}

func (BackupScheduleGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupScheduleGroupState)(nil)).Elem()
}

type backupScheduleGroupArgs struct {
	DeviceName        string  `pulumi:"deviceName"`
	ManagerName       string  `pulumi:"managerName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ScheduleGroupName *string `pulumi:"scheduleGroupName"`
	StartTime         Time    `pulumi:"startTime"`
}


type BackupScheduleGroupArgs struct {
	DeviceName        pulumi.StringInput
	ManagerName       pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ScheduleGroupName pulumi.StringPtrInput
	StartTime         TimeInput
}

func (BackupScheduleGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupScheduleGroupArgs)(nil)).Elem()
}

type BackupScheduleGroupInput interface {
	pulumi.Input

	ToBackupScheduleGroupOutput() BackupScheduleGroupOutput
	ToBackupScheduleGroupOutputWithContext(ctx context.Context) BackupScheduleGroupOutput
}

func (*BackupScheduleGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupScheduleGroup)(nil))
}

func (i *BackupScheduleGroup) ToBackupScheduleGroupOutput() BackupScheduleGroupOutput {
	return i.ToBackupScheduleGroupOutputWithContext(context.Background())
}

func (i *BackupScheduleGroup) ToBackupScheduleGroupOutputWithContext(ctx context.Context) BackupScheduleGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupScheduleGroupOutput)
}

type BackupScheduleGroupOutput struct{ *pulumi.OutputState }

func (BackupScheduleGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupScheduleGroup)(nil))
}

func (o BackupScheduleGroupOutput) ToBackupScheduleGroupOutput() BackupScheduleGroupOutput {
	return o
}

func (o BackupScheduleGroupOutput) ToBackupScheduleGroupOutputWithContext(ctx context.Context) BackupScheduleGroupOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupScheduleGroupInput)(nil)).Elem(), &BackupScheduleGroup{})
	pulumi.RegisterOutputType(BackupScheduleGroupOutput{})
}
