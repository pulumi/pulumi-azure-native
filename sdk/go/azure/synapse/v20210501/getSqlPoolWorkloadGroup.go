


package v20210501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSqlPoolWorkloadGroup(ctx *pulumi.Context, args *LookupSqlPoolWorkloadGroupArgs, opts ...pulumi.InvokeOption) (*LookupSqlPoolWorkloadGroupResult, error) {
	var rv LookupSqlPoolWorkloadGroupResult
	err := ctx.Invoke("azure-native:synapse/v20210501:getSqlPoolWorkloadGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlPoolWorkloadGroupArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SqlPoolName       string `pulumi:"sqlPoolName"`
	WorkloadGroupName string `pulumi:"workloadGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupSqlPoolWorkloadGroupResult struct {
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
