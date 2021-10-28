


package v20170601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackupPolicy(ctx *pulumi.Context, args *LookupBackupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBackupPolicyResult, error) {
	var rv LookupBackupPolicyResult
	err := ctx.Invoke("azure-native:storsimple/v20170601:getBackupPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupPolicyArgs struct {
	BackupPolicyName  string `pulumi:"backupPolicyName"`
	DeviceName        string `pulumi:"deviceName"`
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBackupPolicyResult struct {
	BackupPolicyCreationType string   `pulumi:"backupPolicyCreationType"`
	Id                       string   `pulumi:"id"`
	Kind                     *string  `pulumi:"kind"`
	LastBackupTime           string   `pulumi:"lastBackupTime"`
	Name                     string   `pulumi:"name"`
	NextBackupTime           string   `pulumi:"nextBackupTime"`
	ScheduledBackupStatus    string   `pulumi:"scheduledBackupStatus"`
	SchedulesCount           float64  `pulumi:"schedulesCount"`
	SsmHostName              string   `pulumi:"ssmHostName"`
	Type                     string   `pulumi:"type"`
	VolumeIds                []string `pulumi:"volumeIds"`
}
