


package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppBackupStatusSecretsSlot(ctx *pulumi.Context, args *ListWebAppBackupStatusSecretsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppBackupStatusSecretsSlotResult, error) {
	var rv ListWebAppBackupStatusSecretsSlotResult
	err := ctx.Invoke("azure-native:web/v20200601:listWebAppBackupStatusSecretsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppBackupStatusSecretsSlotArgs struct {
	BackupId          string                  `pulumi:"backupId"`
	BackupName        *string                 `pulumi:"backupName"`
	BackupSchedule    *BackupSchedule         `pulumi:"backupSchedule"`
	Databases         []DatabaseBackupSetting `pulumi:"databases"`
	Enabled           *bool                   `pulumi:"enabled"`
	Kind              *string                 `pulumi:"kind"`
	Name              string                  `pulumi:"name"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	Slot              string                  `pulumi:"slot"`
	StorageAccountUrl string                  `pulumi:"storageAccountUrl"`
}


type ListWebAppBackupStatusSecretsSlotResult struct {
	BackupId             int                             `pulumi:"backupId"`
	BlobName             string                          `pulumi:"blobName"`
	CorrelationId        string                          `pulumi:"correlationId"`
	Created              string                          `pulumi:"created"`
	Databases            []DatabaseBackupSettingResponse `pulumi:"databases"`
	FinishedTimeStamp    string                          `pulumi:"finishedTimeStamp"`
	Id                   string                          `pulumi:"id"`
	Kind                 *string                         `pulumi:"kind"`
	LastRestoreTimeStamp string                          `pulumi:"lastRestoreTimeStamp"`
	Log                  string                          `pulumi:"log"`
	Name                 string                          `pulumi:"name"`
	Scheduled            bool                            `pulumi:"scheduled"`
	SizeInBytes          float64                         `pulumi:"sizeInBytes"`
	Status               string                          `pulumi:"status"`
	StorageAccountUrl    string                          `pulumi:"storageAccountUrl"`
	Type                 string                          `pulumi:"type"`
	WebsiteSizeInBytes   float64                         `pulumi:"websiteSizeInBytes"`
}
