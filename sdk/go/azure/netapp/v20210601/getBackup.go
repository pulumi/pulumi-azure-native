


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackup(ctx *pulumi.Context, args *LookupBackupArgs, opts ...pulumi.InvokeOption) (*LookupBackupResult, error) {
	var rv LookupBackupResult
	err := ctx.Invoke("azure-native:netapp/v20210601:getBackup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupBackupArgs struct {
	AccountName       string `pulumi:"accountName"`
	BackupName        string `pulumi:"backupName"`
	PoolName          string `pulumi:"poolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VolumeName        string `pulumi:"volumeName"`
}


type LookupBackupResult struct {
	BackupId            string  `pulumi:"backupId"`
	BackupType          string  `pulumi:"backupType"`
	CreationDate        string  `pulumi:"creationDate"`
	FailureReason       string  `pulumi:"failureReason"`
	Id                  string  `pulumi:"id"`
	Label               *string `pulumi:"label"`
	Location            string  `pulumi:"location"`
	Name                string  `pulumi:"name"`
	ProvisioningState   string  `pulumi:"provisioningState"`
	Size                float64 `pulumi:"size"`
	Type                string  `pulumi:"type"`
	UseExistingSnapshot *bool   `pulumi:"useExistingSnapshot"`
	VolumeName          string  `pulumi:"volumeName"`
}


func (val *LookupBackupResult) Defaults() *LookupBackupResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseExistingSnapshot) {
		useExistingSnapshot_ := false
		tmp.UseExistingSnapshot = &useExistingSnapshot_
	}
	return &tmp
}
