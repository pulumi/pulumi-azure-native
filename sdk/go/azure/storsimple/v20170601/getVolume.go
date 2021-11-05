


package v20170601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("azure-native:storsimple/v20170601:getVolume", args, &rv, opts...)
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
