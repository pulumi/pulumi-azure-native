


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteBackupConfiguration(ctx *pulumi.Context, args *ListSiteBackupConfigurationArgs, opts ...pulumi.InvokeOption) (*ListSiteBackupConfigurationResult, error) {
	var rv ListSiteBackupConfigurationResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteBackupConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteBackupConfigurationArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListSiteBackupConfigurationResult struct {
	BackupSchedule    *BackupScheduleResponse         `pulumi:"backupSchedule"`
	Databases         []DatabaseBackupSettingResponse `pulumi:"databases"`
	Enabled           *bool                           `pulumi:"enabled"`
	Id                *string                         `pulumi:"id"`
	Kind              *string                         `pulumi:"kind"`
	Location          string                          `pulumi:"location"`
	Name              *string                         `pulumi:"name"`
	StorageAccountUrl *string                         `pulumi:"storageAccountUrl"`
	Tags              map[string]string               `pulumi:"tags"`
	Type              string                          `pulumi:"type"`
}
