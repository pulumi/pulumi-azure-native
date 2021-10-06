


package v20180907preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabase(ctx *pulumi.Context, args *LookupDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseResult, error) {
	var rv LookupDatabaseResult
	err := ctx.Invoke("azure-native:kusto/v20180907preview:getDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatabaseResult struct {
	Etag                   string                     `pulumi:"etag"`
	HotCachePeriodInDays   *int                       `pulumi:"hotCachePeriodInDays"`
	Id                     string                     `pulumi:"id"`
	Location               string                     `pulumi:"location"`
	Name                   string                     `pulumi:"name"`
	ProvisioningState      string                     `pulumi:"provisioningState"`
	SoftDeletePeriodInDays int                        `pulumi:"softDeletePeriodInDays"`
	Statistics             DatabaseStatisticsResponse `pulumi:"statistics"`
	Tags                   map[string]string          `pulumi:"tags"`
	Type                   string                     `pulumi:"type"`
}
