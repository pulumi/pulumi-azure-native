


package sql

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackupShortTermRetentionPolicy(ctx *pulumi.Context, args *LookupBackupShortTermRetentionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBackupShortTermRetentionPolicyResult, error) {
	var rv LookupBackupShortTermRetentionPolicyResult
	err := ctx.Invoke("azure-native:sql:getBackupShortTermRetentionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupShortTermRetentionPolicyArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	PolicyName        string `pulumi:"policyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupBackupShortTermRetentionPolicyResult struct {
	Id            string `pulumi:"id"`
	Name          string `pulumi:"name"`
	RetentionDays *int   `pulumi:"retentionDays"`
	Type          string `pulumi:"type"`
}
