


package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackupVault(ctx *pulumi.Context, args *LookupBackupVaultArgs, opts ...pulumi.InvokeOption) (*LookupBackupVaultResult, error) {
	var rv LookupBackupVaultResult
	err := ctx.Invoke("azure-native:dataprotection/v20210101:getBackupVault", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupVaultArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VaultName         string `pulumi:"vaultName"`
}


type LookupBackupVaultResult struct {
	ETag       *string                     `pulumi:"eTag"`
	Id         string                      `pulumi:"id"`
	Identity   *DppIdentityDetailsResponse `pulumi:"identity"`
	Location   string                      `pulumi:"location"`
	Name       string                      `pulumi:"name"`
	Properties BackupVaultResponse         `pulumi:"properties"`
	SystemData SystemDataResponse          `pulumi:"systemData"`
	Tags       map[string]string           `pulumi:"tags"`
	Type       string                      `pulumi:"type"`
}
