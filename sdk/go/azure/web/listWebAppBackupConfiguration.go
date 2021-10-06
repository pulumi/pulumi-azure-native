


package web

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppBackupConfiguration(ctx *pulumi.Context, args *ListWebAppBackupConfigurationArgs, opts ...pulumi.InvokeOption) (*ListWebAppBackupConfigurationResult, error) {
	var rv ListWebAppBackupConfigurationResult
	err := ctx.Invoke("azure-native:web:listWebAppBackupConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppBackupConfigurationArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppBackupConfigurationResult struct {
	BackupName        *string                         `pulumi:"backupName"`
	BackupSchedule    *BackupScheduleResponse         `pulumi:"backupSchedule"`
	Databases         []DatabaseBackupSettingResponse `pulumi:"databases"`
	Enabled           *bool                           `pulumi:"enabled"`
	Id                string                          `pulumi:"id"`
	Kind              *string                         `pulumi:"kind"`
	Name              string                          `pulumi:"name"`
	StorageAccountUrl string                          `pulumi:"storageAccountUrl"`
	Type              string                          `pulumi:"type"`
}
