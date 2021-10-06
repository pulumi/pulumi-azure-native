


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackupInstance(ctx *pulumi.Context, args *LookupBackupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupBackupInstanceResult, error) {
	var rv LookupBackupInstanceResult
	err := ctx.Invoke("azure-native:dataprotection/v20210601preview:getBackupInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupInstanceArgs struct {
	BackupInstanceName string `pulumi:"backupInstanceName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	VaultName          string `pulumi:"vaultName"`
}


type LookupBackupInstanceResult struct {
	Id         string                 `pulumi:"id"`
	Name       string                 `pulumi:"name"`
	Properties BackupInstanceResponse `pulumi:"properties"`
	SystemData SystemDataResponse     `pulumi:"systemData"`
	Type       string                 `pulumi:"type"`
}
