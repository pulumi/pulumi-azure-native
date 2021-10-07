


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackupPolicy(ctx *pulumi.Context, args *LookupBackupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBackupPolicyResult, error) {
	var rv LookupBackupPolicyResult
	err := ctx.Invoke("azure-native:dataprotection/v20210601preview:getBackupPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupPolicyArgs struct {
	BackupPolicyName  string `pulumi:"backupPolicyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VaultName         string `pulumi:"vaultName"`
}


type LookupBackupPolicyResult struct {
	Id         string               `pulumi:"id"`
	Name       string               `pulumi:"name"`
	Properties BackupPolicyResponse `pulumi:"properties"`
	SystemData SystemDataResponse   `pulumi:"systemData"`
	Type       string               `pulumi:"type"`
}
