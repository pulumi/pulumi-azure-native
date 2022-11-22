


package storsimple

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("azure-native:storsimple:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeArgs struct {
	DeviceName          string `pulumi:"deviceName"`
	ManagerName         string `pulumi:"managerName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	VolumeContainerName string `pulumi:"volumeContainerName"`
	VolumeName          string `pulumi:"volumeName"`
}


type LookupVolumeResult struct {
	AccessControlRecordIds []string `pulumi:"accessControlRecordIds"`
	BackupPolicyIds        []string `pulumi:"backupPolicyIds"`
	BackupStatus           string   `pulumi:"backupStatus"`
	Id                     string   `pulumi:"id"`
	Kind                   *string  `pulumi:"kind"`
	MonitoringStatus       string   `pulumi:"monitoringStatus"`
	Name                   string   `pulumi:"name"`
	OperationStatus        string   `pulumi:"operationStatus"`
	SizeInBytes            float64  `pulumi:"sizeInBytes"`
	Type                   string   `pulumi:"type"`
	VolumeContainerId      string   `pulumi:"volumeContainerId"`
	VolumeStatus           string   `pulumi:"volumeStatus"`
	VolumeType             string   `pulumi:"volumeType"`
}

func LookupVolumeOutput(ctx *pulumi.Context, args LookupVolumeOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeResult, error) {
			args := v.(LookupVolumeArgs)
			r, err := LookupVolume(ctx, &args, opts...)
			var s LookupVolumeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVolumeResultOutput)
}

type LookupVolumeOutputArgs struct {
	DeviceName          pulumi.StringInput `pulumi:"deviceName"`
	ManagerName         pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	VolumeContainerName pulumi.StringInput `pulumi:"volumeContainerName"`
	VolumeName          pulumi.StringInput `pulumi:"volumeName"`
}

func (LookupVolumeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeArgs)(nil)).Elem()
}


type LookupVolumeResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeResult)(nil)).Elem()
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutput() LookupVolumeResultOutput {
	return o
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutputWithContext(ctx context.Context) LookupVolumeResultOutput {
	return o
}

func (o LookupVolumeResultOutput) AccessControlRecordIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVolumeResult) []string { return v.AccessControlRecordIds }).(pulumi.StringArrayOutput)
}

func (o LookupVolumeResultOutput) BackupPolicyIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVolumeResult) []string { return v.BackupPolicyIds }).(pulumi.StringArrayOutput)
}

func (o LookupVolumeResultOutput) BackupStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.BackupStatus }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeResultOutput) MonitoringStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.MonitoringStatus }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) OperationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.OperationStatus }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) SizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVolumeResult) float64 { return v.SizeInBytes }).(pulumi.Float64Output)
}

func (o LookupVolumeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) VolumeContainerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.VolumeContainerId }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) VolumeStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.VolumeStatus }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) VolumeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.VolumeType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeResultOutput{})
}
