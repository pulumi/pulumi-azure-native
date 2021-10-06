


package v20160801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppBackupConfigurationSlot(ctx *pulumi.Context, args *ListWebAppBackupConfigurationSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppBackupConfigurationSlotResult, error) {
	var rv ListWebAppBackupConfigurationSlotResult
	err := ctx.Invoke("azure-native:web/v20160801:listWebAppBackupConfigurationSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppBackupConfigurationSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppBackupConfigurationSlotResult struct {
	BackupRequestName string                          `pulumi:"backupRequestName"`
	BackupSchedule    *BackupScheduleResponse         `pulumi:"backupSchedule"`
	Databases         []DatabaseBackupSettingResponse `pulumi:"databases"`
	Enabled           *bool                           `pulumi:"enabled"`
	Id                string                          `pulumi:"id"`
	Kind              *string                         `pulumi:"kind"`
	Name              string                          `pulumi:"name"`
	StorageAccountUrl string                          `pulumi:"storageAccountUrl"`
	Type              string                          `pulumi:"type"`
}
