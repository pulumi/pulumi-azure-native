


package storsimple

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Volume struct {
	pulumi.CustomResourceState

	AccessControlRecordIds pulumi.StringArrayOutput `pulumi:"accessControlRecordIds"`
	BackupPolicyIds        pulumi.StringArrayOutput `pulumi:"backupPolicyIds"`
	BackupStatus           pulumi.StringOutput      `pulumi:"backupStatus"`
	Kind                   pulumi.StringPtrOutput   `pulumi:"kind"`
	MonitoringStatus       pulumi.StringOutput      `pulumi:"monitoringStatus"`
	Name                   pulumi.StringOutput      `pulumi:"name"`
	OperationStatus        pulumi.StringOutput      `pulumi:"operationStatus"`
	SizeInBytes            pulumi.Float64Output     `pulumi:"sizeInBytes"`
	Type                   pulumi.StringOutput      `pulumi:"type"`
	VolumeContainerId      pulumi.StringOutput      `pulumi:"volumeContainerId"`
	VolumeStatus           pulumi.StringOutput      `pulumi:"volumeStatus"`
	VolumeType             pulumi.StringOutput      `pulumi:"volumeType"`
}


func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessControlRecordIds == nil {
		return nil, errors.New("invalid value for required argument 'AccessControlRecordIds'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.MonitoringStatus == nil {
		return nil, errors.New("invalid value for required argument 'MonitoringStatus'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SizeInBytes == nil {
		return nil, errors.New("invalid value for required argument 'SizeInBytes'")
	}
	if args.VolumeContainerName == nil {
		return nil, errors.New("invalid value for required argument 'VolumeContainerName'")
	}
	if args.VolumeStatus == nil {
		return nil, errors.New("invalid value for required argument 'VolumeStatus'")
	}
	if args.VolumeType == nil {
		return nil, errors.New("invalid value for required argument 'VolumeType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storsimple/v20170601:Volume"),
		},
	})
	opts = append(opts, aliases)
	var resource Volume
	err := ctx.RegisterResource("azure-native:storsimple:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("azure-native:storsimple:Volume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type volumeState struct {
}

type VolumeState struct {
}

func (VolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeState)(nil)).Elem()
}

type volumeArgs struct {
	AccessControlRecordIds []string         `pulumi:"accessControlRecordIds"`
	DeviceName             string           `pulumi:"deviceName"`
	Kind                   *Kind            `pulumi:"kind"`
	ManagerName            string           `pulumi:"managerName"`
	MonitoringStatus       MonitoringStatus `pulumi:"monitoringStatus"`
	ResourceGroupName      string           `pulumi:"resourceGroupName"`
	SizeInBytes            float64          `pulumi:"sizeInBytes"`
	VolumeContainerName    string           `pulumi:"volumeContainerName"`
	VolumeName             *string          `pulumi:"volumeName"`
	VolumeStatus           VolumeStatus     `pulumi:"volumeStatus"`
	VolumeType             VolumeType       `pulumi:"volumeType"`
}


type VolumeArgs struct {
	AccessControlRecordIds pulumi.StringArrayInput
	DeviceName             pulumi.StringInput
	Kind                   KindPtrInput
	ManagerName            pulumi.StringInput
	MonitoringStatus       MonitoringStatusInput
	ResourceGroupName      pulumi.StringInput
	SizeInBytes            pulumi.Float64Input
	VolumeContainerName    pulumi.StringInput
	VolumeName             pulumi.StringPtrInput
	VolumeStatus           VolumeStatusInput
	VolumeType             VolumeTypeInput
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeArgs)(nil)).Elem()
}

type VolumeInput interface {
	pulumi.Input

	ToVolumeOutput() VolumeOutput
	ToVolumeOutputWithContext(ctx context.Context) VolumeOutput
}

func (*Volume) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (i *Volume) ToVolumeOutput() VolumeOutput {
	return i.ToVolumeOutputWithContext(context.Background())
}

func (i *Volume) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeOutput)
}

type VolumeOutput struct{ *pulumi.OutputState }

func (VolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (o VolumeOutput) ToVolumeOutput() VolumeOutput {
	return o
}

func (o VolumeOutput) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return o
}

func (o VolumeOutput) AccessControlRecordIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringArrayOutput { return v.AccessControlRecordIds }).(pulumi.StringArrayOutput)
}

func (o VolumeOutput) BackupPolicyIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringArrayOutput { return v.BackupPolicyIds }).(pulumi.StringArrayOutput)
}

func (o VolumeOutput) BackupStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.BackupStatus }).(pulumi.StringOutput)
}

func (o VolumeOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) MonitoringStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.MonitoringStatus }).(pulumi.StringOutput)
}

func (o VolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VolumeOutput) OperationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.OperationStatus }).(pulumi.StringOutput)
}

func (o VolumeOutput) SizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *Volume) pulumi.Float64Output { return v.SizeInBytes }).(pulumi.Float64Output)
}

func (o VolumeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VolumeOutput) VolumeContainerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.VolumeContainerId }).(pulumi.StringOutput)
}

func (o VolumeOutput) VolumeStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.VolumeStatus }).(pulumi.StringOutput)
}

func (o VolumeOutput) VolumeType() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.VolumeType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeOutput{})
}
