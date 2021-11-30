


package v20201101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkloadGroup(ctx *pulumi.Context, args *LookupWorkloadGroupArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadGroupResult, error) {
	var rv LookupWorkloadGroupResult
	err := ctx.Invoke("azure-native:sql/v20201101preview:getWorkloadGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkloadGroupArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
	WorkloadGroupName string `pulumi:"workloadGroupName"`
}


type LookupWorkloadGroupResult struct {
	Id                           string   `pulumi:"id"`
	Importance                   *string  `pulumi:"importance"`
	MaxResourcePercent           int      `pulumi:"maxResourcePercent"`
	MaxResourcePercentPerRequest *float64 `pulumi:"maxResourcePercentPerRequest"`
	MinResourcePercent           int      `pulumi:"minResourcePercent"`
	MinResourcePercentPerRequest float64  `pulumi:"minResourcePercentPerRequest"`
	Name                         string   `pulumi:"name"`
	QueryExecutionTimeout        *int     `pulumi:"queryExecutionTimeout"`
	Type                         string   `pulumi:"type"`
}
