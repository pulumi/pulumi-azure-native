


package v20200614

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReadWriteDatabase(ctx *pulumi.Context, args *LookupReadWriteDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupReadWriteDatabaseResult, error) {
	var rv LookupReadWriteDatabaseResult
	err := ctx.Invoke("azure-native:kusto/v20200614:getReadWriteDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReadWriteDatabaseArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupReadWriteDatabaseResult struct {
	HotCachePeriod    *string                    `pulumi:"hotCachePeriod"`
	Id                string                     `pulumi:"id"`
	IsFollowed        bool                       `pulumi:"isFollowed"`
	Kind              string                     `pulumi:"kind"`
	Location          *string                    `pulumi:"location"`
	Name              string                     `pulumi:"name"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	SoftDeletePeriod  *string                    `pulumi:"softDeletePeriod"`
	Statistics        DatabaseStatisticsResponse `pulumi:"statistics"`
	Type              string                     `pulumi:"type"`
}
