


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteBackupStatusSecrets(ctx *pulumi.Context, args *ListSiteBackupStatusSecretsArgs, opts ...pulumi.InvokeOption) (*ListSiteBackupStatusSecretsResult, error) {
	var rv ListSiteBackupStatusSecretsResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteBackupStatusSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteBackupStatusSecretsArgs struct {
	BackupId          string                  `pulumi:"backupId"`
	BackupSchedule    *BackupSchedule         `pulumi:"backupSchedule"`
	Databases         []DatabaseBackupSetting `pulumi:"databases"`
	Enabled           *bool                   `pulumi:"enabled"`
	Id                *string                 `pulumi:"id"`
	Kind              *string                 `pulumi:"kind"`
	Location          *string                 `pulumi:"location"`
	Name              string                  `pulumi:"name"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	StorageAccountUrl *string                 `pulumi:"storageAccountUrl"`
	Tags              map[string]string       `pulumi:"tags"`
	Type              string                  `pulumi:"type"`
}


type ListSiteBackupStatusSecretsResult struct {
	BlobName             *string                         `pulumi:"blobName"`
	CorrelationId        *string                         `pulumi:"correlationId"`
	Created              *string                         `pulumi:"created"`
	Databases            []DatabaseBackupSettingResponse `pulumi:"databases"`
	FinishedTimeStamp    *string                         `pulumi:"finishedTimeStamp"`
	Id                   *string                         `pulumi:"id"`
	Kind                 *string                         `pulumi:"kind"`
	LastRestoreTimeStamp *string                         `pulumi:"lastRestoreTimeStamp"`
	Location             string                          `pulumi:"location"`
	Log                  *string                         `pulumi:"log"`
	Name                 *string                         `pulumi:"name"`
	Scheduled            *bool                           `pulumi:"scheduled"`
	SizeInBytes          *float64                        `pulumi:"sizeInBytes"`
	Status               string                          `pulumi:"status"`
	StorageAccountUrl    *string                         `pulumi:"storageAccountUrl"`
	Tags                 map[string]string               `pulumi:"tags"`
	Type                 *string                         `pulumi:"type"`
	WebsiteSizeInBytes   *float64                        `pulumi:"websiteSizeInBytes"`
}
