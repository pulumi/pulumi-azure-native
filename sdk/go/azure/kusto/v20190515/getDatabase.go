


package v20190515

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabase(ctx *pulumi.Context, args *LookupDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseResult, error) {
	var rv LookupDatabaseResult
	err := ctx.Invoke("azure-native:kusto/v20190515:getDatabase", args, &rv, opts...)
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
	HotCachePeriod    *string                    `pulumi:"hotCachePeriod"`
	Id                string                     `pulumi:"id"`
	Location          *string                    `pulumi:"location"`
	Name              string                     `pulumi:"name"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	SoftDeletePeriod  *string                    `pulumi:"softDeletePeriod"`
	Statistics        DatabaseStatisticsResponse `pulumi:"statistics"`
	Type              string                     `pulumi:"type"`
}
