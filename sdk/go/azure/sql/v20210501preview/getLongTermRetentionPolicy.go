


package v20210501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLongTermRetentionPolicy(ctx *pulumi.Context, args *LookupLongTermRetentionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupLongTermRetentionPolicyResult, error) {
	var rv LookupLongTermRetentionPolicyResult
	err := ctx.Invoke("azure-native:sql/v20210501preview:getLongTermRetentionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLongTermRetentionPolicyArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	PolicyName        string `pulumi:"policyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupLongTermRetentionPolicyResult struct {
	Id               string  `pulumi:"id"`
	MonthlyRetention *string `pulumi:"monthlyRetention"`
	Name             string  `pulumi:"name"`
	Type             string  `pulumi:"type"`
	WeekOfYear       *int    `pulumi:"weekOfYear"`
	WeeklyRetention  *string `pulumi:"weeklyRetention"`
	YearlyRetention  *string `pulumi:"yearlyRetention"`
}
