


package storsimple

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BackupPolicy struct {
	pulumi.CustomResourceState

	BackupPolicyCreationType pulumi.StringOutput      `pulumi:"backupPolicyCreationType"`
	Kind                     pulumi.StringPtrOutput   `pulumi:"kind"`
	LastBackupTime           pulumi.StringOutput      `pulumi:"lastBackupTime"`
	Name                     pulumi.StringOutput      `pulumi:"name"`
	NextBackupTime           pulumi.StringOutput      `pulumi:"nextBackupTime"`
	ScheduledBackupStatus    pulumi.StringOutput      `pulumi:"scheduledBackupStatus"`
	SchedulesCount           pulumi.Float64Output     `pulumi:"schedulesCount"`
	SsmHostName              pulumi.StringOutput      `pulumi:"ssmHostName"`
	Type                     pulumi.StringOutput      `pulumi:"type"`
	VolumeIds                pulumi.StringArrayOutput `pulumi:"volumeIds"`
}


func NewBackupPolicy(ctx *pulumi.Context,
	name string, args *BackupPolicyArgs, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
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
	if args.VolumeIds == nil {
		return nil, errors.New("invalid value for required argument 'VolumeIds'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:storsimple:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:storsimple/v20170601:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:storsimple/v20170601:BackupPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource BackupPolicy
	err := ctx.RegisterResource("azure-native:storsimple:BackupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBackupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupPolicyState, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	var resource BackupPolicy
	err := ctx.ReadResource("azure-native:storsimple:BackupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type backupPolicyState struct {
}

type BackupPolicyState struct {
}

func (BackupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyState)(nil)).Elem()
}

type backupPolicyArgs struct {
	BackupPolicyName  *string  `pulumi:"backupPolicyName"`
	DeviceName        string   `pulumi:"deviceName"`
	Kind              *Kind    `pulumi:"kind"`
	ManagerName       string   `pulumi:"managerName"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	VolumeIds         []string `pulumi:"volumeIds"`
}


type BackupPolicyArgs struct {
	BackupPolicyName  pulumi.StringPtrInput
	DeviceName        pulumi.StringInput
	Kind              KindPtrInput
	ManagerName       pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	VolumeIds         pulumi.StringArrayInput
}

func (BackupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyArgs)(nil)).Elem()
}

type BackupPolicyInput interface {
	pulumi.Input

	ToBackupPolicyOutput() BackupPolicyOutput
	ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput
}

func (*BackupPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicy)(nil))
}

func (i *BackupPolicy) ToBackupPolicyOutput() BackupPolicyOutput {
	return i.ToBackupPolicyOutputWithContext(context.Background())
}

func (i *BackupPolicy) ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyOutput)
}

type BackupPolicyOutput struct{ *pulumi.OutputState }

func (BackupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicy)(nil))
}

func (o BackupPolicyOutput) ToBackupPolicyOutput() BackupPolicyOutput {
	return o
}

func (o BackupPolicyOutput) ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupPolicyInput)(nil)).Elem(), &BackupPolicy{})
	pulumi.RegisterOutputType(BackupPolicyOutput{})
}
