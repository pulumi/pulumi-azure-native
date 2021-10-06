


package v20170301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBackupLongTermRetentionPolicy(ctx *pulumi.Context, args *LookupBackupLongTermRetentionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBackupLongTermRetentionPolicyResult, error) {
	var rv LookupBackupLongTermRetentionPolicyResult
	err := ctx.Invoke("azure-native:sql/v20170301preview:getBackupLongTermRetentionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupLongTermRetentionPolicyArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	PolicyName        string `pulumi:"policyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupBackupLongTermRetentionPolicyResult struct {
	Id               string  `pulumi:"id"`
	MonthlyRetention *string `pulumi:"monthlyRetention"`
	Name             string  `pulumi:"name"`
	Type             string  `pulumi:"type"`
	WeekOfYear       *int    `pulumi:"weekOfYear"`
	WeeklyRetention  *string `pulumi:"weeklyRetention"`
	YearlyRetention  *string `pulumi:"yearlyRetention"`
}
