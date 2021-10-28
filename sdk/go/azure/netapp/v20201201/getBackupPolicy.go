


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackupPolicy(ctx *pulumi.Context, args *LookupBackupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBackupPolicyResult, error) {
	var rv LookupBackupPolicyResult
	err := ctx.Invoke("azure-native:netapp/v20201201:getBackupPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupPolicyArgs struct {
	AccountName       string `pulumi:"accountName"`
	BackupPolicyName  string `pulumi:"backupPolicyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupBackupPolicyResult struct {
	DailyBackupsToKeep   *int                    `pulumi:"dailyBackupsToKeep"`
	Enabled              *bool                   `pulumi:"enabled"`
	Id                   string                  `pulumi:"id"`
	Location             string                  `pulumi:"location"`
	MonthlyBackupsToKeep *int                    `pulumi:"monthlyBackupsToKeep"`
	Name                 string                  `pulumi:"name"`
	ProvisioningState    string                  `pulumi:"provisioningState"`
	Tags                 map[string]string       `pulumi:"tags"`
	Type                 string                  `pulumi:"type"`
	VolumeBackups        []VolumeBackupsResponse `pulumi:"volumeBackups"`
	VolumesAssigned      *int                    `pulumi:"volumesAssigned"`
	WeeklyBackupsToKeep  *int                    `pulumi:"weeklyBackupsToKeep"`
	YearlyBackupsToKeep  *int                    `pulumi:"yearlyBackupsToKeep"`
}
